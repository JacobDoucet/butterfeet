// This file is auto-generated. DO NOT EDIT.

import { useMemo } from 'react';
import { useQuery, useMutation, UseQueryOptions } from '@tanstack/react-query';
import { useApiBaseUrl } from '../api';
import { SelectManyResponse, MutationResponse } from '../../api/model';
import { Cart } from '../../model/cart-model';
import { CartWithRefs } from '../../model/cart-api';
import {
    searchCarts, SearchCartsParams,
    selectCartById, SelectCartByIdParams,
    selectCartByReferenceUnique, SelectCartByReferenceUniqueParams,
    createCart, updateCart, deleteCart,
    aggregateCarts, AggregateCartParams, CartAggregateResponse,
} from '../../api/cart-endpoints';
import { ApiError } from '../../api/errors';

type UseSearchCartsProps = Omit<SearchCartsParams, 'baseUrl'> & {
    queryName?: string,
    queryKey?: any[];
}

type SearchQueryOptions = Omit<UseQueryOptions<
    SelectManyResponse<CartWithRefs>,
    ApiError,
    SelectManyResponse<CartWithRefs>,
    any[]
>, 'initialData'>;

export function useSearchCarts(
    { queryKey, queryName, ...params }: UseSearchCartsProps,
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
            `${key}=${JSON.stringify(params.query[key as keyof SearchCartsParams['query']])}`
        );
        return ['searchCarts', queryName, ...searchKey];
    }, [queryName, queryKey, params.query]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => searchCarts({ baseUrl, ...params }),
    });
}
type UseSelectCartByIdProps = Omit<SelectCartByIdParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type SelectCartByIdOptions = Omit<UseQueryOptions<
    CartWithRefs,
    ApiError,
    CartWithRefs,
    any[]
>, 'initialData'>;

export function useSelectCartById(
    { queryKey, queryName, ...params }: UseSelectCartByIdProps,
    queryOptions?: SelectCartByIdOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        return queryKey || ['selectCartById', queryName, params.id];
    }, [queryKey, queryName, params.id]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => selectCartById({ baseUrl, ...params }),
    });
}
type UseSelectCartByReferenceUniqueProps = Omit<SelectCartByReferenceUniqueParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type SelectCartByReferenceUniqueOptions = Omit<UseQueryOptions<
    CartWithRefs,
    ApiError,
    CartWithRefs,
    any[]
>, 'initialData'>;

export function useSelectCartByReferenceUnique(
    { queryKey, queryName, ...params }: UseSelectCartByReferenceUniqueProps,
    queryOptions?: SelectCartByReferenceUniqueOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        return queryKey || ['selectCartByReferenceUnique', queryName, params.referenceCode];
    }, [queryKey, queryName, params.referenceCode]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => selectCartByReferenceUnique({ baseUrl, ...params }),
    });
}

export type CartMutationOptions = {
    onAfterCommit?: () => Promise<void>;
}

export function useCreateCart(options: CartMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<Cart>, ApiError, Cart>(async (cart: Cart) => {
        const res = await createCart({ baseUrl, data: cart });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useUpdateCart(options: CartMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<Cart>, ApiError, Cart>(async (cart: Cart) => {
        const res = await updateCart({ baseUrl, data: cart });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useDeleteCart(options: CartMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<void, ApiError, string>(async (id: string) => {
        const res = await deleteCart({ baseUrl, id });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

type UseAggregateCartsProps = Omit<AggregateCartParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type CartAggregateQueryOptions = Omit<UseQueryOptions<
    CartAggregateResponse,
    ApiError,
    CartAggregateResponse,
    any[]
>, 'initialData'>;

export function useAggregateCarts(
    { queryKey, queryName, ...params }: UseAggregateCartsProps,
    queryOptions?: CartAggregateQueryOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        if (queryKey) {
            return queryKey;
        }
        const queryKeys = Object.keys(params.query);
        queryKeys.sort();
        const searchKey = queryKeys.map((key) =>
            `${key}=${JSON.stringify(params.query[key as keyof AggregateCartParams['query']])}`
        );
        const fieldKeys = params.fields.map((f) => `${f.field}_${f.method}`);
        const groupByKeys = params.groupBy.join(',');
        return ['aggregateCarts', queryName, ...searchKey, ...fieldKeys, groupByKeys];
    }, [queryName, queryKey, params.query, params.fields, params.groupBy]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => aggregateCarts({ baseUrl, ...params }),
    });
}
