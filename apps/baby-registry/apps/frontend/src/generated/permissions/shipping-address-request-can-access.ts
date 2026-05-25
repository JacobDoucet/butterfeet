// This file is auto-generated. DO NOT EDIT.

import { ShippingAddressRequest } from '../model/shipping-address-request-model';
import { NewCanReadActorTrace, NewCanWriteActorTrace } from './actor-trace-can-access';
import { ActorRole } from '../model/actor-role-model';
import { ActorCanAccessFunc } from './actor';

type canAccessShippingAddressRequest<T = ShippingAddressRequest> = ActorCanAccessFunc<T> & {
    field: {
        id: ActorCanAccessFunc<ShippingAddressRequest>; 
        created: ReturnType<typeof NewCanReadActorTrace<ShippingAddressRequest>>,
        decisionReason: ActorCanAccessFunc<ShippingAddressRequest>;
        emailEnc: ActorCanAccessFunc<ShippingAddressRequest>;
        emailHash: ActorCanAccessFunc<ShippingAddressRequest>;
        name: ActorCanAccessFunc<ShippingAddressRequest>;
        note: ActorCanAccessFunc<ShippingAddressRequest>;
        ownerId: ActorCanAccessFunc<ShippingAddressRequest>;
        policyVersion: ActorCanAccessFunc<ShippingAddressRequest>;
        registryId: ActorCanAccessFunc<ShippingAddressRequest>;
        registryItemId: ActorCanAccessFunc<ShippingAddressRequest>;
        status: ActorCanAccessFunc<ShippingAddressRequest>; 
        updated: ReturnType<typeof NewCanReadActorTrace<ShippingAddressRequest>>, 
        updatedByOwnerUser: ReturnType<typeof NewCanReadActorTrace<ShippingAddressRequest>>,
    }
};

const getAbacOwnerId = (obj: ShippingAddressRequest) => obj.ownerId;

export const canReadShippingAddressRequest = NewCanReadShippingAddressRequest(
    (actorRoles: ActorRole[], obj?: ShippingAddressRequest) => {
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

export const canWriteShippingAddressRequest = NewCanWriteShippingAddressRequest(
    (actorRoles: ActorRole[], obj?: ShippingAddressRequest) => {
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

export function NewCanReadShippingAddressRequest<T = ShippingAddressRequest>(canAccessObj: ActorCanAccessFunc<T>): canAccessShippingAddressRequest<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                created:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true),
                decisionReason: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                emailEnc: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                emailHash: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                name: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                note: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                ownerId: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                policyVersion: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                registryId: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                registryItemId: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                status: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                updated:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true),
                updatedByOwnerUser:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true),
            },
        },
    );
}

export function NewCanWriteShippingAddressRequest<T = ShippingAddressRequest>(canAccessObj: ActorCanAccessFunc<T>): canAccessShippingAddressRequest<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                created:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true),
                decisionReason: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                emailEnc: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                emailHash: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                name: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                note: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                ownerId: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                policyVersion: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                registryId: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                registryItemId: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                status: (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true,
                updated:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true),
                updatedByOwnerUser:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: ShippingAddressRequest) =>  true),
            },
        },
    );
}
