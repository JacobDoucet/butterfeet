// This file is auto-generated. DO NOT EDIT.

import { OwnerUserSearchQuery, OwnerUserWithRefs, OwnerUserWithRefsProjection } from '../model/owner-user-api';
import { SelectManyResponse, MutationResponse, AggregateResponse } from './model';
import { ApiError } from './errors';
import { OwnerUser, OwnerUserSortParams } from '../model/owner-user-model';

export type SearchOwnerUsersParams = {
    baseUrl: string;
    query: OwnerUserSearchQuery;
    sort?: OwnerUserSortParams;
    projection?: OwnerUserWithRefsProjection;
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

export function searchOwnerUsers(params: SearchOwnerUsersParams): Promise<SelectManyResponse<OwnerUserWithRefs>> {
    return fetch(`${params.baseUrl}/owner-users/search`, {
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
            const err = await newApiError(response, 'Failed to search OwnerUser');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type SelectOwnerUserByIdParams = {
    baseUrl: string;
    id: string;
    projection?: OwnerUserWithRefsProjection;
}

export function selectOwnerUserById(params: SelectOwnerUserByIdParams): Promise<OwnerUserWithRefs> {
    return fetch(`${params.baseUrl}/owner-users/id/${params.id}`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              projection: params.projection,
          }),
    }).then(async (response) => {
          if (!response.ok) {
              const err = await newApiError(response, 'Failed to select OwnerUser');
              return Promise.reject(err);
          }
          return response.json();
     });
}

export type SelectOwnerUserByEmailUniqueParams = {
    baseUrl: string;
    email: string;
    projection?: OwnerUserWithRefsProjection;
}

export function selectOwnerUserByEmailUnique(params: SelectOwnerUserByEmailUniqueParams): Promise<OwnerUserWithRefs> {
    return fetch(`${params.baseUrl}/owner-users/email/${params.email}`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              projection: params.projection,
          }),
    }).then(async (response) => {
          if (!response.ok) {
              const err = await newApiError(response, 'Failed to select OwnerUser');
              return Promise.reject(err);
          }
          return response.json();
     });
}

export type SaveOwnerUserParams = {
    baseUrl: string;
    data: OwnerUser;
}

export function createOwnerUser(params: SaveOwnerUserParams): Promise<MutationResponse<OwnerUser>> {
    return fetch(`${params.baseUrl}/owner-users/create`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to create OwnerUser');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export function updateOwnerUser(params: SaveOwnerUserParams): Promise<MutationResponse<OwnerUser>> {
    return fetch(`${params.baseUrl}/owner-users/update`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: params.data }),
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to update OwnerUser');
            return Promise.reject(err);
        }
        return response.json();
    });
}

export type DeleteOwnerUserParams = {
    baseUrl: string;
    id: string;
}

export function deleteOwnerUser({ baseUrl, id }: DeleteOwnerUserParams): Promise<void> {
    return fetch(`${baseUrl}/owner-users/delete/${id}`, {
        method: 'DELETE',
    }).then(async (response) => {
        if (!response.ok) {
            const err = await newApiError(response, 'Failed to delete OwnerUser');
            return Promise.reject(err);
        }
        return;
    });
}
