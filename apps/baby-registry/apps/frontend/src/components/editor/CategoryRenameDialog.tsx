import {
  Alert,
  Autocomplete,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Stack,
  TextField,
} from '@mui/material';

export default function CategoryRenameDialog({
  open,
  onClose,
  categories,
  from,
  to,
  onFromChange,
  onToChange,
  error,
  pending,
  onSubmit,
}: {
  open: boolean;
  onClose: () => void;
  categories: string[];
  from: string;
  to: string;
  onFromChange: (v: string) => void;
  onToChange: (v: string) => void;
  error: string | null;
  pending: boolean;
  onSubmit: () => void;
}) {
  return (
    <Dialog open={open} onClose={onClose} fullWidth maxWidth="xs">
      <DialogTitle>Rename category</DialogTitle>
      <DialogContent>
        <Stack spacing={2} sx={{ mt: 1 }}>
          <Autocomplete
            freeSolo
            options={categories}
            value={from}
            inputValue={from}
            onInputChange={(_, v) => onFromChange(v)}
            onChange={(_, v) => onFromChange(typeof v === 'string' ? v : v ?? '')}
            renderInput={(params) => <TextField {...params} label="Current category" autoFocus />}
          />
          <Autocomplete
            freeSolo
            options={categories}
            value={to}
            inputValue={to}
            onInputChange={(_, v) => onToChange(v)}
            onChange={(_, v) => onToChange(typeof v === 'string' ? v : v ?? '')}
            renderInput={(params) => (
              <TextField
                {...params}
                label="New category"
                helperText="Leave blank to clear the category on matching items."
              />
            )}
          />
          {error && <Alert severity="error">{error}</Alert>}
        </Stack>
      </DialogContent>
      <DialogActions>
        <Button onClick={onClose}>Cancel</Button>
        <Button
          variant="contained"
          disabled={pending || from.trim() === to.trim()}
          onClick={onSubmit}
        >
          {pending ? 'Renaming…' : 'Rename'}
        </Button>
      </DialogActions>
    </Dialog>
  );
}
