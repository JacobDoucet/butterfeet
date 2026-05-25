package permissions_api

import (
	"context"
	"fmt"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/coded_error"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/owner_user_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
)

type Client interface {
	SelectActorById(ctx context.Context, actorType permissions.ActorType, actorId string) (permissions.Actor, error)
	UseOwnerUserClient(api owner_user_api.Client) Client
	UseOwnerUserProjection(projection owner_user_api.Projection) Client
}

func New() Client {
	return &client{}
}

type client struct {
	ownerUser           owner_user_api.Client
	ownerUserProjection *owner_user_api.Projection
}

func (c *client) UseOwnerUserClient(api owner_user_api.Client) Client {
	c.ownerUser = api
	return c
}

func (c *client) UseOwnerUserProjection(projection owner_user_api.Projection) Client {
	c.ownerUserProjection = &projection
	return c
}

func (c *client) SelectActorById(ctx context.Context, actorType permissions.ActorType, actorId string) (permissions.Actor, error) {
	switch actorType {
	case permissions.ActorTypeOwnerUser:
		if c.ownerUser == nil {
			return nil, coded_error.NewUnexpectedError("ownerUser api not provided")
		}
		projection := owner_user_api.Projection{
			Projection: owner_user.NewProjection(true),
		}
		if c.ownerUserProjection != nil {
			projection = *c.ownerUserProjection
		}
		actor, _, err := c.ownerUser.SelectById(
			ctx,
			permissions.NewSuperActor(),
			owner_user.SelectByIdQuery{Id: actorId},
			projection,
		)
		if err != nil {
			return nil, err
		}
		return &actor, nil
	}
	return nil, coded_error.NewUnexpectedError(fmt.Sprintf("unhandled actor type %s", actorType))
}
