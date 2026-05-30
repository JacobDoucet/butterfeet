// This file is auto-generated. DO NOT EDIT.

import { ActorTrace, ActorTraceProjection } from './actor-trace-model';
import { ItemSource } from './item-source-enum';

export type RegistryItem = {
  id?: string;
  category?: string;
  created?: ActorTrace;
  currency?: string;
  description?: string;
  imageUrl?: string;
  noSubstitutes?: boolean;
  notes?: string;
  ownerPurchased?: boolean;
  parentItemId?: string;
  position?: number;
  priceCents?: number;
  productUrl?: string;
  quantity?: number;
  quantityUnlimited?: boolean;
  registryId?: string;
  source?: ItemSource;
  title?: string;
  updated?: ActorTrace;
  updatedByOwnerUser?: ActorTrace;
}

export type RegistryItemProjection = {
    id?: boolean;
    category?: boolean;
    created?: boolean;
		createdFields?: ActorTraceProjection;
    currency?: boolean;
    description?: boolean;
    imageUrl?: boolean;
    noSubstitutes?: boolean;
    notes?: boolean;
    ownerPurchased?: boolean;
    parentItemId?: boolean;
    position?: boolean;
    priceCents?: boolean;
    productUrl?: boolean;
    quantity?: boolean;
    quantityUnlimited?: boolean;
    registryId?: boolean;
    source?: boolean;
    title?: boolean;
    updated?: boolean;
		updatedFields?: ActorTraceProjection;
    updatedByOwnerUser?: boolean;
		updatedByOwnerUserFields?: ActorTraceProjection;
}

export type RegistryItemSortParams = {
    createdAt?: -1 | 1;
    position?: -1 | 1;
    registryId?: -1 | 1;
    updatedAt?: -1 | 1;
}
