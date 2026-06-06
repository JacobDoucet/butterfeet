import { useState } from 'react';
import {
  Stack,
  Typography,
  TextField,
  Button,
  Alert,
  Box,
  Card,
  CardContent,
  Chip,
  Divider,
  Switch,
  FormControlLabel,
  IconButton,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  MenuItem,
  Select,
  CircularProgress,
} from '@mui/material';
import DeleteIcon from '@mui/icons-material/DeleteOutline';
import EditIcon from '@mui/icons-material/EditOutlined';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import {
  payments,
  PAYMENT_METHOD_LABELS,
  type Registry,
  type PaymentMethod,
  type PaymentMethodType,
} from '../../api';

const METHOD_TYPES: PaymentMethodType[] = [
  'PaymentLink',
  'InteracETransfer',
  'BankTransfer',
];

// Which detail fields are relevant for each payment method. Fields not listed
// here are hidden in the editor and cleared when switching to that method, so
// guests only ever see the inputs that make sense (e.g. email/phone only for
// Interac e-Transfer, bank details only for a bank transfer).
type MethodFields = {
  url?: boolean;
  email?: boolean;
  phone?: boolean;
  bank?: boolean;
};

const METHOD_FIELDS: Record<PaymentMethodType, MethodFields> = {
  PaymentLink: { url: true },
  InteracETransfer: { email: true, phone: true },
  BankTransfer: { bank: true },
};

function showsPaymentUrl(type: PaymentMethodType): boolean {
  return !!METHOD_FIELDS[type].url;
}

// The payment link is the only way to pay for a payment-link method, so it's
// required there.
function requiresPaymentUrl(type: PaymentMethodType): boolean {
  return type === 'PaymentLink';
}

function showsContactFields(type: PaymentMethodType): boolean {
  return !!METHOD_FIELDS[type].email || !!METHOD_FIELDS[type].phone;
}

// Banking detail fields are only meaningful for bank-style transfers.
function showsBankFields(type: PaymentMethodType): boolean {
  return !!METHOD_FIELDS[type].bank;
}

// clearedForType blanks out any draft fields that aren't relevant for the given
// method type, so switching methods doesn't leave stale values behind.
function clearedForType(draft: MethodDraft, type: PaymentMethodType): MethodDraft {
  const fields = METHOD_FIELDS[type];
  return {
    ...draft,
    type,
    paymentUrl: fields.url ? draft.paymentUrl : '',
    recipientEmail: fields.email ? draft.recipientEmail : '',
    recipientPhone: fields.phone ? draft.recipientPhone : '',
    bankName: fields.bank ? draft.bankName : '',
    bankAccountName: fields.bank ? draft.bankAccountName : '',
    bankAccountNumber: fields.bank ? draft.bankAccountNumber : '',
    bankRoutingNumber: fields.bank ? draft.bankRoutingNumber : '',
    bankIban: fields.bank ? draft.bankIban : '',
    bankSwift: fields.bank ? draft.bankSwift : '',
  };
}

type MethodDraft = Partial<PaymentMethod> & { type: PaymentMethodType };

function emptyDraft(): MethodDraft {
  return { type: 'PaymentLink', enabled: true };
}

