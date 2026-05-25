// This file is auto-generated. DO NOT EDIT.

import { useCallback, useMemo, useState } from 'react';
import { GridFilterModel, GridSortModel, GridColDef, GridColumnVisibilityModel } from '@mui/x-data-grid';
import { RegistryProjection, RegistrySortParams } from '../../model/registry-model';
import { RegistryWithRefs } from '../../model/registry-api';
import { OwnerUser } from '../../model/owner-user-model';
import { useSearchOwnerUsers } from '../tanstack-query/owner-user-queries';

type Options = {
    defaultSort?: RegistrySortParams,
    enableUrlQueryStrings?: boolean,
}

export const RegistryMUIDataGridUrlQueryKeys = {
    sort: 'registrySort' as const,
    filter: 'registryFilter' as const,
    page: 'registryPage' as const,
    pageSize: 'registryPageSize' as const,
} as const;

export function useRegistryMuiDataGridSortModel(options: Options = {}) {
    const [sort, setSort] = useState<RegistrySortParams>(() => {
        if (options.enableUrlQueryStrings) {
            const urlParams = new URLSearchParams(window.location.search);
            const param = urlParams.get(RegistryMUIDataGridUrlQueryKeys.sort);
            if (!param) {
                return options.defaultSort || {};
            }
            return parseUrlSortParams(param) ?? options.defaultSort ?? {};
        }
        return options.defaultSort || {};
    });

    const gridSortModel = useMemo<GridSortModel>(() => {
        const next = [];
        for (const key in sort) {
          next.push({ field: key, sort: sort[key as keyof RegistrySortParams] === 1 ? 'asc' : 'desc' });
        }
        return next as GridSortModel;
    }, [sort]);

    const setGridSortModel = useCallback((model: GridSortModel) => {
        const next: RegistrySortParams = {};
        for (const item of model) {
          next[item.field as keyof RegistrySortParams] = item.sort === 'asc' ? 1 : -1;
        }
        setSort(next);
        if (options.enableUrlQueryStrings) {
            const urlParams = new URLSearchParams(window.location.search);
            urlParams.set(RegistryMUIDataGridUrlQueryKeys.sort, serializeSortParams(next));
            window.history.replaceState({}, '', `${window.location.pathname}?${urlParams.toString()}`);
        }
    }, []);

    return {
        sort,
        setSort,
        gridSortModel,
        setGridSortModel,
    };
}

type FilterOptions = {
    enableUrlQueryStrings?: boolean,
}

