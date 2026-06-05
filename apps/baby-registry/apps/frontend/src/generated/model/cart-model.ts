// This file is auto-generated. DO NOT EDIT.

import { ActorTrace, ActorTraceProjection } from './actor-trace-model';
import { PaymentMethodType } from './payment-method-type-enum';
import { CartStatus } from './cart-status-enum';

export type Cart = {
  id?: string;
  amountCents?: number;
  claimedAt?: string;
  contributorEmail?: string;
  contributorName?: string;
  created?: ActorTrace;
  currency?: string;
  decidedAt?: string;
  decisionReason?: string;
  message?: string;
  methodDisplayName?: string;
  methodType?: PaymentMethodType;
  ownerId?: string;
  paymentMethodId?: string;
  referenceCode?: string;
  registryId?: string;
  status?: CartStatus;
  updated?: ActorTrace;
  updatedByOwnerUser?: ActorTrace;
}

export type CartProjection = {
    id?: boolean;
    amountCents?: boolean;
    claimedAt?: boolean;
    contributorEmail?: boolean;
    contributorName?: boolean;
    created?: boolean;
		createdFields?: ActorTraceProjection;
    currency?: boolean;
    decidedAt?: boolean;
    decisionReason?: boolean;
    message?: boolean;
    methodDisplayName?: boolean;
    methodType?: boolean;
    ownerId?: boolean;
    paymentMethodId?: boolean;
    referenceCode?: boolean;
    registryId?: boolean;
    status?: boolean;
    updated?: boolean;
		updatedFields?: ActorTraceProjection;
    updatedByOwnerUser?: boolean;
		updatedByOwnerUserFields?: ActorTraceProjection;
}

export type CartSortParams = {
    contributorEmail?: -1 | 1;
    createdAt?: -1 | 1;
    ownerId?: -1 | 1;
    referenceCode?: -1 | 1;
    registryId?: -1 | 1;
    updatedAt?: -1 | 1;
}
