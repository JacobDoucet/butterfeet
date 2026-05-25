package shipping_address_request_api

import (
	"context"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/shipping_address_request"
)

type UnimplementedClient struct{}

func (c *UnimplementedClient) Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error) {
	return QueryResult{}, errors.New("search is not implemented")
}

func (c *UnimplementedClient) Create(ctx context.Context, obj shipping_address_request.Model, projection shipping_address_request.Projection) (shipping_address_request.Model, error) {
	return shipping_address_request.Model{}, errors.New("create is not implemented")
}

func (c *UnimplementedClient) Update(ctx context.Context, obj shipping_address_request.Model, where shipping_address_request.WhereClause, projection shipping_address_request.Projection) (shipping_address_request.Model, error) {
	return shipping_address_request.Model{}, errors.New("update is not implemented")
}

func (c *UnimplementedClient) Delete(ctx context.Context, id string) error {
	return errors.New("delete is not implemented")
}

func (c *UnimplementedClient) Aggregate(ctx context.Context, query WhereClause, options AggregateOptions) (AggregateResult, error) {
	return AggregateResult{}, errors.New("aggregate is not implemented")
}
