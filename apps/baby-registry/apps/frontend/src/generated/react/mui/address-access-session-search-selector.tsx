// This file is auto-generated. DO NOT EDIT.

import { useCallback, useMemo, useState, useEffect, ComponentProps } from 'react';
import { Autocomplete, TextField } from '@mui/material';
import { AddressAccessSessionProjection, AddressAccessSessionSortParams } from '../../model/address-access-session-model';
import { AddressAccessSessionWithRefs, AddressAccessSessionSearchQuery } from '../../model/address-access-session-api';
import { useSearchAddressAccessSessions } from '../tanstack-query/address-access-session-queries';

export type AddressAccessSessionSearchSelectorOptions = {
    label?: string;
    placeholder?: string;
    multiple?: boolean;
    disabled?: boolean;
    clearable?: boolean;
    searchProjection?: AddressAccessSessionProjection;
    defaultSort?: AddressAccessSessionSortParams;
    getOptionLabel: (option: AddressAccessSessionWithRefs) => string;
    getOptionDescription?: (option: AddressAccessSessionWithRefs) => string;
    isOptionEqualToValue?: (option: AddressAccessSessionWithRefs, value: AddressAccessSessionWithRefs) => boolean;
    filterOptions?: (options: AddressAccessSessionWithRefs[], state: any) => AddressAccessSessionWithRefs[];
    filter?: AddressAccessSessionSearchQuery;
    enableSearch?: boolean;
    searchDebounceMs?: number;
    maxOptions?: number;
    sx?: ComponentProps<typeof Autocomplete>['sx'];
};

export type AddressAccessSessionSearchSelectorProps = AddressAccessSessionSearchSelectorOptions & {
    value?: AddressAccessSessionWithRefs | AddressAccessSessionWithRefs[] | null;
    onChange: (value: AddressAccessSessionWithRefs | AddressAccessSessionWithRefs[] | null) => void;
    onInputChange?: (inputValue: string) => void;
    size?: 'small' | 'medium';
};

export function useAddressAccessSessionSearchSelector(options: AddressAccessSessionSearchSelectorOptions) {
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
    const { data: searchResults, isLoading, error } = useSearchAddressAccessSessions({
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
        isOptionEqualToValue: options.isOptionEqualToValue || ((option, value) => option.addressAccessSession.id === value.addressAccessSession.id),
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

export type UseAddressAccessSessionSearchSelectorResult = {
    autocompleteProps: ReturnType<typeof useAddressAccessSessionSearchSelector>['autocompleteProps'];
    inputValue: string;
    setInputValue: (value: string) => void;
    searchResults: ReturnType<typeof useSearchAddressAccessSessions>['data'];
    isLoading: boolean;
    error: ReturnType<typeof useSearchAddressAccessSessions>['error'];
    options: AddressAccessSessionWithRefs[];
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
export function useAddressAccessSessionSearchSelectorPresets() {
    return {
        // Basic selector with minimal data
        minimal: (): AddressAccessSessionSearchSelectorOptions => ({
            searchProjection: { id: true },
            getOptionLabel: (option) => option.addressAccessSession.id || '',
            enableSearch: false,
        }),

        // Full-featured selector with search
        searchable: (searchProjection?: AddressAccessSessionProjection): AddressAccessSessionSearchSelectorOptions => ({
            searchProjection: searchProjection || { id: true },
            getOptionLabel: (option) => option.addressAccessSession.id || '',
            enableSearch: true,
            searchDebounceMs: 300,
            clearable: true,
        }),

        // Multiple selection with chips
        multiple: (searchProjection?: AddressAccessSessionProjection): AddressAccessSessionSearchSelectorOptions => ({
            multiple: true,
            searchProjection: searchProjection || { id: true },
            getOptionLabel: (option) => option.addressAccessSession.id || '',
            enableSearch: true,
            clearable: true,
        }),
    };
}

// Main AddressAccessSession selector component
export function AddressAccessSessionSearchSelector(props: AddressAccessSessionSearchSelectorProps) {
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
    } = useAddressAccessSessionSearchSelector(options);

    const handleInputChange = useCallback((event: React.SyntheticEvent, newInputValue: string) => {
        setInputValue(newInputValue);
        onInputChange?.(newInputValue);
    }, [setInputValue, onInputChange]);

    return (
        <Autocomplete<AddressAccessSessionWithRefs, boolean, boolean, boolean>
            {...autocompleteProps}
            sx={options.sx}
            size={props.size || 'small'}
            value={value}
            onChange={(_, newValue) => {
                onChange(newValue as AddressAccessSessionWithRefs | AddressAccessSessionWithRefs[] | null);
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
