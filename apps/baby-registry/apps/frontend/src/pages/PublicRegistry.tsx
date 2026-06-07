import { useState, useEffect, useRef } from 'react';
import { useParams } from 'react-router-dom';
import {
  Container,
  Typography,
  Card,
  CardContent,
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
  Select,
  MenuItem,
  Tabs,
  Tab,
  Fab,
  Badge,
  useMediaQuery,
  useTheme,
} from '@mui/material';
import ShoppingCartOutlinedIcon from '@mui/icons-material/ShoppingCartOutlined';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { pub, buyer, isGatedRegistry, formatPriceCents, publicPayments, exchangeRates, convertCents, type RegistryAccessRequestStatus, type PaymentMethod } from '../api';
import { useSetActiveThemeColor } from '../activeTheme';
import { useViewerCurrency } from '../viewerCurrency';
import ItemCard from '../components/ItemCard';
import CartDrawer, { type CartLine } from '../components/public/CartDrawer';
import BuyerVerifyGate from '../components/public/BuyerVerifyGate';
import RegistryAccessGate from '../components/public/RegistryAccessGate';

type ClickableItem = {
  id: string;
  productUrl?: string;
  affiliateUrl?: string;
  retailer?: string;
};

function purchaseHref(item: ClickableItem | undefined | null): string | undefined {
  if (!item) return undefined;
  return item.affiliateUrl || item.productUrl || undefined;
}

