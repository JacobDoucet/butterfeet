// This file is auto-generated. DO NOT EDIT.

import { Registry, RegistryProjection } from './registry-model';
import { AddressAccessSession, AddressAccessSessionProjection } from './address-access-session-model';
import { OwnerUser, OwnerUserProjection } from './owner-user-model';
import { RegistryApprovedGuest, RegistryApprovedGuestProjection } from './registry-approved-guest-model';
import { RegistryItem, RegistryItemProjection } from './registry-item-model';
import { Reservation, ReservationProjection } from './reservation-model';
import { ShippingAddressRequest, ShippingAddressRequestProjection } from './shipping-address-request-model';
import { ActorTraceSearchQuery } from './actor-trace-api';
import { AddressAccessMode } from './address-access-mode-enum';

export type RegistryWithRefs = {
    registry: Registry;
    addressAccessSessions?: AddressAccessSession[];
    registryApprovedGuests?: RegistryApprovedGuest[];
    registryItems?: RegistryItem[];
    reservations?: Reservation[];
    shippingAddressRequests?: ShippingAddressRequest[];
    owner?: OwnerUser;
}

export type RegistryWithRefsProjection = RegistryProjection & {
    AddressAccessSessions?: AddressAccessSessionProjection;
    RegistryApprovedGuests?: RegistryApprovedGuestProjection;
    RegistryItems?: RegistryItemProjection;
    Reservations?: ReservationProjection;
    ShippingAddressRequests?: ShippingAddressRequestProjection;
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
    // addressAccessMode (AddressAccessMode) search options
    addressAccessModeEq?: AddressAccessMode;
    addressAccessModeNe?: AddressAccessMode;
    addressAccessModeGt?: AddressAccessMode;
    addressAccessModeGte?: AddressAccessMode;
    addressAccessModeLt?: AddressAccessMode;
    addressAccessModeLte?: AddressAccessMode;
    addressAccessModeIn?: AddressAccessMode[];
    addressAccessModeNin?: AddressAccessMode[];
    addressAccessModeExists?: boolean;
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
    // shippingCity (string) search options
    shippingCityEq?: string;
    shippingCityNe?: string;
    shippingCityGt?: string;
    shippingCityGte?: string;
    shippingCityLt?: string;
    shippingCityLte?: string;
    shippingCityIn?: string[];
    shippingCityNin?: string[];
    shippingCityExists?: boolean;
    shippingCityLike?: string;
    shippingCityNlike?: string;
    // shippingCountry (string) search options
    shippingCountryEq?: string;
    shippingCountryNe?: string;
    shippingCountryGt?: string;
    shippingCountryGte?: string;
    shippingCountryLt?: string;
    shippingCountryLte?: string;
    shippingCountryIn?: string[];
    shippingCountryNin?: string[];
    shippingCountryExists?: boolean;
    shippingCountryLike?: string;
    shippingCountryNlike?: string;
    // shippingDeliveryNotes (string) search options
    shippingDeliveryNotesEq?: string;
    shippingDeliveryNotesNe?: string;
    shippingDeliveryNotesGt?: string;
    shippingDeliveryNotesGte?: string;
    shippingDeliveryNotesLt?: string;
    shippingDeliveryNotesLte?: string;
    shippingDeliveryNotesIn?: string[];
    shippingDeliveryNotesNin?: string[];
    shippingDeliveryNotesExists?: boolean;
    shippingDeliveryNotesLike?: string;
    shippingDeliveryNotesNlike?: string;
    // shippingLine1 (string) search options
    shippingLine1Eq?: string;
    shippingLine1Ne?: string;
    shippingLine1Gt?: string;
    shippingLine1Gte?: string;
    shippingLine1Lt?: string;
    shippingLine1Lte?: string;
    shippingLine1In?: string[];
    shippingLine1Nin?: string[];
    shippingLine1Exists?: boolean;
    shippingLine1Like?: string;
    shippingLine1Nlike?: string;
    // shippingLine2 (string) search options
    shippingLine2Eq?: string;
    shippingLine2Ne?: string;
    shippingLine2Gt?: string;
    shippingLine2Gte?: string;
    shippingLine2Lt?: string;
    shippingLine2Lte?: string;
    shippingLine2In?: string[];
    shippingLine2Nin?: string[];
    shippingLine2Exists?: boolean;
    shippingLine2Like?: string;
    shippingLine2Nlike?: string;
    // shippingPolicyVersion (int) search options
    shippingPolicyVersionEq?: number;
    shippingPolicyVersionNe?: number;
    shippingPolicyVersionGt?: number;
    shippingPolicyVersionGte?: number;
    shippingPolicyVersionLt?: number;
    shippingPolicyVersionLte?: number;
    shippingPolicyVersionIn?: number[];
    shippingPolicyVersionNin?: number[];
    shippingPolicyVersionExists?: boolean;
    // shippingPostalCode (string) search options
    shippingPostalCodeEq?: string;
    shippingPostalCodeNe?: string;
    shippingPostalCodeGt?: string;
    shippingPostalCodeGte?: string;
    shippingPostalCodeLt?: string;
    shippingPostalCodeLte?: string;
    shippingPostalCodeIn?: string[];
    shippingPostalCodeNin?: string[];
    shippingPostalCodeExists?: boolean;
    shippingPostalCodeLike?: string;
    shippingPostalCodeNlike?: string;
    // shippingRecipientName (string) search options
    shippingRecipientNameEq?: string;
    shippingRecipientNameNe?: string;
    shippingRecipientNameGt?: string;
    shippingRecipientNameGte?: string;
    shippingRecipientNameLt?: string;
    shippingRecipientNameLte?: string;
    shippingRecipientNameIn?: string[];
    shippingRecipientNameNin?: string[];
    shippingRecipientNameExists?: boolean;
    shippingRecipientNameLike?: string;
    shippingRecipientNameNlike?: string;
    // shippingRegion (string) search options
    shippingRegionEq?: string;
    shippingRegionNe?: string;
    shippingRegionGt?: string;
    shippingRegionGte?: string;
    shippingRegionLt?: string;
    shippingRegionLte?: string;
    shippingRegionIn?: string[];
    shippingRegionNin?: string[];
    shippingRegionExists?: boolean;
    shippingRegionLike?: string;
    shippingRegionNlike?: string;
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
