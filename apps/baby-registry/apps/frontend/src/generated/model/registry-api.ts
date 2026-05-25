// This file is auto-generated. DO NOT EDIT.

import { Registry, RegistryProjection } from './registry-model';
import { OwnerUser, OwnerUserProjection } from './owner-user-model';
import { RegistryItem, RegistryItemProjection } from './registry-item-model';
import { Reservation, ReservationProjection } from './reservation-model';
import { ActorTraceSearchQuery } from './actor-trace-api';

export type RegistryWithRefs = {
    registry: Registry;
    registryItems?: RegistryItem[];
    reservations?: Reservation[];
    owner?: OwnerUser;
}

export type RegistryWithRefsProjection = RegistryProjection & {
    RegistryItems?: RegistryItemProjection;
    Reservations?: ReservationProjection;
    Owner?: OwnerUserProjection;
}

export type SelectRegistryByIdQuery = {
    id: string;
}

export type SelectRegistryBySlugUniqueQuery = {
    slug: string;
}

export type RegistrySearchQuery = {
    // id (Ref<Registry>) search options
    idEq?: string;
    idIn?: string[];
    idNin?: string[];
    idExists?: boolean;
    // coverImageUrl (string) search options
    coverImageUrlEq?: string;
    coverImageUrlNe?: string;
    coverImageUrlGt?: string;
    coverImageUrlGte?: string;
    coverImageUrlLt?: string;
    coverImageUrlLte?: string;
    coverImageUrlIn?: string[];
    coverImageUrlNin?: string[];
    coverImageUrlExists?: boolean;
    coverImageUrlLike?: string;
    coverImageUrlNlike?: string;
    // created (ActorTrace) search options
    created?: ActorTraceSearchQuery;
    // dueDate (timestamp) search options
    dueDateEq?: string;
    dueDateNe?: string;
    dueDateGt?: string;
    dueDateGte?: string;
    dueDateLt?: string;
    dueDateLte?: string;
    dueDateIn?: string[];
    dueDateNin?: string[];
    dueDateExists?: boolean;
    // isPublic (bool) search options
    isPublicEq?: boolean;
    isPublicNe?: boolean;
    isPublicGt?: boolean;
    isPublicGte?: boolean;
    isPublicLt?: boolean;
    isPublicLte?: boolean;
    isPublicIn?: boolean[];
    isPublicNin?: boolean[];
    isPublicExists?: boolean;
    // ownerId (Ref<OwnerUser>) search options
    ownerIdEq?: string;
    ownerIdIn?: string[];
    ownerIdNin?: string[];
    ownerIdExists?: boolean;
    // parentNames (string) search options
    parentNamesEq?: string;
    parentNamesNe?: string;
    parentNamesGt?: string;
    parentNamesGte?: string;
    parentNamesLt?: string;
    parentNamesLte?: string;
    parentNamesIn?: string[];
    parentNamesNin?: string[];
    parentNamesExists?: boolean;
    parentNamesLike?: string;
    parentNamesNlike?: string;
    // slug (string) search options
    slugEq?: string;
    slugNe?: string;
    slugGt?: string;
    slugGte?: string;
    slugLt?: string;
    slugLte?: string;
    slugIn?: string[];
    slugNin?: string[];
    slugExists?: boolean;
    slugLike?: string;
    slugNlike?: string;
    // themeColor (string) search options
    themeColorEq?: string;
    themeColorNe?: string;
    themeColorGt?: string;
    themeColorGte?: string;
    themeColorLt?: string;
    themeColorLte?: string;
    themeColorIn?: string[];
    themeColorNin?: string[];
    themeColorExists?: boolean;
    themeColorLike?: string;
    themeColorNlike?: string;
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
    // welcomeMessage (string) search options
    welcomeMessageEq?: string;
    welcomeMessageNe?: string;
    welcomeMessageGt?: string;
    welcomeMessageGte?: string;
    welcomeMessageLt?: string;
    welcomeMessageLte?: string;
    welcomeMessageIn?: string[];
    welcomeMessageNin?: string[];
    welcomeMessageExists?: boolean;
    welcomeMessageLike?: string;
    welcomeMessageNlike?: string;
}
