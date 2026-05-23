# Repository LLM Context

This file captures practical context for future model sessions.

## Repo Snapshot

- Monorepo root with apps, packages, tooling, and docs.
- First app: Bad Libs Android app at `apps/bad-libs/android`.

## Preferred Dev Commands

- `task --list`
- `task android:doctor`
- `task android:apps:list`
- `task android:app:build-install-run APP=bad-libs`
- `task android:app:cold-start APP=bad-libs AVD=<name>`

## Android App Registry

- Registry file: `tooling/android-apps.tsv`
- Format: `app_key|android_dir|app_module|app_package|app_activity`

## Known Local Gotcha

- Mismatched `ANDROID_HOME` and `ANDROID_SDK_ROOT` can break Gradle SDK resolution.
- Align them to the same path when needed.

## Bad Libs Product Constraints

- Local-only JSON content.
- No backend/auth/ads/billing/analytics in MVP.
- Keep UI playful but implementation simple.
