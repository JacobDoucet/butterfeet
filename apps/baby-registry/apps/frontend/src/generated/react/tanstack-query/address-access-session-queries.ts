// This file is auto-generated. DO NOT EDIT.

import { useMemo } from 'react';
import { useQuery, useMutation, UseQueryOptions } from '@tanstack/react-query';
import { useApiBaseUrl } from '../api';
import { SelectManyResponse, MutationResponse } from '../../api/model';
import { AddressAccessSession } from '../../model/address-access-session-model';
import { AddressAccessSessionWithRefs } from '../../model/address-access-session-api';
import {
    searchAddressAccessSessions, SearchAddressAccessSessionsParams,
    selectAddressAccessSessionById, SelectAddressAccessSessionByIdParams,
    selectAddressAccessSessionByTokenUnique, SelectAddressAccessSessionByTokenUniqueParams,
    createAddressAccessSession, updateAddressAccessSession, deleteAddressAccessSession,
    aggregateAddressAccessSessions, AggregateAddressAccessSessionParams, AddressAccessSessionAggregateResponse,
} from '../../api/address-access-session-endpoints';
import { ApiError } from '../../api/errors';

type UseSearchAddressAccessSessionsProps = Omit<SearchAddressAccessSessionsParams, 'baseUrl'> & {
    queryName?: string,
    queryKey?: any[];
}

type SearchQueryOptions = Omit<UseQueryOptions<
    SelectManyResponse<AddressAccessSessionWithRefs>,
    ApiError,
    SelectManyResponse<AddressAccessSessionWithRefs>,
    any[]
>, 'initialData'>;

export function useSearchAddressAccessSessions(
    { queryKey, queryName, ...params }: UseSearchAddressAccessSessionsProps,
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
            `${key}=${JSON.stringify(params.query[key as keyof SearchAddressAccessSessionsParams['query']])}`
        );
        return ['searchAddressAccessSessions', queryName, ...searchKey];
    }, [queryName, queryKey, params.query]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => searchAddressAccessSessions({ baseUrl, ...params }),
    });
}
type UseSelectAddressAccessSessionByIdProps = Omit<SelectAddressAccessSessionByIdParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type SelectAddressAccessSessionByIdOptions = Omit<UseQueryOptions<
    AddressAccessSessionWithRefs,
    ApiError,
    AddressAccessSessionWithRefs,
    any[]
>, 'initialData'>;

export function useSelectAddressAccessSessionById(
    { queryKey, queryName, ...params }: UseSelectAddressAccessSessionByIdProps,
    queryOptions?: SelectAddressAccessSessionByIdOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        return queryKey || ['selectAddressAccessSessionById', queryName, params.id];
    }, [queryKey, queryName, params.id]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => selectAddressAccessSessionById({ baseUrl, ...params }),
    });
}
type UseSelectAddressAccessSessionByTokenUniqueProps = Omit<SelectAddressAccessSessionByTokenUniqueParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type SelectAddressAccessSessionByTokenUniqueOptions = Omit<UseQueryOptions<
    AddressAccessSessionWithRefs,
    ApiError,
    AddressAccessSessionWithRefs,
    any[]
>, 'initialData'>;

export function useSelectAddressAccessSessionByTokenUnique(
    { queryKey, queryName, ...params }: UseSelectAddressAccessSessionByTokenUniqueProps,
    queryOptions?: SelectAddressAccessSessionByTokenUniqueOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        return queryKey || ['selectAddressAccessSessionByTokenUnique', queryName, params.tokenHash];
    }, [queryKey, queryName, params.tokenHash]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => selectAddressAccessSessionByTokenUnique({ baseUrl, ...params }),
    });
}

export type AddressAccessSessionMutationOptions = {
    onAfterCommit?: () => Promise<void>;
}

export function useCreateAddressAccessSession(options: AddressAccessSessionMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<AddressAccessSession>, ApiError, AddressAccessSession>(async (addressAccessSession: AddressAccessSession) => {
        const res = await createAddressAccessSession({ baseUrl, data: addressAccessSession });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useUpdateAddressAccessSession(options: AddressAccessSessionMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<AddressAccessSession>, ApiError, AddressAccessSession>(async (addressAccessSession: AddressAccessSession) => {
        const res = await updateAddressAccessSession({ baseUrl, data: addressAccessSession });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useDeleteAddressAccessSession(options: AddressAccessSessionMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<void, ApiError, string>(async (id: string) => {
        const res = await deleteAddressAccessSession({ baseUrl, id });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

type UseAggregateAddressAccessSessionsProps = Omit<AggregateAddressAccessSessionParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type AddressAccessSessionAggregateQueryOptions = Omit<UseQueryOptions<
    AddressAccessSessionAggregateResponse,
    ApiError,
    AddressAccessSessionAggregateResponse,
    any[]
>, 'initialData'>;

export function useAggregateAddressAccessSessions(
    { queryKey, queryName, ...params }: UseAggregateAddressAccessSessionsProps,
    queryOptions?: AddressAccessSessionAggregateQueryOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        if (queryKey) {
            return queryKey;
        }
        const queryKeys = Object.keys(params.query);
        queryKeys.sort();
        const searchKey = queryKeys.map((key) =>
            `${key}=${JSON.stringify(params.query[key as keyof AggregateAddressAccessSessionParams['query']])}`
        );
        const fieldKeys = params.fields.map((f) => `${f.field}_${f.method}`);
        const groupByKeys = params.groupBy.join(',');
        return ['aggregateAddressAccessSessions', queryName, ...searchKey, ...fieldKeys, groupByKeys];
    }, [queryName, queryKey, params.query, params.fields, params.groupBy]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => aggregateAddressAccessSessions({ baseUrl, ...params }),
    });
}
