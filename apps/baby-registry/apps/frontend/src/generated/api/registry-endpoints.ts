// This file is auto-generated. DO NOT EDIT.

import { RegistrySearchQuery, RegistryWithRefs, RegistryWithRefsProjection } from '../model/registry-api';
import { SelectManyResponse, MutationResponse, AggregateResponse } from './model';
import { ApiError } from './errors';
import { Registry, RegistrySortParams } from '../model/registry-model';

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
