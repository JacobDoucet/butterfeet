// This file is auto-generated. DO NOT EDIT.

import { RegistryApprovedGuestSearchQuery, RegistryApprovedGuestWithRefs, RegistryApprovedGuestWithRefsProjection } from '../model/registry-approved-guest-api';
import { SelectManyResponse, MutationResponse, AggregateResponse } from './model';
import { ApiError } from './errors';
import { RegistryApprovedGuest, RegistryApprovedGuestSortParams } from '../model/registry-approved-guest-model';

export type SearchRegistryApprovedGuestsParams = {
    baseUrl: string;
    query: RegistryApprovedGuestSearchQuery;
    sort?: RegistryApprovedGuestSortParams;
    projection?: RegistryApprovedGuestWithRefsProjection;
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

export function searchRegistryApprovedGuests(params: SearchRegistryApprovedGuestsParams): Promise<SelectManyResponse<RegistryApprovedGuestWithRefs>> {
    return fetch(`${params.baseUrl}/registry-approved-guests/search`, {
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
            const err = await newApiError(response, 'Failed to search RegistryApprovedGuest');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type SelectRegistryApprovedGuestByIdParams = {
    baseUrl: string;
    id: string;
    projection?: RegistryApprovedGuestWithRefsProjection;
}

export function selectRegistryApprovedGuestById(params: SelectRegistryApprovedGuestByIdParams): Promise<RegistryApprovedGuestWithRefs> {
    return fetch(`${params.baseUrl}/registry-approved-guests/id/${params.id}`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              projection: params.projection,
          }),
    }).then(async (response) => {
          if (!response.ok) {
              const err = await newApiError(response, 'Failed to select RegistryApprovedGuest');
              return Promise.reject(err);
          }
          return response.json();
     });
}

export type SelectRegistryApprovedGuestByRegistryEmailUniqueParams = {
    baseUrl: string;
    registryId: string;
	emailHash: string;
    projection?: RegistryApprovedGuestWithRefsProjection;
}

export function selectRegistryApprovedGuestByRegistryEmailUnique(params: SelectRegistryApprovedGuestByRegistryEmailUniqueParams): Promise<RegistryApprovedGuestWithRefs> {
    return fetch(`${params.baseUrl}/registry-approved-guests/registryId/${params.registryId}/emailHash/${params.emailHash}`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              projection: params.projection,
          }),
    }).then(async (response) => {
          if (!response.ok) {
              const err = await newApiError(response, 'Failed to select RegistryApprovedGuest');
              return Promise.reject(err);
          }
          return response.json();
     });
}

export type SaveRegistryApprovedGuestParams = {
    baseUrl: string;
    data: RegistryApprovedGuest;
}

export function createRegistryApprovedGuest(params: SaveRegistryApprovedGuestParams): Promise<MutationResponse<RegistryApprovedGuest>> {
    return fetch(`${params.baseUrl}/registry-approved-guests/create`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to create RegistryApprovedGuest');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export function updateRegistryApprovedGuest(params: SaveRegistryApprovedGuestParams): Promise<MutationResponse<RegistryApprovedGuest>> {
    return fetch(`${params.baseUrl}/registry-approved-guests/update`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to update RegistryApprovedGuest');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type DeleteRegistryApprovedGuestParams = {
    baseUrl: string;
    id: string;
}

export function deleteRegistryApprovedGuest({ baseUrl, id }: DeleteRegistryApprovedGuestParams): Promise<void> {
    return fetch(`${baseUrl}/registry-approved-guests/delete/${id}`, {
        method: 'DELETE',
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to delete RegistryApprovedGuest');
            return Promise.reject(err);
        }
        return;
    });
}
