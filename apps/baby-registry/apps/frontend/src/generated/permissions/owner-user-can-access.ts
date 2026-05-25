// This file is auto-generated. DO NOT EDIT.

import { OwnerUser } from '../model/owner-user-model';
import { NewCanReadActorRole, NewCanWriteActorRole } from './actor-role-can-access';
import { NewCanReadActorTrace, NewCanWriteActorTrace } from './actor-trace-can-access';
import { ActorRole } from '../model/actor-role-model';
import { ActorCanAccessFunc } from './actor';

type canAccessOwnerUser<T = OwnerUser> = ActorCanAccessFunc<T> & {
    field: {
        id: ActorCanAccessFunc<OwnerUser>; 
        actorRoles: ReturnType<typeof NewCanReadActorRole<OwnerUser>>, 
        created: ReturnType<typeof NewCanReadActorTrace<OwnerUser>>,
        email: ActorCanAccessFunc<OwnerUser>;
        name: ActorCanAccessFunc<OwnerUser>; 
        updated: ReturnType<typeof NewCanReadActorTrace<OwnerUser>>, 
        updatedByOwnerUser: ReturnType<typeof NewCanReadActorTrace<OwnerUser>>,
    }
};

const getAbacActorId = (obj: OwnerUser) => obj.id;

export const canReadOwnerUser = NewCanReadOwnerUser(
    (actorRoles: ActorRole[], obj?: OwnerUser) => {
        for (const actorRole of actorRoles) {
            switch(actorRole.role) {
            case 'Owner':
                if (!obj) {
                    return false;
                }
                // TODO pass actor instead of actorRoles, actor.actorId === getAbacActorId(obj))
                return true;
            case 'Super':
                return true;
            }
        }
        return false;
    },
);

export const canWriteOwnerUser = NewCanWriteOwnerUser(
    (actorRoles: ActorRole[], obj?: OwnerUser) => {
          for (const actorRole of actorRoles) {
              switch(actorRole.role) {
              case 'Owner':
                  if (!obj) {
                      return false;
                  }
                  // TODO pass actor instead of actorRoles, actor.actorId === getAbacActorId(obj))
                  return true;
              case 'Super':
                  return true;
              }
          }
          return false;
    },
);

export function NewCanReadOwnerUser<T = OwnerUser>(canAccessObj: ActorCanAccessFunc<T>): canAccessOwnerUser<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: OwnerUser) =>  true,
                actorRoles:  NewCanReadActorRole( (_actorRoles: ActorRole[], _obj?: OwnerUser) =>  true),
                created:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: OwnerUser) =>  true),
                email: (_actorRoles: ActorRole[], _obj?: OwnerUser) =>  true,
                name: (_actorRoles: ActorRole[], _obj?: OwnerUser) =>  true,
                updated:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: OwnerUser) =>  true),
                updatedByOwnerUser:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: OwnerUser) =>  true),
            },
        },
    );
}

export function NewCanWriteOwnerUser<T = OwnerUser>(canAccessObj: ActorCanAccessFunc<T>): canAccessOwnerUser<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: OwnerUser) =>  true,
                actorRoles:  NewCanWriteActorRole( (_actorRoles: ActorRole[], _obj?: OwnerUser) =>  true),
                created:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: OwnerUser) =>  true),
                email: (_actorRoles: ActorRole[], _obj?: OwnerUser) =>  true,
                name: (_actorRoles: ActorRole[], _obj?: OwnerUser) =>  true,
                updated:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: OwnerUser) =>  true),
                updatedByOwnerUser:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: OwnerUser) =>  true),
            },
        },
    );
}
