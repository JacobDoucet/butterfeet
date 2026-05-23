# AGENTS

Shared instructions for future LLM coding sessions in this repository.

## Purpose

This is an experiment-focused monorepo for Butter Feet Labs.

## Current Android Workflow

- Prefer `Taskfile.yml` for local Android workflows.
- App metadata is registered in `tooling/android-apps.tsv`.
- For a registered app key, prefer:
  - `task android:app:build-install-run APP=<app-key>`
  - `task android:app:cold-start APP=<app-key> AVD=<name>`
- Bad Libs is already registered as `bad-libs`.

## Android Environment Notes

- If Android builds fail with SDK-path conflicts, align `ANDROID_HOME` and `ANDROID_SDK_ROOT` to the same SDK path.
- The Taskfile resolves SDK path from `ANDROID_HOME` first, then `ANDROID_SDK_ROOT`.

## Product and Architecture Defaults

- Keep architecture intentionally simple:
  - No backend/auth/analytics by default.
  - No DI framework unless needed.
  - No database for current Bad Libs MVP.
- Content source lives at `apps/bad-libs/content/packs/starter.json` and is bundled to Android assets.

## Adding New Android Apps

1. Add app metadata to `tooling/android-apps.tsv` in this format:
   - `app_key|android_dir|app_module|app_package|app_activity`
2. Use the app-key tasks rather than adding new per-app task blocks.
3. Keep wrappers only when they improve daily ergonomics.

## Doc Hygiene

- Keep root README command examples aligned with `Taskfile.yml`.
- Prefer boring, reliable defaults over clever abstractions.
