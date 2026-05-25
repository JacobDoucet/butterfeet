// This file is auto-generated. DO NOT EDIT.

import { RegistryItemSearchQuery, RegistryItemWithRefs, RegistryItemWithRefsProjection } from '../model/registry-item-api';
import { SelectManyResponse, MutationResponse, AggregateResponse } from './model';
import { ApiError } from './errors';
import { RegistryItem, RegistryItemSortParams } from '../model/registry-item-model';
import { Registry, RegistryProjection } from '../model/registry-model';
import { Reservation, ReservationProjection } from '../model/reservation-model';
import { ShippingAddressRequest, ShippingAddressRequestProjection } from '../model/shipping-address-request-model';

export type SearchRegistryItemsParams = {
    baseUrl: string;
    query: RegistryItemSearchQuery;
    sort?: RegistryItemSortParams;
    projection?: RegistryItemWithRefsProjection;
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

export function searchRegistryItems(params: SearchRegistryItemsParams): Promise<SelectManyResponse<RegistryItemWithRefs>> {
    return fetch(`${params.baseUrl}/registry-items/search`, {
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
            const err = await newApiError(response, 'Failed to search RegistryItem');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type SelectRegistryItemByIdParams = {
    baseUrl: string;
    id: string;
    projection?: RegistryItemWithRefsProjection;
}

export function selectRegistryItemById(params: SelectRegistryItemByIdParams): Promise<RegistryItemWithRefs> {
    return fetch(`${params.baseUrl}/registry-items/id/${params.id}`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              projection: params.projection,
          }),
    }).then(async (response) => {
          if (!response.ok) {
              const err = await newApiError(response, 'Failed to select RegistryItem');
              return Promise.reject(err);
          }
          return response.json();
     });
}

export type SaveRegistryItemParams = {
    baseUrl: string;
    data: RegistryItem;
}

export function createRegistryItem(params: SaveRegistryItemParams): Promise<MutationResponse<RegistryItem>> {
    return fetch(`${params.baseUrl}/registry-items/create`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to create RegistryItem');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export function updateRegistryItem(params: SaveRegistryItemParams): Promise<MutationResponse<RegistryItem>> {
    return fetch(`${params.baseUrl}/registry-items/update`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to update RegistryItem');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type DeleteRegistryItemParams = {
    baseUrl: string;
    id: string;
}

export function deleteRegistryItem({ baseUrl, id }: DeleteRegistryItemParams): Promise<void> {
    return fetch(`${baseUrl}/registry-items/delete/${id}`, {
        method: 'DELETE',
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to delete RegistryItem');
            return Promise.reject(err);
        }
        return;
    });
}

// Aggregation types
export type AggregateMethod = 'sum' | 'avg' | 'min' | 'max' | 'count' | 'first' | 'last';

// Type-safe aggregatable fields
export const RegistryItemAggregateFields = {
    Position: 'position',
    PriceCents: 'priceCents',
    Quantity: 'quantity',
} as const;

export type RegistryItemAggregateField = typeof RegistryItemAggregateFields[keyof typeof RegistryItemAggregateFields];

// Type-safe group-by fields
export const RegistryItemGroupByFields = {
    Currency: 'currency',
    Description: 'description',
    ImageUrl: 'imageUrl',
    Notes: 'notes',
    OwnerPurchased: 'ownerPurchased',
    Position: 'position',
    PriceCents: 'priceCents',
    ProductUrl: 'productUrl',
    Quantity: 'quantity',
    RegistryId: 'registryId',
    Title: 'title',
} as const;

export type RegistryItemGroupByField = typeof RegistryItemGroupByFields[keyof typeof RegistryItemGroupByFields];

export type AggregateFieldSpec = {
    field: RegistryItemAggregateField;
    method: AggregateMethod;
    alias?: string;
}

// Aggregate result row with partial model fields and metadata
export type RegistryItemAggregateResultRow = {
    // Group-by fields (original types)
    currency?: string | null;
    description?: string | null;
    imageUrl?: string | null;
    notes?: string | null;
    ownerPurchased?: boolean | null;
    position?: number | null;
    priceCents?: number | null;
    productUrl?: string | null;
    quantity?: number | null;
    registryId?: string | null;
    title?: string | null;
    // Aggregate fields - always numbers since they're results of sum/avg/etc
    // Ref field registry
    registry?: Registry | null;
    // Ref field reservations
    reservations?: Reservation[] | null;
    // Ref field shippingAddressRequests
    shippingAddressRequests?: ShippingAddressRequest[] | null;
    // Metadata indicating which fields are populated
    __groupKeys: RegistryItemGroupByField[];
    __aggregateKeys: string[];
}

export type RegistryItemAggregateResponse = {
    data: RegistryItemAggregateResultRow[];
    total: number;
}

export type AggregateRegistryItemParams = {
    baseUrl: string;
    query: RegistryItemSearchQuery;
    fields: AggregateFieldSpec[];
    groupBy: RegistryItemGroupByField[];
    reservationsProjection?: ReservationProjection;
    shippingAddressRequestsProjection?: ShippingAddressRequestProjection;
    registryProjection?: RegistryProjection;
}

export function aggregateRegistryItems(params: AggregateRegistryItemParams): Promise<RegistryItemAggregateResponse> {
    return fetch(`${params.baseUrl}/registry-items/aggregate`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            query: params.query,
            fields: params.fields,
            groupBy: params.groupBy,
            reservationsProjection: params.reservationsProjection,
            shippingAddressRequestsProjection: params.shippingAddressRequestsProjection,
            registryProjection: params.registryProjection,
        }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to aggregate RegistryItem');
            return Promise.reject(err);
        }
        return response.json();
    });
}
