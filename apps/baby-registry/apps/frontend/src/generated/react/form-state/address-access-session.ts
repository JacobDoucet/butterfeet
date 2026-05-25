// This file is auto-generated. DO NOT EDIT.

import { useCallback } from 'react';
import { AddressAccessSession } from '../../model/address-access-session-model';
import { useFormState } from './common';
import {
    AddressAccessSessionMutationOptions,
    useCreateAddressAccessSession,
    useUpdateAddressAccessSession,
} from '../tanstack-query/address-access-session-queries';
import { MutationResponse } from '../../api/model';

type UseAddressAccessSessionFormStateOptions = {
    initialState: AddressAccessSession;
    onSuccess?: (res: MutationResponse<AddressAccessSession>) => void;
    onError?: (error?: any) => void;
    mutationOptions?: AddressAccessSessionMutationOptions;
};

export function useAddressAccessSessionFormState(options: UseAddressAccessSessionFormStateOptions) {
    const { initialState, ...mutationOptions } = options;
    const formState = useFormState(initialState);

    const updateAddressAccessSession = useUpdateAddressAccessSession(options.mutationOptions);
    const createAddressAccessSession = useCreateAddressAccessSession(options.mutationOptions);

    const save = useCallback(() => {
        const opts = {
            onSuccess: options.onSuccess,
            onError: options.onError,
        };
        if (formState.currentState.id) {
            return updateAddressAccessSession.mutate(formState.updates, opts);
        }
        return createAddressAccessSession.mutate(formState.currentState, opts);
    }, [
        formState.currentState,
        formState.updates,
        createAddressAccessSession,
        updateAddressAccessSession,
        options.onSuccess,
        options.onError,
    ]);

    const isLoading = createAddressAccessSession.isLoading || updateAddressAccessSession.isLoading;

    return {
        ...formState,
        save,
        createMutation: createAddressAccessSession,
        updateMutation: updateAddressAccessSession,
        isLoading,
    } as const;
}
