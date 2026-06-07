import { useMemo, useState } from 'react';
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  Stack,
  Box,
  Typography,
  Checkbox,
  Divider,
  Alert,
  CircularProgress,
} from '@mui/material';
import OpenInNewIcon from '@mui/icons-material/OpenInNew';
import Inventory2Icon from '@mui/icons-material/Inventory2Outlined';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';
import StorefrontIcon from '@mui/icons-material/StorefrontOutlined';
import ShoppingCartIcon from '@mui/icons-material/ShoppingCartOutlined';
import LocalMallIcon from '@mui/icons-material/LocalMallOutlined';
import ChairIcon from '@mui/icons-material/ChairOutlined';
import RedeemIcon from '@mui/icons-material/RedeemOutlined';
import HelpOutlineIcon from '@mui/icons-material/HelpOutline';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import {
  reservations,
  formatPriceCents,
  type RegistryItem,
  type Reservation,
} from '../../api';

type Step = 'source' | 'select' | 'review';

// Label used for items that have no source set.
const NO_SOURCE = 'No source';

// Pick an icon that loosely matches a store/source name.
function sourceIcon(source: string) {
  const s = source.toLowerCase();
  if (source === NO_SOURCE) return <HelpOutlineIcon />;
  if (s.includes('amazon')) return <ShoppingCartIcon />;
  if (s.includes('ikea') || s.includes('furniture')) return <ChairIcon />;
  if (s.includes('etsy') || s.includes('handmade')) return <RedeemIcon />;
  if (s.includes('mall') || s.includes('mamas') || s.includes('papas') || s.includes('lewis'))
    return <LocalMallIcon />;
  return <StorefrontIcon />;
}

// An item eligible to be ordered: it has at least one reservation whose payment
// has been received (confirmed by the owner) but the order hasn't been placed
// online yet. `orderReservations` are the holds that move to Purchased.
interface OrderItem {
  item: RegistryItem;
  orderReservations: Reservation[];
  quantity: number;
}

