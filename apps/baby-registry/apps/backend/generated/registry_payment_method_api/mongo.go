package registry_payment_method_api

import (
	"context"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/cart"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_payment_method"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_payment_method_mongo"
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
	mongoWhereClause, err := where.RegistryPaymentMethod.ToMongoWhereClause()
	if err != nil {
		return QueryResult{}, err
	}
	mongoWhereClauseCarts, err := where.Carts.ToMongoWhereClause()
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

	searchResult, err := registry_payment_method_mongo.Search(
		ctx,
		m.db,
		registry_payment_method_mongo.WhereClause{
			RegistryPaymentMethod: mongoWhereClause,
			Carts:                 mongoWhereClauseCarts,
			Owner:                 mongoWhereClauseOwner,
			Registry:              mongoWhereClauseRegistry,
		},
		registry_payment_method_mongo.LookupOptions{
			Projection:         projection.Projection,
			Sort:               options.Sort.ToMongoSortParams(),
			CartsProjection:    projection.Carts,
			OwnerProjection:    projection.Owner,
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

func (m *mongoClient) Create(ctx context.Context, obj registry_payment_method.Model, projection registry_payment_method.Projection) (registry_payment_method.Model, error) {

	createRecord, err := obj.ToMongoRecord(projection)
	if err != nil {
		return registry_payment_method.Model{}, err
	}
	var id primitive.ObjectID
	id, err = registry_payment_method_mongo.Create(ctx, m.db, createRecord)
	if err != nil {
		return registry_payment_method.Model{}, err
	}
	createRecord.Id = &id
	return createRecord.ToModel()
}

func (m *mongoClient) Update(ctx context.Context, obj registry_payment_method.Model, where registry_payment_method.WhereClause, projection registry_payment_method.Projection) (registry_payment_method.Model, error) {
	mongoWhereClause, err := where.ToMongoWhereClause()
	if err != nil {
		return registry_payment_method.Model{}, err
	}

	updateRecord, err := obj.ToMongoRecord(projection)
	if err != nil {
		return registry_payment_method.Model{}, err
	}

	err = registry_payment_method_mongo.Update(ctx, m.db, updateRecord, mongoWhereClause)
	if err != nil {
		return registry_payment_method.Model{}, err
	}

	return updateRecord.ToModel()
}

func (m *mongoClient) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Join(errors.New("invalid id"), err)
	}
	err = registry_payment_method_mongo.Delete(ctx, m.db, oid)
	if err != nil {
		return err
	}
	return nil
}

func FromMongoQueryResultData(r registry_payment_method_mongo.Model) (Model, error) {
	m := Model{}
	var err error
	m.Model, err = r.ToModel()
	if r.Carts != nil {
		val := make([]cart.Model, 0)
		var err error
		for _, rr := range *r.Carts {
			nextVal, nextErr := rr.ToModel()
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}
		m.Carts = &val
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

func FromMongoQueryResultDataList(dbRecords []registry_payment_method_mongo.Model) ([]Model, error) {
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
	mongoWhereClause, err := where.RegistryPaymentMethod.ToMongoWhereClause()
	if err != nil {
		return AggregateResult{}, err
	}

	// Convert API aggregate options to mongo aggregate options
	mongoAggFields := make([]registry_payment_method_mongo.AggregateFieldSpec, len(options.Fields))
	for i, f := range options.Fields {
		mongoAggFields[i] = registry_payment_method_mongo.AggregateFieldSpec{
			Field:  string(f.Field),
			Method: registry_payment_method_mongo.AggregateMethod(f.Method),
			Alias:  f.Alias,
		}
	}

	// Convert group-by fields to strings
	mongoGroupBy := make([]string, len(options.GroupBy))
	for i, g := range options.GroupBy {
		mongoGroupBy[i] = string(g)
	}

	mongoAggOptions := registry_payment_method_mongo.AggregateOptions{
		Fields:             mongoAggFields,
		GroupBy:            mongoGroupBy,
		CartsProjection:    options.CartsProjection,
		OwnerProjection:    options.OwnerProjection,
		RegistryProjection: options.RegistryProjection,
	}

	result, err := registry_payment_method_mongo.Aggregate(
		ctx,
		m.db,
		registry_payment_method_mongo.WhereClause{
			RegistryPaymentMethod: mongoWhereClause,
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
		row.BankAccountName = r.BankAccountName
		row.BankAccountNumber = r.BankAccountNumber
		row.BankIban = r.BankIban
		row.BankName = r.BankName
		row.BankRoutingNumber = r.BankRoutingNumber
		row.BankSwift = r.BankSwift
		row.DisplayName = r.DisplayName
		row.Enabled = r.Enabled
		row.Instructions = r.Instructions
		if r.OwnerId != nil {
			s := r.OwnerId.Hex()
			row.OwnerId = &s
		}
		row.PaymentUrl = r.PaymentUrl
		row.Position = r.Position
		row.RecipientEmail = r.RecipientEmail
		row.RecipientPhone = r.RecipientPhone
		if r.RegistryId != nil {
			s := r.RegistryId.Hex()
			row.RegistryId = &s
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
		// Copy ref field Carts
		if r.Carts != nil {
			val := make([]cart.Model, 0)
			for _, rr := range r.Carts {
				nextVal, nextErr := rr.ToModel()
				if nextErr != nil {
					err = errors.Join(err, nextErr)
				}
				val = append(val, nextVal)
			}
			row.Carts = val
		}
		apiResults[i] = row
	}

	return AggregateResult{
		Data:  apiResults,
		Total: result.Total,
	}, nil
}
