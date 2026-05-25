// This file is auto-generated. DO NOT EDIT.

import { OwnerUser, OwnerUserProjection } from './owner-user-model';
import { Registry, RegistryProjection } from './registry-model';
import { ActorRoleSearchQuery } from './actor-role-api';
import { ActorTraceSearchQuery } from './actor-trace-api';

export type OwnerUserWithRefs = {
    ownerUser: OwnerUser;
    registrys?: Registry[];
}

export type OwnerUserWithRefsProjection = OwnerUserProjection & {
    Registrys?: RegistryProjection;
}

export type SelectOwnerUserByIdQuery = {
    id: string;
}

export type SelectOwnerUserByEmailUniqueQuery = {
    email: string;
}

export type OwnerUserSearchQuery = {
    // id (Ref<OwnerUser>) search options
    idEq?: string;
    idIn?: string[];
    idNin?: string[];
    idExists?: boolean;
    // actorRoles (List<ActorRole>) search options
    actorRoles?: ActorRoleSearchQuery;
    actorRolesEmpty?: boolean;
    // created (ActorTrace) search options
    created?: ActorTraceSearchQuery;
    // email (string) search options
    emailEq?: string;
    emailNe?: string;
    emailGt?: string;
    emailGte?: string;
    emailLt?: string;
    emailLte?: string;
    emailIn?: string[];
    emailNin?: string[];
    emailExists?: boolean;
    emailLike?: string;
    emailNlike?: string;
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
    // updated (ActorTrace) search options
    updated?: ActorTraceSearchQuery;
    // updatedByOwnerUser (ActorTrace) search options
    updatedByOwnerUser?: ActorTraceSearchQuery;
}
