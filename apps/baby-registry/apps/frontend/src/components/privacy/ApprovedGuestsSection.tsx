import { useState } from 'react';
import {
  Alert,
  Box,
  Button,
  Chip,
  CircularProgress,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  IconButton,
  Stack,
  TextField,
  Tooltip,
  Typography,
} from '@mui/material';
import BlockIcon from '@mui/icons-material/Block';
import UndoIcon from '@mui/icons-material/Undo';
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import LinkIcon from '@mui/icons-material/Link';
import DeleteIcon from '@mui/icons-material/DeleteOutline';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { approvedGuests, type ApprovedGuest, type GuestAccessLevel } from '../../api';
import ShareLinkAlert, { type ShareLink } from './ShareLinkAlert';

type GuestAction = 'revoke' | 'block' | 'reactivate' | 'remove';

export default function ApprovedGuestsSection({ registryId }: { registryId: string }) {
  const qc = useQueryClient();
  const guestsQ = useQuery({
    queryKey: ['approvedGuests', registryId],
    queryFn: () => approvedGuests.list(registryId),
  });
  const [email, setEmail] = useState('');
  const [name, setName] = useState('');
  const [accessLevel] = useState<GuestAccessLevel>('ViewShippingAddress');
  const [err, setErr] = useState<string | null>(null);
  const [confirmGuestAction, setConfirmGuestAction] = useState<
    | { action: GuestAction; guest: ApprovedGuest }
    | null
  >(null);
  const [issuedLink, setIssuedLink] = useState<ShareLink | null>(null);

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
  const revokeM = useMutation({ mutationFn: (id: string) => approvedGuests.revoke(id), onSuccess: invalidate });
  const blockM = useMutation({ mutationFn: (id: string) => approvedGuests.block(id), onSuccess: invalidate });
  const reactivateM = useMutation({ mutationFn: (id: string) => approvedGuests.reactivate(id), onSuccess: invalidate });
  const removeM = useMutation({ mutationFn: (id: string) => approvedGuests.remove(id), onSuccess: invalidate });
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
    <GuestRow
      key={g.id}
      guest={g}
      issuing={issueLinkM.isPending}
      onIssueLink={() => issueLinkM.mutate(g)}
      onAction={(action) => setConfirmGuestAction({ action, guest: g })}
    />
  );

  return (
    <Box>
      {err && (
        <Alert severity="error" sx={{ mb: 2 }}>
          {err}
        </Alert>
      )}

      {issuedLink && (
        <ShareLinkAlert
          link={issuedLink}
          onClose={() => setIssuedLink(null)}
          expiresSuffix="Anyone with this link can see your address."
        />
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

function GuestRow({
  guest: g,
  issuing,
  onIssueLink,
  onAction,
}: {
  guest: ApprovedGuest;
  issuing: boolean;
  onIssueLink: () => void;
  onAction: (action: GuestAction) => void;
}) {
  return (
    <Stack
      direction="row"
      alignItems="center"
      spacing={1}
      sx={{ p: 1, borderRadius: 1, bgcolor: 'action.hover' }}
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
          color={g.status === 'Active' ? 'success' : g.status === 'Blocked' ? 'error' : 'default'}
          variant={g.status === 'Active' ? 'filled' : 'outlined'}
        />
      )}
      {g.status === 'Active' && (
        <>
          <Tooltip title="Generate share link">
            <IconButton size="small" onClick={onIssueLink} disabled={issuing}>
              <LinkIcon fontSize="small" />
            </IconButton>
          </Tooltip>
          <Tooltip title="Revoke">
            <IconButton size="small" onClick={() => onAction('revoke')}>
              <UndoIcon fontSize="small" />
            </IconButton>
          </Tooltip>
          <Tooltip title="Block">
            <IconButton size="small" onClick={() => onAction('block')}>
              <BlockIcon fontSize="small" />
            </IconButton>
          </Tooltip>
        </>
      )}
      {g.status === 'Pending' && (
        <>
          <Tooltip title="Approve">
            <IconButton size="small" color="success" onClick={() => onAction('reactivate')}>
              <CheckCircleIcon fontSize="small" />
            </IconButton>
          </Tooltip>
          <Tooltip title="Decline">
            <IconButton size="small" onClick={() => onAction('block')}>
              <BlockIcon fontSize="small" />
            </IconButton>
          </Tooltip>
        </>
      )}
      {g.status !== 'Active' && g.status !== 'Pending' && (
        <Tooltip title="Reactivate">
          <IconButton size="small" onClick={() => onAction('reactivate')}>
            <CheckCircleIcon fontSize="small" />
          </IconButton>
        </Tooltip>
      )}
      <Tooltip title="Remove guest">
        <IconButton size="small" color="error" onClick={() => onAction('remove')}>
          <DeleteIcon fontSize="small" />
        </IconButton>
      </Tooltip>
    </Stack>
  );
}
