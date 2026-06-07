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
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import {
  formatPriceCents,
  convertCents,
  PAYMENT_METHOD_LABELS,
  type PaymentMethod,
  type PaymentIntent,
  type ExchangeRates,
  type MyCart,
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
  heldCart?: MyCart;
  heldCartMethod?: PaymentMethod;
  heldCartLines?: CartLine[];
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
  heldCart,
  heldCartMethod,
  heldCartLines,
}: CartDrawerProps) {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('sm'));
  const [step, setStep] = useState<Step>('cart');
  const [selectedMethodId, setSelectedMethodId] = useState('');
  const [intent, setIntent] = useState<PaymentIntent | null>(null);
  const [payError, setPayError] = useState<string | null>(null);
  const [creatingIntent, setCreatingIntent] = useState(false);
  const [claiming, setClaiming] = useState(false);
  // Snapshot of the cart lines taken when payment starts. During the pay step
  // the gifts are locked into the cart so `lines` is empty, but we still want
  // to show the guest the same cart list they were paying for.
  const [payLines, setPayLines] = useState<CartLine[]>([]);
  // When resuming a previously-held cart, the payment method may no longer be
  // in the enabled list, so we keep a direct reference. `resuming` also tells
  // the back action to leave the cart held rather than cancelling it.
  const [resumeMethod, setResumeMethod] = useState<PaymentMethod | null>(null);
  const [resuming, setResuming] = useState(false);
  const [cancellingHeld, setCancellingHeld] = useState(false);

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
      setResuming(false);
      setResumeMethod(null);
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
      setPayLines(lines);
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
  // the open cart. When resuming a held cart, backing out instead leaves the
  // cart held (the buyer cancels explicitly) and returns to the cart view.
  const handlePayBack = async () => {
    if (resuming) {
      setIntent(null);
      setResuming(false);
      setResumeMethod(null);
      setStep('cart');
      return;
    }
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

  // Resume a held cart: rebuild a synthetic intent from the stored cart so the
  // existing pay step can show the payment instructions and "I've sent it".
  const resumeHeldCart = () => {
    if (!heldCart) return;
    setPayError(null);
    setPayLines(heldCartLines ?? []);
    setResumeMethod(heldCartMethod ?? null);
    if (heldCartMethod) setSelectedMethodId(heldCartMethod.id);
    setIntent({
      ok: true,
      id: heldCart.id,
      referenceCode: heldCart.referenceCode,
      amountCents: heldCart.amountCents,
      currency: heldCart.currency,
      status: heldCart.status,
      items: heldCart.items,
    });
    setResuming(true);
    setStep('pay');
  };

  // Cancel a held cart entirely, releasing its gifts back to the registry.
  const cancelHeldCart = async () => {
    if (!heldCart || !onCancelPaymentIntent) return;
    setCancellingHeld(true);
    setPayError(null);
    try {
      await onCancelPaymentIntent(heldCart.id);
    } catch (err) {
      setPayError(err instanceof Error ? err.message : 'Could not cancel the held cart.');
    } finally {
      setCancellingHeld(false);
    }
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

  // Banner surfacing a held cart (a payment the buyer started but never
  // confirmed). Shown on the cart / empty views so the buyer can resume and
  // confirm, or cancel to release the gifts.
  const heldCartBanner =
    heldCart && step !== 'pay' && step !== 'done' ? (
      <Box
        sx={{
          m: 2.5,
          mb: 0,
          p: 2,
          borderRadius: 2,
          border: 1,
          borderColor: 'warning.light',
          bgcolor: 'warning.50',
        }}
      >
        <Typography sx={{ fontWeight: 700, mb: 0.5 }}>Payment in progress</Typography>
        <Typography variant="body2" color="text.secondary" sx={{ mb: 1.5 }}>
          You started sending {fmt(displayPrice(heldCart.amountCents, heldCart.currency)) ?? formatPriceCents(heldCart.amountCents, heldCart.currency)} for{' '}
          {heldCart.items.length} {heldCart.items.length === 1 ? 'gift' : 'gifts'} but haven't
          confirmed yet. Your gifts stay reserved until you confirm or cancel.
        </Typography>
        {payError && (
          <Alert severity="error" sx={{ mb: 1.5 }}>
            {payError}
          </Alert>
        )}
        <Stack direction="row" spacing={1}>
          <Button variant="contained" size="small" onClick={resumeHeldCart} sx={{ color: '#fff' }}>
            Confirm payment
          </Button>
          <Button
            variant="outlined"
            size="small"
            color="inherit"
            onClick={cancelHeldCart}
            disabled={cancellingHeld}
          >
            {cancellingHeld ? 'Cancelling…' : 'Cancel'}
          </Button>
        </Stack>
      </Box>
    ) : null;

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
        <Button variant="contained" onClick={onClose} sx={{ color: '#fff', mt: 1 }}>
          Back to the registry
        </Button>
      </Stack>
    );
  } else if (lines.length === 0 && step !== 'pay') {
    body = (
      <Box sx={{ flex: 1, minHeight: 0, overflowY: 'auto', display: 'flex', flexDirection: 'column' }}>
        {heldCartBanner}
        {heldCart && heldCartLines && heldCartLines.length > 0 ? (
          <Box sx={{ px: 2.5 }}>
            <Stack divider={<Divider />}>{heldCartLines.map(renderLine)}</Stack>
          </Box>
        ) : (
          <Stack spacing={2} sx={{ px: 3, py: 8, textAlign: 'center', alignItems: 'center', flex: 1, justifyContent: 'center' }}>
            <ShoppingBagOutlinedIcon sx={{ fontSize: 56, color: 'text.disabled' }} />
            <Typography variant="h6" sx={{ fontWeight: 700 }}>
              {heldCart ? 'No new gifts in your cart' : 'Your cart is empty'}
            </Typography>
            <Typography color="text.secondary" sx={{ maxWidth: 300 }}>
              Add a gift from the registry and it will be held just for you.
            </Typography>
            <Button variant="contained" onClick={onContinueShopping} sx={{ color: '#fff', mt: 1 }}>
              Browse the registry
            </Button>
          </Stack>
        )}
      </Box>
    );
  } else if (step === 'cart') {
    body = (
      <>
        <Box sx={{ flex: 1, minHeight: 0, overflowY: 'auto' }}>
          {heldCartBanner}
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
    const payMethod = selectedMethod ?? resumeMethod;
    const methodLabel = payMethod?.displayName?.trim()
      || (payMethod ? PAYMENT_METHOD_LABELS[payMethod.type] : 'the chosen method');
    body = (
      <>
        <Box sx={{ flex: 1, minHeight: 0, overflowY: 'auto', px: 2.5, py: 2 }}>
          <Stack spacing={2.5}>
            {(() => {
              const t = (viewerCurrency || '').toUpperCase();
              const from = (intent?.currency || 'USD').toUpperCase();
              const nativeText = intent
                ? formatPriceCents(intent.amountCents, intent.currency)
                : subtotalLabel ?? '—';
              let viewerText: string | null = null;
              if (intent && t && t !== from) {
                const converted = convertCents(intent.amountCents, from, t, rates);
                if (converted != null) viewerText = formatPriceCents(converted, t);
              }
              return (
                <Stack direction="row" justifyContent="space-between" alignItems="baseline" spacing={1}>
                  <Typography sx={{ fontWeight: 700, fontSize: { xs: '1.1rem', sm: '1rem' } }}>
                    Amount to send
                  </Typography>
                  <Box sx={{ textAlign: 'right' }}>
                    <Typography
                      sx={{ fontWeight: 800, fontSize: { xs: '1.6rem', sm: '1.35rem' }, lineHeight: 1.2 }}
                    >
                      {viewerText ? `≈ ${viewerText}` : nativeText}
                    </Typography>
                    {viewerText && (
                      <Typography sx={{ fontSize: { xs: '0.95rem', sm: '0.85rem' }, color: 'text.secondary' }}>
                        {nativeText} sent to the parents
                      </Typography>
                    )}
                  </Box>
                </Stack>
              );
            })()}

            <Box>
              <Typography variant="overline" sx={{ color: 'text.secondary', fontSize: { xs: '0.8rem', sm: '0.75rem' } }}>
                Pay with {methodLabel}
              </Typography>
              <Stack spacing={0.75} sx={{ mt: 1, p: 2, borderRadius: 2, bgcolor: 'action.hover' }}>
                {selectedMethod?.instructions && (
                  <Typography sx={{ whiteSpace: 'pre-wrap', fontSize: { xs: '1rem', sm: '0.9rem' } }}>
                    {selectedMethod.instructions}
                  </Typography>
                )}
                {selectedMethod?.recipientEmail && (
                  <CopyableDetail label="Email" value={selectedMethod.recipientEmail} />
                )}
                {selectedMethod?.recipientPhone && (
                  <CopyableDetail label="Phone" value={selectedMethod.recipientPhone} />
                )}
                {selectedMethod?.bankName && <PayDetail label="Bank" value={selectedMethod.bankName} />}
                {selectedMethod?.bankAccountName && (
                  <PayDetail label="Account name" value={selectedMethod.bankAccountName} />
                )}
                {selectedMethod?.bankAccountNumber && (
                  <CopyableDetail label="Account number" value={selectedMethod.bankAccountNumber} />
                )}
                {selectedMethod?.bankRoutingNumber && (
                  <CopyableDetail label="Routing / sort code" value={selectedMethod.bankRoutingNumber} />
                )}
                {selectedMethod?.bankIban && (
                  <CopyableDetail label="IBAN" value={selectedMethod.bankIban} />
                )}
                {selectedMethod?.bankSwift && (
                  <CopyableDetail label="SWIFT / BIC" value={selectedMethod.bankSwift} />
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

            <Alert severity="info" icon={false}>
              You are sending this payment directly to the parents. Stork Nest does not process,
              hold, or verify the payment. The parents will confirm once they receive it.
            </Alert>

            {payError && <Alert severity="error">{payError}</Alert>}

            {payLines.length > 0 && (
              <Box>
                <Typography variant="overline" sx={{ color: 'text.secondary', fontSize: { xs: '0.8rem', sm: '0.75rem' } }}>
                  Your cart
                </Typography>
                <Stack divider={<Divider />}>{payLines.map(renderLine)}</Stack>
              </Box>
            )}
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

// PayDetail renders a labelled payment value at a comfortable mobile size.
function PayDetail({ label, value }: { label: string; value: string }) {
  return (
    <Typography sx={{ fontSize: { xs: '1rem', sm: '0.9rem' }, wordBreak: 'break-word' }}>
      <strong>{label}:</strong> {value}
    </Typography>
  );
}

// CopyableDetail renders a payment value the guest can tap to copy (e.g. the
// recipient email for an e-transfer), with brief inline "Copied" feedback.
function CopyableDetail({ label, value }: { label: string; value: string }) {
  const [copied, setCopied] = useState(false);
  const handleCopy = async () => {
    try {
      await navigator.clipboard.writeText(value);
      setCopied(true);
      setTimeout(() => setCopied(false), 1500);
    } catch {
      // Clipboard may be unavailable (e.g. insecure context); ignore.
    }
  };
  return (
    <Box
      role="button"
      tabIndex={0}
      onClick={handleCopy}
      onKeyDown={(e) => {
        if (e.key === 'Enter' || e.key === ' ') {
          e.preventDefault();
          handleCopy();
        }
      }}
      aria-label={`Copy ${label}`}
      sx={{
        display: 'flex',
        alignItems: 'center',
        gap: 1,
        cursor: 'pointer',
        borderRadius: 1,
        px: 1,
        py: 0.75,
        mx: -1,
        '&:hover': { bgcolor: 'action.selected' },
      }}
    >
      <Box sx={{ flex: 1, minWidth: 0 }}>
        <Typography
          component="span"
          sx={{
            display: { xs: 'block', sm: 'inline' },
            fontWeight: 700,
            fontSize: { xs: '0.85rem', sm: '0.9rem' },
            color: { xs: 'text.secondary', sm: 'text.primary' },
          }}
        >
          {label}:
        </Typography>{' '}
        <Typography
          component="span"
          sx={{ fontSize: { xs: '1rem', sm: '0.9rem' }, wordBreak: 'break-all' }}
        >
          {value}
        </Typography>
      </Box>
      {copied ? (
        <Typography
          variant="caption"
          sx={{ color: 'success.main', fontWeight: 600, whiteSpace: 'nowrap' }}
        >
          Copied
        </Typography>
      ) : (
        <ContentCopyIcon sx={{ fontSize: 18, color: 'text.secondary', flexShrink: 0 }} />
      )}
    </Box>
  );
}
