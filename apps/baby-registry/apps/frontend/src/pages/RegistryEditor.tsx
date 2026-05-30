import { useRef, useState } from 'react';
import { Link, useParams } from 'react-router-dom';
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
  IconButton,
  CircularProgress,
  Chip,
  Tabs,
  Tab,
  Box,
  FormControlLabel,
  Checkbox,
  Autocomplete,
} from '@mui/material';
import DeleteIcon from '@mui/icons-material/DeleteOutline';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import Inventory2Icon from '@mui/icons-material/Inventory2Outlined';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import {
  Accordion,
  AccordionSummary,
  AccordionDetails,
  Divider,
  MenuItem,
  Select,
  Tooltip,
} from '@mui/material';
import { registries, items, scrape, reservations, type RegistryItem, type Registry, type Reservation, type ReservationStatus } from '../api';
import PrivacyPanel from './PrivacyPanel';

type DeleteTarget =
  | { kind: 'item'; id: string; title: string }
  | { kind: 'reservation'; id: string; title: string };

export default function RegistryEditor() {
  const { slug = '' } = useParams();
  const qc = useQueryClient();
  const [activeTab, setActiveTab] = useState<'items' | 'shipping' | 'access'>('items');

  const regsQ = useQuery({
    queryKey: ['registries'],
    queryFn: () => registries.list(),
  });
  const reg: Registry | undefined = regsQ.data?.data.find((r) => r.slug === slug);

  const itemsQ = useQuery({
    queryKey: ['items', reg?.id],
    queryFn: () => items.listForRegistry(reg!.id),
    enabled: !!reg,
  });

  const reservationsQ = useQuery({
    queryKey: ['reservations', reg?.id],
    queryFn: () => reservations.listForRegistry(reg!.id),
    enabled: !!reg,
  });

  const setStatusM = useMutation({
    mutationFn: ({ id, status }: { id: string; status: ReservationStatus }) =>
      reservations.setStatus(id, status),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['reservations', reg?.id] });
      qc.invalidateQueries({ queryKey: ['items', reg?.id] });
    },
  });

  const deleteReservationM = useMutation({
    mutationFn: (id: string) => reservations.remove(id),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['reservations', reg?.id] });
      qc.invalidateQueries({ queryKey: ['items', reg?.id] });
    },
  });

  const [open, setOpen] = useState(false);
  const [editOpen, setEditOpen] = useState(false);
  const [editTab, setEditTab] = useState<'details' | 'substitutes' | 'fulfillment'>('details');
  const [editingId, setEditingId] = useState<string | null>(null);
  const [fulfillmentStatus, setFulfillmentStatus] = useState<ReservationStatus>('Purchased');
  const [deleteTarget, setDeleteTarget] = useState<DeleteTarget | null>(null);
  const [url, setUrl] = useState('');
  const [scraping, setScraping] = useState(false);
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [imageUrl, setImageUrl] = useState('');
  const [source, setSource] = useState('');
  const [quantity, setQuantity] = useState('1');
  const [quantityUnlimited, setQuantityUnlimited] = useState(false);
  const [category, setCategory] = useState('');
  const [noSubstitutes, setNoSubstitutes] = useState(false);
  const [parentItemId, setParentItemId] = useState('');
  const [notes, setNotes] = useState('');
  const [error, setError] = useState<string | null>(null);
  const modalPaperSx = {
    height: '80vh',
    maxHeight: '80vh',
    display: 'flex',
    flexDirection: 'column',
  } as const;

  const suppressEditResetRef = useRef(false);

  const reset = () => {
    setUrl('');
    setTitle('');
    setDescription('');
    setImageUrl('');
    setSource('');
    setQuantity('1');
    setQuantityUnlimited(false);
    setCategory('');
    setNoSubstitutes(false);
    setParentItemId('');
    setNotes('');
    setEditingId(null);
    setError(null);
  };

  const beginEdit = (it: RegistryItem) => {
    setEditTab('details');
    setEditingId(it.id);
    setUrl(it.productUrl || '');
    setTitle(it.title || '');
    setDescription(it.description || '');
    setImageUrl(it.imageUrl || '');
    setSource(it.source || '');
    setQuantity(String(it.quantity || 1));
    setQuantityUnlimited(!!it.quantityUnlimited);
    setCategory(it.category || '');
    setNoSubstitutes(!!it.noSubstitutes);
    setParentItemId(it.parentItemId || '');
    setNotes(it.notes || '');
    setFulfillmentStatus('Purchased');
    setError(null);
    setEditOpen(true);
  };

  const beginAddAlternative = (rootId: string) => {
    suppressEditResetRef.current = true;
    setEditOpen(false);
    reset();
    setParentItemId(rootId);
    setOpen(true);
  };

  const doScrape = async () => {
    setError(null);
    setScraping(true);
    try {
      const r = await scrape.url(url);
      setTitle(r.title || '');
      setImageUrl(r.imageUrl || '');
      setSource(r.source || 'Other');
    } catch (err) {
      setError((err as Error).message);
    } finally {
      setScraping(false);
    }
  };

  const createItemM = useMutation({
    mutationFn: async () => {
      const qty = quantityUnlimited ? 0 : Math.max(1, parseInt(quantity || '1', 10) || 1);
      return items.create({
        registryId: reg!.id,
        title,
        description,
        imageUrl,
        productUrl: url,
        source,
        quantity: qty,
        quantityUnlimited,
        category: category.trim(),
        noSubstitutes,
        parentItemId: parentItemId || undefined,
        notes,
      });
    },
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['items', reg?.id] });
      setOpen(false);
    },
    onError: (err) => setError((err as Error).message),
  });

  const deleteM = useMutation({
    mutationFn: (id: string) => items.remove(id),
    onSuccess: () => qc.invalidateQueries({ queryKey: ['items', reg?.id] }),
  });

  const updateItemM = useMutation({
    mutationFn: async () => {
      if (!editingId) return;
      const qty = quantityUnlimited ? 0 : Math.max(1, parseInt(quantity || '1', 10) || 1);
      return items.update(editingId, {
        title,
        description,
        imageUrl,
        productUrl: url,
        source,
        quantity: qty,
        quantityUnlimited,
        category: category.trim(),
        noSubstitutes,
        parentItemId: parentItemId || undefined,
        notes,
      });
    },
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['items', reg?.id] });
      setEditOpen(false);
    },
    onError: (err) => setError((err as Error).message),
  });

  const markPurchasedM = useMutation({
    mutationFn: async ({
      item,
      activeReservations,
      status,
    }: {
      item: RegistryItem;
      activeReservations: Reservation[];
      status: ReservationStatus;
    }) => {
      if (activeReservations.length > 0) {
        await Promise.all(activeReservations.map((r) => reservations.setStatus(r.id, status)));
        return;
      }
      await reservations.create({
        itemId: item.id,
        registryId: reg!.id,
        reserverName: 'Owner',
        isAnonymous: false,
        message: 'Set by owner',
        quantity: Math.max(1, item.quantity || 1),
        status,
      });
    },
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['reservations', reg?.id] });
      qc.invalidateQueries({ queryKey: ['items', reg?.id] });
    },
    onError: (err) => setError((err as Error).message),
  });

  const confirmDelete = () => {
    if (!deleteTarget) return;
    if (deleteTarget.kind === 'item') {
      deleteM.mutate(deleteTarget.id);
    } else {
      deleteReservationM.mutate(deleteTarget.id);
    }
    setDeleteTarget(null);
  };

  if (regsQ.isLoading) return null;
  if (!reg) return (
    <Container sx={{ py: 6 }}>
      <Alert severity="warning">Registry not found.</Alert>
    </Container>
  );

  const list: RegistryItem[] = itemsQ.data?.data ?? [];
  const allReservations: Reservation[] = reservationsQ.data?.data ?? [];
  const reservationsByItem = allReservations.reduce<Record<string, Reservation[]>>((acc, r) => {
    (acc[r.itemId] ??= []).push(r);
    return acc;
  }, {});
  const itemById = list.reduce<Record<string, RegistryItem>>((acc, it) => {
    acc[it.id] = it;
    return acc;
  }, {});
  const categoryOptions = Array.from(new Set(list.map((it) => (it.category || '').trim()).filter(Boolean))).sort((a, b) =>
    a.localeCompare(b),
  );
  const alternativesByRootId = list.reduce<Record<string, RegistryItem[]>>((acc, it) => {
    if (it.parentItemId && itemById[it.parentItemId]) {
      (acc[it.parentItemId] ??= []).push(it);
    }
    return acc;
  }, {});
  const groupRootId = (it: RegistryItem) => (it.parentItemId && itemById[it.parentItemId] ? it.parentItemId : it.id);
  const optionsByRootId = list.reduce<Record<string, number>>((acc, it) => {
    const rootId = groupRootId(it);
    acc[rootId] = (acc[rootId] ?? 0) + 1;
    return acc;
  }, {});
  const topLevelItems = list.filter((it) => !it.parentItemId || !itemById[it.parentItemId]);
  const editingItem = list.find((it) => it.id === editingId) ?? null;
  const editRootId = editingItem ? groupRootId(editingItem) : null;
  const editRootItem = editRootId ? itemById[editRootId] ?? null : null;
  const editAlternatives = editRootId
    ? list.filter((x) => x.parentItemId === editRootId)
    : [];
  const editOptionCount = editRootId ? optionsByRootId[editRootId] ?? 1 : 1;
  const editOptionItems = editRootId
    ? list.filter((it) => it.id === editRootId || it.parentItemId === editRootId)
    : editingItem ? [editingItem] : [];
  const editGroupActiveReservations = editOptionItems
    .flatMap((it) => reservationsByItem[it.id] ?? [])
    .filter((r) => r.status !== 'Cancelled');
  const editActiveReservations = (editingItem ? reservationsByItem[editingItem.id] ?? [] : [])
    .filter((r) => r.status !== 'Cancelled');

  return (
    <Container maxWidth="lg" sx={{ py: 6 }}>
      <Stack direction="row" alignItems="center" sx={{ mb: 4 }}>
        <Stack sx={{ flexGrow: 1 }}>
          <Typography variant="h4">{reg.title}</Typography>
          <Typography color="text.secondary" component={Link} to={`/r/${reg.slug}`}>
            View public page → /r/{reg.slug}
          </Typography>
        </Stack>
        {activeTab === 'items' && (
          <Button variant="contained" onClick={() => { reset(); setOpen(true); }}>Add item</Button>
        )}
      </Stack>

      <Box sx={{ borderBottom: 1, borderColor: 'divider', mb: 3 }}>
        <Tabs
          value={activeTab}
          onChange={(_, value) => setActiveTab(value)}
          variant="scrollable"
          scrollButtons="auto"
        >
          <Tab value="items" label="Items" />
          <Tab value="shipping" label="Shipping info" />
          <Tab value="access" label="User access" />
        </Tabs>
      </Box>

      {activeTab === 'shipping' && <PrivacyPanel reg={reg} section="shipping" />}
      {activeTab === 'access' && <PrivacyPanel reg={reg} section="access" />}

      {activeTab === 'items' && <Grid container spacing={2}>
        {topLevelItems.map((it) => {
          const optionItems = [it, ...(alternativesByRootId[it.id] ?? [])];
          const itemReservations = optionItems.flatMap((opt) => reservationsByItem[opt.id] ?? []);
          const activeCount = itemReservations
            .filter((r) => r.status !== 'Cancelled')
            .reduce((sum, r) => sum + (r.quantity ?? 1), 0);
          const requested = it.quantity ?? 1;
          const isUnlimited = !!it.quantityUnlimited;
          const fulfilled = !isUnlimited && activeCount >= requested;
          const optionCount = optionItems.length;
          return (
            <Grid item xs={12} sm={6} md={4} key={it.id}>
              <Card
                sx={{ cursor: 'pointer' }}
                onClick={() => beginEdit(it)}
                role="button"
                tabIndex={0}
                onKeyDown={(e) => {
                  if (e.key === 'Enter' || e.key === ' ') {
                    e.preventDefault();
                    beginEdit(it);
                  }
                }}
              >
                {it.imageUrl && (
                  <CardMedia component="img" image={it.imageUrl} sx={{ aspectRatio: '1', objectFit: 'contain', bgcolor: '#f4ede3' }} />
                )}
                <CardContent>
                  <Typography
                    variant="subtitle1"
                    sx={{
                      fontWeight: 600,
                      overflow: 'hidden',
                      textOverflow: 'ellipsis',
                      display: '-webkit-box',
                      WebkitLineClamp: 2,
                      WebkitBoxOrient: 'vertical',
                    }}
                  >
                    {it.title}
                  </Typography>
                  <Stack direction="row" spacing={1} sx={{ my: 1, flexWrap: 'wrap', rowGap: 1 }}>
                    {it.source && <Chip size="small" label={it.source} />}
                    {it.category && <Chip size="small" variant="outlined" label={it.category} />}
                    {optionCount > 1 && <Chip size="small" variant="outlined" label={`${optionCount} options`} />}
                    {it.noSubstitutes && <Chip size="small" variant="outlined" label="No substitutes" />}
                    <Chip
                      size="small"
                      color={fulfilled ? 'success' : activeCount > 0 ? 'warning' : 'default'}
                      variant={activeCount > 0 ? 'filled' : 'outlined'}
                      label={
                        isUnlimited
                          ? `${activeCount} / ∞ reserved`
                          : `${activeCount} / ${requested} reserved`
                      }
                    />
                  </Stack>
                  {itemReservations.length > 0 && (
                    <Accordion
                      disableGutters
                      elevation={0}
                      sx={{ mt: 2, bgcolor: 'transparent', '&:before': { display: 'none' } }}
                      onClick={(e) => e.stopPropagation()}
                      onKeyDown={(e) => e.stopPropagation()}
                    >
                      <AccordionSummary
                        expandIcon={<ExpandMoreIcon />}
                        sx={{ px: 0, minHeight: 0, '& .MuiAccordionSummary-content': { my: 0.5 } }}
                      >
                        <Typography variant="caption" color="text.secondary">
                          {itemReservations.length} reservation{itemReservations.length === 1 ? '' : 's'}
                        </Typography>
                      </AccordionSummary>
                      <AccordionDetails sx={{ px: 0, pt: 0 }}>
                        <Stack spacing={1} divider={<Divider flexItem />}>
                          {itemReservations.map((r) => (
                            <ReservationRow
                              key={r.id}
                              reservation={r}
                              optionLabel={r.itemId !== it.id ? (itemById[r.itemId]?.title ?? undefined) : undefined}
                              onSetStatus={(status) => setStatusM.mutate({ id: r.id, status })}
                              onDelete={() => {
                                const who = r.isAnonymous
                                  ? 'Anonymous'
                                  : r.reserverName?.trim() || r.contactEmail?.trim() || 'Someone';
                                setDeleteTarget({ kind: 'reservation', id: r.id, title: who });
                              }}
                            />
                          ))}
                        </Stack>
                      </AccordionDetails>
                    </Accordion>
                  )}
                </CardContent>
              </Card>
            </Grid>
          );
        })}
        {list.length === 0 && (
          <Grid item xs={12}>
            <Typography color="text.secondary">No items yet. Paste a product URL to start.</Typography>
          </Grid>
        )}
      </Grid>}

      <Dialog open={open} onClose={() => setOpen(false)} TransitionProps={{ onExited: () => reset() }} fullWidth maxWidth="sm" PaperProps={{ sx: modalPaperSx }}>
        <DialogTitle>Add item</DialogTitle>
        <DialogContent sx={{ flex: 1, overflow: 'hidden' }}>
          <Grid
            container
            spacing={2}
            sx={{
              mt: 0.5,
              height: '100%',
              minHeight: 0,
              overflowY: 'auto',
              overflowX: 'hidden',
              scrollbarGutter: 'stable',
              pr: 1,
            }}
          >
            <Grid item xs={12} md={5}>
              <Stack spacing={1.25} sx={{ minWidth: 0 }}>
                <Card
                  variant="outlined"
                  sx={{
                    borderRadius: 2,
                    overflow: 'hidden',
                    bgcolor: '#f6efe6',
                    minHeight: 220,
                  }}
                >
                  {imageUrl ? (
                    <CardMedia component="img" image={imageUrl} sx={{ aspectRatio: '1 / 1', objectFit: 'contain' }} />
                  ) : (
                    <Stack sx={{ aspectRatio: '1 / 1' }} alignItems="center" justifyContent="center" spacing={0.5}>
                      <Inventory2Icon color="disabled" />
                      <Typography variant="body2" color="text.secondary">
                        Image preview
                      </Typography>
                    </Stack>
                  )}
                </Card>
                <Typography variant="caption" color="text.secondary">
                  Preview updates as you fetch or paste an image URL.
                </Typography>
              </Stack>
            </Grid>
            <Grid item xs={12} md={7} sx={{ minWidth: 0 }}>
              <Stack spacing={2} sx={{ minWidth: 0 }}>
                <Stack direction={{ xs: 'column', sm: 'row' }} spacing={1}>
                  <TextField fullWidth label="Product URL" value={url} onChange={(e) => setUrl(e.target.value)} />
                  <Button variant="outlined" onClick={doScrape} disabled={!url || scraping} sx={{ minWidth: 110 }}>
                    {scraping ? <CircularProgress size={20} /> : 'Fetch'}
                  </Button>
                </Stack>
                {url.trim() && (
                  <Button variant="text" component="a" href={url} target="_blank" rel="noreferrer" sx={{ alignSelf: 'flex-start' }}>
                    View product page
                  </Button>
                )}
                <TextField fullWidth label="Title" value={title} onChange={(e) => setTitle(e.target.value)} />
                <TextField fullWidth label="Description" value={description} onChange={(e) => setDescription(e.target.value)} multiline minRows={2} />
                <TextField fullWidth label="Image URL" value={imageUrl} onChange={(e) => setImageUrl(e.target.value)} />
                <FormControlLabel
                  control={<Checkbox checked={quantityUnlimited} onChange={(e) => setQuantityUnlimited(e.target.checked)} />}
                  label="Allow unlimited reservations"
                />
                {!quantityUnlimited && (
                  <TextField fullWidth label="Quantity" type="number" inputProps={{ min: 1, step: 1 }} value={quantity} onChange={(e) => setQuantity(e.target.value)} />
                )}
                <Autocomplete
                  freeSolo
                  options={categoryOptions}
                  value={category}
                  inputValue={category}
                  onInputChange={(_, value) => setCategory(value)}
                  onChange={(_, value) => setCategory(typeof value === 'string' ? value : value || '')}
                  renderInput={(params) => (
                    <TextField
                      {...params}
                      fullWidth
                      label="Category"
                      placeholder="Select existing or create new"
                    />
                  )}
                />
                {parentItemId && itemById[parentItemId] && (
                  <Alert severity="info">
                    Adding as an alternative under <strong>{itemById[parentItemId].title}</strong>.
                  </Alert>
                )}
                {!parentItemId && (
                  <FormControlLabel
                    control={<Checkbox checked={noSubstitutes} onChange={(e) => setNoSubstitutes(e.target.checked)} />}
                    label="No substitutes"
                  />
                )}
                <TextField fullWidth label="Source" value={source} onChange={(e) => setSource(e.target.value)} />
                <TextField fullWidth label="Notes" value={notes} onChange={(e) => setNotes(e.target.value)} multiline minRows={2} />
                {error && <Alert severity="error">{error}</Alert>}
              </Stack>
            </Grid>
          </Grid>
        </DialogContent>
        <DialogActions sx={{ px: 3, py: 2, borderTop: 1, borderColor: 'divider' }}>
          <Button onClick={() => setOpen(false)}>Cancel</Button>
          <Button variant="contained" onClick={() => createItemM.mutate()} disabled={!title}>
            Add
          </Button>
        </DialogActions>
      </Dialog>

      <Dialog
        open={editOpen}
        onClose={() => setEditOpen(false)}
        TransitionProps={{
          onExited: () => {
            if (suppressEditResetRef.current) {
              suppressEditResetRef.current = false;
            } else {
              reset();
            }
          },
        }}
        fullWidth
        maxWidth="md"
        PaperProps={{ sx: { ...modalPaperSx, width: 'min(960px, calc(100vw - 32px))' } }}
      >
        <DialogTitle>Edit item</DialogTitle>
        <DialogContent sx={{ flex: 1, overflow: 'hidden' }}>
          <Stack spacing={2} sx={{ mt: 1, height: '100%', minHeight: 0 }}>
            <Tabs
              value={editTab}
              onChange={(_, value) => {
                setEditTab(value);
                if (value === 'fulfillment' && editActiveReservations.length > 0) {
                  setFulfillmentStatus(editActiveReservations[0].status);
                }
              }}
              variant="fullWidth"
            >
              <Tab value="details" label="Details" />
              <Tab value="substitutes" label="Substitutes" />
              <Tab value="fulfillment" label="Fulfillment" />
            </Tabs>
            <Box
              sx={{
                flex: 1,
                minHeight: 0,
                overflowY: 'auto',
                overflowX: 'hidden',
                scrollbarGutter: 'stable',
                pt: 1,
                pr: 1,
                pb: 0.5,
              }}
            >
              {editTab === 'details' && (
                <Stack direction={{ xs: 'column', md: 'row' }} spacing={2} sx={{ minWidth: 0 }}>
                  <Box sx={{ width: { xs: '100%', md: '40%' }, minWidth: 0 }}>
                    <Stack spacing={1.25} sx={{ minWidth: 0 }}>
                      <Card
                        variant="outlined"
                        sx={{
                          borderRadius: 2,
                          overflow: 'hidden',
                          bgcolor: '#f6efe6',
                          minHeight: 220,
                        }}
                      >
                        {imageUrl ? (
                          <CardMedia component="img" image={imageUrl} sx={{ aspectRatio: '1 / 1', objectFit: 'contain' }} />
                        ) : (
                          <Stack sx={{ aspectRatio: '1 / 1' }} alignItems="center" justifyContent="center" spacing={0.5}>
                            <Inventory2Icon color="disabled" />
                            <Typography variant="body2" color="text.secondary">
                              Image preview
                            </Typography>
                          </Stack>
                        )}
                      </Card>
                      <Typography variant="caption" color="text.secondary">
                        Preview updates as you fetch or paste an image URL.
                      </Typography>
                    </Stack>
                  </Box>
                  <Box sx={{ width: { xs: '100%', md: '60%' }, minWidth: 0 }}>
                    <Stack spacing={2} sx={{ minWidth: 0 }}>
                      <Stack direction={{ xs: 'column', sm: 'row' }} spacing={1}>
                        <TextField fullWidth label="Product URL" value={url} onChange={(e) => setUrl(e.target.value)} />
                        <Button variant="outlined" onClick={doScrape} disabled={!url || scraping} sx={{ minWidth: 110 }}>
                          {scraping ? <CircularProgress size={20} /> : 'Fetch'}
                        </Button>
                      </Stack>
                      {url.trim() && (
                        <Button variant="text" component="a" href={url} target="_blank" rel="noreferrer" sx={{ alignSelf: 'flex-start' }}>
                          View product page
                        </Button>
                      )}
                      <TextField fullWidth label="Title" value={title} onChange={(e) => setTitle(e.target.value)} />
                      <TextField fullWidth label="Description" value={description} onChange={(e) => setDescription(e.target.value)} multiline minRows={2} />
                      <TextField fullWidth label="Image URL" value={imageUrl} onChange={(e) => setImageUrl(e.target.value)} />
                      <FormControlLabel
                        control={<Checkbox checked={quantityUnlimited} onChange={(e) => setQuantityUnlimited(e.target.checked)} />}
                        label="Allow unlimited reservations"
                      />
                      {!quantityUnlimited && (
                        <TextField fullWidth label="Quantity" type="number" inputProps={{ min: 1, step: 1 }} value={quantity} onChange={(e) => setQuantity(e.target.value)} />
                      )}
                      <Autocomplete
                        freeSolo
                        options={categoryOptions}
                        value={category}
                        inputValue={category}
                        onInputChange={(_, value) => setCategory(value)}
                        onChange={(_, value) => setCategory(typeof value === 'string' ? value : value || '')}
                        renderInput={(params) => (
                          <TextField
                            {...params}
                            fullWidth
                            label="Category"
                            placeholder="Select existing or create new"
                          />
                        )}
                      />
                      <TextField fullWidth label="Source" value={source} onChange={(e) => setSource(e.target.value)} />
                      <TextField fullWidth label="Notes" value={notes} onChange={(e) => setNotes(e.target.value)} multiline minRows={2} />
                    </Stack>
                  </Box>
                </Stack>
              )}

              {editTab === 'substitutes' && (
                <Stack spacing={2}>
                <Alert severity="info">
                  Can't find the exact one? Add similar options here so your guests can pick whichever version works best for them.
                </Alert>
                {editRootItem && (
                  <Typography variant="body2" color="text.secondary">
                    Top-level item: <strong>{editRootItem.title}</strong> ({editOptionCount} option{editOptionCount === 1 ? '' : 's'})
                  </Typography>
                )}
                {!parentItemId ? (
                  <FormControlLabel
                    control={<Checkbox checked={noSubstitutes} onChange={(e) => setNoSubstitutes(e.target.checked)} />}
                    label="No substitutes"
                  />
                ) : (
                  <Stack direction={{ xs: 'column', sm: 'row' }} spacing={1}>
                    {editRootItem && (
                      <Button variant="text" onClick={() => beginEdit(editRootItem)}>
                        Open top-level item
                      </Button>
                    )}
                    <Button variant="text" color="warning" onClick={() => setParentItemId('')}>
                      Make this item top-level
                    </Button>
                  </Stack>
                )}
                {editRootId && (
                  <Button variant="outlined" onClick={() => beginAddAlternative(editRootId)}>
                    Add alternative
                  </Button>
                )}
                {editAlternatives.length > 0 ? (
                  <Stack spacing={1}>
                    <Typography variant="caption" color="text.secondary">Current alternatives</Typography>
                    {editAlternatives.map((alt) => (
                      <Stack key={alt.id} direction="row" alignItems="center" spacing={1}>
                        <Typography variant="body2" sx={{ flexGrow: 1 }}>
                          {alt.title}
                        </Typography>
                        <Button size="small" variant="text" onClick={() => beginEdit(alt)}>
                          Manage
                        </Button>
                      </Stack>
                    ))}
                  </Stack>
                ) : (
                  <Typography variant="body2" color="text.secondary">No alternatives yet.</Typography>
                )}
                </Stack>
              )}

              {editTab === 'fulfillment' && (
                <Stack spacing={2}>
                  <Alert severity="info">
                    {editActiveReservations.length > 0
                      ? `This will update ${editActiveReservations.length} active reservation${editActiveReservations.length === 1 ? '' : 's'} for this item.`
                      : 'No active reservations found. Saving will create an owner reservation so status stays consistent.'}
                  </Alert>
                  <Select
                    value={fulfillmentStatus}
                    onChange={(e) => setFulfillmentStatus(e.target.value as ReservationStatus)}
                  >
                    <MenuItem value="Reserved">Reserved</MenuItem>
                    <MenuItem value="Purchased">Purchased</MenuItem>
                    <MenuItem value="Received">Received</MenuItem>
                    <MenuItem value="Cancelled">Cancelled</MenuItem>
                  </Select>
                  {editGroupActiveReservations.length > 0 && (
                    <Stack spacing={0}>
                      <Typography variant="overline" color="text.secondary">
                        Who's buying this
                      </Typography>
                      <Stack divider={<Divider flexItem />}>
                        {editGroupActiveReservations.map((r) => {
                          const reservedItem = itemById[r.itemId];
                          const optionLabel =
                            reservedItem && reservedItem.parentItemId
                              ? reservedItem.title ?? undefined
                              : undefined;
                          return (
                            <ReservationRow
                              key={r.id}
                              reservation={r}
                              optionLabel={optionLabel}
                              onSetStatus={(status) => setStatusM.mutate({ id: r.id, status })}
                              onDelete={() => {
                                const who = r.isAnonymous
                                  ? 'Anonymous'
                                  : r.reserverName?.trim() || r.contactEmail?.trim() || 'Someone';
                                setDeleteTarget({ kind: 'reservation', id: r.id, title: who });
                              }}
                            />
                          );
                        })}
                      </Stack>
                    </Stack>
                  )}
                </Stack>
              )}

              {error && <Alert severity="error">{error}</Alert>}
            </Box>
          </Stack>
        </DialogContent>
        <DialogActions sx={{ px: 3, py: 2, borderTop: 1, borderColor: 'divider' }}>
          <Button
            color="error"
            onClick={() => {
              if (!editingId) return;
              setEditOpen(false);
              setDeleteTarget({ kind: 'item', id: editingId, title: title.trim() || 'this item' });
            }}
            disabled={!editingId}
          >
            Delete item
          </Button>
          <Box sx={{ flexGrow: 1 }} />
          <Button onClick={() => setEditOpen(false)}>Cancel</Button>
          <Button
            variant="contained"
            onClick={() => {
              if (editTab === 'fulfillment') {
                if (!editingItem) return;
                markPurchasedM.mutate({
                  item: editingItem,
                  activeReservations: editActiveReservations,
                  status: fulfillmentStatus,
                });
                return;
              }
              updateItemM.mutate();
            }}
            disabled={
              editTab === 'fulfillment'
                ? !editingItem || markPurchasedM.isPending
                : !title || !editingId || updateItemM.isPending
            }
          >
            {editTab === 'fulfillment' ? 'Save status' : 'Save changes'}
          </Button>
        </DialogActions>
      </Dialog>

      <Dialog open={!!deleteTarget} onClose={() => setDeleteTarget(null)} maxWidth="xs" fullWidth>
        <DialogTitle>Confirm deletion</DialogTitle>
        <DialogContent>
          <Typography variant="body2" color="text.secondary">
            {deleteTarget?.kind === 'item'
              ? `Delete item "${deleteTarget.title}"? This cannot be undone.`
              : `Delete reservation from "${deleteTarget?.title}"? This cannot be undone.`}
          </Typography>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setDeleteTarget(null)}>Cancel</Button>
          <Button variant="contained" color="error" onClick={confirmDelete}>
            Delete
          </Button>
        </DialogActions>
      </Dialog>
    </Container>
  );
}

