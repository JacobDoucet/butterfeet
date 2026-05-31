import { Box } from '@mui/material';

type BrandLogoVariant = 'text' | 'lockup';

interface BrandLogoProps {
  variant?: BrandLogoVariant;
  height?: number;
  width?: number | string;
  alt?: string;
  markScale?: number;
  /** Wordmark height as a multiple of `height` when variant="lockup". */
  wordmarkScale?: number;
  /** CSS color (defaults to `primary.main`). */
  color?: string;
}

const MARK_SRC = '/brand/stork-nest-mark-512.png';
const WORDMARK_SRC = '/brand/stork-nest-wordmark-900.png';

function maskedImage(src: string, height: number | string, color: string) {
  return {
    display: 'block',
    height,
    width: 'auto',
    bgcolor: color,
    WebkitMaskImage: `url(${src})`,
    maskImage: `url(${src})`,
    WebkitMaskRepeat: 'no-repeat',
    maskRepeat: 'no-repeat',
    WebkitMaskSize: 'contain',
    maskSize: 'contain',
    WebkitMaskPosition: 'left center',
    maskPosition: 'left center',
  } as const;
}

// To preserve aspect ratio with CSS masks, we still need an underlying <img>
// to size the box. Render an invisible img and overlay the mask.
function MaskedAsset({ src, height, color, ariaHidden, alt }: { src: string; height: number | string; color: string; ariaHidden?: boolean; alt?: string }) {
  return (
    <Box sx={{ position: 'relative', display: 'inline-block', height, flexShrink: 0 }}>
      <Box
        component="img"
        src={src}
        alt={alt ?? ''}
        aria-hidden={ariaHidden}
        sx={{ display: 'block', height, width: 'auto', visibility: 'hidden' }}
      />
      <Box
        aria-hidden
        sx={{
          position: 'absolute',
          inset: 0,
          bgcolor: color,
          WebkitMaskImage: `url(${src})`,
          maskImage: `url(${src})`,
          WebkitMaskRepeat: 'no-repeat',
          maskRepeat: 'no-repeat',
          WebkitMaskSize: 'contain',
          maskSize: 'contain',
          WebkitMaskPosition: 'center',
          maskPosition: 'center',
        }}
      />
    </Box>
  );
}

export default function BrandLogo({
  variant = 'lockup',
  height = 40,
  alt = 'Stork Nest',
  markScale = 1,
  wordmarkScale = 1.4,
  color = 'primary.main',
}: BrandLogoProps) {
  if (variant === 'text') {
    return <MaskedAsset src={WORDMARK_SRC} height={height} color={color} alt={alt} />;
  }

  return (
    <Box
      role="img"
      aria-label={alt}
      sx={{
        display: 'inline-flex',
        alignItems: 'center',
        gap: 1,
        maxWidth: '100%',
        overflow: 'hidden',
      }}
    >
      <MaskedAsset src={MARK_SRC} height={height * markScale} color={color} ariaHidden />
      <MaskedAsset src={WORDMARK_SRC} height={height * wordmarkScale} color={color} ariaHidden />
    </Box>
  );
}

// Silence unused warning for helper retained for reference.
void maskedImage;
