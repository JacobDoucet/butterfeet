package registry_item_api

import (
	"context"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item_mongo"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMongoBackedClient(db *mongo.Database, hooks ...Hooks) Client {
	return New(&mongoClient{db: db}, hooks...)
}

type mongoClient struct {
	db *mongo.Database
}

func (m *mongoClient) Search(ctx context.Context, where WhereClause, options QueryOptions) (QueryResult, error) {
	projection := options.GetProjection()
	mongoWhereClause, err := where.RegistryItem.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}
	mongoWhereClauseReservations, err := where.Reservations.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}
	mongoWhereClauseRegistry, err := where.Registry.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}

	searchResult, err := registry_item_mongo.Search(
		ctx,
		m.db,
		registry_item_mongo.WhereClause{
			RegistryItem: mongoWhereClause,
			Reservations: mongoWhereClauseReservations,
			Registry:     mongoWhereClauseRegistry,
		},
		registry_item_mongo.LookupOptions{
			Projection:             projection.Projection,
			Sort:                   options.Sort.ToMongoSortParams(),
			ReservationsProjection: projection.Reservations,
			RegistryProjection:     projection.Registry,
			Limit:                  options.Limit,
			Skip:                   options.Skip,
		},
	)
	if err != nil {
		return QueryResult{}, err
	}

	modelRecords, err := FromMongoQueryResultDataList(searchResult.Data)
	if err != nil {
		return QueryResult{}, err
	}

	return QueryResult{
		Data:  modelRecords,
		Total: searchResult.Count,
		Skip:  options.Skip,
	}, nil
}

func (m *mongoClient) Create(ctx context.Context, obj registry_item.Model, projection registry_item.Projection) (registry_item.Model, error) {

	createRecord, err := obj.ToMongoRecord(projection)
	if err != nil {
		return registry_item.Model{}, err
	}
	var id primitive.ObjectID
	id, err = registry_item_mongo.Create(ctx, m.db, createRecord)
	if err != nil {
		return registry_item.Model{}, err
	}
	createRecord.Id = &id
	return createRecord.ToModel()
}

func (m *mongoClient) Update(ctx context.Context, obj registry_item.Model, where registry_item.WhereClause, projection registry_item.Projection) (registry_item.Model, error) {
	mongoWhereClause, err := where.ToMongoWhereClause()
	if err != nil {
		return registry_item.Model{}, err
	}

	updateRecord, err := obj.ToMongoRecord(projection)
	if err != nil {
		return registry_item.Model{}, err
	}

	err = registry_item_mongo.Update(ctx, m.db, updateRecord, mongoWhereClause)
	if err != nil {
		return registry_item.Model{}, err
	}

	return updateRecord.ToModel()
}

func (m *mongoClient) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Join(errors.New("invalid id"), err)
	}
	err = registry_item_mongo.Delete(ctx, m.db, oid)
	if err != nil {
		return err
	}
	return nil
}

func FromMongoQueryResultData(r registry_item_mongo.Model) (Model, error) {
	m := Model{}
	var err error
	m.Model, err = r.ToModel()
	if r.Reservations != nil {
		val := make([]reservation.Model, 0)
		var err error
		for _, rr := range *r.Reservations {
			nextVal, nextErr := rr.ToModel()
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}
		m.Reservations = &val
	}
	if r.Registry != nil {
		val, toModelErr := r.Registry.ToModel()
		if toModelErr != nil {
			err = errors.Join(err, toModelErr)
		}
		m.Registry = &val
	}
	return m, err
}

func FromMongoQueryResultDataList(dbRecords []registry_item_mongo.Model) ([]Model, error) {
	ms := make([]Model, len(dbRecords))
	var err error
	for i, r := range dbRecords {
		var iErr error
		ms[i], iErr = FromMongoQueryResultData(r)
		if iErr != nil {
			err = errors.Join(err, iErr)
		}
	}
	return ms, err
}

func (m *mongoClient) Aggregate(ctx context.Context, where WhereClause, options AggregateOptions) (AggregateResult, error) {
	mongoWhereClause, err := where.RegistryItem.ToMongoWhereClause()
	if err != nil {
		return AggregateResult{}, err
	}

	// Convert API aggregate options to mongo aggregate options
	mongoAggFields := make([]registry_item_mongo.AggregateFieldSpec, len(options.Fields))
	for i, f := range options.Fields {
		mongoAggFields[i] = registry_item_mongo.AggregateFieldSpec{
			Field:  string(f.Field),
			Method: registry_item_mongo.AggregateMethod(f.Method),
			Alias:  f.Alias,
		}
	}

	// Convert group-by fields to strings
	mongoGroupBy := make([]string, len(options.GroupBy))
	for i, g := range options.GroupBy {
		mongoGroupBy[i] = string(g)
	}

	mongoAggOptions := registry_item_mongo.AggregateOptions{
		Fields:                 mongoAggFields,
		GroupBy:                mongoGroupBy,
		ReservationsProjection: options.ReservationsProjection,
		RegistryProjection:     options.RegistryProjection,
	}

	result, err := registry_item_mongo.Aggregate(
		ctx,
		m.db,
		registry_item_mongo.WhereClause{
			RegistryItem: mongoWhereClause,
		},
		mongoAggOptions,
	)
	if err != nil {
		return AggregateResult{}, err
	}

	// Convert mongo result to API result
	apiResults := make([]AggregateResultRow, len(result.Data))
	for i, r := range result.Data {
		row := AggregateResultRow{
			// Copy metadata
			GroupKeys:     r.GroupKeys,
			AggregateKeys: r.AggregateKeys,
		}
		// Copy group-by fields (with type conversion for refs)
		row.Currency = r.Currency
		row.Description = r.Description
		row.ImageUrl = r.ImageUrl
		row.Notes = r.Notes
		row.Position = r.Position
		row.PriceCents = r.PriceCents
		row.ProductUrl = r.ProductUrl
		row.Quantity = r.Quantity
		if r.RegistryId != nil {
			s := r.RegistryId.Hex()
			row.RegistryId = &s
		}
		row.Title = r.Title
		// Copy aggregate fields (only those not in group-by)
		// Copy ref field Registry
		if r.Registry != nil {
			val, toModelErr := r.Registry.ToModel()
			if toModelErr != nil {
				err = errors.Join(err, toModelErr)
			}
			row.Registry = &val
		}
		// Copy ref field Reservations
		if r.Reservations != nil {
			val := make([]reservation.Model, 0)
			for _, rr := range r.Reservations {
				nextVal, nextErr := rr.ToModel()
				if nextErr != nil {
					err = errors.Join(err, nextErr)
				}
				val = append(val, nextVal)
			}
			row.Reservations = val
		}
		apiResults[i] = row
	}

	return AggregateResult{
		Data:  apiResults,
		Total: result.Total,
	}, nil
}
