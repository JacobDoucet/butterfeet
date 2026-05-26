import { Box } from '@mui/material';

type BrandLogoVariant = 'text' | 'lockup';

interface BrandLogoProps {
  variant?: BrandLogoVariant;
  height?: number;
  width?: number | string;
  alt?: string;
}

const SRC_BY_VARIANT: Record<BrandLogoVariant, string> = {
  text: '/brand/stork-nest-wordmark-900.png',
  lockup: '/brand/stork-nest-logo-lockup-900.png',
};

export default function BrandLogo({
  variant = 'lockup',
  height = 40,
  width,
  alt = 'Stork Nest',
}: BrandLogoProps) {
  return (
    <Box
      component="img"
      src={SRC_BY_VARIANT[variant]}
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
