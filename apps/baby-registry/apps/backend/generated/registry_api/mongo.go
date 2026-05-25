package registry_api

import (
	"context"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_mongo"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
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
	mongoWhereClause, err := where.Registry.ToMongoWhereClause()
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
	mongoWhereClauseRegistryItems, err := where.RegistryItems.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}
	mongoWhereClauseReservations, err := where.Reservations.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}
	mongoWhereClauseShippingAddressRequests, err := where.ShippingAddressRequests.ToMongoWhereClause()
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
			Registry:                mongoWhereClause,
			AddressAccessSessions:   mongoWhereClauseAddressAccessSessions,
			RegistryApprovedGuests:  mongoWhereClauseRegistryApprovedGuests,
			RegistryItems:           mongoWhereClauseRegistryItems,
			Reservations:            mongoWhereClauseReservations,
			ShippingAddressRequests: mongoWhereClauseShippingAddressRequests,
			Owner:                   mongoWhereClauseOwner,
		},
		registry_mongo.LookupOptions{
			Projection:                        projection.Projection,
			Sort:                              options.Sort.ToMongoSortParams(),
			AddressAccessSessionsProjection:   projection.AddressAccessSessions,
			RegistryApprovedGuestsProjection:  projection.RegistryApprovedGuests,
			RegistryItemsProjection:           projection.RegistryItems,
			ReservationsProjection:            projection.Reservations,
			ShippingAddressRequestsProjection: projection.ShippingAddressRequests,
			OwnerProjection:                   projection.Owner,
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

func (m *mongoClient) Aggregate(ctx context.Context, where WhereClause, options AggregateOptions) (AggregateResult, error) {
	mongoWhereClause, err := where.Registry.ToMongoWhereClause()
	if err != nil {
		return AggregateResult{}, err
	}

	// Convert API aggregate options to mongo aggregate options
	mongoAggFields := make([]registry_mongo.AggregateFieldSpec, len(options.Fields))
	for i, f := range options.Fields {
		mongoAggFields[i] = registry_mongo.AggregateFieldSpec{
			Field:  string(f.Field),
			Method: registry_mongo.AggregateMethod(f.Method),
			Alias:  f.Alias,
		}
	}

	// Convert group-by fields to strings
	mongoGroupBy := make([]string, len(options.GroupBy))
	for i, g := range options.GroupBy {
		mongoGroupBy[i] = string(g)
	}

	mongoAggOptions := registry_mongo.AggregateOptions{
		Fields:                            mongoAggFields,
		GroupBy:                           mongoGroupBy,
		AddressAccessSessionsProjection:   options.AddressAccessSessionsProjection,
		RegistryApprovedGuestsProjection:  options.RegistryApprovedGuestsProjection,
		RegistryItemsProjection:           options.RegistryItemsProjection,
		ReservationsProjection:            options.ReservationsProjection,
		ShippingAddressRequestsProjection: options.ShippingAddressRequestsProjection,
		OwnerProjection:                   options.OwnerProjection,
	}

	result, err := registry_mongo.Aggregate(
		ctx,
		m.db,
		registry_mongo.WhereClause{
			Registry: mongoWhereClause,
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
		row.CoverImageUrl = r.CoverImageUrl
		row.DueDate = r.DueDate
		row.IsPublic = r.IsPublic
		if r.OwnerId != nil {
			s := r.OwnerId.Hex()
			row.OwnerId = &s
		}
		row.ParentNames = r.ParentNames
		row.ShippingCity = r.ShippingCity
		row.ShippingCountry = r.ShippingCountry
		row.ShippingDeliveryNotes = r.ShippingDeliveryNotes
		row.ShippingLine1 = r.ShippingLine1
		row.ShippingLine2 = r.ShippingLine2
		row.ShippingPolicyVersion = r.ShippingPolicyVersion
		row.ShippingPostalCode = r.ShippingPostalCode
		row.ShippingRecipientName = r.ShippingRecipientName
		row.ShippingRegion = r.ShippingRegion
		row.Slug = r.Slug
		row.ThemeColor = r.ThemeColor
		row.Title = r.Title
		row.WelcomeMessage = r.WelcomeMessage
		// Copy aggregate fields (only those not in group-by)
		// Copy ref field Owner
		if r.Owner != nil {
			val, toModelErr := r.Owner.ToModel()
			if toModelErr != nil {
				err = errors.Join(err, toModelErr)
			}
			row.Owner = &val
		}
		// Copy ref field AddressAccessSessions
		if r.AddressAccessSessions != nil {
			val := make([]address_access_session.Model, 0)
			for _, rr := range r.AddressAccessSessions {
				nextVal, nextErr := rr.ToModel()
				if nextErr != nil {
					err = errors.Join(err, nextErr)
				}
				val = append(val, nextVal)
			}
			row.AddressAccessSessions = val
		}
		// Copy ref field RegistryApprovedGuests
		if r.RegistryApprovedGuests != nil {
			val := make([]registry_approved_guest.Model, 0)
			for _, rr := range r.RegistryApprovedGuests {
				nextVal, nextErr := rr.ToModel()
				if nextErr != nil {
					err = errors.Join(err, nextErr)
				}
				val = append(val, nextVal)
			}
			row.RegistryApprovedGuests = val
		}
		// Copy ref field RegistryItems
		if r.RegistryItems != nil {
			val := make([]registry_item.Model, 0)
			for _, rr := range r.RegistryItems {
				nextVal, nextErr := rr.ToModel()
				if nextErr != nil {
					err = errors.Join(err, nextErr)
				}
				val = append(val, nextVal)
			}
			row.RegistryItems = val
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
		// Copy ref field ShippingAddressRequests
		if r.ShippingAddressRequests != nil {
			val := make([]shipping_address_request.Model, 0)
			for _, rr := range r.ShippingAddressRequests {
				nextVal, nextErr := rr.ToModel()
				if nextErr != nil {
					err = errors.Join(err, nextErr)
				}
				val = append(val, nextVal)
			}
			row.ShippingAddressRequests = val
		}
		apiResults[i] = row
	}

	return AggregateResult{
		Data:  apiResults,
		Total: result.Total,
	}, nil
}
