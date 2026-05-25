import { Routes, Route, Link, useNavigate } from 'react-router-dom';
import { AppBar, Toolbar, Box, Button, Typography, Container } from '@mui/material';
import LandingPage from './pages/Landing';
import LoginPage from './pages/Login';
import AuthCallbackPage from './pages/AuthCallback';
import OwnerDashboardPage from './pages/OwnerDashboard';
import RegistryEditorPage from './pages/RegistryEditor';
import PublicRegistryPage from './pages/PublicRegistry';
import { useQuery, useQueryClient } from '@tanstack/react-query';
import { auth, type Me } from './api';

function Shell({ children }: { children: React.ReactNode }) {
  const nav = useNavigate();
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

  return (
    <Box sx={{ minHeight: '100vh', display: 'flex', flexDirection: 'column' }}>
      <AppBar position="sticky" color="transparent" elevation={0} sx={{ bgcolor: 'background.paper', borderBottom: '1px solid rgba(0,0,0,0.06)' }}>
        <Toolbar sx={{ gap: 2 }}>
          <Typography
            component={Link}
            to="/"
            variant="h6"
            sx={{ textDecoration: 'none', color: 'inherit', fontWeight: 700 }}
          >
            🌿 little nest
          </Typography>
          <Box sx={{ flexGrow: 1 }} />
          {me ? (
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
          )}
        </Toolbar>
      </AppBar>
      <Box sx={{ flexGrow: 1 }}>{children}</Box>
      <Box component="footer" sx={{ py: 4, textAlign: 'center', color: 'text.secondary' }}>
        <Container>
          <Typography variant="body2">made with care for new parents</Typography>
        </Container>
      </Box>
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
      </Routes>
    </Shell>
  );
}
