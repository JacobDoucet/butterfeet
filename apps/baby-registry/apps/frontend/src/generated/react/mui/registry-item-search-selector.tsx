// This file is auto-generated. DO NOT EDIT.

import { useCallback, useMemo, useState, useEffect, ComponentProps } from 'react';
import { Autocomplete, TextField } from '@mui/material';
import { RegistryItemProjection, RegistryItemSortParams } from '../../model/registry-item-model';
import { RegistryItemWithRefs, RegistryItemSearchQuery } from '../../model/registry-item-api';
import { useSearchRegistryItems } from '../tanstack-query/registry-item-queries';

export type RegistryItemSearchSelectorOptions = {
    label?: string;
    placeholder?: string;
    multiple?: boolean;
    disabled?: boolean;
    clearable?: boolean;
    searchProjection?: RegistryItemProjection;
    defaultSort?: RegistryItemSortParams;
    getOptionLabel: (option: RegistryItemWithRefs) => string;
    getOptionDescription?: (option: RegistryItemWithRefs) => string;
    isOptionEqualToValue?: (option: RegistryItemWithRefs, value: RegistryItemWithRefs) => boolean;
    filterOptions?: (options: RegistryItemWithRefs[], state: any) => RegistryItemWithRefs[];
    filter?: RegistryItemSearchQuery;
    enableSearch?: boolean;
    searchDebounceMs?: number;
    maxOptions?: number;
    sx?: ComponentProps<typeof Autocomplete>['sx'];
};

export type RegistryItemSearchSelectorProps = RegistryItemSearchSelectorOptions & {
    value?: RegistryItemWithRefs | RegistryItemWithRefs[] | null;
    onChange: (value: RegistryItemWithRefs | RegistryItemWithRefs[] | null) => void;
    onInputChange?: (inputValue: string) => void;
    size?: 'small' | 'medium';
};

export function useRegistryItemSearchSelector(options: RegistryItemSearchSelectorOptions) {
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
    const { data: searchResults, isLoading, error } = useSearchRegistryItems({
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
        isOptionEqualToValue: options.isOptionEqualToValue || ((option, value) => option.registryItem.id === value.registryItem.id),
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

export type UseRegistryItemSearchSelectorResult = {
    autocompleteProps: ReturnType<typeof useRegistryItemSearchSelector>['autocompleteProps'];
    inputValue: string;
    setInputValue: (value: string) => void;
    searchResults: ReturnType<typeof useSearchRegistryItems>['data'];
    isLoading: boolean;
    error: ReturnType<typeof useSearchRegistryItems>['error'];
    options: RegistryItemWithRefs[];
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
export function useRegistryItemSearchSelectorPresets() {
    return {
        // Basic selector with minimal data
        minimal: (): RegistryItemSearchSelectorOptions => ({
            searchProjection: { id: true },
            getOptionLabel: (option) => option.registryItem.id || '',
            enableSearch: false,
        }),

        // Full-featured selector with search
        searchable: (searchProjection?: RegistryItemProjection): RegistryItemSearchSelectorOptions => ({
            searchProjection: searchProjection || { id: true },
            getOptionLabel: (option) => option.registryItem.id || '',
            enableSearch: true,
            searchDebounceMs: 300,
            clearable: true,
        }),

        // Multiple selection with chips
        multiple: (searchProjection?: RegistryItemProjection): RegistryItemSearchSelectorOptions => ({
            multiple: true,
            searchProjection: searchProjection || { id: true },
            getOptionLabel: (option) => option.registryItem.id || '',
            enableSearch: true,
            clearable: true,
        }),
    };
}

// Main RegistryItem selector component
export function RegistryItemSearchSelector(props: RegistryItemSearchSelectorProps) {
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
    } = useRegistryItemSearchSelector(options);

    const handleInputChange = useCallback((event: React.SyntheticEvent, newInputValue: string) => {
        setInputValue(newInputValue);
        onInputChange?.(newInputValue);
    }, [setInputValue, onInputChange]);

    return (
        <Autocomplete<RegistryItemWithRefs, boolean, boolean, boolean>
            {...autocompleteProps}
            sx={options.sx}
            size={props.size || 'small'}
            value={value}
            onChange={(_, newValue) => {
                onChange(newValue as RegistryItemWithRefs | RegistryItemWithRefs[] | null);
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
