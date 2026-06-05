// This file is auto-generated. DO NOT EDIT.

import { useCallback, useMemo, useState, useEffect, ComponentProps } from 'react';
import { Autocomplete, TextField } from '@mui/material';
import { RegistryPaymentMethodProjection, RegistryPaymentMethodSortParams } from '../../model/registry-payment-method-model';
import { RegistryPaymentMethodWithRefs, RegistryPaymentMethodSearchQuery } from '../../model/registry-payment-method-api';
import { useSearchRegistryPaymentMethods } from '../tanstack-query/registry-payment-method-queries';

export type RegistryPaymentMethodSearchSelectorOptions = {
    label?: string;
    placeholder?: string;
    multiple?: boolean;
    disabled?: boolean;
    clearable?: boolean;
    searchProjection?: RegistryPaymentMethodProjection;
    defaultSort?: RegistryPaymentMethodSortParams;
    getOptionLabel: (option: RegistryPaymentMethodWithRefs) => string;
    getOptionDescription?: (option: RegistryPaymentMethodWithRefs) => string;
    isOptionEqualToValue?: (option: RegistryPaymentMethodWithRefs, value: RegistryPaymentMethodWithRefs) => boolean;
    filterOptions?: (options: RegistryPaymentMethodWithRefs[], state: any) => RegistryPaymentMethodWithRefs[];
    filter?: RegistryPaymentMethodSearchQuery;
    enableSearch?: boolean;
    searchDebounceMs?: number;
    maxOptions?: number;
    sx?: ComponentProps<typeof Autocomplete>['sx'];
};

export type RegistryPaymentMethodSearchSelectorProps = RegistryPaymentMethodSearchSelectorOptions & {
    value?: RegistryPaymentMethodWithRefs | RegistryPaymentMethodWithRefs[] | null;
    onChange: (value: RegistryPaymentMethodWithRefs | RegistryPaymentMethodWithRefs[] | null) => void;
    onInputChange?: (inputValue: string) => void;
    size?: 'small' | 'medium';
};

export function useRegistryPaymentMethodSearchSelector(options: RegistryPaymentMethodSearchSelectorOptions) {
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
    const { data: searchResults, isLoading, error } = useSearchRegistryPaymentMethods({
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
        isOptionEqualToValue: options.isOptionEqualToValue || ((option, value) => option.registryPaymentMethod.id === value.registryPaymentMethod.id),
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

export type UseRegistryPaymentMethodSearchSelectorResult = {
    autocompleteProps: ReturnType<typeof useRegistryPaymentMethodSearchSelector>['autocompleteProps'];
    inputValue: string;
    setInputValue: (value: string) => void;
    searchResults: ReturnType<typeof useSearchRegistryPaymentMethods>['data'];
    isLoading: boolean;
    error: ReturnType<typeof useSearchRegistryPaymentMethods>['error'];
    options: RegistryPaymentMethodWithRefs[];
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
export function useRegistryPaymentMethodSearchSelectorPresets() {
    return {
        // Basic selector with minimal data
        minimal: (): RegistryPaymentMethodSearchSelectorOptions => ({
            searchProjection: { id: true },
            getOptionLabel: (option) => option.registryPaymentMethod.id || '',
            enableSearch: false,
        }),

        // Full-featured selector with search
        searchable: (searchProjection?: RegistryPaymentMethodProjection): RegistryPaymentMethodSearchSelectorOptions => ({
            searchProjection: searchProjection || { id: true },
            getOptionLabel: (option) => option.registryPaymentMethod.id || '',
            enableSearch: true,
            searchDebounceMs: 300,
            clearable: true,
        }),

        // Multiple selection with chips
        multiple: (searchProjection?: RegistryPaymentMethodProjection): RegistryPaymentMethodSearchSelectorOptions => ({
            multiple: true,
            searchProjection: searchProjection || { id: true },
            getOptionLabel: (option) => option.registryPaymentMethod.id || '',
            enableSearch: true,
            clearable: true,
        }),
    };
}

// Main RegistryPaymentMethod selector component
export function RegistryPaymentMethodSearchSelector(props: RegistryPaymentMethodSearchSelectorProps) {
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
    } = useRegistryPaymentMethodSearchSelector(options);

    const handleInputChange = useCallback((event: React.SyntheticEvent, newInputValue: string) => {
        setInputValue(newInputValue);
        onInputChange?.(newInputValue);
    }, [setInputValue, onInputChange]);

    return (
        <Autocomplete<RegistryPaymentMethodWithRefs, boolean, boolean, boolean>
            {...autocompleteProps}
            sx={options.sx}
            size={props.size || 'small'}
            value={value}
            onChange={(_, newValue) => {
                onChange(newValue as RegistryPaymentMethodWithRefs | RegistryPaymentMethodWithRefs[] | null);
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
