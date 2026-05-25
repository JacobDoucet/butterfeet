// This file is auto-generated. DO NOT EDIT.

import { ActorTrace, ActorTraceProjection } from './actor-trace-model';

export type Registry = {
  id?: string;
  coverImageUrl?: string;
  created?: ActorTrace;
  dueDate?: string;
  isPublic?: boolean;
  ownerId?: string;
  parentNames?: string;
  slug?: string;
  themeColor?: string;
  title?: string;
  updated?: ActorTrace;
  updatedByOwnerUser?: ActorTrace;
  welcomeMessage?: string;
}

export type RegistryProjection = {
    id?: boolean;
    coverImageUrl?: boolean;
    created?: boolean;
		createdFields?: ActorTraceProjection;
    dueDate?: boolean;
    isPublic?: boolean;
    ownerId?: boolean;
    parentNames?: boolean;
    slug?: boolean;
    themeColor?: boolean;
    title?: boolean;
    updated?: boolean;
		updatedFields?: ActorTraceProjection;
    updatedByOwnerUser?: boolean;
		updatedByOwnerUserFields?: ActorTraceProjection;
    welcomeMessage?: boolean;
}

export type RegistrySortParams = {
    createdAt?: -1 | 1;
    ownerId?: -1 | 1;
    slug?: -1 | 1;
    updatedAt?: -1 | 1;
}
