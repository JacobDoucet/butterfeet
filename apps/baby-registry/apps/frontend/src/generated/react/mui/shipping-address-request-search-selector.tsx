// This file is auto-generated. DO NOT EDIT.

import { useCallback, useMemo, useState, useEffect, ComponentProps } from 'react';
import { Autocomplete, TextField } from '@mui/material';
import { ShippingAddressRequestProjection, ShippingAddressRequestSortParams } from '../../model/shipping-address-request-model';
import { ShippingAddressRequestWithRefs, ShippingAddressRequestSearchQuery } from '../../model/shipping-address-request-api';
import { useSearchShippingAddressRequests } from '../tanstack-query/shipping-address-request-queries';

export type ShippingAddressRequestSearchSelectorOptions = {
    label?: string;
    placeholder?: string;
    multiple?: boolean;
    disabled?: boolean;
    clearable?: boolean;
    searchProjection?: ShippingAddressRequestProjection;
    defaultSort?: ShippingAddressRequestSortParams;
    getOptionLabel: (option: ShippingAddressRequestWithRefs) => string;
    getOptionDescription?: (option: ShippingAddressRequestWithRefs) => string;
    isOptionEqualToValue?: (option: ShippingAddressRequestWithRefs, value: ShippingAddressRequestWithRefs) => boolean;
    filterOptions?: (options: ShippingAddressRequestWithRefs[], state: any) => ShippingAddressRequestWithRefs[];
    filter?: ShippingAddressRequestSearchQuery;
    enableSearch?: boolean;
    searchDebounceMs?: number;
    maxOptions?: number;
    sx?: ComponentProps<typeof Autocomplete>['sx'];
};

export type ShippingAddressRequestSearchSelectorProps = ShippingAddressRequestSearchSelectorOptions & {
    value?: ShippingAddressRequestWithRefs | ShippingAddressRequestWithRefs[] | null;
    onChange: (value: ShippingAddressRequestWithRefs | ShippingAddressRequestWithRefs[] | null) => void;
    onInputChange?: (inputValue: string) => void;
    size?: 'small' | 'medium';
};

export function useShippingAddressRequestSearchSelector(options: ShippingAddressRequestSearchSelectorOptions) {
    const [inputValue, setInputValue] = useState('');
    const [searchQuery, setSearchQuery] = useState(() => options.filter || {});

    // Debounced search effect
    const debouncedInputValue = useDebounce(inputValue, options.searchDebounceMs || 300);

    // Update search query when input changes (if search is enabled)
    useEffect(() => {
        if (!options.enableSearch || !debouncedInputValue.trim()) {
            setSearchQuery(options.filter || {});
            return;
        }

        // Create a text search query - this would depend on your backend implementation
        // For now, assuming there's a text search field or similar
        const textSearchQuery = {
            ...options.filter,
            // Add text search logic here based on your backend API
            // Example: textSearch: debouncedInputValue
        };
        
        setSearchQuery(textSearchQuery);
    }, [debouncedInputValue, options.filter, options.enableSearch]);

    // Fetch data
    const { data: searchResults, isLoading, error } = useSearchShippingAddressRequests({
        query: searchQuery,
        projection: options.searchProjection || { id: true },
        sort: options.defaultSort,
        limit: options.maxOptions || 100,
    });

    const options_data = useMemo(() => {
        return searchResults?.data || [];
    }, [searchResults?.data]);

    const autocompleteProps = useMemo(() => ({
        options: options_data,
        loading: isLoading,
        multiple: options.multiple || false,
        disabled: options.disabled || false,
        disableClearable: !(options.clearable ?? true),
        getOptionLabel: options.getOptionLabel,
        isOptionEqualToValue: options.isOptionEqualToValue || ((option, value) => option.shippingAddressRequest.id === value.shippingAddressRequest.id),
        filterOptions: options.filterOptions || ((opts, state) => {
            // If search is enabled, don't filter client-side (server handles it)
            if (options.enableSearch) {
                return opts;
            }
            // Default MUI filtering
            return opts.filter(opt => 
                options.getOptionLabel(opt).toLowerCase().includes(state.inputValue.toLowerCase())
            );
        }),
    }), [options_data, isLoading, options]);

    return {
        autocompleteProps,
        inputValue,
        setInputValue,
        searchResults,
        isLoading,
        error,
        options: options_data,
    };
}

export type UseShippingAddressRequestSearchSelectorResult = {
    autocompleteProps: ReturnType<typeof useShippingAddressRequestSearchSelector>['autocompleteProps'];
    inputValue: string;
    setInputValue: (value: string) => void;
    searchResults: ReturnType<typeof useSearchShippingAddressRequests>['data'];
    isLoading: boolean;
    error: ReturnType<typeof useSearchShippingAddressRequests>['error'];
    options: ShippingAddressRequestWithRefs[];
};

// Simple debounce hook
function useDebounce<T>(value: T, delay: number): T {
    const [debouncedValue, setDebouncedValue] = useState<T>(value);

    useEffect(() => {
        const handler = setTimeout(() => {
            setDebouncedValue(value);
        }, delay);

        return () => {
            clearTimeout(handler);
        };
    }, [value, delay]);

    return debouncedValue;
}

// Preset configurations for common use cases
export function useShippingAddressRequestSearchSelectorPresets() {
    return {
        // Basic selector with minimal data
        minimal: (): ShippingAddressRequestSearchSelectorOptions => ({
            searchProjection: { id: true },
            getOptionLabel: (option) => option.shippingAddressRequest.id || '',
            enableSearch: false,
        }),

        // Full-featured selector with search
        searchable: (searchProjection?: ShippingAddressRequestProjection): ShippingAddressRequestSearchSelectorOptions => ({
            searchProjection: searchProjection || { id: true },
            getOptionLabel: (option) => option.shippingAddressRequest.id || '',
            enableSearch: true,
            searchDebounceMs: 300,
            clearable: true,
        }),

        // Multiple selection with chips
        multiple: (searchProjection?: ShippingAddressRequestProjection): ShippingAddressRequestSearchSelectorOptions => ({
            multiple: true,
            searchProjection: searchProjection || { id: true },
            getOptionLabel: (option) => option.shippingAddressRequest.id || '',
            enableSearch: true,
            clearable: true,
        }),
    };
}

// Main ShippingAddressRequest selector component
export function ShippingAddressRequestSearchSelector(props: ShippingAddressRequestSearchSelectorProps) {
    const {
        value,
        onChange,
        onInputChange,
        label,
        placeholder,
        ...options
    } = props;

    const {
        autocompleteProps,
        inputValue,
        setInputValue,
    } = useShippingAddressRequestSearchSelector(options);

    const handleInputChange = useCallback((event: React.SyntheticEvent, newInputValue: string) => {
        setInputValue(newInputValue);
        onInputChange?.(newInputValue);
    }, [setInputValue, onInputChange]);

    return (
        <Autocomplete<ShippingAddressRequestWithRefs, boolean, boolean, boolean>
            {...autocompleteProps}
            sx={options.sx}
            size={props.size || 'small'}
            value={value}
            onChange={(_, newValue) => {
                onChange(newValue as ShippingAddressRequestWithRefs | ShippingAddressRequestWithRefs[] | null);
            }}
            getOptionLabel={(option) => {
                if (typeof option === 'string') {
                    return option;
                }
                return options.getOptionLabel(option);
            }}
            inputValue={inputValue}
            onInputChange={handleInputChange}
            renderInput={(params) => (
                <TextField
                    {...params}
                    label={label}
                    placeholder={placeholder}
                    variant="outlined"
                    fullWidth
                />
            )}
        />
    );
}
