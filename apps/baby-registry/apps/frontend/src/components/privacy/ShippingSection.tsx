import { useEffect, useState } from 'react';
import {
  Alert,
  Box,
  Button,
  Stack,
  TextField,
  ToggleButton,
  ToggleButtonGroup,
  Typography,
} from '@mui/material';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { registries, type AddressAccessMode, type Registry } from '../../api';

export default function ShippingSection({ reg }: { reg: Registry }) {
  const qc = useQueryClient();
  const [mode, setMode] = useState<AddressAccessMode>(reg.addressAccessMode ?? 'RequestApproval');
  const [recipientName, setRecipientName] = useState(reg.shippingRecipientName ?? '');
  const [line1, setLine1] = useState(reg.shippingLine1 ?? '');
  const [line2, setLine2] = useState(reg.shippingLine2 ?? '');
  const [city, setCity] = useState(reg.shippingCity ?? '');
  const [region, setRegion] = useState(reg.shippingRegion ?? '');
  const [postalCode, setPostalCode] = useState(reg.shippingPostalCode ?? '');
  const [country, setCountry] = useState(reg.shippingCountry ?? '');
  const [deliveryNotes, setDeliveryNotes] = useState(reg.shippingDeliveryNotes ?? '');
  const [saved, setSaved] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    setMode(reg.addressAccessMode ?? 'RequestApproval');
    setRecipientName(reg.shippingRecipientName ?? '');
    setLine1(reg.shippingLine1 ?? '');
    setLine2(reg.shippingLine2 ?? '');
    setCity(reg.shippingCity ?? '');
    setRegion(reg.shippingRegion ?? '');
    setPostalCode(reg.shippingPostalCode ?? '');
    setCountry(reg.shippingCountry ?? '');
    setDeliveryNotes(reg.shippingDeliveryNotes ?? '');
  }, [reg.id]);

  const saveM = useMutation({
    mutationFn: () =>
      registries.update(reg.id, {
        addressAccessMode: mode,
        shippingRecipientName: recipientName,
        shippingLine1: line1,
        shippingLine2: line2,
        shippingCity: city,
        shippingRegion: region,
        shippingPostalCode: postalCode,
        shippingCountry: country,
        shippingDeliveryNotes: deliveryNotes,
        shippingPolicyVersion: (reg.shippingPolicyVersion ?? 0) + 1,
      }),
    onSuccess: () => {
      setSaved(true);
      setError(null);
      qc.invalidateQueries({ queryKey: ['registries'] });
      window.setTimeout(() => setSaved(false), 2000);
    },
    onError: (err) => setError((err as Error).message),
  });

  const shippingHidden = mode === 'Disabled';

  return (
    <Stack spacing={3}>
      <Box
        sx={{
          p: 2,
          borderRadius: 2,
          bgcolor: shippingHidden ? 'background.default' : 'primary.main',
          border: '1px solid',
          borderColor: shippingHidden ? 'divider' : 'primary.main',
        }}
      >
        <Stack direction="row" alignItems="flex-start" spacing={1.5}>
          <LockOutlinedIcon
            fontSize="small"
            sx={{ color: shippingHidden ? 'text.secondary' : '#fff', mt: 0.3 }}
          />
          <Box sx={{ flex: 1 }}>
            <Typography variant="subtitle2" sx={{ fontWeight: 600, color: shippingHidden ? undefined : '#fff' }}>
              Shipping address visibility
            </Typography>
            <Typography variant="caption" sx={{ color: shippingHidden ? 'text.secondary' : 'rgba(255,255,255,0.85)' }}>
              {shippingHidden
                ? 'Hidden — no one can see your shipping address. Approved guests will need to message you for it.'
                : 'Shown — approved guests can see your shipping address from the registry.'}
            </Typography>
          </Box>
          <ToggleButtonGroup
            exclusive
            size="small"
            value={shippingHidden ? 'hide' : 'show'}
            onChange={(_, v) => v && setMode(v === 'show' ? 'ApprovedGuestsOnly' : 'Disabled')}
            sx={{
              '& .MuiToggleButton-root': {
                color: '#fff',
                borderColor: 'rgba(255,255,255,0.5)',
                '&.Mui-selected': {
                  color: 'primary.main',
                  bgcolor: '#fff',
                  '&:hover': { bgcolor: '#fff' },
                },
                '&:hover': { bgcolor: 'rgba(255,255,255,0.12)' },
              },
            }}
          >
            <ToggleButton value="show" sx={{ textTransform: 'none', px: 2 }}>Show</ToggleButton>
            <ToggleButton value="hide" sx={{ textTransform: 'none', px: 2 }}>Hide</ToggleButton>
          </ToggleButtonGroup>
        </Stack>
      </Box>

      <Box>
        <Typography variant="overline" color="text.secondary">Shipping address</Typography>
        <Stack spacing={2} sx={{ mt: 1, opacity: shippingHidden ? 0.55 : 1 }}>
          <TextField label="Recipient name" value={recipientName} onChange={(e) => setRecipientName(e.target.value)} disabled={shippingHidden} />
          <TextField label="Address line 1" value={line1} onChange={(e) => setLine1(e.target.value)} disabled={shippingHidden} />
          <TextField label="Address line 2" value={line2} onChange={(e) => setLine2(e.target.value)} disabled={shippingHidden} />
          <Stack direction={{ xs: 'column', sm: 'row' }} spacing={2}>
            <TextField label="City" value={city} onChange={(e) => setCity(e.target.value)} sx={{ flex: 1 }} disabled={shippingHidden} />
            <TextField label="Region / State" value={region} onChange={(e) => setRegion(e.target.value)} sx={{ flex: 1 }} disabled={shippingHidden} />
          </Stack>
          <Stack direction={{ xs: 'column', sm: 'row' }} spacing={2}>
            <TextField label="Postal code" value={postalCode} onChange={(e) => setPostalCode(e.target.value)} sx={{ flex: 1 }} disabled={shippingHidden} />
            <TextField label="Country" value={country} onChange={(e) => setCountry(e.target.value)} sx={{ flex: 1 }} disabled={shippingHidden} />
          </Stack>
          <TextField
            label="Delivery notes (optional)"
            value={deliveryNotes}
            onChange={(e) => setDeliveryNotes(e.target.value)}
            multiline
            minRows={2}
            helperText="Buzzer code, where to leave parcels, etc."
            disabled={shippingHidden}
          />
        </Stack>
      </Box>

      {error && <Alert severity="error">{error}</Alert>}
      {saved && <Alert severity="success">Saved.</Alert>}

      <Stack direction="row" justifyContent="flex-end">
        <Button variant="contained" onClick={() => saveM.mutate()} disabled={saveM.isPending} sx={{ color: '#fff' }}>
          Save shipping settings
        </Button>
      </Stack>
    </Stack>
  );
}
