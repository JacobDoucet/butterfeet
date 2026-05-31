import { useState } from 'react';
import {
  Alert,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Stack,
  TextField,
} from '@mui/material';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { registries, type Me, type Registry } from '../../api';

function mapRegistryCreateError(err: unknown, slug: string): string {
  const raw = (err as Error)?.message?.trim() || 'Could not create registry.';
  const msg = raw.toLowerCase();
  if (
    msg.includes('slug_taken') ||
    msg.includes('e11000') ||
    msg.includes('duplicate key') ||
    msg.includes('slug_unique') ||
    msg === 'unexpected' ||
    msg.includes('unexpected')
  ) {
    return `That slug is already taken. Try a different one (for example: ${slug}-2).`;
  }
  return raw;
}

export default function NewRegistryDialog({
  open,
  onClose,
  me,
  existingSlugs,
  onCreated,
}: {
  open: boolean;
  onClose: () => void;
  me: Me;
  existingSlugs: string[];
  onCreated: (reg: Registry) => void;
}) {
  const qc = useQueryClient();
  const [slug, setSlug] = useState('');
  const [title, setTitle] = useState('');
  const [parentNames, setParentNames] = useState('');
  const [error, setError] = useState<string | null>(null);

  const close = () => {
    onClose();
    setError(null);
  };

  const createM = useMutation({
    mutationFn: async () => {
      const normalizedSlug = slug.trim().toLowerCase();
      if (existingSlugs.includes(normalizedSlug)) throw new Error('slug_taken');
      return registries.create({
        slug: normalizedSlug,
        title: title.trim(),
        parentNames: parentNames.trim(),
        addressAccessMode: 'RequestApproval',
        isPublic: true,
        ownerId: me.id,
      });
    },
    onSuccess: (reg) => {
      qc.invalidateQueries({ queryKey: ['registries'] });
      close();
      onCreated(reg);
    },
    onError: (err) => setError(mapRegistryCreateError(err, slug.trim().toLowerCase() || 'my-registry')),
  });

  return (
    <Dialog open={open} onClose={close} fullWidth maxWidth="sm">
      <DialogTitle>New registry</DialogTitle>
      <DialogContent>
        <Stack spacing={2} sx={{ mt: 1 }}>
          <TextField label="Title (e.g. Baby Smith)" value={title} onChange={(e) => setTitle(e.target.value)} />
          <TextField
            label="URL slug"
            helperText="lowercase letters, numbers, dashes"
            value={slug}
            onChange={(e) => setSlug(e.target.value.replace(/[^a-z0-9-]/g, ''))}
          />
          <TextField label="Parent names (optional)" value={parentNames} onChange={(e) => setParentNames(e.target.value)} />
          {error && <Alert severity="error">{error}</Alert>}
        </Stack>
      </DialogContent>
      <DialogActions>
        <Button onClick={close}>Cancel</Button>
        <Button onClick={() => createM.mutate()} variant="contained" disabled={!slug || !title}>
          Create
        </Button>
      </DialogActions>
    </Dialog>
  );
}