export default function PaymentsPanel({ reg }: { reg: Registry }) {
  const qc = useQueryClient();
  const [dialogOpen, setDialogOpen] = useState(false);
  const [draft, setDraft] = useState<MethodDraft>(emptyDraft());
  const [editingId, setEditingId] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);

  const methodsQ = useQuery({
    queryKey: ['payment-methods', reg.id],
    queryFn: () => payments.listMethods(reg.id),
  });

  const invalidateMethods = () => qc.invalidateQueries({ queryKey: ['payment-methods', reg.id] });

  const saveM = useMutation({
    mutationFn: () => {
      const body: Partial<PaymentMethod> = {
        type: draft.type,
        displayName: draft.displayName?.trim() || '',
        instructions: draft.instructions?.trim() || '',
        paymentUrl: draft.paymentUrl?.trim() || '',
        recipientEmail: draft.recipientEmail?.trim() || '',
        recipientPhone: draft.recipientPhone?.trim() || '',
        bankName: draft.bankName?.trim() || '',
        bankAccountName: draft.bankAccountName?.trim() || '',
        bankAccountNumber: draft.bankAccountNumber?.trim() || '',
        bankRoutingNumber: draft.bankRoutingNumber?.trim() || '',
        bankIban: draft.bankIban?.trim() || '',
        bankSwift: draft.bankSwift?.trim() || '',
        enabled: draft.enabled ?? true,
      };
      return editingId ? payments.updateMethod(editingId, body) : payments.createMethod(reg.id, body);
    },
    onSuccess: () => {
      setDialogOpen(false);
      setError(null);
      invalidateMethods();
    },
    onError: (err) => setError((err as Error).message),
  });

  const toggleM = useMutation({
    mutationFn: ({ id, enabled }: { id: string; enabled: boolean }) =>
      payments.updateMethod(id, { enabled }),
    onSuccess: invalidateMethods,
  });

  const deleteM = useMutation({
    mutationFn: (id: string) => payments.removeMethod(id),
    onSuccess: invalidateMethods,
  });

  const openAdd = () => {
    setEditingId(null);
    setDraft(emptyDraft());
    setError(null);
    setDialogOpen(true);
  };
  const openEdit = (m: PaymentMethod) => {
    setEditingId(m.id);
    setDraft({ ...m });
    setError(null);
    setDialogOpen(true);
  };

  return (
    <Stack spacing={4}>
      <Box>
        <Stack direction="row" alignItems="center" justifyContent="space-between">
          <Box>
            <Typography variant="h6">Payment methods</Typography>
            <Typography variant="body2" color="text.secondary">
              Guests send money directly to you using one of these methods. Stork Nest never
              processes, holds, or verifies the payment.
            </Typography>
          </Box>
          <Button variant="contained" onClick={openAdd} sx={{ color: '#fff', flexShrink: 0 }}>
            Add method
          </Button>
        </Stack>

        {methodsQ.isLoading ? (
          <Box sx={{ py: 3, textAlign: 'center' }}><CircularProgress size={24} /></Box>
        ) : (methodsQ.data ?? []).length === 0 ? (
          <Alert severity="info" sx={{ mt: 2 }}>
            No payment methods yet. Add one so guests know how to send their contribution.
          </Alert>
        ) : (
          <Stack spacing={1.5} sx={{ mt: 2 }}>
            {(methodsQ.data ?? []).map((m) => (
              <Card key={m.id} variant="outlined">
                <CardContent sx={{ pb: '12px !important' }}>
                  <Stack direction="row" alignItems="flex-start" justifyContent="space-between" spacing={1}>
                    <Box>
                      <Stack direction="row" spacing={1} alignItems="center">
                        <Typography variant="subtitle1" fontWeight={600}>
                          {m.displayName?.trim() || PAYMENT_METHOD_LABELS[m.type]}
                        </Typography>
                        <Chip size="small" label={PAYMENT_METHOD_LABELS[m.type]} />
                      </Stack>
                      {m.recipientEmail && (
                        <Typography variant="body2" color="text.secondary">{m.recipientEmail}</Typography>
                      )}
                      {m.recipientPhone && (
                        <Typography variant="body2" color="text.secondary">{m.recipientPhone}</Typography>
                      )}
                      {m.paymentUrl && (
                        <Typography variant="body2" color="text.secondary">{m.paymentUrl}</Typography>
                      )}
                    </Box>
                    <Stack direction="row" alignItems="center" spacing={0.5}>
                      <FormControlLabel
                        sx={{ mr: 0 }}
                        control={
                          <Switch
                            checked={m.enabled ?? false}
                            onChange={(e) => toggleM.mutate({ id: m.id, enabled: e.target.checked })}
                          />
                        }
                        label={m.enabled ? 'On' : 'Off'}
                      />
                      <IconButton size="small" onClick={() => openEdit(m)}><EditIcon fontSize="small" /></IconButton>
                      <IconButton size="small" onClick={() => deleteM.mutate(m.id)}><DeleteIcon fontSize="small" /></IconButton>
                    </Stack>
                  </Stack>
                </CardContent>
              </Card>
            ))}
          </Stack>
        )}
      </Box>

      <MethodDialog
        open={dialogOpen}
        draft={draft}
        setDraft={setDraft}
        editing={!!editingId}
        error={error}
        saving={saveM.isPending}
        onClose={() => setDialogOpen(false)}
        onSave={() => saveM.mutate()}
      />
    </Stack>
  );
}

