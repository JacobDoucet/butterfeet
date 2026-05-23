#!/usr/bin/env bash
# sync-bad-libs-assets.sh
# Syncs Bad Libs content packs from the content source into Android assets.
#
# Run this after editing any pack JSON under apps/bad-libs/content/packs/.
# Mirrors the content directory into the Android bundle directory exactly,
# removing any stale pack files that no longer exist in content.
#
# Usage:
#   ./tooling/scripts/sync-bad-libs-assets.sh

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
SRC="$ROOT_DIR/apps/bad-libs/content/packs"
DEST="$ROOT_DIR/apps/bad-libs/android/app/src/main/assets/packs"

rsync -a --delete "$SRC/" "$DEST/"

echo "Synced content packs -> Android assets"
echo "Files now in assets:"
ls "$DEST/"
