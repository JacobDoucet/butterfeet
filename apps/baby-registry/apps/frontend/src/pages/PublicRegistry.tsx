import { useState, useEffect, useRef } from 'react';
import { useParams } from 'react-router-dom';
import {
  Container,
  Typography,
  Card,
  CardContent,
  CardMedia,
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
} from '@mui/material';
import MarkEmailReadIcon from '@mui/icons-material/MarkEmailRead';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { pub, buyer } from '../api';

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

  const [target, setTarget] = useState<string | null>(null);
  const [name, setName] = useState('');
  const [anon, setAnon] = useState(false);
  const [message, setMessage] = useState('');
  const [reserveQtyMode, setReserveQtyMode] = useState<'one' | 'all'>('one');
  const [reserveQty, setReserveQty] = useState('1');
  const [selectedOptionId, setSelectedOptionId] = useState<string | null>(null);
  const [categoryFilter, setCategoryFilter] = useState<string>('__all__');
  const [celebrate, setCelebrate] = useState<{ title: string } | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [snack, setSnack] = useState<string | null>(null);
  const [accessNote, setAccessNote] = useState('');
  const [accessRequested, setAccessRequested] = useState(false);
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
    height: '80vh',
    maxHeight: '80vh',
    display: 'flex',
    flexDirection: 'column',
  } as const;

  // Reset selected option whenever a new card is opened
  useEffect(() => {
    if (target) setSelectedOptionId(null);
  }, [target]);

  // Set page title from registry once loaded
  useEffect(() => {
    if (regQ.data?.title) {
      const prev = document.title;
      document.title = `${regQ.data.title} · baby registry`;
      return () => { document.title = prev; };
    }
  }, [regQ.data?.title]);

  const reserveM = useMutation({
    mutationFn: async () => {
      if (!target) return;
      const effectiveId = selectedOptionId ?? target;
      const item = regQ.data?.items.find((it) => it.id === effectiveId);
      const rootItem = regQ.data?.items.find((it) => it.id === target);
      const qty = rootItem?.quantityUnlimited
        ? Math.max(1, parseInt(reserveQty || '1', 10) || 1)
        : reserveQtyMode === 'all'
        ? Math.max(1, (rootItem?.quantity || 1) - (rootItem?.reserved || 0))
        : 1;
      return pub.reserve(effectiveId, { reserverName: name, isAnonymous: anon, message, quantity: qty });
    },
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['public', slug] });
      const rootItem = regQ.data?.items.find((it) => it.id === target);
      const effectiveId = selectedOptionId ?? target;
      const purchasedItem = regQ.data?.items.find((it) => it.id === effectiveId) ?? rootItem;
      setCelebrate({ title: purchasedItem?.title ?? 'this gift' });
      setTarget(null); setSelectedOptionId(null); setName(''); setAnon(false); setMessage(''); setError(null); setReserveQtyMode('one'); setReserveQty('1');
    },
    onError: (err) => setError((err as Error).message),
  });

  const requestAddressM = useMutation({
    mutationFn: () =>
      pub.requestAddress({
        slug,
        itemId: target ?? undefined,
        name: name.trim() || undefined,
        note: accessNote.trim() || undefined,
      }),
    onSuccess: () => {
      setAccessRequested(true);
      setAccessNote('');
      setSnack("Request sent \u2014 we'll email you the address once the parents approve.");
    },
    onError: (err) => setError((err as Error).message),
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
  const isClaimed = (rootId: string) => {
    const r = itemById[rootId];
    if (!r) return false;
    if (r.quantityUnlimited) return false;
    return Math.max(0, (r.quantity || 1) - (r.reserved || 0)) === 0;
  };
  const sortedTopLevelItems = [...topLevelItems].sort((a, b) =>
    (a.title || '').localeCompare(b.title || '', undefined, { sensitivity: 'base' }),
  );
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
    const optionCount = options.length;
    const lowStock = !root.quantityUnlimited && remaining > 0 && remaining < (root.quantity || 1);
    const swatchImages = options
      .filter((o) => !!o.imageUrl)
      .slice(0, 4)
      .map((o) => o.imageUrl as string);

    return (
      <Grid item xs={12} sm={6} md={4} lg={3} key={it.id}>
        <Card
          elevation={0}
          sx={{
            height: '100%',
            display: 'flex',
            flexDirection: 'column',
            cursor: claimed ? 'default' : 'pointer',
            border: '1px solid',
            borderColor: 'divider',
            borderRadius: 3,
            overflow: 'hidden',
            position: 'relative',
            bgcolor: 'background.paper',
            transition: 'transform 0.18s ease, box-shadow 0.18s ease, border-color 0.18s ease',
            ...(claimed
              ? {}
              : {
                  '&:hover': {
                    transform: 'translateY(-3px)',
                    boxShadow: '0 12px 28px rgba(0,0,0,0.08)',
                    borderColor: 'primary.light',
                  },
                  '&:focus-visible': {
                    outline: '2px solid',
                    outlineColor: 'primary.main',
                    outlineOffset: 2,
                  },
                }),
          }}
          onClick={() => {
            if (!claimed) setTarget(it.id);
          }}
          role={claimed ? undefined : 'button'}
          tabIndex={claimed ? -1 : 0}
          onKeyDown={(e) => {
            if (claimed) return;
            if (e.key === 'Enter' || e.key === ' ') {
              e.preventDefault();
              setTarget(it.id);
            }
          }}
        >
          <Box sx={{ position: 'relative', bgcolor: it.imageBgColor || '#ffffff' }}>
            {it.imageUrl ? (
              <CardMedia
                component="img"
                image={it.imageUrl}
                sx={{
                  aspectRatio: '1',
                  objectFit: 'contain',
                  p: 2,
                  filter: claimed ? 'grayscale(0.4)' : 'none',
                  transition: 'filter 0.2s',
                }}
              />
            ) : (
              <Box
                sx={{
                  aspectRatio: '1',
                  display: 'flex',
                  alignItems: 'center',
                  justifyContent: 'center',
                  color: 'text.disabled',
                }}
              >
                <Typography variant="caption">No image</Typography>
              </Box>
            )}
            {optionCount > 1 && (
              <Box
                sx={{
                  position: 'absolute',
                  top: 10,
                  left: 10,
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
            )}
            {claimed && (
              <Box
                sx={{
                  position: 'absolute',
                  top: 12,
                  right: 12,
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
            )}
            {!claimed && lowStock && (
              <Box
                sx={{
                  position: 'absolute',
                  top: 12,
                  right: 12,
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
            )}
          </Box>
          <CardContent sx={{ flexGrow: 1, display: 'flex', flexDirection: 'column', p: 2.5 }}>
            <Typography
              variant="subtitle1"
              sx={{
                fontWeight: 600,
                lineHeight: 1.35,
                display: '-webkit-box',
                WebkitLineClamp: 2,
                WebkitBoxOrient: 'vertical',
                overflow: 'hidden',
                mb: 1,
                color: claimed ? 'text.secondary' : 'text.primary',
              }}
            >
              {it.title}
            </Typography>
            <Box sx={{ flexGrow: 1 }} />
            <Stack direction="row" alignItems="center" justifyContent="space-between" spacing={1}>
              {it.source ? (
                <Typography variant="caption" color="text.secondary" sx={{ textTransform: 'uppercase', letterSpacing: 0.5, fontWeight: 600 }}>
                  {it.source}
                </Typography>
              ) : (
                <span />
              )}
              {!claimed && (
                <Typography variant="caption" color="primary.main" sx={{ fontWeight: 600 }}>
                  Gift this →
                </Typography>
              )}
            </Stack>
          </CardContent>
        </Card>
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
  const hasShippingAddress = Boolean(
    reg.shippingRecipientName || reg.shippingLine1 || reg.shippingCity || reg.shippingRegion || reg.shippingPostalCode || reg.shippingCountry,
  );

  return (
    <Box
      sx={{
        minHeight: '100%',
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
              <Typography variant="subtitle1" color="text.secondary">
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

      <Container maxWidth="lg" sx={{ py: { xs: 4, md: 6 } }}>
        {categories.length > 0 && (
          <Box
            sx={{
              mb: 4,
              position: 'sticky',
              top: { xs: 104, sm: 108 },
              zIndex: 3,
              borderRadius: 0,
              px: 1,
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
                    <Typography variant="h5" sx={{ fontWeight: 600, letterSpacing: '-0.01em' }}>
                      {cat}
                    </Typography>
                    {hasMore && (
                      <Button
                        size="small"
                        onClick={() => setCategoryFilter(cat)}
                        sx={{ textTransform: 'none', fontWeight: 600 }}
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
          open={!!target}
          onClose={() => setTarget(null)}
          scroll="paper"
          fullWidth
          maxWidth="md"
          TransitionProps={{
            onExited: () => {
              targetSnapshotRef.current = null;
              setAccessRequested(false);
              setAccessNote('');
            },
          }}
          PaperProps={{ sx: { ...modalPaperSx, width: 'min(960px, calc(100vw - 32px))' } }}
        >
          <DialogTitle sx={{ pb: 1 }}>
            <Typography variant="overline" color="text.secondary" sx={{ display: 'block', lineHeight: 1, mb: 0.5 }}>
              Get this gift
            </Typography>
            <Typography variant="h6" sx={{ fontWeight: 700, lineHeight: 1.2 }}>
              {targetRootItem?.title ?? 'Item details'}
            </Typography>
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
              {targetOptions.length === 1 && targetOptions[0] && (
                <Card variant="outlined" sx={{ borderRadius: 2 }}>
                  <Stack direction={{ xs: 'column', sm: 'row' }} spacing={0} sx={{ overflow: 'hidden' }}>
                    <Box
                      component={targetOptions[0].productUrl ? 'a' : 'div'}
                      {...(targetOptions[0].productUrl ? { href: targetOptions[0].productUrl, target: '_blank', rel: 'noreferrer' } : {})}
                      sx={{
                        width: { xs: '100%', sm: 160 },
                        flexShrink: 0,
                        bgcolor: targetOptions[0].imageBgColor || '#ffffff',
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'center',
                        aspectRatio: { xs: '16 / 9', sm: '1 / 1' },
                        overflow: 'hidden',
                        textDecoration: 'none',
                        ...(targetOptions[0].productUrl && { cursor: 'pointer', '&:hover': { opacity: 0.85 } }),
                      }}
                    >
                      {targetOptions[0].imageUrl ? (
                        <Box
                          component="img"
                          src={targetOptions[0].imageUrl}
                          alt={targetOptions[0].title}
                          sx={{ width: '100%', height: '100%', objectFit: 'contain' }}
                        />
                      ) : (
                        <Typography variant="caption" color="text.disabled">No image</Typography>
                      )}
                    </Box>
                    <Stack spacing={1} sx={{ p: 2, minWidth: 0, flex: 1 }}>
                      <Typography
                        variant="subtitle1"
                        component={targetOptions[0].productUrl ? 'a' : 'span'}
                        {...(targetOptions[0].productUrl ? { href: targetOptions[0].productUrl, target: '_blank', rel: 'noreferrer' } : {})}
                        sx={{
                          fontWeight: 700,
                          lineHeight: 1.3,
                          textDecoration: targetOptions[0].productUrl ? 'underline' : 'none',
                          color: targetOptions[0].productUrl ? 'primary.main' : 'text.primary',
                          '&:hover': targetOptions[0].productUrl ? { opacity: 0.8 } : {},
                        }}
                      >
                        {targetOptions[0].title}
                      </Typography>
                      {targetOptions[0].description && (
                        <Typography variant="body2" color="text.secondary" sx={{ whiteSpace: 'pre-wrap' }}>
                          {targetOptions[0].description}
                        </Typography>
                      )}
                    </Stack>
                  </Stack>
                </Card>
              )}

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
                            {opt.productUrl && (
                              <Box
                                component="a"
                                href={opt.productUrl}
                                target="_blank"
                                rel="noreferrer"
                                onClick={(e: React.MouseEvent) => e.stopPropagation()}
                                sx={{ textDecoration: 'none' }}
                              >
                                <Typography variant="caption" color="primary" sx={{ display: 'block', mt: 0.5 }}>
                                  Shop →
                                </Typography>
                              </Box>
                            )}
                          </Box>
                        </Box>
                      );
                    })}
                  </Box>
                </Box>
              )}

              {targetRootItem && (
                <Typography variant="body2" color="text.secondary">
                  {targetRootItem.quantityUnlimited
                    ? 'No reservation limit for this item.'
                    : `${targetRemaining} remaining out of ${targetRootItem.quantity || 1}.`}
                </Typography>
              )}
              <Card variant="outlined" sx={{ borderRadius: 2 }}>
                <CardContent sx={{ minWidth: 0 }}>
                  <Stack spacing={1}>
                    <Typography variant="overline" color="text.secondary">How it works</Typography>
                    <Typography variant="body2">
                      1.{' '}
                      {(() => {
                        const effectiveOpt = itemById[selectedOptionId ?? target ?? ''] ?? targetOptions[0];
                        return effectiveOpt?.productUrl ? (
                          <>
                            <Box
                              component="a"
                              href={effectiveOpt.productUrl}
                              target="_blank"
                              rel="noreferrer"
                              sx={{ color: 'primary.main', textDecoration: 'underline' }}
                            >
                              Open the product page
                            </Box>
                            {' '}and complete your purchase.
                          </>
                        ) : (
                          'Open the product page and complete your purchase.'
                        );
                      })()}
                    </Typography>
                    <Typography variant="body2">2. Ship it to the delivery address below.</Typography>
                    <Typography variant="body2">3. Click <strong>I've bought this</strong> when you’re done.</Typography>
                  </Stack>
                </CardContent>
              </Card>

              <Card variant="outlined" sx={{ borderRadius: 2 }}>
                <CardContent sx={{ minWidth: 0 }}>
                  <Stack spacing={2} sx={{ minWidth: 0 }}>
                    <Typography variant="overline" color="text.secondary">Purchase details</Typography>
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
                        <MenuItem value="one">Buy 1</MenuItem>
                        <MenuItem value="all">Buy all remaining ({targetRemaining})</MenuItem>
                      </Select>
                    ) : null}

                    {hasShippingAddress ? (
                      <Box sx={{ p: 2, borderRadius: 1, bgcolor: 'action.hover' }}>
                        <Typography variant="overline" color="text.secondary">Delivery address</Typography>
                        {reg.shippingRecipientName && <Typography>{reg.shippingRecipientName}</Typography>}
                        {reg.shippingLine1 && <Typography>{reg.shippingLine1}</Typography>}
                        {reg.shippingLine2 && <Typography>{reg.shippingLine2}</Typography>}
                        {(reg.shippingCity || reg.shippingRegion || reg.shippingPostalCode) && (
                          <Typography>{[reg.shippingCity, reg.shippingRegion, reg.shippingPostalCode].filter(Boolean).join(' ')}</Typography>
                        )}
                        {reg.shippingCountry && <Typography>{reg.shippingCountry}</Typography>}
                        {reg.shippingDeliveryNotes && (
                          <Typography variant="body2" color="text.secondary" sx={{ mt: 1, whiteSpace: 'pre-wrap' }}>
                            Note: {reg.shippingDeliveryNotes}
                          </Typography>
                        )}
                      </Box>
                    ) : (
                      <Alert severity="warning" sx={{ '& .MuiAlert-message': { width: '100%' } }}>
                        <Stack spacing={1.5}>
                          <Box>
                            Delivery address is protected and not shown on the public registry page.
                            You can ask the parents to send it to you privately.
                          </Box>
                          {accessRequested ? (
                            <Box sx={{ fontSize: 14, color: 'text.secondary' }}>
                              Request sent. We'll email <strong>{meQ.data?.email}</strong> a private link
                              as soon as the parents approve.
                            </Box>
                          ) : (
                            <>
                              <TextField
                                size="small"
                                label="Add a short note (optional)"
                                placeholder="Example: For the Pottery Barn glider I'm shipping next week."
                                value={accessNote}
                                onChange={(e) => setAccessNote(e.target.value)}
                                multiline
                                minRows={2}
                                fullWidth
                              />
                              <Box>
                                <Button
                                  variant="contained"
                                  size="small"
                                  onClick={() => requestAddressM.mutate()}
                                  disabled={requestAddressM.isPending}
                                >
                                  {requestAddressM.isPending ? 'Sending\u2026' : 'Request shipping address'}
                                </Button>
                              </Box>
                            </>
                          )}
                        </Stack>
                      </Alert>
                    )}

                    <TextField label="Your name" value={name} onChange={(e) => setName(e.target.value)} disabled={anon} />
                    <FormControlLabel
                      control={<Checkbox checked={anon} onChange={(e) => setAnon(e.target.checked)} />}
                      label="Keep me anonymous"
                    />
                    <TextField
                      label="Message to the parents (optional)"
                      placeholder="Example: Ordered from Amazon, arrives next Tuesday."
                      multiline
                      minRows={3}
                      value={message}
                      onChange={(e) => setMessage(e.target.value)}
                    />
                  </Stack>
                </CardContent>
              </Card>
              <Typography variant="caption" color="text.secondary">
                Verified as <strong>{meQ.data?.email}</strong>. The parents will see this email so they can follow up.
              </Typography>
              {error && <Alert severity="error">{error}</Alert>}
            </Stack>
          </DialogContent>
          <DialogActions sx={{ px: 3, py: 2, borderTop: 1, borderColor: 'divider' }}>
            <Button onClick={() => setTarget(null)}>Cancel</Button>
            <Button
              variant="contained"
              onClick={() => reserveM.mutate()}
              disabled={(!anon && !name.trim()) || (!targetRootItem?.quantityUnlimited && targetRemaining <= 0)}
            >
              I've bought this
            </Button>
          </DialogActions>
        </Dialog>

        <Snackbar open={!!snack} autoHideDuration={4000} onClose={() => setSnack(null)} message={snack ?? ''} />

        <Dialog
          open={!!celebrate}
          onClose={() => setCelebrate(null)}
          maxWidth="xs"
          fullWidth
          PaperProps={{
            sx: {
              borderRadius: 4,
              overflow: 'hidden',
              position: 'relative',
              textAlign: 'center',
            },
          }}
        >
          {celebrate && (
            <>
              <Box
                sx={{
                  position: 'absolute',
                  inset: 0,
                  pointerEvents: 'none',
                  overflow: 'hidden',
                  '@keyframes confetti-fall': {
                    '0%': { transform: 'translateY(-20px) rotate(0deg)', opacity: 1 },
                    '100%': { transform: 'translateY(420px) rotate(720deg)', opacity: 0 },
                  },
                }}
              >
                {Array.from({ length: 28 }).map((_, i) => {
                  const colors = ['#f4a4a4', '#ffd56b', '#9cd6c1', '#c1b2e0', '#f49a78', '#7fb8e3'];
                  const left = (i * 37) % 100;
                  const delay = (i % 7) * 0.15;
                  const duration = 2 + ((i * 13) % 10) / 10;
                  const color = colors[i % colors.length];
                  const size = 6 + (i % 4) * 2;
                  return (
                    <Box
                      key={i}
                      sx={{
                        position: 'absolute',
                        left: `${left}%`,
                        top: -10,
                        width: size,
                        height: size * 1.6,
                        bgcolor: color,
                        borderRadius: '2px',
                        animation: `confetti-fall ${duration}s ${delay}s ease-in forwards`,
                      }}
                    />
                  );
                })}
              </Box>
              <Box sx={{ pt: 5, pb: 4, px: 4, position: 'relative' }}>
                <Typography sx={{ fontSize: 56, lineHeight: 1, mb: 2 }}>🎁</Typography>
                <Typography variant="h5" sx={{ fontWeight: 700, mb: 1 }}>
                  Thank you!
                </Typography>
                <Typography variant="body1" color="text.secondary" sx={{ mb: 3 }}>
                  Your gift of <strong>{celebrate.title}</strong> means the world. The parents will be over the moon.
                </Typography>
                <Button variant="contained" onClick={() => setCelebrate(null)} sx={{ borderRadius: 5, px: 4 }}>
                  Close
                </Button>
              </Box>
            </>
          )}
        </Dialog>
      </Container>
    </Box>
  );
}

function BuyerVerifyGate({ slug, onVerified }: { slug: string; onVerified: () => void }) {
  const [step, setStep] = useState<'email' | 'code'>('email');
  const [email, setEmail] = useState('');
  const [code, setCode] = useState('');
  const [err, setErr] = useState<string | null>(null);

  const requestM = useMutation({
    mutationFn: () => buyer.request(slug, email.trim()),
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
                label="Your email"
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                autoFocus
                fullWidth
              />
              {err && <Alert severity="error">{err}</Alert>}
              <Button
                variant="contained"
                size="large"
                onClick={() => requestM.mutate()}
                disabled={!email.trim() || requestM.isPending}
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
