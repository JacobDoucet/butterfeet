import {
  Stack,
  Typography,
  Button,
  Alert,
  Box,
  Card,
  CardContent,
  Chip,
  CircularProgress,
  Link,
} from '@mui/material';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { payments, items, formatPriceCents, type Registry, type Cart } from '../../api';

type Mode = 'held' | 'to-review' | 'completed';

const HEADINGS: Record<Mode, string> = {
  held: 'Held carts',
  'to-review': 'To review',
  completed: 'Completed',
};

// CartsPanel renders the owner's cart surfaces. In "held" mode it lists carts a
// guest has started but not yet marked as paid (Pending), with a delete action
// to free up abandoned/stuck holds. In "to-review" mode it lists carts a guest
// has marked as paid (AwaitingConfirmation) with approve/reject actions. In
// "completed" mode it lists carts the owner has already confirmed.
export default function CartsPanel({ reg, mode }: { reg: Registry; mode: Mode }) {
  const qc = useQueryClient();
  // Each surface maps to a cart status: Pending = guest mid-flow (held);
  // AwaitingConfirmation = guest says they've paid; Completed = owner confirmed.
  const status = mode === 'held' ? 'Pending' : mode === 'to-review' ? 'AwaitingConfirmation' : 'Completed';

  const cartsQ = useQuery({
    queryKey: ['carts', reg.id, status],
    queryFn: () => payments.listCarts(reg.id, status),
  });

  // Cart item snapshots only carry the title, not a product URL. Join against
  // the registry's items by itemId so each bullet can link out to the product
  // (preferring the affiliate URL when present).
  const itemsQ = useQuery({
    queryKey: ['items', reg.id],
    queryFn: () => items.listForRegistry(reg.id),
  });
  const urlByItemId: Record<string, string> = {};
  for (const it of itemsQ.data?.data ?? []) {
    const url = it.affiliateUrl || it.productUrl;
    if (url) urlByItemId[it.id] = url;
  }

  const invalidate = () => {
    qc.invalidateQueries({ queryKey: ['carts', reg.id] });
    qc.invalidateQueries({ queryKey: ['reservations', reg.id] });
    qc.invalidateQueries({ queryKey: ['items', reg.id] });
  };

  const approveM = useMutation({
    mutationFn: (id: string) => payments.approve(id),
    onSuccess: invalidate,
  });
  const rejectM = useMutation({
    mutationFn: (id: string) => payments.reject(id),
    onSuccess: invalidate,
  });
  const deleteM = useMutation({
    mutationFn: (id: string) => payments.deleteCart(id),
    onSuccess: invalidate,
  });

  const carts = cartsQ.data ?? [];
  const busy = approveM.isPending || rejectM.isPending || deleteM.isPending;

  const emptyText =
    mode === 'held'
      ? 'No carts are currently being held.'
      : mode === 'to-review'
        ? 'No payments are waiting for your review.'
        : 'No confirmed payments yet.';

  return (
    <Stack spacing={2}>
      <Box>
        <Typography variant="h6">{HEADINGS[mode]}</Typography>
        {mode === 'held' ? (
          <Typography variant="body2" color="text.secondary" sx={{ mt: 0.5 }}>
            Gifts guests are holding while they pay, but haven't yet marked as sent. Deleting a cart
            frees its gifts so other guests can claim them.
          </Typography>
        ) : mode === 'to-review' ? (
          <Alert severity="warning" sx={{ mt: 1 }}>
            Only approve a payment after you have confirmed the money has arrived. Rejecting frees
            the gifts so other guests can claim them.
          </Alert>
        ) : (
          <Typography variant="body2" color="text.secondary" sx={{ mt: 0.5 }}>
            Payments you've confirmed. The matching gifts are marked as purchased.
          </Typography>
        )}
      </Box>

      {cartsQ.isLoading ? (
        <Box sx={{ py: 3, textAlign: 'center' }}>
          <CircularProgress size={24} />
        </Box>
      ) : carts.length === 0 ? (
        <Typography variant="body2" color="text.secondary">
          {emptyText}
        </Typography>
      ) : (
        <Stack spacing={1.5}>
          {carts.map((c) => (
            <CartCard
              key={c.id}
              cart={c}
              mode={mode}
              busy={busy}
              urlByItemId={urlByItemId}
              onApprove={() => approveM.mutate(c.id)}
              onReject={() => rejectM.mutate(c.id)}
              onDelete={() => deleteM.mutate(c.id)}
            />
          ))}
        </Stack>
      )}
    </Stack>
  );
}

