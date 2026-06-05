import { useEffect, useMemo, useState } from 'react';
import {
  Alert,
  Box,
  Button,
  Divider,
  Drawer,
  FormControl,
  FormControlLabel,
  IconButton,
  Link as MuiLink,
  Radio,
  RadioGroup,
  Stack,
  Typography,
  useMediaQuery,
  useTheme,
} from '@mui/material';
import CloseIcon from '@mui/icons-material/Close';
import ShoppingBagOutlinedIcon from '@mui/icons-material/ShoppingBagOutlined';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import CheckCircleRoundedIcon from '@mui/icons-material/CheckCircleRounded';
import OpenInNewIcon from '@mui/icons-material/OpenInNew';
import {
  formatPriceCents,
  convertCents,
  PAYMENT_METHOD_LABELS,
  type PaymentMethod,
  type PaymentIntent,
  type ExchangeRates,
} from '../../api';

export interface CartLine {
  reservationId: string;
  itemId: string;
  title: string;
  imageUrl?: string;
  imageBgColor?: string;
  priceCents?: number;
  currency?: string;
  quantity: number;
  productUrl?: string;
  retailer?: string;
}

interface CartDrawerProps {
  open: boolean;
  onClose: () => void;
  lines: CartLine[];
  buyerEmail: string;
  onRemove: (reservationId: string) => void;
  removingId?: string | null;
  onViewProduct?: (line: CartLine) => void;
  onContinueShopping: () => void;
  paymentMethods: PaymentMethod[];
  onCreatePaymentIntent: (paymentMethodId: string) => Promise<PaymentIntent>;
  onClaimPayment: (paymentId: string, message: string) => Promise<void>;
  onCancelPaymentIntent?: (paymentId: string) => Promise<void>;
  viewerCurrency?: string;
  rates?: ExchangeRates;
}

type Step = 'cart' | 'checkout' | 'pay' | 'done';

