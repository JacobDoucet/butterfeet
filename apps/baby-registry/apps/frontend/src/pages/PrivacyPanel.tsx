import { useEffect, useState } from 'react';
import {
  Accordion,
  AccordionSummary,
  AccordionDetails,
  Typography,
  Stack,
  TextField,
  ToggleButton,
  ToggleButtonGroup,
  Button,
  Alert,
  Box,
  Divider,
  Chip,
  IconButton,
  Tooltip,
  CircularProgress,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
} from '@mui/material';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import BlockIcon from '@mui/icons-material/Block';
import UndoIcon from '@mui/icons-material/Undo';
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import LinkIcon from '@mui/icons-material/Link';
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import DeleteIcon from '@mui/icons-material/DeleteOutline';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import {
  registries,
  approvedGuests,
  addressRequests,
  type Registry,
  type AddressAccessMode,
  type GuestAccessLevel,
  type ApprovedGuest,
  type AddressRequest,
} from '../api';

const MODES: { value: AddressAccessMode; label: string; help: string }[] = [
  {
    value: 'RequestApproval',
    label: 'Request approval',
    help: 'Buyers can request your shipping address. You approve each request.',
  },
  {
    value: 'ApprovedGuestsOnly',
    label: 'Approved guests only',
    help: 'Only people on your approved guest list can see the address. Everyone else is silently ignored.',
  },
  {
    value: 'Disabled',
    label: 'Disabled',
    help: 'No one can request or view your shipping address.',
  },
];

type PrivacyPanelSection = 'all' | 'shipping' | 'access';

