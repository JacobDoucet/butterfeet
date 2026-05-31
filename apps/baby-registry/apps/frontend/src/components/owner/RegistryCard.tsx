import { Link } from 'react-router-dom';
import { Button, Card, CardContent, Stack, Typography } from '@mui/material';
import type { Registry } from '../../api';

export default function RegistryCard({ reg }: { reg: Registry }) {
  return (
    <Card>
      <CardContent>
        <Typography variant="h6">{reg.title}</Typography>
        <Typography color="text.secondary" sx={{ mb: 2 }}>
          /r/{reg.slug}
        </Typography>
        <Stack direction="row" spacing={1}>
          <Button size="small" component={Link} to={`/owner/r/${reg.slug}`}>
            Manage
          </Button>
          <Button size="small" component="a" href={`/r/${reg.slug}`} target="_blank" rel="noreferrer">
            View public
          </Button>
        </Stack>
      </CardContent>
    </Card>
  );
}
