// This file is auto-generated. DO NOT EDIT.

import { useCallback, useMemo, useState } from 'react';
import { GridFilterModel, GridSortModel, GridColDef, GridColumnVisibilityModel } from '@mui/x-data-grid';
import { ShippingAddressRequestProjection, ShippingAddressRequestSortParams } from '../../model/shipping-address-request-model';
import { ShippingAddressRequestWithRefs } from '../../model/shipping-address-request-api';
import { OwnerUser } from '../../model/owner-user-model';
import { useSearchOwnerUsers } from '../tanstack-query/owner-user-queries';
import { Registry } from '../../model/registry-model';
import { useSearchRegistrys } from '../tanstack-query/registry-queries';
import { RegistryItem } from '../../model/registry-item-model';
import { useSearchRegistryItems } from '../tanstack-query/registry-item-queries';

type Options = {
    defaultSort?: ShippingAddressRequestSortParams,
    enableUrlQueryStrings?: boolean,
}

export const ShippingAddressRequestMUIDataGridUrlQueryKeys = {
    sort: 'shippingAddressRequestSort' as const,
    filter: 'shippingAddressRequestFilter' as const,
    page: 'shippingAddressRequestPage' as const,
    pageSize: 'shippingAddressRequestPageSize' as const,
} as const;

