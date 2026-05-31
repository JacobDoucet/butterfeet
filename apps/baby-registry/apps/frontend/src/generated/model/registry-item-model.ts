// This file is auto-generated. DO NOT EDIT.

import { ActorTrace, ActorTraceProjection } from './actor-trace-model';
import { ItemSource } from './item-source-enum';

export type RegistryItem = {
  id?: string;
  affiliateUrl?: string;
  canonicalUrl?: string;
  category?: string;
  created?: ActorTrace;
  currency?: string;
  description?: string;
  imageBgColor?: string;
  imageUrl?: string;
  noSubstitutes?: boolean;
  notes?: string;
  originalUrl?: string;
  ownerPurchased?: boolean;
  parentItemId?: string;
  position?: number;
  priceCents?: number;
  productUrl?: string;
  quantity?: number;
  quantityUnlimited?: boolean;
  registryId?: string;
  retailer?: string;
  source?: ItemSource;
  title?: string;
  updated?: ActorTrace;
  updatedByOwnerUser?: ActorTrace;
}

export type RegistryItemProjection = {
    id?: boolean;
    affiliateUrl?: boolean;
    canonicalUrl?: boolean;
    category?: boolean;
    created?: boolean;
		createdFields?: ActorTraceProjection;
    currency?: boolean;
    description?: boolean;
    imageBgColor?: boolean;
    imageUrl?: boolean;
    noSubstitutes?: boolean;
    notes?: boolean;
    originalUrl?: boolean;
    ownerPurchased?: boolean;
    parentItemId?: boolean;
    position?: boolean;
    priceCents?: boolean;
    productUrl?: boolean;
    quantity?: boolean;
    quantityUnlimited?: boolean;
    registryId?: boolean;
    retailer?: boolean;
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
