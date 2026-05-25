// This file is auto-generated. DO NOT EDIT.

import { useCallback, useMemo, useState } from 'react';
import { GridFilterModel, GridSortModel, GridColDef, GridColumnVisibilityModel } from '@mui/x-data-grid';
import { AddressAccessSessionProjection, AddressAccessSessionSortParams } from '../../model/address-access-session-model';
import { AddressAccessSessionWithRefs } from '../../model/address-access-session-api';
import { OwnerUser } from '../../model/owner-user-model';
import { useSearchOwnerUsers } from '../tanstack-query/owner-user-queries';
import { Registry } from '../../model/registry-model';
import { useSearchRegistrys } from '../tanstack-query/registry-queries';
import { RegistryApprovedGuest } from '../../model/registry-approved-guest-model';
import { useSearchRegistryApprovedGuests } from '../tanstack-query/registry-approved-guest-queries';

type Options = {
    defaultSort?: AddressAccessSessionSortParams,
    enableUrlQueryStrings?: boolean,
}

export const AddressAccessSessionMUIDataGridUrlQueryKeys = {
    sort: 'addressAccessSessionSort' as const,
    filter: 'addressAccessSessionFilter' as const,
    page: 'addressAccessSessionPage' as const,
    pageSize: 'addressAccessSessionPageSize' as const,
} as const;

