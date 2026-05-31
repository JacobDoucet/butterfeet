import { Box, Card, CardContent, CardMedia, Stack, Typography } from '@mui/material';
import type { ReactNode } from 'react';

interface ItemCardProps {
  imageUrl?: string;
  imageBgColor?: string;
  title: string;
  onClick?: () => void;
  disabled?: boolean;
  imageFilter?: string;
  topLeftOverlay?: ReactNode;
  topRightOverlay?: ReactNode;
  belowTitle?: ReactNode;
  footer?: ReactNode;
  dimTitle?: boolean;
}

export default function ItemCard({
  imageUrl,
  imageBgColor,
  title,
  onClick,
  disabled,
  imageFilter,
  topLeftOverlay,
  topRightOverlay,
  belowTitle,
  footer,
  dimTitle,
}: ItemCardProps) {
  const interactive = !!onClick && !disabled;
  return (
    <Card
      elevation={0}
      sx={{
        height: '100%',
        width: '100%',
        display: 'flex',
        flexDirection: 'column',
        cursor: interactive ? 'pointer' : 'default',
        border: '1px solid',
        borderColor: 'divider',
        borderRadius: 3,
        overflow: 'hidden',
        position: 'relative',
        bgcolor: 'background.paper',
        transition: 'transform 0.18s ease, box-shadow 0.18s ease, border-color 0.18s ease',
        ...(interactive
          ? {
              '&:hover': {
                transform: 'translateY(-3px)',
                boxShadow: '0 12px 28px rgba(0,0,0,0.08)',
                borderColor: 'primary.light',
              },
              '&:focus-visible': {
                outline: '2px solid',
                outlineColor: 'primary.main',
                outlineOffset: 2,
              },
            }
          : {}),
      }}
      onClick={interactive ? onClick : undefined}
      role={interactive ? 'button' : undefined}
      tabIndex={interactive ? 0 : -1}
      onKeyDown={(e) => {
        if (!interactive) return;
        if (e.key === 'Enter' || e.key === ' ') {
          e.preventDefault();
          onClick?.();
        }
      }}
    >
      <Box sx={{ position: 'relative', bgcolor: imageBgColor || '#ffffff' }}>
        {imageUrl ? (
          <CardMedia
            component="img"
            image={imageUrl}
            sx={{
              aspectRatio: '1',
              objectFit: 'contain',
              p: 2,
              filter: imageFilter,
              transition: 'filter 0.2s',
            }}
          />
        ) : (
          <Box
            sx={{
              aspectRatio: '1',
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              color: 'text.disabled',
            }}
          >
            <Typography variant="caption">No image</Typography>
          </Box>
        )}
        {topLeftOverlay && (
          <Box sx={{ position: 'absolute', top: 10, left: 10 }}>{topLeftOverlay}</Box>
        )}
        {topRightOverlay && (
          <Box sx={{ position: 'absolute', top: 12, right: 12 }}>{topRightOverlay}</Box>
        )}
      </Box>
      <CardContent sx={{ flexGrow: 1, display: 'flex', flexDirection: 'column', p: 2.5 }}>
        <Typography
          variant="subtitle1"
          sx={{
            fontWeight: 600,
            lineHeight: 1.35,
            display: '-webkit-box',
            WebkitLineClamp: 2,
            WebkitBoxOrient: 'vertical',
            overflow: 'hidden',
            mb: belowTitle ? 1 : 1,
            color: dimTitle ? 'text.secondary' : 'text.primary',
          }}
        >
          {title}
        </Typography>
        {belowTitle}
        <Box sx={{ flexGrow: 1 }} />
        {footer && (
          <Stack direction="row" alignItems="center" justifyContent="space-between" spacing={1}>
            {footer}
          </Stack>
        )}
      </CardContent>
    </Card>
  );
}
