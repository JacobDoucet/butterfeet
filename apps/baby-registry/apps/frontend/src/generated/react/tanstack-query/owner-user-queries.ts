// This file is auto-generated. DO NOT EDIT.

import { useMemo } from 'react';
import { useQuery, useMutation, UseQueryOptions } from '@tanstack/react-query';
import { useApiBaseUrl } from '../api';
import { SelectManyResponse, MutationResponse } from '../../api/model';
import { OwnerUser } from '../../model/owner-user-model';
import { OwnerUserWithRefs } from '../../model/owner-user-api';
import {
    searchOwnerUsers, SearchOwnerUsersParams,
    selectOwnerUserById, SelectOwnerUserByIdParams,
    selectOwnerUserByEmailUnique, SelectOwnerUserByEmailUniqueParams,
    createOwnerUser, updateOwnerUser, deleteOwnerUser,
} from '../../api/owner-user-endpoints';
import { ApiError } from '../../api/errors';

type UseSearchOwnerUsersProps = Omit<SearchOwnerUsersParams, 'baseUrl'> & {
    queryName?: string,
    queryKey?: any[];
}

type SearchQueryOptions = Omit<UseQueryOptions<
    SelectManyResponse<OwnerUserWithRefs>,
    ApiError,
    SelectManyResponse<OwnerUserWithRefs>,
    any[]
>, 'initialData'>;

export function useSearchOwnerUsers(
    { queryKey, queryName, ...params }: UseSearchOwnerUsersProps,
    queryOptions?: SearchQueryOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        if (queryKey) {
            return queryKey;
        }
        const keys = Object.keys(params.query);
        keys.sort();
        const searchKey = keys.map((key) =>
            `${key}=${JSON.stringify(params.query[key as keyof SearchOwnerUsersParams['query']])}`
        );
        return ['searchOwnerUsers', queryName, ...searchKey];
    }, [queryName, queryKey, params.query]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => searchOwnerUsers({ baseUrl, ...params }),
    });
}
type UseSelectOwnerUserByIdProps = Omit<SelectOwnerUserByIdParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type SelectOwnerUserByIdOptions = Omit<UseQueryOptions<
    OwnerUserWithRefs,
    ApiError,
    OwnerUserWithRefs,
    any[]
>, 'initialData'>;

export function useSelectOwnerUserById(
    { queryKey, queryName, ...params }: UseSelectOwnerUserByIdProps,
    queryOptions?: SelectOwnerUserByIdOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        return queryKey || ['selectOwnerUserById', queryName, params.id];
    }, [queryKey, queryName, params.id]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => selectOwnerUserById({ baseUrl, ...params }),
    });
}
type UseSelectOwnerUserByEmailUniqueProps = Omit<SelectOwnerUserByEmailUniqueParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type SelectOwnerUserByEmailUniqueOptions = Omit<UseQueryOptions<
    OwnerUserWithRefs,
    ApiError,
    OwnerUserWithRefs,
    any[]
>, 'initialData'>;

export function useSelectOwnerUserByEmailUnique(
    { queryKey, queryName, ...params }: UseSelectOwnerUserByEmailUniqueProps,
    queryOptions?: SelectOwnerUserByEmailUniqueOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        return queryKey || ['selectOwnerUserByEmailUnique', queryName, params.email];
    }, [queryKey, queryName, params.email]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => selectOwnerUserByEmailUnique({ baseUrl, ...params }),
    });
}

export type OwnerUserMutationOptions = {
    onAfterCommit?: () => Promise<void>;
}

export function useCreateOwnerUser(options: OwnerUserMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<OwnerUser>, ApiError, OwnerUser>(async (ownerUser: OwnerUser) => {
        const res = await createOwnerUser({ baseUrl, data: ownerUser });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useUpdateOwnerUser(options: OwnerUserMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<OwnerUser>, ApiError, OwnerUser>(async (ownerUser: OwnerUser) => {
        const res = await updateOwnerUser({ baseUrl, data: ownerUser });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useDeleteOwnerUser(options: OwnerUserMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<void, ApiError, string>(async (id: string) => {
        const res = await deleteOwnerUser({ baseUrl, id });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}
