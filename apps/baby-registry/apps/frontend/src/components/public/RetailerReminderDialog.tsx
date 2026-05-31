import { Button, Dialog, DialogActions, DialogContent, DialogTitle, Stack, Typography } from '@mui/material';

export default function RetailerReminderDialog({
  open,
  onClose,
  onConfirm,
}: {
  open: boolean;
  onClose: () => void;
  onConfirm: () => void;
}) {
  return (
    <Dialog open={open} onClose={onClose} maxWidth="xs" fullWidth>
      <DialogTitle sx={{ fontWeight: 700 }}>Come back to confirm</DialogTitle>
      <DialogContent>
        <Stack spacing={1.5} sx={{ mt: 0.5 }}>
          <Typography variant="body2">
            We'll open the retailer in a new tab and hold this gift for you for 24 hours.
          </Typography>
          <Typography variant="body2" sx={{ fontWeight: 600 }}>
            Keep this tab open. Once you've completed checkout, come back here and tap
            “I've bought this” so the parents know it's on the way.
          </Typography>
          <Typography variant="body2" color="text.secondary">
            If you change your mind, just release the reservation so another guest can grab it.
          </Typography>
        </Stack>
      </DialogContent>
      <DialogActions sx={{ px: 3, py: 2 }}>
        <Button onClick={onClose}>Not yet</Button>
        <Button variant="contained" sx={{ color: '#fff' }} onClick={onConfirm}>
          Got it - open retailer
        </Button>
      </DialogActions>
    </Dialog>
  );
}
