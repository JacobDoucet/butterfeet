// This file is auto-generated. DO NOT EDIT.

import { useCallback } from 'react';
import { RegistryItem } from '../../model/registry-item-model';
import { useFormState } from './common';
import {
    RegistryItemMutationOptions,
    useCreateRegistryItem,
    useUpdateRegistryItem,
} from '../tanstack-query/registry-item-queries';
import { MutationResponse } from '../../api/model';

type UseRegistryItemFormStateOptions = {
    initialState: RegistryItem;
    onSuccess?: (res: MutationResponse<RegistryItem>) => void;
    onError?: (error?: any) => void;
    mutationOptions?: RegistryItemMutationOptions;
};

export function useRegistryItemFormState(options: UseRegistryItemFormStateOptions) {
    const { initialState, ...mutationOptions } = options;
    const formState = useFormState(initialState);

    const updateRegistryItem = useUpdateRegistryItem(options.mutationOptions);
    const createRegistryItem = useCreateRegistryItem(options.mutationOptions);

    const save = useCallback(() => {
        const opts = {
            onSuccess: options.onSuccess,
            onError: options.onError,
        };
        if (formState.currentState.id) {
            return updateRegistryItem.mutate(formState.updates, opts);
        }
        return createRegistryItem.mutate(formState.currentState, opts);
    }, [
        formState.currentState,
        formState.updates,
        createRegistryItem,
        updateRegistryItem,
        options.onSuccess,
        options.onError,
    ]);

    const isLoading = createRegistryItem.isLoading || updateRegistryItem.isLoading;

    return {
        ...formState,
        save,
        createMutation: createRegistryItem,
        updateMutation: updateRegistryItem,
        isLoading,
    } as const;
}
