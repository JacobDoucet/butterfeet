// This file is auto-generated. DO NOT EDIT.

import { Reservation } from '../model/reservation-model';
import { NewCanReadActorTrace, NewCanWriteActorTrace } from './actor-trace-can-access';
import { ActorRole } from '../model/actor-role-model';
import { ActorCanAccessFunc } from './actor';

type canAccessReservation<T = Reservation> = ActorCanAccessFunc<T> & {
    field: {
        id: ActorCanAccessFunc<Reservation>;
        contactEmail: ActorCanAccessFunc<Reservation>; 
        created: ReturnType<typeof NewCanReadActorTrace<Reservation>>,
        isAnonymous: ActorCanAccessFunc<Reservation>;
        itemId: ActorCanAccessFunc<Reservation>;
        message: ActorCanAccessFunc<Reservation>;
        quantity: ActorCanAccessFunc<Reservation>;
        registryId: ActorCanAccessFunc<Reservation>;
        reserverName: ActorCanAccessFunc<Reservation>;
        status: ActorCanAccessFunc<Reservation>; 
        updated: ReturnType<typeof NewCanReadActorTrace<Reservation>>, 
        updatedByOwnerUser: ReturnType<typeof NewCanReadActorTrace<Reservation>>,
    }
};

export const canReadReservation = NewCanReadReservation(
    (actorRoles: ActorRole[], obj?: Reservation) => {
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

export const canWriteReservation = NewCanWriteReservation(
    (actorRoles: ActorRole[], obj?: Reservation) => {
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

export function NewCanReadReservation<T = Reservation>(canAccessObj: ActorCanAccessFunc<T>): canAccessReservation<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: Reservation) =>  true,
                contactEmail: (_actorRoles: ActorRole[], _obj?: Reservation) =>  true,
                created:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Reservation) =>  true),
                isAnonymous: (_actorRoles: ActorRole[], _obj?: Reservation) =>  true,
                itemId: (_actorRoles: ActorRole[], _obj?: Reservation) =>  true,
                message: (_actorRoles: ActorRole[], _obj?: Reservation) =>  true,
                quantity: (_actorRoles: ActorRole[], _obj?: Reservation) =>  true,
                registryId: (_actorRoles: ActorRole[], _obj?: Reservation) =>  true,
                reserverName: (_actorRoles: ActorRole[], _obj?: Reservation) =>  true,
                status: (_actorRoles: ActorRole[], _obj?: Reservation) =>  true,
                updated:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Reservation) =>  true),
                updatedByOwnerUser:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Reservation) =>  true),
            },
        },
    );
}

export function NewCanWriteReservation<T = Reservation>(canAccessObj: ActorCanAccessFunc<T>): canAccessReservation<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: Reservation) =>  true,
                contactEmail: (_actorRoles: ActorRole[], _obj?: Reservation) =>  true,
                created:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Reservation) =>  true),
                isAnonymous: (_actorRoles: ActorRole[], _obj?: Reservation) =>  true,
                itemId: (_actorRoles: ActorRole[], _obj?: Reservation) =>  true,
                message: (_actorRoles: ActorRole[], _obj?: Reservation) =>  true,
                quantity: (_actorRoles: ActorRole[], _obj?: Reservation) =>  true,
                registryId: (_actorRoles: ActorRole[], _obj?: Reservation) =>  true,
                reserverName: (_actorRoles: ActorRole[], _obj?: Reservation) =>  true,
                status: (_actorRoles: ActorRole[], _obj?: Reservation) =>  true,
                updated:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Reservation) =>  true),
                updatedByOwnerUser:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Reservation) =>  true),
            },
        },
    );
}
