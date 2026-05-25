// This file is auto-generated. DO NOT EDIT.

import { RegistrySearchQuery, RegistryWithRefs, RegistryWithRefsProjection } from '../model/registry-api';
import { SelectManyResponse, MutationResponse, AggregateResponse } from './model';
import { ApiError } from './errors';
import { Registry, RegistrySortParams } from '../model/registry-model';
import { AddressAccessSession, AddressAccessSessionProjection } from '../model/address-access-session-model';
import { OwnerUser, OwnerUserProjection } from '../model/owner-user-model';
import { RegistryApprovedGuest, RegistryApprovedGuestProjection } from '../model/registry-approved-guest-model';
import { RegistryItem, RegistryItemProjection } from '../model/registry-item-model';
import { Reservation, ReservationProjection } from '../model/reservation-model';
import { ShippingAddressRequest, ShippingAddressRequestProjection } from '../model/shipping-address-request-model';

export type SearchRegistrysParams = {
    baseUrl: string;
    query: RegistrySearchQuery;
    sort?: RegistrySortParams;
    projection?: RegistryWithRefsProjection;
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

export function searchRegistrys(params: SearchRegistrysParams): Promise<SelectManyResponse<RegistryWithRefs>> {
    return fetch(`${params.baseUrl}/registries/search`, {
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
            const err = await newApiError(response, 'Failed to search Registry');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type SelectRegistryByIdParams = {
    baseUrl: string;
    id: string;
    projection?: RegistryWithRefsProjection;
}

export function selectRegistryById(params: SelectRegistryByIdParams): Promise<RegistryWithRefs> {
    return fetch(`${params.baseUrl}/registries/id/${params.id}`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              projection: params.projection,
          }),
    }).then(async (response) => {
          if (!response.ok) {
              const err = await newApiError(response, 'Failed to select Registry');
              return Promise.reject(err);
          }
          return response.json();
     });
}

export type SelectRegistryBySlugUniqueParams = {
    baseUrl: string;
    slug: string;
    projection?: RegistryWithRefsProjection;
}

export function selectRegistryBySlugUnique(params: SelectRegistryBySlugUniqueParams): Promise<RegistryWithRefs> {
    return fetch(`${params.baseUrl}/registries/slug/${params.slug}`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              projection: params.projection,
          }),
    }).then(async (response) => {
          if (!response.ok) {
              const err = await newApiError(response, 'Failed to select Registry');
              return Promise.reject(err);
          }
          return response.json();
     });
}

export type SaveRegistryParams = {
    baseUrl: string;
    data: Registry;
}

export function createRegistry(params: SaveRegistryParams): Promise<MutationResponse<Registry>> {
    return fetch(`${params.baseUrl}/registries/create`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to create Registry');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export function updateRegistry(params: SaveRegistryParams): Promise<MutationResponse<Registry>> {
    return fetch(`${params.baseUrl}/registries/update`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to update Registry');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type DeleteRegistryParams = {
    baseUrl: string;
    id: string;
}

export function deleteRegistry({ baseUrl, id }: DeleteRegistryParams): Promise<void> {
    return fetch(`${baseUrl}/registries/delete/${id}`, {
        method: 'DELETE',
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to delete Registry');
            return Promise.reject(err);
        }
        return;
    });
}

// Aggregation types
export type AggregateMethod = 'sum' | 'avg' | 'min' | 'max' | 'count' | 'first' | 'last';

// Type-safe aggregatable fields
export const RegistryAggregateFields = {
    ShippingPolicyVersion: 'shippingPolicyVersion',
} as const;

export type RegistryAggregateField = typeof RegistryAggregateFields[keyof typeof RegistryAggregateFields];

// Type-safe group-by fields
export const RegistryGroupByFields = {
    CoverImageUrl: 'coverImageUrl',
    DueDate: 'dueDate',
    IsPublic: 'isPublic',
    OwnerId: 'ownerId',
    ParentNames: 'parentNames',
    ShippingCity: 'shippingCity',
    ShippingCountry: 'shippingCountry',
    ShippingDeliveryNotes: 'shippingDeliveryNotes',
    ShippingLine1: 'shippingLine1',
    ShippingLine2: 'shippingLine2',
    ShippingPolicyVersion: 'shippingPolicyVersion',
    ShippingPostalCode: 'shippingPostalCode',
    ShippingRecipientName: 'shippingRecipientName',
    ShippingRegion: 'shippingRegion',
    Slug: 'slug',
    ThemeColor: 'themeColor',
    Title: 'title',
    WelcomeMessage: 'welcomeMessage',
} as const;

export type RegistryGroupByField = typeof RegistryGroupByFields[keyof typeof RegistryGroupByFields];

export type AggregateFieldSpec = {
    field: RegistryAggregateField;
    method: AggregateMethod;
    alias?: string;
}

// Aggregate result row with partial model fields and metadata
export type RegistryAggregateResultRow = {
    // Group-by fields (original types)
    coverImageUrl?: string | null;
    dueDate?: string | null;
    isPublic?: boolean | null;
    ownerId?: string | null;
    parentNames?: string | null;
    shippingCity?: string | null;
    shippingCountry?: string | null;
    shippingDeliveryNotes?: string | null;
    shippingLine1?: string | null;
    shippingLine2?: string | null;
    shippingPolicyVersion?: number | null;
    shippingPostalCode?: string | null;
    shippingRecipientName?: string | null;
    shippingRegion?: string | null;
    slug?: string | null;
    themeColor?: string | null;
    title?: string | null;
    welcomeMessage?: string | null;
    // Aggregate fields - always numbers since they're results of sum/avg/etc
    // Ref field owner
    owner?: OwnerUser | null;
    // Ref field addressAccessSessions
    addressAccessSessions?: AddressAccessSession[] | null;
    // Ref field registryApprovedGuests
    registryApprovedGuests?: RegistryApprovedGuest[] | null;
    // Ref field registryItems
    registryItems?: RegistryItem[] | null;
    // Ref field reservations
    reservations?: Reservation[] | null;
    // Ref field shippingAddressRequests
    shippingAddressRequests?: ShippingAddressRequest[] | null;
    // Metadata indicating which fields are populated
    __groupKeys: RegistryGroupByField[];
    __aggregateKeys: string[];
}

export type RegistryAggregateResponse = {
    data: RegistryAggregateResultRow[];
    total: number;
}

export type AggregateRegistryParams = {
    baseUrl: string;
    query: RegistrySearchQuery;
    fields: AggregateFieldSpec[];
    groupBy: RegistryGroupByField[];
    addressAccessSessionsProjection?: AddressAccessSessionProjection;
    registryApprovedGuestsProjection?: RegistryApprovedGuestProjection;
    registryItemsProjection?: RegistryItemProjection;
    reservationsProjection?: ReservationProjection;
    shippingAddressRequestsProjection?: ShippingAddressRequestProjection;
    ownerProjection?: OwnerUserProjection;
}

export function aggregateRegistrys(params: AggregateRegistryParams): Promise<RegistryAggregateResponse> {
    return fetch(`${params.baseUrl}/registries/aggregate`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            query: params.query,
            fields: params.fields,
            groupBy: params.groupBy,
            addressAccessSessionsProjection: params.addressAccessSessionsProjection,
            registryApprovedGuestsProjection: params.registryApprovedGuestsProjection,
            registryItemsProjection: params.registryItemsProjection,
            reservationsProjection: params.reservationsProjection,
            shippingAddressRequestsProjection: params.shippingAddressRequestsProjection,
            ownerProjection: params.ownerProjection,
        }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to aggregate Registry');
            return Promise.reject(err);
        }
        return response.json();
    });
}
