// This file is auto-generated. DO NOT EDIT.

import { ActorTrace, ActorTraceProjection } from './actor-trace-model';

export type AddressAccessSession = {
  id?: string;
  approvedGuestId?: string;
  created?: ActorTrace;
  emailHash?: string;
  expiresAt?: string;
  ownerId?: string;
  policyVersionAtIssue?: number;
  registryId?: string;
  tokenHash?: string;
  updated?: ActorTrace;
  updatedByOwnerUser?: ActorTrace;
}

export type AddressAccessSessionProjection = {
    id?: boolean;
    approvedGuestId?: boolean;
    created?: boolean;
		createdFields?: ActorTraceProjection;
    emailHash?: boolean;
    expiresAt?: boolean;
    ownerId?: boolean;
    policyVersionAtIssue?: boolean;
    registryId?: boolean;
    tokenHash?: boolean;
    updated?: boolean;
		updatedFields?: ActorTraceProjection;
    updatedByOwnerUser?: boolean;
		updatedByOwnerUserFields?: ActorTraceProjection;
}

export type AddressAccessSessionSortParams = {
    createdAt?: -1 | 1;
    emailHash?: -1 | 1;
    ownerId?: -1 | 1;
    registryId?: -1 | 1;
    tokenHash?: -1 | 1;
    updatedAt?: -1 | 1;
}
