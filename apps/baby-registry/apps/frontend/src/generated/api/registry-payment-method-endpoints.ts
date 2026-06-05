// This file is auto-generated. DO NOT EDIT.

import { RegistryPaymentMethodSearchQuery, RegistryPaymentMethodWithRefs, RegistryPaymentMethodWithRefsProjection } from '../model/registry-payment-method-api';
import { SelectManyResponse, MutationResponse, AggregateResponse } from './model';
import { ApiError } from './errors';
import { RegistryPaymentMethod, RegistryPaymentMethodSortParams } from '../model/registry-payment-method-model';
import { Cart, CartProjection } from '../model/cart-model';
import { OwnerUser, OwnerUserProjection } from '../model/owner-user-model';
import { Registry, RegistryProjection } from '../model/registry-model';

export type SearchRegistryPaymentMethodsParams = {
    baseUrl: string;
    query: RegistryPaymentMethodSearchQuery;
    sort?: RegistryPaymentMethodSortParams;
    projection?: RegistryPaymentMethodWithRefsProjection;
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

export function searchRegistryPaymentMethods(params: SearchRegistryPaymentMethodsParams): Promise<SelectManyResponse<RegistryPaymentMethodWithRefs>> {
    return fetch(`${params.baseUrl}/registry-payment-methods/search`, {
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
            const err = await newApiError(response, 'Failed to search RegistryPaymentMethod');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type SelectRegistryPaymentMethodByIdParams = {
    baseUrl: string;
    id: string;
    projection?: RegistryPaymentMethodWithRefsProjection;
}

export function selectRegistryPaymentMethodById(params: SelectRegistryPaymentMethodByIdParams): Promise<RegistryPaymentMethodWithRefs> {
    return fetch(`${params.baseUrl}/registry-payment-methods/id/${params.id}`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              projection: params.projection,
          }),
    }).then(async (response) => {
          if (!response.ok) {
              const err = await newApiError(response, 'Failed to select RegistryPaymentMethod');
              return Promise.reject(err);
          }
          return response.json();
     });
}

export type SaveRegistryPaymentMethodParams = {
    baseUrl: string;
    data: RegistryPaymentMethod;
}

export function createRegistryPaymentMethod(params: SaveRegistryPaymentMethodParams): Promise<MutationResponse<RegistryPaymentMethod>> {
    return fetch(`${params.baseUrl}/registry-payment-methods/create`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to create RegistryPaymentMethod');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export function updateRegistryPaymentMethod(params: SaveRegistryPaymentMethodParams): Promise<MutationResponse<RegistryPaymentMethod>> {
    return fetch(`${params.baseUrl}/registry-payment-methods/update`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to update RegistryPaymentMethod');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type DeleteRegistryPaymentMethodParams = {
    baseUrl: string;
    id: string;
}

export function deleteRegistryPaymentMethod({ baseUrl, id }: DeleteRegistryPaymentMethodParams): Promise<void> {
    return fetch(`${baseUrl}/registry-payment-methods/delete/${id}`, {
        method: 'DELETE',
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to delete RegistryPaymentMethod');
            return Promise.reject(err);
        }
        return;
    });
}

// Aggregation types
export type AggregateMethod = 'sum' | 'avg' | 'min' | 'max' | 'count' | 'first' | 'last';

// Type-safe aggregatable fields
export const RegistryPaymentMethodAggregateFields = {
    Position: 'position',
} as const;

export type RegistryPaymentMethodAggregateField = typeof RegistryPaymentMethodAggregateFields[keyof typeof RegistryPaymentMethodAggregateFields];

// Type-safe group-by fields
export const RegistryPaymentMethodGroupByFields = {
    BankAccountName: 'bankAccountName',
    BankAccountNumber: 'bankAccountNumber',
    BankIban: 'bankIban',
    BankName: 'bankName',
    BankRoutingNumber: 'bankRoutingNumber',
    BankSwift: 'bankSwift',
    DisplayName: 'displayName',
    Enabled: 'enabled',
    Instructions: 'instructions',
    OwnerId: 'ownerId',
    PaymentUrl: 'paymentUrl',
    Position: 'position',
    RecipientEmail: 'recipientEmail',
    RecipientPhone: 'recipientPhone',
    RegistryId: 'registryId',
} as const;

export type RegistryPaymentMethodGroupByField = typeof RegistryPaymentMethodGroupByFields[keyof typeof RegistryPaymentMethodGroupByFields];

export type AggregateFieldSpec = {
    field: RegistryPaymentMethodAggregateField;
    method: AggregateMethod;
    alias?: string;
}

// Aggregate result row with partial model fields and metadata
export type RegistryPaymentMethodAggregateResultRow = {
    // Group-by fields (original types)
    bankAccountName?: string | null;
    bankAccountNumber?: string | null;
    bankIban?: string | null;
    bankName?: string | null;
    bankRoutingNumber?: string | null;
    bankSwift?: string | null;
    displayName?: string | null;
    enabled?: boolean | null;
    instructions?: string | null;
    ownerId?: string | null;
    paymentUrl?: string | null;
    position?: number | null;
    recipientEmail?: string | null;
    recipientPhone?: string | null;
    registryId?: string | null;
    // Aggregate fields - always numbers since they're results of sum/avg/etc
    // Ref field owner
    owner?: OwnerUser | null;
    // Ref field registry
    registry?: Registry | null;
    // Ref field carts
    carts?: Cart[] | null;
    // Metadata indicating which fields are populated
    __groupKeys: RegistryPaymentMethodGroupByField[];
    __aggregateKeys: string[];
}

export type RegistryPaymentMethodAggregateResponse = {
    data: RegistryPaymentMethodAggregateResultRow[];
    total: number;
}

export type AggregateRegistryPaymentMethodParams = {
    baseUrl: string;
    query: RegistryPaymentMethodSearchQuery;
    fields: AggregateFieldSpec[];
    groupBy: RegistryPaymentMethodGroupByField[];
    cartsProjection?: CartProjection;
    ownerProjection?: OwnerUserProjection;
    registryProjection?: RegistryProjection;
}

export function aggregateRegistryPaymentMethods(params: AggregateRegistryPaymentMethodParams): Promise<RegistryPaymentMethodAggregateResponse> {
    return fetch(`${params.baseUrl}/registry-payment-methods/aggregate`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            query: params.query,
            fields: params.fields,
            groupBy: params.groupBy,
            cartsProjection: params.cartsProjection,
            ownerProjection: params.ownerProjection,
            registryProjection: params.registryProjection,
        }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to aggregate RegistryPaymentMethod');
            return Promise.reject(err);
        }
        return response.json();
    });
}
