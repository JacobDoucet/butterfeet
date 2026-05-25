// This file is auto-generated. DO NOT EDIT.

import { AddressAccessSessionSearchQuery, AddressAccessSessionWithRefs, AddressAccessSessionWithRefsProjection } from '../model/address-access-session-api';
import { SelectManyResponse, MutationResponse, AggregateResponse } from './model';
import { ApiError } from './errors';
import { AddressAccessSession, AddressAccessSessionSortParams } from '../model/address-access-session-model';
import { OwnerUser, OwnerUserProjection } from '../model/owner-user-model';
import { Registry, RegistryProjection } from '../model/registry-model';
import { RegistryApprovedGuest, RegistryApprovedGuestProjection } from '../model/registry-approved-guest-model';

export type SearchAddressAccessSessionsParams = {
    baseUrl: string;
    query: AddressAccessSessionSearchQuery;
    sort?: AddressAccessSessionSortParams;
    projection?: AddressAccessSessionWithRefsProjection;
    limit?: number;
    skip?: number;
}

async function newApiError(response: Response, defaultText: string): Promise<ApiError> {
    let text = defaultText;
    try {
        text = await response.text();
    } catch(_) {}
    return new ApiError(text);
}

export function searchAddressAccessSessions(params: SearchAddressAccessSessionsParams): Promise<SelectManyResponse<AddressAccessSessionWithRefs>> {
    return fetch(`${params.baseUrl}/address-access-sessions/search`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            query: params.query,
            sort: params.sort,
            projection: params.projection,
            limit: params.limit,
            skip: params.skip,
        }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to search AddressAccessSession');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type SelectAddressAccessSessionByIdParams = {
    baseUrl: string;
    id: string;
    projection?: AddressAccessSessionWithRefsProjection;
}

export function selectAddressAccessSessionById(params: SelectAddressAccessSessionByIdParams): Promise<AddressAccessSessionWithRefs> {
    return fetch(`${params.baseUrl}/address-access-sessions/id/${params.id}`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              projection: params.projection,
          }),
    }).then(async (response) => {
          if (!response.ok) {
              const err = await newApiError(response, 'Failed to select AddressAccessSession');
              return Promise.reject(err);
          }
          return response.json();
     });
}

export type SelectAddressAccessSessionByTokenUniqueParams = {
    baseUrl: string;
    tokenHash: string;
    projection?: AddressAccessSessionWithRefsProjection;
}

export function selectAddressAccessSessionByTokenUnique(params: SelectAddressAccessSessionByTokenUniqueParams): Promise<AddressAccessSessionWithRefs> {
    return fetch(`${params.baseUrl}/address-access-sessions/tokenHash/${params.tokenHash}`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              projection: params.projection,
          }),
    }).then(async (response) => {
          if (!response.ok) {
              const err = await newApiError(response, 'Failed to select AddressAccessSession');
              return Promise.reject(err);
          }
          return response.json();
     });
}

export type SaveAddressAccessSessionParams = {
    baseUrl: string;
    data: AddressAccessSession;
}

export function createAddressAccessSession(params: SaveAddressAccessSessionParams): Promise<MutationResponse<AddressAccessSession>> {
    return fetch(`${params.baseUrl}/address-access-sessions/create`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to create AddressAccessSession');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export function updateAddressAccessSession(params: SaveAddressAccessSessionParams): Promise<MutationResponse<AddressAccessSession>> {
    return fetch(`${params.baseUrl}/address-access-sessions/update`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to update AddressAccessSession');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type DeleteAddressAccessSessionParams = {
    baseUrl: string;
    id: string;
}

export function deleteAddressAccessSession({ baseUrl, id }: DeleteAddressAccessSessionParams): Promise<void> {
    return fetch(`${baseUrl}/address-access-sessions/delete/${id}`, {
        method: 'DELETE',
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to delete AddressAccessSession');
            return Promise.reject(err);
        }
        return;
    });
}

// Aggregation types
export type AggregateMethod = 'sum' | 'avg' | 'min' | 'max' | 'count' | 'first' | 'last';

// Type-safe aggregatable fields
export const AddressAccessSessionAggregateFields = {
    PolicyVersionAtIssue: 'policyVersionAtIssue',
} as const;

export type AddressAccessSessionAggregateField = typeof AddressAccessSessionAggregateFields[keyof typeof AddressAccessSessionAggregateFields];

// Type-safe group-by fields
export const AddressAccessSessionGroupByFields = {
    ApprovedGuestId: 'approvedGuestId',
    EmailHash: 'emailHash',
    ExpiresAt: 'expiresAt',
    OwnerId: 'ownerId',
    PolicyVersionAtIssue: 'policyVersionAtIssue',
    RegistryId: 'registryId',
    TokenHash: 'tokenHash',
} as const;

export type AddressAccessSessionGroupByField = typeof AddressAccessSessionGroupByFields[keyof typeof AddressAccessSessionGroupByFields];

export type AggregateFieldSpec = {
    field: AddressAccessSessionAggregateField;
    method: AggregateMethod;
    alias?: string;
}

// Aggregate result row with partial model fields and metadata
export type AddressAccessSessionAggregateResultRow = {
    // Group-by fields (original types)
    approvedGuestId?: string | null;
    emailHash?: string | null;
    expiresAt?: string | null;
    ownerId?: string | null;
    policyVersionAtIssue?: number | null;
    registryId?: string | null;
    tokenHash?: string | null;
    // Aggregate fields - always numbers since they're results of sum/avg/etc
    // Ref field approvedGuest
    approvedGuest?: RegistryApprovedGuest | null;
    // Ref field owner
    owner?: OwnerUser | null;
    // Ref field registry
    registry?: Registry | null;
    // Metadata indicating which fields are populated
    __groupKeys: AddressAccessSessionGroupByField[];
    __aggregateKeys: string[];
}

export type AddressAccessSessionAggregateResponse = {
    data: AddressAccessSessionAggregateResultRow[];
    total: number;
}

export type AggregateAddressAccessSessionParams = {
    baseUrl: string;
    query: AddressAccessSessionSearchQuery;
    fields: AggregateFieldSpec[];
    groupBy: AddressAccessSessionGroupByField[];
    approvedGuestProjection?: RegistryApprovedGuestProjection;
    ownerProjection?: OwnerUserProjection;
    registryProjection?: RegistryProjection;
}

export function aggregateAddressAccessSessions(params: AggregateAddressAccessSessionParams): Promise<AddressAccessSessionAggregateResponse> {
    return fetch(`${params.baseUrl}/address-access-sessions/aggregate`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            query: params.query,
            fields: params.fields,
            groupBy: params.groupBy,
            approvedGuestProjection: params.approvedGuestProjection,
            ownerProjection: params.ownerProjection,
            registryProjection: params.registryProjection,
        }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to aggregate AddressAccessSession');
            return Promise.reject(err);
        }
        return response.json();
    });
}
