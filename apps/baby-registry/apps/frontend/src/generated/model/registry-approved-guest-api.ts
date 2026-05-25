// This file is auto-generated. DO NOT EDIT.

import { RegistryApprovedGuest, RegistryApprovedGuestProjection } from './registry-approved-guest-model';
import { AddressAccessSession, AddressAccessSessionProjection } from './address-access-session-model';
import { OwnerUser, OwnerUserProjection } from './owner-user-model';
import { Registry, RegistryProjection } from './registry-model';
import { ActorTraceSearchQuery } from './actor-trace-api';
import { GuestAccessLevel } from './guest-access-level-enum';
import { GuestStatus } from './guest-status-enum';

export type RegistryApprovedGuestWithRefs = {
    registryApprovedGuest: RegistryApprovedGuest;
    addressAccessSessions?: AddressAccessSession[];
    owner?: OwnerUser;
    registry?: Registry;
}

export type RegistryApprovedGuestWithRefsProjection = RegistryApprovedGuestProjection & {
    AddressAccessSessions?: AddressAccessSessionProjection;
    Owner?: OwnerUserProjection;
    Registry?: RegistryProjection;
}

export type SelectRegistryApprovedGuestByIdQuery = {
    id: string;
}

export type SelectRegistryApprovedGuestByRegistryEmailUniqueQuery = {
    registryId: string;
    emailHash: string;
}

export type RegistryApprovedGuestSearchQuery = {
    // id (Ref<RegistryApprovedGuest>) search options
    idEq?: string;
    idIn?: string[];
    idNin?: string[];
    idExists?: boolean;
    // accessLevel (GuestAccessLevel) search options
    accessLevelEq?: GuestAccessLevel;
    accessLevelNe?: GuestAccessLevel;
    accessLevelGt?: GuestAccessLevel;
    accessLevelGte?: GuestAccessLevel;
    accessLevelLt?: GuestAccessLevel;
    accessLevelLte?: GuestAccessLevel;
    accessLevelIn?: GuestAccessLevel[];
    accessLevelNin?: GuestAccessLevel[];
    accessLevelExists?: boolean;
    // created (ActorTrace) search options
    created?: ActorTraceSearchQuery;
    // emailEnc (string) search options
    emailEncEq?: string;
    emailEncNe?: string;
    emailEncGt?: string;
    emailEncGte?: string;
    emailEncLt?: string;
    emailEncLte?: string;
    emailEncIn?: string[];
    emailEncNin?: string[];
    emailEncExists?: boolean;
    emailEncLike?: string;
    emailEncNlike?: string;
    // emailHash (string) search options
    emailHashEq?: string;
    emailHashNe?: string;
    emailHashGt?: string;
    emailHashGte?: string;
    emailHashLt?: string;
    emailHashLte?: string;
    emailHashIn?: string[];
    emailHashNin?: string[];
    emailHashExists?: boolean;
    emailHashLike?: string;
    emailHashNlike?: string;
    // name (string) search options
    nameEq?: string;
    nameNe?: string;
    nameGt?: string;
    nameGte?: string;
    nameLt?: string;
    nameLte?: string;
    nameIn?: string[];
    nameNin?: string[];
    nameExists?: boolean;
    nameLike?: string;
    nameNlike?: string;
    // ownerId (Ref<OwnerUser>) search options
    ownerIdEq?: string;
    ownerIdIn?: string[];
    ownerIdNin?: string[];
    ownerIdExists?: boolean;
    // registryId (Ref<Registry>) search options
    registryIdEq?: string;
    registryIdIn?: string[];
    registryIdNin?: string[];
    registryIdExists?: boolean;
    // status (GuestStatus) search options
    statusEq?: GuestStatus;
    statusNe?: GuestStatus;
    statusGt?: GuestStatus;
    statusGte?: GuestStatus;
    statusLt?: GuestStatus;
    statusLte?: GuestStatus;
    statusIn?: GuestStatus[];
    statusNin?: GuestStatus[];
    statusExists?: boolean;
    // updated (ActorTrace) search options
    updated?: ActorTraceSearchQuery;
    // updatedByOwnerUser (ActorTrace) search options
    updatedByOwnerUser?: ActorTraceSearchQuery;
}
