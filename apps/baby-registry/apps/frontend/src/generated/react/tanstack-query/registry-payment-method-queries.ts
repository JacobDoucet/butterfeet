// This file is auto-generated. DO NOT EDIT.

import { useMemo } from 'react';
import { useQuery, useMutation, UseQueryOptions } from '@tanstack/react-query';
import { useApiBaseUrl } from '../api';
import { SelectManyResponse, MutationResponse } from '../../api/model';
import { RegistryPaymentMethod } from '../../model/registry-payment-method-model';
import { RegistryPaymentMethodWithRefs } from '../../model/registry-payment-method-api';
import {
    searchRegistryPaymentMethods, SearchRegistryPaymentMethodsParams,
    selectRegistryPaymentMethodById, SelectRegistryPaymentMethodByIdParams,
    createRegistryPaymentMethod, updateRegistryPaymentMethod, deleteRegistryPaymentMethod,
    aggregateRegistryPaymentMethods, AggregateRegistryPaymentMethodParams, RegistryPaymentMethodAggregateResponse,
} from '../../api/registry-payment-method-endpoints';
import { ApiError } from '../../api/errors';

type UseSearchRegistryPaymentMethodsProps = Omit<SearchRegistryPaymentMethodsParams, 'baseUrl'> & {
    queryName?: string,
    queryKey?: any[];
}

type SearchQueryOptions = Omit<UseQueryOptions<
    SelectManyResponse<RegistryPaymentMethodWithRefs>,
    ApiError,
    SelectManyResponse<RegistryPaymentMethodWithRefs>,
    any[]
>, 'initialData'>;

export function useSearchRegistryPaymentMethods(
    { queryKey, queryName, ...params }: UseSearchRegistryPaymentMethodsProps,
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
            `${key}=${JSON.stringify(params.query[key as keyof SearchRegistryPaymentMethodsParams['query']])}`
        );
        return ['searchRegistryPaymentMethods', queryName, ...searchKey];
    }, [queryName, queryKey, params.query]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => searchRegistryPaymentMethods({ baseUrl, ...params }),
    });
}
type UseSelectRegistryPaymentMethodByIdProps = Omit<SelectRegistryPaymentMethodByIdParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type SelectRegistryPaymentMethodByIdOptions = Omit<UseQueryOptions<
    RegistryPaymentMethodWithRefs,
    ApiError,
    RegistryPaymentMethodWithRefs,
    any[]
>, 'initialData'>;

export function useSelectRegistryPaymentMethodById(
    { queryKey, queryName, ...params }: UseSelectRegistryPaymentMethodByIdProps,
    queryOptions?: SelectRegistryPaymentMethodByIdOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        return queryKey || ['selectRegistryPaymentMethodById', queryName, params.id];
    }, [queryKey, queryName, params.id]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => selectRegistryPaymentMethodById({ baseUrl, ...params }),
    });
}

export type RegistryPaymentMethodMutationOptions = {
    onAfterCommit?: () => Promise<void>;
}

export function useCreateRegistryPaymentMethod(options: RegistryPaymentMethodMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<RegistryPaymentMethod>, ApiError, RegistryPaymentMethod>(async (registryPaymentMethod: RegistryPaymentMethod) => {
        const res = await createRegistryPaymentMethod({ baseUrl, data: registryPaymentMethod });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useUpdateRegistryPaymentMethod(options: RegistryPaymentMethodMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<RegistryPaymentMethod>, ApiError, RegistryPaymentMethod>(async (registryPaymentMethod: RegistryPaymentMethod) => {
        const res = await updateRegistryPaymentMethod({ baseUrl, data: registryPaymentMethod });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useDeleteRegistryPaymentMethod(options: RegistryPaymentMethodMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<void, ApiError, string>(async (id: string) => {
        const res = await deleteRegistryPaymentMethod({ baseUrl, id });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

type UseAggregateRegistryPaymentMethodsProps = Omit<AggregateRegistryPaymentMethodParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type RegistryPaymentMethodAggregateQueryOptions = Omit<UseQueryOptions<
    RegistryPaymentMethodAggregateResponse,
    ApiError,
    RegistryPaymentMethodAggregateResponse,
    any[]
>, 'initialData'>;

export function useAggregateRegistryPaymentMethods(
    { queryKey, queryName, ...params }: UseAggregateRegistryPaymentMethodsProps,
    queryOptions?: RegistryPaymentMethodAggregateQueryOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        if (queryKey) {
            return queryKey;
        }
        const queryKeys = Object.keys(params.query);
        queryKeys.sort();
        const searchKey = queryKeys.map((key) =>
            `${key}=${JSON.stringify(params.query[key as keyof AggregateRegistryPaymentMethodParams['query']])}`
        );
        const fieldKeys = params.fields.map((f) => `${f.field}_${f.method}`);
        const groupByKeys = params.groupBy.join(',');
        return ['aggregateRegistryPaymentMethods', queryName, ...searchKey, ...fieldKeys, groupByKeys];
    }, [queryName, queryKey, params.query, params.fields, params.groupBy]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => aggregateRegistryPaymentMethods({ baseUrl, ...params }),
    });
}
