// This file is auto-generated. DO NOT EDIT.

import { ActorTrace, ActorTraceProjection } from './actor-trace-model';
import { AddressAccessMode } from './address-access-mode-enum';

export type Registry = {
  id?: string;
  addressAccessMode?: AddressAccessMode;
  coverImageUrl?: string;
  created?: ActorTrace;
  dueDate?: string;
  isPublic?: boolean;
  ownerId?: string;
  parentNames?: string;
  shippingCity?: string;
  shippingCountry?: string;
  shippingDeliveryNotes?: string;
  shippingLine1?: string;
  shippingLine2?: string;
  shippingPolicyVersion?: number;
  shippingPostalCode?: string;
  shippingRecipientName?: string;
  shippingRegion?: string;
  slug?: string;
  themeColor?: string;
  title?: string;
  updated?: ActorTrace;
  updatedByOwnerUser?: ActorTrace;
  welcomeMessage?: string;
}

export type RegistryProjection = {
    id?: boolean;
    addressAccessMode?: boolean;
    coverImageUrl?: boolean;
    created?: boolean;
		createdFields?: ActorTraceProjection;
    dueDate?: boolean;
    isPublic?: boolean;
    ownerId?: boolean;
    parentNames?: boolean;
    shippingCity?: boolean;
    shippingCountry?: boolean;
    shippingDeliveryNotes?: boolean;
    shippingLine1?: boolean;
    shippingLine2?: boolean;
    shippingPolicyVersion?: boolean;
    shippingPostalCode?: boolean;
    shippingRecipientName?: boolean;
    shippingRegion?: boolean;
    slug?: boolean;
    themeColor?: boolean;
    title?: boolean;
    updated?: boolean;
		updatedFields?: ActorTraceProjection;
    updatedByOwnerUser?: boolean;
		updatedByOwnerUserFields?: ActorTraceProjection;
    welcomeMessage?: boolean;
}

export type RegistrySortParams = {
    createdAt?: -1 | 1;
    ownerId?: -1 | 1;
    slug?: -1 | 1;
    updatedAt?: -1 | 1;
}
