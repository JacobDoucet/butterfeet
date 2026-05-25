// This file is auto-generated. DO NOT EDIT.

import { RegistryItem } from '../model/registry-item-model';
import { NewCanReadActorTrace, NewCanWriteActorTrace } from './actor-trace-can-access';
import { ActorRole } from '../model/actor-role-model';
import { ActorCanAccessFunc } from './actor';

type canAccessRegistryItem<T = RegistryItem> = ActorCanAccessFunc<T> & {
    field: {
        id: ActorCanAccessFunc<RegistryItem>; 
        created: ReturnType<typeof NewCanReadActorTrace<RegistryItem>>,
        currency: ActorCanAccessFunc<RegistryItem>;
        description: ActorCanAccessFunc<RegistryItem>;
        imageUrl: ActorCanAccessFunc<RegistryItem>;
        notes: ActorCanAccessFunc<RegistryItem>;
        ownerPurchased: ActorCanAccessFunc<RegistryItem>;
        position: ActorCanAccessFunc<RegistryItem>;
        priceCents: ActorCanAccessFunc<RegistryItem>;
        productUrl: ActorCanAccessFunc<RegistryItem>;
        quantity: ActorCanAccessFunc<RegistryItem>;
        registryId: ActorCanAccessFunc<RegistryItem>;
        source: ActorCanAccessFunc<RegistryItem>;
        title: ActorCanAccessFunc<RegistryItem>; 
        updated: ReturnType<typeof NewCanReadActorTrace<RegistryItem>>, 
        updatedByOwnerUser: ReturnType<typeof NewCanReadActorTrace<RegistryItem>>,
    }
};

export const canReadRegistryItem = NewCanReadRegistryItem(
    (actorRoles: ActorRole[], obj?: RegistryItem) => {
        for (const actorRole of actorRoles) {
            switch(actorRole.role) {
            case 'Owner':
                return true;
            case 'Super':
                return true;
            }
        }
        return false;
    },
);

export const canWriteRegistryItem = NewCanWriteRegistryItem(
    (actorRoles: ActorRole[], obj?: RegistryItem) => {
          for (const actorRole of actorRoles) {
              switch(actorRole.role) {
              case 'Owner':
                  return true;
              case 'Super':
                  return true;
              }
          }
          return false;
    },
);

export function NewCanReadRegistryItem<T = RegistryItem>(canAccessObj: ActorCanAccessFunc<T>): canAccessRegistryItem<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                created:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true),
                currency: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                description: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                imageUrl: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                notes: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                ownerPurchased: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                position: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                priceCents: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                productUrl: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                quantity: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                registryId: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                source: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                title: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                updated:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true),
                updatedByOwnerUser:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true),
            },
        },
    );
}

export function NewCanWriteRegistryItem<T = RegistryItem>(canAccessObj: ActorCanAccessFunc<T>): canAccessRegistryItem<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                created:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true),
                currency: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                description: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                imageUrl: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                notes: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                ownerPurchased: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                position: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                priceCents: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                productUrl: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                quantity: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                registryId: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                source: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                title: (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true,
                updated:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true),
                updatedByOwnerUser:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: RegistryItem) =>  true),
            },
        },
    );
}
