// This file is auto-generated. DO NOT EDIT.

import { RegistryApprovedGuest } from '../model/registry-approved-guest-model';
import { NewCanReadActorTrace, NewCanWriteActorTrace } from './actor-trace-can-access';
import { ActorRole } from '../model/actor-role-model';
import { ActorCanAccessFunc } from './actor';

type canAccessRegistryApprovedGuest<T = RegistryApprovedGuest> = ActorCanAccessFunc<T> & {
    field: {
        id: ActorCanAccessFunc<RegistryApprovedGuest>;
        accessLevel: ActorCanAccessFunc<RegistryApprovedGuest>; 
        created: ReturnType<typeof NewCanReadActorTrace<RegistryApprovedGuest>>,
        emailEnc: ActorCanAccessFunc<RegistryApprovedGuest>;
        emailHash: ActorCanAccessFunc<RegistryApprovedGuest>;
        name: ActorCanAccessFunc<RegistryApprovedGuest>;
        ownerId: ActorCanAccessFunc<RegistryApprovedGuest>;
        registryId: ActorCanAccessFunc<RegistryApprovedGuest>;
        status: ActorCanAccessFunc<RegistryApprovedGuest>; 
        updated: ReturnType<typeof NewCanReadActorTrace<RegistryApprovedGuest>>, 
        updatedByOwnerUser: ReturnType<typeof NewCanReadActorTrace<RegistryApprovedGuest>>,
    }
};

const getAbacOwnerId = (obj: RegistryApprovedGuest) => obj.ownerId;

export const canReadRegistryApprovedGuest = NewCanReadRegistryApprovedGuest(
    (actorRoles: ActorRole[], obj?: RegistryApprovedGuest) => {
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

export const canWriteRegistryApprovedGuest = NewCanWriteRegistryApprovedGuest(
    (actorRoles: ActorRole[], obj?: RegistryApprovedGuest) => {
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

export function NewCanReadRegistryApprovedGuest<T = RegistryApprovedGuest>(canAccessObj: ActorCanAccessFunc<T>): canAccessRegistryApprovedGuest<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true,
                accessLevel: (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true,
                created:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true),
                emailEnc: (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true,
                emailHash: (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true,
                name: (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true,
                ownerId: (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true,
                registryId: (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true,
                status: (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true,
                updated:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true),
                updatedByOwnerUser:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true),
            },
        },
    );
}

export function NewCanWriteRegistryApprovedGuest<T = RegistryApprovedGuest>(canAccessObj: ActorCanAccessFunc<T>): canAccessRegistryApprovedGuest<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true,
                accessLevel: (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true,
                created:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true),
                emailEnc: (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true,
                emailHash: (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true,
                name: (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true,
                ownerId: (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true,
                registryId: (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true,
                status: (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true,
                updated:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true),
                updatedByOwnerUser:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: RegistryApprovedGuest) =>  true),
            },
        },
    );
}
