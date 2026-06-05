// This file is auto-generated. DO NOT EDIT.

import { Cart } from '../model/cart-model';
import { NewCanReadActorTrace, NewCanWriteActorTrace } from './actor-trace-can-access';
import { ActorRole } from '../model/actor-role-model';
import { ActorCanAccessFunc } from './actor';

type canAccessCart<T = Cart> = ActorCanAccessFunc<T> & {
    field: {
        id: ActorCanAccessFunc<Cart>;
        amountCents: ActorCanAccessFunc<Cart>;
        claimedAt: ActorCanAccessFunc<Cart>;
        contributorEmail: ActorCanAccessFunc<Cart>;
        contributorName: ActorCanAccessFunc<Cart>; 
        created: ReturnType<typeof NewCanReadActorTrace<Cart>>,
        currency: ActorCanAccessFunc<Cart>;
        decidedAt: ActorCanAccessFunc<Cart>;
        decisionReason: ActorCanAccessFunc<Cart>;
        message: ActorCanAccessFunc<Cart>;
        methodDisplayName: ActorCanAccessFunc<Cart>;
        methodType: ActorCanAccessFunc<Cart>;
        ownerId: ActorCanAccessFunc<Cart>;
        paymentMethodId: ActorCanAccessFunc<Cart>;
        referenceCode: ActorCanAccessFunc<Cart>;
        registryId: ActorCanAccessFunc<Cart>;
        status: ActorCanAccessFunc<Cart>; 
        updated: ReturnType<typeof NewCanReadActorTrace<Cart>>, 
        updatedByOwnerUser: ReturnType<typeof NewCanReadActorTrace<Cart>>,
    }
};

const getAbacOwnerId = (obj: Cart) => obj.ownerId;

export const canReadCart = NewCanReadCart(
    (actorRoles: ActorRole[], obj?: Cart) => {
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

export const canWriteCart = NewCanWriteCart(
    (actorRoles: ActorRole[], obj?: Cart) => {
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

export function NewCanReadCart<T = Cart>(canAccessObj: ActorCanAccessFunc<T>): canAccessCart<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                amountCents: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                claimedAt: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                contributorEmail: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                contributorName: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                created:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Cart) =>  true),
                currency: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                decidedAt: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                decisionReason: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                message: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                methodDisplayName: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                methodType: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                ownerId: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                paymentMethodId: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                referenceCode: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                registryId: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                status: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                updated:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Cart) =>  true),
                updatedByOwnerUser:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Cart) =>  true),
            },
        },
    );
}

export function NewCanWriteCart<T = Cart>(canAccessObj: ActorCanAccessFunc<T>): canAccessCart<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                amountCents: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                claimedAt: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                contributorEmail: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                contributorName: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                created:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Cart) =>  true),
                currency: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                decidedAt: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                decisionReason: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                message: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                methodDisplayName: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                methodType: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                ownerId: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                paymentMethodId: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                referenceCode: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                registryId: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                status: (_actorRoles: ActorRole[], _obj?: Cart) =>  true,
                updated:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Cart) =>  true),
                updatedByOwnerUser:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Cart) =>  true),
            },
        },
    );
}
