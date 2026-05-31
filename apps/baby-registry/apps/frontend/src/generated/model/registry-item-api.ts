// This file is auto-generated. DO NOT EDIT.

import { RegistryItem, RegistryItemProjection } from './registry-item-model';
import { Registry, RegistryProjection } from './registry-model';
import { Reservation, ReservationProjection } from './reservation-model';
import { ShippingAddressRequest, ShippingAddressRequestProjection } from './shipping-address-request-model';
import { ActorTraceSearchQuery } from './actor-trace-api';
import { ItemSource } from './item-source-enum';

export type RegistryItemWithRefs = {
    registryItem: RegistryItem;
    reservations?: Reservation[];
    shippingAddressRequests?: ShippingAddressRequest[];
    registry?: Registry;
}

export type RegistryItemWithRefsProjection = RegistryItemProjection & {
    Reservations?: ReservationProjection;
    ShippingAddressRequests?: ShippingAddressRequestProjection;
    Registry?: RegistryProjection;
}

export type SelectRegistryItemByIdQuery = {
    id: string;
}

export type RegistryItemSearchQuery = {
    // id (Ref<RegistryItem>) search options
    idEq?: string;
    idIn?: string[];
    idNin?: string[];
    idExists?: boolean;
    // affiliateUrl (string) search options
    affiliateUrlEq?: string;
    affiliateUrlNe?: string;
    affiliateUrlGt?: string;
    affiliateUrlGte?: string;
    affiliateUrlLt?: string;
    affiliateUrlLte?: string;
    affiliateUrlIn?: string[];
    affiliateUrlNin?: string[];
    affiliateUrlExists?: boolean;
    affiliateUrlLike?: string;
    affiliateUrlNlike?: string;
    // canonicalUrl (string) search options
    canonicalUrlEq?: string;
    canonicalUrlNe?: string;
    canonicalUrlGt?: string;
    canonicalUrlGte?: string;
    canonicalUrlLt?: string;
    canonicalUrlLte?: string;
    canonicalUrlIn?: string[];
    canonicalUrlNin?: string[];
    canonicalUrlExists?: boolean;
    canonicalUrlLike?: string;
    canonicalUrlNlike?: string;
    // category (string) search options
    categoryEq?: string;
    categoryNe?: string;
    categoryGt?: string;
    categoryGte?: string;
    categoryLt?: string;
    categoryLte?: string;
    categoryIn?: string[];
    categoryNin?: string[];
    categoryExists?: boolean;
    categoryLike?: string;
    categoryNlike?: string;
    // created (ActorTrace) search options
    created?: ActorTraceSearchQuery;
    // currency (string) search options
    currencyEq?: string;
    currencyNe?: string;
    currencyGt?: string;
    currencyGte?: string;
    currencyLt?: string;
    currencyLte?: string;
    currencyIn?: string[];
    currencyNin?: string[];
    currencyExists?: boolean;
    currencyLike?: string;
    currencyNlike?: string;
    // description (string) search options
    descriptionEq?: string;
    descriptionNe?: string;
    descriptionGt?: string;
    descriptionGte?: string;
    descriptionLt?: string;
    descriptionLte?: string;
    descriptionIn?: string[];
    descriptionNin?: string[];
    descriptionExists?: boolean;
    descriptionLike?: string;
    descriptionNlike?: string;
    // imageBgColor (string) search options
    imageBgColorEq?: string;
    imageBgColorNe?: string;
    imageBgColorGt?: string;
    imageBgColorGte?: string;
    imageBgColorLt?: string;
    imageBgColorLte?: string;
    imageBgColorIn?: string[];
    imageBgColorNin?: string[];
    imageBgColorExists?: boolean;
    imageBgColorLike?: string;
    imageBgColorNlike?: string;
    // imageUrl (string) search options
    imageUrlEq?: string;
    imageUrlNe?: string;
    imageUrlGt?: string;
    imageUrlGte?: string;
    imageUrlLt?: string;
    imageUrlLte?: string;
    imageUrlIn?: string[];
    imageUrlNin?: string[];
    imageUrlExists?: boolean;
    imageUrlLike?: string;
    imageUrlNlike?: string;
    // noSubstitutes (bool) search options
    noSubstitutesEq?: boolean;
    noSubstitutesNe?: boolean;
    noSubstitutesGt?: boolean;
    noSubstitutesGte?: boolean;
    noSubstitutesLt?: boolean;
    noSubstitutesLte?: boolean;
    noSubstitutesIn?: boolean[];
    noSubstitutesNin?: boolean[];
    noSubstitutesExists?: boolean;
    // notes (string) search options
    notesEq?: string;
    notesNe?: string;
    notesGt?: string;
    notesGte?: string;
    notesLt?: string;
    notesLte?: string;
    notesIn?: string[];
    notesNin?: string[];
    notesExists?: boolean;
    notesLike?: string;
    notesNlike?: string;
    // originalUrl (string) search options
    originalUrlEq?: string;
    originalUrlNe?: string;
    originalUrlGt?: string;
    originalUrlGte?: string;
    originalUrlLt?: string;
    originalUrlLte?: string;
    originalUrlIn?: string[];
    originalUrlNin?: string[];
    originalUrlExists?: boolean;
    originalUrlLike?: string;
    originalUrlNlike?: string;
    // ownerPurchased (bool) search options
    ownerPurchasedEq?: boolean;
    ownerPurchasedNe?: boolean;
    ownerPurchasedGt?: boolean;
    ownerPurchasedGte?: boolean;
    ownerPurchasedLt?: boolean;
    ownerPurchasedLte?: boolean;
    ownerPurchasedIn?: boolean[];
    ownerPurchasedNin?: boolean[];
    ownerPurchasedExists?: boolean;
    // parentItemId (Ref<RegistryItem>) search options
    parentItemIdEq?: string;
    parentItemIdIn?: string[];
    parentItemIdNin?: string[];
    parentItemIdExists?: boolean;
    // position (int) search options
    positionEq?: number;
    positionNe?: number;
    positionGt?: number;
    positionGte?: number;
    positionLt?: number;
    positionLte?: number;
    positionIn?: number[];
    positionNin?: number[];
    positionExists?: boolean;
    // priceCents (int) search options
    priceCentsEq?: number;
    priceCentsNe?: number;
    priceCentsGt?: number;
    priceCentsGte?: number;
    priceCentsLt?: number;
    priceCentsLte?: number;
    priceCentsIn?: number[];
    priceCentsNin?: number[];
    priceCentsExists?: boolean;
    // productUrl (string) search options
    productUrlEq?: string;
    productUrlNe?: string;
    productUrlGt?: string;
    productUrlGte?: string;
    productUrlLt?: string;
    productUrlLte?: string;
    productUrlIn?: string[];
    productUrlNin?: string[];
    productUrlExists?: boolean;
    productUrlLike?: string;
    productUrlNlike?: string;
    // quantity (int) search options
    quantityEq?: number;
    quantityNe?: number;
    quantityGt?: number;
    quantityGte?: number;
    quantityLt?: number;
    quantityLte?: number;
    quantityIn?: number[];
    quantityNin?: number[];
    quantityExists?: boolean;
    // quantityUnlimited (bool) search options
    quantityUnlimitedEq?: boolean;
    quantityUnlimitedNe?: boolean;
    quantityUnlimitedGt?: boolean;
    quantityUnlimitedGte?: boolean;
    quantityUnlimitedLt?: boolean;
    quantityUnlimitedLte?: boolean;
    quantityUnlimitedIn?: boolean[];
    quantityUnlimitedNin?: boolean[];
    quantityUnlimitedExists?: boolean;
    // registryId (ParentRef<Registry>) search options
    registryIdEq?: string;
    registryIdIn?: string[];
    registryIdNin?: string[];
    registryIdExists?: boolean;
    // retailer (string) search options
    retailerEq?: string;
    retailerNe?: string;
    retailerGt?: string;
    retailerGte?: string;
    retailerLt?: string;
    retailerLte?: string;
    retailerIn?: string[];
    retailerNin?: string[];
    retailerExists?: boolean;
    retailerLike?: string;
    retailerNlike?: string;
    // source (ItemSource) search options
    sourceEq?: ItemSource;
    sourceNe?: ItemSource;
    sourceGt?: ItemSource;
    sourceGte?: ItemSource;
    sourceLt?: ItemSource;
    sourceLte?: ItemSource;
    sourceIn?: ItemSource[];
    sourceNin?: ItemSource[];
    sourceExists?: boolean;
    // title (string) search options
    titleEq?: string;
    titleNe?: string;
    titleGt?: string;
    titleGte?: string;
    titleLt?: string;
    titleLte?: string;
    titleIn?: string[];
    titleNin?: string[];
    titleExists?: boolean;
    titleLike?: string;
    titleNlike?: string;
    // updated (ActorTrace) search options
    updated?: ActorTraceSearchQuery;
    // updatedByOwnerUser (ActorTrace) search options
    updatedByOwnerUser?: ActorTraceSearchQuery;
}
