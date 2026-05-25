package registry_api

import (
	"context"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_mongo"
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
	mongoWhereClause, err := where.Registry.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}
	mongoWhereClauseRegistryItems, err := where.RegistryItems.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}
	mongoWhereClauseReservations, err := where.Reservations.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}
	mongoWhereClauseOwner, err := where.Owner.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}

	searchResult, err := registry_mongo.Search(
		ctx,
		m.db,
		registry_mongo.WhereClause{
			Registry:      mongoWhereClause,
			RegistryItems: mongoWhereClauseRegistryItems,
			Reservations:  mongoWhereClauseReservations,
			Owner:         mongoWhereClauseOwner,
		},
		registry_mongo.LookupOptions{
			Projection:              projection.Projection,
			Sort:                    options.Sort.ToMongoSortParams(),
			RegistryItemsProjection: projection.RegistryItems,
			ReservationsProjection:  projection.Reservations,
			OwnerProjection:         projection.Owner,
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

func (m *mongoClient) Create(ctx context.Context, obj registry.Model, projection registry.Projection) (registry.Model, error) {

	createRecord, err := obj.ToMongoRecord(projection)
	if err != nil {
		return registry.Model{}, err
	}
	var id primitive.ObjectID
	id, err = registry_mongo.Create(ctx, m.db, createRecord)
	if err != nil {
		return registry.Model{}, err
	}
	createRecord.Id = &id
	return createRecord.ToModel()
}

func (m *mongoClient) Update(ctx context.Context, obj registry.Model, where registry.WhereClause, projection registry.Projection) (registry.Model, error) {
	mongoWhereClause, err := where.ToMongoWhereClause()
	if err != nil {
		return registry.Model{}, err
	}

	updateRecord, err := obj.ToMongoRecord(projection)
	if err != nil {
		return registry.Model{}, err
	}

	err = registry_mongo.Update(ctx, m.db, updateRecord, mongoWhereClause)
	if err != nil {
		return registry.Model{}, err
	}

	return updateRecord.ToModel()
}

func (m *mongoClient) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Join(errors.New("invalid id"), err)
	}
	err = registry_mongo.Delete(ctx, m.db, oid)
	if err != nil {
		return err
	}
	return nil
}

func FromMongoQueryResultData(r registry_mongo.Model) (Model, error) {
	m := Model{}
	var err error
	m.Model, err = r.ToModel()
	if r.RegistryItems != nil {
		val := make([]registry_item.Model, 0)
		var err error
		for _, rr := range *r.RegistryItems {
			nextVal, nextErr := rr.ToModel()
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}
		m.RegistryItems = &val
	}
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
	if r.Owner != nil {
		val, toModelErr := r.Owner.ToModel()
		if toModelErr != nil {
			err = errors.Join(err, toModelErr)
		}
		m.Owner = &val
	}
	return m, err
}

func FromMongoQueryResultDataList(dbRecords []registry_mongo.Model) ([]Model, error) {
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
