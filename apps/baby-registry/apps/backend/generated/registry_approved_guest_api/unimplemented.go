package registry_approved_guest_api

import (
	"context"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_approved_guest"
)

type UnimplementedClient struct{}

func (c *UnimplementedClient) Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error) {
	return QueryResult{}, errors.New("search is not implemented")
}

func (c *UnimplementedClient) Create(ctx context.Context, obj registry_approved_guest.Model, projection registry_approved_guest.Projection) (registry_approved_guest.Model, error) {
	return registry_approved_guest.Model{}, errors.New("create is not implemented")
}

func (c *UnimplementedClient) Update(ctx context.Context, obj registry_approved_guest.Model, where registry_approved_guest.WhereClause, projection registry_approved_guest.Projection) (registry_approved_guest.Model, error) {
	return registry_approved_guest.Model{}, errors.New("update is not implemented")
}

func (c *UnimplementedClient) Delete(ctx context.Context, id string) error {
	return errors.New("delete is not implemented")
}
