import { useMemo, useRef, useState } from 'react';
import {
  Alert, Box, Button, Checkbox, Dialog, DialogActions, DialogContent, DialogTitle,
  IconButton, LinearProgress, MenuItem, Select, Stack, Step, StepLabel, Stepper,
  Table, TableBody, TableCell, TableContainer, TableHead, TableRow, TextField,
  Tooltip, Typography,
} from '@mui/material';
import CloseIcon from '@mui/icons-material/Close';
import UploadFileIcon from '@mui/icons-material/UploadFile';
import { items, scrape, type RegistryItem } from '../api';

interface Props {
  open: boolean;
  onClose: () => void;
  registryId: string;
  existingItems: RegistryItem[];
  onComplete: (createdCount: number, skippedCount: number) => void;
}

type TargetField =
  | 'title'
  | 'productUrl'
  | 'description'
  | 'imageUrl'
  | 'price'
  | 'currency'
  | 'quantity'
  | 'category'
  | 'notes';

interface FieldDef {
  key: TargetField;
  label: string;
  required?: boolean;
  hints: string[]; // lowercase substrings to match in CSV headers
}

const FIELDS: FieldDef[] = [
  { key: 'title', label: 'Title', required: true, hints: ['title', 'name', 'product', 'item'] },
  { key: 'productUrl', label: 'Product URL', hints: ['url', 'link', 'product url', 'product link', 'href'] },
  { key: 'imageUrl', label: 'Image URL', hints: ['image', 'photo', 'picture', 'img'] },
  { key: 'price', label: 'Price', hints: ['price', 'cost'] },
  { key: 'currency', label: 'Currency', hints: ['currency', 'ccy'] },
  { key: 'quantity', label: 'Quantity', hints: ['qty', 'quantity', 'count', 'amount'] },
  { key: 'category', label: 'Category', hints: ['category', 'cat', 'type', 'room'] },
  { key: 'description', label: 'Description', hints: ['description', 'desc', 'details'] },
  { key: 'notes', label: 'Notes', hints: ['note', 'comment', 'message'] },
];

interface ParsedRow {
  values: Record<string, string>; // keyed by CSV header
  include: boolean;
  duplicate: boolean;
  edits: Partial<Record<TargetField, string>>; // user overrides
}

const STEPS = ['Upload', 'Map columns', 'Preview & edit', 'Importing'];

