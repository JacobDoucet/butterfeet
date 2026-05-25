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

export const registries = {
  list: () => api.get<{ data: Registry[]; total: number }>('/api/registries/'),
  create: (body: Partial<Registry>) => api.post<Registry>('/api/registries/', body),
  update: (id: string, body: Partial<Registry>) => api.patch<Registry>(`/api/registries/${id}`, body),
  remove: (id: string) => api.del<void>(`/api/registries/${id}`),
};

export const items = {
  listForRegistry: (registryId: string) =>
    api.get<{ data: RegistryItem[]; total: number }>(`/api/registry-items/?registryIdEq=${registryId}`),
  create: (body: Partial<RegistryItem>) => api.post<RegistryItem>('/api/registry-items/', body),
  update: (id: string, body: Partial<RegistryItem>) => api.patch<RegistryItem>(`/api/registry-items/${id}`, body),
  remove: (id: string) => api.del<void>(`/api/registry-items/${id}`),
};

export const pub = {
  registry: (slug: string) => api.get<PublicRegistry>(`/api/public/r/${encodeURIComponent(slug)}`),
  reserve: (itemId: string, body: { reserverName: string; isAnonymous: boolean; message: string; contactEmail?: string; quantity?: number }) =>
    api.post<{ ok: boolean }>(`/api/public/items/${itemId}/reserve`, body),
};

export const scrape = {
  url: (u: string) => api.get<ScrapeResult>(`/api/scrape?url=${encodeURIComponent(u)}`),
};
