// This file is auto-generated. DO NOT EDIT.

import { useCallback, useMemo, useState } from 'react';
import { GridFilterModel, GridSortModel, GridColDef, GridColumnVisibilityModel } from '@mui/x-data-grid';
import { RegistryApprovedGuestProjection, RegistryApprovedGuestSortParams } from '../../model/registry-approved-guest-model';
import { RegistryApprovedGuestWithRefs } from '../../model/registry-approved-guest-api';
import { OwnerUser } from '../../model/owner-user-model';
import { useSearchOwnerUsers } from '../tanstack-query/owner-user-queries';
import { Registry } from '../../model/registry-model';
import { useSearchRegistrys } from '../tanstack-query/registry-queries';

type Options = {
    defaultSort?: RegistryApprovedGuestSortParams,
    enableUrlQueryStrings?: boolean,
}

export const RegistryApprovedGuestMUIDataGridUrlQueryKeys = {
    sort: 'registryApprovedGuestSort' as const,
    filter: 'registryApprovedGuestFilter' as const,
    page: 'registryApprovedGuestPage' as const,
    pageSize: 'registryApprovedGuestPageSize' as const,
} as const;

export function useRegistryApprovedGuestMuiDataGridSortModel(options: Options = {}) {
    const [sort, setSort] = useState<RegistryApprovedGuestSortParams>(() => {
        if (options.enableUrlQueryStrings) {
            const urlParams = new URLSearchParams(window.location.search);
            const param = urlParams.get(RegistryApprovedGuestMUIDataGridUrlQueryKeys.sort);
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
          next.push({ field: key, sort: sort[key as keyof RegistryApprovedGuestSortParams] === 1 ? 'asc' : 'desc' });
        }
        return next as GridSortModel;
    }, [sort]);

    const setGridSortModel = useCallback((model: GridSortModel) => {
        const next: RegistryApprovedGuestSortParams = {};
        for (const item of model) {
          next[item.field as keyof RegistryApprovedGuestSortParams] = item.sort === 'asc' ? 1 : -1;
        }
        setSort(next);
        if (options.enableUrlQueryStrings) {
            const urlParams = new URLSearchParams(window.location.search);
            urlParams.set(RegistryApprovedGuestMUIDataGridUrlQueryKeys.sort, serializeSortParams(next));
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

export function useRegistryApprovedGuestMuiDataGridFilterModel(options: FilterOptions = {}) {
    const [gridFilterModel, setGridFilterModel] = useState<GridFilterModel | undefined>(() => {
        if (!options.enableUrlQueryStrings) {
            return undefined;
        }
        const urlParams = new URLSearchParams(window.location.search);
        const filterParam = urlParams.get(RegistryApprovedGuestMUIDataGridUrlQueryKeys.filter);
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
            if (!urlParams.has(RegistryApprovedGuestMUIDataGridUrlQueryKeys.filter)) {
                return;
            }
            urlParams.delete(RegistryApprovedGuestMUIDataGridUrlQueryKeys.filter);
            window.history.replaceState({}, '', `${window.location.pathname}?${urlParams.toString()}`);
            return;
        }
        const urlParams = new URLSearchParams(window.location.search);
        if (model) {
            urlParams.set(RegistryApprovedGuestMUIDataGridUrlQueryKeys.filter, serializeFilterModel(model));
        } else {
            urlParams.delete(RegistryApprovedGuestMUIDataGridUrlQueryKeys.filter);
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
            case "accessLevel":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                accessLevelEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                accessLevelEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                accessLevelEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                accessLevelEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                accessLevelNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                accessLevelNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                accessLevelGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                accessLevelGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                accessLevelGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                accessLevelGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                accessLevelGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                accessLevelGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                accessLevelGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                accessLevelGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                accessLevelLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                accessLevelLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                accessLevelLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                accessLevelLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                accessLevelLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                accessLevelLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                accessLevelLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                accessLevelLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                accessLevelIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                accessLevelIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                accessLevelNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                accessLevelNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                accessLevelExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                accessLevelExists: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "created":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                }
            case "emailEnc":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                emailEncEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailEncEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                emailEncEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailEncEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                emailEncNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailEncNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                emailEncGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailEncGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                emailEncGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailEncGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                emailEncGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailEncGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                emailEncGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailEncGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                emailEncLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailEncLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                emailEncLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailEncLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                emailEncLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailEncLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                emailEncLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailEncLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                emailEncIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                emailEncIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                emailEncNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                emailEncNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                emailEncExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                emailEncExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                emailEncLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailEncLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                emailEncNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailEncNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "emailHash":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                emailHashEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailHashEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                emailHashEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailHashEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                emailHashNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailHashNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                emailHashGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailHashGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                emailHashGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailHashGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                emailHashGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailHashGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                emailHashGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailHashGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                emailHashLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailHashLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                emailHashLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailHashLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                emailHashLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailHashLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                emailHashLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailHashLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                emailHashIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                emailHashIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                emailHashNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                emailHashNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                emailHashExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                emailHashExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                emailHashLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailHashLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                emailHashNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                emailHashNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "name":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                nameEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                nameEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                nameEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                nameEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                nameNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                nameNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                nameGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                nameGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                nameGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                nameGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                nameGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                nameGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                nameGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                nameGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                nameLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                nameLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                nameLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                nameLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                nameLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                nameLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                nameLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                nameLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                nameIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                nameIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                nameNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                nameNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                nameExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                nameExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                nameLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                nameLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                nameNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                nameNlike: gridFilterModel?.items?.[0]?.value,
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
            case "registryId":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                registryIdEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                registryIdEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                registryIdEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                registryIdEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                registryIdIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                registryIdIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                registryIdNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                registryIdNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                registryIdExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                registryIdExists: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "status":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                statusEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                statusEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                statusEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                statusEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                statusNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                statusNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                statusGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                statusGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                statusGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                statusGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                statusGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                statusGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                statusGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                statusGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                statusLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                statusLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                statusLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                statusLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                statusLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                statusLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                statusLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                statusLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                statusIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                statusIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                statusNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                statusNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                statusExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                statusExists: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "updated":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                }
            case "updatedByOwnerUser":
                switch(gridFilterModel?.items?.[0]?.operator) { 
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

type RegistryApprovedGuestIdDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue: (obj: RegistryApprovedGuestWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryApprovedGuestWithRefs>['renderCell']; 
};

export const RegistryApprovedGuestIdDataGridColumnKey = 'id' as const;

export function useRegistryApprovedGuestIdDataGridColumn(options: RegistryApprovedGuestIdDataGridColumnOptions) { 

    return useMemo<GridColDef<RegistryApprovedGuestWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryApprovedGuestIdDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryApprovedGuestIdDataGridColumnKey,
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

type RegistryApprovedGuestAccessLevelDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryApprovedGuestWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryApprovedGuestWithRefs>['renderCell'];
    getOptionLabel: (opt: RegistryApprovedGuestWithRefs['registryApprovedGuest']['accessLevel']) => string;
};

export const RegistryApprovedGuestAccessLevelDataGridColumnKey = 'accessLevel' as const;

export function useRegistryApprovedGuestAccessLevelDataGridColumn(options: RegistryApprovedGuestAccessLevelDataGridColumnOptions) { 
    const values = useMemo(() => [
            {
                value: 'ViewShippingAddress',
                label: options.getOptionLabel('ViewShippingAddress'),
            },
            {
                value: 'ReserveOnly',
                label: options.getOptionLabel('ReserveOnly'),
            },
    ], [options.getOptionLabel]);

    return useMemo<GridColDef<RegistryApprovedGuestWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryApprovedGuestAccessLevelDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryApprovedGuestAccessLevelDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryApprovedGuest.accessLevel;
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

type RegistryApprovedGuestCreatedDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryApprovedGuestWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryApprovedGuestWithRefs>['renderCell'];
};

export const RegistryApprovedGuestCreatedDataGridColumnKey = 'created' as const;

export function useRegistryApprovedGuestCreatedDataGridColumn(options: RegistryApprovedGuestCreatedDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryApprovedGuestWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryApprovedGuestCreatedDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryApprovedGuestCreatedDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryApprovedGuest.created;
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

type RegistryApprovedGuestEmailEncDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryApprovedGuestWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryApprovedGuestWithRefs>['renderCell'];
};

export const RegistryApprovedGuestEmailEncDataGridColumnKey = 'emailEnc' as const;

export function useRegistryApprovedGuestEmailEncDataGridColumn(options: RegistryApprovedGuestEmailEncDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryApprovedGuestWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryApprovedGuestEmailEncDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryApprovedGuestEmailEncDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryApprovedGuest.emailEnc;
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

type RegistryApprovedGuestEmailHashDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryApprovedGuestWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryApprovedGuestWithRefs>['renderCell'];
};

export const RegistryApprovedGuestEmailHashDataGridColumnKey = 'emailHash' as const;

export function useRegistryApprovedGuestEmailHashDataGridColumn(options: RegistryApprovedGuestEmailHashDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryApprovedGuestWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryApprovedGuestEmailHashDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryApprovedGuestEmailHashDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryApprovedGuest.emailHash;
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

type RegistryApprovedGuestNameDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryApprovedGuestWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryApprovedGuestWithRefs>['renderCell'];
};

export const RegistryApprovedGuestNameDataGridColumnKey = 'name' as const;

export function useRegistryApprovedGuestNameDataGridColumn(options: RegistryApprovedGuestNameDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryApprovedGuestWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryApprovedGuestNameDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryApprovedGuestNameDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryApprovedGuest.name;
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

type RegistryApprovedGuestOwnerIdDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue: (obj: RegistryApprovedGuestWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryApprovedGuestWithRefs>['renderCell']; 
    optionLabelField: keyof OwnerUser;
};

export const RegistryApprovedGuestOwnerIdDataGridColumnKey = 'ownerId' as const;

export function useRegistryApprovedGuestOwnerIdDataGridColumn(options: RegistryApprovedGuestOwnerIdDataGridColumnOptions) { 
    const ownerUsers = useSearchOwnerUsers({
        queryKey: ['ownerUser-data-grid-options'],
        query: {},
        projection: { id: true, [options.optionLabelField]: true },
    });

    return useMemo<GridColDef<RegistryApprovedGuestWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryApprovedGuestOwnerIdDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryApprovedGuestOwnerIdDataGridColumnKey,
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

type RegistryApprovedGuestRegistryIdDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue: (obj: RegistryApprovedGuestWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryApprovedGuestWithRefs>['renderCell']; 
    optionLabelField: keyof Registry;
};

export const RegistryApprovedGuestRegistryIdDataGridColumnKey = 'registryId' as const;

export function useRegistryApprovedGuestRegistryIdDataGridColumn(options: RegistryApprovedGuestRegistryIdDataGridColumnOptions) { 
    const registrys = useSearchRegistrys({
        queryKey: ['registry-data-grid-options'],
        query: {},
        projection: { id: true, [options.optionLabelField]: true },
    });

    return useMemo<GridColDef<RegistryApprovedGuestWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryApprovedGuestRegistryIdDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryApprovedGuestRegistryIdDataGridColumnKey,
        valueGetter: (_, row) => {
            return options.getValue(row);
        }, 
        type: 'singleSelect',
        valueOptions: registrys.data?.data?.map((obj) => ({
            value: obj.registry.id,
            label: obj.registry[options.optionLabelField],
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

type RegistryApprovedGuestStatusDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryApprovedGuestWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryApprovedGuestWithRefs>['renderCell'];
    getOptionLabel: (opt: RegistryApprovedGuestWithRefs['registryApprovedGuest']['status']) => string;
};

export const RegistryApprovedGuestStatusDataGridColumnKey = 'status' as const;

export function useRegistryApprovedGuestStatusDataGridColumn(options: RegistryApprovedGuestStatusDataGridColumnOptions) { 
    const values = useMemo(() => [
            {
                value: 'Active',
                label: options.getOptionLabel('Active'),
            },
            {
                value: 'Revoked',
                label: options.getOptionLabel('Revoked'),
            },
            {
                value: 'Blocked',
                label: options.getOptionLabel('Blocked'),
            },
    ], [options.getOptionLabel]);

    return useMemo<GridColDef<RegistryApprovedGuestWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryApprovedGuestStatusDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryApprovedGuestStatusDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryApprovedGuest.status;
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

type RegistryApprovedGuestUpdatedDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryApprovedGuestWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryApprovedGuestWithRefs>['renderCell'];
};

export const RegistryApprovedGuestUpdatedDataGridColumnKey = 'updated' as const;

export function useRegistryApprovedGuestUpdatedDataGridColumn(options: RegistryApprovedGuestUpdatedDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryApprovedGuestWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryApprovedGuestUpdatedDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryApprovedGuestUpdatedDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryApprovedGuest.updated;
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

type RegistryApprovedGuestUpdatedByOwnerUserDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryApprovedGuestWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryApprovedGuestWithRefs>['renderCell'];
};

export const RegistryApprovedGuestUpdatedByOwnerUserDataGridColumnKey = 'updatedByOwnerUser' as const;

export function useRegistryApprovedGuestUpdatedByOwnerUserDataGridColumn(options: RegistryApprovedGuestUpdatedByOwnerUserDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryApprovedGuestWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryApprovedGuestUpdatedByOwnerUserDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryApprovedGuestUpdatedByOwnerUserDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryApprovedGuest.updatedByOwnerUser;
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

export function getRegistryApprovedGuestColumnVisibilityModel(
    defaultValue: boolean,
    projection: RegistryApprovedGuestProjection,
): GridColumnVisibilityModel {
    return {
        id: projection.id ?? false,
        accessLevel: projection.accessLevel ?? false,
        created: projection.created ?? false,
        emailEnc: projection.emailEnc ?? false,
        emailHash: projection.emailHash ?? false,
        name: projection.name ?? false,
        ownerId: projection.ownerId ?? false,
        registryId: projection.registryId ?? false,
        status: projection.status ?? false,
        updated: projection.updated ?? false,
        updatedByOwnerUser: projection.updatedByOwnerUser ?? false,
    };
}

// URL query string helpers
function parseUrlSortParams(sortParam: string): RegistryApprovedGuestSortParams | null {
    try {
        return JSON.parse(sortParam);
    } catch (error) {
        return null;
    }
}

function serializeSortParams(sort: RegistryApprovedGuestSortParams): string {
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