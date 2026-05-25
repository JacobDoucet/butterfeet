import { Box, Container, Typography, Button, Stack } from '@mui/material';
import { Link } from 'react-router-dom';
import BrandLogo from '../components/BrandLogo';

export default function Landing() {
  return (
    <Container maxWidth="md" sx={{ py: { xs: 6, md: 12 }, textAlign: 'center' }}>
      <Box sx={{ display: 'flex', justifyContent: 'center', mb: 2 }}>
        <BrandLogo variant="lockup" height={56} />
      </Box>
      <Typography variant="h2" gutterBottom>
        A gentle, beautiful baby registry.
      </Typography>
      <Typography variant="h6" color="text.secondary" sx={{ mb: 4 }}>
        Curate gifts from anywhere on the web. Share one link with friends and family.
        Track what's been picked up — without the spreadsheets.
      </Typography>
      <Stack direction={{ xs: 'column', sm: 'row' }} spacing={2} justifyContent="center">
        <Button component={Link} to="/login" size="large" variant="contained" color="primary">
          Create your registry
        </Button>
      </Stack>
      <Box sx={{ mt: 8, opacity: 0.9 }}>
        <Typography variant="body2" color="text.secondary">
          Import from Amazon, Mamas & Papas, Etsy, John Lewis, IKEA — or any product page.
        </Typography>
      </Box>
    </Container>
  );
}
