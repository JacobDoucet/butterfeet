import { Routes, Route, Link, useNavigate, useLocation } from 'react-router-dom';
import { AppBar, Toolbar, Box, Button, Typography, Container } from '@mui/material';
import LandingPage from './pages/Landing';
import LoginPage from './pages/Login';
import AuthCallbackPage from './pages/AuthCallback';
import OwnerDashboardPage from './pages/OwnerDashboard';
import RegistryEditorPage from './pages/RegistryEditor';
import PublicRegistryPage from './pages/PublicRegistry';
import ShipPage from './pages/Ship';
import { useQuery, useQueryClient } from '@tanstack/react-query';
import { auth, type Me } from './api';
import BrandLogo from './components/BrandLogo';

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
    ? { gap: 2, minHeight: { xs: 64, sm: 68 }, justifyContent: 'center' }
    : { gap: 2 };

  return (
    <Box sx={{ minHeight: '100vh', display: 'flex', flexDirection: 'column' }}>
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
            backgroundColor: 'background.paper',
            WebkitMaskImage: `url("data:image/svg+xml;utf8,<svg xmlns='http://www.w3.org/2000/svg' width='80' height='40' viewBox='0 0 80 40' preserveAspectRatio='none'><path d='M0 0 Q40 56 80 0 Z' fill='black'/></svg>")`,
            maskImage: `url("data:image/svg+xml;utf8,<svg xmlns='http://www.w3.org/2000/svg' width='80' height='40' viewBox='0 0 80 40' preserveAspectRatio='none'><path d='M0 0 Q40 56 80 0 Z' fill='black'/></svg>")`,
            WebkitMaskSize: '80px 40px',
            maskSize: '80px 40px',
            WebkitMaskRepeat: 'repeat-x',
            maskRepeat: 'repeat-x',
            WebkitMaskPosition: 'top left',
            maskPosition: 'top left',
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
        <Box component="footer" sx={{ py: 4, textAlign: 'center', color: 'text.secondary' }}>
          <Container>
            <Typography variant="body2">made with care for new parents</Typography>
          </Container>
        </Box>
      )}
    </Box>
  );
}

export default function App() {
  return (
    <Shell>
      <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/auth/callback" element={<AuthCallbackPage />} />
        <Route path="/owner" element={<OwnerDashboardPage />} />
        <Route path="/owner/r/:slug" element={<RegistryEditorPage />} />
        <Route path="/r/:slug" element={<PublicRegistryPage />} />
        <Route path="/ship" element={<ShipPage />} />
      </Routes>
    </Shell>
  );
}
