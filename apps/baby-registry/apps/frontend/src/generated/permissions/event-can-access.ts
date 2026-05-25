// This file is auto-generated. DO NOT EDIT.

import { Event } from '../model/event-model';
import { NewCanReadActorTrace, NewCanWriteActorTrace } from './actor-trace-can-access';
import { NewCanReadEventSubject, NewCanWriteEventSubject } from './event-subject-can-access';
import { ActorRole } from '../model/actor-role-model';
import { ActorCanAccessFunc } from './actor';

type canAccessEvent<T = Event> = ActorCanAccessFunc<T> & {
    field: {
        id: ActorCanAccessFunc<Event>; 
        created: ReturnType<typeof NewCanReadActorTrace<Event>>, 
        subjects: ReturnType<typeof NewCanReadEventSubject<Event>>,
        type: ActorCanAccessFunc<Event>; 
        updated: ReturnType<typeof NewCanReadActorTrace<Event>>, 
        updatedByOwnerUser: ReturnType<typeof NewCanReadActorTrace<Event>>,
    }
};

export const canReadEvent = NewCanReadEvent(
    (actorRoles: ActorRole[], obj?: Event) => {
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

export const canWriteEvent = NewCanWriteEvent(
    (actorRoles: ActorRole[], obj?: Event) => {
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

export function NewCanReadEvent<T = Event>(canAccessObj: ActorCanAccessFunc<T>): canAccessEvent<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: Event) =>  true,
                created:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Event) =>  true),
                subjects:  NewCanReadEventSubject( (_actorRoles: ActorRole[], _obj?: Event) =>  true),
                type: (_actorRoles: ActorRole[], _obj?: Event) =>  true,
                updated:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Event) =>  true),
                updatedByOwnerUser:  NewCanReadActorTrace( (_actorRoles: ActorRole[], _obj?: Event) =>  true),
            },
        },
    );
}

export function NewCanWriteEvent<T = Event>(canAccessObj: ActorCanAccessFunc<T>): canAccessEvent<T> {
    return Object.assign(
        function (actorRoles: ActorRole[], obj?: T) {
            return canAccessObj(actorRoles, obj);
        },
        {
            field: {
                id: (_actorRoles: ActorRole[], _obj?: Event) =>  true,
                created:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Event) =>  true),
                subjects:  NewCanWriteEventSubject( (_actorRoles: ActorRole[], _obj?: Event) =>  true),
                type: (_actorRoles: ActorRole[], _obj?: Event) =>  true,
                updated:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Event) =>  true),
                updatedByOwnerUser:  NewCanWriteActorTrace( (_actorRoles: ActorRole[], _obj?: Event) =>  true),
            },
        },
    );
}
