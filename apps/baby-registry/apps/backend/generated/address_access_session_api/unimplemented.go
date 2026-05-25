package address_access_session_api

import (
	"context"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/address_access_session"
)

type UnimplementedClient struct{}

func (c *UnimplementedClient) Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error) {
	return QueryResult{}, errors.New("search is not implemented")
}

func (c *UnimplementedClient) Create(ctx context.Context, obj address_access_session.Model, projection address_access_session.Projection) (address_access_session.Model, error) {
	return address_access_session.Model{}, errors.New("create is not implemented")
}

func (c *UnimplementedClient) Update(ctx context.Context, obj address_access_session.Model, where address_access_session.WhereClause, projection address_access_session.Projection) (address_access_session.Model, error) {
	return address_access_session.Model{}, errors.New("update is not implemented")
}

func (c *UnimplementedClient) Delete(ctx context.Context, id string) error {
	return errors.New("delete is not implemented")
}

func (c *UnimplementedClient) Aggregate(ctx context.Context, query WhereClause, options AggregateOptions) (AggregateResult, error) {
	return AggregateResult{}, errors.New("aggregate is not implemented")
}
