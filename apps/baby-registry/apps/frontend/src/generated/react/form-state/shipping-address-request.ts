// This file is auto-generated. DO NOT EDIT.

import { useCallback } from 'react';
import { ShippingAddressRequest } from '../../model/shipping-address-request-model';
import { useFormState } from './common';
import {
    ShippingAddressRequestMutationOptions,
    useCreateShippingAddressRequest,
    useUpdateShippingAddressRequest,
} from '../tanstack-query/shipping-address-request-queries';
import { MutationResponse } from '../../api/model';

type UseShippingAddressRequestFormStateOptions = {
    initialState: ShippingAddressRequest;
    onSuccess?: (res: MutationResponse<ShippingAddressRequest>) => void;
    onError?: (error?: any) => void;
    mutationOptions?: ShippingAddressRequestMutationOptions;
};

export function useShippingAddressRequestFormState(options: UseShippingAddressRequestFormStateOptions) {
    const { initialState, ...mutationOptions } = options;
    const formState = useFormState(initialState);

    const updateShippingAddressRequest = useUpdateShippingAddressRequest(options.mutationOptions);
    const createShippingAddressRequest = useCreateShippingAddressRequest(options.mutationOptions);

    const save = useCallback(() => {
        const opts = {
            onSuccess: options.onSuccess,
            onError: options.onError,
        };
        if (formState.currentState.id) {
            return updateShippingAddressRequest.mutate(formState.updates, opts);
        }
        return createShippingAddressRequest.mutate(formState.currentState, opts);
    }, [
        formState.currentState,
        formState.updates,
        createShippingAddressRequest,
        updateShippingAddressRequest,
        options.onSuccess,
        options.onError,
    ]);

    const isLoading = createShippingAddressRequest.isLoading || updateShippingAddressRequest.isLoading;

    return {
        ...formState,
        save,
        createMutation: createShippingAddressRequest,
        updateMutation: updateShippingAddressRequest,
        isLoading,
    } as const;
}
