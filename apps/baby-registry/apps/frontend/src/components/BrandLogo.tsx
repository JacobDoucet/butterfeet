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
}

const MARK_SRC = '/brand/stork-nest-mark-512.png';
const WORDMARK_SRC = '/brand/stork-nest-wordmark-900.png';

export default function BrandLogo({
  variant = 'lockup',
  height = 40,
  width,
  alt = 'Stork Nest',
  markScale = 1,
  wordmarkScale = 1.4,
}: BrandLogoProps) {
  if (variant === 'text') {
    return (
      <Box
        component="img"
        src={WORDMARK_SRC}
        alt={alt}
        sx={{
          display: 'block',
          height,
          width: width ?? 'auto',
          maxWidth: '100%',
        }}
      />
    );
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
      <Box
        component="img"
        src={MARK_SRC}
        alt=""
        aria-hidden="true"
        sx={{
          display: 'block',
          height: height * markScale,
          maxHeight: '13vw',
          width: 'auto',
          flexShrink: 0,
        }}
      />
      <Box
        component="img"
        src={WORDMARK_SRC}
        alt=""
        aria-hidden="true"
        sx={{
          display: 'block',
          height: height * wordmarkScale,
          maxHeight: '16vw',
          width: 'auto',
          minWidth: 0,
          flexShrink: 1,
        }}
      />
    </Box>
  );
}
