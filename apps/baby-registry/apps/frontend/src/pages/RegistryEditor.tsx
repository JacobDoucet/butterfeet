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
} from '@mui/material';
import DeleteIcon from '@mui/icons-material/DeleteOutline';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import EditIcon from '@mui/icons-material/EditOutlined';
import CheckCircleIcon from '@mui/icons-material/CheckCircleOutline';
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

export default function RegistryEditor() {
  const { slug = '' } = useParams();
  const qc = useQueryClient();

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
  const [url, setUrl] = useState('');
  const [scraping, setScraping] = useState(false);
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [imageUrl, setImageUrl] = useState('');
  const [price, setPrice] = useState('');
  const [currency, setCurrency] = useState('');
  const [source, setSource] = useState('');
  const [quantity, setQuantity] = useState('1');
  const [notes, setNotes] = useState('');
  const [error, setError] = useState<string | null>(null);

  const reset = () => {
    setUrl('');
    setTitle('');
    setDescription('');
    setImageUrl('');
    setPrice('');
    setCurrency('');
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
    setPrice(it.priceCents ? String((it.priceCents || 0) / 100) : '');
    setCurrency(it.currency || '');
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
      setPrice(r.price ? String(r.price) : '');
      setCurrency(r.currency || '');
      setSource(r.source || 'Other');
    } catch (err) {
      setError((err as Error).message);
    } finally {
      setScraping(false);
    }
  };

  const createItemM = useMutation({
    mutationFn: async () => {
      const priceCents = price ? Math.round(parseFloat(price) * 100) : 0;
      const qty = Math.max(1, parseInt(quantity || '1', 10) || 1);
      return items.create({
        registryId: reg!.id,
        title,
        description,
        imageUrl,
        productUrl: url,
        source,
        priceCents,
        currency,
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

  const setOwnerPurchasedM = useMutation({
    mutationFn: ({ id, ownerPurchased }: { id: string; ownerPurchased: boolean }) =>
      items.update(id, { ownerPurchased }),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['items', reg?.id] });
      qc.invalidateQueries({ queryKey: ['reservations', reg?.id] });
    },
  });

  const updateItemM = useMutation({
    mutationFn: async () => {
      if (!editingId) return;
      const priceCents = price ? Math.round(parseFloat(price) * 100) : 0;
      const qty = Math.max(1, parseInt(quantity || '1', 10) || 1);
      return items.update(editingId, {
        title,
        description,
        imageUrl,
        productUrl: url,
        source,
        priceCents,
        currency,
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
    mutationFn: async (reservationIds: string[]) => {
      await Promise.all(reservationIds.map((id) => reservations.setStatus(id, 'Purchased')));
    },
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['reservations', reg?.id] });
      qc.invalidateQueries({ queryKey: ['items', reg?.id] });
    },
  });

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

  return (
    <Container maxWidth="lg" sx={{ py: 6 }}>
      <Stack direction="row" alignItems="center" sx={{ mb: 4 }}>
        <Stack sx={{ flexGrow: 1 }}>
          <Typography variant="h4">{reg.title}</Typography>
          <Typography color="text.secondary" component={Link} to={`/r/${reg.slug}`}>
            View public page → /r/{reg.slug}
          </Typography>
        </Stack>
        <Button variant="contained" onClick={() => setOpen(true)}>Add item</Button>
      </Stack>

      <PrivacyPanel reg={reg} />

      <Grid container spacing={2}>
        {list.map((it) => {
          const itemReservations = reservationsByItem[it.id] ?? [];
          const activeReservations = itemReservations.filter((r) => r.status !== 'Cancelled');
          const canMarkPurchased = activeReservations.length > 0 && activeReservations.some((r) => r.status !== 'Purchased');
          const ownerPurchased = !!it.ownerPurchased;
          const activeCount = itemReservations
            .filter((r) => r.status !== 'Cancelled')
            .reduce((sum, r) => sum + (r.quantity ?? 1), 0);
          const requested = it.quantity ?? 1;
          const fulfilled = ownerPurchased || activeCount >= requested;
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
                    {it.priceCents ? (
                      <Chip size="small" label={`${(it.priceCents / 100).toFixed(2)} ${it.currency || ''}`.trim()} variant="outlined" />
                    ) : null}
                    <Chip
                      size="small"
                      color={fulfilled ? 'success' : activeCount > 0 ? 'warning' : 'default'}
                      variant={activeCount > 0 ? 'filled' : 'outlined'}
                      label={
                        ownerPurchased
                          ? 'Purchased (Owner)'
                          : fulfilled
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
                    <Tooltip title={canMarkPurchased ? 'Mark all active reservations as purchased' : 'No active reservations to mark purchased'}>
                      <span>
                        <IconButton
                          size="small"
                          color="success"
                          disabled={!canMarkPurchased || markPurchasedM.isPending}
                          onClick={() => markPurchasedM.mutate(activeReservations.map((r) => r.id))}
                          aria-label="mark purchased"
                        >
                          <CheckCircleIcon fontSize="small" />
                        </IconButton>
                      </span>
                    </Tooltip>
                    <Button
                      size="small"
                      variant={ownerPurchased ? 'outlined' : 'contained'}
                      color={ownerPurchased ? 'inherit' : 'success'}
                      onClick={() => setOwnerPurchasedM.mutate({ id: it.id, ownerPurchased: !ownerPurchased })}
                      disabled={setOwnerPurchasedM.isPending}
                    >
                      {ownerPurchased ? 'Remove purchased' : 'Mark purchased'}
                    </Button>
                    <IconButton size="small" onClick={() => deleteM.mutate(it.id)} aria-label="delete">
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
                              onDelete={() => deleteReservationM.mutate(r.id)}
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
      </Grid>

      <Dialog open={open} onClose={() => { setOpen(false); reset(); }} fullWidth maxWidth="sm">
        <DialogTitle>Add item</DialogTitle>
        <DialogContent>
          <Stack spacing={2} sx={{ mt: 1 }}>
            <Stack direction="row" spacing={1}>
              <TextField label="Product URL" value={url} onChange={(e) => setUrl(e.target.value)} />
              <Button variant="outlined" onClick={doScrape} disabled={!url || scraping}>
                {scraping ? <CircularProgress size={20} /> : 'Fetch'}
              </Button>
            </Stack>
            <TextField label="Title" value={title} onChange={(e) => setTitle(e.target.value)} />
            <TextField label="Description" value={description} onChange={(e) => setDescription(e.target.value)} multiline minRows={2} />
            <TextField label="Image URL" value={imageUrl} onChange={(e) => setImageUrl(e.target.value)} />
            <Stack direction="row" spacing={2}>
              <TextField label="Price" value={price} onChange={(e) => setPrice(e.target.value)} />
              <TextField label="Currency" value={currency} onChange={(e) => setCurrency(e.target.value)} />
            </Stack>
            <TextField label="Quantity" type="number" inputProps={{ min: 1, step: 1 }} value={quantity} onChange={(e) => setQuantity(e.target.value)} />
            <TextField label="Source" value={source} onChange={(e) => setSource(e.target.value)} />
            <TextField label="Notes" value={notes} onChange={(e) => setNotes(e.target.value)} multiline minRows={2} />
            {error && <Alert severity="error">{error}</Alert>}
          </Stack>
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
          <Stack spacing={2} sx={{ mt: 1 }}>
            <Stack direction="row" spacing={1}>
              <TextField label="Product URL" value={url} onChange={(e) => setUrl(e.target.value)} />
              <Button variant="outlined" onClick={doScrape} disabled={!url || scraping}>
                {scraping ? <CircularProgress size={20} /> : 'Fetch'}
              </Button>
            </Stack>
            <TextField label="Title" value={title} onChange={(e) => setTitle(e.target.value)} />
            <TextField label="Description" value={description} onChange={(e) => setDescription(e.target.value)} multiline minRows={2} />
            <TextField label="Image URL" value={imageUrl} onChange={(e) => setImageUrl(e.target.value)} />
            <Stack direction="row" spacing={2}>
              <TextField label="Price" value={price} onChange={(e) => setPrice(e.target.value)} />
              <TextField label="Currency" value={currency} onChange={(e) => setCurrency(e.target.value)} />
            </Stack>
            <TextField label="Quantity" type="number" inputProps={{ min: 1, step: 1 }} value={quantity} onChange={(e) => setQuantity(e.target.value)} />
            <TextField label="Source" value={source} onChange={(e) => setSource(e.target.value)} />
            <TextField label="Notes" value={notes} onChange={(e) => setNotes(e.target.value)} multiline minRows={2} />
            {error && <Alert severity="error">{error}</Alert>}
          </Stack>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => { setEditOpen(false); reset(); }}>Cancel</Button>
          <Button variant="contained" onClick={() => updateItemM.mutate()} disabled={!title || !editingId || updateItemM.isPending}>
            Save changes
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
