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
  CircularProgress,
} from '@mui/material';
import MarkEmailReadIcon from '@mui/icons-material/MarkEmailRead';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { pub, buyer } from '../api';

export default function PublicRegistry() {
  const { slug = '' } = useParams();
  const qc = useQueryClient();

  const meQ = useQuery({
    queryKey: ['buyer', slug],
    queryFn: () => buyer.me(slug),
    retry: false,
  });

  const verified = !!meQ.data?.email;

  const regQ = useQuery({
    queryKey: ['public', slug],
    queryFn: () => pub.registry(slug),
    enabled: verified,
  });

  const [target, setTarget] = useState<string | null>(null);
  const [name, setName] = useState('');
  const [anon, setAnon] = useState(false);
  const [message, setMessage] = useState('');
  const [error, setError] = useState<string | null>(null);
  const [snack, setSnack] = useState<string | null>(null);

  const reserveM = useMutation({
    mutationFn: async () => {
      if (!target) return;
      return pub.reserve(target, { reserverName: name, isAnonymous: anon, message });
    },
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['public', slug] });
      setSnack('Perfect. We marked this as claimed. Thank you for purchasing it.');
      setTarget(null); setName(''); setAnon(false); setMessage(''); setError(null);
    },
    onError: (err) => setError((err as Error).message),
  });

  if (meQ.isLoading) {
    return (
      <Container sx={{ py: 8, textAlign: 'center' }}>
        <CircularProgress />
      </Container>
    );
  }

  if (!verified) {
    return <BuyerVerifyGate slug={slug} onVerified={() => qc.invalidateQueries({ queryKey: ['buyer', slug] })} />;
  }

  if (regQ.isLoading) return null;
  if (regQ.error || !regQ.data) {
    return (
      <Container sx={{ py: 8 }}>
        <Alert severity="warning">This registry isn't available.</Alert>
      </Container>
    );
  }
  const reg = regQ.data;
  const targetItem = reg.items.find((it) => it.id === target) ?? null;
  const hasShippingAddress = Boolean(
    reg.shippingRecipientName || reg.shippingLine1 || reg.shippingCity || reg.shippingRegion || reg.shippingPostalCode || reg.shippingCountry,
  );

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
                    <Typography
                      variant="h6"
                      sx={{
                        mb: 1,
                        overflow: 'hidden',
                        textOverflow: 'ellipsis',
                        display: '-webkit-box',
                        WebkitLineClamp: 2,
                        WebkitBoxOrient: 'vertical',
                      }}
                    >
                      {it.title}
                    </Typography>
                    <Stack direction="row" spacing={1} sx={{ mb: 2 }}>
                      {it.source && <Chip size="small" label={it.source} />}
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
          <DialogTitle>Get this gift</DialogTitle>
          <DialogContent>
            <Stack spacing={2} sx={{ mt: 1 }}>
              {targetItem && (
                <Typography variant="subtitle1" sx={{ fontWeight: 700 }}>
                  {targetItem.title}
                </Typography>
              )}
              <Alert severity="info" sx={{ '& .MuiAlert-message': { width: '100%' } }}>
                <Stack spacing={1}>
                  <Typography variant="body2"><strong>How this works</strong></Typography>
                  <Typography variant="body2">1. Open the product page and complete your purchase.</Typography>
                  <Typography variant="body2">2. Ship it to the delivery address below.</Typography>
                  <Typography variant="body2">3. Click <strong>I've bought this</strong> to mark it as claimed.</Typography>
                </Stack>
              </Alert>

              {targetItem?.productUrl ? (
                <Button
                  variant="outlined"
                  component="a"
                  href={targetItem.productUrl}
                  target="_blank"
                  rel="noreferrer"
                >
                  Open product link
                </Button>
              ) : (
                <Alert severity="warning">No product link was provided for this item.</Alert>
              )}

              {hasShippingAddress ? (
                <Box sx={{ p: 2, borderRadius: 1, bgcolor: 'action.hover' }}>
                  <Typography variant="overline" color="text.secondary">Delivery address</Typography>
                  {reg.shippingRecipientName && <Typography>{reg.shippingRecipientName}</Typography>}
                  {reg.shippingLine1 && <Typography>{reg.shippingLine1}</Typography>}
                  {reg.shippingLine2 && <Typography>{reg.shippingLine2}</Typography>}
                  {(reg.shippingCity || reg.shippingRegion || reg.shippingPostalCode) && (
                    <Typography>{[reg.shippingCity, reg.shippingRegion, reg.shippingPostalCode].filter(Boolean).join(' ')}</Typography>
                  )}
                  {reg.shippingCountry && <Typography>{reg.shippingCountry}</Typography>}
                  {reg.shippingDeliveryNotes && (
                    <Typography variant="body2" color="text.secondary" sx={{ mt: 1, whiteSpace: 'pre-wrap' }}>
                      Note: {reg.shippingDeliveryNotes}
                    </Typography>
                  )}
                </Box>
              ) : (
                <Alert severity="warning">
                  Delivery address is protected and not shown on the public registry page.
                  Share your purchase note below and the parents can send shipping details via a private link.
                </Alert>
              )}

              <TextField label="Your name" value={name} onChange={(e) => setName(e.target.value)} disabled={anon} />
              <FormControlLabel
                control={<Checkbox checked={anon} onChange={(e) => setAnon(e.target.checked)} />}
                label="Keep me anonymous"
              />
              <TextField
                label="Message to the parents (optional)"
                placeholder="Example: Ordered from Amazon, arrives next Tuesday."
                multiline
                minRows={3}
                value={message}
                onChange={(e) => setMessage(e.target.value)}
              />
              <Typography variant="caption" color="text.secondary">
                Verified as <strong>{meQ.data?.email}</strong>. The parents will see this email so they can follow up.
              </Typography>
              {error && <Alert severity="error">{error}</Alert>}
            </Stack>
          </DialogContent>
          <DialogActions>
            <Button onClick={() => setTarget(null)}>Cancel</Button>
            <Button variant="contained" onClick={() => reserveM.mutate()} disabled={!anon && !name.trim()}>
              I've bought this
            </Button>
          </DialogActions>
        </Dialog>

        <Snackbar open={!!snack} autoHideDuration={4000} onClose={() => setSnack(null)} message={snack ?? ''} />
      </Container>
    </Box>
  );
}

