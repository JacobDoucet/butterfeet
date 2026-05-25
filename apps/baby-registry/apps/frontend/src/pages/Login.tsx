import { useState } from 'react';
import { Container, Typography, TextField, Button, Stack, Alert, Card, CardContent } from '@mui/material';
import { auth } from '../api';

export default function Login() {
  const [email, setEmail] = useState('');
  const [name, setName] = useState('');
  const [sent, setSent] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const onSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    setLoading(true);
    try {
      await auth.request(email, name);
      setSent(true);
    } catch (err) {
      setError((err as Error).message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <Container maxWidth="sm" sx={{ py: 8 }}>
      <Card>
        <CardContent sx={{ p: 4 }}>
          <Typography variant="h4" gutterBottom>
            Sign in
          </Typography>
          <Typography color="text.secondary" sx={{ mb: 3 }}>
            We'll email you a magic link. No passwords.
          </Typography>
          {sent ? (
            <Alert severity="success">
              Check your inbox (or the server log in dev) for your sign-in link.
            </Alert>
          ) : (
            <form onSubmit={onSubmit}>
              <Stack spacing={2}>
                <TextField
                  label="Email"
                  type="email"
                  required
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                />
                <TextField
                  label="Your name (optional, for first-time)"
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                />
                {error && <Alert severity="error">{error}</Alert>}
                <Button type="submit" variant="contained" size="large" disabled={loading}>
                  {loading ? 'Sending…' : 'Email me a link'}
                </Button>
              </Stack>
            </form>
          )}
        </CardContent>
      </Card>
    </Container>
  );
}
