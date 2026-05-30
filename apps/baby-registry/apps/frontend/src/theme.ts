import { createTheme } from '@mui/material/styles';

export const theme = createTheme({
  palette: {
    mode: 'light',
    primary: { main: '#7a9e7e' },
    secondary: { main: '#e8a87c' },
    background: { default: '#fbf7f2', paper: '#ffffff' },
    text: { primary: '#2d2a26' },
  },
  shape: { borderRadius: 14 },
  typography: {
    fontFamily: '"Inter", "Helvetica", "Arial", sans-serif',
    h1: { fontWeight: 700, letterSpacing: '0.02em', textTransform: 'uppercase' },
    h2: { fontWeight: 700, letterSpacing: '0.02em', textTransform: 'uppercase' },
    h3: { fontWeight: 700, letterSpacing: '0.02em', textTransform: 'uppercase' },
    h4: { fontWeight: 700, letterSpacing: '0.02em', textTransform: 'uppercase' },
    button: { textTransform: 'none', fontWeight: 600 },
  },
  components: {
    MuiButton: {
      defaultProps: { disableElevation: true },
      styleOverrides: { root: { borderRadius: 12, paddingInline: 18 } },
    },
    MuiCard: {
      styleOverrides: { root: { borderRadius: 18 } },
    },
    MuiTextField: { defaultProps: { fullWidth: true } },
    MuiTab: {
      styleOverrides: {
        root: { textTransform: 'uppercase', letterSpacing: '0.08em', fontWeight: 600 },
      },
    },
  },
});
