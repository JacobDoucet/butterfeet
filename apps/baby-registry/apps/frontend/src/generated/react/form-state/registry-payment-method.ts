// This file is auto-generated. DO NOT EDIT.

import { useCallback } from 'react';
import { RegistryPaymentMethod } from '../../model/registry-payment-method-model';
import { useFormState } from './common';
import {
    RegistryPaymentMethodMutationOptions,
    useCreateRegistryPaymentMethod,
    useUpdateRegistryPaymentMethod,
} from '../tanstack-query/registry-payment-method-queries';
import { MutationResponse } from '../../api/model';

type UseRegistryPaymentMethodFormStateOptions = {
    initialState: RegistryPaymentMethod;
    onSuccess?: (res: MutationResponse<RegistryPaymentMethod>) => void;
    onError?: (error?: any) => void;
    mutationOptions?: RegistryPaymentMethodMutationOptions;
};

export function useRegistryPaymentMethodFormState(options: UseRegistryPaymentMethodFormStateOptions) {
    const { initialState, ...mutationOptions } = options;
    const formState = useFormState(initialState);

    const updateRegistryPaymentMethod = useUpdateRegistryPaymentMethod(options.mutationOptions);
    const createRegistryPaymentMethod = useCreateRegistryPaymentMethod(options.mutationOptions);

    const save = useCallback(() => {
        const opts = {
            onSuccess: options.onSuccess,
            onError: options.onError,
        };
        if (formState.currentState.id) {
            return updateRegistryPaymentMethod.mutate(formState.updates, opts);
        }
        return createRegistryPaymentMethod.mutate(formState.currentState, opts);
    }, [
        formState.currentState,
        formState.updates,
        createRegistryPaymentMethod,
        updateRegistryPaymentMethod,
        options.onSuccess,
        options.onError,
    ]);

    const isLoading = createRegistryPaymentMethod.isLoading || updateRegistryPaymentMethod.isLoading;

    return {
        ...formState,
        save,
        createMutation: createRegistryPaymentMethod,
        updateMutation: updateRegistryPaymentMethod,
        isLoading,
    } as const;
}
