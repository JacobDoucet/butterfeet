#!/usr/bin/env python3
"""Rebuild Stork Nest brand assets from sources in apps/baby-registry/temp.

- Recolors the new terracotta/copper sources to sage green (#7a9e7e).
- Makes backgrounds transparent.
- Emits the size variants the frontend expects.

Run: python3 apps/baby-registry/scripts/rebuild-brand.py
"""
from __future__ import annotations

from pathlib import Path

import numpy as np
from PIL import Image

SAGE = (122, 158, 126)  # #7a9e7e
ROOT = Path(__file__).resolve().parents[1]
SRC_ICON = ROOT / "temp" / "icon.png"
SRC_WORDMARK = ROOT / "temp" / "stork-list-wordmark.png"
OUT_DIR = ROOT / "apps" / "frontend" / "public" / "brand"


def load_rgba(path: Path) -> np.ndarray:
    return np.array(Image.open(path).convert("RGBA"), dtype=np.uint8)


def save(img: Image.Image, name: str) -> None:
    out = OUT_DIR / name
    img.save(out, optimize=True)
    print(f"wrote {out.relative_to(ROOT)}  ({img.size[0]}x{img.size[1]})")


def recolor_from_bg(arr: np.ndarray, bg_rgb: tuple[int, int, int], fg_rgb: tuple[int, int, int]) -> Image.Image:
    """Replace ``fg`` pixels with sage, treating ``bg`` as transparent.

    Per-pixel alpha = projection of (pixel - bg) onto (fg - bg) direction, clamped.
    """
    rgb = arr[..., :3].astype(np.float32)
    src_alpha = arr[..., 3].astype(np.float32) / 255.0

    bg = np.array(bg_rgb, dtype=np.float32)
    fg = np.array(fg_rgb, dtype=np.float32)
    direction = fg - bg
    denom = float(np.dot(direction, direction)) or 1.0

    diff = rgb - bg
    t = (diff @ direction) / denom
    t = np.clip(t, 0.0, 1.0)

    alpha = (t * src_alpha * 255.0).round().astype(np.uint8)
    out = np.zeros_like(arr)
    out[..., 0] = SAGE[0]
    out[..., 1] = SAGE[1]
    out[..., 2] = SAGE[2]
    out[..., 3] = alpha
    return Image.fromarray(out, mode="RGBA")


def trim(img: Image.Image, padding: int = 0) -> Image.Image:
    """Crop to the non-transparent bounding box, with optional padding."""
    alpha = np.array(img)[..., 3]
    ys, xs = np.where(alpha > 4)
    if not len(xs):
        return img
    x0, x1 = xs.min(), xs.max() + 1
    y0, y1 = ys.min(), ys.max() + 1
    x0 = max(0, x0 - padding)
    y0 = max(0, y0 - padding)
    x1 = min(img.size[0], x1 + padding)
    y1 = min(img.size[1], y1 + padding)
    return img.crop((x0, y0, x1, y1))