export default function CsvImportDialog({ open, onClose, registryId, existingItems, onComplete }: Props) {
  const [step, setStep] = useState(0);
  const [fileName, setFileName] = useState('');
  const [headers, setHeaders] = useState<string[]>([]);
  const [rows, setRows] = useState<ParsedRow[]>([]);
  const [mapping, setMapping] = useState<Record<TargetField, string>>(() => emptyMapping());
  const [parseError, setParseError] = useState<string | null>(null);
  const [importProgress, setImportProgress] = useState({ done: 0, total: 0, errors: 0 });
  const [enrichProgress, setEnrichProgress] = useState({ done: 0, total: 0, active: false });
  const enrichRunIdRef = useRef(0);
  const fileInputRef = useRef<HTMLInputElement | null>(null);

  const existingUrls = useMemo(
    () => new Set(existingItems.map((it) => normalizeUrl(it.productUrl ?? '')).filter(Boolean)),
    [existingItems],
  );

  const handleFile = async (file: File) => {
    setParseError(null);
    try {
      const text = await file.text();
      const { headers: hs, rows: rs } = parseCsv(text);
      if (hs.length === 0 || rs.length === 0) {
        setParseError('No data found in this file.');
        return;
      }
      setFileName(file.name);
      setHeaders(hs);
      const initialMapping = autoMap(hs, rs);
      setMapping(initialMapping);
      const parsedRows: ParsedRow[] = rs.map((values) => {
        const url = normalizeUrl(values[initialMapping.productUrl] ?? '');
        const dup = !!url && existingUrls.has(url);
        return { values, include: !dup, duplicate: dup, edits: {} };
      });
      setRows(parsedRows);
      setStep(1);
    } catch (e) {
      setParseError((e as Error).message || 'Could not read this file.');
    }
  };

  const recomputeDuplicates = (m: Record<TargetField, string>, rs: ParsedRow[]): ParsedRow[] =>
    rs.map((r) => {
      const url = normalizeUrl((r.edits.productUrl ?? r.values[m.productUrl]) ?? '');
      const dup = !!url && existingUrls.has(url);
      return { ...r, duplicate: dup, include: dup ? false : r.include };
    });

  const updateMapping = (field: TargetField, header: string) => {
    const next = { ...mapping, [field]: header };
    setMapping(next);
    if (field === 'productUrl') setRows((rs) => recomputeDuplicates(next, rs));
  };

  const resolved = useMemo(
    () => rows.map((r) => resolveRow(r, mapping)),
    [rows, mapping],
  );

  const eligibleCount = resolved.filter((r, i) => rows[i].include && !rows[i].duplicate && r.title.trim()).length;
  const duplicateCount = rows.filter((r) => r.duplicate).length;
  const missingTitleCount = resolved.filter((r, i) => rows[i].include && !rows[i].duplicate && !r.title.trim()).length;

  const handleClose = () => {
    if (step === 3 && importProgress.done < importProgress.total) return; // block close mid-import
    enrichRunIdRef.current += 1; // cancel any in-flight enrichment
    onClose();
    // reset on next open
    setTimeout(() => {
      setStep(0); setFileName(''); setHeaders([]); setRows([]); setMapping(emptyMapping());
      setParseError(null); setImportProgress({ done: 0, total: 0, errors: 0 });
      setEnrichProgress({ done: 0, total: 0, active: false });
    }, 200);
  };

  const goToPreview = () => {
    setStep(2);
    runEnrichment();
  };

  const runEnrichment = () => {
    enrichRunIdRef.current += 1;
    const runId = enrichRunIdRef.current;
    // Determine which rows need enrichment: have a productUrl and at least one missing scrapeable field.
    const snapshot = rows.map((r, i) => ({ r, i }));
    const targets = snapshot.filter(({ r }) => {
      const url = (r.edits.productUrl ?? r.values[mapping.productUrl] ?? '').trim();
      if (!url || !/^https?:\/\//i.test(url)) return false;
      // Always scrape when we have a productUrl: image is overwritten on
      // every run, and other fields backfill only when empty.
      return true;
    });
    if (targets.length === 0) {
      setEnrichProgress({ done: 0, total: 0, active: false });
      return;
    }
    setEnrichProgress({ done: 0, total: targets.length, active: true });

    const CONCURRENCY = 4;
    let cursor = 0;
    let completed = 0;

    const worker = async (): Promise<void> => {
      while (true) {
        if (runId !== enrichRunIdRef.current) return;
        const myIdx = cursor++;
        if (myIdx >= targets.length) return;
        const { r, i } = targets[myIdx];
        const url = (r.edits.productUrl ?? r.values[mapping.productUrl] ?? '').trim();
        try {
          const result = await scrape.url(url);
          if (runId !== enrichRunIdRef.current) return;
          setRows((current) => {
            const next = [...current];
            const cur = next[i];
            const getResolved = (f: TargetField) => (cur.edits[f] ?? (mapping[f] ? cur.values[mapping[f]] : '') ?? '').trim();
            const newEdits = { ...cur.edits };
            // Always prefer a freshly scraped image when available — CSV
            // image columns are often missing or stale for Amazon links.
            if (result.imageUrl) newEdits.imageUrl = result.imageUrl;
            if (!getResolved('title') && result.title) newEdits.title = result.title;
            if (!getResolved('price') && result.price && result.price > 0) newEdits.price = String(result.price);
            if (!getResolved('currency') && result.currency) newEdits.currency = result.currency;
            next[i] = { ...cur, edits: newEdits };
            return next;
          });
        } catch {
          // swallow per-row failures; the row simply stays as-is
        }
        completed += 1;
        setEnrichProgress({ done: completed, total: targets.length, active: completed < targets.length });
      }
    };

    Promise.all(Array.from({ length: Math.min(CONCURRENCY, targets.length) }, worker));
  };

  const handleImport = async () => {
    const toCreate = resolved
      .map((r, i) => ({ r, row: rows[i] }))
      .filter(({ row, r }) => row.include && !row.duplicate && r.title.trim());
    setImportProgress({ done: 0, total: toCreate.length, errors: 0 });
    setStep(3);
    let done = 0;
    let errors = 0;
    for (const { r } of toCreate) {
      try {
        await items.create({
          registryId,
          title: r.title.trim(),
          description: r.description?.trim() || undefined,
          productUrl: r.productUrl?.trim() || undefined,
          imageUrl: r.imageUrl?.trim() || undefined,
          priceCents: r.priceCents,
          currency: r.currency?.trim() || undefined,
          quantity: r.quantity,
          category: r.category?.trim() || undefined,
          notes: r.notes?.trim() || undefined,
        });
      } catch {
        errors += 1;
      }
      done += 1;
      setImportProgress({ done, total: toCreate.length, errors });
    }
    onComplete(done - errors, duplicateCount);
  };

  return (
    <Dialog open={open} onClose={handleClose} maxWidth="lg" fullWidth>
      <DialogTitle sx={{ display: 'flex', alignItems: 'center', pr: 1 }}>
        <Box sx={{ flexGrow: 1 }}>Import items from CSV</Box>
        <IconButton onClick={handleClose} size="small"><CloseIcon /></IconButton>
      </DialogTitle>
      <DialogContent dividers>
        <Stepper activeStep={step} sx={{ mb: 3 }}>
          {STEPS.map((label) => (
            <Step key={label}><StepLabel>{label}</StepLabel></Step>
          ))}
        </Stepper>

        {step === 0 && (
          <Stack spacing={2} alignItems="center" sx={{ py: 6 }}>
            <UploadFileIcon sx={{ fontSize: 56, color: 'text.secondary' }} />
            <Typography variant="body1">Upload a CSV exported from Google Sheets, Excel, or any spreadsheet.</Typography>
            <Typography variant="body2" color="text.secondary">
              The first row should be your column headers. You'll map them to item fields in the next step.
            </Typography>
            <Button
              variant="contained"
              startIcon={<UploadFileIcon />}
              onClick={() => fileInputRef.current?.click()}
            >
              Choose CSV file
            </Button>
            <input
              ref={fileInputRef}
              type="file"
              accept=".csv,text/csv"
              style={{ display: 'none' }}
              onChange={(e) => {
                const f = e.target.files?.[0];
                if (f) handleFile(f);
                e.target.value = '';
              }}
            />
            {parseError && <Alert severity="error" sx={{ width: '100%' }}>{parseError}</Alert>}
          </Stack>
        )}

        {step === 1 && (
          <Stack spacing={2}>
            <Typography variant="body2" color="text.secondary">
              File: <strong>{fileName}</strong> · {rows.length} rows · {headers.length} columns
            </Typography>
            <Typography variant="body2">
              Map your CSV columns to item fields. <em>Title</em> is required. The <em>Product URL</em>
              {' '}column is used to skip items already in your registry.
            </Typography>
            <Box sx={{ display: 'grid', gridTemplateColumns: { xs: '1fr', sm: '180px 1fr' }, gap: 1.5, alignItems: 'center' }}>
              {FIELDS.map((f) => (
                <Box key={f.key} sx={{ display: 'contents' }}>
                  <Typography variant="body2">
                    {f.label}{f.required && <span style={{ color: '#c62828' }}> *</span>}
                  </Typography>
                  <Select
                    size="small"
                    value={mapping[f.key]}
                    onChange={(e) => updateMapping(f.key, e.target.value)}
                    displayEmpty
                  >
                    <MenuItem value=""><em>— Not mapped —</em></MenuItem>
                    {headers.map((h) => (
                      <MenuItem key={h} value={h}>{h}</MenuItem>
                    ))}
                  </Select>
                </Box>
              ))}
            </Box>
          </Stack>
        )}

        {step === 2 && (
          <Stack spacing={2}>
            <Stack direction="row" spacing={2} alignItems="center" flexWrap="wrap">
              <Typography variant="body2">
                <strong>{eligibleCount}</strong> ready to import
              </Typography>
              {duplicateCount > 0 && (
                <Typography variant="body2" color="text.secondary">
                  · {duplicateCount} already in registry (skipped)
                </Typography>
              )}
              {missingTitleCount > 0 && (
                <Typography variant="body2" color="warning.main">
                  · {missingTitleCount} missing title
                </Typography>
              )}
            </Stack>
            {enrichProgress.total > 0 && (
              <Alert
                severity={enrichProgress.active ? 'info' : 'success'}
                action={enrichProgress.active ? undefined : (
                  <Button color="inherit" size="small" onClick={runEnrichment}>Re-fetch</Button>
                )}
              >
                {enrichProgress.active
                  ? `Fetching missing details from product pages (${enrichProgress.done}/${enrichProgress.total})\u2026`
                  : `Filled in details for ${enrichProgress.done} row${enrichProgress.done === 1 ? '' : 's'}.`}
                {enrichProgress.active && (
                  <LinearProgress
                    variant="determinate"
                    value={(enrichProgress.done / enrichProgress.total) * 100}
                    sx={{ mt: 0.5 }}
                  />
                )}
              </Alert>
            )}
            <TableContainer sx={{ maxHeight: '50vh' }}>
              <Table size="small" stickyHeader>
                <TableHead>
                  <TableRow>
                    <TableCell padding="checkbox" />
                    <TableCell sx={{ width: 56 }}>Img</TableCell>
                    <TableCell>Title *</TableCell>
                    <TableCell>Product URL</TableCell>
                    <TableCell>Price</TableCell>
                    <TableCell>Qty</TableCell>
                    <TableCell>Category</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {resolved.map((r, i) => {
                    const row = rows[i];
                    return (
                      <TableRow key={i} sx={{ opacity: row.duplicate ? 0.5 : 1 }}>
                        <TableCell padding="checkbox">
                          <Checkbox
                            size="small"
                            checked={row.include && !row.duplicate}
                            disabled={row.duplicate}
                            onChange={(e) => {
                              const next = [...rows];
                              next[i] = { ...next[i], include: e.target.checked };
                              setRows(next);
                            }}
                          />
                        </TableCell>
                        <TableCell sx={{ width: 56 }}>
                          {r.imageUrl ? (
                            <Box
                              component="img"
                              src={r.imageUrl}
                              alt=""
                              sx={{ width: 40, height: 40, objectFit: 'cover', borderRadius: 1, bgcolor: 'action.hover' }}
                              onError={(e) => { (e.currentTarget as HTMLImageElement).style.visibility = 'hidden'; }}
                            />
                          ) : (
                            <Box sx={{ width: 40, height: 40, borderRadius: 1, bgcolor: 'action.hover' }} />
                          )}
                        </TableCell>
                        <TableCell>
                          <TextField
                            size="small" variant="standard"
                            fullWidth
                            value={r.title}
                            onChange={(e) => updateRowEdit(i, 'title', e.target.value, rows, setRows)}
                            error={row.include && !row.duplicate && !r.title.trim()}
                            disabled={row.duplicate}
                          />
                        </TableCell>
                        <TableCell sx={{ minWidth: 220 }}>
                          {row.duplicate ? (
                            <Tooltip title="A matching URL is already in this registry">
                              <Typography variant="caption" color="text.secondary">
                                Already in registry
                              </Typography>
                            </Tooltip>
                          ) : (
                            <TextField
                              size="small" variant="standard" fullWidth
                              value={r.productUrl ?? ''}
                              onChange={(e) => updateRowEdit(i, 'productUrl', e.target.value, rows, setRows, mapping, existingUrls)}
                            />
                          )}
                        </TableCell>
                        <TableCell sx={{ width: 100 }}>
                          <TextField
                            size="small" variant="standard"
                            value={r.priceDisplay}
                            onChange={(e) => updateRowEdit(i, 'price', e.target.value, rows, setRows)}
                            disabled={row.duplicate}
                          />
                        </TableCell>
                        <TableCell sx={{ width: 70 }}>
                          <TextField
                            size="small" variant="standard"
                            value={r.quantity ?? ''}
                            onChange={(e) => updateRowEdit(i, 'quantity', e.target.value, rows, setRows)}
                            disabled={row.duplicate}
                          />
                        </TableCell>
                        <TableCell>
                          <TextField
                            size="small" variant="standard" fullWidth
                            value={r.category ?? ''}
                            onChange={(e) => updateRowEdit(i, 'category', e.target.value, rows, setRows)}
                            disabled={row.duplicate}
                          />
                        </TableCell>
                      </TableRow>
                    );
                  })}
                </TableBody>
              </Table>
            </TableContainer>
          </Stack>
        )}

        {step === 3 && (
          <Stack spacing={2} sx={{ py: 4 }}>
            <Typography>
              Importing {importProgress.done} of {importProgress.total}…
            </Typography>
            <LinearProgress
              variant="determinate"
              value={importProgress.total === 0 ? 100 : (importProgress.done / importProgress.total) * 100}
            />
            {importProgress.errors > 0 && (
              <Alert severity="warning">{importProgress.errors} row(s) failed to import.</Alert>
            )}
            {importProgress.done === importProgress.total && importProgress.total > 0 && (
              <Alert severity="success">
                Imported {importProgress.done - importProgress.errors} item(s).
              </Alert>
            )}
          </Stack>
        )}
      </DialogContent>
      <DialogActions sx={{ px: 3, py: 2 }}>
        {step > 0 && step < 3 && (
          <Button onClick={() => setStep((s) => s - 1)}>Back</Button>
        )}
        <Box sx={{ flexGrow: 1 }} />
        {step === 1 && (
          <Button
            variant="contained"
            onClick={goToPreview}
            disabled={!mapping.title}
          >
            Next: preview
          </Button>
        )}
        {step === 2 && (
          <Button
            variant="contained"
            onClick={handleImport}
            disabled={eligibleCount === 0 || enrichProgress.active}
          >
            {enrichProgress.active
              ? `Enriching ${enrichProgress.done}/${enrichProgress.total}…`
              : `Import ${eligibleCount} item${eligibleCount === 1 ? '' : 's'}`}
          </Button>
        )}
        {step === 3 && importProgress.done === importProgress.total && (
          <Button variant="contained" onClick={handleClose}>Done</Button>
        )}
      </DialogActions>
    </Dialog>
  );
}

// ---------- helpers ----------

function emptyMapping(): Record<TargetField, string> {
  return {
    title: '', productUrl: '', description: '', imageUrl: '',
    price: '', currency: '', quantity: '', category: '', notes: '',
  };
}

function autoMap(headers: string[], sampleRows: Record<string, string>[]): Record<TargetField, string> {
  const URL_FIELDS = new Set<TargetField>(['productUrl', 'imageUrl']);
  const looksLikeUrl = (header: string): boolean => {
    const samples = sampleRows.slice(0, 10).map((r) => (r[header] ?? '').trim()).filter(Boolean);
    if (samples.length === 0) return false;
    const urlish = samples.filter((v) => /^https?:\/\//i.test(v)).length;
    return urlish / samples.length >= 0.5;
  };
  const m = emptyMapping();
  const used = new Set<string>();
  for (const f of FIELDS) {
    const match = headers.find((h) => {
      if (used.has(h)) return false;
      const norm = h.toLowerCase().trim();
      if (!f.hints.some((hint) => norm === hint || norm.includes(hint))) return false;
      if (URL_FIELDS.has(f.key) && !looksLikeUrl(h)) return false;
      return true;
    });
    if (match) { m[f.key] = match; used.add(match); }
  }
  return m;
}

interface ResolvedRow {
  title: string;
  productUrl?: string;
  imageUrl?: string;
  description?: string;
  category?: string;
  notes?: string;
  currency?: string;
  quantity?: number;
  priceCents?: number;
  priceDisplay: string;
}

function resolveRow(row: ParsedRow, mapping: Record<TargetField, string>): ResolvedRow {
  const raw = (field: TargetField): string => {
    if (field in row.edits) return row.edits[field] ?? '';
    const header = mapping[field];
    return header ? (row.values[header] ?? '') : '';
  };
  const priceText = raw('price');
  const priceCents = parsePriceCents(priceText);
  const qtyText = raw('quantity');
  const qty = qtyText ? Math.max(1, parseInt(qtyText.replace(/[^0-9]/g, ''), 10) || 0) : undefined;
  return {
    title: raw('title'),
    productUrl: raw('productUrl') || undefined,
    imageUrl: raw('imageUrl') || undefined,
    description: raw('description') || undefined,
    category: raw('category') || undefined,
    notes: raw('notes') || undefined,
    currency: raw('currency') || undefined,
    quantity: qty || undefined,
    priceCents,
    priceDisplay: priceText,
  };
}

function updateRowEdit(
  i: number,
  field: TargetField,
  value: string,
  rows: ParsedRow[],
  setRows: (r: ParsedRow[]) => void,
  mapping?: Record<TargetField, string>,
  existingUrls?: Set<string>,
) {
  const next = [...rows];
  next[i] = { ...next[i], edits: { ...next[i].edits, [field]: value } };
  if (field === 'productUrl' && mapping && existingUrls) {
    const dup = !!normalizeUrl(value) && existingUrls.has(normalizeUrl(value));
    next[i] = { ...next[i], duplicate: dup, include: dup ? false : next[i].include };
  }
  setRows(next);
}

function normalizeUrl(u: string): string {
  if (!u) return '';
  try {
    const url = new URL(u.trim());
    url.hash = '';
    // strip common tracking params
    const skip = new Set(['utm_source','utm_medium','utm_campaign','utm_term','utm_content','gclid','fbclid','ref','ref_']);
    [...url.searchParams.keys()].forEach((k) => { if (skip.has(k.toLowerCase())) url.searchParams.delete(k); });
    return url.toString().replace(/\/+$/, '').toLowerCase();
  } catch {
    return u.trim().toLowerCase().replace(/\/+$/, '');
  }
}

function parsePriceCents(s: string): number | undefined {
  if (!s) return undefined;
  // pull the first number, supporting "$24.99", "24,99 €", "1,299.00", etc.
  const cleaned = s.replace(/[^0-9.,-]/g, '');
  if (!cleaned) return undefined;
  // If both . and , present, the last separator is the decimal.
  let normalized = cleaned;
  if (cleaned.includes('.') && cleaned.includes(',')) {
    const lastDot = cleaned.lastIndexOf('.');
    const lastComma = cleaned.lastIndexOf(',');
    if (lastComma > lastDot) {
      normalized = cleaned.replace(/\./g, '').replace(',', '.');
    } else {
      normalized = cleaned.replace(/,/g, '');
    }
  } else if (cleaned.includes(',') && !cleaned.includes('.')) {
    // Treat comma as decimal if 1–2 digits after it; else thousands.
    const parts = cleaned.split(',');
    if (parts.length === 2 && parts[1].length <= 2) normalized = cleaned.replace(',', '.');
    else normalized = cleaned.replace(/,/g, '');
  }
  const n = parseFloat(normalized);
  if (!isFinite(n) || n < 0) return undefined;
  return Math.round(n * 100);
}

// Minimal RFC4180-ish CSV parser. Handles quoted fields, escaped quotes, CRLF.
function parseCsv(text: string): { headers: string[]; rows: Record<string, string>[] } {
  // Strip BOM
  if (text.charCodeAt(0) === 0xfeff) text = text.slice(1);
  const records: string[][] = [];
  let field = '';
  let row: string[] = [];
  let inQuotes = false;
  for (let i = 0; i < text.length; i++) {
    const c = text[i];
    if (inQuotes) {
      if (c === '"') {
        if (text[i + 1] === '"') { field += '"'; i++; }
        else inQuotes = false;
      } else {
        field += c;
      }
    } else {
      if (c === '"') inQuotes = true;
      else if (c === ',') { row.push(field); field = ''; }
      else if (c === '\n' || c === '\r') {
        // commit row; handle \r\n
        if (c === '\r' && text[i + 1] === '\n') i++;
        row.push(field); field = '';
        if (row.length > 1 || row[0] !== '') records.push(row);
        row = [];
      } else {
        field += c;
      }
    }
  }
  if (field.length > 0 || row.length > 0) {
    row.push(field);
    if (row.length > 1 || row[0] !== '') records.push(row);
  }
  if (records.length === 0) return { headers: [], rows: [] };
  const headers = records[0].map((h) => h.trim());
  const rows = records.slice(1).map((r) => {
    const o: Record<string, string> = {};
    headers.forEach((h, idx) => { o[h] = (r[idx] ?? '').trim(); });
    return o;
  }).filter((o) => Object.values(o).some((v) => v !== ''));
  return { headers, rows };
}
