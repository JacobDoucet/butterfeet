// This file is auto-generated. DO NOT EDIT.

import { useCallback, useMemo, useState } from 'react';
import { GridFilterModel, GridSortModel, GridColDef, GridColumnVisibilityModel } from '@mui/x-data-grid';
import { RegistryItemProjection, RegistryItemSortParams } from '../../model/registry-item-model';
import { RegistryItemWithRefs } from '../../model/registry-item-api';
import { Registry } from '../../model/registry-model';
import { useSearchRegistrys } from '../tanstack-query/registry-queries';

type Options = {
    defaultSort?: RegistryItemSortParams,
    enableUrlQueryStrings?: boolean,
}

export const RegistryItemMUIDataGridUrlQueryKeys = {
    sort: 'registryItemSort' as const,
    filter: 'registryItemFilter' as const,
    page: 'registryItemPage' as const,
    pageSize: 'registryItemPageSize' as const,
} as const;

export function useRegistryItemMuiDataGridSortModel(options: Options = {}) {
    const [sort, setSort] = useState<RegistryItemSortParams>(() => {
        if (options.enableUrlQueryStrings) {
            const urlParams = new URLSearchParams(window.location.search);
            const param = urlParams.get(RegistryItemMUIDataGridUrlQueryKeys.sort);
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
          next.push({ field: key, sort: sort[key as keyof RegistryItemSortParams] === 1 ? 'asc' : 'desc' });
        }
        return next as GridSortModel;
    }, [sort]);

    const setGridSortModel = useCallback((model: GridSortModel) => {
        const next: RegistryItemSortParams = {};
        for (const item of model) {
          next[item.field as keyof RegistryItemSortParams] = item.sort === 'asc' ? 1 : -1;
        }
        setSort(next);
        if (options.enableUrlQueryStrings) {
            const urlParams = new URLSearchParams(window.location.search);
            urlParams.set(RegistryItemMUIDataGridUrlQueryKeys.sort, serializeSortParams(next));
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

export function useRegistryItemMuiDataGridFilterModel(options: FilterOptions = {}) {
    const [gridFilterModel, setGridFilterModel] = useState<GridFilterModel | undefined>(() => {
        if (!options.enableUrlQueryStrings) {
            return undefined;
        }
        const urlParams = new URLSearchParams(window.location.search);
        const filterParam = urlParams.get(RegistryItemMUIDataGridUrlQueryKeys.filter);
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
            if (!urlParams.has(RegistryItemMUIDataGridUrlQueryKeys.filter)) {
                return;
            }
            urlParams.delete(RegistryItemMUIDataGridUrlQueryKeys.filter);
            window.history.replaceState({}, '', `${window.location.pathname}?${urlParams.toString()}`);
            return;
        }
        const urlParams = new URLSearchParams(window.location.search);
        if (model) {
            urlParams.set(RegistryItemMUIDataGridUrlQueryKeys.filter, serializeFilterModel(model));
        } else {
            urlParams.delete(RegistryItemMUIDataGridUrlQueryKeys.filter);
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
            case "created":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                }
            case "currency":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                currencyEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                currencyEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                currencyEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                currencyEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                currencyNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                currencyNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                currencyGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                currencyGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                currencyGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                currencyGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                currencyGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                currencyGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                currencyGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                currencyGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                currencyLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                currencyLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                currencyLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                currencyLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                currencyLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                currencyLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                currencyLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                currencyLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                currencyIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                currencyIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                currencyNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                currencyNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                currencyExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                currencyExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                currencyLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                currencyLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                currencyNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                currencyNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "description":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                descriptionEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                descriptionEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                descriptionEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                descriptionEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                descriptionNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                descriptionNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                descriptionGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                descriptionGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                descriptionGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                descriptionGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                descriptionGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                descriptionGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                descriptionGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                descriptionGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                descriptionLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                descriptionLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                descriptionLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                descriptionLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                descriptionLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                descriptionLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                descriptionLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                descriptionLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                descriptionIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                descriptionIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                descriptionNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                descriptionNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                descriptionExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                descriptionExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                descriptionLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                descriptionLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                descriptionNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                descriptionNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "imageUrl":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                imageUrlEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                imageUrlEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                imageUrlEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                imageUrlEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                imageUrlNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                imageUrlNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                imageUrlGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                imageUrlGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                imageUrlGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                imageUrlGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                imageUrlGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                imageUrlGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                imageUrlGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                imageUrlGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                imageUrlLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                imageUrlLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                imageUrlLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                imageUrlLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                imageUrlLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                imageUrlLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                imageUrlLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                imageUrlLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                imageUrlIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                imageUrlIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                imageUrlNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                imageUrlNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                imageUrlExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                imageUrlExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                imageUrlLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                imageUrlLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                imageUrlNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                imageUrlNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "notes":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                notesEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                notesEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                notesEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                notesEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                notesNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                notesNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                notesGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                notesGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                notesGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                notesGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                notesGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                notesGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                notesGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                notesGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                notesLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                notesLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                notesLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                notesLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                notesLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                notesLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                notesLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                notesLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                notesIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                notesIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                notesNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                notesNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                notesExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                notesExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                notesLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                notesLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                notesNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                notesNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "ownerPurchased":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                ownerPurchasedEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                ownerPurchasedEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                ownerPurchasedEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                ownerPurchasedEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                ownerPurchasedNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                ownerPurchasedNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                ownerPurchasedGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                ownerPurchasedGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                ownerPurchasedGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                ownerPurchasedGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                ownerPurchasedGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                ownerPurchasedGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                ownerPurchasedGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                ownerPurchasedGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                ownerPurchasedLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                ownerPurchasedLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                ownerPurchasedLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                ownerPurchasedLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                ownerPurchasedLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                ownerPurchasedLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                ownerPurchasedLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                ownerPurchasedLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                ownerPurchasedIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                ownerPurchasedIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                ownerPurchasedNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                ownerPurchasedNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                ownerPurchasedExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                ownerPurchasedExists: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "position":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                positionEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                positionEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                positionEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                positionEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                positionNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                positionNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                positionGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                positionGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                positionGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                positionGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                positionGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                positionGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                positionGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                positionGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                positionLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                positionLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                positionLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                positionLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                positionLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                positionLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                positionLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                positionLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                positionIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                positionIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                positionNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                positionNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                positionExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                positionExists: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "priceCents":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                priceCentsEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                priceCentsEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                priceCentsEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                priceCentsEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                priceCentsNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                priceCentsNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                priceCentsGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                priceCentsGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                priceCentsGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                priceCentsGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                priceCentsGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                priceCentsGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                priceCentsGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                priceCentsGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                priceCentsLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                priceCentsLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                priceCentsLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                priceCentsLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                priceCentsLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                priceCentsLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                priceCentsLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                priceCentsLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                priceCentsIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                priceCentsIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                priceCentsNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                priceCentsNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                priceCentsExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                priceCentsExists: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "productUrl":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                productUrlEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                productUrlEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                productUrlEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                productUrlEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                productUrlNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                productUrlNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                productUrlGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                productUrlGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                productUrlGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                productUrlGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                productUrlGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                productUrlGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                productUrlGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                productUrlGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                productUrlLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                productUrlLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                productUrlLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                productUrlLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                productUrlLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                productUrlLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                productUrlLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                productUrlLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                productUrlIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                productUrlIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                productUrlNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                productUrlNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                productUrlExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                productUrlExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                productUrlLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                productUrlLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                productUrlNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                productUrlNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "quantity":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                quantityEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                quantityEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                quantityEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                quantityEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                quantityNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                quantityNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                quantityGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                quantityGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                quantityGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                quantityGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                quantityGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                quantityGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                quantityGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                quantityGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                quantityLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                quantityLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                quantityLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                quantityLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                quantityLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                quantityLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                quantityLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                quantityLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                quantityIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                quantityIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                quantityNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                quantityNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                quantityExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                quantityExists: gridFilterModel?.items?.[0]?.value,
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
            case "source":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                sourceEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                sourceEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                sourceEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                sourceEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                sourceNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                sourceNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                sourceGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                sourceGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                sourceGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                sourceGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                sourceGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                sourceGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                sourceGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                sourceGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                sourceLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                sourceLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                sourceLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                sourceLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                sourceLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                sourceLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                sourceLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters:  { 
                                sourceLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                sourceIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                sourceIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                sourceNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                sourceNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                sourceExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                sourceExists: gridFilterModel?.items?.[0]?.value,
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

type RegistryItemIdDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue: (obj: RegistryItemWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryItemWithRefs>['renderCell']; 
};

export const RegistryItemIdDataGridColumnKey = 'id' as const;

export function useRegistryItemIdDataGridColumn(options: RegistryItemIdDataGridColumnOptions) { 

    return useMemo<GridColDef<RegistryItemWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryItemIdDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryItemIdDataGridColumnKey,
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

type RegistryItemCreatedDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryItemWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryItemWithRefs>['renderCell'];
};

export const RegistryItemCreatedDataGridColumnKey = 'created' as const;

export function useRegistryItemCreatedDataGridColumn(options: RegistryItemCreatedDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryItemWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryItemCreatedDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryItemCreatedDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryItem.created;
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

type RegistryItemCurrencyDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryItemWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryItemWithRefs>['renderCell'];
};

export const RegistryItemCurrencyDataGridColumnKey = 'currency' as const;

export function useRegistryItemCurrencyDataGridColumn(options: RegistryItemCurrencyDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryItemWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryItemCurrencyDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryItemCurrencyDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryItem.currency;
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

type RegistryItemDescriptionDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryItemWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryItemWithRefs>['renderCell'];
};

export const RegistryItemDescriptionDataGridColumnKey = 'description' as const;

export function useRegistryItemDescriptionDataGridColumn(options: RegistryItemDescriptionDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryItemWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryItemDescriptionDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryItemDescriptionDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryItem.description;
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

type RegistryItemImageUrlDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryItemWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryItemWithRefs>['renderCell'];
};

export const RegistryItemImageUrlDataGridColumnKey = 'imageUrl' as const;

export function useRegistryItemImageUrlDataGridColumn(options: RegistryItemImageUrlDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryItemWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryItemImageUrlDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryItemImageUrlDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryItem.imageUrl;
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

type RegistryItemNotesDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryItemWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryItemWithRefs>['renderCell'];
};

export const RegistryItemNotesDataGridColumnKey = 'notes' as const;

export function useRegistryItemNotesDataGridColumn(options: RegistryItemNotesDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryItemWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryItemNotesDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryItemNotesDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryItem.notes;
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

type RegistryItemOwnerPurchasedDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryItemWithRefs | undefined) => boolean;
    renderCell?: GridColDef<RegistryItemWithRefs>['renderCell'];
};

export const RegistryItemOwnerPurchasedDataGridColumnKey = 'ownerPurchased' as const;

export function useRegistryItemOwnerPurchasedDataGridColumn(options: RegistryItemOwnerPurchasedDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryItemWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryItemOwnerPurchasedDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryItemOwnerPurchasedDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryItem.ownerPurchased;
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

type RegistryItemPositionDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryItemWithRefs | undefined) => number;
    renderCell?: GridColDef<RegistryItemWithRefs>['renderCell'];
};

export const RegistryItemPositionDataGridColumnKey = 'position' as const;

export function useRegistryItemPositionDataGridColumn(options: RegistryItemPositionDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryItemWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryItemPositionDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryItemPositionDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryItem.position;
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

type RegistryItemPriceCentsDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryItemWithRefs | undefined) => number;
    renderCell?: GridColDef<RegistryItemWithRefs>['renderCell'];
};

export const RegistryItemPriceCentsDataGridColumnKey = 'priceCents' as const;

export function useRegistryItemPriceCentsDataGridColumn(options: RegistryItemPriceCentsDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryItemWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryItemPriceCentsDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryItemPriceCentsDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryItem.priceCents;
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

type RegistryItemProductUrlDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryItemWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryItemWithRefs>['renderCell'];
};

export const RegistryItemProductUrlDataGridColumnKey = 'productUrl' as const;

export function useRegistryItemProductUrlDataGridColumn(options: RegistryItemProductUrlDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryItemWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryItemProductUrlDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryItemProductUrlDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryItem.productUrl;
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

type RegistryItemQuantityDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryItemWithRefs | undefined) => number;
    renderCell?: GridColDef<RegistryItemWithRefs>['renderCell'];
};

export const RegistryItemQuantityDataGridColumnKey = 'quantity' as const;

export function useRegistryItemQuantityDataGridColumn(options: RegistryItemQuantityDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryItemWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryItemQuantityDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryItemQuantityDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryItem.quantity;
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

type RegistryItemRegistryIdDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue: (obj: RegistryItemWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryItemWithRefs>['renderCell']; 
    optionLabelField: keyof Registry;
};

export const RegistryItemRegistryIdDataGridColumnKey = 'registryId' as const;

export function useRegistryItemRegistryIdDataGridColumn(options: RegistryItemRegistryIdDataGridColumnOptions) { 
    const registrys = useSearchRegistrys({
        queryKey: ['registry-data-grid-options'],
        query: {},
        projection: { id: true, [options.optionLabelField]: true },
    });

    return useMemo<GridColDef<RegistryItemWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryItemRegistryIdDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryItemRegistryIdDataGridColumnKey,
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

type RegistryItemSourceDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryItemWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryItemWithRefs>['renderCell'];
    getOptionLabel: (opt: RegistryItemWithRefs['registryItem']['source']) => string;
};

export const RegistryItemSourceDataGridColumnKey = 'source' as const;

export function useRegistryItemSourceDataGridColumn(options: RegistryItemSourceDataGridColumnOptions) { 
    const values = useMemo(() => [
            {
                value: 'Amazon',
                label: options.getOptionLabel('Amazon'),
            },
            {
                value: 'MamasAndPapas',
                label: options.getOptionLabel('MamasAndPapas'),
            },
            {
                value: 'Etsy',
                label: options.getOptionLabel('Etsy'),
            },
            {
                value: 'JohnLewis',
                label: options.getOptionLabel('JohnLewis'),
            },
            {
                value: 'IKEA',
                label: options.getOptionLabel('IKEA'),
            },
            {
                value: 'Other',
                label: options.getOptionLabel('Other'),
            },
    ], [options.getOptionLabel]);

    return useMemo<GridColDef<RegistryItemWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryItemSourceDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryItemSourceDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryItem.source;
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

type RegistryItemTitleDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryItemWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryItemWithRefs>['renderCell'];
};

export const RegistryItemTitleDataGridColumnKey = 'title' as const;

export function useRegistryItemTitleDataGridColumn(options: RegistryItemTitleDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryItemWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryItemTitleDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryItemTitleDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryItem.title;
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

type RegistryItemUpdatedDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryItemWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryItemWithRefs>['renderCell'];
};

export const RegistryItemUpdatedDataGridColumnKey = 'updated' as const;

export function useRegistryItemUpdatedDataGridColumn(options: RegistryItemUpdatedDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryItemWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryItemUpdatedDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryItemUpdatedDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryItem.updated;
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

type RegistryItemUpdatedByOwnerUserDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: RegistryItemWithRefs | undefined) => string;
    renderCell?: GridColDef<RegistryItemWithRefs>['renderCell'];
};

export const RegistryItemUpdatedByOwnerUserDataGridColumnKey = 'updatedByOwnerUser' as const;

export function useRegistryItemUpdatedByOwnerUserDataGridColumn(options: RegistryItemUpdatedByOwnerUserDataGridColumnOptions) {

    return useMemo<GridColDef<RegistryItemWithRefs>>(() => ({
        headerName: options.headerName ?? RegistryItemUpdatedByOwnerUserDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: RegistryItemUpdatedByOwnerUserDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.registryItem.updatedByOwnerUser;
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

export function getRegistryItemColumnVisibilityModel(
    defaultValue: boolean,
    projection: RegistryItemProjection,
): GridColumnVisibilityModel {
    return {
        id: projection.id ?? false,
        created: projection.created ?? false,
        currency: projection.currency ?? false,
        description: projection.description ?? false,
        imageUrl: projection.imageUrl ?? false,
        notes: projection.notes ?? false,
        ownerPurchased: projection.ownerPurchased ?? false,
        position: projection.position ?? false,
        priceCents: projection.priceCents ?? false,
        productUrl: projection.productUrl ?? false,
        quantity: projection.quantity ?? false,
        registryId: projection.registryId ?? false,
        source: projection.source ?? false,
        title: projection.title ?? false,
        updated: projection.updated ?? false,
        updatedByOwnerUser: projection.updatedByOwnerUser ?? false,
    };
}

// URL query string helpers
function parseUrlSortParams(sortParam: string): RegistryItemSortParams | null {
    try {
        return JSON.parse(sortParam);
    } catch (error) {
        return null;
    }
}

function serializeSortParams(sort: RegistryItemSortParams): string {
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