import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import {
  Container,
  Typography,
  Card,
  CardContent,
  Button,
  Stack,
  Grid,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  Alert,
} from '@mui/material';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { registries, auth, type Registry, type Me } from '../api';

export default function OwnerDashboard() {
  const nav = useNavigate();
  const qc = useQueryClient();

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
  const [slug, setSlug] = useState('');
  const [title, setTitle] = useState('');
  const [parentNames, setParentNames] = useState('');
  const [error, setError] = useState<string | null>(null);

  const createM = useMutation({
    mutationFn: async () => {
      return registries.create({
        slug: slug.trim().toLowerCase(),
        title: title.trim(),
        parentNames: parentNames.trim(),
        isPublic: true,
        ownerId: meQ.data?.id,
      });
    },
    onSuccess: (reg) => {
      qc.invalidateQueries({ queryKey: ['registries'] });
      setOpen(false);
      nav(`/owner/r/${reg.slug}`);
    },
    onError: (err) => setError((err as Error).message),
  });

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
        {(listQ.data?.data ?? []).map((r: Registry) => (
          <Grid item xs={12} sm={6} key={r.id}>
            <Card>
              <CardContent>
                <Typography variant="h6">{r.title}</Typography>
                <Typography color="text.secondary" sx={{ mb: 2 }}>
                  /r/{r.slug}
                </Typography>
                <Stack direction="row" spacing={1}>
                  <Button size="small" component={Link} to={`/owner/r/${r.slug}`}>
                    Manage
                  </Button>
                  <Button size="small" component={Link} to={`/r/${r.slug}`}>
                    View public
                  </Button>
                </Stack>
              </CardContent>
            </Card>
          </Grid>
        ))}
        {listQ.data && listQ.data.data.length === 0 && (
          <Grid item xs={12}>
            <Typography color="text.secondary">No registries yet. Create your first!</Typography>
          </Grid>
        )}
      </Grid>

      <Dialog open={open} onClose={() => setOpen(false)} fullWidth maxWidth="sm">
        <DialogTitle>New registry</DialogTitle>
        <DialogContent>
          <Stack spacing={2} sx={{ mt: 1 }}>
            <TextField label="Title (e.g. Baby Smith)" value={title} onChange={(e) => setTitle(e.target.value)} />
            <TextField
              label="URL slug"
              helperText="lowercase letters, numbers, dashes"
              value={slug}
              onChange={(e) => setSlug(e.target.value.replace(/[^a-z0-9-]/g, ''))}
            />
            <TextField label="Parent names (optional)" value={parentNames} onChange={(e) => setParentNames(e.target.value)} />
            {error && <Alert severity="error">{error}</Alert>}
          </Stack>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpen(false)}>Cancel</Button>
          <Button onClick={() => createM.mutate()} variant="contained" disabled={!slug || !title}>
            Create
          </Button>
        </DialogActions>
      </Dialog>
    </Container>
  );
}
