// This file is auto-generated. DO NOT EDIT.

import { useMemo } from 'react';
import { useQuery, useMutation, UseQueryOptions } from '@tanstack/react-query';
import { useApiBaseUrl } from '../api';
import { SelectManyResponse, MutationResponse } from '../../api/model';
import { ShippingAddressRequest } from '../../model/shipping-address-request-model';
import { ShippingAddressRequestWithRefs } from '../../model/shipping-address-request-api';
import {
    searchShippingAddressRequests, SearchShippingAddressRequestsParams,
    selectShippingAddressRequestById, SelectShippingAddressRequestByIdParams,
    createShippingAddressRequest, updateShippingAddressRequest, deleteShippingAddressRequest,
    aggregateShippingAddressRequests, AggregateShippingAddressRequestParams, ShippingAddressRequestAggregateResponse,
} from '../../api/shipping-address-request-endpoints';
import { ApiError } from '../../api/errors';

type UseSearchShippingAddressRequestsProps = Omit<SearchShippingAddressRequestsParams, 'baseUrl'> & {
    queryName?: string,
    queryKey?: any[];
}

type SearchQueryOptions = Omit<UseQueryOptions<
    SelectManyResponse<ShippingAddressRequestWithRefs>,
    ApiError,
    SelectManyResponse<ShippingAddressRequestWithRefs>,
    any[]
>, 'initialData'>;

export function useSearchShippingAddressRequests(
    { queryKey, queryName, ...params }: UseSearchShippingAddressRequestsProps,
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
            `${key}=${JSON.stringify(params.query[key as keyof SearchShippingAddressRequestsParams['query']])}`
        );
        return ['searchShippingAddressRequests', queryName, ...searchKey];
    }, [queryName, queryKey, params.query]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => searchShippingAddressRequests({ baseUrl, ...params }),
    });
}
type UseSelectShippingAddressRequestByIdProps = Omit<SelectShippingAddressRequestByIdParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type SelectShippingAddressRequestByIdOptions = Omit<UseQueryOptions<
    ShippingAddressRequestWithRefs,
    ApiError,
    ShippingAddressRequestWithRefs,
    any[]
>, 'initialData'>;

export function useSelectShippingAddressRequestById(
    { queryKey, queryName, ...params }: UseSelectShippingAddressRequestByIdProps,
    queryOptions?: SelectShippingAddressRequestByIdOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        return queryKey || ['selectShippingAddressRequestById', queryName, params.id];
    }, [queryKey, queryName, params.id]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => selectShippingAddressRequestById({ baseUrl, ...params }),
    });
}

export type ShippingAddressRequestMutationOptions = {
    onAfterCommit?: () => Promise<void>;
}

export function useCreateShippingAddressRequest(options: ShippingAddressRequestMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<ShippingAddressRequest>, ApiError, ShippingAddressRequest>(async (shippingAddressRequest: ShippingAddressRequest) => {
        const res = await createShippingAddressRequest({ baseUrl, data: shippingAddressRequest });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useUpdateShippingAddressRequest(options: ShippingAddressRequestMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<ShippingAddressRequest>, ApiError, ShippingAddressRequest>(async (shippingAddressRequest: ShippingAddressRequest) => {
        const res = await updateShippingAddressRequest({ baseUrl, data: shippingAddressRequest });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useDeleteShippingAddressRequest(options: ShippingAddressRequestMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<void, ApiError, string>(async (id: string) => {
        const res = await deleteShippingAddressRequest({ baseUrl, id });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

type UseAggregateShippingAddressRequestsProps = Omit<AggregateShippingAddressRequestParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type ShippingAddressRequestAggregateQueryOptions = Omit<UseQueryOptions<
    ShippingAddressRequestAggregateResponse,
    ApiError,
    ShippingAddressRequestAggregateResponse,
    any[]
>, 'initialData'>;

export function useAggregateShippingAddressRequests(
    { queryKey, queryName, ...params }: UseAggregateShippingAddressRequestsProps,
    queryOptions?: ShippingAddressRequestAggregateQueryOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        if (queryKey) {
            return queryKey;
        }
        const queryKeys = Object.keys(params.query);
        queryKeys.sort();
        const searchKey = queryKeys.map((key) =>
            `${key}=${JSON.stringify(params.query[key as keyof AggregateShippingAddressRequestParams['query']])}`
        );
        const fieldKeys = params.fields.map((f) => `${f.field}_${f.method}`);
        const groupByKeys = params.groupBy.join(',');
        return ['aggregateShippingAddressRequests', queryName, ...searchKey, ...fieldKeys, groupByKeys];
    }, [queryName, queryKey, params.query, params.fields, params.groupBy]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => aggregateShippingAddressRequests({ baseUrl, ...params }),
    });
}
