// This file is auto-generated. DO NOT EDIT.

import { ShippingAddressRequestSearchQuery, ShippingAddressRequestWithRefs, ShippingAddressRequestWithRefsProjection } from '../model/shipping-address-request-api';
import { SelectManyResponse, MutationResponse, AggregateResponse } from './model';
import { ApiError } from './errors';
import { ShippingAddressRequest, ShippingAddressRequestSortParams } from '../model/shipping-address-request-model';
import { OwnerUser, OwnerUserProjection } from '../model/owner-user-model';
import { Registry, RegistryProjection } from '../model/registry-model';
import { RegistryItem, RegistryItemProjection } from '../model/registry-item-model';

export type SearchShippingAddressRequestsParams = {
    baseUrl: string;
    query: ShippingAddressRequestSearchQuery;
    sort?: ShippingAddressRequestSortParams;
    projection?: ShippingAddressRequestWithRefsProjection;
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

export function searchShippingAddressRequests(params: SearchShippingAddressRequestsParams): Promise<SelectManyResponse<ShippingAddressRequestWithRefs>> {
    return fetch(`${params.baseUrl}/shipping-address-requests/search`, {
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
            const err = await newApiError(response, 'Failed to search ShippingAddressRequest');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type SelectShippingAddressRequestByIdParams = {
    baseUrl: string;
    id: string;
    projection?: ShippingAddressRequestWithRefsProjection;
}

export function selectShippingAddressRequestById(params: SelectShippingAddressRequestByIdParams): Promise<ShippingAddressRequestWithRefs> {
    return fetch(`${params.baseUrl}/shipping-address-requests/id/${params.id}`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              projection: params.projection,
          }),
    }).then(async (response) => {
          if (!response.ok) {
              const err = await newApiError(response, 'Failed to select ShippingAddressRequest');
              return Promise.reject(err);
          }
          return response.json();
     });
}

export type SaveShippingAddressRequestParams = {
    baseUrl: string;
    data: ShippingAddressRequest;
}

export function createShippingAddressRequest(params: SaveShippingAddressRequestParams): Promise<MutationResponse<ShippingAddressRequest>> {
    return fetch(`${params.baseUrl}/shipping-address-requests/create`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to create ShippingAddressRequest');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export function updateShippingAddressRequest(params: SaveShippingAddressRequestParams): Promise<MutationResponse<ShippingAddressRequest>> {
    return fetch(`${params.baseUrl}/shipping-address-requests/update`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to update ShippingAddressRequest');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type DeleteShippingAddressRequestParams = {
    baseUrl: string;
    id: string;
}

export function deleteShippingAddressRequest({ baseUrl, id }: DeleteShippingAddressRequestParams): Promise<void> {
    return fetch(`${baseUrl}/shipping-address-requests/delete/${id}`, {
        method: 'DELETE',
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to delete ShippingAddressRequest');
            return Promise.reject(err);
        }
        return;
    });
}

// Aggregation types
export type AggregateMethod = 'sum' | 'avg' | 'min' | 'max' | 'count' | 'first' | 'last';

// Type-safe aggregatable fields
export const ShippingAddressRequestAggregateFields = {
    PolicyVersion: 'policyVersion',
} as const;

export type ShippingAddressRequestAggregateField = typeof ShippingAddressRequestAggregateFields[keyof typeof ShippingAddressRequestAggregateFields];

// Type-safe group-by fields
export const ShippingAddressRequestGroupByFields = {
    DecisionReason: 'decisionReason',
    EmailEnc: 'emailEnc',
    EmailHash: 'emailHash',
    Name: 'name',
    Note: 'note',
    OwnerId: 'ownerId',
    PolicyVersion: 'policyVersion',
    RegistryId: 'registryId',
    RegistryItemId: 'registryItemId',
} as const;

export type ShippingAddressRequestGroupByField = typeof ShippingAddressRequestGroupByFields[keyof typeof ShippingAddressRequestGroupByFields];

export type AggregateFieldSpec = {
    field: ShippingAddressRequestAggregateField;
    method: AggregateMethod;
    alias?: string;
}

// Aggregate result row with partial model fields and metadata
export type ShippingAddressRequestAggregateResultRow = {
    // Group-by fields (original types)
    decisionReason?: string | null;
    emailEnc?: string | null;
    emailHash?: string | null;
    name?: string | null;
    note?: string | null;
    ownerId?: string | null;
    policyVersion?: number | null;
    registryId?: string | null;
    registryItemId?: string | null;
    // Aggregate fields - always numbers since they're results of sum/avg/etc
    // Ref field owner
    owner?: OwnerUser | null;
    // Ref field registry
    registry?: Registry | null;
    // Ref field registryItem
    registryItem?: RegistryItem | null;
    // Metadata indicating which fields are populated
    __groupKeys: ShippingAddressRequestGroupByField[];
    __aggregateKeys: string[];
}

export type ShippingAddressRequestAggregateResponse = {
    data: ShippingAddressRequestAggregateResultRow[];
    total: number;
}

export type AggregateShippingAddressRequestParams = {
    baseUrl: string;
    query: ShippingAddressRequestSearchQuery;
    fields: AggregateFieldSpec[];
    groupBy: ShippingAddressRequestGroupByField[];
    ownerProjection?: OwnerUserProjection;
    registryProjection?: RegistryProjection;
    registryItemProjection?: RegistryItemProjection;
}

export function aggregateShippingAddressRequests(params: AggregateShippingAddressRequestParams): Promise<ShippingAddressRequestAggregateResponse> {
    return fetch(`${params.baseUrl}/shipping-address-requests/aggregate`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            query: params.query,
            fields: params.fields,
            groupBy: params.groupBy,
            ownerProjection: params.ownerProjection,
            registryProjection: params.registryProjection,
            registryItemProjection: params.registryItemProjection,
        }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to aggregate ShippingAddressRequest');
            return Promise.reject(err);
        }
        return response.json();
    });
}
