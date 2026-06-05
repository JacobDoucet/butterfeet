// This file is auto-generated. DO NOT EDIT.

import { useCallback } from 'react';
import { Cart } from '../../model/cart-model';
import { useFormState } from './common';
import {
    CartMutationOptions,
    useCreateCart,
    useUpdateCart,
} from '../tanstack-query/cart-queries';
import { MutationResponse } from '../../api/model';

type UseCartFormStateOptions = {
    initialState: Cart;
    onSuccess?: (res: MutationResponse<Cart>) => void;
    onError?: (error?: any) => void;
    mutationOptions?: CartMutationOptions;
};

export function useCartFormState(options: UseCartFormStateOptions) {
    const { initialState, ...mutationOptions } = options;
    const formState = useFormState(initialState);

    const updateCart = useUpdateCart(options.mutationOptions);
    const createCart = useCreateCart(options.mutationOptions);

    const save = useCallback(() => {
        const opts = {
            onSuccess: options.onSuccess,
            onError: options.onError,
        };
        if (formState.currentState.id) {
            return updateCart.mutate(formState.updates, opts);
        }
        return createCart.mutate(formState.currentState, opts);
    }, [
        formState.currentState,
        formState.updates,
        createCart,
        updateCart,
        options.onSuccess,
        options.onError,
    ]);

    const isLoading = createCart.isLoading || updateCart.isLoading;

    return {
        ...formState,
        save,
        createMutation: createCart,
        updateMutation: updateCart,
        isLoading,
    } as const;
}
