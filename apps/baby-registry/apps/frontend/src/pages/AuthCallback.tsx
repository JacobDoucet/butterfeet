import { useEffect, useState } from 'react';
import { useSearchParams, useNavigate } from 'react-router-dom';
import { Container, Typography, CircularProgress, Alert, Stack } from '@mui/material';
import { useQueryClient } from '@tanstack/react-query';
import { auth } from '../api';

export default function AuthCallback() {
  const [sp] = useSearchParams();
  const nav = useNavigate();
  const qc = useQueryClient();
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const token = sp.get('token');
    if (!token) {
      setError('Missing token.');
      return;
    }
    auth
      .verify(token)
      .then(async () => {
        await qc.invalidateQueries({ queryKey: ['me'] });
        nav('/owner', { replace: true });
      })
      .catch((err) => setError((err as Error).message));
  }, [sp, nav, qc]);

  return (
    <Container maxWidth="sm" sx={{ py: 10 }}>
      <Stack spacing={2} alignItems="center">
        {error ? <Alert severity="error">{error}</Alert> : <CircularProgress />}
        <Typography color="text.secondary">{error ? 'Try requesting a new link.' : 'Signing you in…'}</Typography>
      </Stack>
    </Container>
  );
}
