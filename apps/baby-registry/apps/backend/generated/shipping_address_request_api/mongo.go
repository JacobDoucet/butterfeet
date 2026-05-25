package shipping_address_request_api

import (
	"context"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request_mongo"
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
	mongoWhereClause, err := where.ShippingAddressRequest.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}
	mongoWhereClauseOwner, err := where.Owner.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}
	mongoWhereClauseRegistry, err := where.Registry.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}
	mongoWhereClauseRegistryItem, err := where.RegistryItem.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}

	searchResult, err := shipping_address_request_mongo.Search(
		ctx,
		m.db,
		shipping_address_request_mongo.WhereClause{
			ShippingAddressRequest: mongoWhereClause,
			Owner:                  mongoWhereClauseOwner,
			Registry:               mongoWhereClauseRegistry,
			RegistryItem:           mongoWhereClauseRegistryItem,
		},
		shipping_address_request_mongo.LookupOptions{
			Projection:             projection.Projection,
			Sort:                   options.Sort.ToMongoSortParams(),
			OwnerProjection:        projection.Owner,
			RegistryProjection:     projection.Registry,
			RegistryItemProjection: projection.RegistryItem,
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

func (m *mongoClient) Create(ctx context.Context, obj shipping_address_request.Model, projection shipping_address_request.Projection) (shipping_address_request.Model, error) {

	createRecord, err := obj.ToMongoRecord(projection)
	if err != nil {
		return shipping_address_request.Model{}, err
	}
	var id primitive.ObjectID
	id, err = shipping_address_request_mongo.Create(ctx, m.db, createRecord)
	if err != nil {
		return shipping_address_request.Model{}, err
	}
	createRecord.Id = &id
	return createRecord.ToModel()
}

func (m *mongoClient) Update(ctx context.Context, obj shipping_address_request.Model, where shipping_address_request.WhereClause, projection shipping_address_request.Projection) (shipping_address_request.Model, error) {
	mongoWhereClause, err := where.ToMongoWhereClause()
	if err != nil {
		return shipping_address_request.Model{}, err
	}

	updateRecord, err := obj.ToMongoRecord(projection)
	if err != nil {
		return shipping_address_request.Model{}, err
	}

	err = shipping_address_request_mongo.Update(ctx, m.db, updateRecord, mongoWhereClause)
	if err != nil {
		return shipping_address_request.Model{}, err
	}

	return updateRecord.ToModel()
}

func (m *mongoClient) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Join(errors.New("invalid id"), err)
	}
	err = shipping_address_request_mongo.Delete(ctx, m.db, oid)
	if err != nil {
		return err
	}
	return nil
}

func FromMongoQueryResultData(r shipping_address_request_mongo.Model) (Model, error) {
	m := Model{}
	var err error
	m.Model, err = r.ToModel()
	if r.Owner != nil {
		val, toModelErr := r.Owner.ToModel()
		if toModelErr != nil {
			err = errors.Join(err, toModelErr)
		}
		m.Owner = &val
	}
	if r.Registry != nil {
		val, toModelErr := r.Registry.ToModel()
		if toModelErr != nil {
			err = errors.Join(err, toModelErr)
		}
		m.Registry = &val
	}
	if r.RegistryItem != nil {
		val, toModelErr := r.RegistryItem.ToModel()
		if toModelErr != nil {
			err = errors.Join(err, toModelErr)
		}
		m.RegistryItem = &val
	}
	return m, err
}

func FromMongoQueryResultDataList(dbRecords []shipping_address_request_mongo.Model) ([]Model, error) {
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
	mongoWhereClause, err := where.ShippingAddressRequest.ToMongoWhereClause()
	if err != nil {
		return AggregateResult{}, err
	}

	// Convert API aggregate options to mongo aggregate options
	mongoAggFields := make([]shipping_address_request_mongo.AggregateFieldSpec, len(options.Fields))
	for i, f := range options.Fields {
		mongoAggFields[i] = shipping_address_request_mongo.AggregateFieldSpec{
			Field:  string(f.Field),
			Method: shipping_address_request_mongo.AggregateMethod(f.Method),
			Alias:  f.Alias,
		}
	}

	// Convert group-by fields to strings
	mongoGroupBy := make([]string, len(options.GroupBy))
	for i, g := range options.GroupBy {
		mongoGroupBy[i] = string(g)
	}

	mongoAggOptions := shipping_address_request_mongo.AggregateOptions{
		Fields:                 mongoAggFields,
		GroupBy:                mongoGroupBy,
		OwnerProjection:        options.OwnerProjection,
		RegistryProjection:     options.RegistryProjection,
		RegistryItemProjection: options.RegistryItemProjection,
	}

	result, err := shipping_address_request_mongo.Aggregate(
		ctx,
		m.db,
		shipping_address_request_mongo.WhereClause{
			ShippingAddressRequest: mongoWhereClause,
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
		row.DecisionReason = r.DecisionReason
		row.EmailEnc = r.EmailEnc
		row.EmailHash = r.EmailHash
		row.Name = r.Name
		row.Note = r.Note
		if r.OwnerId != nil {
			s := r.OwnerId.Hex()
			row.OwnerId = &s
		}
		row.PolicyVersion = r.PolicyVersion
		if r.RegistryId != nil {
			s := r.RegistryId.Hex()
			row.RegistryId = &s
		}
		if r.RegistryItemId != nil {
			s := r.RegistryItemId.Hex()
			row.RegistryItemId = &s
		}
		// Copy aggregate fields (only those not in group-by)
		// Copy ref field Owner
		if r.Owner != nil {
			val, toModelErr := r.Owner.ToModel()
			if toModelErr != nil {
				err = errors.Join(err, toModelErr)
			}
			row.Owner = &val
		}
		// Copy ref field Registry
		if r.Registry != nil {
			val, toModelErr := r.Registry.ToModel()
			if toModelErr != nil {
				err = errors.Join(err, toModelErr)
			}
			row.Registry = &val
		}
		// Copy ref field RegistryItem
		if r.RegistryItem != nil {
			val, toModelErr := r.RegistryItem.ToModel()
			if toModelErr != nil {
				err = errors.Join(err, toModelErr)
			}
			row.RegistryItem = &val
		}
		apiResults[i] = row
	}

	return AggregateResult{
		Data:  apiResults,
		Total: result.Total,
	}, nil
}