export default function PrivacyPanel({
  reg,
  section = 'all',
}: {
  reg: Registry;
  section?: PrivacyPanelSection;
}) {
  const qc = useQueryClient();
  const [mode, setMode] = useState<AddressAccessMode>(reg.addressAccessMode ?? 'RequestApproval');
  const [requireGuestApproval, setRequireGuestApproval] = useState<boolean>(!(reg.allowOpenAccess ?? false));
  const [recipientName, setRecipientName] = useState(reg.shippingRecipientName ?? '');
  const [line1, setLine1] = useState(reg.shippingLine1 ?? '');
  const [line2, setLine2] = useState(reg.shippingLine2 ?? '');
  const [city, setCity] = useState(reg.shippingCity ?? '');
  const [region, setRegion] = useState(reg.shippingRegion ?? '');
  const [postalCode, setPostalCode] = useState(reg.shippingPostalCode ?? '');
  const [country, setCountry] = useState(reg.shippingCountry ?? '');
  const [deliveryNotes, setDeliveryNotes] = useState(reg.shippingDeliveryNotes ?? '');
  const [saved, setSaved] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    setMode(reg.addressAccessMode ?? 'RequestApproval');
    setRequireGuestApproval(!(reg.allowOpenAccess ?? false));
    setRecipientName(reg.shippingRecipientName ?? '');
    setLine1(reg.shippingLine1 ?? '');
    setLine2(reg.shippingLine2 ?? '');
    setCity(reg.shippingCity ?? '');
    setRegion(reg.shippingRegion ?? '');
    setPostalCode(reg.shippingPostalCode ?? '');
    setCountry(reg.shippingCountry ?? '');
    setDeliveryNotes(reg.shippingDeliveryNotes ?? '');
  }, [reg.id]);

  const saveM = useMutation({
    mutationFn: () =>
      registries.update(reg.id, {
        addressAccessMode: mode,
        shippingRecipientName: recipientName,
        shippingLine1: line1,
        shippingLine2: line2,
        shippingCity: city,
        shippingRegion: region,
        shippingPostalCode: postalCode,
        shippingCountry: country,
        shippingDeliveryNotes: deliveryNotes,
        shippingPolicyVersion: (reg.shippingPolicyVersion ?? 0) + 1,
      }),
    onSuccess: () => {
      setSaved(true);
      setError(null);
      qc.invalidateQueries({ queryKey: ['registries'] });
      window.setTimeout(() => setSaved(false), 2000);
    },
    onError: (err) => setError((err as Error).message),
  });

  const saveAccessM = useMutation({
    mutationFn: (next: boolean) =>
      registries.update(reg.id, { allowOpenAccess: !next }),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['registries'] });
    },
    onError: (err) => setError((err as Error).message),
  });

  const shippingHidden = mode === 'Disabled';
  const shippingContent = (
    <Stack spacing={3}>
      <Box
        sx={{
          p: 2,
          borderRadius: 2,
          bgcolor: shippingHidden ? 'background.default' : 'primary.main',
          border: '1px solid',
          borderColor: shippingHidden ? 'divider' : 'primary.main',
        }}
      >
        <Stack direction="row" alignItems="flex-start" spacing={1.5}>
          <LockOutlinedIcon
            fontSize="small"
            sx={{ color: shippingHidden ? 'text.secondary' : '#fff', mt: 0.3 }}
          />
          <Box sx={{ flex: 1 }}>
            <Typography variant="subtitle2" sx={{ fontWeight: 600, color: shippingHidden ? undefined : '#fff' }}>
              Shipping address visibility
            </Typography>
            <Typography variant="caption" sx={{ color: shippingHidden ? 'text.secondary' : 'rgba(255,255,255,0.85)' }}>
              {shippingHidden
                ? 'Hidden — no one can see your shipping address. Approved guests will need to message you for it.'
                : 'Shown — approved guests can see your shipping address from the registry.'}
            </Typography>
          </Box>
          <ToggleButtonGroup
            exclusive
            size="small"
            value={shippingHidden ? 'hide' : 'show'}
            onChange={(_, v) => v && setMode(v === 'show' ? 'ApprovedGuestsOnly' : 'Disabled')}
            sx={{
              '& .MuiToggleButton-root': {
                color: '#fff',
                borderColor: 'rgba(255,255,255,0.5)',
                '&.Mui-selected': {
                  color: 'primary.main',
                  bgcolor: '#fff',
                  '&:hover': { bgcolor: '#fff' },
                },
                '&:hover': { bgcolor: 'rgba(255,255,255,0.12)' },
              },
            }}
          >
            <ToggleButton value="show" sx={{ textTransform: 'none', px: 2 }}>Show</ToggleButton>
            <ToggleButton value="hide" sx={{ textTransform: 'none', px: 2 }}>Hide</ToggleButton>
          </ToggleButtonGroup>
        </Stack>
      </Box>

      <Box>
        <Typography variant="overline" color="text.secondary">Shipping address</Typography>
        <Stack spacing={2} sx={{ mt: 1, opacity: shippingHidden ? 0.55 : 1 }}>
          <TextField label="Recipient name" value={recipientName} onChange={(e) => setRecipientName(e.target.value)} disabled={shippingHidden} />
          <TextField label="Address line 1" value={line1} onChange={(e) => setLine1(e.target.value)} disabled={shippingHidden} />
          <TextField label="Address line 2" value={line2} onChange={(e) => setLine2(e.target.value)} disabled={shippingHidden} />
          <Stack direction={{ xs: 'column', sm: 'row' }} spacing={2}>
            <TextField label="City" value={city} onChange={(e) => setCity(e.target.value)} sx={{ flex: 1 }} disabled={shippingHidden} />
            <TextField label="Region / State" value={region} onChange={(e) => setRegion(e.target.value)} sx={{ flex: 1 }} disabled={shippingHidden} />
          </Stack>
          <Stack direction={{ xs: 'column', sm: 'row' }} spacing={2}>
            <TextField label="Postal code" value={postalCode} onChange={(e) => setPostalCode(e.target.value)} sx={{ flex: 1 }} disabled={shippingHidden} />
            <TextField label="Country" value={country} onChange={(e) => setCountry(e.target.value)} sx={{ flex: 1 }} disabled={shippingHidden} />
          </Stack>
          <TextField
            label="Delivery notes (optional)"
            value={deliveryNotes}
            onChange={(e) => setDeliveryNotes(e.target.value)}
            multiline
            minRows={2}
            helperText="Buzzer code, where to leave parcels, etc."
            disabled={shippingHidden}
          />
        </Stack>
      </Box>

      {error && <Alert severity="error">{error}</Alert>}

      <Stack direction="row" justifyContent="flex-end">
        <Button variant="contained" onClick={() => saveM.mutate()} disabled={saveM.isPending} sx={{ color: '#fff' }}>
          Save shipping settings
        </Button>
      </Stack>
    </Stack>
  );

  const accessContent = (
    <Stack spacing={3}>
      <Box
        sx={{
          p: 2,
          borderRadius: 2,
          bgcolor: requireGuestApproval ? 'primary.main' : 'background.default',
          border: '1px solid',
          borderColor: requireGuestApproval ? 'primary.main' : 'divider',
        }}
      >
        <Stack direction="row" alignItems="flex-start" spacing={1.5}>
          <LockOutlinedIcon
            fontSize="small"
            sx={{ color: requireGuestApproval ? '#fff' : 'text.secondary', mt: 0.3 }}
          />
          <Box sx={{ flex: 1 }}>
            <Typography variant="subtitle2" sx={{ fontWeight: 600, color: requireGuestApproval ? '#fff' : undefined }}>
              Require approval to view your registry
            </Typography>
            <Typography variant="caption" sx={{ color: requireGuestApproval ? 'rgba(255,255,255,0.85)' : 'text.secondary' }}>
              When on, guests must request access from you before they can see any of your registry. You decide who gets in.
            </Typography>
          </Box>
          <ToggleButtonGroup
            exclusive
            size="small"
            value={requireGuestApproval ? 'on' : 'off'}
            disabled={saveAccessM.isPending}
            onChange={(_, v) => {
              if (v == null) return;
              const next = v === 'on';
              setRequireGuestApproval(next);
              saveAccessM.mutate(next);
            }}
            sx={{
              '& .MuiToggleButton-root': {
                color: requireGuestApproval ? '#fff' : undefined,
                borderColor: requireGuestApproval ? 'rgba(255,255,255,0.5)' : undefined,
                '&.Mui-selected': requireGuestApproval
                  ? {
                      color: 'primary.main',
                      bgcolor: '#fff',
                      '&:hover': { bgcolor: '#fff' },
                    }
                  : undefined,
                '&:hover': requireGuestApproval ? { bgcolor: 'rgba(255,255,255,0.12)' } : undefined,
              },
            }}
          >
            <ToggleButton value="off" sx={{ textTransform: 'none', px: 2 }}>Off</ToggleButton>
            <ToggleButton value="on" sx={{ textTransform: 'none', px: 2 }}>On</ToggleButton>
          </ToggleButtonGroup>
        </Stack>
      </Box>
      <ApprovedGuestsSection registryId={reg.id} />
    </Stack>
  );

  if (section === 'shipping') {
    return shippingContent;
  }

  if (section === 'access') {
    return accessContent;
  }

  return (
    <Accordion sx={{ mb: 3 }}>
      <AccordionSummary expandIcon={<ExpandMoreIcon />}>
        <Stack direction="row" alignItems="center" spacing={1}>
          <LockOutlinedIcon fontSize="small" />
          <Typography variant="subtitle1" sx={{ fontWeight: 600 }}>Private shipping</Typography>
          <Typography variant="body2" color="text.secondary">
            ({MODES.find((m) => m.value === mode)?.label})
          </Typography>
        </Stack>
      </AccordionSummary>
      <AccordionDetails>
        <Stack spacing={3}>
          {shippingContent}

          <Divider />

          {accessContent}
        </Stack>
      </AccordionDetails>
    </Accordion>
  );
}