export function useShippingAddressRequestMuiDataGridSortModel(options: Options = {}) {
    const [sort, setSort] = useState<ShippingAddressRequestSortParams>(() => {
        if (options.enableUrlQueryStrings) {
            const urlParams = new URLSearchParams(window.location.search);
            const param = urlParams.get(ShippingAddressRequestMUIDataGridUrlQueryKeys.sort);
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
          next.push({ field: key, sort: sort[key as keyof ShippingAddressRequestSortParams] === 1 ? 'asc' : 'desc' });
        }
        return next as GridSortModel;
    }, [sort]);

    const setGridSortModel = useCallback((model: GridSortModel) => {
        const next: ShippingAddressRequestSortParams = {};
        for (const item of model) {
          next[item.field as keyof ShippingAddressRequestSortParams] = item.sort === 'asc' ? 1 : -1;
        }
        setSort(next);
        if (options.enableUrlQueryStrings) {
            const urlParams = new URLSearchParams(window.location.search);
            urlParams.set(ShippingAddressRequestMUIDataGridUrlQueryKeys.sort, serializeSortParams(next));
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

export function useShippingAddressRequestMuiDataGridFilterModel(options: FilterOptions = {}) {
    const [gridFilterModel, setGridFilterModel] = useState<GridFilterModel | undefined>(() => {
        if (!options.enableUrlQueryStrings) {
            return undefined;
        }
        const urlParams = new URLSearchParams(window.location.search);
        const filterParam = urlParams.get(ShippingAddressRequestMUIDataGridUrlQueryKeys.filter);
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
            if (!urlParams.has(ShippingAddressRequestMUIDataGridUrlQueryKeys.filter)) {
                return;
            }
            urlParams.delete(ShippingAddressRequestMUIDataGridUrlQueryKeys.filter);
            window.history.replaceState({}, '', `${window.location.pathname}?${urlParams.toString()}`);
            return;
        }
        const urlParams = new URLSearchParams(window.location.search);
        if (model) {
            urlParams.set(ShippingAddressRequestMUIDataGridUrlQueryKeys.filter, serializeFilterModel(model));
        } else {
            urlParams.delete(ShippingAddressRequestMUIDataGridUrlQueryKeys.filter);
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
            case "decisionReason":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                decisionReasonEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                decisionReasonEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                decisionReasonEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                decisionReasonEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                decisionReasonNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                decisionReasonNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                decisionReasonGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                decisionReasonGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                decisionReasonGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                decisionReasonGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                decisionReasonGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                decisionReasonGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                decisionReasonGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                decisionReasonGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                decisionReasonLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                decisionReasonLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                decisionReasonLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                decisionReasonLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                decisionReasonLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                decisionReasonLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                decisionReasonLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                decisionReasonLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                decisionReasonIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                decisionReasonIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                decisionReasonNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                decisionReasonNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                decisionReasonExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                decisionReasonExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                decisionReasonLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                decisionReasonLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                decisionReasonNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                decisionReasonNlike: gridFilterModel?.items?.[0]?.value,
                            },
                        };
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
            case "note":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                noteEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                noteEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                noteEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                noteEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                noteNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                noteNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                noteGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                noteGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                noteGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                noteGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                noteGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                noteGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                noteGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                noteGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                noteLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                noteLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                noteLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                noteLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                noteLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                noteLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                noteLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                noteLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                noteIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                noteIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                noteNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                noteNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                noteExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                noteExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                noteLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                noteLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                noteNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                noteNlike: gridFilterModel?.items?.[0]?.value,
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
            case "policyVersion":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                policyVersionEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                policyVersionEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                policyVersionNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                policyVersionGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                policyVersionGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                policyVersionGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                policyVersionGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                policyVersionLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                policyVersionLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                policyVersionLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                policyVersionLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                policyVersionIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                policyVersionIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                policyVersionNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                policyVersionNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                policyVersionExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionExists: gridFilterModel?.items?.[0]?.value,
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
            case "registryItemId":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                registryItemIdEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                registryItemIdEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                registryItemIdEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                registryItemIdEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                registryItemIdIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                registryItemIdIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                registryItemIdNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                registryItemIdNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                registryItemIdExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                registryItemIdExists: gridFilterModel?.items?.[0]?.value,
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

type ShippingAddressRequestIdDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue: (obj: ShippingAddressRequestWithRefs | undefined) => string;
    renderCell?: GridColDef<ShippingAddressRequestWithRefs>['renderCell']; 
};

export const ShippingAddressRequestIdDataGridColumnKey = 'id' as const;

export function useShippingAddressRequestIdDataGridColumn(options: ShippingAddressRequestIdDataGridColumnOptions) { 

    return useMemo<GridColDef<ShippingAddressRequestWithRefs>>(() => ({
        headerName: options.headerName ?? ShippingAddressRequestIdDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: ShippingAddressRequestIdDataGridColumnKey,
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

type ShippingAddressRequestCreatedDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: ShippingAddressRequestWithRefs | undefined) => string;
    renderCell?: GridColDef<ShippingAddressRequestWithRefs>['renderCell'];
};

export const ShippingAddressRequestCreatedDataGridColumnKey = 'created' as const;

export function useShippingAddressRequestCreatedDataGridColumn(options: ShippingAddressRequestCreatedDataGridColumnOptions) {

    return useMemo<GridColDef<ShippingAddressRequestWithRefs>>(() => ({
        headerName: options.headerName ?? ShippingAddressRequestCreatedDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: ShippingAddressRequestCreatedDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.shippingAddressRequest.created;
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

type ShippingAddressRequestDecisionReasonDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: ShippingAddressRequestWithRefs | undefined) => string;
    renderCell?: GridColDef<ShippingAddressRequestWithRefs>['renderCell'];
};

export const ShippingAddressRequestDecisionReasonDataGridColumnKey = 'decisionReason' as const;

export function useShippingAddressRequestDecisionReasonDataGridColumn(options: ShippingAddressRequestDecisionReasonDataGridColumnOptions) {

    return useMemo<GridColDef<ShippingAddressRequestWithRefs>>(() => ({
        headerName: options.headerName ?? ShippingAddressRequestDecisionReasonDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: ShippingAddressRequestDecisionReasonDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.shippingAddressRequest.decisionReason;
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

type ShippingAddressRequestEmailEncDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: ShippingAddressRequestWithRefs | undefined) => string;
    renderCell?: GridColDef<ShippingAddressRequestWithRefs>['renderCell'];
};

export const ShippingAddressRequestEmailEncDataGridColumnKey = 'emailEnc' as const;

export function useShippingAddressRequestEmailEncDataGridColumn(options: ShippingAddressRequestEmailEncDataGridColumnOptions) {

    return useMemo<GridColDef<ShippingAddressRequestWithRefs>>(() => ({
        headerName: options.headerName ?? ShippingAddressRequestEmailEncDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: ShippingAddressRequestEmailEncDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.shippingAddressRequest.emailEnc;
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

type ShippingAddressRequestEmailHashDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: ShippingAddressRequestWithRefs | undefined) => string;
    renderCell?: GridColDef<ShippingAddressRequestWithRefs>['renderCell'];
};

export const ShippingAddressRequestEmailHashDataGridColumnKey = 'emailHash' as const;

export function useShippingAddressRequestEmailHashDataGridColumn(options: ShippingAddressRequestEmailHashDataGridColumnOptions) {

    return useMemo<GridColDef<ShippingAddressRequestWithRefs>>(() => ({
        headerName: options.headerName ?? ShippingAddressRequestEmailHashDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: ShippingAddressRequestEmailHashDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.shippingAddressRequest.emailHash;
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

type ShippingAddressRequestNameDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: ShippingAddressRequestWithRefs | undefined) => string;
    renderCell?: GridColDef<ShippingAddressRequestWithRefs>['renderCell'];
};

export const ShippingAddressRequestNameDataGridColumnKey = 'name' as const;

export function useShippingAddressRequestNameDataGridColumn(options: ShippingAddressRequestNameDataGridColumnOptions) {

    return useMemo<GridColDef<ShippingAddressRequestWithRefs>>(() => ({
        headerName: options.headerName ?? ShippingAddressRequestNameDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: ShippingAddressRequestNameDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.shippingAddressRequest.name;
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

type ShippingAddressRequestNoteDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: ShippingAddressRequestWithRefs | undefined) => string;
    renderCell?: GridColDef<ShippingAddressRequestWithRefs>['renderCell'];
};

export const ShippingAddressRequestNoteDataGridColumnKey = 'note' as const;

export function useShippingAddressRequestNoteDataGridColumn(options: ShippingAddressRequestNoteDataGridColumnOptions) {

    return useMemo<GridColDef<ShippingAddressRequestWithRefs>>(() => ({
        headerName: options.headerName ?? ShippingAddressRequestNoteDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: ShippingAddressRequestNoteDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.shippingAddressRequest.note;
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

type ShippingAddressRequestOwnerIdDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue: (obj: ShippingAddressRequestWithRefs | undefined) => string;
    renderCell?: GridColDef<ShippingAddressRequestWithRefs>['renderCell']; 
    optionLabelField: keyof OwnerUser;
};

export const ShippingAddressRequestOwnerIdDataGridColumnKey = 'ownerId' as const;

export function useShippingAddressRequestOwnerIdDataGridColumn(options: ShippingAddressRequestOwnerIdDataGridColumnOptions) { 
    const ownerUsers = useSearchOwnerUsers({
        queryKey: ['ownerUser-data-grid-options'],
        query: {},
        projection: { id: true, [options.optionLabelField]: true },
    });

    return useMemo<GridColDef<ShippingAddressRequestWithRefs>>(() => ({
        headerName: options.headerName ?? ShippingAddressRequestOwnerIdDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: ShippingAddressRequestOwnerIdDataGridColumnKey,
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

type ShippingAddressRequestPolicyVersionDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: ShippingAddressRequestWithRefs | undefined) => number;
    renderCell?: GridColDef<ShippingAddressRequestWithRefs>['renderCell'];
};

export const ShippingAddressRequestPolicyVersionDataGridColumnKey = 'policyVersion' as const;

export function useShippingAddressRequestPolicyVersionDataGridColumn(options: ShippingAddressRequestPolicyVersionDataGridColumnOptions) {

    return useMemo<GridColDef<ShippingAddressRequestWithRefs>>(() => ({
        headerName: options.headerName ?? ShippingAddressRequestPolicyVersionDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: ShippingAddressRequestPolicyVersionDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.shippingAddressRequest.policyVersion;
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

type ShippingAddressRequestRegistryIdDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue: (obj: ShippingAddressRequestWithRefs | undefined) => string;
    renderCell?: GridColDef<ShippingAddressRequestWithRefs>['renderCell']; 
    optionLabelField: keyof Registry;
};

export const ShippingAddressRequestRegistryIdDataGridColumnKey = 'registryId' as const;

export function useShippingAddressRequestRegistryIdDataGridColumn(options: ShippingAddressRequestRegistryIdDataGridColumnOptions) { 
    const registrys = useSearchRegistrys({
        queryKey: ['registry-data-grid-options'],
        query: {},
        projection: { id: true, [options.optionLabelField]: true },
    });

    return useMemo<GridColDef<ShippingAddressRequestWithRefs>>(() => ({
        headerName: options.headerName ?? ShippingAddressRequestRegistryIdDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: ShippingAddressRequestRegistryIdDataGridColumnKey,
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

type ShippingAddressRequestRegistryItemIdDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue: (obj: ShippingAddressRequestWithRefs | undefined) => string;
    renderCell?: GridColDef<ShippingAddressRequestWithRefs>['renderCell']; 
    optionLabelField: keyof RegistryItem;
};

export const ShippingAddressRequestRegistryItemIdDataGridColumnKey = 'registryItemId' as const;

export function useShippingAddressRequestRegistryItemIdDataGridColumn(options: ShippingAddressRequestRegistryItemIdDataGridColumnOptions) { 
    const registryItems = useSearchRegistryItems({
        queryKey: ['registryItem-data-grid-options'],
        query: {},
        projection: { id: true, [options.optionLabelField]: true },
    });

    return useMemo<GridColDef<ShippingAddressRequestWithRefs>>(() => ({
        headerName: options.headerName ?? ShippingAddressRequestRegistryItemIdDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: ShippingAddressRequestRegistryItemIdDataGridColumnKey,
        valueGetter: (_, row) => {
            return options.getValue(row);
        }, 
        type: 'singleSelect',
        valueOptions: registryItems.data?.data?.map((obj) => ({
            value: obj.registryItem.id,
            label: obj.registryItem[options.optionLabelField],
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

type ShippingAddressRequestStatusDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: ShippingAddressRequestWithRefs | undefined) => string;
    renderCell?: GridColDef<ShippingAddressRequestWithRefs>['renderCell'];
    getOptionLabel: (opt: ShippingAddressRequestWithRefs['shippingAddressRequest']['status']) => string;
};

export const ShippingAddressRequestStatusDataGridColumnKey = 'status' as const;

export function useShippingAddressRequestStatusDataGridColumn(options: ShippingAddressRequestStatusDataGridColumnOptions) { 
    const values = useMemo(() => [
            {
                value: 'Pending',
                label: options.getOptionLabel('Pending'),
            },
            {
                value: 'Approved',
                label: options.getOptionLabel('Approved'),
            },
            {
                value: 'AutoApproved',
                label: options.getOptionLabel('AutoApproved'),
            },
            {
                value: 'Rejected',
                label: options.getOptionLabel('Rejected'),
            },
            {
                value: 'Blocked',
                label: options.getOptionLabel('Blocked'),
            },
    ], [options.getOptionLabel]);

    return useMemo<GridColDef<ShippingAddressRequestWithRefs>>(() => ({
        headerName: options.headerName ?? ShippingAddressRequestStatusDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: ShippingAddressRequestStatusDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.shippingAddressRequest.status;
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

type ShippingAddressRequestUpdatedDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: ShippingAddressRequestWithRefs | undefined) => string;
    renderCell?: GridColDef<ShippingAddressRequestWithRefs>['renderCell'];
};

export const ShippingAddressRequestUpdatedDataGridColumnKey = 'updated' as const;

export function useShippingAddressRequestUpdatedDataGridColumn(options: ShippingAddressRequestUpdatedDataGridColumnOptions) {

    return useMemo<GridColDef<ShippingAddressRequestWithRefs>>(() => ({
        headerName: options.headerName ?? ShippingAddressRequestUpdatedDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: ShippingAddressRequestUpdatedDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.shippingAddressRequest.updated;
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

type ShippingAddressRequestUpdatedByOwnerUserDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: ShippingAddressRequestWithRefs | undefined) => string;
    renderCell?: GridColDef<ShippingAddressRequestWithRefs>['renderCell'];
};

export const ShippingAddressRequestUpdatedByOwnerUserDataGridColumnKey = 'updatedByOwnerUser' as const;

export function useShippingAddressRequestUpdatedByOwnerUserDataGridColumn(options: ShippingAddressRequestUpdatedByOwnerUserDataGridColumnOptions) {

    return useMemo<GridColDef<ShippingAddressRequestWithRefs>>(() => ({
        headerName: options.headerName ?? ShippingAddressRequestUpdatedByOwnerUserDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: ShippingAddressRequestUpdatedByOwnerUserDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.shippingAddressRequest.updatedByOwnerUser;
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

export function getShippingAddressRequestColumnVisibilityModel(
    defaultValue: boolean,
    projection: ShippingAddressRequestProjection,
): GridColumnVisibilityModel {
    return {
        id: projection.id ?? false,
        created: projection.created ?? false,
        decisionReason: projection.decisionReason ?? false,
        emailEnc: projection.emailEnc ?? false,
        emailHash: projection.emailHash ?? false,
        name: projection.name ?? false,
        note: projection.note ?? false,
        ownerId: projection.ownerId ?? false,
        policyVersion: projection.policyVersion ?? false,
        registryId: projection.registryId ?? false,
        registryItemId: projection.registryItemId ?? false,
        status: projection.status ?? false,
        updated: projection.updated ?? false,
        updatedByOwnerUser: projection.updatedByOwnerUser ?? false,
    };
}

// URL query string helpers
function parseUrlSortParams(sortParam: string): ShippingAddressRequestSortParams | null {
    try {
        return JSON.parse(sortParam);
    } catch (error) {
        return null;
    }
}

function serializeSortParams(sort: ShippingAddressRequestSortParams): string {
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