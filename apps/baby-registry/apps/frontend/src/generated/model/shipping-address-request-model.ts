// This file is auto-generated. DO NOT EDIT.

import { ActorTrace, ActorTraceProjection } from './actor-trace-model';
import { AddressRequestStatus } from './address-request-status-enum';

export type ShippingAddressRequest = {
  id?: string;
  created?: ActorTrace;
  decisionReason?: string;
  emailEnc?: string;
  emailHash?: string;
  name?: string;
  note?: string;
  ownerId?: string;
  policyVersion?: number;
  registryId?: string;
  registryItemId?: string;
  status?: AddressRequestStatus;
  updated?: ActorTrace;
  updatedByOwnerUser?: ActorTrace;
}

export type ShippingAddressRequestProjection = {
    id?: boolean;
    created?: boolean;
		createdFields?: ActorTraceProjection;
    decisionReason?: boolean;
    emailEnc?: boolean;
    emailHash?: boolean;
    name?: boolean;
    note?: boolean;
    ownerId?: boolean;
    policyVersion?: boolean;
    registryId?: boolean;
    registryItemId?: boolean;
    status?: boolean;
    updated?: boolean;
		updatedFields?: ActorTraceProjection;
    updatedByOwnerUser?: boolean;
		updatedByOwnerUserFields?: ActorTraceProjection;
}

export type ShippingAddressRequestSortParams = {
    createdAt?: -1 | 1;
    emailHash?: -1 | 1;
    ownerId?: -1 | 1;
    registryId?: -1 | 1;
    status?: -1 | 1;
    updatedAt?: -1 | 1;
}
