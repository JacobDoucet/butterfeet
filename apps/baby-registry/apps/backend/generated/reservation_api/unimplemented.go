package reservation_api

import (
	"context"
	"errors"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
)

type UnimplementedClient struct{}

func (c *UnimplementedClient) Search(ctx context.Context, query WhereClause, options QueryOptions) (QueryResult, error) {
	return QueryResult{}, errors.New("search is not implemented")
}

func (c *UnimplementedClient) Create(ctx context.Context, obj reservation.Model, projection reservation.Projection) (reservation.Model, error) {
	return reservation.Model{}, errors.New("create is not implemented")
}

func (c *UnimplementedClient) Update(ctx context.Context, obj reservation.Model, where reservation.WhereClause, projection reservation.Projection) (reservation.Model, error) {
	return reservation.Model{}, errors.New("update is not implemented")
}

func (c *UnimplementedClient) Delete(ctx context.Context, id string) error {
	return errors.New("delete is not implemented")
}

func (c *UnimplementedClient) Aggregate(ctx context.Context, query WhereClause, options AggregateOptions) (AggregateResult, error) {
	return AggregateResult{}, errors.New("aggregate is not implemented")
}
