// This file is auto-generated. DO NOT EDIT.

import { CartSearchQuery, CartWithRefs, CartWithRefsProjection } from '../model/cart-api';
import { SelectManyResponse, MutationResponse, AggregateResponse } from './model';
import { ApiError } from './errors';
import { Cart, CartSortParams } from '../model/cart-model';
import { OwnerUser, OwnerUserProjection } from '../model/owner-user-model';
import { Registry, RegistryProjection } from '../model/registry-model';
import { RegistryPaymentMethod, RegistryPaymentMethodProjection } from '../model/registry-payment-method-model';
import { Reservation, ReservationProjection } from '../model/reservation-model';

export type SearchCartsParams = {
    baseUrl: string;
    query: CartSearchQuery;
    sort?: CartSortParams;
    projection?: CartWithRefsProjection;
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

export function searchCarts(params: SearchCartsParams): Promise<SelectManyResponse<CartWithRefs>> {
    return fetch(`${params.baseUrl}/carts/search`, {
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
            const err = await newApiError(response, 'Failed to search Cart');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type SelectCartByIdParams = {
    baseUrl: string;
    id: string;
    projection?: CartWithRefsProjection;
}

export function selectCartById(params: SelectCartByIdParams): Promise<CartWithRefs> {
    return fetch(`${params.baseUrl}/carts/id/${params.id}`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              projection: params.projection,
          }),
    }).then(async (response) => {
          if (!response.ok) {
              const err = await newApiError(response, 'Failed to select Cart');
              return Promise.reject(err);
          }
          return response.json();
     });
}

export type SelectCartByReferenceUniqueParams = {
    baseUrl: string;
    referenceCode: string;
    projection?: CartWithRefsProjection;
}

export function selectCartByReferenceUnique(params: SelectCartByReferenceUniqueParams): Promise<CartWithRefs> {
    return fetch(`${params.baseUrl}/carts/referenceCode/${params.referenceCode}`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              projection: params.projection,
          }),
    }).then(async (response) => {
          if (!response.ok) {
              const err = await newApiError(response, 'Failed to select Cart');
              return Promise.reject(err);
          }
          return response.json();
     });
}

export type SaveCartParams = {
    baseUrl: string;
    data: Cart;
}

export function createCart(params: SaveCartParams): Promise<MutationResponse<Cart>> {
    return fetch(`${params.baseUrl}/carts/create`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to create Cart');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export function updateCart(params: SaveCartParams): Promise<MutationResponse<Cart>> {
    return fetch(`${params.baseUrl}/carts/update`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to update Cart');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type DeleteCartParams = {
    baseUrl: string;
    id: string;
}

export function deleteCart({ baseUrl, id }: DeleteCartParams): Promise<void> {
    return fetch(`${baseUrl}/carts/delete/${id}`, {
        method: 'DELETE',
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to delete Cart');
            return Promise.reject(err);
        }
        return;
    });
}

// Aggregation types
export type AggregateMethod = 'sum' | 'avg' | 'min' | 'max' | 'count' | 'first' | 'last';

// Type-safe aggregatable fields
export const CartAggregateFields = {
    AmountCents: 'amountCents',
} as const;

export type CartAggregateField = typeof CartAggregateFields[keyof typeof CartAggregateFields];

// Type-safe group-by fields
export const CartGroupByFields = {
    AmountCents: 'amountCents',
    ClaimedAt: 'claimedAt',
    ContributorEmail: 'contributorEmail',
    ContributorName: 'contributorName',
    Currency: 'currency',
    DecidedAt: 'decidedAt',
    DecisionReason: 'decisionReason',
    Message: 'message',
    MethodDisplayName: 'methodDisplayName',
    OwnerId: 'ownerId',
    PaymentMethodId: 'paymentMethodId',
    ReferenceCode: 'referenceCode',
    RegistryId: 'registryId',
} as const;

export type CartGroupByField = typeof CartGroupByFields[keyof typeof CartGroupByFields];

export type AggregateFieldSpec = {
    field: CartAggregateField;
    method: AggregateMethod;
    alias?: string;
}

// Aggregate result row with partial model fields and metadata
export type CartAggregateResultRow = {
    // Group-by fields (original types)
    amountCents?: number | null;
    claimedAt?: string | null;
    contributorEmail?: string | null;
    contributorName?: string | null;
    currency?: string | null;
    decidedAt?: string | null;
    decisionReason?: string | null;
    message?: string | null;
    methodDisplayName?: string | null;
    ownerId?: string | null;
    paymentMethodId?: string | null;
    referenceCode?: string | null;
    registryId?: string | null;
    // Aggregate fields - always numbers since they're results of sum/avg/etc
    // Ref field owner
    owner?: OwnerUser | null;
    // Ref field paymentMethod
    paymentMethod?: RegistryPaymentMethod | null;
    // Ref field registry
    registry?: Registry | null;
    // Ref field reservations
    reservations?: Reservation[] | null;
    // Metadata indicating which fields are populated
    __groupKeys: CartGroupByField[];
    __aggregateKeys: string[];
}

export type CartAggregateResponse = {
    data: CartAggregateResultRow[];
    total: number;
}

export type AggregateCartParams = {
    baseUrl: string;
    query: CartSearchQuery;
    fields: AggregateFieldSpec[];
    groupBy: CartGroupByField[];
    reservationsProjection?: ReservationProjection;
    ownerProjection?: OwnerUserProjection;
    paymentMethodProjection?: RegistryPaymentMethodProjection;
    registryProjection?: RegistryProjection;
}

export function aggregateCarts(params: AggregateCartParams): Promise<CartAggregateResponse> {
    return fetch(`${params.baseUrl}/carts/aggregate`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            query: params.query,
            fields: params.fields,
            groupBy: params.groupBy,
            reservationsProjection: params.reservationsProjection,
            ownerProjection: params.ownerProjection,
            paymentMethodProjection: params.paymentMethodProjection,
            registryProjection: params.registryProjection,
        }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to aggregate Cart');
            return Promise.reject(err);
        }
        return response.json();
    });
}
