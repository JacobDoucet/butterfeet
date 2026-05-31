import { Box, Container, Typography, Button, Stack } from '@mui/material';
import { Link } from 'react-router-dom';
import BrandLogo from '../components/BrandLogo';

export default function Landing() {
  return (
    <Container maxWidth="md" sx={{ py: { xs: 6, md: 12 }, textAlign: 'center' }}>
      <Box sx={{ display: 'flex', justifyContent: 'center', mb: 2, maxWidth: '100%', overflow: 'hidden' }}>
        <BrandLogo variant="lockup" height={60} markScale={1.3} wordmarkScale={1.26} />
      </Box>
      <Typography
        variant="h2"
        gutterBottom
        sx={{ fontSize: { xs: 'clamp(1.4rem, 6vw, 3rem)', md: '3.75rem' } }}
      >
        A baby registry for everywhere you actually shop.
      </Typography>
      <Typography variant="h6" color="text.secondary" sx={{ mb: 4, maxWidth: 640, mx: 'auto' }}>
        Add any product from any site. Share one link with friends and family.
        We make sure no one gifts the same thing twice.
      </Typography>
      <Stack direction={{ xs: 'column', sm: 'row' }} spacing={2} justifyContent="center">
        <Button component={Link} to="/login" size="large" variant="contained" color="primary">
          Create your registry — free
        </Button>
      </Stack>
      <Box sx={{ mt: 6 }}>
        <Typography variant="body2" color="text.secondary">
          Works with Amazon, IKEA, John Lewis, Etsy, Mamas &amp; Papas,
          and any product page on the web.
        </Typography>
      </Box>
      <Box sx={{ mt: 3 }}>
        <Typography variant="caption" color="text.secondary" sx={{ letterSpacing: '0.04em' }}>
          No ads. No selling your data. Just your registry.
        </Typography>
      </Box>
    </Container>
  );
}

