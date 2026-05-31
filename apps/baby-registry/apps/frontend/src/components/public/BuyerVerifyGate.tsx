import { useState } from 'react';
import {
  Alert,
  Box,
  Button,
  Card,
  CardContent,
  CircularProgress,
  Container,
  Stack,
  TextField,
  Typography,
} from '@mui/material';
import MarkEmailReadIcon from '@mui/icons-material/MarkEmailRead';
import { useMutation } from '@tanstack/react-query';
import { buyer } from '../../api';

export default function BuyerVerifyGate({ slug, onVerified }: { slug: string; onVerified: () => void }) {
  const [step, setStep] = useState<'email' | 'code'>('email');
  const [email, setEmail] = useState('');
  const [name, setName] = useState('');
  const [code, setCode] = useState('');
  const [err, setErr] = useState<string | null>(null);

  const requestM = useMutation({
    mutationFn: () => buyer.request(slug, email.trim(), name.trim() || undefined),
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
                label="Your name"
                value={name}
                onChange={(e) => setName(e.target.value)}
                autoFocus
                fullWidth
                helperText="So the parents know who reserved each gift."
              />
              <TextField
                label="Your email"
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                fullWidth
              />
              {err && <Alert severity="error">{err}</Alert>}
              <Button
                variant="contained"
                size="large"
                onClick={() => requestM.mutate()}
                disabled={!email.trim() || !name.trim() || requestM.isPending}
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
