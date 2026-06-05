// This file is auto-generated. DO NOT EDIT.

import { RegistryPaymentMethod } from '../model/registry-payment-method-model';
import { NewCanReadActorTrace, NewCanWriteActorTrace } from './actor-trace-can-access';
import { ActorRole } from '../model/actor-role-model';
import { ActorCanAccessFunc } from './actor';

type canAccessRegistryPaymentMethod<T = RegistryPaymentMethod> = ActorCanAccessFunc<T> & {
    field: {
        id: ActorCanAccessFunc<RegistryPaymentMethod>;
        bankAccountName: ActorCanAccessFunc<RegistryPaymentMethod>;
        bankAccountNumber: ActorCanAccessFunc<RegistryPaymentMethod>;
        bankIban: ActorCanAccessFunc<RegistryPaymentMethod>;
        bankName: ActorCanAccessFunc<RegistryPaymentMethod>;
        bankRoutingNumber: ActorCanAccessFunc<RegistryPaymentMethod>;
        bankSwift: ActorCanAccessFunc<RegistryPaymentMethod>; 
        created: ReturnType<typeof NewCanReadActorTrace<RegistryPaymentMethod>>,
        displayName: ActorCanAccessFunc<RegistryPaymentMethod>;
        enabled: ActorCanAccessFunc<RegistryPaymentMethod>;
        instructions: ActorCanAccessFunc<RegistryPaymentMethod>;
        ownerId: ActorCanAccessFunc<RegistryPaymentMethod>;
        paymentUrl: ActorCanAccessFunc<RegistryPaymentMethod>;
        position: ActorCanAccessFunc<RegistryPaymentMethod>;
        recipientEmail: ActorCanAccessFunc<RegistryPaymentMethod>;
        recipientPhone: ActorCanAccessFunc<RegistryPaymentMethod>;
        registryId: ActorCanAccessFunc<RegistryPaymentMethod>;
        type: ActorCanAccessFunc<RegistryPaymentMethod>; 
        updated: ReturnType<typeof NewCanReadActorTrace<RegistryPaymentMethod>>, 
        updatedByOwnerUser: ReturnType<typeof NewCanReadActorTrace<RegistryPaymentMethod>>,
    }
};

const getAbacOwnerId = (obj: RegistryPaymentMethod) => obj.ownerId;

export const canReadRegistryPaymentMethod = NewCanReadRegistryPaymentMethod(
    (actorRoles: ActorRole[], obj?: RegistryPaymentMethod) => {
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

export const canWriteRegistryPaymentMethod = NewCanWriteRegistryPaymentMethod(
    (actorRoles: ActorRole[], obj?: RegistryPaymentMethod) => {
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

export function NewCanReadRegistryPaymentMethod<T = RegistryPaymentMethod>(canAccessObj: ActorCanAccessFunc<T>): canAccessRegistryPaymentMethod<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                bankAccountName: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                bankAccountNumber: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                bankIban: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                bankName: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                bankRoutingNumber: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                bankSwift: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                created:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true),
                displayName: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                enabled: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                instructions: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                ownerId: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                paymentUrl: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                position: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                recipientEmail: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                recipientPhone: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                registryId: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                type: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                updated:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true),
                updatedByOwnerUser:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true),
            },
        },
    );
}

export function NewCanWriteRegistryPaymentMethod<T = RegistryPaymentMethod>(canAccessObj: ActorCanAccessFunc<T>): canAccessRegistryPaymentMethod<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                bankAccountName: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                bankAccountNumber: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                bankIban: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                bankName: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                bankRoutingNumber: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                bankSwift: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                created:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true),
                displayName: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                enabled: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                instructions: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                ownerId: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                paymentUrl: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                position: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                recipientEmail: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                recipientPhone: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                registryId: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                type: (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true,
                updated:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true),
                updatedByOwnerUser:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: RegistryPaymentMethod) =>  true),
            },
        },
    );
}
