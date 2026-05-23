#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
CONTENT_FILE="$ROOT_DIR/apps/bad-libs/content/packs/starter.json"

if [[ ! -f "$CONTENT_FILE" ]]; then
  echo "Missing content file: $CONTENT_FILE" >&2
  exit 1
fi

if ! python3 -m json.tool "$CONTENT_FILE" >/dev/null 2>&1; then
  echo "Invalid JSON in: $CONTENT_FILE" >&2
  exit 1
fi

echo "Content looks good: $CONTENT_FILE"
