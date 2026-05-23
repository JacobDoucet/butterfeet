# .llms

This directory contains optional, model-facing context documents.

Notes:

- `.llms` is a useful team convention, but not a universal standard across all tools.
- `AGENTS.md` at repo root is also provided because many agent frameworks look for top-level guidance files.
- Keep both files short and maintained when workflows change.

## Update Checklist

When changing Android workflow, keep these files in sync:

1. If Taskfile commands or variable names change, update `README.md` command examples.
2. If app registry behavior changes, update `tooling/android-apps.tsv` docs in both `README.md` and `.llms/context.md`.
3. If preferred commands change, update `AGENTS.md` and `.llms/context.md`.
4. If environment gotchas change, update the Android environment notes in `AGENTS.md`.
