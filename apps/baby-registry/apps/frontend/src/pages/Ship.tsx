import { useEffect, useState } from 'react';
import { Box, Card, CardContent, Stack, Typography, Alert, Button, CircularProgress } from '@mui/material';
import LockOpenIcon from '@mui/icons-material/LockOpen';
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import { API_BASE } from '../api';

interface ResolvedAddress {
  registryTitle: string;
  recipientName: string;
  line1: string;
  line2?: string;
  city: string;
  region: string;
  postalCode: string;
  country: string;
  deliveryNotes?: string;
  expiresAt: string;
}

export default function ShipPage() {
  const [state, setState] = useState<
    { kind: 'loading' } | { kind: 'error'; message: string } | { kind: 'ok'; address: ResolvedAddress }
  >({ kind: 'loading' });

  useEffect(() => {
    const token = window.location.hash.replace(/^#/, '').split('&')
      .map((p) => p.split('='))
      .find(([k]) => k === 'tok')?.[1];

    if (!token) {
      setState({ kind: 'error', message: 'No access token in URL.' });
      return;
    }
    fetch(`${API_BASE}/api/public/shipping/resolve`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ token }),
    })
      .then(async (r) => {
        const data = await r.json().catch(() => ({}));
        if (!r.ok) {
          setState({ kind: 'error', message: data.error ?? `Error (${r.status})` });
          return;
        }
        setState({ kind: 'ok', address: data as ResolvedAddress });
      })
      .catch((e) => setState({ kind: 'error', message: e.message ?? 'Network error' }));
  }, []);

  const copyAddress = (addr: ResolvedAddress) => {
    const lines = [
      addr.recipientName,
      addr.line1,
      addr.line2,
      [addr.city, addr.region, addr.postalCode].filter(Boolean).join(' '),
      addr.country,
    ].filter(Boolean);
    navigator.clipboard.writeText(lines.join('\n'));
  };

  return (
    <Box sx={{ maxWidth: 600, mx: 'auto', mt: 6, px: 2 }}>
      <Card>
        <CardContent>
          <Stack direction="row" alignItems="center" spacing={1} sx={{ mb: 2 }}>
            <LockOpenIcon color="primary" />
            <Typography variant="h6">Shipping address</Typography>
          </Stack>

          {state.kind === 'loading' && (
            <Stack direction="row" alignItems="center" spacing={2}>
              <CircularProgress size={20} />
              <Typography variant="body2" color="text.secondary">Verifying your link…</Typography>
            </Stack>
          )}

          {state.kind === 'error' && <Alert severity="error">{state.message}</Alert>}

          {state.kind === 'ok' && (
            <Stack spacing={2}>
              <Typography variant="body2" color="text.secondary">
                For: <strong>{state.address.registryTitle}</strong>
              </Typography>
              <Box sx={{ p: 2, bgcolor: 'action.hover', borderRadius: 1, fontFamily: 'inherit' }}>
                <Typography>{state.address.recipientName}</Typography>
                <Typography>{state.address.line1}</Typography>
                {state.address.line2 && <Typography>{state.address.line2}</Typography>}
                <Typography>
                  {[state.address.city, state.address.region, state.address.postalCode]
                    .filter(Boolean)
                    .join(' ')}
                </Typography>
                <Typography>{state.address.country}</Typography>
              </Box>
              {state.address.deliveryNotes && (
                <Alert severity="info" sx={{ '& .MuiAlert-message': { whiteSpace: 'pre-wrap' } }}>
                  {state.address.deliveryNotes}
                </Alert>
              )}
              <Stack direction="row" spacing={1} justifyContent="flex-end">
                <Button
                  startIcon={<ContentCopyIcon />}
                  variant="outlined"
                  onClick={() => copyAddress(state.address)}
                >
                  Copy address
                </Button>
              </Stack>
              <Typography variant="caption" color="text.secondary">
                This link expires on {new Date(state.address.expiresAt).toLocaleString()}.
                Please don’t share it — it grants address access to anyone who has it.
              </Typography>
            </Stack>
          )}
        </CardContent>
      </Card>
    </Box>
  );
}
