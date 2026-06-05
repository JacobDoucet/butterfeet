import { Routes, Route, Link, useNavigate, useLocation } from 'react-router-dom';
import { Suspense, lazy } from 'react';
import { AppBar, Toolbar, Box, Button, Typography, Container, Select, MenuItem, CircularProgress } from '@mui/material';
import { ThemeProvider } from '@mui/material/styles';
import LandingPage from './pages/Landing';
const LoginPage = lazy(() => import('./pages/Login'));
const AuthCallbackPage = lazy(() => import('./pages/AuthCallback'));
const OwnerDashboardPage = lazy(() => import('./pages/OwnerDashboard'));
const RegistryEditorPage = lazy(() => import('./pages/RegistryEditor'));
const PublicRegistryPage = lazy(() => import('./pages/PublicRegistry'));
const ShipPage = lazy(() => import('./pages/Ship'));
import { useQuery, useQueryClient } from '@tanstack/react-query';
import { auth, type Me, SUPPORTED_CURRENCIES } from './api';
import BrandLogo from './components/BrandLogo';
import { useActiveThemeColor } from './activeTheme';
import { useViewerCurrency } from './viewerCurrency';
import { themeForColor } from './themePalettes';

function Shell({ children }: { children: React.ReactNode }) {
  const nav = useNavigate();
  const location = useLocation();
  const qc = useQueryClient();
  const { data: me } = useQuery<Me | null>({
    queryKey: ['me'],
    queryFn: async () => {
      try {
        return await auth.me();
      } catch {
        return null;
      }
    },
  });

  const handleLogout = async () => {
    await auth.logout().catch(() => undefined);
    qc.setQueryData(['me'], null);
    nav('/');
  };

  const hideHeaderActions = location.pathname.startsWith('/r/');
  const hideLogo = location.pathname === '/';
  const toolbarSx = hideHeaderActions
    ? { gap: 2, minHeight: { xs: 64, sm: 68 }, justifyContent: 'center', position: 'relative' }
    : { gap: 2 };

  const { currency, setCurrency } = useViewerCurrency();

  const activeColor = useActiveThemeColor();
  const themed = themeForColor(activeColor);

  return (
    <ThemeProvider theme={themed}>
    <Box sx={{ minHeight: '100vh', display: 'flex', flexDirection: 'column', bgcolor: activeColor ?? undefined }}>
      <AppBar
        position="sticky"
        color="transparent"
        elevation={0}
        sx={{
          bgcolor: 'background.paper',
          overflow: 'visible',
          '&::after': {
            content: '""',
            position: 'absolute',
            top: '100%',
            left: 0,
            right: 0,
            height: 40,
            pointerEvents: 'none',
            backgroundImage: `url("data:image/svg+xml;utf8,<svg xmlns='http://www.w3.org/2000/svg' width='80' height='40' viewBox='0 0 80 40' preserveAspectRatio='none'><path d='M0 0 Q40 56 80 0 Z' fill='%23ffffff'/><path d='M0 0 Q40 56 80 0' fill='none' stroke='rgba(0,0,0,0.12)' stroke-width='1' vector-effect='non-scaling-stroke'/></svg>")`,
            backgroundRepeat: 'repeat-x',
            backgroundSize: '80px 40px',
            backgroundPosition: 'top left',
          },
        }}
      >
        <Toolbar sx={toolbarSx}>
          <Typography
            component={Link}
            to="/"
            variant="h6"
            sx={{ textDecoration: 'none', color: 'inherit', fontWeight: 700, display: hideLogo ? 'none' : 'flex', alignItems: 'center' }}
          >
            <BrandLogo variant="lockup" height={35} markScale={1.12} wordmarkScale={1.1} />
          </Typography>
          {!hideHeaderActions && <Box sx={{ flexGrow: 1 }} />}
          {hideHeaderActions && (
            <Select
              size="small"
              value={currency}
              onChange={(e) => setCurrency(e.target.value)}
              sx={{
                position: 'absolute',
                right: { xs: 8, sm: 16 },
                top: '50%',
                transform: 'translateY(-50%)',
                '& .MuiSelect-select': { py: 0.5, fontWeight: 600, fontSize: '0.85rem' },
                borderRadius: 2,
                bgcolor: 'background.paper',
              }}
            >
              {SUPPORTED_CURRENCIES.map((c) => (
                <MenuItem key={c.code} value={c.code}>{c.code}</MenuItem>
              ))}
            </Select>
          )}
          {!hideHeaderActions && (me ? (
            <>
              <Button component={Link} to="/owner" color="primary">
                My registries
              </Button>
              <Button onClick={handleLogout} color="inherit">
                Sign out
              </Button>
            </>
          ) : (
            <Button component={Link} to="/login" variant="contained" color="primary">
              Sign in
            </Button>
          ))}
        </Toolbar>
      </AppBar>
      <Box sx={{ flexGrow: 1 }}>{children}</Box>
      {!hideHeaderActions && (
        <Box component="footer" sx={{ py: 4, textAlign: 'center', color: 'text.secondary', bgcolor: activeColor ?? undefined }}>
          <Container>
            <Typography variant="body2">made with care for new parents</Typography>
          </Container>
        </Box>
      )}
    </Box>
    </ThemeProvider>
  );
}

export default function App() {
  return (
    <Shell>
      <Suspense
        fallback={
          <Box sx={{ display: 'flex', justifyContent: 'center', py: 8 }}>
            <CircularProgress />
          </Box>
        }
      >
        <Routes>
          <Route path="/" element={<LandingPage />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/auth/callback" element={<AuthCallbackPage />} />
          <Route path="/owner" element={<OwnerDashboardPage />} />
          <Route path="/owner/r/:slug" element={<RegistryEditorPage />} />
          <Route path="/r/:slug" element={<PublicRegistryPage />} />
          <Route path="/ship" element={<ShipPage />} />
        </Routes>
      </Suspense>
    </Shell>
  );
}
