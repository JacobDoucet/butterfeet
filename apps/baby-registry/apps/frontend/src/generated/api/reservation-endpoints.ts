// This file is auto-generated. DO NOT EDIT.

import { ReservationSearchQuery, ReservationWithRefs, ReservationWithRefsProjection } from '../model/reservation-api';
import { SelectManyResponse, MutationResponse, AggregateResponse } from './model';
import { ApiError } from './errors';
import { Reservation, ReservationSortParams } from '../model/reservation-model';
import { Registry, RegistryProjection } from '../model/registry-model';
import { RegistryItem, RegistryItemProjection } from '../model/registry-item-model';

export type SearchReservationsParams = {
    baseUrl: string;
    query: ReservationSearchQuery;
    sort?: ReservationSortParams;
    projection?: ReservationWithRefsProjection;
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

export function searchReservations(params: SearchReservationsParams): Promise<SelectManyResponse<ReservationWithRefs>> {
    return fetch(`${params.baseUrl}/reservations/search`, {
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
            const err = await newApiError(response, 'Failed to search Reservation');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type SelectReservationByIdParams = {
    baseUrl: string;
    id: string;
    projection?: ReservationWithRefsProjection;
}

export function selectReservationById(params: SelectReservationByIdParams): Promise<ReservationWithRefs> {
    return fetch(`${params.baseUrl}/reservations/id/${params.id}`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              projection: params.projection,
          }),
    }).then(async (response) => {
          if (!response.ok) {
              const err = await newApiError(response, 'Failed to select Reservation');
              return Promise.reject(err);
          }
          return response.json();
     });
}

export type SaveReservationParams = {
    baseUrl: string;
    data: Reservation;
}

export function createReservation(params: SaveReservationParams): Promise<MutationResponse<Reservation>> {
    return fetch(`${params.baseUrl}/reservations/create`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to create Reservation');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export function updateReservation(params: SaveReservationParams): Promise<MutationResponse<Reservation>> {
    return fetch(`${params.baseUrl}/reservations/update`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to update Reservation');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type DeleteReservationParams = {
    baseUrl: string;
    id: string;
}

export function deleteReservation({ baseUrl, id }: DeleteReservationParams): Promise<void> {
    return fetch(`${baseUrl}/reservations/delete/${id}`, {
        method: 'DELETE',
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to delete Reservation');
            return Promise.reject(err);
        }
        return;
    });
}

// Aggregation types
export type AggregateMethod = 'sum' | 'avg' | 'min' | 'max' | 'count' | 'first' | 'last';

// Type-safe aggregatable fields
export const ReservationAggregateFields = {
    Quantity: 'quantity',
} as const;

export type ReservationAggregateField = typeof ReservationAggregateFields[keyof typeof ReservationAggregateFields];

// Type-safe group-by fields
export const ReservationGroupByFields = {
    ContactEmail: 'contactEmail',
    IsAnonymous: 'isAnonymous',
    ItemId: 'itemId',
    Message: 'message',
    Quantity: 'quantity',
    RegistryId: 'registryId',
    ReserverName: 'reserverName',
} as const;

export type ReservationGroupByField = typeof ReservationGroupByFields[keyof typeof ReservationGroupByFields];

export type AggregateFieldSpec = {
    field: ReservationAggregateField;
    method: AggregateMethod;
    alias?: string;
}

// Aggregate result row with partial model fields and metadata
export type ReservationAggregateResultRow = {
    // Group-by fields (original types)
    contactEmail?: string | null;
    isAnonymous?: boolean | null;
    itemId?: string | null;
    message?: string | null;
    quantity?: number | null;
    registryId?: string | null;
    reserverName?: string | null;
    // Aggregate fields - always numbers since they're results of sum/avg/etc
    // Ref field item
    item?: RegistryItem | null;
    // Ref field registry
    registry?: Registry | null;
    // Metadata indicating which fields are populated
    __groupKeys: ReservationGroupByField[];
    __aggregateKeys: string[];
}

export type ReservationAggregateResponse = {
    data: ReservationAggregateResultRow[];
    total: number;
}

export type AggregateReservationParams = {
    baseUrl: string;
    query: ReservationSearchQuery;
    fields: AggregateFieldSpec[];
    groupBy: ReservationGroupByField[];
    itemProjection?: RegistryItemProjection;
    registryProjection?: RegistryProjection;
}

export function aggregateReservations(params: AggregateReservationParams): Promise<ReservationAggregateResponse> {
    return fetch(`${params.baseUrl}/reservations/aggregate`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            query: params.query,
            fields: params.fields,
            groupBy: params.groupBy,
            itemProjection: params.itemProjection,
            registryProjection: params.registryProjection,
        }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to aggregate Reservation');
            return Promise.reject(err);
        }
        return response.json();
    });
}
