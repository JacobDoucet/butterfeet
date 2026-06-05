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
} from '@mui/material';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { payments, formatPriceCents, type Registry, type Cart } from '../../api';

type Mode = 'to-review' | 'completed';

// CartsPanel renders the owner's cart-review surface. In "to-review" mode it
// lists carts a guest has marked as paid (AwaitingConfirmation) with
// approve/reject actions. In "completed" mode it lists carts the owner has
// already confirmed.
export default function CartsPanel({ reg, mode }: { reg: Registry; mode: Mode }) {
  const qc = useQueryClient();
  // "To review" surfaces carts a guest has locked in (Pending = mid-flow, money
  // not yet marked sent; AwaitingConfirmation = guest says they've paid).
  const status = mode === 'to-review' ? 'Pending,AwaitingConfirmation' : 'Completed';

  const cartsQ = useQuery({
    queryKey: ['carts', reg.id, status],
    queryFn: () => payments.listCarts(reg.id, status),
  });

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

  const carts = cartsQ.data ?? [];
  const busy = approveM.isPending || rejectM.isPending;

  return (
    <Stack spacing={2}>
      <Box>
        <Typography variant="h6">{mode === 'to-review' ? 'To review' : 'Completed'}</Typography>
        {mode === 'to-review' ? (
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
          {mode === 'to-review'
            ? 'No payments are waiting for your review.'
            : 'No confirmed payments yet.'}
        </Typography>
      ) : (
        <Stack spacing={1.5}>
          {carts.map((c) => (
            <CartCard
              key={c.id}
              cart={c}
              mode={mode}
              busy={busy}
              onApprove={() => approveM.mutate(c.id)}
              onReject={() => rejectM.mutate(c.id)}
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
  onApprove,
  onReject,
}: {
  cart: Cart;
  mode: Mode;
  busy: boolean;
  onApprove: () => void;
  onReject: () => void;
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
            {mode === 'to-review' &&
              (cart.status === 'AwaitingConfirmation' ? (
                <Chip size="small" color="warning" label="Says they've paid" />
              ) : (
                <Chip size="small" variant="outlined" label="Awaiting payment" />
              ))}
          </Stack>
        </Stack>

        {cart.items.length > 0 && (
          <Stack spacing={0.25} sx={{ mt: 1.5 }}>
            {cart.items.map((it) => (
              <Typography key={it.reservationId} variant="body2">
                {it.quantity}× {it.title}
                {it.priceCents ? ` — ${formatPriceCents(it.priceCents, it.currency)}` : ''}
              </Typography>
            ))}
          </Stack>
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
      </CardContent>
    </Card>
  );
}
