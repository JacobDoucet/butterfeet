// This file is auto-generated. DO NOT EDIT.

import { Reservation, ReservationProjection } from './reservation-model';
import { Registry, RegistryProjection } from './registry-model';
import { RegistryItem, RegistryItemProjection } from './registry-item-model';
import { ActorTraceSearchQuery } from './actor-trace-api';
import { ReservationStatus } from './reservation-status-enum';

export type ReservationWithRefs = {
    reservation: Reservation;
    item?: RegistryItem;
    registry?: Registry;
}

export type ReservationWithRefsProjection = ReservationProjection & {
    Item?: RegistryItemProjection;
    Registry?: RegistryProjection;
}

export type SelectReservationByIdQuery = {
    id: string;
}

export type ReservationSearchQuery = {
    // id (Ref<Reservation>) search options
    idEq?: string;
    idIn?: string[];
    idNin?: string[];
    idExists?: boolean;
    // contactEmail (string) search options
    contactEmailEq?: string;
    contactEmailNe?: string;
    contactEmailGt?: string;
    contactEmailGte?: string;
    contactEmailLt?: string;
    contactEmailLte?: string;
    contactEmailIn?: string[];
    contactEmailNin?: string[];
    contactEmailExists?: boolean;
    contactEmailLike?: string;
    contactEmailNlike?: string;
    // created (ActorTrace) search options
    created?: ActorTraceSearchQuery;
    // isAnonymous (bool) search options
    isAnonymousEq?: boolean;
    isAnonymousNe?: boolean;
    isAnonymousGt?: boolean;
    isAnonymousGte?: boolean;
    isAnonymousLt?: boolean;
    isAnonymousLte?: boolean;
    isAnonymousIn?: boolean[];
    isAnonymousNin?: boolean[];
    isAnonymousExists?: boolean;
    // itemId (ParentRef<RegistryItem>) search options
    itemIdEq?: string;
    itemIdIn?: string[];
    itemIdNin?: string[];
    itemIdExists?: boolean;
    // message (string) search options
    messageEq?: string;
    messageNe?: string;
    messageGt?: string;
    messageGte?: string;
    messageLt?: string;
    messageLte?: string;
    messageIn?: string[];
    messageNin?: string[];
    messageExists?: boolean;
    messageLike?: string;
    messageNlike?: string;
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
    // registryId (Ref<Registry>) search options
    registryIdEq?: string;
    registryIdIn?: string[];
    registryIdNin?: string[];
    registryIdExists?: boolean;
    // reserverName (string) search options
    reserverNameEq?: string;
    reserverNameNe?: string;
    reserverNameGt?: string;
    reserverNameGte?: string;
    reserverNameLt?: string;
    reserverNameLte?: string;
    reserverNameIn?: string[];
    reserverNameNin?: string[];
    reserverNameExists?: boolean;
    reserverNameLike?: string;
    reserverNameNlike?: string;
    // status (ReservationStatus) search options
    statusEq?: ReservationStatus;
    statusNe?: ReservationStatus;
    statusGt?: ReservationStatus;
    statusGte?: ReservationStatus;
    statusLt?: ReservationStatus;
    statusLte?: ReservationStatus;
    statusIn?: ReservationStatus[];
    statusNin?: ReservationStatus[];
    statusExists?: boolean;
    // updated (ActorTrace) search options
    updated?: ActorTraceSearchQuery;
    // updatedByOwnerUser (ActorTrace) search options
    updatedByOwnerUser?: ActorTraceSearchQuery;
}