function ReservationRow({
  reservation,
  optionLabel,
  onSetStatus,
  onDelete,
}: {
  reservation: Reservation;
  optionLabel?: string;
  onSetStatus: (status: ReservationStatus) => void;
  onDelete: () => void;
}) {
  const who = reservation.isAnonymous
    ? 'Anonymous'
    : reservation.reserverName?.trim() || reservation.contactEmail?.trim() || 'Someone';
  const qty = reservation.quantity ?? 1;
  return (
    <Stack direction="row" spacing={1} alignItems="flex-start" sx={{ py: 1 }}>
      <Stack sx={{ flex: 1, minWidth: 0 }}>
        <Typography variant="body2" sx={{ fontWeight: 500 }} noWrap>
          {who}
          {qty > 1 ? ` · ×${qty}` : ''}
        </Typography>
        {optionLabel && (
          <Typography variant="caption" color="primary.main" noWrap>
            {optionLabel}
          </Typography>
        )}
        {!reservation.isAnonymous && reservation.contactEmail && (
          <Typography variant="caption" color="text.secondary" noWrap>
            {reservation.contactEmail}
          </Typography>
        )}
        {reservation.message && (
          <Typography
            variant="caption"
            color="text.secondary"
            sx={{ display: 'block', whiteSpace: 'pre-wrap', mt: 0.5 }}
          >
            “{reservation.message}”
          </Typography>
        )}
      </Stack>
      <Select
        size="small"
        value={reservation.status}
        onChange={(e) => onSetStatus(e.target.value as ReservationStatus)}
        sx={{ minWidth: 130 }}
      >
        <MenuItem value="Reserved">Reserved</MenuItem>
        <MenuItem value="Purchased">Purchased</MenuItem>
        <MenuItem value="Received">Received</MenuItem>
        <MenuItem value="Cancelled">Cancelled</MenuItem>
      </Select>
      <Tooltip title="Delete reservation">
        <IconButton size="small" onClick={onDelete} aria-label="delete reservation">
          <DeleteIcon fontSize="small" />
        </IconButton>
      </Tooltip>
    </Stack>
  );
}
