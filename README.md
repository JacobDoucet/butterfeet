# Butter Feet Labs

Butter Feet Labs is an experimental app studio for shipping small, weird, fun software ideas quickly.

## Monorepo Shape

- Apps live under `apps/<app-name>`.
- Shared code belongs in `packages/` **only after** it has been reused by more than one app.
- Tooling, scripts, and templates live in `tooling/`.

## First App: Bad Stories

The first app is **Bad Stories**, an Android word-game app inspired by fill-in-the-blank party games.

## CLI Workflow (Taskfile)

This repo includes a root `Taskfile.yml` to streamline emulator and app commands.

LLM session context is captured in:

- `AGENTS.md`
- `.llms/context.md`

Install Task (if needed):

```bash
brew install go-task/tap/go-task
```

Common commands:

```bash
# Check SDK paths and tool availability
task android:doctor

# List emulators
task android:list-avds

# Start emulator (uses first AVD if AVD is not provided)
task android:emu
task android:emu AVD=Pixel_8

# Start emulator in background and wait for full boot
task android:emu-bg
task android:wait-for-boot

# Build/install/run any Android app
task android:assemble ANDROID_DIR=apps/<app-name>/android APP_MODULE=app
task android:install ANDROID_DIR=apps/<app-name>/android APP_MODULE=app
task android:run APP_PACKAGE=com.example.app APP_ACTIVITY=.MainActivity
task android:build-install-run ANDROID_DIR=apps/<app-name>/android APP_MODULE=app APP_PACKAGE=com.example.app APP_ACTIVITY=.MainActivity
task android:cold-start ANDROID_DIR=apps/<app-name>/android APP_MODULE=app APP_PACKAGE=com.example.app APP_ACTIVITY=.MainActivity AVD=Pixel_8

# Registry-based shortcuts (recommended for multiple apps)
task android:apps:list
task android:app:build-install-run APP=bad-stories
task android:app:cold-start APP=bad-stories AVD=Pixel_8

# Bad Stories shortcuts
task bad-stories:assemble
task bad-stories:install
task bad-stories:run

# One-shot build + install + run
task bad-stories:build-install-run

# One-shot cold start: boot emulator + build + install + run
task bad-stories:cold-start
```

Note: if both `ANDROID_HOME` and `ANDROID_SDK_ROOT` are set differently on your machine, align them to the same SDK path.

### Add A New Android App To Taskfile Registry

Add one line to `tooling/android-apps.tsv` in this format:

```text
app_key|android_dir|app_module|app_package|app_activity
```

Example:

```text
word-bonks|apps/word-bonks/android|app|com.butterfeetlabs.wordbonks|.MainActivity
```

Then run:

```bash
task android:app:build-install-run APP=word-bonks
task android:app:cold-start APP=word-bonks AVD=Pixel_8
```

## Repository Layout

```text
butter-feet-labs/
  apps/
    bad-stories/
      android/
      content/
      docs/
      README.md
  packages/
    android/
  tooling/
    scripts/
    templates/
  docs/
    ideas.md
    app-launch-checklist.md
  README.md
```