export function useRegistryMuiDataGridFilterModel(options: FilterOptions = {}) {
    const [gridFilterModel, setGridFilterModel] = useState<GridFilterModel | undefined>(() => {
        if (!options.enableUrlQueryStrings) {
            return undefined;
        }
        const urlParams = new URLSearchParams(window.location.search);
        const filterParam = urlParams.get(RegistryMUIDataGridUrlQueryKeys.filter);
        if (filterParam) {
            return parseUrlFilterModel(filterParam) || { items: [] };
        }
        return undefined;
    });

    const updateGridFilterModel = useCallback((model: GridFilterModel | undefined) => {
        setGridFilterModel(model);
        if (!options.enableUrlQueryStrings) {
            return;
        }
        if (!model?.items?.some((item) => item.value !== undefined)) {
            // If no filters are set, remove the filter param from the URL
            const urlParams = new URLSearchParams(window.location.search);
            if (!urlParams.has(RegistryMUIDataGridUrlQueryKeys.filter)) {
                return;
            }
            urlParams.delete(RegistryMUIDataGridUrlQueryKeys.filter);
            window.history.replaceState({}, '', `${window.location.pathname}?${urlParams.toString()}`);
            return;
        }
        const urlParams = new URLSearchParams(window.location.search);
        if (model) {
            urlParams.set(RegistryMUIDataGridUrlQueryKeys.filter, serializeFilterModel(model));
        } else {
            urlParams.delete(RegistryMUIDataGridUrlQueryKeys.filter);
        }
        window.history.replaceState({}, '', `${window.location.pathname}?${urlParams.toString()}`);
    }, [options.enableUrlQueryStrings]);

    const { searchQuery, searchQueryExcludingIncompleteFilters } = useMemo(() => {
        switch(gridFilterModel?.items?.[0]?.field) {
            case "id":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                idEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                idEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                idEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                idEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                idIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                idIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                idNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                idNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                idExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                idExists: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "addressAccessMode":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                addressAccessModeEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                addressAccessModeEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                addressAccessModeEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                addressAccessModeEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                addressAccessModeNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                addressAccessModeNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                addressAccessModeGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                addressAccessModeGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                addressAccessModeGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                addressAccessModeGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                addressAccessModeGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                addressAccessModeGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                addressAccessModeGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                addressAccessModeGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                addressAccessModeLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                addressAccessModeLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                addressAccessModeLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                addressAccessModeLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                addressAccessModeLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                addressAccessModeLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                addressAccessModeLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                addressAccessModeLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                addressAccessModeIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                addressAccessModeIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                addressAccessModeNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                addressAccessModeNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                addressAccessModeExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                addressAccessModeExists: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "coverImageUrl":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                coverImageUrlEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                coverImageUrlEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                coverImageUrlEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                coverImageUrlEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                coverImageUrlNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                coverImageUrlNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                coverImageUrlGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                coverImageUrlGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                coverImageUrlGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                coverImageUrlGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                coverImageUrlGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                coverImageUrlGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                coverImageUrlGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                coverImageUrlGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                coverImageUrlLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                coverImageUrlLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                coverImageUrlLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                coverImageUrlLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                coverImageUrlLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                coverImageUrlLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                coverImageUrlLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                coverImageUrlLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                coverImageUrlIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                coverImageUrlIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                coverImageUrlNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                coverImageUrlNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                coverImageUrlExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                coverImageUrlExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                coverImageUrlLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                coverImageUrlLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                coverImageUrlNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                coverImageUrlNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "created":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                }
            case "dueDate":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                dueDateEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                dueDateEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                dueDateEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                dueDateEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                dueDateNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                dueDateNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                dueDateGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                dueDateGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                dueDateGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                dueDateGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                dueDateGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                dueDateGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                dueDateGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                dueDateGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                dueDateLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                dueDateLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                dueDateLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                dueDateLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                dueDateLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                dueDateLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                dueDateLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                dueDateLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                dueDateIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                dueDateIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                dueDateNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                dueDateNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                dueDateExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                dueDateExists: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "isPublic":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                isPublicEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                isPublicEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                isPublicEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                isPublicEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                isPublicNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                isPublicNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                isPublicGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                isPublicGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                isPublicGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                isPublicGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                isPublicGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                isPublicGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                isPublicGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                isPublicGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                isPublicLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                isPublicLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                isPublicLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                isPublicLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                isPublicLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                isPublicLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                isPublicLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                isPublicLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                isPublicIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                isPublicIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                isPublicNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                isPublicNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                isPublicExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                isPublicExists: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "ownerId":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                ownerIdEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                ownerIdEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                ownerIdEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                ownerIdEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                ownerIdIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                ownerIdIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                ownerIdNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                ownerIdNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                ownerIdExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                ownerIdExists: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "parentNames":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                parentNamesEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                parentNamesEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                parentNamesEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                parentNamesEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                parentNamesNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                parentNamesNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                parentNamesGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                parentNamesGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                parentNamesGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                parentNamesGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                parentNamesGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                parentNamesGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                parentNamesGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                parentNamesGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                parentNamesLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                parentNamesLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                parentNamesLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                parentNamesLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                parentNamesLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                parentNamesLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                parentNamesLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                parentNamesLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                parentNamesIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                parentNamesIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                parentNamesNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                parentNamesNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                parentNamesExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                parentNamesExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                parentNamesLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                parentNamesLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                parentNamesNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                parentNamesNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "shippingCity":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                shippingCityEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCityEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                shippingCityEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCityEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                shippingCityNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCityNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                shippingCityGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCityGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                shippingCityGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCityGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                shippingCityGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCityGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                shippingCityGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCityGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                shippingCityLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCityLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                shippingCityLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCityLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                shippingCityLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCityLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                shippingCityLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCityLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                shippingCityIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                shippingCityIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                shippingCityNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                shippingCityNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                shippingCityExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingCityExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                shippingCityLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCityLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                shippingCityNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCityNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "shippingCountry":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                shippingCountryEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCountryEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                shippingCountryEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCountryEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                shippingCountryNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCountryNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                shippingCountryGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCountryGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                shippingCountryGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCountryGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                shippingCountryGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCountryGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                shippingCountryGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCountryGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                shippingCountryLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCountryLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                shippingCountryLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCountryLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                shippingCountryLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCountryLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                shippingCountryLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCountryLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                shippingCountryIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                shippingCountryIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                shippingCountryNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                shippingCountryNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                shippingCountryExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingCountryExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                shippingCountryLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCountryLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                shippingCountryNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingCountryNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "shippingDeliveryNotes":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                shippingDeliveryNotesEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingDeliveryNotesEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                shippingDeliveryNotesEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingDeliveryNotesEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                shippingDeliveryNotesNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingDeliveryNotesNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                shippingDeliveryNotesGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingDeliveryNotesGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                shippingDeliveryNotesGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingDeliveryNotesGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                shippingDeliveryNotesGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingDeliveryNotesGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                shippingDeliveryNotesGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingDeliveryNotesGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                shippingDeliveryNotesLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingDeliveryNotesLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                shippingDeliveryNotesLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingDeliveryNotesLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                shippingDeliveryNotesLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingDeliveryNotesLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                shippingDeliveryNotesLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingDeliveryNotesLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                shippingDeliveryNotesIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                shippingDeliveryNotesIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                shippingDeliveryNotesNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                shippingDeliveryNotesNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                shippingDeliveryNotesExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingDeliveryNotesExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                shippingDeliveryNotesLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingDeliveryNotesLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                shippingDeliveryNotesNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingDeliveryNotesNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "shippingLine1":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                shippingLine1Eq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine1Eq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                shippingLine1Eq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine1Eq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                shippingLine1Ne: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine1Ne: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                shippingLine1Gt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine1Gt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                shippingLine1Gt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine1Gt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                shippingLine1Gte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine1Gte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                shippingLine1Gte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine1Gte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                shippingLine1Lt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine1Lt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                shippingLine1Lt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine1Lt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                shippingLine1Lte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine1Lte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                shippingLine1Lte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine1Lte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                shippingLine1In: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                shippingLine1In: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                shippingLine1Nin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                shippingLine1Nin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                shippingLine1Exists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingLine1Exists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                shippingLine1Like: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine1Like: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                shippingLine1Nlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine1Nlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "shippingLine2":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                shippingLine2Eq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine2Eq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                shippingLine2Eq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine2Eq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                shippingLine2Ne: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine2Ne: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                shippingLine2Gt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine2Gt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                shippingLine2Gt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine2Gt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                shippingLine2Gte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine2Gte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                shippingLine2Gte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine2Gte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                shippingLine2Lt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine2Lt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                shippingLine2Lt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine2Lt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                shippingLine2Lte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine2Lte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                shippingLine2Lte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine2Lte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                shippingLine2In: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                shippingLine2In: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                shippingLine2Nin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                shippingLine2Nin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                shippingLine2Exists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingLine2Exists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                shippingLine2Like: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine2Like: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                shippingLine2Nlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingLine2Nlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "shippingPolicyVersion":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                shippingPolicyVersionEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingPolicyVersionEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                shippingPolicyVersionEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingPolicyVersionEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                shippingPolicyVersionNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingPolicyVersionNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                shippingPolicyVersionGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingPolicyVersionGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                shippingPolicyVersionGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingPolicyVersionGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                shippingPolicyVersionGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingPolicyVersionGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                shippingPolicyVersionGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingPolicyVersionGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                shippingPolicyVersionLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingPolicyVersionLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                shippingPolicyVersionLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingPolicyVersionLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                shippingPolicyVersionLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingPolicyVersionLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                shippingPolicyVersionLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingPolicyVersionLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                shippingPolicyVersionIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                shippingPolicyVersionIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                shippingPolicyVersionNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                shippingPolicyVersionNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                shippingPolicyVersionExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingPolicyVersionExists: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "shippingPostalCode":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                shippingPostalCodeEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingPostalCodeEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                shippingPostalCodeEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingPostalCodeEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                shippingPostalCodeNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingPostalCodeNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                shippingPostalCodeGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingPostalCodeGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                shippingPostalCodeGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingPostalCodeGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                shippingPostalCodeGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingPostalCodeGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                shippingPostalCodeGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingPostalCodeGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                shippingPostalCodeLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingPostalCodeLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                shippingPostalCodeLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingPostalCodeLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                shippingPostalCodeLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingPostalCodeLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                shippingPostalCodeLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingPostalCodeLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                shippingPostalCodeIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                shippingPostalCodeIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                shippingPostalCodeNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                shippingPostalCodeNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                shippingPostalCodeExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingPostalCodeExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                shippingPostalCodeLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingPostalCodeLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                shippingPostalCodeNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingPostalCodeNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "shippingRecipientName":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                shippingRecipientNameEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRecipientNameEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                shippingRecipientNameEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRecipientNameEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                shippingRecipientNameNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRecipientNameNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                shippingRecipientNameGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRecipientNameGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                shippingRecipientNameGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRecipientNameGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                shippingRecipientNameGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRecipientNameGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                shippingRecipientNameGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRecipientNameGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                shippingRecipientNameLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRecipientNameLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                shippingRecipientNameLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRecipientNameLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                shippingRecipientNameLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRecipientNameLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                shippingRecipientNameLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRecipientNameLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                shippingRecipientNameIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                shippingRecipientNameIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                shippingRecipientNameNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                shippingRecipientNameNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                shippingRecipientNameExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingRecipientNameExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                shippingRecipientNameLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRecipientNameLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                shippingRecipientNameNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRecipientNameNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "shippingRegion":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                shippingRegionEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRegionEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                shippingRegionEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRegionEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                shippingRegionNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRegionNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                shippingRegionGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRegionGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                shippingRegionGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRegionGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                shippingRegionGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRegionGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                shippingRegionGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRegionGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                shippingRegionLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRegionLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                shippingRegionLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRegionLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                shippingRegionLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRegionLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                shippingRegionLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRegionLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                shippingRegionIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                shippingRegionIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                shippingRegionNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                shippingRegionNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                shippingRegionExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                shippingRegionExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                shippingRegionLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRegionLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                shippingRegionNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                shippingRegionNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "slug":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                slugEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                slugEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                slugEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                slugEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                slugNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                slugNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                slugGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                slugGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                slugGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                slugGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                slugGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                slugGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                slugGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                slugGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                slugLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                slugLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                slugLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                slugLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                slugLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                slugLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                slugLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                slugLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                slugIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                slugIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                slugNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                slugNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                slugExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                slugExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                slugLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                slugLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                slugNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                slugNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "themeColor":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                themeColorEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                themeColorEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                themeColorEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                themeColorEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                themeColorNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                themeColorNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                themeColorGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                themeColorGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                themeColorGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                themeColorGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                themeColorGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                themeColorGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                themeColorGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                themeColorGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                themeColorLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                themeColorLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                themeColorLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                themeColorLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                themeColorLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                themeColorLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                themeColorLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                themeColorLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                themeColorIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                themeColorIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                themeColorNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                themeColorNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                themeColorExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                themeColorExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                themeColorLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                themeColorLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                themeColorNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                themeColorNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "title":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                titleEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                titleEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                titleEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                titleEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                titleNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                titleNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                titleGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                titleGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                titleGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                titleGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                titleGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                titleGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                titleGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                titleGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                titleLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                titleLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                titleLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                titleLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                titleLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                titleLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                titleLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                titleLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                titleIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                titleIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                titleNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                titleNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                titleExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                titleExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                titleLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                titleLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                titleNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                titleNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "updated":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                }
            case "updatedByOwnerUser":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                }
            case "welcomeMessage":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                welcomeMessageEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                welcomeMessageEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                welcomeMessageEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                welcomeMessageEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                welcomeMessageNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                welcomeMessageNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                welcomeMessageGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                welcomeMessageGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                welcomeMessageGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                welcomeMessageGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                welcomeMessageGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                welcomeMessageGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                welcomeMessageGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                welcomeMessageGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                welcomeMessageLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                welcomeMessageLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                welcomeMessageLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                welcomeMessageLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                welcomeMessageLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                welcomeMessageLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                welcomeMessageLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                welcomeMessageLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                welcomeMessageIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                welcomeMessageIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                welcomeMessageNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                welcomeMessageNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                welcomeMessageExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                welcomeMessageExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                welcomeMessageLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                welcomeMessageLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                welcomeMessageNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                welcomeMessageNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
        }
        if (!!gridFilterModel?.items?.[0]?.field) {
            console.warn(`Unsupported filter model for field ${gridFilterModel?.items?.[0]?.field} and operator ${gridFilterModel?.items?.[0]?.operator}`);
        }
        return { searchQuery: undefined, searchQueryExcludingIncompleteFilters: undefined };
    }, [gridFilterModel]);

    return {
        searchQuery,
        searchQueryExcludingIncompleteFilters,
        gridFilterModel,
        setGridFilterModel: updateGridFilterModel,
    };
}

type RegistryIdDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell']; 
};

export const RegistryIdDataGridColumnKey = 'id' as const;

export function useRegistryIdDataGridColumn(options: RegistryIdDataGridColumnOptions) { 

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryIdDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryIdDataGridColumnKey,
        valueGetter: (_, row) => {
            return options.getValue(row);
        }, 
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryAddressAccessModeDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
    getOptionLabel: (opt: RegistryWithRefs['registry']['addressAccessMode']) => string;
};

export const RegistryAddressAccessModeDataGridColumnKey = 'addressAccessMode' as const;

export function useRegistryAddressAccessModeDataGridColumn(options: RegistryAddressAccessModeDataGridColumnOptions) { 
    const values = useMemo(() => [
            {
                value: 'RequestApproval',
                label: options.getOptionLabel('RequestApproval'),
            },
            {
                value: 'ApprovedGuestsOnly',
                label: options.getOptionLabel('ApprovedGuestsOnly'),
            },
            {
                value: 'Disabled',
                label: options.getOptionLabel('Disabled'),
            },
    ], [options.getOptionLabel]);

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryAddressAccessModeDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryAddressAccessModeDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.addressAccessMode;
        },
        type: 'singleSelect',
        valueOptions: values,
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryCoverImageUrlDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistryCoverImageUrlDataGridColumnKey = 'coverImageUrl' as const;

