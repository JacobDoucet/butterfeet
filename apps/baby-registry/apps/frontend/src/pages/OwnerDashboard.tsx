import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { Alert, Button, Container, Grid, Stack, Typography } from '@mui/material';
import { useQuery } from '@tanstack/react-query';
import { registries, auth, type Registry, type Me } from '../api';
import NewRegistryDialog from '../components/owner/NewRegistryDialog';
import RegistryCard from '../components/owner/RegistryCard';

export default function OwnerDashboard() {
  const nav = useNavigate();

  const meQ = useQuery<Me | null>({
    queryKey: ['me'],
    queryFn: async () => {
      try {
        return await auth.me();
      } catch {
        return null;
      }
    },
  });

  const listQ = useQuery({
    queryKey: ['registries'],
    queryFn: () => registries.list(),
    enabled: !!meQ.data,
  });

  const [open, setOpen] = useState(false);

  if (meQ.isLoading) return null;
  if (!meQ.data) {
    return (
      <Container maxWidth="sm" sx={{ py: 8 }}>
        <Alert severity="info">
          Please <Link to="/login">sign in</Link> to manage registries.
        </Alert>
      </Container>
    );
  }

  const list = listQ.data?.data ?? [];

  return (
    <Container maxWidth="md" sx={{ py: 6 }}>
      <Stack direction="row" alignItems="center" sx={{ mb: 4 }}>
        <Typography variant="h4" sx={{ flexGrow: 1 }}>
          Your registries
        </Typography>
        <Button variant="contained" onClick={() => setOpen(true)}>
          New registry
        </Button>
      </Stack>

      <Grid container spacing={2}>
        {list.map((r: Registry) => (
          <Grid item xs={12} sm={6} key={r.id}>
            <RegistryCard reg={r} />
          </Grid>
        ))}
        {listQ.data && list.length === 0 && (
          <Grid item xs={12}>
            <Typography color="text.secondary">No registries yet. Create your first!</Typography>
          </Grid>
        )}
      </Grid>

      <NewRegistryDialog
        open={open}
        onClose={() => setOpen(false)}
        me={meQ.data}
        existingSlugs={list.map((r) => r.slug)}
        onCreated={(reg) => nav(`/owner/r/${reg.slug}`)}
      />
    </Container>
  );
}
