// This file is auto-generated. DO NOT EDIT.

import { Registry } from '../model/registry-model';
import { NewCanReadActorTrace, NewCanWriteActorTrace } from './actor-trace-can-access';
import { ActorRole } from '../model/actor-role-model';
import { ActorCanAccessFunc } from './actor';

type canAccessRegistry<T = Registry> = ActorCanAccessFunc<T> & {
    field: {
        id: ActorCanAccessFunc<Registry>;
        coverImageUrl: ActorCanAccessFunc<Registry>; 
        created: ReturnType<typeof NewCanReadActorTrace<Registry>>,
        dueDate: ActorCanAccessFunc<Registry>;
        isPublic: ActorCanAccessFunc<Registry>;
        ownerId: ActorCanAccessFunc<Registry>;
        parentNames: ActorCanAccessFunc<Registry>;
        slug: ActorCanAccessFunc<Registry>;
        themeColor: ActorCanAccessFunc<Registry>;
        title: ActorCanAccessFunc<Registry>; 
        updated: ReturnType<typeof NewCanReadActorTrace<Registry>>, 
        updatedByOwnerUser: ReturnType<typeof NewCanReadActorTrace<Registry>>,
        welcomeMessage: ActorCanAccessFunc<Registry>;
    }
};

const getAbacOwnerId = (obj: Registry) => obj.ownerId;

export const canReadRegistry = NewCanReadRegistry(
    (actorRoles: ActorRole[], obj?: Registry) => {
        for (const actorRole of actorRoles) {
            switch(actorRole.role) {
            case 'Owner':
                if (!obj) {
                    return false;
                }
                if (actorRole.ownerId === getAbacOwnerId(obj)) {
                    return true;
                }
                return true;
            case 'Super':
                return true;
            }
        }
        return false;
    },
);

export const canWriteRegistry = NewCanWriteRegistry(
    (actorRoles: ActorRole[], obj?: Registry) => {
          for (const actorRole of actorRoles) {
              switch(actorRole.role) {
              case 'Owner':
                  if (!obj) {
                      return false;
                  }
                  if (actorRole.ownerId !== getAbacOwnerId(obj)) {
                      return false;
                  }
                  return true;
              case 'Super':
                  return true;
              }
          }
          return false;
    },
);

export function NewCanReadRegistry<T = Registry>(canAccessObj: ActorCanAccessFunc<T>): canAccessRegistry<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
                coverImageUrl: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
                created:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Registry) =>  true),
                dueDate: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
                isPublic: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
                ownerId: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
                parentNames: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
                slug: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
                themeColor: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
                title: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
                updated:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Registry) =>  true),
                updatedByOwnerUser:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Registry) =>  true),
                welcomeMessage: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
            },
        },
    );
}

export function NewCanWriteRegistry<T = Registry>(canAccessObj: ActorCanAccessFunc<T>): canAccessRegistry<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
                coverImageUrl: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
                created:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Registry) =>  true),
                dueDate: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
                isPublic: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
                ownerId: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
                parentNames: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
                slug: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
                themeColor: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
                title: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
                updated:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Registry) =>  true),
                updatedByOwnerUser:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Registry) =>  true),
                welcomeMessage: (_actorRoles: ActorRole[], _obj?: Registry) =>  true,
            },
        },
    );
}