export function useRegistryCoverImageUrlDataGridColumn(options: RegistryCoverImageUrlDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryCoverImageUrlDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryCoverImageUrlDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.coverImageUrl;
        },
        type: "string",
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryCreatedDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistryCreatedDataGridColumnKey = 'created' as const;

export function useRegistryCreatedDataGridColumn(options: RegistryCreatedDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryCreatedDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryCreatedDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.created;
        },
        type: undefined,
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryDueDateDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => Date;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistryDueDateDataGridColumnKey = 'dueDate' as const;

export function useRegistryDueDateDataGridColumn(options: RegistryDueDateDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryDueDateDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryDueDateDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : new Date(row.registry.dueDate ?? 0);
        },
        type: "date",
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryIsPublicDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => boolean;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistryIsPublicDataGridColumnKey = 'isPublic' as const;

export function useRegistryIsPublicDataGridColumn(options: RegistryIsPublicDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryIsPublicDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryIsPublicDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.isPublic;
        },
        type: "boolean",
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryOwnerIdDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell']; 
    optionLabelField: keyof OwnerUser;
};

export const RegistryOwnerIdDataGridColumnKey = 'ownerId' as const;

export function useRegistryOwnerIdDataGridColumn(options: RegistryOwnerIdDataGridColumnOptions) { 
    const ownerUsers = useSearchOwnerUsers({
        queryKey: ['ownerUser-data-grid-options'],
        query: {},
        projection: { id: true, [options.optionLabelField]: true },
    });

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryOwnerIdDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryOwnerIdDataGridColumnKey,
        valueGetter: (_, row) => {
            return options.getValue(row);
        }, 
        type: 'singleSelect',
        valueOptions: ownerUsers.data?.data?.map((obj) => ({
            value: obj.ownerUser.id,
            label: obj.ownerUser[options.optionLabelField],
        })),
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryParentNamesDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistryParentNamesDataGridColumnKey = 'parentNames' as const;

export function useRegistryParentNamesDataGridColumn(options: RegistryParentNamesDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryParentNamesDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryParentNamesDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.parentNames;
        },
        type: "string",
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryShippingCityDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistryShippingCityDataGridColumnKey = 'shippingCity' as const;

export function useRegistryShippingCityDataGridColumn(options: RegistryShippingCityDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryShippingCityDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryShippingCityDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.shippingCity;
        },
        type: "string",
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryShippingCountryDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistryShippingCountryDataGridColumnKey = 'shippingCountry' as const;

export function useRegistryShippingCountryDataGridColumn(options: RegistryShippingCountryDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryShippingCountryDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryShippingCountryDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.shippingCountry;
        },
        type: "string",
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryShippingDeliveryNotesDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistryShippingDeliveryNotesDataGridColumnKey = 'shippingDeliveryNotes' as const;

export function useRegistryShippingDeliveryNotesDataGridColumn(options: RegistryShippingDeliveryNotesDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryShippingDeliveryNotesDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryShippingDeliveryNotesDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.shippingDeliveryNotes;
        },
        type: "string",
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryShippingLine1DataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistryShippingLine1DataGridColumnKey = 'shippingLine1' as const;

export function useRegistryShippingLine1DataGridColumn(options: RegistryShippingLine1DataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryShippingLine1DataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryShippingLine1DataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.shippingLine1;
        },
        type: "string",
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryShippingLine2DataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistryShippingLine2DataGridColumnKey = 'shippingLine2' as const;

export function useRegistryShippingLine2DataGridColumn(options: RegistryShippingLine2DataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryShippingLine2DataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryShippingLine2DataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.shippingLine2;
        },
        type: "string",
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryShippingPolicyVersionDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => number;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistryShippingPolicyVersionDataGridColumnKey = 'shippingPolicyVersion' as const;

export function useRegistryShippingPolicyVersionDataGridColumn(options: RegistryShippingPolicyVersionDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryShippingPolicyVersionDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryShippingPolicyVersionDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.shippingPolicyVersion;
        },
        type: "number",
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryShippingPostalCodeDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistryShippingPostalCodeDataGridColumnKey = 'shippingPostalCode' as const;

export function useRegistryShippingPostalCodeDataGridColumn(options: RegistryShippingPostalCodeDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryShippingPostalCodeDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryShippingPostalCodeDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.shippingPostalCode;
        },
        type: "string",
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryShippingRecipientNameDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistryShippingRecipientNameDataGridColumnKey = 'shippingRecipientName' as const;

export function useRegistryShippingRecipientNameDataGridColumn(options: RegistryShippingRecipientNameDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryShippingRecipientNameDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryShippingRecipientNameDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.shippingRecipientName;
        },
        type: "string",
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryShippingRegionDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistryShippingRegionDataGridColumnKey = 'shippingRegion' as const;

export function useRegistryShippingRegionDataGridColumn(options: RegistryShippingRegionDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryShippingRegionDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryShippingRegionDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.shippingRegion;
        },
        type: "string",
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistrySlugDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistrySlugDataGridColumnKey = 'slug' as const;

export function useRegistrySlugDataGridColumn(options: RegistrySlugDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistrySlugDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistrySlugDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.slug;
        },
        type: "string",
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryThemeColorDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistryThemeColorDataGridColumnKey = 'themeColor' as const;

export function useRegistryThemeColorDataGridColumn(options: RegistryThemeColorDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryThemeColorDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryThemeColorDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.themeColor;
        },
        type: "string",
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryTitleDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistryTitleDataGridColumnKey = 'title' as const;

export function useRegistryTitleDataGridColumn(options: RegistryTitleDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryTitleDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryTitleDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.title;
        },
        type: "string",
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryUpdatedDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistryUpdatedDataGridColumnKey = 'updated' as const;

export function useRegistryUpdatedDataGridColumn(options: RegistryUpdatedDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryUpdatedDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryUpdatedDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.updated;
        },
        type: undefined,
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryUpdatedByOwnerUserDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistryUpdatedByOwnerUserDataGridColumnKey = 'updatedByOwnerUser' as const;

export function useRegistryUpdatedByOwnerUserDataGridColumn(options: RegistryUpdatedByOwnerUserDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryUpdatedByOwnerUserDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryUpdatedByOwnerUserDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.updatedByOwnerUser;
        },
        type: undefined,
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

type RegistryWelcomeMessageDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryWithRefs>['renderCell'];
};

