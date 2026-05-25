package owner_user_api

import (
	"context"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user_mongo"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request"
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
	mongoWhereClause, err := where.OwnerUser.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}
	mongoWhereClauseAddressAccessSessions, err := where.AddressAccessSessions.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}
	mongoWhereClauseRegistryApprovedGuests, err := where.RegistryApprovedGuests.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}
	mongoWhereClauseRegistrys, err := where.Registrys.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}
	mongoWhereClauseShippingAddressRequests, err := where.ShippingAddressRequests.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}

	searchResult, err := owner_user_mongo.Search(
		ctx,
		m.db,
		owner_user_mongo.WhereClause{
			OwnerUser:               mongoWhereClause,
			AddressAccessSessions:   mongoWhereClauseAddressAccessSessions,
			RegistryApprovedGuests:  mongoWhereClauseRegistryApprovedGuests,
			Registrys:               mongoWhereClauseRegistrys,
			ShippingAddressRequests: mongoWhereClauseShippingAddressRequests,
		},
		owner_user_mongo.LookupOptions{
			Projection:                        projection.Projection,
			Sort:                              options.Sort.ToMongoSortParams(),
			AddressAccessSessionsProjection:   projection.AddressAccessSessions,
			RegistryApprovedGuestsProjection:  projection.RegistryApprovedGuests,
			RegistrysProjection:               projection.Registrys,
			ShippingAddressRequestsProjection: projection.ShippingAddressRequests,
			Limit:                             options.Limit,
			Skip:                              options.Skip,
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

func (m *mongoClient) Create(ctx context.Context, obj owner_user.Model, projection owner_user.Projection) (owner_user.Model, error) {

	createRecord, err := obj.ToMongoRecord(projection)
	if err != nil {
		return owner_user.Model{}, err
	}
	var id primitive.ObjectID
	id, err = owner_user_mongo.Create(ctx, m.db, createRecord)
	if err != nil {
		return owner_user.Model{}, err
	}
	createRecord.Id = &id
	return createRecord.ToModel()
}

func (m *mongoClient) Update(ctx context.Context, obj owner_user.Model, where owner_user.WhereClause, projection owner_user.Projection) (owner_user.Model, error) {
	mongoWhereClause, err := where.ToMongoWhereClause()
	if err != nil {
		return owner_user.Model{}, err
	}

	updateRecord, err := obj.ToMongoRecord(projection)
	if err != nil {
		return owner_user.Model{}, err
	}

	err = owner_user_mongo.Update(ctx, m.db, updateRecord, mongoWhereClause)
	if err != nil {
		return owner_user.Model{}, err
	}

	return updateRecord.ToModel()
}

func (m *mongoClient) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Join(errors.New("invalid id"), err)
	}
	err = owner_user_mongo.Delete(ctx, m.db, oid)
	if err != nil {
		return err
	}
	return nil
}

func FromMongoQueryResultData(r owner_user_mongo.Model) (Model, error) {
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
	if r.RegistryApprovedGuests != nil {
		val := make([]registry_approved_guest.Model, 0)
		var err error
		for _, rr := range *r.RegistryApprovedGuests {
			nextVal, nextErr := rr.ToModel()
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}
		m.RegistryApprovedGuests = &val
	}
	if r.Registrys != nil {
		val := make([]registry.Model, 0)
		var err error
		for _, rr := range *r.Registrys {
			nextVal, nextErr := rr.ToModel()
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}
		m.Registrys = &val
	}
	if r.ShippingAddressRequests != nil {
		val := make([]shipping_address_request.Model, 0)
		var err error
		for _, rr := range *r.ShippingAddressRequests {
			nextVal, nextErr := rr.ToModel()
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}
		m.ShippingAddressRequests = &val
	}
	return m, err
}

func FromMongoQueryResultDataList(dbRecords []owner_user_mongo.Model) ([]Model, error) {
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
