// This file is auto-generated. DO NOT EDIT.

import { useMemo } from 'react';
import { useQuery, useMutation, UseQueryOptions } from '@tanstack/react-query';
import { useApiBaseUrl } from '../api';
import { SelectManyResponse, MutationResponse } from '../../api/model';
import { RegistryItem } from '../../model/registry-item-model';
import { RegistryItemWithRefs } from '../../model/registry-item-api';
import {
    searchRegistryItems, SearchRegistryItemsParams,
    selectRegistryItemById, SelectRegistryItemByIdParams,
    createRegistryItem, updateRegistryItem, deleteRegistryItem,
    aggregateRegistryItems, AggregateRegistryItemParams, RegistryItemAggregateResponse,
} from '../../api/registry-item-endpoints';
import { ApiError } from '../../api/errors';

type UseSearchRegistryItemsProps = Omit<SearchRegistryItemsParams, 'baseUrl'> & {
    queryName?: string,
    queryKey?: any[];
}

type SearchQueryOptions = Omit<UseQueryOptions<
    SelectManyResponse<RegistryItemWithRefs>,
    ApiError,
    SelectManyResponse<RegistryItemWithRefs>,
    any[]
>, 'initialData'>;

export function useSearchRegistryItems(
    { queryKey, queryName, ...params }: UseSearchRegistryItemsProps,
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
            `${key}=${JSON.stringify(params.query[key as keyof SearchRegistryItemsParams['query']])}`
        );
        return ['searchRegistryItems', queryName, ...searchKey];
    }, [queryName, queryKey, params.query]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => searchRegistryItems({ baseUrl, ...params }),
    });
}
type UseSelectRegistryItemByIdProps = Omit<SelectRegistryItemByIdParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type SelectRegistryItemByIdOptions = Omit<UseQueryOptions<
    RegistryItemWithRefs,
    ApiError,
    RegistryItemWithRefs,
    any[]
>, 'initialData'>;

export function useSelectRegistryItemById(
    { queryKey, queryName, ...params }: UseSelectRegistryItemByIdProps,
    queryOptions?: SelectRegistryItemByIdOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        return queryKey || ['selectRegistryItemById', queryName, params.id];
    }, [queryKey, queryName, params.id]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => selectRegistryItemById({ baseUrl, ...params }),
    });
}

export type RegistryItemMutationOptions = {
    onAfterCommit?: () => Promise<void>;
}

export function useCreateRegistryItem(options: RegistryItemMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<RegistryItem>, ApiError, RegistryItem>(async (registryItem: RegistryItem) => {
        const res = await createRegistryItem({ baseUrl, data: registryItem });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useUpdateRegistryItem(options: RegistryItemMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<RegistryItem>, ApiError, RegistryItem>(async (registryItem: RegistryItem) => {
        const res = await updateRegistryItem({ baseUrl, data: registryItem });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useDeleteRegistryItem(options: RegistryItemMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<void, ApiError, string>(async (id: string) => {
        const res = await deleteRegistryItem({ baseUrl, id });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

type UseAggregateRegistryItemsProps = Omit<AggregateRegistryItemParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type RegistryItemAggregateQueryOptions = Omit<UseQueryOptions<
    RegistryItemAggregateResponse,
    ApiError,
    RegistryItemAggregateResponse,
    any[]
>, 'initialData'>;

export function useAggregateRegistryItems(
    { queryKey, queryName, ...params }: UseAggregateRegistryItemsProps,
    queryOptions?: RegistryItemAggregateQueryOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        if (queryKey) {
            return queryKey;
        }
        const queryKeys = Object.keys(params.query);
        queryKeys.sort();
        const searchKey = queryKeys.map((key) =>
            `${key}=${JSON.stringify(params.query[key as keyof AggregateRegistryItemParams['query']])}`
        );
        const fieldKeys = params.fields.map((f) => `${f.field}_${f.method}`);
        const groupByKeys = params.groupBy.join(',');
        return ['aggregateRegistryItems', queryName, ...searchKey, ...fieldKeys, groupByKeys];
    }, [queryName, queryKey, params.query, params.fields, params.groupBy]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => aggregateRegistryItems({ baseUrl, ...params }),
    });
}