export function useAddressAccessSessionMuiDataGridSortModel(options: Options = {}) {
    const [sort, setSort] = useState<AddressAccessSessionSortParams>(() => {
        if (options.enableUrlQueryStrings) {
            const urlParams = new URLSearchParams(window.location.search);
            const param = urlParams.get(AddressAccessSessionMUIDataGridUrlQueryKeys.sort);
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
          next.push({ field: key, sort: sort[key as keyof AddressAccessSessionSortParams] === 1 ? 'asc' : 'desc' });
        }
        return next as GridSortModel;
    }, [sort]);

    const setGridSortModel = useCallback((model: GridSortModel) => {
        const next: AddressAccessSessionSortParams = {};
        for (const item of model) {
          next[item.field as keyof AddressAccessSessionSortParams] = item.sort === 'asc' ? 1 : -1;
        }
        setSort(next);
        if (options.enableUrlQueryStrings) {
            const urlParams = new URLSearchParams(window.location.search);
            urlParams.set(AddressAccessSessionMUIDataGridUrlQueryKeys.sort, serializeSortParams(next));
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

export function useAddressAccessSessionMuiDataGridFilterModel(options: FilterOptions = {}) {
    const [gridFilterModel, setGridFilterModel] = useState<GridFilterModel | undefined>(() => {
        if (!options.enableUrlQueryStrings) {
            return undefined;
        }
        const urlParams = new URLSearchParams(window.location.search);
        const filterParam = urlParams.get(AddressAccessSessionMUIDataGridUrlQueryKeys.filter);
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
            if (!urlParams.has(AddressAccessSessionMUIDataGridUrlQueryKeys.filter)) {
                return;
            }
            urlParams.delete(AddressAccessSessionMUIDataGridUrlQueryKeys.filter);
            window.history.replaceState({}, '', `${window.location.pathname}?${urlParams.toString()}`);
            return;
        }
        const urlParams = new URLSearchParams(window.location.search);
        if (model) {
            urlParams.set(AddressAccessSessionMUIDataGridUrlQueryKeys.filter, serializeFilterModel(model));
        } else {
            urlParams.delete(AddressAccessSessionMUIDataGridUrlQueryKeys.filter);
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
            case "approvedGuestId":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                approvedGuestIdEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                approvedGuestIdEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                approvedGuestIdEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                approvedGuestIdEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                approvedGuestIdIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                approvedGuestIdIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                approvedGuestIdNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                approvedGuestIdNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                approvedGuestIdExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                approvedGuestIdExists: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                }
            case "created":
                switch(gridFilterModel?.items?.[0]?.operator) { 
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
            case "expiresAt":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                expiresAtEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                expiresAtEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                expiresAtEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                expiresAtEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                expiresAtNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                expiresAtNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                expiresAtGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                expiresAtGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                expiresAtGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                expiresAtGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                expiresAtGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                expiresAtGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                expiresAtGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                expiresAtGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                expiresAtLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                expiresAtLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                expiresAtLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                expiresAtLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                expiresAtLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                expiresAtLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                expiresAtLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                expiresAtLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                expiresAtIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                expiresAtIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                expiresAtNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                expiresAtNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                expiresAtExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                expiresAtExists: gridFilterModel?.items?.[0]?.value,
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
            case "policyVersionAtIssue":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                policyVersionAtIssueEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionAtIssueEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                policyVersionAtIssueEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionAtIssueEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                policyVersionAtIssueNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionAtIssueNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                policyVersionAtIssueGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionAtIssueGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                policyVersionAtIssueGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionAtIssueGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                policyVersionAtIssueGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionAtIssueGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                policyVersionAtIssueGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionAtIssueGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                policyVersionAtIssueLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionAtIssueLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                policyVersionAtIssueLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionAtIssueLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                policyVersionAtIssueLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionAtIssueLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                policyVersionAtIssueLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionAtIssueLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                policyVersionAtIssueIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                policyVersionAtIssueIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                policyVersionAtIssueNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                policyVersionAtIssueNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                policyVersionAtIssueExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                policyVersionAtIssueExists: gridFilterModel?.items?.[0]?.value,
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
            case "tokenHash":
                switch(gridFilterModel?.items?.[0]?.operator) { 
                    case "equals":
                        return { 
                            searchQuery: { 
                                tokenHashEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                tokenHashEq: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "is":
                        return { 
                            searchQuery: { 
                                tokenHashEq: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                tokenHashEq: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notEquals":
                        return { 
                            searchQuery: { 
                                tokenHashNe: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                tokenHashNe: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThan":
                        return { 
                            searchQuery: { 
                                tokenHashGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                tokenHashGt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "after":
                        return { 
                            searchQuery: { 
                                tokenHashGt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                tokenHashGt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "greaterThanOrEqual":
                        return { 
                            searchQuery: { 
                                tokenHashGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                tokenHashGte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrAfter":
                        return { 
                            searchQuery: { 
                                tokenHashGte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                tokenHashGte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThan":
                        return { 
                            searchQuery: { 
                                tokenHashLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                tokenHashLt: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "before":
                        return { 
                            searchQuery: { 
                                tokenHashLt: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                tokenHashLt: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "lessThanOrEqual":
                        return { 
                            searchQuery: { 
                                tokenHashLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                tokenHashLte: gridFilterModel?.items?.[0]?.value,
                            },
                        };
                    case "onOrBefore":
                        return { 
                            searchQuery: { 
                                tokenHashLte: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                tokenHashLte: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "in":
                        return { 
                            searchQuery: { 
                                tokenHashIn: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                tokenHashIn: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notIn":
                        return { 
                            searchQuery: { 
                                tokenHashNin: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value?.length ? undefined :  { 
                                tokenHashNin: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "isEmpty":
                        return { 
                            searchQuery: { 
                                tokenHashExists: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: gridFilterModel?.items?.[0]?.value == null ? undefined :  { 
                                tokenHashExists: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "contains":
                        return { 
                            searchQuery: { 
                                tokenHashLike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                tokenHashLike: gridFilterModel?.items?.[0]?.value,
                            },
                        }; 
                    case "notContains":
                        return { 
                            searchQuery: { 
                                tokenHashNlike: gridFilterModel?.items?.[0]?.value,
                            },
                            searchQueryExcludingIncompleteFilters: !gridFilterModel?.items?.[0]?.value ? undefined :  { 
                                tokenHashNlike: gridFilterModel?.items?.[0]?.value,
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

type AddressAccessSessionIdDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue: (obj: AddressAccessSessionWithRefs | undefined) => string;
    renderCell?: GridColDef<AddressAccessSessionWithRefs>['renderCell']; 
};

export const AddressAccessSessionIdDataGridColumnKey = 'id' as const;

export function useAddressAccessSessionIdDataGridColumn(options: AddressAccessSessionIdDataGridColumnOptions) { 

    return useMemo<GridColDef<AddressAccessSessionWithRefs>>(() => ({
        headerName: options.headerName ?? AddressAccessSessionIdDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: AddressAccessSessionIdDataGridColumnKey,
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

type AddressAccessSessionApprovedGuestIdDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue: (obj: AddressAccessSessionWithRefs | undefined) => string;
    renderCell?: GridColDef<AddressAccessSessionWithRefs>['renderCell']; 
    optionLabelField: keyof RegistryApprovedGuest;
};

export const AddressAccessSessionApprovedGuestIdDataGridColumnKey = 'approvedGuestId' as const;

export function useAddressAccessSessionApprovedGuestIdDataGridColumn(options: AddressAccessSessionApprovedGuestIdDataGridColumnOptions) { 
    const registryApprovedGuests = useSearchRegistryApprovedGuests({
        queryKey: ['registryApprovedGuest-data-grid-options'],
        query: {},
        projection: { id: true, [options.optionLabelField]: true },
    });

    return useMemo<GridColDef<AddressAccessSessionWithRefs>>(() => ({
        headerName: options.headerName ?? AddressAccessSessionApprovedGuestIdDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: AddressAccessSessionApprovedGuestIdDataGridColumnKey,
        valueGetter: (_, row) => {
            return options.getValue(row);
        }, 
        type: 'singleSelect',
        valueOptions: registryApprovedGuests.data?.data?.map((obj) => ({
            value: obj.registryApprovedGuest.id,
            label: obj.registryApprovedGuest[options.optionLabelField],
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

type AddressAccessSessionCreatedDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: AddressAccessSessionWithRefs | undefined) => string;
    renderCell?: GridColDef<AddressAccessSessionWithRefs>['renderCell'];
};

export const AddressAccessSessionCreatedDataGridColumnKey = 'created' as const;

export function useAddressAccessSessionCreatedDataGridColumn(options: AddressAccessSessionCreatedDataGridColumnOptions) {

    return useMemo<GridColDef<AddressAccessSessionWithRefs>>(() => ({
        headerName: options.headerName ?? AddressAccessSessionCreatedDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: AddressAccessSessionCreatedDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.addressAccessSession.created;
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

type AddressAccessSessionEmailHashDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: AddressAccessSessionWithRefs | undefined) => string;
    renderCell?: GridColDef<AddressAccessSessionWithRefs>['renderCell'];
};

export const AddressAccessSessionEmailHashDataGridColumnKey = 'emailHash' as const;

export function useAddressAccessSessionEmailHashDataGridColumn(options: AddressAccessSessionEmailHashDataGridColumnOptions) {

    return useMemo<GridColDef<AddressAccessSessionWithRefs>>(() => ({
        headerName: options.headerName ?? AddressAccessSessionEmailHashDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: AddressAccessSessionEmailHashDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.addressAccessSession.emailHash;
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

type AddressAccessSessionExpiresAtDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: AddressAccessSessionWithRefs | undefined) => Date;
    renderCell?: GridColDef<AddressAccessSessionWithRefs>['renderCell'];
};

export const AddressAccessSessionExpiresAtDataGridColumnKey = 'expiresAt' as const;

export function useAddressAccessSessionExpiresAtDataGridColumn(options: AddressAccessSessionExpiresAtDataGridColumnOptions) {

    return useMemo<GridColDef<AddressAccessSessionWithRefs>>(() => ({
        headerName: options.headerName ?? AddressAccessSessionExpiresAtDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: AddressAccessSessionExpiresAtDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : new Date(row.addressAccessSession.expiresAt ?? 0);
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

type AddressAccessSessionOwnerIdDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue: (obj: AddressAccessSessionWithRefs | undefined) => string;
    renderCell?: GridColDef<AddressAccessSessionWithRefs>['renderCell']; 
    optionLabelField: keyof OwnerUser;
};

export const AddressAccessSessionOwnerIdDataGridColumnKey = 'ownerId' as const;

export function useAddressAccessSessionOwnerIdDataGridColumn(options: AddressAccessSessionOwnerIdDataGridColumnOptions) { 
    const ownerUsers = useSearchOwnerUsers({
        queryKey: ['ownerUser-data-grid-options'],
        query: {},
        projection: { id: true, [options.optionLabelField]: true },
    });

    return useMemo<GridColDef<AddressAccessSessionWithRefs>>(() => ({
        headerName: options.headerName ?? AddressAccessSessionOwnerIdDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: AddressAccessSessionOwnerIdDataGridColumnKey,
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

type AddressAccessSessionPolicyVersionAtIssueDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: AddressAccessSessionWithRefs | undefined) => number;
    renderCell?: GridColDef<AddressAccessSessionWithRefs>['renderCell'];
};

export const AddressAccessSessionPolicyVersionAtIssueDataGridColumnKey = 'policyVersionAtIssue' as const;

export function useAddressAccessSessionPolicyVersionAtIssueDataGridColumn(options: AddressAccessSessionPolicyVersionAtIssueDataGridColumnOptions) {

    return useMemo<GridColDef<AddressAccessSessionWithRefs>>(() => ({
        headerName: options.headerName ?? AddressAccessSessionPolicyVersionAtIssueDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: AddressAccessSessionPolicyVersionAtIssueDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.addressAccessSession.policyVersionAtIssue;
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

type AddressAccessSessionRegistryIdDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue: (obj: AddressAccessSessionWithRefs | undefined) => string;
    renderCell?: GridColDef<AddressAccessSessionWithRefs>['renderCell']; 
    optionLabelField: keyof Registry;
};

export const AddressAccessSessionRegistryIdDataGridColumnKey = 'registryId' as const;

export function useAddressAccessSessionRegistryIdDataGridColumn(options: AddressAccessSessionRegistryIdDataGridColumnOptions) { 
    const registrys = useSearchRegistrys({
        queryKey: ['registry-data-grid-options'],
        query: {},
        projection: { id: true, [options.optionLabelField]: true },
    });

    return useMemo<GridColDef<AddressAccessSessionWithRefs>>(() => ({
        headerName: options.headerName ?? AddressAccessSessionRegistryIdDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: AddressAccessSessionRegistryIdDataGridColumnKey,
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

type AddressAccessSessionTokenHashDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: AddressAccessSessionWithRefs | undefined) => string;
    renderCell?: GridColDef<AddressAccessSessionWithRefs>['renderCell'];
};

export const AddressAccessSessionTokenHashDataGridColumnKey = 'tokenHash' as const;

export function useAddressAccessSessionTokenHashDataGridColumn(options: AddressAccessSessionTokenHashDataGridColumnOptions) {

    return useMemo<GridColDef<AddressAccessSessionWithRefs>>(() => ({
        headerName: options.headerName ?? AddressAccessSessionTokenHashDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: AddressAccessSessionTokenHashDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.addressAccessSession.tokenHash;
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

type AddressAccessSessionUpdatedDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: AddressAccessSessionWithRefs | undefined) => string;
    renderCell?: GridColDef<AddressAccessSessionWithRefs>['renderCell'];
};

export const AddressAccessSessionUpdatedDataGridColumnKey = 'updated' as const;

export function useAddressAccessSessionUpdatedDataGridColumn(options: AddressAccessSessionUpdatedDataGridColumnOptions) {

    return useMemo<GridColDef<AddressAccessSessionWithRefs>>(() => ({
        headerName: options.headerName ?? AddressAccessSessionUpdatedDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: AddressAccessSessionUpdatedDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.addressAccessSession.updated;
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

type AddressAccessSessionUpdatedByOwnerUserDataGridColumnOptions = {
    headerName?: string;
    width?: number;
    sortable?: boolean;
    hideable?: boolean;
    getValue?: (obj: AddressAccessSessionWithRefs | undefined) => string;
    renderCell?: GridColDef<AddressAccessSessionWithRefs>['renderCell'];
};

export const AddressAccessSessionUpdatedByOwnerUserDataGridColumnKey = 'updatedByOwnerUser' as const;

export function useAddressAccessSessionUpdatedByOwnerUserDataGridColumn(options: AddressAccessSessionUpdatedByOwnerUserDataGridColumnOptions) {

    return useMemo<GridColDef<AddressAccessSessionWithRefs>>(() => ({
        headerName: options.headerName ?? AddressAccessSessionUpdatedByOwnerUserDataGridColumnKey,
        width: options.width,
        sortable: options.sortable,
        hideable: options.hideable,
        field: AddressAccessSessionUpdatedByOwnerUserDataGridColumnKey,
        valueGetter: (_, row) => { 
            return options.getValue ? options.getValue(row) : row.addressAccessSession.updatedByOwnerUser;
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

export function getAddressAccessSessionColumnVisibilityModel(
    defaultValue: boolean,
    projection: AddressAccessSessionProjection,
): GridColumnVisibilityModel {
    return {
        id: projection.id ?? false,
        approvedGuestId: projection.approvedGuestId ?? false,
        created: projection.created ?? false,
        emailHash: projection.emailHash ?? false,
        expiresAt: projection.expiresAt ?? false,
        ownerId: projection.ownerId ?? false,
        policyVersionAtIssue: projection.policyVersionAtIssue ?? false,
        registryId: projection.registryId ?? false,
        tokenHash: projection.tokenHash ?? false,
        updated: projection.updated ?? false,
        updatedByOwnerUser: projection.updatedByOwnerUser ?? false,
    };
}

// URL query string helpers
function parseUrlSortParams(sortParam: string): AddressAccessSessionSortParams | null {
    try {
        return JSON.parse(sortParam);
    } catch (error) {
        return null;
    }
}

function serializeSortParams(sort: AddressAccessSessionSortParams): string {
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