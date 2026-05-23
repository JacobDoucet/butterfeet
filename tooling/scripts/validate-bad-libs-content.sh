#!/usr/bin/env bash
# validate-bad-libs-content.sh
# Validates all story pack JSON files under apps/bad-libs/content/packs.
#
# Fails if:
#   - JSON is invalid
#   - a pack is missing required fields (id, title, description, rating, status, emoji, tags)
#   - a pack or story has more than 3 tags
#   - a coming_soon pack has any playable stories
#   - a story has more than 15 prompts
#   - a story contains a rating field
#   - duplicate story IDs exist within a pack
#   - duplicate prompt keys exist within a story
#
# Requires: jq

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
PACKS_DIR="$ROOT_DIR/apps/bad-libs/content/packs"
ERRORS=0

if ! command -v jq &>/dev/null; then
  echo "ERROR: jq is required. Install with: brew install jq" >&2
  exit 1
fi

for FILE in "$PACKS_DIR"/*.json; do
  PACK="${FILE##*/}"

  # 1. Valid JSON
  if ! jq empty "$FILE" 2>/dev/null; then
    echo "FAIL [$PACK] Invalid JSON" >&2
    ERRORS=$((ERRORS + 1))
    continue
  fi

  # 2. Required pack fields
  for FIELD in id title description rating status emoji tags; do
    VALUE="$(jq -r --arg f "$FIELD" '.[$f] // empty' "$FILE")"
    if [[ -z "$VALUE" ]]; then
      echo "FAIL [$PACK] Missing required pack field: $FIELD" >&2
      ERRORS=$((ERRORS + 1))
    fi
  done

  STATUS="$(jq -r '.status' "$FILE")"

  # 2b. Pack tag cap (≤3)
  PACK_TAG_COUNT="$(jq '.tags | length' "$FILE")"
  if [[ "$PACK_TAG_COUNT" -gt 3 ]]; then
    echo "FAIL [$PACK] Pack has $PACK_TAG_COUNT tags (max 3)" >&2
    ERRORS=$((ERRORS + 1))
  fi

  # 3. coming_soon packs must have no stories
  if [[ "$STATUS" == "coming_soon" ]]; then
    STORY_COUNT="$(jq '.stories | length' "$FILE")"
    if [[ "$STORY_COUNT" -gt 0 ]]; then
      echo "FAIL [$PACK] coming_soon pack must have empty stories array (found $STORY_COUNT)" >&2
      ERRORS=$((ERRORS + 1))
    fi
    continue
  fi

  # 4. Per-story validation (available packs only)
  STORY_COUNT="$(jq '.stories | length' "$FILE")"
  for ((i = 0; i < STORY_COUNT; i++)); do

    STORY_ID="$(jq -r --argjson i "$i" '.stories[$i].id' "$FILE")"

    # 4a. Story has a rating field (forbidden)
    HAS_RATING="$(jq -r --argjson i "$i" '.stories[$i] | has("rating")' "$FILE")"
    if [[ "$HAS_RATING" == "true" ]]; then
      echo "FAIL [$PACK] Story '$STORY_ID' must not have a rating field (rating lives on the pack)" >&2
      ERRORS=$((ERRORS + 1))
    fi

    # 4b. Too many prompts
    PROMPT_COUNT="$(jq --argjson i "$i" '.stories[$i].prompts | length' "$FILE")"
    if [[ "$PROMPT_COUNT" -gt 15 ]]; then
      echo "FAIL [$PACK] Story '$STORY_ID' has $PROMPT_COUNT prompts (max 15)" >&2
      ERRORS=$((ERRORS + 1))
    fi

    # 4c. Story tag cap (≤3)
    STORY_TAG_COUNT="$(jq --argjson i "$i" '.stories[$i].tags | length' "$FILE")"
    if [[ "$STORY_TAG_COUNT" -gt 3 ]]; then
      echo "FAIL [$PACK] Story '$STORY_ID' has $STORY_TAG_COUNT tags (max 3)" >&2
      ERRORS=$((ERRORS + 1))
    fi

    # 4d. Duplicate prompt keys within a story
    DUP_KEYS="$(jq -r --argjson i "$i" '[.stories[$i].prompts[].key] | group_by(.) | map(select(length > 1)) | .[] | .[0]' "$FILE")"
    if [[ -n "$DUP_KEYS" ]]; then
      echo "FAIL [$PACK] Story '$STORY_ID' has duplicate prompt keys: $DUP_KEYS" >&2
      ERRORS=$((ERRORS + 1))
    fi
  done

  # 5. Duplicate story IDs within a pack
  DUP_STORY_IDS="$(jq -r '[.stories[].id] | group_by(.) | map(select(length > 1)) | .[] | .[0]' "$FILE")"
  if [[ -n "$DUP_STORY_IDS" ]]; then
    echo "FAIL [$PACK] Duplicate story IDs: $DUP_STORY_IDS" >&2
    ERRORS=$((ERRORS + 1))
  fi

  echo "OK   [$PACK] $STORY_COUNT stories"
done

echo ""
if [[ "$ERRORS" -gt 0 ]]; then
  echo "Validation failed with $ERRORS error(s)." >&2
  exit 1
else
  echo "All packs valid."
fi
