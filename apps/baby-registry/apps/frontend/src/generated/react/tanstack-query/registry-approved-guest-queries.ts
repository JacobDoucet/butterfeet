// This file is auto-generated. DO NOT EDIT.

import { useMemo } from 'react';
import { useQuery, useMutation, UseQueryOptions } from '@tanstack/react-query';
import { useApiBaseUrl } from '../api';
import { SelectManyResponse, MutationResponse } from '../../api/model';
import { RegistryApprovedGuest } from '../../model/registry-approved-guest-model';
import { RegistryApprovedGuestWithRefs } from '../../model/registry-approved-guest-api';
import {
    searchRegistryApprovedGuests, SearchRegistryApprovedGuestsParams,
    selectRegistryApprovedGuestById, SelectRegistryApprovedGuestByIdParams,
    selectRegistryApprovedGuestByRegistryEmailUnique, SelectRegistryApprovedGuestByRegistryEmailUniqueParams,
    createRegistryApprovedGuest, updateRegistryApprovedGuest, deleteRegistryApprovedGuest,
} from '../../api/registry-approved-guest-endpoints';
import { ApiError } from '../../api/errors';

type UseSearchRegistryApprovedGuestsProps = Omit<SearchRegistryApprovedGuestsParams, 'baseUrl'> & {
    queryName?: string,
    queryKey?: any[];
}

type SearchQueryOptions = Omit<UseQueryOptions<
    SelectManyResponse<RegistryApprovedGuestWithRefs>,
    ApiError,
    SelectManyResponse<RegistryApprovedGuestWithRefs>,
    any[]
>, 'initialData'>;

export function useSearchRegistryApprovedGuests(
    { queryKey, queryName, ...params }: UseSearchRegistryApprovedGuestsProps,
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
            `${key}=${JSON.stringify(params.query[key as keyof SearchRegistryApprovedGuestsParams['query']])}`
        );
        return ['searchRegistryApprovedGuests', queryName, ...searchKey];
    }, [queryName, queryKey, params.query]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => searchRegistryApprovedGuests({ baseUrl, ...params }),
    });
}
type UseSelectRegistryApprovedGuestByIdProps = Omit<SelectRegistryApprovedGuestByIdParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type SelectRegistryApprovedGuestByIdOptions = Omit<UseQueryOptions<
    RegistryApprovedGuestWithRefs,
    ApiError,
    RegistryApprovedGuestWithRefs,
    any[]
>, 'initialData'>;

export function useSelectRegistryApprovedGuestById(
    { queryKey, queryName, ...params }: UseSelectRegistryApprovedGuestByIdProps,
    queryOptions?: SelectRegistryApprovedGuestByIdOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        return queryKey || ['selectRegistryApprovedGuestById', queryName, params.id];
    }, [queryKey, queryName, params.id]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => selectRegistryApprovedGuestById({ baseUrl, ...params }),
    });
}
type UseSelectRegistryApprovedGuestByRegistryEmailUniqueProps = Omit<SelectRegistryApprovedGuestByRegistryEmailUniqueParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type SelectRegistryApprovedGuestByRegistryEmailUniqueOptions = Omit<UseQueryOptions<
    RegistryApprovedGuestWithRefs,
    ApiError,
    RegistryApprovedGuestWithRefs,
    any[]
>, 'initialData'>;

export function useSelectRegistryApprovedGuestByRegistryEmailUnique(
    { queryKey, queryName, ...params }: UseSelectRegistryApprovedGuestByRegistryEmailUniqueProps,
    queryOptions?: SelectRegistryApprovedGuestByRegistryEmailUniqueOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        return queryKey || ['selectRegistryApprovedGuestByRegistryEmailUnique', queryName, params.registryId, params.emailHash];
    }, [queryKey, queryName, params.registryId, params.emailHash]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => selectRegistryApprovedGuestByRegistryEmailUnique({ baseUrl, ...params }),
    });
}

export type RegistryApprovedGuestMutationOptions = {
    onAfterCommit?: () => Promise<void>;
}

export function useCreateRegistryApprovedGuest(options: RegistryApprovedGuestMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<RegistryApprovedGuest>, ApiError, RegistryApprovedGuest>(async (registryApprovedGuest: RegistryApprovedGuest) => {
        const res = await createRegistryApprovedGuest({ baseUrl, data: registryApprovedGuest });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useUpdateRegistryApprovedGuest(options: RegistryApprovedGuestMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<RegistryApprovedGuest>, ApiError, RegistryApprovedGuest>(async (registryApprovedGuest: RegistryApprovedGuest) => {
        const res = await updateRegistryApprovedGuest({ baseUrl, data: registryApprovedGuest });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useDeleteRegistryApprovedGuest(options: RegistryApprovedGuestMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<void, ApiError, string>(async (id: string) => {
        const res = await deleteRegistryApprovedGuest({ baseUrl, id });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}
