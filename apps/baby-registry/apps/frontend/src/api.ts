export const API_BASE = (import.meta.env.VITE_API_BASE_URL as string | undefined) ?? '';

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const res = await fetch(`${API_BASE}${path}`, {
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',
      ...(init?.headers ?? {}),
    },
    ...init,
  });
  if (!res.ok) {
    const text = await res.text().catch(() => '');
    throw new Error(text || `${res.status} ${res.statusText}`);
  }
  if (res.status === 204) return undefined as T;
  const ct = res.headers.get('content-type') ?? '';
  if (ct.includes('application/json')) return res.json() as Promise<T>;
  return undefined as T;
}

export const api = {
  get: <T>(path: string) => request<T>(path),
  post: <T>(path: string, body?: unknown) =>
    request<T>(path, { method: 'POST', body: body ? JSON.stringify(body) : undefined }),
  patch: <T>(path: string, body?: unknown) =>
    request<T>(path, { method: 'PATCH', body: body ? JSON.stringify(body) : undefined }),
  del: <T>(path: string) => request<T>(path, { method: 'DELETE' }),
};

// Domain helpers
export interface Me { id: string; email: string; name: string; }

export type AddressAccessMode = 'RequestApproval' | 'ApprovedGuestsOnly' | 'Disabled';

export interface Registry {
  id: string;
  slug: string;
  title: string;
  parentNames?: string;
  welcomeMessage?: string;
  themeColor?: string;
  coverImageUrl?: string;
  dueDate?: string;
  isPublic?: boolean;
  ownerId?: string;
  addressAccessMode?: AddressAccessMode;
  shippingPolicyVersion?: number;
  shippingRecipientName?: string;
  shippingLine1?: string;
  shippingLine2?: string;
  shippingCity?: string;
  shippingRegion?: string;
  shippingPostalCode?: string;
  shippingCountry?: string;
  shippingDeliveryNotes?: string;
}

export interface RegistryItem {
  id: string;
  registryId: string;
  title: string;
  description?: string;
  imageUrl?: string;
  productUrl?: string;
  source?: string;
  priceCents?: number;
  currency?: string;
  quantity?: number;
  quantityUnlimited?: boolean;
  category?: string;
  noSubstitutes?: boolean;
  parentItemId?: string;
  ownerPurchased?: boolean;
  notes?: string;
  position?: number;
}

export interface PublicRegistry extends Registry {
  items: (RegistryItem & { reserved: number })[];
}

export interface ScrapeResult {
  title: string;
  imageUrl: string;
  productUrl: string;
  price: number;
  currency: string;
  source: string;
}

export const auth = {
  me: () => api.get<Me>('/api/auth/me'),
  request: (email: string, name?: string) =>
    api.post<{ ok: boolean }>('/api/auth/magic/request', { email, name }),
  verify: (token: string) =>
    api.post<{ ok: boolean; ownerId: string }>('/api/auth/magic/verify', { token }),
  logout: () => api.post<void>('/api/auth/logout'),
};

interface QueryResult<T> { data: T[]; total: number; skip: number; metadata?: unknown }
interface MutationResult<T> { data: T; metadata?: unknown }

function unwrapList<T>(key: string) {
  return (r: { data: Array<Record<string, T>>; total: number; skip: number; metadata?: unknown }): QueryResult<T> => ({
    ...r,
    data: r.data.map((row) => row[key]),
  });
}

export const registries = {
  list: () =>
    api.post<{ data: Array<{ registry: Registry }>; total: number; skip: number; metadata?: unknown }>(
      '/api/registries/search',
      {},
    ).then(unwrapList<Registry>('registry')),
  create: (body: Partial<Registry>) =>
    api.post<MutationResult<Registry>>('/api/registries/create', { data: body }).then((r) => r.data),
  update: (id: string, body: Partial<Registry>) =>
    api.patch<MutationResult<Registry>>('/api/registries/update', { data: { id, ...body } }).then((r) => r.data),
  remove: (id: string) => api.del<void>(`/api/registries/delete/${id}`),
};

export const items = {
  listForRegistry: (registryId: string) =>
    api.post<{ data: Array<{ registryItem: RegistryItem }>; total: number; skip: number; metadata?: unknown }>(
      '/api/registry-items/search',
      {
        query: { registryIdEq: registryId },
        sort: { position: 1 },
        limit: 500,
      },
    ).then(unwrapList<RegistryItem>('registryItem')),
  create: (body: Partial<RegistryItem>) =>
    api.post<MutationResult<RegistryItem>>('/api/registry-items/create', { data: body }).then((r) => r.data),
  update: (id: string, body: Partial<RegistryItem>) =>
    api.patch<MutationResult<RegistryItem>>('/api/registry-items/update', { data: { id, ...body } }).then((r) => r.data),
  remove: (id: string) => api.del<void>(`/api/registry-items/delete/${id}`),
};

export type ReservationStatus = 'Reserved' | 'Purchased' | 'Received' | 'Cancelled';

