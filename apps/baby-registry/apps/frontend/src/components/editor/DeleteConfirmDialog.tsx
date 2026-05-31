import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Typography,
} from '@mui/material';

export type DeleteTarget =
  | { kind: 'item'; id: string; title: string }
  | { kind: 'reservation'; id: string; title: string };

export default function DeleteConfirmDialog({
  target,
  onClose,
  onConfirm,
}: {
  target: DeleteTarget | null;
  onClose: () => void;
  onConfirm: () => void;
}) {
  return (
    <Dialog open={!!target} onClose={onClose} maxWidth="xs" fullWidth>
      <DialogTitle>Confirm deletion</DialogTitle>
      <DialogContent>
        <Typography variant="body2" color="text.secondary">
          {target?.kind === 'item'
            ? `Delete item "${target.title}"? This cannot be undone.`
            : `Delete reservation from "${target?.title}"? This cannot be undone.`}
        </Typography>
      </DialogContent>
      <DialogActions>
        <Button onClick={onClose}>Cancel</Button>
        <Button variant="contained" color="error" onClick={onConfirm}>
          Delete
        </Button>
      </DialogActions>
    </Dialog>
  );
}
