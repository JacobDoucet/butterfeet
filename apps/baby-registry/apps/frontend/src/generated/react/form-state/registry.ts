// This file is auto-generated. DO NOT EDIT.

import { useCallback } from 'react';
import { Registry } from '../../model/registry-model';
import { useFormState } from './common';
import {
    RegistryMutationOptions,
    useCreateRegistry,
    useUpdateRegistry,
} from '../tanstack-query/registry-queries';
import { MutationResponse } from '../../api/model';

type UseRegistryFormStateOptions = {
    initialState: Registry;
    onSuccess?: (res: MutationResponse<Registry>) => void;
    onError?: (error?: any) => void;
    mutationOptions?: RegistryMutationOptions;
};

export function useRegistryFormState(options: UseRegistryFormStateOptions) {
    const { initialState, ...mutationOptions } = options;
    const formState = useFormState(initialState);

    const updateRegistry = useUpdateRegistry(options.mutationOptions);
    const createRegistry = useCreateRegistry(options.mutationOptions);

    const save = useCallback(() => {
        const opts = {
            onSuccess: options.onSuccess,
            onError: options.onError,
        };
        if (formState.currentState.id) {
            return updateRegistry.mutate(formState.updates, opts);
        }
        return createRegistry.mutate(formState.currentState, opts);
    }, [
        formState.currentState,
        formState.updates,
        createRegistry,
        updateRegistry,
        options.onSuccess,
        options.onError,
    ]);

    const isLoading = createRegistry.isLoading || updateRegistry.isLoading;

    return {
        ...formState,
        save,
        createMutation: createRegistry,
        updateMutation: updateRegistry,
        isLoading,
    } as const;
}
