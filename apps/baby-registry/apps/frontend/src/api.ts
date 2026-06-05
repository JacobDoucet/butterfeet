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
  allowOpenAccess?: boolean;
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
  imageBgColor?: string;
  productUrl?: string;
  affiliateUrl?: string;
  retailer?: string;
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
  myReservations?: MyReservation[];
  myCarts?: MyCart[];
}

export interface MyReservation {
  id: string;
  itemId: string;
  itemTitle: string;
  quantity: number;
  createdAt: string;
  expiresAt: string;
}

export interface CartItem {
  reservationId: string;
  itemId: string;
  title: string;
  quantity: number;
  priceCents: number;
  currency: string;
}

export interface MyCart {
  id: string;
  referenceCode: string;
  status: CartStatus;
  amountCents: number;
  currency: string;
  methodDisplayName: string;
  createdAt: string;
  items: CartItem[];
}

export type RegistryAccessRequestStatus = 'none' | 'pending' | 'rejected';

export interface GatedRegistry {
  accessGated: true;
  slug: string;
  title: string;
  parentNames?: string;
  themeColor?: string;
  coverImageUrl?: string;
  welcomeMessage?: string;
  ownerDisplayName?: string;
  accessRequestStatus: RegistryAccessRequestStatus;
}

export type PublicRegistryResponse = PublicRegistry | GatedRegistry;

export function isGatedRegistry(r: PublicRegistryResponse): r is GatedRegistry {
  return (r as GatedRegistry).accessGated === true;
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
  renameCategory: (registryId: string, oldCategory: string, newCategory: string) =>
    api.post<{ ok: boolean; modifiedCount: number }>('/api/registry-admin/rename-category', {
      registryId,
      oldCategory,
      newCategory,
    }),
};

export type ReservationStatus = 'Reserved' | 'AwaitingConfirmation' | 'Purchased' | 'Received' | 'Cancelled';

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
  created?: { actorId?: string; actorName?: string; actorType?: string; at?: string };
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
  registry: (slug: string) => api.get<PublicRegistryResponse>(`/api/public/r/${encodeURIComponent(slug)}`),
  reserve: (itemId: string, body: { reserverName: string; isAnonymous: boolean; message: string; contactEmail?: string; quantity?: number }) =>
    api.post<{ ok: boolean; id: string }>(`/api/public/items/${itemId}/reserve`, body),
  confirmReservation: (id: string) =>
    api.post<{ ok: boolean; status: ReservationStatus }>(`/api/public/reservations/${encodeURIComponent(id)}/confirm`, {}),
  cancelReservation: (id: string) =>
    api.post<{ ok: boolean; status: ReservationStatus }>(`/api/public/reservations/${encodeURIComponent(id)}/cancel`, {}),
  requestAddress: (body: { slug: string; itemId?: string; name?: string; note?: string }) =>
    api.post<{ ok: boolean; status?: 'pending'; id?: string }>('/api/public/address-requests', body),
  requestRegistryAccess: (body: { slug: string; name?: string; note?: string }) =>
    api.post<{ ok: boolean; status: 'pending' | 'approved' | 'rejected' }>('/api/public/registry-access/request', body),
};

export interface ExchangeRates {
  base: string;
  rates: Record<string, number>;
  date: string;
  fetchedAt: string;
}

export const exchangeRates = {
  get: () => api.get<ExchangeRates>('/api/public/exchange-rates'),
};

// Currencies the public page lets a viewer convert prices into.
export const SUPPORTED_CURRENCIES: { code: string; label: string }[] = [
  { code: 'USD', label: 'USD · US Dollar' },
  { code: 'CAD', label: 'CAD · Canadian Dollar' },
  { code: 'GBP', label: 'GBP · British Pound' },
  { code: 'EUR', label: 'EUR · Euro' },
];

// convertCents converts an amount in minor units from one currency to another
// using a rates table keyed against `rates.base`. Returns null when either
// currency is missing from the table so callers can fall back to the original.
export function convertCents(
  cents: number,
  from: string,
  to: string,
  rates?: ExchangeRates,
): number | null {
  const fromCur = (from || 'USD').toUpperCase();
  const toCur = (to || 'USD').toUpperCase();
  if (fromCur === toCur) return cents;
  if (!rates) return null;
  const fromRate = rates.rates[fromCur];
  const toRate = rates.rates[toCur];
  if (!fromRate || !toRate) return null;
  // Convert via the base currency: amount_base = amount_from / fromRate.
  return Math.round((cents / fromRate) * toRate);
}

export interface BuyerSession {
  email: string;
  name?: string;
}

export const buyer = {
  request: (slug: string, email: string, name?: string) =>
    api.post<{ ok: boolean }>('/api/public/buyer/verify/request', { slug, email, name }),
  confirm: (slug: string, email: string, code: string) =>
    api.post<{ ok: boolean; email: string; name?: string }>('/api/public/buyer/verify/confirm', { slug, email, code }),
  me: (slug: string) =>
    api.get<BuyerSession>(`/api/public/buyer/me?slug=${encodeURIComponent(slug)}`),
  logout: (slug: string) =>
    api.post<void>(`/api/public/buyer/logout?slug=${encodeURIComponent(slug)}`),
};

