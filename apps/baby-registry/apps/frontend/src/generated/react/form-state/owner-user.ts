// This file is auto-generated. DO NOT EDIT.

import { useCallback } from 'react';
import { OwnerUser } from '../../model/owner-user-model';
import { useFormState } from './common';
import {
    OwnerUserMutationOptions,
    useCreateOwnerUser,
    useUpdateOwnerUser,
} from '../tanstack-query/owner-user-queries';
import { MutationResponse } from '../../api/model';

type UseOwnerUserFormStateOptions = {
    initialState: OwnerUser;
    onSuccess?: (res: MutationResponse<OwnerUser>) => void;
    onError?: (error?: any) => void;
    mutationOptions?: OwnerUserMutationOptions;
};

export function useOwnerUserFormState(options: UseOwnerUserFormStateOptions) {
    const { initialState, ...mutationOptions } = options;
    const formState = useFormState(initialState);

    const updateOwnerUser = useUpdateOwnerUser(options.mutationOptions);
    const createOwnerUser = useCreateOwnerUser(options.mutationOptions);

    const save = useCallback(() => {
        const opts = {
            onSuccess: options.onSuccess,
            onError: options.onError,
        };
        if (formState.currentState.id) {
            return updateOwnerUser.mutate(formState.updates, opts);
        }
        return createOwnerUser.mutate(formState.currentState, opts);
    }, [
        formState.currentState,
        formState.updates,
        createOwnerUser,
        updateOwnerUser,
        options.onSuccess,
        options.onError,
    ]);

    const isLoading = createOwnerUser.isLoading || updateOwnerUser.isLoading;

    return {
        ...formState,
        save,
        createMutation: createOwnerUser,
        updateMutation: updateOwnerUser,
        isLoading,
    } as const;
}
