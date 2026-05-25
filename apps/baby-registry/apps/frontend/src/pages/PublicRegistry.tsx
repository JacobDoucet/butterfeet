import { useState } from 'react';
import { useParams } from 'react-router-dom';
import {
  Container,
  Typography,
  Card,
  CardContent,
  CardMedia,
  Button,
  Stack,
  Grid,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  Alert,
  Chip,
  Box,
  FormControlLabel,
  Checkbox,
  Snackbar,
} from '@mui/material';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { pub } from '../api';

export default function PublicRegistry() {
  const { slug = '' } = useParams();
  const qc = useQueryClient();
  const regQ = useQuery({ queryKey: ['public', slug], queryFn: () => pub.registry(slug) });

  const [target, setTarget] = useState<string | null>(null);
  const [name, setName] = useState('');
  const [anon, setAnon] = useState(false);
  const [message, setMessage] = useState('');
  const [email, setEmail] = useState('');
  const [error, setError] = useState<string | null>(null);
  const [snack, setSnack] = useState<string | null>(null);

  const reserveM = useMutation({
    mutationFn: async () => {
      if (!target) return;
      return pub.reserve(target, { reserverName: name, isAnonymous: anon, message, contactEmail: email });
    },
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['public', slug] });
      setSnack('Thank you! Your reservation is recorded.');
      setTarget(null); setName(''); setAnon(false); setMessage(''); setEmail(''); setError(null);
    },
    onError: (err) => setError((err as Error).message),
  });

  if (regQ.isLoading) return null;
  if (regQ.error || !regQ.data) {
    return (
      <Container sx={{ py: 8 }}>
        <Alert severity="warning">This registry isn't available.</Alert>
      </Container>
    );
  }
  const reg = regQ.data;

  return (
    <Box sx={{ bgcolor: reg.themeColor || 'background.default', minHeight: '100%' }}>
      <Container maxWidth="lg" sx={{ py: 6 }}>
        <Stack alignItems="center" sx={{ textAlign: 'center', mb: 5 }}>
          <Typography variant="h3" gutterBottom>{reg.title}</Typography>
          {reg.parentNames && (
            <Typography variant="subtitle1" color="text.secondary">for {reg.parentNames}</Typography>
          )}
          {reg.welcomeMessage && (
            <Typography sx={{ mt: 2, maxWidth: 640 }}>{reg.welcomeMessage}</Typography>
          )}
        </Stack>

        <Grid container spacing={3}>
          {reg.items.map((it) => {
            const remaining = Math.max(0, (it.quantity || 1) - (it.reserved || 0));
            const claimed = remaining === 0;
            return (
              <Grid item xs={12} sm={6} md={4} key={it.id}>
                <Card sx={{ opacity: claimed ? 0.6 : 1, height: '100%', display: 'flex', flexDirection: 'column' }}>
                  {it.imageUrl && (
                    <CardMedia component="img" image={it.imageUrl} sx={{ aspectRatio: '1', objectFit: 'contain', bgcolor: '#f4ede3' }} />
                  )}
                  <CardContent sx={{ flexGrow: 1, display: 'flex', flexDirection: 'column' }}>
                    <Typography variant="h6" sx={{ mb: 1 }}>{it.title}</Typography>
                    <Stack direction="row" spacing={1} sx={{ mb: 2 }}>
                      {it.source && <Chip size="small" label={it.source} />}
                      {it.priceCents ? (
                        <Chip size="small" label={`${((it.priceCents || 0) / 100).toFixed(2)} ${it.currency || ''}`.trim()} variant="outlined" />
                      ) : null}
                      {claimed && <Chip size="small" color="success" label="Reserved" />}
                    </Stack>
                    <Box sx={{ flexGrow: 1 }} />
                    <Stack direction="row" spacing={1}>
                      {it.productUrl && (
                        <Button size="small" variant="outlined" component="a" href={it.productUrl} target="_blank" rel="noreferrer">
                          View product
                        </Button>
                      )}
                      <Button
                        size="small"
                        variant="contained"
                        disabled={claimed}
                        onClick={() => setTarget(it.id)}
                      >
                        {claimed ? 'Taken' : "I'll get this"}
                      </Button>
                    </Stack>
                  </CardContent>
                </Card>
              </Grid>
            );
          })}
        </Grid>

        <Dialog open={!!target} onClose={() => setTarget(null)} fullWidth maxWidth="sm">
          <DialogTitle>Reserve this gift</DialogTitle>
          <DialogContent>
            <Stack spacing={2} sx={{ mt: 1 }}>
              <TextField label="Your name" value={name} onChange={(e) => setName(e.target.value)} disabled={anon} />
              <FormControlLabel
                control={<Checkbox checked={anon} onChange={(e) => setAnon(e.target.checked)} />}
                label="Keep me anonymous"
              />
              <TextField label="Message to the parents (optional)" multiline minRows={3} value={message} onChange={(e) => setMessage(e.target.value)} />
              <TextField label="Your email (optional)" value={email} onChange={(e) => setEmail(e.target.value)} />
              {error && <Alert severity="error">{error}</Alert>}
            </Stack>
          </DialogContent>
          <DialogActions>
            <Button onClick={() => setTarget(null)}>Cancel</Button>
            <Button variant="contained" onClick={() => reserveM.mutate()} disabled={!anon && !name.trim()}>
              Reserve
            </Button>
          </DialogActions>
        </Dialog>

        <Snackbar open={!!snack} autoHideDuration={4000} onClose={() => setSnack(null)} message={snack ?? ''} />
      </Container>
    </Box>
  );
}
