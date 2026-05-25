package registry_item_api

import (
	"context"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
)

type UnimplementedClient struct{}

func (c *UnimplementedClient) Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error) {
	return QueryResult{}, errors.New("search is not implemented")
}

func (c *UnimplementedClient) Create(ctx context.Context, obj registry_item.Model, projection registry_item.Projection) (registry_item.Model, error) {
	return registry_item.Model{}, errors.New("create is not implemented")
}

func (c *UnimplementedClient) Update(ctx context.Context, obj registry_item.Model, where registry_item.WhereClause, projection registry_item.Projection) (registry_item.Model, error) {
	return registry_item.Model{}, errors.New("update is not implemented")
}

func (c *UnimplementedClient) Delete(ctx context.Context, id string) error {
	return errors.New("delete is not implemented")
}

func (c *UnimplementedClient) Aggregate(ctx context.Context, query WhereClause, options AggregateOptions) (AggregateResult, error) {
	return AggregateResult{}, errors.New("aggregate is not implemented")
}
