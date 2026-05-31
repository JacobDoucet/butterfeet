import { createTheme, type Theme } from '@mui/material/styles';
import { theme as baseTheme } from './theme';

export interface ThemePalette {
  bg: string;
  primary: string;
}

export const THEME_PALETTES: ThemePalette[] = [
  { bg: '#fbf7f2', primary: '#7a9e7e' },
  { bg: '#fdecd2', primary: '#d97757' },
  { bg: '#fde2e4', primary: '#c97b8b' },
  { bg: '#f4d8ff', primary: '#8b6fbf' },
  { bg: '#dbeafe', primary: '#5b8def' },
  { bg: '#d6f0f5', primary: '#3aa6a0' },
  { bg: '#fff1c1', primary: '#c79a2a' },
];

export const DEFAULT_PALETTE = THEME_PALETTES[0];

export function paletteForColor(bg?: string | null): ThemePalette {
  if (!bg) return DEFAULT_PALETTE;
  const match = THEME_PALETTES.find((p) => p.bg.toLowerCase() === bg.toLowerCase());
  return match ?? DEFAULT_PALETTE;
}

export function themeForColor(bg?: string | null): Theme {
  const p = paletteForColor(bg);
  if (p === DEFAULT_PALETTE) return baseTheme;
  return createTheme({
    ...baseTheme,
    palette: {
      ...baseTheme.palette,
      primary: { main: p.primary },
      background: { ...baseTheme.palette.background, default: p.bg },
    },
  });
}
