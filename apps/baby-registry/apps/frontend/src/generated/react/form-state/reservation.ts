// This file is auto-generated. DO NOT EDIT.

import { useCallback } from 'react';
import { Reservation } from '../../model/reservation-model';
import { useFormState } from './common';
import {
    ReservationMutationOptions,
    useCreateReservation,
    useUpdateReservation,
} from '../tanstack-query/reservation-queries';
import { MutationResponse } from '../../api/model';

type UseReservationFormStateOptions = {
    initialState: Reservation;
    onSuccess?: (res: MutationResponse<Reservation>) => void;
    onError?: (error?: any) => void;
    mutationOptions?: ReservationMutationOptions;
};

export function useReservationFormState(options: UseReservationFormStateOptions) {
    const { initialState, ...mutationOptions } = options;
    const formState = useFormState(initialState);

    const updateReservation = useUpdateReservation(options.mutationOptions);
    const createReservation = useCreateReservation(options.mutationOptions);

    const save = useCallback(() => {
        const opts = {
            onSuccess: options.onSuccess,
            onError: options.onError,
        };
        if (formState.currentState.id) {
            return updateReservation.mutate(formState.updates, opts);
        }
        return createReservation.mutate(formState.currentState, opts);
    }, [
        formState.currentState,
        formState.updates,
        createReservation,
        updateReservation,
        options.onSuccess,
        options.onError,
    ]);

    const isLoading = createReservation.isLoading || updateReservation.isLoading;

    return {
        ...formState,
        save,
        createMutation: createReservation,
        updateMutation: updateReservation,
        isLoading,
    } as const;
}