export const scrape = {
  url: (u: string) => api.get<ScrapeResult>(`/api/scrape?url=${encodeURIComponent(u)}`),
};

// formatPriceCents renders a price stored in minor units (cents) using the
// shopper's locale. Returns null when no price is available so callers can
// gracefully omit the amount.
export function formatPriceCents(cents?: number, currency?: string): string | null {
  if (cents == null || Number.isNaN(cents)) return null;
  try {
    return new Intl.NumberFormat(undefined, {
      style: 'currency',
      currency: (currency || 'USD').toUpperCase(),
    }).format(cents / 100);
  } catch {
    return `$${(cents / 100).toFixed(2)}`;
  }
}

export type GuestAccessLevel = 'ViewShippingAddress' | 'ReserveOnly';
export type GuestStatus = 'Pending' | 'Active' | 'Revoked' | 'Blocked';

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

// ---------------------------------------------------------------------------
// Manual (parent-to-parent) payments
// ---------------------------------------------------------------------------

export type PaymentMethodType =
  | 'PayPal'
  | 'Revolut'
  | 'Wise'
  | 'InteracETransfer'
  | 'BankTransfer'
  | 'Other';

export const PAYMENT_METHOD_LABELS: Record<PaymentMethodType, string> = {
  PayPal: 'PayPal',
  Revolut: 'Revolut',
  Wise: 'Wise',
  InteracETransfer: 'Interac e-Transfer',
  BankTransfer: 'Bank transfer',
  Other: 'Other',
};

export interface PaymentMethod {
  id: string;
  registryId: string;
  type: PaymentMethodType;
  displayName?: string;
  instructions?: string;
  paymentUrl?: string;
  recipientEmail?: string;
  recipientPhone?: string;
  bankName?: string;
  bankAccountName?: string;
  bankAccountNumber?: string;
  bankRoutingNumber?: string;
  bankIban?: string;
  bankSwift?: string;
  enabled?: boolean;
  position?: number;
}

export type CartStatus = 'Pending' | 'AwaitingConfirmation' | 'Completed' | 'Rejected';

export interface Cart {
  id: string;
  registryId: string;
  paymentMethodId?: string;
  methodType: PaymentMethodType;
  methodDisplayName?: string;
  referenceCode: string;
  amountCents: number;
  currency: string;
  contributorEmail?: string;
  contributorName?: string;
  message?: string;
  status: CartStatus;
  decisionReason?: string;
  createdAt?: string;
  claimedAt?: string;
  decidedAt?: string;
  items: CartItem[];
}

// Owner-facing payment configuration + cart review.
export const payments = {
  listMethods: (registryId: string) =>
    api.get<{ data: PaymentMethod[] }>(`/api/payments/registries/${registryId}/payment-methods`).then((r) => r.data),
  createMethod: (registryId: string, body: Partial<PaymentMethod>) =>
    api.post<PaymentMethod>(`/api/payments/registries/${registryId}/payment-methods`, body),
  updateMethod: (id: string, body: Partial<PaymentMethod>) =>
    api.patch<PaymentMethod>(`/api/payments/payment-methods/${id}`, body),
  removeMethod: (id: string) => api.del<{ ok: boolean }>(`/api/payments/payment-methods/${id}`),
  listCarts: (registryId: string, status?: string) =>
    api
      .get<{ data: Cart[] }>(
        `/api/payments/registries/${registryId}/carts${status ? `?status=${encodeURIComponent(status)}` : ''}`,
      )
      .then((r) => r.data),
  approve: (id: string) => api.post<Cart>(`/api/payments/carts/${id}/approve`, {}),
  reject: (id: string, reason?: string) =>
    api.post<Cart>(`/api/payments/carts/${id}/reject`, { reason }),
};

// Public (contributor-facing) payment helpers.
export interface PaymentIntent {
  ok: boolean;
  id: string;
  referenceCode: string;
  amountCents: number;
  currency: string;
  status: CartStatus;
  items: CartItem[];
}

export const publicPayments = {
  methods: (slug: string) =>
    api.get<{ data: PaymentMethod[] }>(`/api/public/payments/methods?slug=${encodeURIComponent(slug)}`).then((r) => r.data),
  createIntent: (slug: string, paymentMethodId: string) =>
    api.post<PaymentIntent>('/api/public/payments/intent', { slug, paymentMethodId }),
  claim: (id: string, slug: string, message?: string) =>
    api.post<{ ok: boolean; status: CartStatus }>(`/api/public/payments/${encodeURIComponent(id)}/claim`, { slug, message }),
  cancel: (id: string, slug: string) =>
    api.post<{ ok: boolean; status: string }>(`/api/public/payments/${encodeURIComponent(id)}/cancel`, { slug }),
};

