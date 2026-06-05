// This file is auto-generated. DO NOT EDIT.

import { ActorTrace, ActorTraceProjection } from './actor-trace-model';
import { PaymentMethodType } from './payment-method-type-enum';

export type RegistryPaymentMethod = {
  id?: string;
  bankAccountName?: string;
  bankAccountNumber?: string;
  bankIban?: string;
  bankName?: string;
  bankRoutingNumber?: string;
  bankSwift?: string;
  created?: ActorTrace;
  displayName?: string;
  enabled?: boolean;
  instructions?: string;
  ownerId?: string;
  paymentUrl?: string;
  position?: number;
  recipientEmail?: string;
  recipientPhone?: string;
  registryId?: string;
  type?: PaymentMethodType;
  updated?: ActorTrace;
  updatedByOwnerUser?: ActorTrace;
}

export type RegistryPaymentMethodProjection = {
    id?: boolean;
    bankAccountName?: boolean;
    bankAccountNumber?: boolean;
    bankIban?: boolean;
    bankName?: boolean;
    bankRoutingNumber?: boolean;
    bankSwift?: boolean;
    created?: boolean;
		createdFields?: ActorTraceProjection;
    displayName?: boolean;
    enabled?: boolean;
    instructions?: boolean;
    ownerId?: boolean;
    paymentUrl?: boolean;
    position?: boolean;
    recipientEmail?: boolean;
    recipientPhone?: boolean;
    registryId?: boolean;
    type?: boolean;
    updated?: boolean;
		updatedFields?: ActorTraceProjection;
    updatedByOwnerUser?: boolean;
		updatedByOwnerUserFields?: ActorTraceProjection;
}

export type RegistryPaymentMethodSortParams = {
    createdAt?: -1 | 1;
    ownerId?: -1 | 1;
    registryId?: -1 | 1;
    updatedAt?: -1 | 1;
}