function MethodDialog({
  open,
  draft,
  setDraft,
  editing,
  error,
  saving,
  onClose,
  onSave,
}: {
  open: boolean;
  draft: MethodDraft;
  setDraft: (d: MethodDraft) => void;
  editing: boolean;
  error: string | null;
  saving: boolean;
  onClose: () => void;
  onSave: () => void;
}) {
  const set = (patch: Partial<MethodDraft>) => setDraft({ ...draft, ...patch });
  return (
    <Dialog open={open} onClose={onClose} fullWidth maxWidth="sm">
      <DialogTitle>{editing ? 'Edit payment method' : 'Add payment method'}</DialogTitle>
      <DialogContent>
        <Stack spacing={2} sx={{ mt: 1 }}>
          <Select
            value={draft.type}
            onChange={(e) => setDraft(clearedForType(draft, e.target.value as PaymentMethodType))}
            fullWidth
          >
            {METHOD_TYPES.map((t) => (
              <MenuItem key={t} value={t}>{PAYMENT_METHOD_LABELS[t]}</MenuItem>
            ))}
          </Select>
          <TextField
            label="Display name"
            value={draft.displayName ?? ''}
            onChange={(e) => set({ displayName: e.target.value })}
            placeholder={PAYMENT_METHOD_LABELS[draft.type]}
            helperText="What guests will see for this option."
          />
          <TextField
            label="Instructions"
            value={draft.instructions ?? ''}
            onChange={(e) => set({ instructions: e.target.value })}
            multiline
            minRows={2}
            helperText="e.g. Send as Friends & Family and add the reference code in the note."
          />
          {showsPaymentUrl(draft.type) && (
            <TextField
              label={requiresPaymentUrl(draft.type) ? 'Payment link' : 'Payment link (optional)'}
              required={requiresPaymentUrl(draft.type)}
              value={draft.paymentUrl ?? ''}
              onChange={(e) => set({ paymentUrl: e.target.value })}
              placeholder="https://paypal.me/… or wise.com/pay/me/…"
            />
          )}
          {showsContactFields(draft.type) && (
            <>
              {METHOD_FIELDS[draft.type].email && (
                <TextField
                  label="Recipient email"
                  value={draft.recipientEmail ?? ''}
                  onChange={(e) => set({ recipientEmail: e.target.value })}
                />
              )}
              {METHOD_FIELDS[draft.type].phone && (
                <TextField
                  label="Recipient phone"
                  value={draft.recipientPhone ?? ''}
                  onChange={(e) => set({ recipientPhone: e.target.value })}
                />
              )}
            </>
          )}
          {showsBankFields(draft.type) && (
            <>
              <Divider>Bank details</Divider>
              <TextField label="Bank name" value={draft.bankName ?? ''} onChange={(e) => set({ bankName: e.target.value })} />
              <TextField label="Account name" value={draft.bankAccountName ?? ''} onChange={(e) => set({ bankAccountName: e.target.value })} />
              <TextField label="Account number" value={draft.bankAccountNumber ?? ''} onChange={(e) => set({ bankAccountNumber: e.target.value })} />
              <TextField label="Routing / sort code" value={draft.bankRoutingNumber ?? ''} onChange={(e) => set({ bankRoutingNumber: e.target.value })} />
              <TextField label="IBAN" value={draft.bankIban ?? ''} onChange={(e) => set({ bankIban: e.target.value })} />
              <TextField label="SWIFT / BIC" value={draft.bankSwift ?? ''} onChange={(e) => set({ bankSwift: e.target.value })} />
            </>
          )}
          <FormControlLabel
            control={<Switch checked={draft.enabled ?? true} onChange={(e) => set({ enabled: e.target.checked })} />}
            label="Enabled (shown to guests)"
          />
          {error && <Alert severity="error">{error}</Alert>}
        </Stack>
      </DialogContent>
      <DialogActions>
        <Button onClick={onClose}>Cancel</Button>
        <Button
          variant="contained"
          onClick={onSave}
          disabled={saving || (requiresPaymentUrl(draft.type) && !draft.paymentUrl?.trim())}
          sx={{ color: '#fff' }}
        >
          {editing ? 'Save' : 'Add method'}
        </Button>
      </DialogActions>
    </Dialog>
  );
}
