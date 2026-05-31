import { useEffect, useState } from 'react';
import { Box, Stack, ToggleButton, ToggleButtonGroup, Typography } from '@mui/material';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { registries, type Registry } from '../../api';
import ApprovedGuestsSection from './ApprovedGuestsSection';

export default function AccessControlSection({ reg }: { reg: Registry }) {
  const qc = useQueryClient();
  const [requireGuestApproval, setRequireGuestApproval] = useState<boolean>(!(reg.allowOpenAccess ?? false));

  useEffect(() => {
    setRequireGuestApproval(!(reg.allowOpenAccess ?? false));
  }, [reg.id]);

  const saveAccessM = useMutation({
    mutationFn: (next: boolean) => registries.update(reg.id, { allowOpenAccess: !next }),
    onSuccess: () => qc.invalidateQueries({ queryKey: ['registries'] }),
  });

  return (
    <Stack spacing={3}>
      <Box
        sx={{
          p: 2,
          borderRadius: 2,
          bgcolor: requireGuestApproval ? 'primary.main' : 'background.default',
          border: '1px solid',
          borderColor: requireGuestApproval ? 'primary.main' : 'divider',
        }}
      >
        <Stack direction="row" alignItems="flex-start" spacing={1.5}>
          <LockOutlinedIcon
            fontSize="small"
            sx={{ color: requireGuestApproval ? '#fff' : 'text.secondary', mt: 0.3 }}
          />
          <Box sx={{ flex: 1 }}>
            <Typography variant="subtitle2" sx={{ fontWeight: 600, color: requireGuestApproval ? '#fff' : undefined }}>
              Require approval to view your registry
            </Typography>
            <Typography variant="caption" sx={{ color: requireGuestApproval ? 'rgba(255,255,255,0.85)' : 'text.secondary' }}>
              When on, guests must request access from you before they can see any of your registry. You decide who gets in.
            </Typography>
          </Box>
          <ToggleButtonGroup
            exclusive
            size="small"
            value={requireGuestApproval ? 'on' : 'off'}
            disabled={saveAccessM.isPending}
            onChange={(_, v) => {
              if (v == null) return;
              const next = v === 'on';
              setRequireGuestApproval(next);
              saveAccessM.mutate(next);
            }}
            sx={{
              '& .MuiToggleButton-root': {
                color: requireGuestApproval ? '#fff' : undefined,
                borderColor: requireGuestApproval ? 'rgba(255,255,255,0.5)' : undefined,
                '&.Mui-selected': requireGuestApproval
                  ? {
                      color: 'primary.main',
                      bgcolor: '#fff',
                      '&:hover': { bgcolor: '#fff' },
                    }
                  : undefined,
                '&:hover': requireGuestApproval ? { bgcolor: 'rgba(255,255,255,0.12)' } : undefined,
              },
            }}
          >
            <ToggleButton value="off" sx={{ textTransform: 'none', px: 2 }}>Off</ToggleButton>
            <ToggleButton value="on" sx={{ textTransform: 'none', px: 2 }}>On</ToggleButton>
          </ToggleButtonGroup>
        </Stack>
      </Box>
      <ApprovedGuestsSection registryId={reg.id} />
    </Stack>
  );
}
