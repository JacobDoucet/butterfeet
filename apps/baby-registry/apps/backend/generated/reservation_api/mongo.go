package reservation_api

import (
	"context"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation_mongo"
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
	mongoWhereClause, err := where.Reservation.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}
	mongoWhereClauseItem, err := where.Item.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}
	mongoWhereClauseRegistry, err := where.Registry.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}

	searchResult, err := reservation_mongo.Search(
		ctx,
		m.db,
		reservation_mongo.WhereClause{
			Reservation: mongoWhereClause,
			Item:        mongoWhereClauseItem,
			Registry:    mongoWhereClauseRegistry,
		},
		reservation_mongo.LookupOptions{
			Projection:         projection.Projection,
			Sort:               options.Sort.ToMongoSortParams(),
			ItemProjection:     projection.Item,
			RegistryProjection: projection.Registry,
			Limit:              options.Limit,
			Skip:               options.Skip,
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

func (m *mongoClient) Create(ctx context.Context, obj reservation.Model, projection reservation.Projection) (reservation.Model, error) {

	createRecord, err := obj.ToMongoRecord(projection)
	if err != nil {
		return reservation.Model{}, err
	}
	var id primitive.ObjectID
	id, err = reservation_mongo.Create(ctx, m.db, createRecord)
	if err != nil {
		return reservation.Model{}, err
	}
	createRecord.Id = &id
	return createRecord.ToModel()
}

func (m *mongoClient) Update(ctx context.Context, obj reservation.Model, where reservation.WhereClause, projection reservation.Projection) (reservation.Model, error) {
	mongoWhereClause, err := where.ToMongoWhereClause()
	if err != nil {
		return reservation.Model{}, err
	}

	updateRecord, err := obj.ToMongoRecord(projection)
	if err != nil {
		return reservation.Model{}, err
	}

	err = reservation_mongo.Update(ctx, m.db, updateRecord, mongoWhereClause)
	if err != nil {
		return reservation.Model{}, err
	}

	return updateRecord.ToModel()
}

func (m *mongoClient) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Join(errors.New("invalid id"), err)
	}
	err = reservation_mongo.Delete(ctx, m.db, oid)
	if err != nil {
		return err
	}
	return nil
}

func FromMongoQueryResultData(r reservation_mongo.Model) (Model, error) {
	m := Model{}
	var err error
	m.Model, err = r.ToModel()
	if r.Item != nil {
		val, toModelErr := r.Item.ToModel()
		if toModelErr != nil {
			err = errors.Join(err, toModelErr)
		}
		m.Item = &val
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

func FromMongoQueryResultDataList(dbRecords []reservation_mongo.Model) ([]Model, error) {
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
	mongoWhereClause, err := where.Reservation.ToMongoWhereClause()
	if err != nil {
		return AggregateResult{}, err
	}

	// Convert API aggregate options to mongo aggregate options
	mongoAggFields := make([]reservation_mongo.AggregateFieldSpec, len(options.Fields))
	for i, f := range options.Fields {
		mongoAggFields[i] = reservation_mongo.AggregateFieldSpec{
			Field:  string(f.Field),
			Method: reservation_mongo.AggregateMethod(f.Method),
			Alias:  f.Alias,
		}
	}

	// Convert group-by fields to strings
	mongoGroupBy := make([]string, len(options.GroupBy))
	for i, g := range options.GroupBy {
		mongoGroupBy[i] = string(g)
	}

	mongoAggOptions := reservation_mongo.AggregateOptions{
		Fields:             mongoAggFields,
		GroupBy:            mongoGroupBy,
		ItemProjection:     options.ItemProjection,
		RegistryProjection: options.RegistryProjection,
	}

	result, err := reservation_mongo.Aggregate(
		ctx,
		m.db,
		reservation_mongo.WhereClause{
			Reservation: mongoWhereClause,
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
		row.ContactEmail = r.ContactEmail
		row.IsAnonymous = r.IsAnonymous
		if r.ItemId != nil {
			s := r.ItemId.Hex()
			row.ItemId = &s
		}
		row.Message = r.Message
		row.Quantity = r.Quantity
		if r.RegistryId != nil {
			s := r.RegistryId.Hex()
			row.RegistryId = &s
		}
		row.ReserverName = r.ReserverName
		// Copy aggregate fields (only those not in group-by)
		// Copy ref field Item
		if r.Item != nil {
			val, toModelErr := r.Item.ToModel()
			if toModelErr != nil {
				err = errors.Join(err, toModelErr)
			}
			row.Item = &val
		}
		// Copy ref field Registry
		if r.Registry != nil {
			val, toModelErr := r.Registry.ToModel()
			if toModelErr != nil {
				err = errors.Join(err, toModelErr)
			}
			row.Registry = &val
		}
		apiResults[i] = row
	}

	return AggregateResult{
		Data:  apiResults,
		Total: result.Total,
	}, nil
}
