// This file is auto-generated. DO NOT EDIT.

import { ActorTrace, ActorTraceProjection } from './actor-trace-model';
import { ReservationStatus } from './reservation-status-enum';

export type Reservation = {
  id?: string;
  cartId?: string;
  contactEmail?: string;
  created?: ActorTrace;
  expiresAt?: string;
  isAnonymous?: boolean;
  itemId?: string;
  message?: string;
  quantity?: number;
  registryId?: string;
  reminderSentAt?: string;
  reserverName?: string;
  status?: ReservationStatus;
  updated?: ActorTrace;
  updatedByOwnerUser?: ActorTrace;
}

export type ReservationProjection = {
    id?: boolean;
    cartId?: boolean;
    contactEmail?: boolean;
    created?: boolean;
		createdFields?: ActorTraceProjection;
    expiresAt?: boolean;
    isAnonymous?: boolean;
    itemId?: boolean;
    message?: boolean;
    quantity?: boolean;
    registryId?: boolean;
    reminderSentAt?: boolean;
    reserverName?: boolean;
    status?: boolean;
    updated?: boolean;
		updatedFields?: ActorTraceProjection;
    updatedByOwnerUser?: boolean;
		updatedByOwnerUserFields?: ActorTraceProjection;
}

export type ReservationSortParams = {
    cartId?: -1 | 1;
    createdAt?: -1 | 1;
    expiresAt?: -1 | 1;
    itemId?: -1 | 1;
    registryId?: -1 | 1;
    updatedAt?: -1 | 1;
}
