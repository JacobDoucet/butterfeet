import { createContext, useContext, useEffect, useMemo, useState, type ReactNode } from 'react';
import { SUPPORTED_CURRENCIES } from './api';

// detectViewerCurrency picks a sensible default display currency from the
// browser locale, limited to the currencies we support. Falls back to USD.
function detectViewerCurrency(): string {
  const supported = SUPPORTED_CURRENCIES.map((c) => c.code);
  try {
    const stored = localStorage.getItem('viewerCurrency');
    if (stored && supported.includes(stored)) return stored;
  } catch {
    // ignore storage access errors
  }
  const localeToCurrency: Record<string, string> = {
    US: 'USD', CA: 'CAD', GB: 'GBP',
    IE: 'EUR', DE: 'EUR', FR: 'EUR', ES: 'EUR', IT: 'EUR', NL: 'EUR',
    BE: 'EUR', AT: 'EUR', PT: 'EUR', FI: 'EUR', GR: 'EUR',
  };
  try {
    const locales = typeof navigator !== 'undefined' ? [navigator.language, ...(navigator.languages || [])] : [];
    for (const loc of locales) {
      const region = loc?.split('-')[1]?.toUpperCase();
      if (region && localeToCurrency[region]) return localeToCurrency[region];
    }
  } catch {
    // ignore locale detection errors
  }
  return 'USD';
}

interface ViewerCurrencyContextValue {
  currency: string;
  setCurrency: (currency: string) => void;
}

const ViewerCurrencyContext = createContext<ViewerCurrencyContextValue>({
  currency: 'USD',
  setCurrency: () => undefined,
});

export function ViewerCurrencyProvider({ children }: { children: ReactNode }) {
  const [currency, setCurrency] = useState<string>(detectViewerCurrency);
  useEffect(() => {
    try {
      localStorage.setItem('viewerCurrency', currency);
    } catch {
      // ignore storage access errors
    }
  }, [currency]);
  const value = useMemo(() => ({ currency, setCurrency }), [currency]);
  return <ViewerCurrencyContext.Provider value={value}>{children}</ViewerCurrencyContext.Provider>;
}

export function useViewerCurrency(): ViewerCurrencyContextValue {
  return useContext(ViewerCurrencyContext);
}
