// This file is auto-generated. DO NOT EDIT.

import { ActorRole, ActorRoleProjection } from './actor-role-model';
import { ActorTrace, ActorTraceProjection } from './actor-trace-model';

export type OwnerUser = {
  id?: string;
  actorRoles?: ActorRole[];
  created?: ActorTrace;
  email?: string;
  name?: string;
  updated?: ActorTrace;
  updatedByOwnerUser?: ActorTrace;
}

export function ownerUserAsActor(ownerUser: OwnerUser) {

}

export type OwnerUserProjection = {
    id?: boolean;
    actorRoles?: boolean;
		actorRolesFields?: ActorRoleProjection;
    created?: boolean;
		createdFields?: ActorTraceProjection;
    email?: boolean;
    name?: boolean;
    updated?: boolean;
		updatedFields?: ActorTraceProjection;
    updatedByOwnerUser?: boolean;
		updatedByOwnerUserFields?: ActorTraceProjection;
}

export type OwnerUserSortParams = {
    createdAt?: -1 | 1;
    email?: -1 | 1;
    updatedAt?: -1 | 1;
}
