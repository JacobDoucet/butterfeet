import { IconButton, MenuItem, Select, Stack, Tooltip, Typography } from '@mui/material';
import DeleteIcon from '@mui/icons-material/DeleteOutline';
import type { Reservation, ReservationStatus } from '../../api';

export default function ReservationRow({
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
        <MenuItem value="AwaitingConfirmation" disabled>
          Awaiting payment
        </MenuItem>
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
