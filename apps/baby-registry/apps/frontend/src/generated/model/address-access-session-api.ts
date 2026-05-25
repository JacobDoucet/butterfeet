// This file is auto-generated. DO NOT EDIT.

import { AddressAccessSession, AddressAccessSessionProjection } from './address-access-session-model';
import { OwnerUser, OwnerUserProjection } from './owner-user-model';
import { Registry, RegistryProjection } from './registry-model';
import { RegistryApprovedGuest, RegistryApprovedGuestProjection } from './registry-approved-guest-model';
import { ActorTraceSearchQuery } from './actor-trace-api';

export type AddressAccessSessionWithRefs = {
    addressAccessSession: AddressAccessSession;
    approvedGuest?: RegistryApprovedGuest;
    owner?: OwnerUser;
    registry?: Registry;
}

export type AddressAccessSessionWithRefsProjection = AddressAccessSessionProjection & {
    ApprovedGuest?: RegistryApprovedGuestProjection;
    Owner?: OwnerUserProjection;
    Registry?: RegistryProjection;
}

export type SelectAddressAccessSessionByIdQuery = {
    id: string;
}

export type SelectAddressAccessSessionByTokenUniqueQuery = {
    tokenHash: string;
}

export type AddressAccessSessionSearchQuery = {
    // id (Ref<AddressAccessSession>) search options
    idEq?: string;
    idIn?: string[];
    idNin?: string[];
    idExists?: boolean;
    // approvedGuestId (Ref<RegistryApprovedGuest>) search options
    approvedGuestIdEq?: string;
    approvedGuestIdIn?: string[];
    approvedGuestIdNin?: string[];
    approvedGuestIdExists?: boolean;
    // created (ActorTrace) search options
    created?: ActorTraceSearchQuery;
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
    // expiresAt (timestamp) search options
    expiresAtEq?: string;
    expiresAtNe?: string;
    expiresAtGt?: string;
    expiresAtGte?: string;
    expiresAtLt?: string;
    expiresAtLte?: string;
    expiresAtIn?: string[];
    expiresAtNin?: string[];
    expiresAtExists?: boolean;
    // ownerId (Ref<OwnerUser>) search options
    ownerIdEq?: string;
    ownerIdIn?: string[];
    ownerIdNin?: string[];
    ownerIdExists?: boolean;
    // policyVersionAtIssue (int) search options
    policyVersionAtIssueEq?: number;
    policyVersionAtIssueNe?: number;
    policyVersionAtIssueGt?: number;
    policyVersionAtIssueGte?: number;
    policyVersionAtIssueLt?: number;
    policyVersionAtIssueLte?: number;
    policyVersionAtIssueIn?: number[];
    policyVersionAtIssueNin?: number[];
    policyVersionAtIssueExists?: boolean;
    // registryId (Ref<Registry>) search options
    registryIdEq?: string;
    registryIdIn?: string[];
    registryIdNin?: string[];
    registryIdExists?: boolean;
    // tokenHash (string) search options
    tokenHashEq?: string;
    tokenHashNe?: string;
    tokenHashGt?: string;
    tokenHashGte?: string;
    tokenHashLt?: string;
    tokenHashLte?: string;
    tokenHashIn?: string[];
    tokenHashNin?: string[];
    tokenHashExists?: boolean;
    tokenHashLike?: string;
    tokenHashNlike?: string;
    // updated (ActorTrace) search options
    updated?: ActorTraceSearchQuery;
    // updatedByOwnerUser (ActorTrace) search options
    updatedByOwnerUser?: ActorTraceSearchQuery;
}