function BuyerVerifyGate({ slug, onVerified }: { slug: string; onVerified: () => void }) {
  const [step, setStep] = useState<'email' | 'code'>('email');
  const [email, setEmail] = useState('');
  const [code, setCode] = useState('');
  const [err, setErr] = useState<string | null>(null);

  const requestM = useMutation({
    mutationFn: () => buyer.request(slug, email.trim()),
    onSuccess: () => {
      setErr(null);
      setStep('code');
    },
    onError: (e) => setErr((e as Error).message),
  });
  const confirmM = useMutation({
    mutationFn: () => buyer.confirm(slug, email.trim(), code.trim()),
    onSuccess: () => {
      setErr(null);
      onVerified();
    },
    onError: (e) => setErr((e as Error).message),
  });

  return (
    <Container maxWidth="sm" sx={{ py: 8 }}>
      <Card>
        <CardContent>
          <Stack alignItems="center" spacing={1} sx={{ mb: 3 }}>
            <MarkEmailReadIcon color="primary" sx={{ fontSize: 40 }} />
            <Typography variant="h5" textAlign="center">
              Verify your email to view this registry
            </Typography>
            <Typography variant="body2" color="text.secondary" textAlign="center">
              We send a 6-digit code to your inbox so the parents know who's coming to their gift list.
              No account, no marketing.
            </Typography>
          </Stack>

          {step === 'email' && (
            <Stack spacing={2}>
              <TextField
                label="Your email"
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                autoFocus
                fullWidth
              />
              {err && <Alert severity="error">{err}</Alert>}
              <Button
                variant="contained"
                size="large"
                onClick={() => requestM.mutate()}
                disabled={!email.trim() || requestM.isPending}
              >
                {requestM.isPending ? <CircularProgress size={20} /> : 'Send code'}
              </Button>
            </Stack>
          )}

          {step === 'code' && (
            <Stack spacing={2}>
              <Typography variant="body2" color="text.secondary">
                We sent a code to <strong>{email}</strong>. It expires in 15 minutes.
              </Typography>
              <TextField
                label="6-digit code"
                value={code}
                onChange={(e) => setCode(e.target.value.replace(/\D/g, '').slice(0, 6))}
                autoFocus
                fullWidth
                inputProps={{ inputMode: 'numeric', pattern: '[0-9]*', style: { letterSpacing: 6, fontSize: 22, textAlign: 'center' } }}
              />
              {err && <Alert severity="error">{err}</Alert>}
              <Stack direction="row" spacing={1}>
                <Button variant="text" onClick={() => { setStep('email'); setCode(''); setErr(null); }}>
                  Use a different email
                </Button>
                <Box sx={{ flex: 1 }} />
                <Button
                  variant="contained"
                  onClick={() => confirmM.mutate()}
                  disabled={code.length !== 6 || confirmM.isPending}
                >
                  {confirmM.isPending ? <CircularProgress size={20} /> : 'Verify'}
                </Button>
              </Stack>
              <Button
                size="small"
                variant="text"
                onClick={() => requestM.mutate()}
                disabled={requestM.isPending}
              >
                Resend code
              </Button>
            </Stack>
          )}
        </CardContent>
      </Card>
    </Container>
  );
}