function CartCard({
  cart,
  mode,
  busy,
  urlByItemId,
  onApprove,
  onReject,
  onDelete,
}: {
  cart: Cart;
  mode: Mode;
  busy: boolean;
  urlByItemId: Record<string, string>;
  onApprove: () => void;
  onReject: () => void;
  onDelete: () => void;
}) {
  return (
    <Card variant="outlined">
      <CardContent>
        <Stack direction="row" justifyContent="space-between" alignItems="flex-start" spacing={1}>
          <Box>
            <Typography variant="subtitle1" fontWeight={600}>
              {formatPriceCents(cart.amountCents, cart.currency)} ·{' '}
              {cart.methodDisplayName || cart.methodType}
            </Typography>
            <Typography variant="body2" color="text.secondary">
              {cart.contributorName || 'Guest'}
              {cart.contributorEmail ? ` · ${cart.contributorEmail}` : ''}
            </Typography>
          </Box>
          <Stack spacing={0.5} alignItems="flex-end">
            <Chip size="small" label={cart.referenceCode} />
            {mode === 'held' && (
              <Chip size="small" variant="outlined" label="Awaiting payment" />
            )}
            {mode === 'to-review' &&
              (cart.status === 'AwaitingConfirmation' ? (
                <Chip size="small" color="warning" label="Says they've paid" />
              ) : (
                <Chip size="small" variant="outlined" label="Awaiting payment" />
              ))}
          </Stack>
        </Stack>

        {cart.items.length > 0 && (
          <Box
            component="ul"
            sx={{ mt: 1.5, mb: 0, pl: 3, display: 'flex', flexDirection: 'column', gap: 0.25 }}
          >
            {cart.items.map((it) => {
              const url = urlByItemId[it.itemId];
              const priceSuffix = it.priceCents
                ? ` — ${formatPriceCents(it.priceCents, it.currency)}`
                : '';
              return (
                <Typography key={it.reservationId} component="li" variant="body2">
                  {it.quantity}×{' '}
                  {url ? (
                    <Link href={url} target="_blank" rel="noopener noreferrer" underline="hover">
                      {it.title}
                    </Link>
                  ) : (
                    it.title
                  )}
                  {priceSuffix}
                </Typography>
              );
            })}
          </Box>
        )}

        {cart.message && (
          <Typography variant="body2" sx={{ mt: 1.5, fontStyle: 'italic' }}>
            “{cart.message}”
          </Typography>
        )}

        {cart.claimedAt && (
          <Typography variant="caption" color="text.secondary" sx={{ display: 'block', mt: 1 }}>
            Marked sent {new Date(cart.claimedAt).toLocaleString()}
          </Typography>
        )}
        {mode === 'completed' && cart.decidedAt && (
          <Typography variant="caption" color="text.secondary" sx={{ display: 'block' }}>
            Confirmed {new Date(cart.decidedAt).toLocaleString()}
          </Typography>
        )}

        {mode === 'to-review' && (
          <Stack direction="row" spacing={1} sx={{ mt: 2 }}>
            <Button
              variant="contained"
              color="success"
              disabled={busy}
              onClick={onApprove}
              sx={{ color: '#fff' }}
            >
              Approve
            </Button>
            <Button variant="outlined" color="error" disabled={busy} onClick={onReject}>
              Reject
            </Button>
          </Stack>
        )}

        {mode === 'held' && (
          <Stack direction="row" spacing={1} sx={{ mt: 2 }}>
            <Button variant="outlined" color="error" disabled={busy} onClick={onDelete}>
              Delete cart
            </Button>
          </Stack>
        )}
      </CardContent>
    </Card>
  );
}