function ApprovedGuestsSection({ registryId }: { registryId: string }) {
  const qc = useQueryClient();
  const guestsQ = useQuery({
    queryKey: ['approvedGuests', registryId],
    queryFn: () => approvedGuests.list(registryId),
  });
  const [email, setEmail] = useState('');
  const [name, setName] = useState('');
  const [accessLevel, setAccessLevel] = useState<GuestAccessLevel>('ViewShippingAddress');
  const [err, setErr] = useState<string | null>(null);
  const [confirmGuestAction, setConfirmGuestAction] = useState<
    | { action: 'revoke' | 'block' | 'reactivate' | 'remove'; guest: ApprovedGuest }
    | null
  >(null);

  const invalidate = () => qc.invalidateQueries({ queryKey: ['approvedGuests', registryId] });

  const addM = useMutation({
    mutationFn: () =>
      approvedGuests.add(registryId, { email: email.trim(), name: name.trim() || undefined, accessLevel }),
    onSuccess: () => {
      setEmail('');
      setName('');
      setErr(null);
      invalidate();
    },
    onError: (e) => setErr((e as Error).message),
  });
  const revokeM = useMutation({
    mutationFn: (id: string) => approvedGuests.revoke(id),
    onSuccess: invalidate,
  });
  const blockM = useMutation({
    mutationFn: (id: string) => approvedGuests.block(id),
    onSuccess: invalidate,
  });
  const reactivateM = useMutation({
    mutationFn: (id: string) => approvedGuests.reactivate(id),
    onSuccess: invalidate,
  });
  const removeM = useMutation({
    mutationFn: (id: string) => approvedGuests.remove(id),
    onSuccess: invalidate,
  });
  const [issuedLink, setIssuedLink] = useState<{ email: string; url: string; expiresAt: string } | null>(null);
  const issueLinkM = useMutation({
    mutationFn: (g: ApprovedGuest) => approvedGuests.issueLink(g.id).then((r) => ({ g, r })),
    onSuccess: ({ g, r }) => {
      const url = `${window.location.origin}/ship#tok=${r.token}`;
      setIssuedLink({ email: g.email, url, expiresAt: r.expiresAt });
    },
    onError: (e) => setErr((e as Error).message),
  });

  const guests: ApprovedGuest[] = guestsQ.data ?? [];

  const confirmGuestActionText =
    confirmGuestAction?.action === 'revoke'
      ? 'Revoke this guest? They will lose active access until reactivated.'
      : confirmGuestAction?.action === 'block'
      ? confirmGuestAction.guest.status === 'Pending'
        ? 'Decline this access request? They will not be allowed in until you reactivate them.'
        : 'Block this guest? This prevents access until manually reactivated.'
      : confirmGuestAction?.action === 'reactivate'
      ? confirmGuestAction.guest.status === 'Pending'
        ? 'Approve this guest? They will be able to view your registry immediately.'
        : 'Reactivate this guest? They will regain access immediately.'
      : confirmGuestAction?.action === 'remove'
      ? 'Remove this guest entry permanently?'
      : '';

  const runGuestAction = () => {
    if (!confirmGuestAction) return;
    const id = confirmGuestAction.guest.id;
    if (confirmGuestAction.action === 'revoke') revokeM.mutate(id);
    if (confirmGuestAction.action === 'block') blockM.mutate(id);
    if (confirmGuestAction.action === 'reactivate') reactivateM.mutate(id);
    if (confirmGuestAction.action === 'remove') removeM.mutate(id);
    setConfirmGuestAction(null);
  };

  const pendingGuests = guests.filter((g) => g.status === 'Pending');
  const otherGuests = guests.filter((g) => g.status !== 'Pending');

  const renderGuestRow = (g: ApprovedGuest) => (
    <Stack
      key={g.id}
      direction="row"
      alignItems="center"
      spacing={1}
      sx={{
        p: 1,
        borderRadius: 1,
        bgcolor: 'action.hover',
      }}
    >
      <Box sx={{ flex: 1, minWidth: 0 }}>
        <Typography variant="body2" sx={{ fontWeight: 500 }} noWrap>
          {g.name ? `${g.name} · ${g.email}` : g.email}
        </Typography>
      </Box>
      {g.status !== 'Pending' && (
        <Chip
          size="small"
          label={g.status}
          color={
            g.status === 'Active'
              ? 'success'
              : g.status === 'Blocked'
              ? 'error'
              : 'default'
          }
          variant={g.status === 'Active' ? 'filled' : 'outlined'}
        />
      )}
      {g.status === 'Active' && (
        <>
          <Tooltip title="Generate share link">
            <IconButton
              size="small"
              onClick={() => issueLinkM.mutate(g)}
              disabled={issueLinkM.isPending}
            >
              <LinkIcon fontSize="small" />
            </IconButton>
          </Tooltip>
          <Tooltip title="Revoke">
            <IconButton size="small" onClick={() => setConfirmGuestAction({ action: 'revoke', guest: g })}>
              <UndoIcon fontSize="small" />
            </IconButton>
          </Tooltip>
          <Tooltip title="Block">
            <IconButton size="small" onClick={() => setConfirmGuestAction({ action: 'block', guest: g })}>
              <BlockIcon fontSize="small" />
            </IconButton>
          </Tooltip>
        </>
      )}
      {g.status === 'Pending' && (
        <>
          <Tooltip title="Approve">
            <IconButton
              size="small"
              color="success"
              onClick={() => setConfirmGuestAction({ action: 'reactivate', guest: g })}
            >
              <CheckCircleIcon fontSize="small" />
            </IconButton>
          </Tooltip>
          <Tooltip title="Decline">
            <IconButton
              size="small"
              onClick={() => setConfirmGuestAction({ action: 'block', guest: g })}
            >
              <BlockIcon fontSize="small" />
            </IconButton>
          </Tooltip>
        </>
      )}
      {g.status !== 'Active' && g.status !== 'Pending' && (
        <Tooltip title="Reactivate">
          <IconButton
            size="small"
            onClick={() => setConfirmGuestAction({ action: 'reactivate', guest: g })}
          >
            <CheckCircleIcon fontSize="small" />
          </IconButton>
        </Tooltip>
      )}
      <Tooltip title="Remove guest">
        <IconButton size="small" color="error" onClick={() => setConfirmGuestAction({ action: 'remove', guest: g })}>
          <DeleteIcon fontSize="small" />
        </IconButton>
      </Tooltip>
    </Stack>
  );

  return (
    <Box>
      {err && (
        <Alert severity="error" sx={{ mb: 2 }}>
          {err}
        </Alert>
      )}

      {issuedLink && (
        <Alert
          severity="success"
          sx={{ mb: 2, '& .MuiAlert-message': { width: '100%' } }}
          onClose={() => setIssuedLink(null)}
        >
          <Typography variant="body2" sx={{ fontWeight: 500, mb: 0.5 }}>
            Share this link with {issuedLink.email}
          </Typography>
          <Stack direction="row" spacing={1} alignItems="center">
            <TextField
              size="small"
              value={issuedLink.url}
              fullWidth
              InputProps={{ readOnly: true }}
              onFocus={(e) => e.target.select()}
            />
            <Tooltip title="Copy">
              <IconButton size="small" onClick={() => navigator.clipboard.writeText(issuedLink.url)}>
                <ContentCopyIcon fontSize="small" />
              </IconButton>
            </Tooltip>
          </Stack>
          <Typography variant="caption" color="text.secondary">
            Expires {new Date(issuedLink.expiresAt).toLocaleString()}. Anyone with this link can see your address.
          </Typography>
        </Alert>
      )}

      {guestsQ.isLoading ? (
        <CircularProgress size={20} />
      ) : (
        <Stack spacing={3}>
          <Box>
            <Typography variant="overline" color="text.secondary">
              Pending requests
            </Typography>
            <Typography variant="caption" color="text.secondary" sx={{ display: 'block', mb: 1.5 }}>
              People who have asked to view your registry. Approve to let them in.
            </Typography>
            {pendingGuests.length === 0 ? (
              <Typography variant="body2" color="text.secondary">
                No pending requests.
              </Typography>
            ) : (
              <Stack spacing={1}>{pendingGuests.map(renderGuestRow)}</Stack>
            )}
          </Box>

          <Box>
            <Typography variant="overline" color="text.secondary">
              Approved guests
            </Typography>
            <Typography variant="caption" color="text.secondary" sx={{ display: 'block', mb: 1.5 }}>
              People who can see your shipping address without asking.
            </Typography>

            <Stack
              direction={{ xs: 'column', sm: 'row' }}
              spacing={1}
              sx={{ mb: 2 }}
              alignItems={{ sm: 'center' }}
            >
              <TextField
                size="small"
                label="Email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                sx={{ flex: 2 }}
              />
              <TextField
                size="small"
                label="Name (optional)"
                value={name}
                onChange={(e) => setName(e.target.value)}
                sx={{ flex: 1 }}
              />
              <Button
                variant="outlined"
                onClick={() => addM.mutate()}
                disabled={!email.trim() || addM.isPending}
              >
                Add
              </Button>
            </Stack>

            {otherGuests.length === 0 ? (
              <Typography variant="body2" color="text.secondary">
                No approved guests yet.
              </Typography>
            ) : (
              <Stack spacing={1}>{otherGuests.map(renderGuestRow)}</Stack>
            )}
          </Box>
        </Stack>
      )}

      <Dialog open={!!confirmGuestAction} onClose={() => setConfirmGuestAction(null)} maxWidth="xs" fullWidth>
        <DialogTitle>Confirm action</DialogTitle>
        <DialogContent>
          <Typography variant="body2" color="text.secondary">
            {confirmGuestActionText}
          </Typography>
          {confirmGuestAction?.guest && (
            <Typography variant="caption" color="text.secondary" sx={{ display: 'block', mt: 1 }}>
              {confirmGuestAction.guest.name
                ? `${confirmGuestAction.guest.name} · ${confirmGuestAction.guest.email}`
                : confirmGuestAction.guest.email}
            </Typography>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setConfirmGuestAction(null)}>Cancel</Button>
          <Button variant="contained" color="error" onClick={runGuestAction}>
            Confirm
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
}

function AddressRequestsSection({ registryId }: { registryId: string }) {
  const qc = useQueryClient();
  const reqsQ = useQuery({
    queryKey: ['addressRequests', registryId],
    queryFn: () => addressRequests.list(registryId),
  });
  const invalidate = () => {
    qc.invalidateQueries({ queryKey: ['addressRequests', registryId] });
    qc.invalidateQueries({ queryKey: ['approvedGuests', registryId] });
  };
  const [issuedLink, setIssuedLink] = useState<{ email: string; url: string; expiresAt: string } | null>(null);
  const [confirmRequestAction, setConfirmRequestAction] = useState<
    | { action: 'reject' | 'block'; req: AddressRequest }
    | null
  >(null);
  const approveM = useMutation({
    mutationFn: ({ id, permanent }: { id: string; permanent: boolean }) =>
      addressRequests.approve(id, { permanent }),
    onSuccess: (req) => {
      invalidate();
      if (req.token && req.tokenExpiresAt) {
        setIssuedLink({
          email: req.email,
          url: `${window.location.origin}/ship#tok=${req.token}`,
          expiresAt: req.tokenExpiresAt,
        });
      }
    },
  });
  const rejectM = useMutation({
    mutationFn: (id: string) => addressRequests.reject(id),
    onSuccess: invalidate,
  });
  const blockReqM = useMutation({
    mutationFn: (id: string) => addressRequests.block(id),
    onSuccess: invalidate,
  });

  const requests: AddressRequest[] = reqsQ.data ?? [];
  const pending = requests.filter((r) => r.status === 'Pending');
  const decided = requests.filter((r) => r.status !== 'Pending');

  const runRequestAction = () => {
    if (!confirmRequestAction) return;
    if (confirmRequestAction.action === 'reject') rejectM.mutate(confirmRequestAction.req.id);
    if (confirmRequestAction.action === 'block') blockReqM.mutate(confirmRequestAction.req.id);
    setConfirmRequestAction(null);
  };

  return (
    <Box>
      <Typography variant="overline" color="text.secondary">
        Address requests
      </Typography>
      <Typography variant="caption" color="text.secondary" sx={{ display: 'block', mb: 2 }}>
        Buyers who asked to see your shipping address. Approve once or permanently (adds them to approved guests).
      </Typography>

      {issuedLink && (
        <Alert
          severity="success"
          sx={{ mb: 2, '& .MuiAlert-message': { width: '100%' } }}
          onClose={() => setIssuedLink(null)}
        >
          <Typography variant="body2" sx={{ fontWeight: 500, mb: 0.5 }}>
            Share this link with {issuedLink.email}
          </Typography>
          <Stack direction="row" spacing={1} alignItems="center">
            <TextField
              size="small"
              value={issuedLink.url}
              fullWidth
              InputProps={{ readOnly: true }}
              onFocus={(e) => e.target.select()}
            />
            <Tooltip title="Copy">
              <IconButton size="small" onClick={() => navigator.clipboard.writeText(issuedLink.url)}>
                <ContentCopyIcon fontSize="small" />
              </IconButton>
            </Tooltip>
          </Stack>
          <Typography variant="caption" color="text.secondary">
            Expires {new Date(issuedLink.expiresAt).toLocaleString()}.
          </Typography>
        </Alert>
      )}
      {reqsQ.isLoading ? (
        <CircularProgress size={20} />
      ) : requests.length === 0 ? (
        <Typography variant="body2" color="text.secondary">
          No requests yet.
        </Typography>
      ) : (
        <Stack spacing={1}>
          {[...pending, ...decided].map((req) => (
            <Stack
              key={req.id}
              direction={{ xs: 'column', sm: 'row' }}
              alignItems={{ sm: 'center' }}
              spacing={1}
              sx={{ p: 1, borderRadius: 1, bgcolor: 'action.hover' }}
            >
              <Box sx={{ flex: 1, minWidth: 0 }}>
                <Typography variant="body2" sx={{ fontWeight: 500 }} noWrap>
                  {req.name ? `${req.name} · ${req.email}` : req.email}
                </Typography>
                {req.note && (
                  <Typography variant="caption" color="text.secondary" sx={{ display: 'block' }}>
                    “{req.note}”
                  </Typography>
                )}
              </Box>
              <Chip
                size="small"
                label={req.status}
                color={
                  req.status === 'Approved' || req.status === 'AutoApproved'
                    ? 'success'
                    : req.status === 'Blocked' || req.status === 'Rejected'
                    ? 'default'
                    : 'warning'
                }
                variant={req.status === 'Pending' ? 'filled' : 'outlined'}
              />
              {req.status === 'Pending' && (
                <Stack direction="row" spacing={1}>
                  <Button
                    size="small"
                    variant="outlined"
                    onClick={() => approveM.mutate({ id: req.id, permanent: false })}
                    disabled={approveM.isPending}
                  >
                    Approve once
                  </Button>
                  <Button
                    size="small"
                    variant="contained"
                    onClick={() => approveM.mutate({ id: req.id, permanent: true })}
                    disabled={approveM.isPending}
                  >
                    Approve permanently
                  </Button>
                  <Button
                    size="small"
                    color="inherit"
                    onClick={() => setConfirmRequestAction({ action: 'reject', req })}
                    disabled={rejectM.isPending}
                  >
                    Reject
                  </Button>
                  <Button
                    size="small"
                    color="error"
                    onClick={() => setConfirmRequestAction({ action: 'block', req })}
                    disabled={blockReqM.isPending}
                  >
                    Block
                  </Button>
                </Stack>
              )}
            </Stack>
          ))}
        </Stack>
      )}

      <Dialog open={!!confirmRequestAction} onClose={() => setConfirmRequestAction(null)} maxWidth="xs" fullWidth>
        <DialogTitle>Confirm action</DialogTitle>
        <DialogContent>
          <Typography variant="body2" color="text.secondary">
            {confirmRequestAction?.action === 'block'
              ? 'Block this requester? They will not be able to access the shipping flow until manually reactivated.'
              : 'Reject this request?'}
          </Typography>
          {confirmRequestAction?.req && (
            <Typography variant="caption" color="text.secondary" sx={{ display: 'block', mt: 1 }}>
              {confirmRequestAction.req.name
                ? `${confirmRequestAction.req.name} · ${confirmRequestAction.req.email}`
                : confirmRequestAction.req.email}
            </Typography>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setConfirmRequestAction(null)}>Cancel</Button>
          <Button variant="contained" color="error" onClick={runRequestAction}>
            Confirm
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
}
