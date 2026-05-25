// This file is auto-generated. DO NOT EDIT.

import { ShippingAddressRequest, ShippingAddressRequestProjection } from './shipping-address-request-model';
import { OwnerUser, OwnerUserProjection } from './owner-user-model';
import { Registry, RegistryProjection } from './registry-model';
import { RegistryItem, RegistryItemProjection } from './registry-item-model';
import { ActorTraceSearchQuery } from './actor-trace-api';
import { AddressRequestStatus } from './address-request-status-enum';

export type ShippingAddressRequestWithRefs = {
    shippingAddressRequest: ShippingAddressRequest;
    owner?: OwnerUser;
    registry?: Registry;
    registryItem?: RegistryItem;
}

export type ShippingAddressRequestWithRefsProjection = ShippingAddressRequestProjection & {
    Owner?: OwnerUserProjection;
    Registry?: RegistryProjection;
    RegistryItem?: RegistryItemProjection;
}

export type SelectShippingAddressRequestByIdQuery = {
    id: string;
}

export type ShippingAddressRequestSearchQuery = {
    // id (Ref<ShippingAddressRequest>) search options
    idEq?: string;
    idIn?: string[];
    idNin?: string[];
    idExists?: boolean;
    // created (ActorTrace) search options
    created?: ActorTraceSearchQuery;
    // decisionReason (string) search options
    decisionReasonEq?: string;
    decisionReasonNe?: string;
    decisionReasonGt?: string;
    decisionReasonGte?: string;
    decisionReasonLt?: string;
    decisionReasonLte?: string;
    decisionReasonIn?: string[];
    decisionReasonNin?: string[];
    decisionReasonExists?: boolean;
    decisionReasonLike?: string;
    decisionReasonNlike?: string;
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
    // note (string) search options
    noteEq?: string;
    noteNe?: string;
    noteGt?: string;
    noteGte?: string;
    noteLt?: string;
    noteLte?: string;
    noteIn?: string[];
    noteNin?: string[];
    noteExists?: boolean;
    noteLike?: string;
    noteNlike?: string;
    // ownerId (Ref<OwnerUser>) search options
    ownerIdEq?: string;
    ownerIdIn?: string[];
    ownerIdNin?: string[];
    ownerIdExists?: boolean;
    // policyVersion (int) search options
    policyVersionEq?: number;
    policyVersionNe?: number;
    policyVersionGt?: number;
    policyVersionGte?: number;
    policyVersionLt?: number;
    policyVersionLte?: number;
    policyVersionIn?: number[];
    policyVersionNin?: number[];
    policyVersionExists?: boolean;
    // registryId (Ref<Registry>) search options
    registryIdEq?: string;
    registryIdIn?: string[];
    registryIdNin?: string[];
    registryIdExists?: boolean;
    // registryItemId (Ref<RegistryItem>) search options
    registryItemIdEq?: string;
    registryItemIdIn?: string[];
    registryItemIdNin?: string[];
    registryItemIdExists?: boolean;
    // status (AddressRequestStatus) search options
    statusEq?: AddressRequestStatus;
    statusNe?: AddressRequestStatus;
    statusGt?: AddressRequestStatus;
    statusGte?: AddressRequestStatus;
    statusLt?: AddressRequestStatus;
    statusLte?: AddressRequestStatus;
    statusIn?: AddressRequestStatus[];
    statusNin?: AddressRequestStatus[];
    statusExists?: boolean;
    // updated (ActorTrace) search options
    updated?: ActorTraceSearchQuery;
    // updatedByOwnerUser (ActorTrace) search options
    updatedByOwnerUser?: ActorTraceSearchQuery;
}
