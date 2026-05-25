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
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { registries, items, scrape, type RegistryItem, type Registry } from '../api';

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

  const [open, setOpen] = useState(false);
  const [url, setUrl] = useState('');
  const [scraping, setScraping] = useState(false);
  const [title, setTitle] = useState('');
  const [imageUrl, setImageUrl] = useState('');
  const [price, setPrice] = useState('');
  const [currency, setCurrency] = useState('');
  const [source, setSource] = useState('');
  const [error, setError] = useState<string | null>(null);

  const reset = () => {
    setUrl(''); setTitle(''); setImageUrl(''); setPrice(''); setCurrency(''); setSource(''); setError(null);
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
      return items.create({
        registryId: reg!.id,
        title,
        imageUrl,
        productUrl: url,
        source,
        priceCents,
        currency,
        quantity: 1,
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

  if (regsQ.isLoading) return null;
  if (!reg) return (
    <Container sx={{ py: 6 }}>
      <Alert severity="warning">Registry not found.</Alert>
    </Container>
  );

  const list: RegistryItem[] = itemsQ.data?.data ?? [];

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

      <Grid container spacing={2}>
        {list.map((it) => (
          <Grid item xs={12} sm={6} md={4} key={it.id}>
            <Card>
              {it.imageUrl && (
                <CardMedia component="img" image={it.imageUrl} sx={{ aspectRatio: '1', objectFit: 'contain', bgcolor: '#f4ede3' }} />
              )}
              <CardContent>
                <Typography variant="subtitle1" sx={{ fontWeight: 600 }} noWrap>
                  {it.title}
                </Typography>
                <Stack direction="row" spacing={1} sx={{ my: 1 }}>
                  {it.source && <Chip size="small" label={it.source} />}
                  {it.priceCents ? (
                    <Chip size="small" label={`${(it.priceCents / 100).toFixed(2)} ${it.currency || ''}`.trim()} variant="outlined" />
                  ) : null}
                </Stack>
                <Stack direction="row" spacing={1}>
                  {it.productUrl && (
                    <Button size="small" component="a" href={it.productUrl} target="_blank" rel="noreferrer">
                      Open
                    </Button>
                  )}
                  <IconButton size="small" onClick={() => deleteM.mutate(it.id)} aria-label="delete">
                    <DeleteIcon fontSize="small" />
                  </IconButton>
                </Stack>
              </CardContent>
            </Card>
          </Grid>
        ))}
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
            <TextField label="Image URL" value={imageUrl} onChange={(e) => setImageUrl(e.target.value)} />
            <Stack direction="row" spacing={2}>
              <TextField label="Price" value={price} onChange={(e) => setPrice(e.target.value)} />
              <TextField label="Currency" value={currency} onChange={(e) => setCurrency(e.target.value)} />
            </Stack>
            <TextField label="Source" value={source} onChange={(e) => setSource(e.target.value)} />
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
    </Container>
  );
}
