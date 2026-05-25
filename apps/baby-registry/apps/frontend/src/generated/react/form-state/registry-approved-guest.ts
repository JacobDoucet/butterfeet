// This file is auto-generated. DO NOT EDIT.

import { useCallback } from 'react';
import { RegistryApprovedGuest } from '../../model/registry-approved-guest-model';
import { useFormState } from './common';
import {
    RegistryApprovedGuestMutationOptions,
    useCreateRegistryApprovedGuest,
    useUpdateRegistryApprovedGuest,
} from '../tanstack-query/registry-approved-guest-queries';
import { MutationResponse } from '../../api/model';

type UseRegistryApprovedGuestFormStateOptions = {
    initialState: RegistryApprovedGuest;
    onSuccess?: (res: MutationResponse<RegistryApprovedGuest>) => void;
    onError?: (error?: any) => void;
    mutationOptions?: RegistryApprovedGuestMutationOptions;
};

export function useRegistryApprovedGuestFormState(options: UseRegistryApprovedGuestFormStateOptions) {
    const { initialState, ...mutationOptions } = options;
    const formState = useFormState(initialState);

    const updateRegistryApprovedGuest = useUpdateRegistryApprovedGuest(options.mutationOptions);
    const createRegistryApprovedGuest = useCreateRegistryApprovedGuest(options.mutationOptions);

    const save = useCallback(() => {
        const opts = {
            onSuccess: options.onSuccess,
            onError: options.onError,
        };
        if (formState.currentState.id) {
            return updateRegistryApprovedGuest.mutate(formState.updates, opts);
        }
        return createRegistryApprovedGuest.mutate(formState.currentState, opts);
    }, [
        formState.currentState,
        formState.updates,
        createRegistryApprovedGuest,
        updateRegistryApprovedGuest,
        options.onSuccess,
        options.onError,
    ]);

    const isLoading = createRegistryApprovedGuest.isLoading || updateRegistryApprovedGuest.isLoading;

    return {
        ...formState,
        save,
        createMutation: createRegistryApprovedGuest,
        updateMutation: updateRegistryApprovedGuest,
        isLoading,
    } as const;
}