export default function CartDrawer({
  open,
  onClose,
  lines,
  buyerEmail,
  onRemove,
  removingId,
  onViewProduct,
  onContinueShopping,
  paymentMethods,
  onCreatePaymentIntent,
  onClaimPayment,
  onCancelPaymentIntent,
  viewerCurrency,
  rates,
}: CartDrawerProps) {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('sm'));
  const [step, setStep] = useState<Step>('cart');
  const [selectedMethodId, setSelectedMethodId] = useState('');
  const [intent, setIntent] = useState<PaymentIntent | null>(null);
  const [payError, setPayError] = useState<string | null>(null);
  const [creatingIntent, setCreatingIntent] = useState(false);
  const [claiming, setClaiming] = useState(false);

  const enabledMethods = useMemo(
    () => paymentMethods.filter((m) => m.enabled !== false),
    [paymentMethods],
  );
  const selectedMethod = useMemo(
    () => enabledMethods.find((m) => m.id === selectedMethodId) ?? null,
    [enabledMethods, selectedMethodId],
  );

  // Reset to the cart view each time the drawer opens.
  useEffect(() => {
    if (open) {
      setStep('cart');
      setIntent(null);
      setPayError(null);
    }
  }, [open]);

  // Default the selected method to the first enabled one.
  useEffect(() => {
    if (!selectedMethodId && enabledMethods.length > 0) {
      setSelectedMethodId(enabledMethods[0].id);
    }
  }, [enabledMethods, selectedMethodId]);

  // If the cart empties while open during checkout, fall back to the cart
  // view. The pay step is excluded: by then the gifts are locked into the
  // cart (so the open-cart list is intentionally empty) and the order details
  // live on the intent.
  useEffect(() => {
    if (lines.length === 0 && step === 'checkout') setStep('cart');
  }, [lines.length, step]);

  const handleProceedToPayment = async () => {
    if (!selectedMethod) {
      setPayError('Please choose a payment method.');
      return;
    }
    setPayError(null);
    setCreatingIntent(true);
    try {
      const created = await onCreatePaymentIntent(selectedMethod.id);
      setIntent(created);
      setStep('pay');
    } catch (err) {
      setPayError(err instanceof Error ? err.message : 'Could not start payment.');
    } finally {
      setCreatingIntent(false);
    }
  };

  const handleClaim = async () => {
    if (!intent) return;
    setClaiming(true);
    setPayError(null);
    try {
      await onClaimPayment(intent.id, '');
      setStep('done');
    } catch (err) {
      setPayError(err instanceof Error ? err.message : 'Could not confirm payment.');
    } finally {
      setClaiming(false);
    }
  };

  // Backing out of the pay step releases the locked gifts so they return to
  // the open cart.
  const handlePayBack = async () => {
    if (intent && onCancelPaymentIntent) {
      try {
        await onCancelPaymentIntent(intent.id);
      } catch {
        // Non-fatal: the gifts stay reserved if the release fails.
      }
    }
    setIntent(null);
    setStep('checkout');
  };

  const { subtotalCents, hasAnyPrice, currency } = useMemo(() => {
    let sum = 0;
    let any = false;
    let cur: string | undefined;
    for (const l of lines) {
      if (l.priceCents != null) {
        sum += l.priceCents * l.quantity;
        any = true;
        cur = cur ?? l.currency;
      }
    }
    return { subtotalCents: sum, hasAnyPrice: any, currency: cur };
  }, [lines]);

  const target = (viewerCurrency || '').toUpperCase();

  // displayPrice formats an amount in the viewer's chosen currency, falling
  // back to the original when no rate is available. `approx` marks converted
  // estimates so the UI can prefix them with "≈".
  const displayPrice = (
    cents?: number,
    cur?: string,
  ): { text: string; approx: boolean } | null => {
    if (cents == null || Number.isNaN(cents)) return null;
    const from = (cur || 'USD').toUpperCase();
    if (!target || from === target) {
      const text = formatPriceCents(cents, from);
      return text ? { text, approx: false } : null;
    }
    const converted = convertCents(cents, from, target, rates);
    if (converted == null) {
      const text = formatPriceCents(cents, from);
      return text ? { text, approx: false } : null;
    }
    const text = formatPriceCents(converted, target);
    return text ? { text, approx: true } : null;
  };

  const fmt = (dp: { text: string; approx: boolean } | null): string | null =>
    dp ? (dp.approx ? `≈ ${dp.text}` : dp.text) : null;

  const itemCount = lines.reduce((n, l) => n + l.quantity, 0);
  const subtotalLabel = hasAnyPrice ? fmt(displayPrice(subtotalCents, currency)) : null;

  const renderLine = (line: CartLine) => {
    const price = fmt(displayPrice(line.priceCents, line.currency));
    const lineTotal =
      line.priceCents != null ? fmt(displayPrice(line.priceCents * line.quantity, line.currency)) : null;
    return (
      <Stack key={line.reservationId} direction="row" spacing={1.5} sx={{ py: 2 }}>
        <Box
          sx={{
            width: 72,
            height: 72,
            flexShrink: 0,
            borderRadius: 2,
            bgcolor: line.imageBgColor || '#ffffff',
            border: '1px solid',
            borderColor: 'divider',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            overflow: 'hidden',
          }}
        >
          {line.imageUrl ? (
            <Box
              component="img"
              src={line.imageUrl}
              alt={line.title}
              sx={{ width: '100%', height: '100%', objectFit: 'contain' }}
            />
          ) : (
            <Typography variant="caption" color="text.disabled">
              No image
            </Typography>
          )}
        </Box>
        <Stack spacing={0.25} sx={{ flex: 1, minWidth: 0 }}>
          <Typography
            variant="subtitle2"
            sx={{
              fontWeight: 600,
              lineHeight: 1.3,
              display: '-webkit-box',
              WebkitLineClamp: 2,
              WebkitBoxOrient: 'vertical',
              overflow: 'hidden',
            }}
          >
            {line.title}
          </Typography>
          <Typography variant="caption" color="text.secondary">
            Qty {line.quantity}
            {price ? ` · ${price} each` : ''}
          </Typography>
          <Stack direction="row" spacing={1.5} alignItems="center" sx={{ mt: 0.5 }}>
            <Button
              size="small"
              color="inherit"
              onClick={() => onRemove(line.reservationId)}
              disabled={removingId === line.reservationId}
              sx={{ p: 0, minWidth: 0, color: 'text.secondary', textTransform: 'none', fontWeight: 500 }}
            >
              {removingId === line.reservationId ? 'Removing…' : 'Remove'}
            </Button>
            {line.productUrl && onViewProduct && (
              <Button
                size="small"
                color="inherit"
                onClick={() => onViewProduct(line)}
                startIcon={<OpenInNewIcon sx={{ fontSize: 14 }} />}
                sx={{ p: 0, minWidth: 0, color: 'text.secondary', textTransform: 'none', fontWeight: 500 }}
              >
                View item
              </Button>
            )}
          </Stack>
        </Stack>
        {lineTotal && (
          <Typography variant="subtitle2" sx={{ fontWeight: 700, whiteSpace: 'nowrap' }}>
            {lineTotal}
          </Typography>
        )}
      </Stack>
    );
  };

  const header = (title: string, back?: Step, onBack?: () => void) => (
    <Stack
      direction="row"
      alignItems="center"
      spacing={1}
      sx={{ px: 2.5, py: 2, borderBottom: 1, borderColor: 'divider' }}
    >
      {back ? (
        <IconButton edge="start" onClick={() => (onBack ? onBack() : setStep(back))} aria-label="Back">
          <ArrowBackIcon />
        </IconButton>
      ) : (
        <ShoppingBagOutlinedIcon sx={{ color: 'primary.main' }} />
      )}
      <Typography variant="h6" sx={{ fontWeight: 700, flex: 1 }}>
        {title}
      </Typography>
      <IconButton edge="end" onClick={onClose} aria-label="Close cart">
        <CloseIcon />
      </IconButton>
    </Stack>
  );

  let body: React.ReactNode;

  if (step === 'done') {
    body = (
      <Stack spacing={2.5} sx={{ px: 3, py: 6, textAlign: 'center', alignItems: 'center', flex: 1, justifyContent: 'center' }}>
        <CheckCircleRoundedIcon sx={{ fontSize: 64, color: 'success.main' }} />
        <Typography variant="h5" sx={{ fontWeight: 700 }}>
          Thank you!
        </Typography>
        <Typography color="text.secondary" sx={{ maxWidth: 340 }}>
          We've let the parents know you've sent your payment. They'll confirm once the money
          arrives, and your gifts stay reserved in the meantime.
        </Typography>
        {intent && (
          <Typography variant="body2" color="text.secondary">
            Reference code: <strong>{intent.referenceCode}</strong>
          </Typography>
        )}
        <Button variant="contained" onClick={onClose} sx={{ color: '#fff', mt: 1 }}>
          Back to the registry
        </Button>
      </Stack>
    );
  } else if (lines.length === 0 && step !== 'pay') {
    body = (
      <Box sx={{ flex: 1, minHeight: 0, overflowY: 'auto', display: 'flex', flexDirection: 'column' }}>
        <Stack spacing={2} sx={{ px: 3, py: 8, textAlign: 'center', alignItems: 'center', flex: 1, justifyContent: 'center' }}>
          <ShoppingBagOutlinedIcon sx={{ fontSize: 56, color: 'text.disabled' }} />
          <Typography variant="h6" sx={{ fontWeight: 700 }}>
            Your cart is empty
          </Typography>
          <Typography color="text.secondary" sx={{ maxWidth: 300 }}>
            Add a gift from the registry and it will be held just for you.
          </Typography>
          <Button variant="contained" onClick={onContinueShopping} sx={{ color: '#fff', mt: 1 }}>
            Browse the registry
          </Button>
        </Stack>
      </Box>
    );
  } else if (step === 'cart') {
    body = (
      <>
        <Box sx={{ flex: 1, minHeight: 0, overflowY: 'auto' }}>
          <Box sx={{ px: 2.5 }}>
            <Stack divider={<Divider />}>{lines.map(renderLine)}</Stack>
          </Box>
        </Box>
        <Box sx={{ px: 2.5, py: 2, borderTop: 1, borderColor: 'divider' }}>
          <Stack direction="row" justifyContent="space-between" alignItems="baseline" sx={{ mb: 0.5 }}>
            <Typography sx={{ fontWeight: 600 }}>Subtotal</Typography>
            <Typography sx={{ fontWeight: 700 }}>{subtotalLabel ?? '—'}</Typography>
          </Stack>
          <Typography variant="caption" color="text.secondary" sx={{ display: 'block', mb: 2 }}>
            {itemCount} {itemCount === 1 ? 'gift' : 'gifts'} held for you · taxes & shipping calculated at checkout
          </Typography>
          <Button fullWidth variant="contained" size="large" onClick={() => setStep('checkout')} sx={{ color: '#fff' }}>
            Proceed to checkout
          </Button>
          <Button fullWidth onClick={onContinueShopping} sx={{ mt: 1 }}>
            Keep browsing
          </Button>
        </Box>
      </>
    );
  } else if (step === 'checkout') {
    // checkout
    body = (
      <>
        <Box sx={{ flex: 1, minHeight: 0, overflowY: 'auto', px: 2.5, py: 2 }}>
          <Stack spacing={2.5}>
            <Box>
              <Typography variant="overline" color="text.secondary">
                Order summary
              </Typography>
              <Stack spacing={1} sx={{ mt: 1 }}>
                {lines.map((line) => {
                  const lineTotal =
                    line.priceCents != null
                      ? fmt(displayPrice(line.priceCents * line.quantity, line.currency))
                      : null;
                  return (
                    <Stack key={line.reservationId} direction="row" justifyContent="space-between" spacing={2}>
                      <Typography variant="body2" sx={{ minWidth: 0 }}>
                        {line.quantity}× {line.title}
                      </Typography>
                      <Typography variant="body2" sx={{ fontWeight: 600, whiteSpace: 'nowrap' }}>
                        {lineTotal ?? '—'}
                      </Typography>
                    </Stack>
                  );
                })}
              </Stack>
              <Divider sx={{ my: 1.5 }} />
              <Stack direction="row" justifyContent="space-between">
                <Typography sx={{ fontWeight: 700 }}>Subtotal</Typography>
                <Typography sx={{ fontWeight: 700 }}>{subtotalLabel ?? '—'}</Typography>
              </Stack>
            </Box>

            <Box>
              <Typography variant="overline" color="text.secondary">
                Payment method
              </Typography>
              {enabledMethods.length === 0 ? (
                <Alert severity="info" sx={{ mt: 1 }}>
                  The parents haven't set up a way to receive payments yet. Your gifts are still
                  reserved — check back soon.
                </Alert>
              ) : (
                <FormControl sx={{ mt: 1, width: '100%' }}>
                  <RadioGroup
                    value={selectedMethodId}
                    onChange={(e) => setSelectedMethodId(e.target.value)}
                  >
                    {enabledMethods.map((m) => (
                      <FormControlLabel
                        key={m.id}
                        value={m.id}
                        control={<Radio />}
                        label={m.displayName?.trim() || PAYMENT_METHOD_LABELS[m.type]}
                      />
                    ))}
                  </RadioGroup>
                </FormControl>
              )}
            </Box>

            <Alert severity="info" icon={false}>
              You are sending this payment directly to the parents. Stork Nest does not process,
              hold, or verify the payment. The parents will confirm once they receive it.
            </Alert>

            {payError && <Alert severity="error">{payError}</Alert>}

            <Typography variant="caption" color="text.secondary">
              Confirming as <strong>{buyerEmail}</strong>. The parents will see this email so they can follow up.
            </Typography>
          </Stack>
        </Box>
        <Box sx={{ px: 2.5, py: 2, borderTop: 1, borderColor: 'divider' }}>
          <Stack direction="row" justifyContent="space-between" alignItems="baseline" sx={{ mb: 1.5 }}>
            <Typography sx={{ fontWeight: 600 }}>Total to send</Typography>
            <Typography sx={{ fontWeight: 700 }}>{subtotalLabel ?? '—'}</Typography>
          </Stack>
          <Button
            fullWidth
            variant="contained"
            size="large"
            onClick={handleProceedToPayment}
            disabled={creatingIntent || enabledMethods.length === 0}
            sx={{ color: '#fff' }}
          >
            {creatingIntent ? 'Preparing…' : 'Continue to payment'}
          </Button>
        </Box>
      </>
    );
  } else {
    // pay
    const refCode = intent?.referenceCode ?? '';
    const methodLabel = selectedMethod?.displayName?.trim()
      || (selectedMethod ? PAYMENT_METHOD_LABELS[selectedMethod.type] : 'the chosen method');
    body = (
      <>
        <Box sx={{ flex: 1, minHeight: 0, overflowY: 'auto', px: 2.5, py: 2 }}>
          <Stack spacing={2.5}>
            <Stack direction="row" justifyContent="space-between" alignItems="baseline">
              <Typography sx={{ fontWeight: 700 }}>Amount to send</Typography>
              <Box sx={{ textAlign: 'right' }}>
                <Typography sx={{ fontWeight: 700 }}>
                  {intent ? formatPriceCents(intent.amountCents, intent.currency) : subtotalLabel ?? '—'}
                </Typography>
                {intent && (() => {
                  const t = (viewerCurrency || '').toUpperCase();
                  const from = (intent.currency || 'USD').toUpperCase();
                  if (!t || t === from) return null;
                  const converted = convertCents(intent.amountCents, from, t, rates);
                  if (converted == null) return null;
                  const est = formatPriceCents(converted, t);
                  return est ? (
                    <Typography variant="caption" color="text.secondary">
                      ≈ {est} for you
                    </Typography>
                  ) : null;
                })()}
              </Box>
            </Stack>
            {intent && (viewerCurrency || '').toUpperCase() !== (intent.currency || 'USD').toUpperCase() && (
              <Typography variant="caption" color="text.secondary">
                Please send the amount in {(intent.currency || 'USD').toUpperCase()} shown above — that's what the parents will receive. The {(viewerCurrency || '').toUpperCase()} figure is an approximate estimate.
              </Typography>
            )}

            <Box>
              <Typography variant="overline" color="text.secondary">
                Pay with {methodLabel}
              </Typography>
              <Stack spacing={0.75} sx={{ mt: 1, p: 2, borderRadius: 2, bgcolor: 'action.hover' }}>
                {selectedMethod?.instructions && (
                  <Typography variant="body2" sx={{ whiteSpace: 'pre-wrap' }}>
                    {selectedMethod.instructions}
                  </Typography>
                )}
                {selectedMethod?.recipientEmail && (
                  <Typography variant="body2"><strong>Email:</strong> {selectedMethod.recipientEmail}</Typography>
                )}
                {selectedMethod?.recipientPhone && (
                  <Typography variant="body2"><strong>Phone:</strong> {selectedMethod.recipientPhone}</Typography>
                )}
                {selectedMethod?.bankName && (
                  <Typography variant="body2"><strong>Bank:</strong> {selectedMethod.bankName}</Typography>
                )}
                {selectedMethod?.bankAccountName && (
                  <Typography variant="body2"><strong>Account name:</strong> {selectedMethod.bankAccountName}</Typography>
                )}
                {selectedMethod?.bankAccountNumber && (
                  <Typography variant="body2"><strong>Account number:</strong> {selectedMethod.bankAccountNumber}</Typography>
                )}
                {selectedMethod?.bankRoutingNumber && (
                  <Typography variant="body2"><strong>Routing / sort code:</strong> {selectedMethod.bankRoutingNumber}</Typography>
                )}
                {selectedMethod?.bankIban && (
                  <Typography variant="body2"><strong>IBAN:</strong> {selectedMethod.bankIban}</Typography>
                )}
                {selectedMethod?.bankSwift && (
                  <Typography variant="body2"><strong>SWIFT / BIC:</strong> {selectedMethod.bankSwift}</Typography>
                )}
                {selectedMethod?.paymentUrl && (
                  <Button
                    variant="outlined"
                    size="small"
                    component={MuiLink}
                    href={selectedMethod.paymentUrl}
                    target="_blank"
                    rel="noreferrer"
                    startIcon={<OpenInNewIcon sx={{ fontSize: 16 }} />}
                    sx={{ alignSelf: 'flex-start', mt: 0.5 }}
                  >
                    Open payment link
                  </Button>
                )}
              </Stack>
            </Box>

            <Box>
              <Typography variant="overline" color="text.secondary">
                Reference code
              </Typography>
              <Box
                sx={{
                  mt: 1,
                  p: 2,
                  borderRadius: 2,
                  border: '2px dashed',
                  borderColor: 'primary.main',
                  textAlign: 'center',
                }}
              >
                <Typography variant="h5" sx={{ fontWeight: 800, letterSpacing: 1 }}>
                  {refCode}
                </Typography>
              </Box>
              <Typography variant="caption" color="text.secondary" sx={{ display: 'block', mt: 1 }}>
                Please include this code in your payment note so the parents can match it to you.
              </Typography>
            </Box>

            <Alert severity="info" icon={false}>
              You are sending this payment directly to the parents. Stork Nest does not process,
              hold, or verify the payment. The parents will confirm once they receive it.
            </Alert>

            {payError && <Alert severity="error">{payError}</Alert>}
          </Stack>
        </Box>
        <Box sx={{ px: 2.5, py: 2, borderTop: 1, borderColor: 'divider' }}>
          <Button
            fullWidth
            variant="contained"
            size="large"
            onClick={handleClaim}
            disabled={claiming}
            sx={{ color: '#fff' }}
          >
            {claiming ? 'Confirming…' : "I've sent the payment"}
          </Button>
        </Box>
      </>
    );
  }

  return (
    <Drawer
      anchor="right"
      open={open}
      onClose={onClose}
      PaperProps={{
        sx: {
          width: isMobile ? '100%' : 420,
          maxWidth: '100vw',
          display: 'flex',
          flexDirection: 'column',
        },
      }}
    >
      {step === 'done'
        ? header('Payment sent', undefined)
        : step === 'pay'
        ? header('Send payment', 'checkout', handlePayBack)
        : step === 'checkout'
        ? header('Checkout', 'cart')
        : header(`Your cart${itemCount ? ` · ${itemCount}` : ''}`, undefined)}
      {body}
    </Drawer>
  );
}
