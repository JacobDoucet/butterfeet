import { useState } from 'react';
import {
  Alert,
  Box,
  Button,
  Card,
  CardContent,
  Chip,
  CircularProgress,
  Container,
  Stack,
  TextField,
  Typography,
} from '@mui/material';
import CheckCircleOutlineIcon from '@mui/icons-material/CheckCircleOutline';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import { useMutation } from '@tanstack/react-query';
import { pub, type GatedRegistry, type RegistryAccessRequestStatus } from '../../api';

export default function RegistryAccessGate({
  slug,
  viewerEmail,
  gated,
  onSubmitted,
}: {
  slug: string;
  viewerEmail: string;
  gated: GatedRegistry;
  onSubmitted: () => void;
}) {
  const initialStatus: RegistryAccessRequestStatus = gated.accessRequestStatus;
  const [name, setName] = useState('');
  const [note, setNote] = useState('');
  const [status, setStatus] = useState<RegistryAccessRequestStatus>(initialStatus);
  const [err, setErr] = useState<string | null>(null);

  const requestM = useMutation({
    mutationFn: () =>
      pub.requestRegistryAccess({ slug, name: name.trim() || undefined, note: note.trim() || undefined }),
    onSuccess: (res) => {
      setErr(null);
      const next: RegistryAccessRequestStatus =
        res.status === 'pending' ? 'pending' : res.status === 'rejected' ? 'rejected' : 'pending';
      setStatus(next);
      if (res.status === 'approved') onSubmitted();
    },
    onError: (e) => setErr((e as Error).message),
  });

  const parentsLabel = gated.ownerDisplayName?.trim() || gated.parentNames?.trim() || 'the parents';
  const titleLabel = gated.title || 'their registry';

  return (
    <Container maxWidth="sm" sx={{ py: { xs: 4, sm: 8 } }}>
      <Card sx={{ overflow: 'hidden' }}>
        {gated.coverImageUrl && (
          <Box
            sx={{
              height: 140,
              backgroundImage: `url(${gated.coverImageUrl})`,
              backgroundSize: 'cover',
              backgroundPosition: 'center',
            }}
          />
        )}
        <CardContent sx={{ p: { xs: 3, sm: 5 } }}>
          <Stack alignItems="center" spacing={1.5} sx={{ mb: 3, textAlign: 'center' }}>
            <Box
              sx={{
                width: 56,
                height: 56,
                borderRadius: '50%',
                bgcolor: status === 'pending' ? 'secondary.light' : 'primary.light',
                color: status === 'pending' ? 'secondary.dark' : 'primary.dark',
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
                mb: 0.5,
              }}
            >
              {status === 'pending' ? (
                <CheckCircleOutlineIcon sx={{ fontSize: 32 }} />
              ) : (
                <LockOutlinedIcon sx={{ fontSize: 30 }} />
              )}
            </Box>
            <Typography variant="overline" color="text.secondary" sx={{ letterSpacing: 1.4 }}>
              {titleLabel}
            </Typography>
            <Typography variant="h5" sx={{ fontWeight: 700 }}>
              {status === 'pending'
                ? "You're on the list"
                : status === 'rejected'
                  ? 'Access not available'
                  : `Ask ${parentsLabel} for access`}
            </Typography>
            <Typography variant="body2" color="text.secondary" sx={{ maxWidth: 380 }}>
              {status === 'pending' ? (
                <>
                  We've let {parentsLabel} know you'd like to view their registry. You'll get an email at{' '}
                  <strong>{viewerEmail}</strong> as soon as they approve.
                </>
              ) : status === 'rejected' ? (
                <>
                  {parentsLabel} have chosen to keep this registry private. If you think this is a mistake, please
                  reach out to them directly.
                </>
              ) : (
                <>
                  This registry is private. Send a short note to {parentsLabel} and they'll let you in. You're
                  verified as <strong>{viewerEmail}</strong>.
                </>
              )}
            </Typography>
          </Stack>

          {status === 'none' && (
            <Stack spacing={2}>
              <TextField
                label="Your name (optional)"
                value={name}
                onChange={(e) => setName(e.target.value)}
                placeholder="So they know who's asking"
                fullWidth
              />
              <TextField
                label="A short note (optional)"
                value={note}
                onChange={(e) => setNote(e.target.value)}
                placeholder="e.g. Auntie Em — congrats on the new baby!"
                multiline
                minRows={3}
                fullWidth
              />
              {err && <Alert severity="error">{err}</Alert>}
              <Button
                variant="contained"
                size="large"
                onClick={() => requestM.mutate()}
                disabled={requestM.isPending}
                sx={{ py: 1.4 }}
              >
                {requestM.isPending ? <CircularProgress size={20} /> : 'Request access'}
              </Button>
              <Typography variant="caption" color="text.secondary" textAlign="center">
                Nothing on the registry is shared with anyone until {parentsLabel} approve you.
              </Typography>
            </Stack>
          )}

          {status === 'pending' && (
            <Stack spacing={1.5} alignItems="center">
              <Chip label="Awaiting approval" color="secondary" sx={{ fontWeight: 600 }} />
              <Typography variant="caption" color="text.secondary" textAlign="center" sx={{ maxWidth: 320 }}>
                You can close this page and come back later — the link will work the moment access is granted.
              </Typography>
            </Stack>
          )}
        </CardContent>
      </Card>
    </Container>
  );
}
