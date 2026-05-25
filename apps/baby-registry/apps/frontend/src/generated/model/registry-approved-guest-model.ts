// This file is auto-generated. DO NOT EDIT.

import { ActorTrace, ActorTraceProjection } from './actor-trace-model';
import { GuestAccessLevel } from './guest-access-level-enum';
import { GuestStatus } from './guest-status-enum';

export type RegistryApprovedGuest = {
  id?: string;
  accessLevel?: GuestAccessLevel;
  created?: ActorTrace;
  emailEnc?: string;
  emailHash?: string;
  name?: string;
  ownerId?: string;
  registryId?: string;
  status?: GuestStatus;
  updated?: ActorTrace;
  updatedByOwnerUser?: ActorTrace;
}

export type RegistryApprovedGuestProjection = {
    id?: boolean;
    accessLevel?: boolean;
    created?: boolean;
		createdFields?: ActorTraceProjection;
    emailEnc?: boolean;
    emailHash?: boolean;
    name?: boolean;
    ownerId?: boolean;
    registryId?: boolean;
    status?: boolean;
    updated?: boolean;
		updatedFields?: ActorTraceProjection;
    updatedByOwnerUser?: boolean;
		updatedByOwnerUserFields?: ActorTraceProjection;
}

export type RegistryApprovedGuestSortParams = {
    createdAt?: -1 | 1;
    emailHash?: -1 | 1;
    ownerId?: -1 | 1;
    registryId?: -1 | 1;
    updatedAt?: -1 | 1;
}