// detectViewerCurrency picks a sensible default display currency from the
// browser locale, limited to the currencies we support. Falls back to USD.
function trackPurchaseClick(item: ClickableItem | undefined | null) {
  if (!item || !item.id) return;
  try {
    const payload = JSON.stringify({ retailer: item.retailer || 'unknown' });
    if (typeof navigator !== 'undefined' && navigator.sendBeacon) {
      const blob = new Blob([payload], { type: 'application/json' });
      navigator.sendBeacon(`/api/public/items/${item.id}/click`, blob);
    } else {
      fetch(`/api/public/items/${item.id}/click`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: payload,
        keepalive: true,
      }).catch(() => {});
    }
  } catch {
    // best-effort tracking; never block navigation
  }
}

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

  useSetActiveThemeColor(regQ.data?.themeColor);

  const paymentMethodsQ = useQuery({
    queryKey: ['public-payment-methods', slug],
    queryFn: () => publicPayments.methods(slug),
    enabled: verified && !!regQ.data && !isGatedRegistry(regQ.data),
  });

  const ratesQ = useQuery({
    queryKey: ['exchange-rates'],
    queryFn: () => exchangeRates.get(),
    staleTime: 60 * 60 * 1000,
    retry: false,
  });

  const { currency: viewerCurrency } = useViewerCurrency();

  // displayPrice renders a price (stored in `cents` of `currency`) in the
  // viewer's chosen currency. When the original already matches, or no rate is
  // available, it returns the original formatting. The `approx` flag tells the
  // caller whether the figure is a converted estimate (prefixed with "≈").
  const displayPrice = (
    cents?: number,
    currency?: string,
  ): { text: string; approx: boolean } | null => {
    if (cents == null || Number.isNaN(cents)) return null;
    const from = (currency || 'USD').toUpperCase();
    if (from === viewerCurrency) {
      const text = formatPriceCents(cents, from);
      return text ? { text, approx: false } : null;
    }
    const converted = convertCents(cents, from, viewerCurrency, ratesQ.data);
    if (converted == null) {
      const text = formatPriceCents(cents, from);
      return text ? { text, approx: false } : null;
    }
    const text = formatPriceCents(converted, viewerCurrency);
    return text ? { text, approx: true } : null;
  };

  const [target, setTarget] = useState<string | null>(null);
  const [name, setName] = useState('');
  const [message, setMessage] = useState('');
  const [reserveQtyMode, setReserveQtyMode] = useState<'one' | 'all'>('one');
  const [reserveQty, setReserveQty] = useState('1');
  const [selectedOptionId, setSelectedOptionId] = useState<string | null>(null);
  const [categoryFilter, setCategoryFilter] = useState<string>('__all__');
  const [reservedId, setReservedId] = useState<string | null>(null);
  const [cartOpen, setCartOpen] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [snack, setSnack] = useState<{ msg: string; cart?: boolean } | null>(null);
  // Auto-open the cart once when the buyer lands with a held (Pending) cart —
  // a payment they started but never confirmed — so they're nudged to finish
  // or cancel it. The ref guards against re-opening after they close it.
  const autoOpenedHeldRef = useRef(false);
  useEffect(() => {
    if (autoOpenedHeldRef.current) return;
    const data = regQ.data;
    if (!data || isGatedRegistry(data)) return;
    const hasHeld = (data.myCarts ?? []).some((c) => c.status === 'Pending');
    if (hasHeld) {
      autoOpenedHeldRef.current = true;
      setCartOpen(true);
    }
  }, [regQ.data]);
  // Snapshot of the last opened target so the public Dialog stays
  // populated through its exit animation. Declared at the top to satisfy
  // React's rules of hooks (early returns happen further below).
  const targetSnapshotRef = useRef<{
    item: any;
    root: any;
    options: any[];
    remaining: number;
  } | null>(null);
  const modalPaperSx = {
    maxHeight: '90vh',
    display: 'flex',
    flexDirection: 'column',
  } as const;
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('sm'));

  // Reset selected option whenever a new card is opened
  useEffect(() => {
    if (target) setSelectedOptionId(null);
  }, [target]);

  // Seed the reservation form's name field with the verified buyer's name.
  useEffect(() => {
    const buyerName = meQ.data?.name?.trim();
    if (buyerName) setName((cur) => (cur ? cur : buyerName));
  }, [meQ.data?.name]);

  useEffect(() => {
    const title = regQ.data?.title;
    if (!title) return;
    const prev = document.title;
    document.title = `${title} · baby registry`;
    return () => { document.title = prev; };
  }, [regQ.data?.title]);

  const reserveM = useMutation({
    mutationFn: async () => {
      if (!target) return;
      const data = regQ.data;
      if (!data || isGatedRegistry(data)) return;
      const effectiveId = selectedOptionId ?? target;
      const rootItem = data.items.find((it) => it.id === target);
      const qty = rootItem?.quantityUnlimited
        ? Math.max(1, parseInt(reserveQty || '1', 10) || 1)
        : reserveQtyMode === 'all'
        ? Math.max(1, (rootItem?.quantity || 1) - (rootItem?.reserved || 0))
        : 1;
      const reserverName = (name.trim() || meQ.data?.name?.trim() || meQ.data?.email || 'Guest');
      return pub.reserve(effectiveId, { reserverName, isAnonymous: false, message, quantity: qty });
    },
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['public', slug] });
      setError(null);
      setTarget(null);
      setSelectedOptionId(null);
      setMessage('');
      setReserveQtyMode('one');
      setReserveQty('1');
      setSnack({ msg: 'Added to your cart — held for you for 24 hours.', cart: true });
    },
    onError: (err) => setError((err as Error).message),
  });

  const resetReserveDialog = () => {
    setTarget(null);
    setReservedId(null);
    setSelectedOptionId(null);
    setName('');
    setMessage('');
    setError(null);
    setReserveQtyMode('one');
    setReserveQty('1');
  };

  const cancelRsvM = useMutation({
    mutationFn: (id: string) => pub.cancelReservation(id),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['public', slug] });
      setSnack({ msg: 'Removed from your cart \u2014 the gift is available again.' });
      resetReserveDialog();
    },
    onError: (err) => setSnack({ msg: (err as Error).message }),
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
  if (isGatedRegistry(regQ.data)) {
    return (
      <RegistryAccessGate
        slug={slug}
        viewerEmail={meQ.data?.email ?? ''}
        gated={regQ.data}
        onSubmitted={() => qc.invalidateQueries({ queryKey: ['public', slug] })}
      />
    );
  }
  const reg = regQ.data;
  const itemById = reg.items.reduce<Record<string, typeof reg.items[number]>>((acc, it) => {
    acc[it.id] = it;
    return acc;
  }, {});
  const rootItemById = reg.items.reduce<Record<string, typeof reg.items[number]>>((acc, it) => {
    const root = it.parentItemId && itemById[it.parentItemId] ? itemById[it.parentItemId] : it;
    acc[it.id] = root;
    return acc;
  }, {});
  const topLevelItems = reg.items.filter((it) => !it.parentItemId || !itemById[it.parentItemId]);
  const myReservationByRootId: Record<string, NonNullable<typeof reg.myReservations>[number]> = {};
  for (const rsv of reg.myReservations ?? []) {
    const root = rootItemById[rsv.itemId];
    if (root) myReservationByRootId[root.id] = rsv;
  }
  // Cart lines are the buyer's active reservations, joined with live item
  // details so we can show images, prices, and retailer links in the cart.
  const cartLines: CartLine[] = (reg.myReservations ?? []).map((rsv) => {
    const it = itemById[rsv.itemId];
    return {
      reservationId: rsv.id,
      itemId: rsv.itemId,
      title: it?.title ?? rsv.itemTitle ?? 'Gift',
      imageUrl: it?.imageUrl,
      imageBgColor: it?.imageBgColor,
      priceCents: it?.priceCents,
      currency: it?.currency,
      quantity: rsv.quantity || 1,
      productUrl: it ? it.affiliateUrl || it.productUrl : undefined,
      retailer: it?.retailer,
    };
  });
  // A held cart is one the buyer started paying for (status Pending) but never
  // confirmed. Its reservations are locked (AwaitingConfirmation) so they don't
  // appear in the open cart above — surface it so the buyer can resume and
  // confirm, or cancel to release the gifts.
  const heldCart = (reg.myCarts ?? []).find((c) => c.status === 'Pending');
  const heldCartMethod =
    heldCart && paymentMethodsQ.data
      ? (paymentMethodsQ.data as PaymentMethod[]).find((m) => m.id === heldCart.paymentMethodId)
      : undefined;
  const heldCartLines: CartLine[] = (heldCart?.items ?? []).map((ci) => {
    const it = itemById[ci.itemId];
    return {
      reservationId: ci.reservationId,
      itemId: ci.itemId,
      title: it?.title ?? ci.title ?? 'Gift',
      imageUrl: it?.imageUrl,
      imageBgColor: it?.imageBgColor,
      priceCents: it?.priceCents ?? ci.priceCents,
      currency: it?.currency ?? ci.currency,
      quantity: ci.quantity || 1,
      productUrl: it ? it.affiliateUrl || it.productUrl : undefined,
      retailer: it?.retailer,
    };
  });
  const isClaimed = (rootId: string) => {
    const r = itemById[rootId];
    if (!r) return false;
    if (r.quantityUnlimited) return false;
    return Math.max(0, (r.quantity || 1) - (r.reserved || 0)) === 0;
  };
  const isItemClaimed = (it: typeof reg.items[number]) => {
    const root = rootItemById[it.id] ?? it;
    if (root.quantityUnlimited) return false;
    return Math.max(0, (root.quantity || 1) - (root.reserved || 0)) === 0;
  };
  const sortedTopLevelItems = [...topLevelItems].sort((a, b) => {
    // Items in the buyer's own cart should stay in place, not get pushed to the
    // end like items claimed by others.
    const aInCart = !!myReservationByRootId[(rootItemById[a.id] ?? a).id];
    const bInCart = !!myReservationByRootId[(rootItemById[b.id] ?? b).id];
    const aClaimed = isItemClaimed(a) && !aInCart;
    const bClaimed = isItemClaimed(b) && !bInCart;
    if (aClaimed !== bClaimed) return aClaimed ? 1 : -1;
    return (a.title || '').localeCompare(b.title || '', undefined, { sensitivity: 'base' });
  });
  const categories = Array.from(
    new Set(topLevelItems.map((it) => (it.category || '').trim()).filter(Boolean)),
  ).sort((a, b) => a.localeCompare(b));
  const filteredTopLevelItems =
    categoryFilter === '__all__'
      ? sortedTopLevelItems
      : sortedTopLevelItems.filter((it) => (it.category || '').trim() === categoryFilter);
  const itemsByCategory: Record<string, typeof topLevelItems> = {};
  categories.forEach((cat) => {
    itemsByCategory[cat] = sortedTopLevelItems.filter((it) => (it.category || '').trim() === cat);
  });
  const uncategorisedItems = sortedTopLevelItems.filter((it) => !(it.category || '').trim());

  const renderItemCard = (it: typeof reg.items[number]) => {
    const root = rootItemById[it.id] ?? it;
    const options = optionsByRootId[root.id] ?? [root];
    const remaining = root.quantityUnlimited ? Infinity : Math.max(0, (root.quantity || 1) - (root.reserved || 0));
    const claimed = !root.quantityUnlimited && remaining === 0;
    const myRsv = myReservationByRootId[root.id];
    const optionCount = options.length;
    const lowStock = !root.quantityUnlimited && remaining > 0 && remaining < (root.quantity || 1);
    const optionPrices = options
      .map((o) => o.priceCents)
      .filter((p): p is number => p != null && p > 0);
    const minPrice = optionPrices.length ? Math.min(...optionPrices) : null;
    const maxPrice = optionPrices.length ? Math.max(...optionPrices) : null;
    const priceCurrency = root.currency || options.find((o) => o.priceCents)?.currency;
    const minPriceDisplay = minPrice != null ? displayPrice(minPrice, priceCurrency) : null;
    const priceLabel =
      minPriceDisplay != null
        ? minPrice === maxPrice
          ? minPriceDisplay.text
          : `from ${minPriceDisplay.text}`
        : null;
    const swatchImages = options
      .filter((o) => !!o.imageUrl)
      .slice(0, 4)
      .map((o) => o.imageUrl as string);

    const topLeftOverlay = optionCount > 1 ? (
      <Box
        sx={{
          bgcolor: 'rgba(255,255,255,0.95)',
          borderRadius: 5,
          px: 1.25,
          py: 0.5,
          display: 'flex',
          alignItems: 'center',
          gap: 0.5,
          boxShadow: '0 1px 3px rgba(0,0,0,0.08)',
        }}
      >
        {swatchImages.slice(0, 3).map((src, idx) => (
          <Box
            key={`${it.id}-sw-${idx}`}
            component="img"
            src={src}
            alt=""
            sx={{
              width: 18,
              height: 18,
              borderRadius: '50%',
              objectFit: 'cover',
              border: '1px solid #fff',
              ml: idx === 0 ? 0 : -0.75,
              bgcolor: '#f4ede3',
            }}
          />
        ))}
        <Typography variant="caption" sx={{ fontWeight: 600, ml: 0.5 }}>
          +{optionCount - 1} more
        </Typography>
      </Box>
    ) : null;

    const topRightOverlay = myRsv ? (
      <Box
        sx={{
          bgcolor: 'primary.main',
          color: '#fff',
          px: 1.5,
          py: 0.5,
          borderRadius: 5,
          fontSize: '0.75rem',
          fontWeight: 700,
          letterSpacing: 0.3,
          display: 'flex',
          alignItems: 'center',
          gap: 0.5,
          boxShadow: '0 2px 6px rgba(0,0,0,0.15)',
        }}
      >
        <ShoppingCartOutlinedIcon sx={{ fontSize: '0.95rem', color: '#fff' }} />
        In your cart
      </Box>
    ) : claimed ? (
      <Box
        sx={{
          bgcolor: 'success.main',
          color: '#fff',
          px: 1.5,
          py: 0.5,
          borderRadius: 5,
          fontSize: '0.75rem',
          fontWeight: 700,
          letterSpacing: 0.3,
          display: 'flex',
          alignItems: 'center',
          gap: 0.5,
          boxShadow: '0 2px 6px rgba(0,0,0,0.15)',
        }}
      >
        ❤ Claimed
      </Box>
    ) : lowStock ? (
      <Box
        sx={{
          bgcolor: 'warning.main',
          color: '#fff',
          px: 1.25,
          py: 0.5,
          borderRadius: 5,
          fontSize: '0.7rem',
          fontWeight: 700,
          letterSpacing: 0.3,
        }}
      >
        Only {remaining} left
      </Box>
    ) : null;

    return (
      <Grid item xs={12} sm={6} md={4} lg={3} key={it.id}>
        <ItemCard
          imageUrl={it.imageUrl}
          imageBgColor={it.imageBgColor}
          title={it.title}
          onClick={() => {
            if (myRsv) {
              setTarget(it.id);
              setReservedId(myRsv.id);
            } else {
              setTarget(it.id);
            }
          }}
          disabled={claimed && !myRsv}
          dimTitle={claimed && !myRsv}
          imageFilter={claimed && !myRsv ? 'grayscale(0.4)' : undefined}
          topLeftOverlay={topLeftOverlay}
          topRightOverlay={topRightOverlay}
          belowTitle={
            priceLabel ? (
              <Typography
                sx={{
                  fontWeight: 700,
                  fontSize: '1.05rem',
                  color: claimed && !myRsv ? 'text.secondary' : 'text.primary',
                  letterSpacing: '0.01em',
                  mb: 0.5,
                }}
              >
                {priceLabel}
              </Typography>
            ) : undefined
          }
          footer={
            <>
              {it.source ? (
                <Typography variant="caption" color="text.secondary" sx={{ textTransform: 'uppercase', letterSpacing: 0.5, fontWeight: 600 }}>
                  {it.source}
                </Typography>
              ) : (
                <span />
              )}
              {myRsv ? (
                <Typography variant="caption" color="primary.main" sx={{ fontWeight: 600 }}>
                  In cart · manage →
                </Typography>
              ) : !claimed ? (
                <Typography variant="caption" color="primary.main" sx={{ fontWeight: 600 }}>
                  Add to cart →
                </Typography>
              ) : null}
            </>
          }
        />
      </Grid>
    );
  };
  const optionsByRootId = reg.items.reduce<Record<string, typeof reg.items[number][]>>((acc, it) => {
    const root = rootItemById[it.id] ?? it;
    (acc[root.id] ??= []).push(it);
    return acc;
  }, {});
  const liveTargetItem = reg.items.find((it) => it.id === target) ?? null;
  const liveTargetRoot = target ? rootItemById[target] ?? liveTargetItem : liveTargetItem;
  const liveTargetOptions = liveTargetRoot ? optionsByRootId[liveTargetRoot.id] ?? [liveTargetRoot] : [];
  const liveTargetRemaining = liveTargetRoot?.quantityUnlimited
    ? Infinity
    : Math.max(0, (liveTargetRoot?.quantity || 1) - (liveTargetRoot?.reserved || 0));
  if (liveTargetRoot) {
    targetSnapshotRef.current = {
      item: liveTargetItem,
      root: liveTargetRoot,
      options: liveTargetOptions,
      remaining: liveTargetRemaining,
    };
  }
  const targetItem = liveTargetItem ?? targetSnapshotRef.current?.item ?? null;
  const targetRootItem = liveTargetRoot ?? targetSnapshotRef.current?.root ?? null;
  const targetOptions = (liveTargetRoot ? liveTargetOptions : targetSnapshotRef.current?.options) ?? [];
  const targetRemaining = liveTargetRoot ? liveTargetRemaining : targetSnapshotRef.current?.remaining ?? 0;

  return (
    <Box
      sx={{
        minHeight: '100%',
        overflowX: 'hidden',
        maxWidth: '100vw',
        ...(reg.themeColor
          ? { bgcolor: reg.themeColor }
          : {
              background: [
                'radial-gradient(ellipse 80% 60% at 0% 0%, rgba(232,168,124,0.18), transparent 60%)',
                'radial-gradient(ellipse 70% 50% at 100% 10%, rgba(122,158,126,0.16), transparent 60%)',
                'radial-gradient(ellipse 90% 70% at 50% 100%, rgba(232,168,124,0.10), transparent 65%)',
                'linear-gradient(180deg, #fbf7f2 0%, #f6efe6 100%)',
              ].join(', '),
              backgroundAttachment: 'fixed',
              backgroundRepeat: 'no-repeat',
            }),
      }}
    >
      <Box
        sx={{
          pt: { xs: 4, md: 6 },
        }}
      >
        <Container maxWidth="lg">
          <Stack alignItems="center" sx={{ textAlign: 'center' }}>
            <Typography
              variant="overline"
              sx={{ letterSpacing: 2, color: 'text.secondary', mb: 1 }}
            >
              A baby registry
            </Typography>
            <Typography
              variant="h2"
              sx={{
                fontWeight: 700,
                fontSize: { xs: '2rem', sm: '2.75rem', md: '3.25rem' },
                lineHeight: 1.1,
                letterSpacing: '0.04em',
              }}
              gutterBottom
            >
              {reg.title}
            </Typography>
            {reg.parentNames && (
              <Typography variant="subtitle1" color="text.secondary" sx={{ textTransform: 'uppercase', letterSpacing: '0.06em', fontSize: '0.75rem' }}>
                with love, for {reg.parentNames}
              </Typography>
            )}
            {reg.welcomeMessage && (
              <Typography sx={{ mt: 2, maxWidth: 640, color: 'text.secondary', lineHeight: 1.6 }}>
                {reg.welcomeMessage}
              </Typography>
            )}
          </Stack>
        </Container>
      </Box>

      <Container maxWidth="lg" sx={{ pt: { xs: 1, md: 2 }, pb: { xs: 4, md: 6 } }}>
        {categories.length > 0 && (
          <Box
            sx={{
              mb: 4,
              position: 'sticky',
              top: { xs: 64, sm: 68 },
              zIndex: 3,
              borderRadius: 0,
              mx: 'calc(50% - 50vw)',
              px: 'calc(50vw - 50%)',
              pt: '24px',
              bgcolor: 'transparent',
              backdropFilter: 'blur(14px) saturate(140%)',
              WebkitBackdropFilter: 'blur(14px) saturate(140%)',
            }}
          >
            <Tabs
              value={categoryFilter}
              onChange={(_, v) => setCategoryFilter(v)}
              variant="scrollable"
              scrollButtons="auto"
              allowScrollButtonsMobile
              sx={{
                minHeight: 52,
                '& .MuiTab-root': {
                  fontWeight: 600,
                  fontSize: '0.85rem',
                  minHeight: 52,
                  px: 2.5,
                },
              }}
            >
              <Tab
                value="__all__"
                label={
                  <span>
                    All
                    <Box component="span" sx={{ color: 'text.disabled', fontWeight: 400, ml: 0.75 }}>
                      · {topLevelItems.length}
                    </Box>
                  </span>
                }
              />
              {categories.map((cat) => {
                const count = topLevelItems.filter((it) => (it.category || '').trim() === cat).length;
                return (
                  <Tab
                    key={cat}
                    value={cat}
                    label={
                      <span>
                        {cat}
                        <Box component="span" sx={{ color: 'text.disabled', fontWeight: 400, ml: 0.75 }}>
                          · {count}
                        </Box>
                      </span>
                    }
                  />
                );
              })}
            </Tabs>
          </Box>
        )}

        {categoryFilter === '__all__' && categories.length > 1 ? (
          <Stack spacing={5}>
            {categories.map((cat) => {
              const all = itemsByCategory[cat];
              const preview = all.slice(0, 4);
              const hasMore = all.length > preview.length;
              return (
                <Box key={cat}>
                  <Stack
                    direction="row"
                    alignItems="baseline"
                    justifyContent="space-between"
                    sx={{ mb: 2 }}
                  >
                    <Typography variant="h5" sx={{ fontWeight: 600, letterSpacing: '0.04em', textTransform: 'uppercase' }}>
                      {cat}
                    </Typography>
                    {hasMore && (
                      <Button
                        size="small"
                        onClick={() => setCategoryFilter(cat)}
                        sx={{ textTransform: 'uppercase', fontWeight: 600, letterSpacing: '0.06em', fontSize: '0.7rem' }}
                      >
                        See all ({all.length}) →
                      </Button>
                    )}
                  </Stack>
                  <Grid container spacing={3}>
                    {preview.map((it) => renderItemCard(it))}
                  </Grid>
                </Box>
              );
            })}
            {uncategorisedItems.length > 0 && (
              <Box>
                <Typography variant="h5" sx={{ fontWeight: 600, letterSpacing: '-0.01em', mb: 2 }}>
                  Everything else
                </Typography>
                <Grid container spacing={3}>
                  {uncategorisedItems.slice(0, 4).map((it) => renderItemCard(it))}
                </Grid>
              </Box>
            )}
          </Stack>
        ) : (
          <Grid container spacing={3}>
            {filteredTopLevelItems.map((it) => renderItemCard(it))}
            {filteredTopLevelItems.length === 0 && (
              <Grid item xs={12}>
                <Box sx={{ py: 8, textAlign: 'center', color: 'text.secondary' }}>
                  <Typography variant="body1">No items in this category yet.</Typography>
                </Box>
              </Grid>
            )}
          </Grid>
        )}

        <Dialog
          open={!!target || !!reservedId}
          onClose={() => { setReservedId(null); setTarget(null); }}
          scroll="paper"
          fullWidth
          maxWidth="sm"
          fullScreen={isMobile}
          TransitionProps={{
            onExited: () => {
              targetSnapshotRef.current = null;
            },
          }}
          PaperProps={{ sx: { ...modalPaperSx, width: { xs: '100%', sm: 480 }, maxWidth: '100vw' } }}
        >
          <DialogTitle sx={{ pb: 1, px: { xs: 2, sm: 3 }, pt: { xs: 1.5, sm: 2 } }}>
            {reservedId && (
              <Typography variant="overline" color="text.secondary" sx={{ display: 'block', lineHeight: 1, mb: 1 }}>
                Held for you
              </Typography>
            )}
            {targetOptions.length === 1 && targetOptions[0] ? (
              <Stack direction="row" spacing={2} alignItems="center">
                {(() => {
                  const href = purchaseHref(targetOptions[0]);
                  const imageInner = targetOptions[0].imageUrl ? (
                    <Box
                      component="img"
                      src={targetOptions[0].imageUrl}
                      alt={targetOptions[0].title}
                      sx={{ width: '100%', height: '100%', objectFit: 'contain' }}
                    />
                  ) : (
                    <Typography variant="caption" color="text.disabled">No image</Typography>
                  );
                  return (
                    <Box
                      {...(href
                        ? {
                            component: 'a',
                            href,
                            target: '_blank',
                            rel: 'noopener noreferrer',
                          }
                        : {})}
                      sx={{
                        width: { xs: 64, sm: 120 },
                        height: { xs: 64, sm: 120 },
                        flexShrink: 0,
                        bgcolor: targetOptions[0].imageBgColor || '#ffffff',
                        borderRadius: 2,
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'center',
                        overflow: 'hidden',
                        cursor: href ? 'pointer' : 'default',
                      }}
                    >
                      {imageInner}
                    </Box>
                  );
                })()}
                <Stack spacing={0.5} sx={{ minWidth: 0, flex: 1 }}>
                  <Typography variant="h6" sx={{ fontWeight: 700, lineHeight: 1.2, fontSize: { xs: '1rem', sm: '1.25rem' } }}>
                    {targetOptions[0].title}
                  </Typography>
                  {targetOptions[0].priceCents != null && (() => {
                    const dp = displayPrice(targetOptions[0].priceCents, targetOptions[0].currency);
                    return dp ? (
                      <Typography variant="subtitle1" sx={{ fontWeight: 700, color: 'primary.main' }}>
                        {dp.approx ? `≈ ${dp.text}` : dp.text}
                      </Typography>
                    ) : null;
                  })()}
                  {targetOptions[0].description && (
                    <Typography
                      variant="body2"
                      color="text.secondary"
                      sx={{
                        whiteSpace: 'pre-wrap',
                        wordBreak: 'break-word',
                        fontWeight: 400,
                        display: '-webkit-box',
                        WebkitLineClamp: { xs: 2, sm: 'unset' },
                        WebkitBoxOrient: 'vertical',
                        overflow: 'hidden',
                      }}
                    >
                      {targetOptions[0].description}
                    </Typography>
                  )}
                </Stack>
              </Stack>
            ) : (
              <Stack spacing={0.5}>
                <Typography variant="h6" sx={{ fontWeight: 700, lineHeight: 1.2 }}>
                  {targetRootItem?.title ?? 'Item details'}
                </Typography>
                {(() => {
                  const opt = itemById[selectedOptionId ?? target ?? ''] ?? targetOptions[0];
                  if (opt?.priceCents == null) return null;
                  const dp = displayPrice(opt.priceCents, opt.currency);
                  return dp ? (
                    <Typography variant="subtitle1" sx={{ fontWeight: 700, color: 'primary.main' }}>
                      {dp.approx ? `≈ ${dp.text}` : dp.text}
                    </Typography>
                  ) : null;
                })()}
              </Stack>
            )}
          </DialogTitle>
          <DialogContent
            sx={{
              flex: 1,
              minHeight: 0,
              overflowY: 'auto',
              overflowX: 'hidden',
              scrollbarGutter: 'stable',
              pr: 1,
            }}
          >
            <Stack spacing={2} sx={{ mt: 1 }}>
              {reservedId ? (
                <>
                  <Alert
                    severity="info"
                    icon={false}
                    sx={{ borderRadius: 2, bgcolor: 'primary.main', color: '#fff' }}
                  >
                    <Typography sx={{ fontWeight: 600, color: '#fff', mb: 0.5 }}>
                      This gift is in your cart.
                    </Typography>
                    <Typography variant="body2" sx={{ color: 'rgba(255,255,255,0.9)' }}>
                      We're holding it for you for 24 hours. Head to your cart to review everything and
                      check out, or remove it to make it available again.
                    </Typography>
                  </Alert>
                  {(() => {
                    const effectiveOpt = itemById[selectedOptionId ?? target ?? ''] ?? targetOptions[0];
                    const href = purchaseHref(effectiveOpt);
                    if (!href) return null;
                    return (
                      <Button
                        variant="outlined"
                        onClick={() => { window.open(href, '_blank', 'noopener,noreferrer'); trackPurchaseClick(effectiveOpt); }}
                        sx={{ alignSelf: 'flex-start' }}
                      >
                        View item at retailer
                      </Button>
                    );
                  })()}
                </>
              ) : (
                <>
              {targetOptions.length > 1 && (
                <Box>
                  <Typography variant="overline" color="text.secondary" sx={{ px: 0.5 }}>
                    Which one are you buying?
                  </Typography>
                  <Box
                    sx={{
                      display: 'flex',
                      flexDirection: 'row',
                      gap: 1.5,
                      overflowX: 'auto',
                      py: 1,
                      px: 0.5,
                      scrollbarWidth: 'none',
                      '&::-webkit-scrollbar': { display: 'none' },
                    }}
                  >
                    {targetOptions.map((opt) => {
                      const isSelected = (selectedOptionId ?? target) === opt.id;
                      return (
                        <Box
                          key={opt.id}
                          onClick={() => setSelectedOptionId(opt.id)}
                          sx={{
                            flexShrink: 0,
                            width: 140,
                            borderRadius: 2,
                            border: '2px solid',
                            borderColor: isSelected ? 'primary.main' : 'divider',
                            overflow: 'hidden',
                            display: 'flex',
                            flexDirection: 'column',
                            bgcolor: isSelected ? 'primary.50' : 'background.paper',
                            cursor: 'pointer',
                            transition: 'box-shadow 0.15s, border-color 0.15s, background-color 0.15s',
                            '&:hover': {
                              boxShadow: 2,
                              borderColor: 'primary.main',
                            },
                          }}
                        >
                          <Box
                            sx={{
                              width: '100%',
                              aspectRatio: '1 / 1',
                              bgcolor: opt.imageBgColor || '#ffffff',
                              overflow: 'hidden',
                              display: 'flex',
                              alignItems: 'center',
                              justifyContent: 'center',
                              position: 'relative',
                            }}
                          >
                            {opt.imageUrl ? (
                              <Box
                                component="img"
                                src={opt.imageUrl}
                                alt={opt.title}
                                sx={{ width: '100%', height: '100%', objectFit: 'contain' }}
                              />
                            ) : (
                              <Typography variant="caption" color="text.disabled">No image</Typography>
                            )}
                            {isSelected && (
                              <Box
                                sx={{
                                  position: 'absolute',
                                  top: 4,
                                  right: 4,
                                  width: 20,
                                  height: 20,
                                  borderRadius: '50%',
                                  bgcolor: 'primary.main',
                                  display: 'flex',
                                  alignItems: 'center',
                                  justifyContent: 'center',
                                }}
                              >
                                <Typography sx={{ color: '#fff', fontSize: 12, lineHeight: 1 }}>✓</Typography>
                              </Box>
                            )}
                          </Box>
                          <Box sx={{ p: 1 }}>
                            <Typography
                              variant="caption"
                              sx={{
                                fontWeight: 600,
                                display: '-webkit-box',
                                WebkitLineClamp: 2,
                                WebkitBoxOrient: 'vertical',
                                overflow: 'hidden',
                                color: 'text.primary',
                              }}
                            >
                              {opt.title}
                            </Typography>
                          </Box>
                        </Box>
                      );
                    })}
                  </Box>
                </Box>
              )}

              {targetRootItem && !targetRootItem.quantityUnlimited && targetRemaining > 1 && (
                <Typography variant="body2" color="text.secondary">
                  {targetRemaining} remaining out of {targetRootItem.quantity || 1}.
                </Typography>
              )}

              <Stack spacing={2} sx={{ minWidth: 0 }}>
                {targetRootItem?.quantityUnlimited ? (
                  <TextField
                    label="Quantity you're buying"
                    type="number"
                    inputProps={{ min: 1, step: 1 }}
                    value={reserveQty}
                    onChange={(e) => setReserveQty(e.target.value)}
                  />
                ) : targetRemaining > 1 ? (
                  <Select value={reserveQtyMode} onChange={(e) => setReserveQtyMode(e.target.value as 'one' | 'all')}>
                    <MenuItem value="one">Add 1</MenuItem>
                    <MenuItem value="all">Add all remaining ({targetRemaining})</MenuItem>
                  </Select>
                ) : null}
              </Stack>
              <Typography variant="caption" color="text.secondary">
                Verified as <strong>{meQ.data?.email}</strong>. The parents will see this email so they can follow up.
              </Typography>
              {error && <Alert severity="error">{error}</Alert>}
                </>
              )}
            </Stack>
          </DialogContent>
          <DialogActions sx={{ px: 3, py: 2, borderTop: 1, borderColor: 'divider' }}>
            {reservedId ? (
              <>
                <Button
                  color="inherit"
                  onClick={() => cancelRsvM.mutate(reservedId)}
                  disabled={cancelRsvM.isPending}
                >
                  Remove from cart
                </Button>
                <Button
                  variant="contained"
                  onClick={() => { resetReserveDialog(); setCartOpen(true); }}
                  disabled={cancelRsvM.isPending}
                  sx={{ color: '#fff' }}
                >
                  View cart
                </Button>
              </>
            ) : (
              <>
                <Button onClick={() => setTarget(null)}>Cancel</Button>
                <Button
                  variant="contained"
                  onClick={() => reserveM.mutate()}
                  disabled={(!targetRootItem?.quantityUnlimited && targetRemaining <= 0) || reserveM.isPending}
                  sx={{ color: '#fff' }}
                  startIcon={<ShoppingCartOutlinedIcon />}
                >
                  {reserveM.isPending ? 'Adding\u2026' : 'Add to cart'}
                </Button>
              </>
            )}
          </DialogActions>
        </Dialog>

        <Snackbar
          open={!!snack}
          autoHideDuration={4000}
          onClose={() => setSnack(null)}
          message={snack?.msg ?? ''}
          action={
            snack?.cart ? (
              <Button
                color="inherit"
                size="small"
                onClick={() => {
                  setSnack(null);
                  setCartOpen(true);
                }}
              >
                View cart
              </Button>
            ) : undefined
          }
        />

        <CartDrawer
          open={cartOpen}
          onClose={() => setCartOpen(false)}
          lines={cartLines}
          buyerEmail={meQ.data?.email ?? ''}
          onRemove={(id) => cancelRsvM.mutate(id)}
          removingId={cancelRsvM.isPending ? (cancelRsvM.variables as string) : null}
          onContinueShopping={() => setCartOpen(false)}
          paymentMethods={(paymentMethodsQ.data as PaymentMethod[] | undefined) ?? []}
          onCreatePaymentIntent={async (methodId) => {
            const created = await publicPayments.createIntent(slug, methodId);
            qc.invalidateQueries({ queryKey: ['public', slug] });
            return created;
          }}
          onClaimPayment={async (paymentId, msg) => {
            await publicPayments.claim(paymentId, slug, msg);
            qc.invalidateQueries({ queryKey: ['public', slug] });
          }}
          onCancelPaymentIntent={async (paymentId) => {
            await publicPayments.cancel(paymentId, slug);
            qc.invalidateQueries({ queryKey: ['public', slug] });
          }}
          onViewProduct={(line) => {
            if (!line.productUrl) return;
            window.open(line.productUrl, '_blank', 'noopener,noreferrer');
            trackPurchaseClick({ id: line.itemId, productUrl: line.productUrl, retailer: line.retailer });
          }}
          viewerCurrency={viewerCurrency}
          rates={ratesQ.data}
          heldCart={heldCart}
          heldCartMethod={heldCartMethod}
          heldCartLines={heldCartLines}
        />

        {(cartLines.length > 0 || !!heldCart) && (
          <Fab
            color="primary"
            aria-label="Open cart"
            onClick={() => setCartOpen(true)}
            sx={{
              position: 'fixed',
              bottom: { xs: 20, md: 32 },
              right: { xs: 20, md: 32 },
              color: '#fff',
              zIndex: (t) => t.zIndex.drawer - 1,
            }}
          >
            <Badge badgeContent={cartLines.length + (heldCart ? heldCartLines.length : 0)} color="error" overlap="circular">
              <ShoppingCartOutlinedIcon />
            </Badge>
          </Fab>
        )}
      </Container>
    </Box>
  );
}

