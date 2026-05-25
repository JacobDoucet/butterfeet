// This file is auto-generated. DO NOT EDIT.

import { Role } from './role-enum';

export type ActorRoleSearchQuery = {
    // ownerId (string) search options
    ownerIdEq?: string;
    ownerIdNe?: string;
    ownerIdGt?: string;
    ownerIdGte?: string;
    ownerIdLt?: string;
    ownerIdLte?: string;
    ownerIdIn?: string[];
    ownerIdNin?: string[];
    ownerIdExists?: boolean;
    ownerIdLike?: string;
    ownerIdNlike?: string;
    // role (Role) search options
    roleEq?: Role;
    roleNe?: Role;
    roleGt?: Role;
    roleGte?: Role;
    roleLt?: Role;
    roleLte?: Role;
    roleIn?: Role[];
    roleNin?: Role[];
    roleExists?: boolean;
}
