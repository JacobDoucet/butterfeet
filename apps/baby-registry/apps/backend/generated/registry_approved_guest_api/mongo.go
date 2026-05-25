package registry_approved_guest_api

import (
	"context"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest_mongo"
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
	mongoWhereClause, err := where.RegistryApprovedGuest.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}
	mongoWhereClauseAddressAccessSessions, err := where.AddressAccessSessions.ToMongoWhereClause()
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

	searchResult, err := registry_approved_guest_mongo.Search(
		ctx,
		m.db,
		registry_approved_guest_mongo.WhereClause{
			RegistryApprovedGuest: mongoWhereClause,
			AddressAccessSessions: mongoWhereClauseAddressAccessSessions,
			Owner:                 mongoWhereClauseOwner,
			Registry:              mongoWhereClauseRegistry,
		},
		registry_approved_guest_mongo.LookupOptions{
			Projection:                      projection.Projection,
			Sort:                            options.Sort.ToMongoSortParams(),
			AddressAccessSessionsProjection: projection.AddressAccessSessions,
			OwnerProjection:                 projection.Owner,
			RegistryProjection:              projection.Registry,
			Limit:                           options.Limit,
			Skip:                            options.Skip,
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

func (m *mongoClient) Create(ctx context.Context, obj registry_approved_guest.Model, projection registry_approved_guest.Projection) (registry_approved_guest.Model, error) {

	createRecord, err := obj.ToMongoRecord(projection)
	if err != nil {
		return registry_approved_guest.Model{}, err
	}
	var id primitive.ObjectID
	id, err = registry_approved_guest_mongo.Create(ctx, m.db, createRecord)
	if err != nil {
		return registry_approved_guest.Model{}, err
	}
	createRecord.Id = &id
	return createRecord.ToModel()
}

func (m *mongoClient) Update(ctx context.Context, obj registry_approved_guest.Model, where registry_approved_guest.WhereClause, projection registry_approved_guest.Projection) (registry_approved_guest.Model, error) {
	mongoWhereClause, err := where.ToMongoWhereClause()
	if err != nil {
		return registry_approved_guest.Model{}, err
	}

	updateRecord, err := obj.ToMongoRecord(projection)
	if err != nil {
		return registry_approved_guest.Model{}, err
	}

	err = registry_approved_guest_mongo.Update(ctx, m.db, updateRecord, mongoWhereClause)
	if err != nil {
		return registry_approved_guest.Model{}, err
	}

	return updateRecord.ToModel()
}

func (m *mongoClient) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Join(errors.New("invalid id"), err)
	}
	err = registry_approved_guest_mongo.Delete(ctx, m.db, oid)
	if err != nil {
		return err
	}
	return nil
}

func FromMongoQueryResultData(r registry_approved_guest_mongo.Model) (Model, error) {
	m := Model{}
	var err error
	m.Model, err = r.ToModel()
	if r.AddressAccessSessions != nil {
		val := make([]address_access_session.Model, 0)
		var err error
		for _, rr := range *r.AddressAccessSessions {
			nextVal, nextErr := rr.ToModel()
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}
		m.AddressAccessSessions = &val
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

func FromMongoQueryResultDataList(dbRecords []registry_approved_guest_mongo.Model) ([]Model, error) {
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
