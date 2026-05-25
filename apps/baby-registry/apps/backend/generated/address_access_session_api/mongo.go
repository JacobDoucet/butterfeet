package address_access_session_api

import (
	"context"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session_mongo"
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
	mongoWhereClause, err := where.AddressAccessSession.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}
	mongoWhereClauseApprovedGuest, err := where.ApprovedGuest.ToMongoWhereClause()
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

	searchResult, err := address_access_session_mongo.Search(
		ctx,
		m.db,
		address_access_session_mongo.WhereClause{
			AddressAccessSession: mongoWhereClause,
			ApprovedGuest:        mongoWhereClauseApprovedGuest,
			Owner:                mongoWhereClauseOwner,
			Registry:             mongoWhereClauseRegistry,
		},
		address_access_session_mongo.LookupOptions{
			Projection:              projection.Projection,
			Sort:                    options.Sort.ToMongoSortParams(),
			ApprovedGuestProjection: projection.ApprovedGuest,
			OwnerProjection:         projection.Owner,
			RegistryProjection:      projection.Registry,
			Limit:                   options.Limit,
			Skip:                    options.Skip,
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

func (m *mongoClient) Create(ctx context.Context, obj address_access_session.Model, projection address_access_session.Projection) (address_access_session.Model, error) {

	createRecord, err := obj.ToMongoRecord(projection)
	if err != nil {
		return address_access_session.Model{}, err
	}
	var id primitive.ObjectID
	id, err = address_access_session_mongo.Create(ctx, m.db, createRecord)
	if err != nil {
		return address_access_session.Model{}, err
	}
	createRecord.Id = &id
	return createRecord.ToModel()
}

func (m *mongoClient) Update(ctx context.Context, obj address_access_session.Model, where address_access_session.WhereClause, projection address_access_session.Projection) (address_access_session.Model, error) {
	mongoWhereClause, err := where.ToMongoWhereClause()
	if err != nil {
		return address_access_session.Model{}, err
	}

	updateRecord, err := obj.ToMongoRecord(projection)
	if err != nil {
		return address_access_session.Model{}, err
	}

	err = address_access_session_mongo.Update(ctx, m.db, updateRecord, mongoWhereClause)
	if err != nil {
		return address_access_session.Model{}, err
	}

	return updateRecord.ToModel()
}

func (m *mongoClient) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Join(errors.New("invalid id"), err)
	}
	err = address_access_session_mongo.Delete(ctx, m.db, oid)
	if err != nil {
		return err
	}
	return nil
}

func FromMongoQueryResultData(r address_access_session_mongo.Model) (Model, error) {
	m := Model{}
	var err error
	m.Model, err = r.ToModel()
	if r.ApprovedGuest != nil {
		val, toModelErr := r.ApprovedGuest.ToModel()
		if toModelErr != nil {
			err = errors.Join(err, toModelErr)
		}
		m.ApprovedGuest = &val
	}
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
	return m, err
}

func FromMongoQueryResultDataList(dbRecords []address_access_session_mongo.Model) ([]Model, error) {
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
	mongoWhereClause, err := where.AddressAccessSession.ToMongoWhereClause()
	if err != nil {
		return AggregateResult{}, err
	}

	// Convert API aggregate options to mongo aggregate options
	mongoAggFields := make([]address_access_session_mongo.AggregateFieldSpec, len(options.Fields))
	for i, f := range options.Fields {
		mongoAggFields[i] = address_access_session_mongo.AggregateFieldSpec{
			Field:  string(f.Field),
			Method: address_access_session_mongo.AggregateMethod(f.Method),
			Alias:  f.Alias,
		}
	}

	// Convert group-by fields to strings
	mongoGroupBy := make([]string, len(options.GroupBy))
	for i, g := range options.GroupBy {
		mongoGroupBy[i] = string(g)
	}

	mongoAggOptions := address_access_session_mongo.AggregateOptions{
		Fields:                  mongoAggFields,
		GroupBy:                 mongoGroupBy,
		ApprovedGuestProjection: options.ApprovedGuestProjection,
		OwnerProjection:         options.OwnerProjection,
		RegistryProjection:      options.RegistryProjection,
	}

	result, err := address_access_session_mongo.Aggregate(
		ctx,
		m.db,
		address_access_session_mongo.WhereClause{
			AddressAccessSession: mongoWhereClause,
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
		if r.ApprovedGuestId != nil {
			s := r.ApprovedGuestId.Hex()
			row.ApprovedGuestId = &s
		}
		row.EmailHash = r.EmailHash
		row.ExpiresAt = r.ExpiresAt
		if r.OwnerId != nil {
			s := r.OwnerId.Hex()
			row.OwnerId = &s
		}
		row.PolicyVersionAtIssue = r.PolicyVersionAtIssue
		if r.RegistryId != nil {
			s := r.RegistryId.Hex()
			row.RegistryId = &s
		}
		row.TokenHash = r.TokenHash
		// Copy aggregate fields (only those not in group-by)
		// Copy ref field ApprovedGuest
		if r.ApprovedGuest != nil {
			val, toModelErr := r.ApprovedGuest.ToModel()
			if toModelErr != nil {
				err = errors.Join(err, toModelErr)
			}
			row.ApprovedGuest = &val
		}
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
		apiResults[i] = row
	}

	return AggregateResult{
		Data:  apiResults,
		Total: result.Total,
	}, nil
}
