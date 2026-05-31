import { useState } from 'react';
import {
  Box,
  Card,
  CardContent,
  Chip,
  IconButton,
  MenuItem,
  Select,
  Stack,
  Tooltip,
  Typography,
} from '@mui/material';
import DeleteIcon from '@mui/icons-material/DeleteOutline';
import { format, formatDistanceToNowStrict, isToday, isYesterday } from 'date-fns';
import type { RegistryItem, Reservation, ReservationStatus } from '../../api';
import { STATUS_STYLE } from './reservationStatusStyle';

const STATUS_VALUES: ReservationStatus[] = ['Reserved', 'Purchased', 'Received', 'Cancelled'];

export default function PurchasesPanel({
  reservations: allReservations,
  itemById,
  onOpenItem,
  onSetStatus,
  onDelete,
}: {
  reservations: Reservation[];
  itemById: Record<string, RegistryItem>;
  onOpenItem: (itemId: string) => void;
  onSetStatus: (id: string, status: ReservationStatus) => void;
  onDelete: (reservation: Reservation) => void;
}) {
  const [statusFilter, setStatusFilter] = useState<'all' | ReservationStatus>('all');

  const counts = allReservations.reduce<Record<ReservationStatus, number>>(
    (acc, r) => {
      acc[r.status] = (acc[r.status] ?? 0) + 1;
      return acc;
    },
    { Reserved: 0, Purchased: 0, Received: 0, Cancelled: 0 },
  );

  const filtered = allReservations
    .filter((r) => (statusFilter === 'all' ? true : r.status === statusFilter))
    .slice()
    .sort((a, b) => String(b.created?.at ?? '').localeCompare(String(a.created?.at ?? '')));

  return (
    <Stack spacing={2}>
      <Stack direction="row" spacing={1} sx={{ flexWrap: 'wrap', rowGap: 1 }}>
        <Chip
          label={`All (${allReservations.length})`}
          onClick={() => setStatusFilter('all')}
          color={statusFilter === 'all' ? 'primary' : 'default'}
          variant={statusFilter === 'all' ? 'filled' : 'outlined'}
        />
        {STATUS_VALUES.map((s) => {
          const style = STATUS_STYLE[s];
          const active = statusFilter === s;
          return (
            <Chip
              key={s}
              label={`${s} (${counts[s]})`}
              onClick={() => setStatusFilter(s)}
              sx={{
                bgcolor: active ? style.bg : 'transparent',
                color: style.fg,
                borderColor: style.border,
                border: '1px solid',
                fontWeight: active ? 600 : 400,
              }}
            />
          );
        })}
      </Stack>

      {filtered.length === 0 ? (
        <Typography color="text.secondary">
          {statusFilter === 'all' ? 'No reservations yet.' : `No ${statusFilter.toLowerCase()} reservations.`}
        </Typography>
      ) : (
        filtered.map((r) => (
          <PurchaseCard
            key={r.id}
            reservation={r}
            item={itemById[r.itemId]}
            onOpenItem={() => onOpenItem(r.itemId)}
            onSetStatus={(status) => onSetStatus(r.id, status)}
            onDelete={() => onDelete(r)}
          />
        ))
      )}
    </Stack>
  );
}

function PurchaseCard({
  reservation: r,
  item,
  onOpenItem,
  onSetStatus,
  onDelete,
}: {
  reservation: Reservation;
  item: RegistryItem | undefined;
  onOpenItem: () => void;
  onSetStatus: (status: ReservationStatus) => void;
  onDelete: () => void;
}) {
  const who = r.isAnonymous
    ? 'Anonymous'
    : r.reserverName?.trim() || r.contactEmail?.trim() || 'Someone';
  const qty = r.quantity ?? 1;
  const style = STATUS_STYLE[r.status];

  return (
    <Card variant="outlined" sx={{ borderLeft: '4px solid', borderLeftColor: style.border }}>
      <CardContent>
        <Stack direction={{ xs: 'column', sm: 'row' }} spacing={2}>
          {item?.imageUrl && (
            <Box
              onClick={onOpenItem}
              sx={{
                width: 96,
                height: 96,
                flexShrink: 0,
                borderRadius: 1,
                overflow: 'hidden',
                bgcolor: item.imageBgColor || 'grey.100',
                cursor: 'pointer',
              }}
            >
              <Box
                component="img"
                src={item.imageUrl}
                alt=""
                sx={{ width: '100%', height: '100%', objectFit: 'contain' }}
              />
            </Box>
          )}
          <Stack sx={{ flex: 1, minWidth: 0 }} spacing={0.5}>
            <Stack direction="row" spacing={1} alignItems="center" sx={{ flexWrap: 'wrap', rowGap: 0.5 }}>
              <Typography
                variant="subtitle1"
                sx={{ fontWeight: 600, cursor: item ? 'pointer' : 'default' }}
                onClick={() => item && onOpenItem()}
              >
                {item?.title || 'Unknown item'}
              </Typography>
              {qty > 1 && <Chip size="small" label={`×${qty}`} />}
            </Stack>
            <Typography variant="body2">
              <strong>{who}</strong>
              {!r.isAnonymous && r.contactEmail && (
                <Typography component="span" variant="body2" color="text.secondary">
                  {' '}· {r.contactEmail}
                </Typography>
              )}
            </Typography>
            {r.message && (
              <Typography
                variant="body2"
                color="text.secondary"
                sx={{ whiteSpace: 'pre-wrap', mt: 0.5 }}
              >
                “{r.message}”
              </Typography>
            )}
            <CreatedAt at={r.created?.at} />
          </Stack>
          <Stack direction="row" spacing={1} alignItems="center" sx={{ flexShrink: 0 }}>
            <Select
              size="small"
              value={r.status}
              onChange={(e) => onSetStatus(e.target.value as ReservationStatus)}
              sx={{
                minWidth: 130,
                bgcolor: style.bg,
                color: style.fg,
                fontWeight: 600,
                '& .MuiOutlinedInput-notchedOutline': { borderColor: style.border },
                '&:hover .MuiOutlinedInput-notchedOutline': { borderColor: style.border },
                '& .MuiSelect-icon': { color: style.fg },
              }}
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
        </Stack>
      </CardContent>
    </Card>
  );
}

function CreatedAt({ at }: { at?: string }) {
  if (!at) return null;
  const d = new Date(at);
  if (isNaN(d.getTime())) return null;
  const relative = formatDistanceToNowStrict(d, { addSuffix: true });
  let absolute: string;
  if (isToday(d)) absolute = `Today at ${format(d, 'h:mm a')}`;
  else if (isYesterday(d)) absolute = `Yesterday at ${format(d, 'h:mm a')}`;
  else absolute = format(d, 'MMM d, yyyy · h:mm a');
  return (
    <Tooltip title={absolute}>
      <Typography variant="caption" color="text.secondary">
        {absolute} · {relative}
      </Typography>
    </Tooltip>
  );
}