export const RegistryWelcomeMessageDataGridColumnKey = 'welcomeMessage' as const;

export function useRegistryWelcomeMessageDataGridColumn(options: RegistryWelcomeMessageDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryWelcomeMessageDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryWelcomeMessageDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registry.welcomeMessage;
        },
        type: "string",
        renderCell: options.renderCell,
    }), [
        options.headerName, 
        options.width, 
        options.sortable,
        options.hideable,
        options.getValue,
        options.renderCell,
    ]);
}

export function getRegistryColumnVisibilityModel(
    defaultValue: boolean,
    projection: RegistryProjection,
): GridColumnVisibilityModel {
    return {
        id: projection.id ?? false,
        addressAccessMode: projection.addressAccessMode ?? false,
        coverImageUrl: projection.coverImageUrl ?? false,
        created: projection.created ?? false,
        dueDate: projection.dueDate ?? false,
        isPublic: projection.isPublic ?? false,
        ownerId: projection.ownerId ?? false,
        parentNames: projection.parentNames ?? false,
        shippingCity: projection.shippingCity ?? false,
        shippingCountry: projection.shippingCountry ?? false,
        shippingDeliveryNotes: projection.shippingDeliveryNotes ?? false,
        shippingLine1: projection.shippingLine1 ?? false,
        shippingLine2: projection.shippingLine2 ?? false,
        shippingPolicyVersion: projection.shippingPolicyVersion ?? false,
        shippingPostalCode: projection.shippingPostalCode ?? false,
        shippingRecipientName: projection.shippingRecipientName ?? false,
        shippingRegion: projection.shippingRegion ?? false,
        slug: projection.slug ?? false,
        themeColor: projection.themeColor ?? false,
        title: projection.title ?? false,
        updated: projection.updated ?? false,
        updatedByOwnerUser: projection.updatedByOwnerUser ?? false,
        welcomeMessage: projection.welcomeMessage ?? false,
    };
}

// URL query string helpers
function parseUrlSortParams(sortParam: string): RegistrySortParams | null {
    try {
        return JSON.parse(sortParam);
    } catch (error) {
        return null;
    }
}

function serializeSortParams(sort: RegistrySortParams): string {
    return JSON.stringify(sort);
}

function parseUrlFilterModel(filterParam: string): GridFilterModel | null {
    try {
        return JSON.parse(filterParam);
    } catch (error) {
        return null;
    }
}

function serializeFilterModel(filterModel: GridFilterModel): string {
    return JSON.stringify(filterModel);
}