// BuildOrderDialog walks the owner through placing the gifts whose payment has
// arrived. Orders are built one source at a time: step 1 picks the source,
// step 2 lists that source's "Payment received" items by category with product
// links and checkboxes (plus select-all), and step 3 reviews the picks and, on
// "I've made the order", bulk-advances their reservations to Purchased.
export default function BuildOrderDialog({
  open,
  onClose,
  items,
  reservationsByItem,
  registryId,
}: {
  open: boolean;
  onClose: () => void;
  items: RegistryItem[];
  reservationsByItem: Record<string, Reservation[]>;
  registryId: string;
}) {
  const qc = useQueryClient();
  const [step, setStep] = useState<Step>('source');
  const [selectedSource, setSelectedSource] = useState<string | null>(null);
  const [selectedIds, setSelectedIds] = useState<Set<string>>(new Set());
  const [error, setError] = useState<string | null>(null);

  // Every item with money in hand but not yet ordered, regardless of source.
  const allOrderItems = useMemo<OrderItem[]>(() => {
    const out: OrderItem[] = [];
    for (const item of items) {
      const held = (reservationsByItem[item.id] ?? []).filter(
        (r) => r.status === 'PaymentReceived',
      );
      if (held.length === 0) continue;
      const quantity = held.reduce((n, r) => n + (r.quantity ?? 1), 0);
      out.push({ item, orderReservations: held, quantity });
    }
    return out.sort((a, b) =>
      (a.item.title || '').localeCompare(b.item.title || '', undefined, { sensitivity: 'base' }),
    );
  }, [items, reservationsByItem]);

  // Sources that have at least one eligible item, with their counts. Orders are
  // placed per source, so the owner picks one before choosing items.
  const sources = useMemo(() => {
    const counts = new Map<string, number>();
    for (const oi of allOrderItems) {
      const src = (oi.item.source || '').trim() || NO_SOURCE;
      counts.set(src, (counts.get(src) ?? 0) + 1);
    }
    return Array.from(counts.entries())
      .map(([source, count]) => ({ source, count }))
      .sort((a, b) => a.source.localeCompare(b.source, undefined, { sensitivity: 'base' }));
  }, [allOrderItems]);

  // Eligible items for the chosen source only.
  const orderItems = useMemo<OrderItem[]>(() => {
    if (!selectedSource) return [];
    return allOrderItems.filter(
      (oi) => ((oi.item.source || '').trim() || NO_SOURCE) === selectedSource,
    );
  }, [allOrderItems, selectedSource]);

  // Group the chosen source's items by category for the selection step.
  const groupedByCategory = useMemo(() => {
    const groups = new Map<string, OrderItem[]>();
    for (const oi of orderItems) {
      const cat = (oi.item.category || '').trim() || 'Uncategorised';
      const arr = groups.get(cat) ?? [];
      arr.push(oi);
      groups.set(cat, arr);
    }
    return Array.from(groups.entries()).sort((a, b) =>
      a[0].localeCompare(b[0], undefined, { sensitivity: 'base' }),
    );
  }, [orderItems]);

  const allSelected = orderItems.length > 0 && selectedIds.size === orderItems.length;
  const someSelected = selectedIds.size > 0 && !allSelected;

  const selectedOrderItems = orderItems.filter((oi) => selectedIds.has(oi.item.id));

  const pickSource = (source: string) => {
    setSelectedSource(source);
    setSelectedIds(new Set(orderItemsForSource(source).map((oi) => oi.item.id)));
    setError(null);
    setStep('select');
  };

  const orderItemsForSource = (source: string) =>
    allOrderItems.filter((oi) => ((oi.item.source || '').trim() || NO_SOURCE) === source);

  const toggleItem = (id: string) => {
    setSelectedIds((prev) => {
      const next = new Set(prev);
      if (next.has(id)) next.delete(id);
      else next.add(id);
      return next;
    });
  };

  const toggleAll = () => {
    setSelectedIds((prev) =>
      prev.size === orderItems.length ? new Set() : new Set(orderItems.map((oi) => oi.item.id)),
    );
  };

  const reset = () => {
    setStep('source');
    setSelectedSource(null);
    setSelectedIds(new Set());
    setError(null);
  };

  const placeOrderM = useMutation({
    mutationFn: async () => {
      const toUpdate = selectedOrderItems.flatMap((oi) => oi.orderReservations);
      await Promise.all(toUpdate.map((r) => reservations.setStatus(r.id, 'Purchased')));
    },
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['reservations', registryId] });
      qc.invalidateQueries({ queryKey: ['items', registryId] });
      reset();
    },
    onError: (err) => setError((err as Error).message),
  });

  const handleClose = () => {
    if (placeOrderM.isPending) return;
    onClose();
  };

  const renderRow = (oi: OrderItem, withCheckbox: boolean) => {
    const { item } = oi;
    const url = item.affiliateUrl || item.productUrl;
    const price = formatPriceCents(item.priceCents, item.currency);
    const checked = selectedIds.has(item.id);
    return (
      <Stack key={item.id} direction="row" spacing={1.5} alignItems="center" sx={{ py: 1.5 }}>
        {withCheckbox && (
          <Checkbox
            checked={checked}
            onChange={() => toggleItem(item.id)}
            sx={{ p: 0.5 }}
            inputProps={{ 'aria-label': `Select ${item.title}` }}
          />
        )}
        <Box
          sx={{
            width: 56,
            height: 56,
            flexShrink: 0,
            borderRadius: 2,
            bgcolor: item.imageBgColor || '#ffffff',
            border: '1px solid',
            borderColor: 'divider',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            overflow: 'hidden',
          }}
        >
          {item.imageUrl ? (
            <Box
              component="img"
              src={item.imageUrl}
              alt={item.title}
              sx={{ width: '100%', height: '100%', objectFit: 'contain' }}
            />
          ) : (
            <Typography variant="caption" color="text.disabled">
              No image
            </Typography>
          )}
        </Box>
        <Stack spacing={0.25} sx={{ flex: 1, minWidth: 0 }}>
          <Typography
            variant="subtitle2"
            sx={{
              fontWeight: 600,
              lineHeight: 1.3,
              display: '-webkit-box',
              WebkitLineClamp: 2,
              WebkitBoxOrient: 'vertical',
              overflow: 'hidden',
            }}
          >
            {item.title}
          </Typography>
          <Typography variant="caption" color="text.secondary">
            Qty {oi.quantity}
            {price ? ` · ${price} each` : ''}
          </Typography>
          {url && (
            <Button
              size="small"
              color="inherit"
              component="a"
              href={url}
              target="_blank"
              rel="noopener noreferrer"
              startIcon={<OpenInNewIcon sx={{ fontSize: 14 }} />}
              sx={{
                p: 0,
                minWidth: 0,
                alignSelf: 'flex-start',
                color: 'primary.main',
                textTransform: 'none',
                fontWeight: 500,
              }}
            >
              Open product page
            </Button>
          )}
        </Stack>
      </Stack>
    );
  };

  return (
    <Dialog
      open={open}
      onClose={handleClose}
      TransitionProps={{ onExited: reset }}
      fullWidth
      maxWidth="sm"
      PaperProps={{ sx: { maxHeight: '90vh', display: 'flex', flexDirection: 'column' } }}
    >
      <DialogTitle>
        {step === 'source'
          ? 'Build an order'
          : step === 'select'
            ? `Order from ${selectedSource ?? ''}`
            : 'Confirm your order'}
      </DialogTitle>
      <DialogContent dividers sx={{ flex: 1, overflowY: 'auto' }}>
        {allOrderItems.length === 0 ? (
          <Stack spacing={1} sx={{ py: 4, textAlign: 'center', alignItems: 'center' }}>
            <Box
              sx={{
                width: 56,
                height: 56,
                borderRadius: '50%',
                bgcolor: 'action.hover',
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
              }}
            >
              <Inventory2Icon sx={{ color: 'text.disabled' }} />
            </Box>
            <Typography variant="h6" sx={{ fontWeight: 700 }}>
              Nothing to order yet
            </Typography>
            <Typography color="text.secondary" sx={{ maxWidth: 360 }}>
              Items appear here once you approve a payment in the “To review” tab. Approved gifts
              move to “Payment received” and are ready to be ordered.
            </Typography>
          </Stack>
        ) : step === 'source' ? (
          <Stack spacing={1}>
            <Typography variant="body2" color="text.secondary">
              Pick a store to order from. You build one order per store.
            </Typography>
            <Stack divider={<Divider />}>
              {sources.map(({ source, count }) => (
                <Stack
                  key={source}
                  direction="row"
                  alignItems="center"
                  spacing={1}
                  role="button"
                  tabIndex={0}
                  onClick={() => pickSource(source)}
                  onKeyDown={(e) => {
                    if (e.key === 'Enter' || e.key === ' ') {
                      e.preventDefault();
                      pickSource(source);
                    }
                  }}
                  sx={{
                    py: 1.5,
                    px: 1,
                    cursor: 'pointer',
                    borderRadius: 1,
                    '&:hover': { bgcolor: 'action.hover' },
                  }}
                >
                  <Box
                    sx={{
                      width: 40,
                      height: 40,
                      borderRadius: 1,
                      bgcolor: 'action.hover',
                      display: 'flex',
                      alignItems: 'center',
                      justifyContent: 'center',
                      color: 'text.secondary',
                      flexShrink: 0,
                    }}
                  >
                    {sourceIcon(source)}
                  </Box>
                  <Stack sx={{ flex: 1, minWidth: 0 }}>
                    <Typography variant="subtitle2" sx={{ fontWeight: 600 }}>
                      {source}
                    </Typography>
                    <Typography variant="caption" color="text.secondary">
                      {count} {count === 1 ? 'item' : 'items'} ready to order
                    </Typography>
                  </Stack>
                  <ChevronRightIcon sx={{ color: 'text.disabled' }} />
                </Stack>
              ))}
            </Stack>
          </Stack>
        ) : step === 'select' ? (
          <Stack spacing={1}>
            <Typography variant="body2" color="text.secondary">
              Select the gifts you're ordering now. Use the product links to place each order, then
              continue to confirm.
            </Typography>
            <Stack
              direction="row"
              alignItems="center"
              spacing={1.5}
              sx={{ py: 1, borderBottom: 1, borderColor: 'divider' }}
            >
              <Checkbox
                checked={allSelected}
                indeterminate={someSelected}
                onChange={toggleAll}
                sx={{ p: 0.5 }}
                inputProps={{ 'aria-label': 'Select all items' }}
              />
              <Typography variant="subtitle2" sx={{ fontWeight: 700 }}>
                Select all ({orderItems.length})
              </Typography>
            </Stack>
            {groupedByCategory.map(([category, group]) => (
              <Box key={category}>
                <Typography
                  variant="overline"
                  color="text.secondary"
                  sx={{ display: 'block', mt: 1.5 }}
                >
                  {category}
                </Typography>
                <Stack divider={<Divider />}>{group.map((oi) => renderRow(oi, true))}</Stack>
              </Box>
            ))}
          </Stack>
        ) : (
          <Stack spacing={1}>
            <Alert severity="info" icon={false}>
              You're about to mark {selectedOrderItems.length}{' '}
              {selectedOrderItems.length === 1 ? 'gift' : 'gifts'} from {selectedSource} as
              purchased. Do this once you've placed the order online.
            </Alert>
            {error && <Alert severity="error">{error}</Alert>}
            <Stack divider={<Divider />}>{selectedOrderItems.map((oi) => renderRow(oi, false))}</Stack>
          </Stack>
        )}
      </DialogContent>
      <DialogActions sx={{ px: 3, py: 2 }}>
        {step === 'source' ? (
          <Button onClick={handleClose}>{allOrderItems.length === 0 ? 'Close' : 'Cancel'}</Button>
        ) : step === 'select' ? (
          <>
            <Button onClick={() => setStep('source')}>Back</Button>
            <Box sx={{ flexGrow: 1 }} />
            <Button
              variant="contained"
              disabled={selectedIds.size === 0}
              onClick={() => {
                setError(null);
                setStep('review');
              }}
            >
              Review order ({selectedIds.size})
            </Button>
          </>
        ) : (
          <>
            <Button onClick={() => setStep('select')} disabled={placeOrderM.isPending}>
              Back
            </Button>
            <Box sx={{ flexGrow: 1 }} />
            <Button
              variant="contained"
              onClick={() => placeOrderM.mutate()}
              disabled={placeOrderM.isPending}
              startIcon={placeOrderM.isPending ? <CircularProgress size={16} color="inherit" /> : undefined}
            >
              {placeOrderM.isPending ? 'Saving…' : "I've made the order"}
            </Button>
          </>
        )}
      </DialogActions>
    </Dialog>
  );
}