def fit_square(img: Image.Image, size: int) -> Image.Image:
    """Letterbox into a transparent square of ``size`` px."""
    w, h = img.size
    scale = min(size / w, size / h)
    new_w, new_h = max(1, int(round(w * scale))), max(1, int(round(h * scale)))
    resized = img.resize((new_w, new_h), Image.LANCZOS)
    canvas = Image.new("RGBA", (size, size), (0, 0, 0, 0))
    canvas.paste(resized, ((size - new_w) // 2, (size - new_h) // 2), resized)
    return canvas


def fit_height(img: Image.Image, height: int) -> Image.Image:
    w, h = img.size
    scale = height / h
    return img.resize((max(1, int(round(w * scale))), height), Image.LANCZOS)


def make_lockup(mark: Image.Image, word: Image.Image, height: int) -> Image.Image:
    """Combine mark + wordmark horizontally with a small gap, vertically centered."""
    mark_h = int(round(height * 0.95))
    word_h = int(round(height * 0.55))
    gap = int(round(height * 0.18))

    m = fit_height(mark, mark_h)
    w = fit_height(word, word_h)
    total_w = m.size[0] + gap + w.size[0]
    canvas = Image.new("RGBA", (total_w, height), (0, 0, 0, 0))
    canvas.paste(m, (0, (height - m.size[1]) // 2), m)
    canvas.paste(w, (m.size[0] + gap, (height - w.size[1]) // 2), w)
    return canvas


def sample_corner_bg(arr: np.ndarray) -> tuple[int, int, int]:
    """Sample the inner-background color near the top edge, away from the outer corner."""
    h, w = arr.shape[:2]
    sample = arr[h // 2, w // 8, :3]
    return tuple(int(c) for c in sample)


def sample_dominant_fg(arr: np.ndarray, bg: tuple[int, int, int]) -> tuple[int, int, int]:
    """Pick the pixel farthest from bg as a representative foreground sample."""
    rgb = arr[..., :3].astype(np.float32)
    bg_arr = np.array(bg, dtype=np.float32)
    dist = np.linalg.norm(rgb - bg_arr, axis=-1) * (arr[..., 3] / 255.0)
    idx = np.unravel_index(int(np.argmax(dist)), dist.shape)
    return tuple(int(c) for c in arr[idx][:3])


def main() -> None:
    OUT_DIR.mkdir(parents=True, exist_ok=True)

    # --- ICON / MARK ---------------------------------------------------------
    icon_arr = load_rgba(SRC_ICON)
    # The icon has rounded corners with a white outside and a peach inside.
    # Force outside white (and near-white) to transparent first.
    rgb = icon_arr[..., :3].astype(np.int32)
    white_mask = (rgb[..., 0] > 240) & (rgb[..., 1] > 240) & (rgb[..., 2] > 240)
    icon_arr[white_mask, 3] = 0

    bg = sample_corner_bg(icon_arr)
    fg = sample_dominant_fg(icon_arr, bg)
    print(f"icon: bg={bg}  fg={fg}  -> sage={SAGE}")
    mark_full = recolor_from_bg(icon_arr, bg, fg)
    mark_full = trim(mark_full, padding=8)
    mark_square = fit_square(mark_full, 1024)

    for size, name in [
        (1024, "stork-nest-mark-1024.png"),
        (512, "stork-nest-mark-512.png"),
        (192, "stork-nest-mark-192.png"),
        (64, "stork-nest-mark-64.png"),
    ]:
        save(mark_square.resize((size, size), Image.LANCZOS), name)

    # --- WORDMARK ------------------------------------------------------------
    word_arr = load_rgba(SRC_WORDMARK)
    # White background, copper text.
    rgb = word_arr[..., :3].astype(np.float32)
    word_bg = (255, 255, 255)
    word_fg = sample_dominant_fg(word_arr, word_bg)
    print(f"wordmark: fg={word_fg}  -> sage={SAGE}")
    word_recolored = recolor_from_bg(word_arr, word_bg, word_fg)
    word_recolored = trim(word_recolored, padding=4)

    base_h = 900
    word_900 = fit_height(word_recolored, base_h)
    save(word_900, "stork-nest-wordmark-900.png")
    save(fit_height(word_recolored, 600), "stork-nest-wordmark-600.png")
    save(fit_height(word_recolored, 200), "stork-nest-wordmark.png")

    # --- LOCKUP --------------------------------------------------------------
    for h, name in [
        (1200, "stork-nest-logo-lockup-1200.png"),
        (900, "stork-nest-logo-lockup-900.png"),
        (200, "stork-nest-logo-lockup.png"),
    ]:
        save(make_lockup(mark_full, word_recolored, h), name)

    # --- FAVICONS ------------------------------------------------------------
    public_dir = OUT_DIR.parent
    fav_src = mark_square
    fav32 = fav_src.resize((32, 32), Image.LANCZOS)
    fav16 = fav_src.resize((16, 16), Image.LANCZOS)
    apple = fav_src.resize((180, 180), Image.LANCZOS)
    fav32.save(public_dir / "favicon-32x32.png", optimize=True)
    fav16.save(public_dir / "favicon-16x16.png", optimize=True)
    apple.save(public_dir / "apple-touch-icon.png", optimize=True)
    ico = fav_src.resize((48, 48), Image.LANCZOS)
    ico.save(
        public_dir / "favicon.ico",
        sizes=[(16, 16), (24, 24), (32, 32), (48, 48)],
    )
    print(f"wrote favicons in {public_dir.relative_to(ROOT)}")


if __name__ == "__main__":
    main()
