# Bad Libs Product Notes

## Launch Content Strategy (V1)

- Lead with personality-rich content, not a tutorial/demo vibe.
- Starter Pack is removed from product and content sources.
- App should feel immediately deep with multiple distinct tones.
- Core flow stays local-first: Home -> Packs -> Stories -> Prompts -> Reveal -> Share.

## Pack File Naming Convention

Pack filenames are **permanent and stable**. A pack's availability is controlled exclusively by its `status` metadata field, never by the filename.

- `"status": "available"` — pack is playable and fully shown in the UI.
- `"status": "coming_soon"` — pack is visible in browse but locked and not playable.

Examples of correct filenames:

- `office.json` (not `office-coming-soon.json`)
- `wedding.json` (not `wedding-coming-soon.json`)

When a coming-soon pack is ready to ship, change `"status"` to `"available"` and populate its `stories` array. No file renames required.

## Pack Size Target

- **15–20 stories** per major available pack.
- Distribution goal: **40% short (3–5 prompts), 40% medium (6–9 prompts), 20% long (10–15 prompts)**.

## Pack Lineup

### Available at launch

- 🎉 Party Pack
- 🧠 Internet Brainrot
- ❤️ Couples Pack
- 🦖 Kids Pack

### Visible but coming soon

- 💼 Office Pack
- 📺 Reality TV Pack
- 🚀 Science Fiction Pack
- 🕵️ Crime Documentary Pack
- ☢️ Gen Alpha Pack

Coming-soon packs are intentionally visible in browse UI to build anticipation, but are not playable yet.

## Content Schema Notes

- `StoryPack` owns rating and metadata:
  - `id`, `title`, `description`, `rating`, `status`, `emoji`, `tags` (max 3), optional `accentName`, `stories`
- `Story` owns content-level metadata:
  - `id`, `title`, `tags` (max 3), `prompts`, `template`
- No story-level rating.
- No dedicated story length field.
- Length UX is inferred from prompt count:
  - 3-5 prompts -> ⚡ Quick
  - 6-9 prompts -> 🍔 Medium
  - 10+ prompts -> 🧠 Brain Damage

## Product UX Rules

- Pack browsing should feel funny, branded, and a little chaotic.
- Available packs are vibrant and tappable.
- Coming-soon packs are visible, stylized, and non-playable.
- Tapping a coming-soon pack should give friendly feedback ("This pack is still being cooked.").

## Roadmap (Content First)

- Continue expanding stories and themed packs before monetization.
- Improve per-pack voice and comedic identity.
- Add more seasonal/event pack drops.

## Monetization Status

- Monetization is intentionally deferred.
- No ads, billing, subscriptions, premium unlocks, backend, accounts, or analytics in current launch scope.
