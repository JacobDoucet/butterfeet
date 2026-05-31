import { useEffect, useState } from 'react';
import {
  Stack,
  Typography,
  TextField,
  Button,
  Alert,
  Box,
} from '@mui/material';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { registries, type Registry } from '../api';
import { THEME_PALETTES, DEFAULT_PALETTE } from '../themePalettes';

export default function BasicInfoPanel({ reg }: { reg: Registry }) {
  const qc = useQueryClient();
  const [title, setTitle] = useState(reg.title ?? '');
  const [parentNames, setParentNames] = useState(reg.parentNames ?? '');
  const [welcomeMessage, setWelcomeMessage] = useState(reg.welcomeMessage ?? '');
  const [themeColor, setThemeColor] = useState(reg.themeColor ?? '');
  const [saved, setSaved] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    setTitle(reg.title ?? '');
    setParentNames(reg.parentNames ?? '');
    setWelcomeMessage(reg.welcomeMessage ?? '');
    setThemeColor(reg.themeColor ?? '');
  }, [reg.id]);

  const saveM = useMutation({
    mutationFn: () =>
      registries.update(reg.id, {
        title: title.trim(),
        parentNames: parentNames.trim(),
        welcomeMessage: welcomeMessage.trim(),
        themeColor: themeColor.trim(),
      }),
    onSuccess: () => {
      setSaved(true);
      setError(null);
      qc.invalidateQueries({ queryKey: ['registries'] });
      window.setTimeout(() => setSaved(false), 2000);
    },
    onError: (err) => setError((err as Error).message),
  });

  return (
    <Stack spacing={3}>
      <Box>
        <Typography variant="overline" color="text.secondary">Basics</Typography>
        <Stack spacing={2} sx={{ mt: 1 }}>
          <TextField
            label="Registry name"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
          />
          <TextField
            label="Parents' names"
            value={parentNames}
            onChange={(e) => setParentNames(e.target.value)}
            placeholder="e.g. Sam & Riley"
          />
          <TextField
            label="Welcome message"
            value={welcomeMessage}
            onChange={(e) => setWelcomeMessage(e.target.value)}
            multiline
            minRows={3}
            helperText="Shown at the top of your public registry page."
          />
        </Stack>
      </Box>

      <Box>
        <Typography variant="overline" color="text.secondary">Look & feel</Typography>
        <Stack direction="row" spacing={1.5} alignItems="center" flexWrap="wrap" sx={{ mt: 1 }}>
          {THEME_PALETTES.map((p) => {
            const effective = themeColor || DEFAULT_PALETTE.bg;
            const selected = effective.toLowerCase() === p.bg.toLowerCase();
            return (
              <Box
                key={p.bg}
                role="button"
                aria-label={`Use ${p.bg}`}
                onClick={() => setThemeColor(p.bg)}
                sx={{
                  width: 32,
                  height: 32,
                  borderRadius: '50%',
                  cursor: 'pointer',
                  border: '2px solid',
                  borderColor: selected ? 'text.primary' : 'transparent',
                  boxShadow: selected ? 'none' : 'inset 0 0 0 1px rgba(0,0,0,0.12)',
                  background: `linear-gradient(135deg, ${p.bg} 0%, ${p.bg} 50%, ${p.primary} 50%, ${p.primary} 100%)`,
                }}
              />
            );
          })}
        </Stack>
      </Box>

      {error && <Alert severity="error">{error}</Alert>}

      <Stack direction="row" justifyContent="flex-end">
        <Button
          variant="contained"
          onClick={() => saveM.mutate()}
          disabled={saveM.isPending || !title.trim()}
        >
          Save details
        </Button>
      </Stack>
    </Stack>
  );
}
