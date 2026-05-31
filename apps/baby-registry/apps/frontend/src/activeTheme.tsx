import { createContext, useContext, useEffect, useMemo, useState, type ReactNode } from 'react';

interface ActiveThemeContextValue {
  color: string | null;
  setColor: (color: string | null) => void;
}

const ActiveThemeContext = createContext<ActiveThemeContextValue>({
  color: null,
  setColor: () => undefined,
});

export function ActiveThemeProvider({ children }: { children: ReactNode }) {
  const [color, setColor] = useState<string | null>(null);
  const value = useMemo(() => ({ color, setColor }), [color]);
  return <ActiveThemeContext.Provider value={value}>{children}</ActiveThemeContext.Provider>;
}

export function useActiveThemeColor(): string | null {
  return useContext(ActiveThemeContext).color;
}

export function useSetActiveThemeColor(color: string | null | undefined) {
  const { setColor } = useContext(ActiveThemeContext);
  useEffect(() => {
    setColor(color ?? null);
    return () => setColor(null);
  }, [color, setColor]);
}