export interface Reservation {
  id: string;
  itemId: string;
  registryId: string;
  reserverName?: string;
  isAnonymous?: boolean;
  message?: string;
  contactEmail?: string;
  quantity?: number;
  status: ReservationStatus;
  created?: string;
}

export const reservations = {
  listForRegistry: (registryId: string) =>
    api.post<{ data: Array<{ reservation: Reservation }>; total: number; skip: number; metadata?: unknown }>(
      '/api/reservations/search',
      {
        query: { registryIdEq: registryId },
        sort: { created: -1 },
        limit: 1000,
      },
    ).then(unwrapList<Reservation>('reservation')),
  create: (body: Partial<Reservation>) =>
    api.post<MutationResult<Reservation>>('/api/reservations/create', { data: body }).then((r) => r.data),
  setStatus: (id: string, status: ReservationStatus) =>
    api.patch<MutationResult<Reservation>>('/api/reservations/update', { data: { id, status } }).then((r) => r.data),
  remove: (id: string) => api.del<void>(`/api/reservations/delete/${id}`),
};

export const pub = {
  registry: (slug: string) => api.get<PublicRegistry>(`/api/public/r/${encodeURIComponent(slug)}`),
  reserve: (itemId: string, body: { reserverName: string; isAnonymous: boolean; message: string; contactEmail?: string; quantity?: number }) =>
    api.post<{ ok: boolean }>(`/api/public/items/${itemId}/reserve`, body),
};

export interface BuyerSession {
  email: string;
}

export const buyer = {
  request: (slug: string, email: string) =>
    api.post<{ ok: boolean }>('/api/public/buyer/verify/request', { slug, email }),
  confirm: (slug: string, email: string, code: string) =>
    api.post<{ ok: boolean; email: string }>('/api/public/buyer/verify/confirm', { slug, email, code }),
  me: (slug: string) =>
    api.get<BuyerSession>(`/api/public/buyer/me?slug=${encodeURIComponent(slug)}`),
  logout: (slug: string) =>
    api.post<void>(`/api/public/buyer/logout?slug=${encodeURIComponent(slug)}`),
};

export const scrape = {
  url: (u: string) => api.get<ScrapeResult>(`/api/scrape?url=${encodeURIComponent(u)}`),
};

export type GuestAccessLevel = 'ViewShippingAddress' | 'ReserveOnly';
export type GuestStatus = 'Active' | 'Revoked' | 'Blocked';

export interface ApprovedGuest {
  id: string;
  registryId: string;
  email: string;
  name?: string;
  accessLevel: GuestAccessLevel;
  status: GuestStatus;
}

export const approvedGuests = {
  list: (registryId: string) =>
    api.get<{ data: ApprovedGuest[] }>(`/api/shipping/registries/${registryId}/approved-guests`).then((r) => r.data),
  add: (registryId: string, body: { email: string; name?: string; accessLevel?: GuestAccessLevel }) =>
    api.post<ApprovedGuest>(`/api/shipping/registries/${registryId}/approved-guests`, body),
  revoke: (id: string) => api.post<ApprovedGuest>(`/api/shipping/approved-guests/${id}/revoke`),
  block: (id: string) => api.post<ApprovedGuest>(`/api/shipping/approved-guests/${id}/block`),
  reactivate: (id: string) => api.post<ApprovedGuest>(`/api/shipping/approved-guests/${id}/reactivate`),
  remove: (id: string) => api.del<{ ok: boolean }>(`/api/shipping/approved-guests/${id}`),
  issueLink: (id: string) =>
    api.post<{ token: string; expiresAt: string }>(`/api/shipping/approved-guests/${id}/issue-link`),
};

export type AddressRequestStatus = 'Pending' | 'Approved' | 'AutoApproved' | 'Rejected' | 'Blocked';

export interface AddressRequest {
  id: string;
  registryId: string;
  registryItemId?: string;
  email: string;
  name?: string;
  note?: string;
  status: AddressRequestStatus;
  decisionReason?: string;
  policyVersion?: number;
  createdAt?: string;
  token?: string;
  tokenExpiresAt?: string;
}

export const addressRequests = {
  list: (registryId: string) =>
    api.get<{ data: AddressRequest[] }>(`/api/shipping/registries/${registryId}/requests`).then((r) => r.data),
  approve: (id: string, opts: { permanent?: boolean; reason?: string } = {}) =>
    api.post<AddressRequest>(`/api/shipping/requests/${id}/approve`, opts),
  reject: (id: string, reason?: string) =>
    api.post<AddressRequest>(`/api/shipping/requests/${id}/reject`, { reason }),
  block: (id: string, reason?: string) =>
    api.post<AddressRequest>(`/api/shipping/requests/${id}/block`, { reason }),
};

export interface ResolvedShippingAddress {
  registryTitle: string;
  recipientName: string;
  line1: string;
  line2?: string;
  city: string;
  region: string;
  postalCode: string;
  country: string;
  deliveryNotes?: string;
  expiresAt: string;
}

export const shippingShare = {
  resolve: (token: string) =>
    api.post<ResolvedShippingAddress>('/api/public/shipping/resolve', { token }),
};
