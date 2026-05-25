// This file is auto-generated. DO NOT EDIT.

import { RegistryItem, RegistryItemProjection } from './registry-item-model';
import { Registry, RegistryProjection } from './registry-model';
import { Reservation, ReservationProjection } from './reservation-model';
import { ActorTraceSearchQuery } from './actor-trace-api';
import { ItemSource } from './item-source-enum';

export type RegistryItemWithRefs = {
    registryItem: RegistryItem;
    reservations?: Reservation[];
    registry?: Registry;
}

export type RegistryItemWithRefsProjection = RegistryItemProjection & {
    Reservations?: ReservationProjection;
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
    // registryId (ParentRef<Registry>) search options
    registryIdEq?: string;
    registryIdIn?: string[];
    registryIdNin?: string[];
    registryIdExists?: boolean;
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
