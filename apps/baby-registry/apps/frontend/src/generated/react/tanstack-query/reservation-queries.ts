// This file is auto-generated. DO NOT EDIT.

import { useMemo } from 'react';
import { useQuery, useMutation, UseQueryOptions } from '@tanstack/react-query';
import { useApiBaseUrl } from '../api';
import { SelectManyResponse, MutationResponse } from '../../api/model';
import { Reservation } from '../../model/reservation-model';
import { ReservationWithRefs } from '../../model/reservation-api';
import {
    searchReservations, SearchReservationsParams,
    selectReservationById, SelectReservationByIdParams,
    createReservation, updateReservation, deleteReservation,
    aggregateReservations, AggregateReservationParams, ReservationAggregateResponse,
} from '../../api/reservation-endpoints';
import { ApiError } from '../../api/errors';

type UseSearchReservationsProps = Omit<SearchReservationsParams, 'baseUrl'> & {
    queryName?: string,
    queryKey?: any[];
}

type SearchQueryOptions = Omit<UseQueryOptions<
    SelectManyResponse<ReservationWithRefs>,
    ApiError,
    SelectManyResponse<ReservationWithRefs>,
    any[]
>, 'initialData'>;

export function useSearchReservations(
    { queryKey, queryName, ...params }: UseSearchReservationsProps,
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
            `${key}=${JSON.stringify(params.query[key as keyof SearchReservationsParams['query']])}`
        );
        return ['searchReservations', queryName, ...searchKey];
    }, [queryName, queryKey, params.query]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => searchReservations({ baseUrl, ...params }),
    });
}
type UseSelectReservationByIdProps = Omit<SelectReservationByIdParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type SelectReservationByIdOptions = Omit<UseQueryOptions<
    ReservationWithRefs,
    ApiError,
    ReservationWithRefs,
    any[]
>, 'initialData'>;

export function useSelectReservationById(
    { queryKey, queryName, ...params }: UseSelectReservationByIdProps,
    queryOptions?: SelectReservationByIdOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        return queryKey || ['selectReservationById', queryName, params.id];
    }, [queryKey, queryName, params.id]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => selectReservationById({ baseUrl, ...params }),
    });
}

export type ReservationMutationOptions = {
    onAfterCommit?: () => Promise<void>;
}

export function useCreateReservation(options: ReservationMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<Reservation>, ApiError, Reservation>(async (reservation: Reservation) => {
        const res = await createReservation({ baseUrl, data: reservation });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useUpdateReservation(options: ReservationMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<MutationResponse<Reservation>, ApiError, Reservation>(async (reservation: Reservation) => {
        const res = await updateReservation({ baseUrl, data: reservation });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

export function useDeleteReservation(options: ReservationMutationOptions = {}) {
    const baseUrl = useApiBaseUrl();
    return useMutation<void, ApiError, string>(async (id: string) => {
        const res = await deleteReservation({ baseUrl, id });
        if (options.onAfterCommit) {
            await options.onAfterCommit();
        }
        return res;
    });
}

type UseAggregateReservationsProps = Omit<AggregateReservationParams, 'baseUrl'> & {
    queryName?: string;
    queryKey?: any[];
}

type ReservationAggregateQueryOptions = Omit<UseQueryOptions<
    ReservationAggregateResponse,
    ApiError,
    ReservationAggregateResponse,
    any[]
>, 'initialData'>;

export function useAggregateReservations(
    { queryKey, queryName, ...params }: UseAggregateReservationsProps,
    queryOptions?: ReservationAggregateQueryOptions,
) {
    const baseUrl = useApiBaseUrl();

    const memoizedQueryKey = useMemo(() => {
        if (queryKey) {
            return queryKey;
        }
        const queryKeys = Object.keys(params.query);
        queryKeys.sort();
        const searchKey = queryKeys.map((key) =>
            `${key}=${JSON.stringify(params.query[key as keyof AggregateReservationParams['query']])}`
        );
        const fieldKeys = params.fields.map((f) => `${f.field}_${f.method}`);
        const groupByKeys = params.groupBy.join(',');
        return ['aggregateReservations', queryName, ...searchKey, ...fieldKeys, groupByKeys];
    }, [queryName, queryKey, params.query, params.fields, params.groupBy]);

    return useQuery({
        ...queryOptions,
        queryKey: memoizedQueryKey,
        queryFn: () => aggregateReservations({ baseUrl, ...params }),
    });
}
