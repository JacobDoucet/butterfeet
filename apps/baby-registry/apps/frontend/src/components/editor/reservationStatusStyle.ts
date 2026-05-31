import type { ReservationStatus } from '../../api';

export interface StatusStyle {
  bg: string;
  fg: string;
  border: string;
}

export const STATUS_STYLE: Record<ReservationStatus, StatusStyle> = {
  Reserved:  { bg: '#FFF4E5', fg: '#8A4B00', border: '#FFB85C' },
  Purchased: { bg: '#E3F2FD', fg: '#0D3C61', border: '#64B5F6' },
  Received:  { bg: '#E8F5E9', fg: '#1B5E20', border: '#66BB6A' },
  Cancelled: { bg: '#F5F5F5', fg: '#616161', border: '#BDBDBD' },
};
