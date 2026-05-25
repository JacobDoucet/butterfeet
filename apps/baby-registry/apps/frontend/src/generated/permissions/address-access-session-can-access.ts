// This file is auto-generated. DO NOT EDIT.

import { AddressAccessSession } from '../model/address-access-session-model';
import { NewCanReadActorTrace, NewCanWriteActorTrace } from './actor-trace-can-access';
import { ActorRole } from '../model/actor-role-model';
import { ActorCanAccessFunc } from './actor';

type canAccessAddressAccessSession<T = AddressAccessSession> = ActorCanAccessFunc<T> & {
    field: {
        id: ActorCanAccessFunc<AddressAccessSession>;
        approvedGuestId: ActorCanAccessFunc<AddressAccessSession>; 
        created: ReturnType<typeof NewCanReadActorTrace<AddressAccessSession>>,
        emailHash: ActorCanAccessFunc<AddressAccessSession>;
        expiresAt: ActorCanAccessFunc<AddressAccessSession>;
        ownerId: ActorCanAccessFunc<AddressAccessSession>;
        policyVersionAtIssue: ActorCanAccessFunc<AddressAccessSession>;
        registryId: ActorCanAccessFunc<AddressAccessSession>;
        tokenHash: ActorCanAccessFunc<AddressAccessSession>; 
        updated: ReturnType<typeof NewCanReadActorTrace<AddressAccessSession>>, 
        updatedByOwnerUser: ReturnType<typeof NewCanReadActorTrace<AddressAccessSession>>,
    }
};

const getAbacOwnerId = (obj: AddressAccessSession) => obj.ownerId;

export const canReadAddressAccessSession = NewCanReadAddressAccessSession(
    (actorRoles: ActorRole[], obj?: AddressAccessSession) => {
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

export const canWriteAddressAccessSession = NewCanWriteAddressAccessSession(
    (actorRoles: ActorRole[], obj?: AddressAccessSession) => {
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

export function NewCanReadAddressAccessSession<T = AddressAccessSession>(canAccessObj: ActorCanAccessFunc<T>): canAccessAddressAccessSession<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true,
                approvedGuestId: (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true,
                created:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true),
                emailHash: (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true,
                expiresAt: (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true,
                ownerId: (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true,
                policyVersionAtIssue: (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true,
                registryId: (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true,
                tokenHash: (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true,
                updated:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true),
                updatedByOwnerUser:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true),
            },
        },
    );
}

export function NewCanWriteAddressAccessSession<T = AddressAccessSession>(canAccessObj: ActorCanAccessFunc<T>): canAccessAddressAccessSession<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true,
                approvedGuestId: (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true,
                created:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true),
                emailHash: (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true,
                expiresAt: (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true,
                ownerId: (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true,
                policyVersionAtIssue: (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true,
                registryId: (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true,
                tokenHash: (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true,
                updated:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true),
                updatedByOwnerUser:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: AddressAccessSession) =>  true),
            },
        },
    );
}
