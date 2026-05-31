import { Alert, IconButton, Stack, TextField, Tooltip, Typography } from '@mui/material';
import ContentCopyIcon from '@mui/icons-material/ContentCopy';

export interface ShareLink {
  email: string;
  url: string;
  expiresAt: string;
}

export default function ShareLinkAlert({
  link,
  onClose,
  expiresSuffix,
}: {
  link: ShareLink;
  onClose: () => void;
  expiresSuffix?: string;
}) {
  return (
    <Alert
      severity="success"
      sx={{ mb: 2, '& .MuiAlert-message': { width: '100%' } }}
      onClose={onClose}
    >
      <Typography variant="body2" sx={{ fontWeight: 500, mb: 0.5 }}>
        Share this link with {link.email}
      </Typography>
      <Stack direction="row" spacing={1} alignItems="center">
        <TextField
          size="small"
          value={link.url}
          fullWidth
          InputProps={{ readOnly: true }}
          onFocus={(e) => e.target.select()}
        />
        <Tooltip title="Copy">
          <IconButton size="small" onClick={() => navigator.clipboard.writeText(link.url)}>
            <ContentCopyIcon fontSize="small" />
          </IconButton>
        </Tooltip>
      </Stack>
      <Typography variant="caption" color="text.secondary">
        Expires {new Date(link.expiresAt).toLocaleString()}.
        {expiresSuffix ? ` ${expiresSuffix}` : ''}
      </Typography>
    </Alert>
  );
}
