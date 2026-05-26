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
        Feather the nest, your way.
      </Typography>
      <Typography variant="h6" color="text.secondary" sx={{ mb: 4 }}>
        Drop a link from anywhere on the web, send it to your people,
        and we'll quietly keep track of who got what.
      </Typography>
      <Stack direction={{ xs: 'column', sm: 'row' }} spacing={2} justifyContent="center">
        <Button component={Link} to="/login" size="large" variant="contained" color="primary">
          Start your nest
        </Button>
      </Stack>
      <Box sx={{ mt: 8, opacity: 0.9 }}>
        <Typography variant="body2" color="text.secondary">
          Plays nice with Amazon, IKEA, John Lewis, Etsy, Mamas &amp; Papas,
          and pretty much any product page.
        </Typography>
      </Box>
    </Container>
  );
}
