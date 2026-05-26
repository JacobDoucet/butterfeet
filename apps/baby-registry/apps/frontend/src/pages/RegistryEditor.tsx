import { useState } from 'react';
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
} from '@mui/material';
import DeleteIcon from '@mui/icons-material/DeleteOutline';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import EditIcon from '@mui/icons-material/EditOutlined';
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
  const [editingId, setEditingId] = useState<string | null>(null);
  const [fulfillmentOpen, setFulfillmentOpen] = useState(false);
  const [fulfillmentItemId, setFulfillmentItemId] = useState<string | null>(null);
  const [fulfillmentStatus, setFulfillmentStatus] = useState<ReservationStatus>('Purchased');
  const [deleteTarget, setDeleteTarget] = useState<DeleteTarget | null>(null);
  const [url, setUrl] = useState('');
  const [scraping, setScraping] = useState(false);
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [imageUrl, setImageUrl] = useState('');
  const [source, setSource] = useState('');
  const [quantity, setQuantity] = useState('1');
  const [notes, setNotes] = useState('');
  const [error, setError] = useState<string | null>(null);

  const reset = () => {
    setUrl('');
    setTitle('');
    setDescription('');
    setImageUrl('');
    setSource('');
    setQuantity('1');
    setNotes('');
    setEditingId(null);
    setError(null);
  };

  const beginEdit = (it: RegistryItem) => {
    setEditingId(it.id);
    setUrl(it.productUrl || '');
    setTitle(it.title || '');
    setDescription(it.description || '');
    setImageUrl(it.imageUrl || '');
    setSource(it.source || '');
    setQuantity(String(it.quantity || 1));
    setNotes(it.notes || '');
    setError(null);
    setEditOpen(true);
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
      const qty = Math.max(1, parseInt(quantity || '1', 10) || 1);
      return items.create({
        registryId: reg!.id,
        title,
        description,
        imageUrl,
        productUrl: url,
        source,
        quantity: qty,
        notes,
      });
    },
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['items', reg?.id] });
      setOpen(false);
      reset();
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
      const qty = Math.max(1, parseInt(quantity || '1', 10) || 1);
      return items.update(editingId, {
        title,
        description,
        imageUrl,
        productUrl: url,
        source,
        quantity: qty,
        notes,
      });
    },
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['items', reg?.id] });
      setEditOpen(false);
      reset();
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
      setFulfillmentOpen(false);
      setFulfillmentItemId(null);
      setFulfillmentStatus('Purchased');
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
  const fulfillmentItem = list.find((it) => it.id === fulfillmentItemId) ?? null;
  const fulfillmentActiveReservations = (fulfillmentItem ? reservationsByItem[fulfillmentItem.id] ?? [] : [])
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
          <Button variant="contained" onClick={() => setOpen(true)}>Add item</Button>
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
        {list.map((it) => {
          const itemReservations = reservationsByItem[it.id] ?? [];
          const activeReservations = itemReservations.filter((r) => r.status !== 'Cancelled');
          const activeCount = itemReservations
            .filter((r) => r.status !== 'Cancelled')
            .reduce((sum, r) => sum + (r.quantity ?? 1), 0);
          const requested = it.quantity ?? 1;
          const fulfilled = activeCount >= requested;
          return (
            <Grid item xs={12} sm={6} md={4} key={it.id}>
              <Card>
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
                    <Chip
                      size="small"
                      color={fulfilled ? 'success' : activeCount > 0 ? 'warning' : 'default'}
                      variant={activeCount > 0 ? 'filled' : 'outlined'}
                      label={
                        fulfilled
                          ? 'Reserved'
                          : activeCount > 0
                          ? `${activeCount} / ${requested} reserved`
                          : 'Unclaimed'
                      }
                    />
                  </Stack>
                  <Stack direction="row" spacing={1}>
                    {it.productUrl && (
                      <Button size="small" component="a" href={it.productUrl} target="_blank" rel="noreferrer">
                        Open
                      </Button>
                    )}
                    <Tooltip title="Edit item details">
                      <IconButton size="small" onClick={() => beginEdit(it)} aria-label="edit item">
                        <EditIcon fontSize="small" />
                      </IconButton>
                    </Tooltip>
                    <Button
                      size="small"
                      variant="outlined"
                      startIcon={<Inventory2Icon />}
                      onClick={() => {
                        setFulfillmentItemId(it.id);
                        setFulfillmentStatus('Purchased');
                        setError(null);
                        setFulfillmentOpen(true);
                      }}
                    >
                      Fulfillment
                    </Button>
                    <IconButton
                      size="small"
                      onClick={() => setDeleteTarget({ kind: 'item', id: it.id, title: it.title })}
                      aria-label="delete"
                    >
                      <DeleteIcon fontSize="small" />
                    </IconButton>
                  </Stack>

                  {itemReservations.length > 0 && (
                    <Accordion
                      disableGutters
                      elevation={0}
                      sx={{ mt: 2, bgcolor: 'transparent', '&:before': { display: 'none' } }}
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

      <Dialog open={open} onClose={() => { setOpen(false); reset(); }} fullWidth maxWidth="sm">
        <DialogTitle>Add item</DialogTitle>
        <DialogContent>
          <Grid container spacing={2} sx={{ mt: 0.5 }}>
            <Grid item xs={12} md={5}>
              <Stack spacing={1.25}>
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
            <Grid item xs={12} md={7}>
              <Stack spacing={2}>
                <Stack direction={{ xs: 'column', sm: 'row' }} spacing={1}>
                  <TextField fullWidth label="Product URL" value={url} onChange={(e) => setUrl(e.target.value)} />
                  <Button variant="outlined" onClick={doScrape} disabled={!url || scraping} sx={{ minWidth: 110 }}>
                    {scraping ? <CircularProgress size={20} /> : 'Fetch'}
                  </Button>
                </Stack>
                <TextField fullWidth label="Title" value={title} onChange={(e) => setTitle(e.target.value)} />
                <TextField fullWidth label="Description" value={description} onChange={(e) => setDescription(e.target.value)} multiline minRows={2} />
                <TextField fullWidth label="Image URL" value={imageUrl} onChange={(e) => setImageUrl(e.target.value)} />
                <TextField fullWidth label="Quantity" type="number" inputProps={{ min: 1, step: 1 }} value={quantity} onChange={(e) => setQuantity(e.target.value)} />
                <TextField fullWidth label="Source" value={source} onChange={(e) => setSource(e.target.value)} />
                <TextField fullWidth label="Notes" value={notes} onChange={(e) => setNotes(e.target.value)} multiline minRows={2} />
                {error && <Alert severity="error">{error}</Alert>}
              </Stack>
            </Grid>
          </Grid>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => { setOpen(false); reset(); }}>Cancel</Button>
          <Button variant="contained" onClick={() => createItemM.mutate()} disabled={!title}>
            Add
          </Button>
        </DialogActions>
      </Dialog>

      <Dialog open={editOpen} onClose={() => { setEditOpen(false); reset(); }} fullWidth maxWidth="sm">
        <DialogTitle>Edit item</DialogTitle>
        <DialogContent>
          <Grid container spacing={2} sx={{ mt: 0.5 }}>
            <Grid item xs={12} md={5}>
              <Stack spacing={1.25}>
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
            <Grid item xs={12} md={7}>
              <Stack spacing={2}>
                <Stack direction={{ xs: 'column', sm: 'row' }} spacing={1}>
                  <TextField fullWidth label="Product URL" value={url} onChange={(e) => setUrl(e.target.value)} />
                  <Button variant="outlined" onClick={doScrape} disabled={!url || scraping} sx={{ minWidth: 110 }}>
                    {scraping ? <CircularProgress size={20} /> : 'Fetch'}
                  </Button>
                </Stack>
                <TextField fullWidth label="Title" value={title} onChange={(e) => setTitle(e.target.value)} />
                <TextField fullWidth label="Description" value={description} onChange={(e) => setDescription(e.target.value)} multiline minRows={2} />
                <TextField fullWidth label="Image URL" value={imageUrl} onChange={(e) => setImageUrl(e.target.value)} />
                <TextField fullWidth label="Quantity" type="number" inputProps={{ min: 1, step: 1 }} value={quantity} onChange={(e) => setQuantity(e.target.value)} />
                <TextField fullWidth label="Source" value={source} onChange={(e) => setSource(e.target.value)} />
                <TextField fullWidth label="Notes" value={notes} onChange={(e) => setNotes(e.target.value)} multiline minRows={2} />
                {error && <Alert severity="error">{error}</Alert>}
              </Stack>
            </Grid>
          </Grid>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => { setEditOpen(false); reset(); }}>Cancel</Button>
          <Button variant="contained" onClick={() => updateItemM.mutate()} disabled={!title || !editingId || updateItemM.isPending}>
            Save changes
          </Button>
        </DialogActions>
      </Dialog>

      <Dialog
        open={fulfillmentOpen}
        onClose={() => {
          setFulfillmentOpen(false);
          setFulfillmentItemId(null);
          setFulfillmentStatus('Purchased');
        }}
        fullWidth
        maxWidth="sm"
      >
        <DialogTitle>Update fulfillment status</DialogTitle>
        <DialogContent>
          <Stack spacing={2} sx={{ mt: 1 }}>
            {fulfillmentItem && (
              <Typography variant="subtitle2" color="text.secondary">
                {fulfillmentItem.title}
              </Typography>
            )}
            <Alert severity="info">
              {fulfillmentActiveReservations.length > 0
                ? `This will update ${fulfillmentActiveReservations.length} active reservation${fulfillmentActiveReservations.length === 1 ? '' : 's'} for this item.`
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
            {error && <Alert severity="error">{error}</Alert>}
          </Stack>
        </DialogContent>
        <DialogActions>
          <Button
            onClick={() => {
              setFulfillmentOpen(false);
              setFulfillmentItemId(null);
              setFulfillmentStatus('Purchased');
            }}
          >
            Cancel
          </Button>
          <Button
            variant="contained"
            disabled={!fulfillmentItem || markPurchasedM.isPending}
            onClick={() => {
              if (!fulfillmentItem) return;
              markPurchasedM.mutate({
                item: fulfillmentItem,
                activeReservations: fulfillmentActiveReservations,
                status: fulfillmentStatus,
              });
            }}
          >
            Save status
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
  onSetStatus,
  onDelete,
}: {
  reservation: Reservation;
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
