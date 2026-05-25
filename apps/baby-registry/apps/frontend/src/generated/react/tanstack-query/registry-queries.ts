// This file is auto-generated. DO NOT EDIT.

import { useMemo } from 'react';
import { useQuery, useMutation, UseQueryOptions } from '@tanstack/react-query';
import { useApiBaseUrl } from '../api';
import { SelectManyResponse, MutationResponse } from '../../api/model';
import { Registry } from '../../model/registry-model';
import { RegistryWithRefs } from '../../model/registry-api';
import {
    searchRegistrys, SearchRegistrysParams,
    selectRegistryById, SelectRegistryByIdParams,
    selectRegistryBySlugUnique, SelectRegistryBySlugUniqueParams,
    createRegistry, updateRegistry, deleteRegistry,
    aggregateRegistrys, AggregateRegistryParams, RegistryAggregateResponse,
} from '../../api/registry-endpoints';
import { ApiError } from '../../api/errors';

type UseSearchRegistrysProps = Omit<SearchRegistrysParams, 'baseUrl'> & {
    queryName?: string,
    queryKey?: any[];
}

type SearchQueryOptions = Omit<UseQueryOptions<
    SelectManyResponse<RegistryWithRefs>,
    ApiError,
    SelectManyResponse<RegistryWithRefs>,
    any[]
>, 'initialData'>;

export function useSearchRegistrys(
    { queryKey, queryName, ...params }: UseSearchRegistrysProps,
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
            `${key}=${JSON.stringify(params.query[key as keyof SearchRegistrysParams['query']])}`
        );
        return ['searchRegistrys', queryName, ...searchKey];
    }, [queryName, queryKey, params.query]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => searchRegistrys({ baseUrl, ...params }),
    });
}
type UseSelectRegistryByIdProps = Omit<SelectRegistryByIdParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type SelectRegistryByIdOptions = Omit<UseQueryOptions<
    RegistryWithRefs,
    ApiError,
    RegistryWithRefs,
    any[]
>, 'initialData'>;

export function useSelectRegistryById(
    { queryKey, queryName, ...params }: UseSelectRegistryByIdProps,
    queryOptions?: SelectRegistryByIdOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        return queryKey || ['selectRegistryById', queryName, params.id];
    }, [queryKey, queryName, params.id]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => selectRegistryById({ baseUrl, ...params }),
    });
}
type UseSelectRegistryBySlugUniqueProps = Omit<SelectRegistryBySlugUniqueParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type SelectRegistryBySlugUniqueOptions = Omit<UseQueryOptions<
    RegistryWithRefs,
    ApiError,
    RegistryWithRefs,
    any[]
>, 'initialData'>;

export function useSelectRegistryBySlugUnique(
    { queryKey, queryName, ...params }: UseSelectRegistryBySlugUniqueProps,
    queryOptions?: SelectRegistryBySlugUniqueOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        return queryKey || ['selectRegistryBySlugUnique', queryName, params.slug];
    }, [queryKey, queryName, params.slug]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => selectRegistryBySlugUnique({ baseUrl, ...params }),
    });
}

export type RegistryMutationOptions = {
    onAfterCommit?: () => Promise<void>;
}

export function useCreateRegistry(options: RegistryMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<Registry>, ApiError, Registry>(async (registry: Registry) => {
        const res = await createRegistry({ baseUrl, data: registry });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useUpdateRegistry(options: RegistryMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<Registry>, ApiError, Registry>(async (registry: Registry) => {
        const res = await updateRegistry({ baseUrl, data: registry });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useDeleteRegistry(options: RegistryMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<void, ApiError, string>(async (id: string) => {
        const res = await deleteRegistry({ baseUrl, id });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

type UseAggregateRegistrysProps = Omit<AggregateRegistryParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type RegistryAggregateQueryOptions = Omit<UseQueryOptions<
    RegistryAggregateResponse,
    ApiError,
    RegistryAggregateResponse,
    any[]
>, 'initialData'>;

export function useAggregateRegistrys(
    { queryKey, queryName, ...params }: UseAggregateRegistrysProps,
    queryOptions?: RegistryAggregateQueryOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        if (queryKey) {
            return queryKey;
        }
        const queryKeys = Object.keys(params.query);
        queryKeys.sort();
        const searchKey = queryKeys.map((key) =>
            `${key}=${JSON.stringify(params.query[key as keyof AggregateRegistryParams['query']])}`
        );
        const fieldKeys = params.fields.map((f) => `${f.field}_${f.method}`);
        const groupByKeys = params.groupBy.join(',');
        return ['aggregateRegistrys', queryName, ...searchKey, ...fieldKeys, groupByKeys];
    }, [queryName, queryKey, params.query, params.fields, params.groupBy]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => aggregateRegistrys({ baseUrl, ...params }),
    });
}
