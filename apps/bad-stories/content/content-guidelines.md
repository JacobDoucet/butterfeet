# Bad Stories Content Guidelines

## Global Story Rules

- Every story needs three beats: **setup**, **escalation**, and **payoff**.
- Avoid pure randomness. Each word/phrase substitution should feel like it could actually belong in the sentence — and be funnier for not.
- Avoid repeated sentence structures across stories in the same pack.
- Prefer socially recognizable chaos over abstract weirdness. Readers should immediately picture the scenario.
- Use custom phrase prompts (e.g. `"label": "a suspicious excuse"`) when they improve replayability over generic parts of speech.
- Keep templates readable out loud. The reveal should work as a spoken performance.

## Prompt Rules

- Max **15 prompts** per story.
- Short stories: **3–5 prompts** — punchy, single-punchline.
- Medium stories: **6–9 prompts** — escalating situation, satisfying close.
- Long stories: **10–15 prompts** — full narrative arc, multiple characters/beats.
- Prompt `key` must be unique within a story (e.g. `noun1`, `noun2` not `noun`, `noun`).
- Prompt `label` should be a **vague category**, not a narrative slot. The fun comes from the surprise of seeing a random word land in a specific role — that surprise dies the moment the label leaks the context.
  - ✅ Good: `Adjective`, `Noun`, `Celebrity`, `City`, `Food`, `a name`, `a profession`, `a noise`, `an emotion`, `a body part`, `a piece of clothing`, `a song title`, `a drink`, `a number`, `a verb ending in -ing`
  - ❌ Bad: `the therapist's name`, `the post that broke the camel's back`, `a hyper-specific rule the mod enforced`, `the answer your partner gave that blindsided you`
  - Rule of thumb: a player reading the prompt out of context should have no idea what story it belongs to.
- Template `{{key}}` placeholders must match the prompt keys exactly.
- **Template syntax is `{{key}}` — double braces, always.** Single-brace `{key}` is invalid and will be rejected by the content validator. The app renderer only substitutes `{{key}}`; anything else will leak into the rendered story (or worse, render as a stray `}`). Every `{{key}}` in the template must have a matching prompt `key`, and every prompt `key` must appear in the template.
- Don't pad stories with throwaway prompts just to hit a word count.

## Pack Target Size

- **15–20 stories** per major available pack at launch.
- Distribution goal: **40% short, 40% medium, 20% long**.

---

## Pack Tone Guides

### 🎉 Party Pack (`party`)

**Rating:** teen

**Vibe:** Loud, physical, and socially chaotic. Stories should feel like they happen at midnight with people who are making increasingly questionable decisions.

**Good topics:** party games gone wrong, suspicious drinks, dancing disasters, someone definitely texting their ex, the host's bathroom situation.

**Avoid:** anything mean-spirited or that punches down. Chaos, not cruelty.

**Example prompt labels:** "a drinking game you invented", "something you find in the couch", "a bad dance move", "the reason you're blocked on three apps"

---

### 🧠 Internet Brainrot (`internet-brainrot`)

**Rating:** teen

**Vibe:** Deeply online. Stories should feel like they were written by someone who has been on their phone for 14 hours straight and has lost the ability to perceive time.

**Good topics:** viral moments, comment sections, posting something unhinged, parasocial spirals, the algorithm, discord drama, niche fandoms.

**Avoid:** targeting specific real people or brands. Keep it recognizable but fictional.

**Example prompt labels:** "a post that should not have gone viral", "what the algorithm recommended at 3am", "a sentence that only makes sense in context", "a reply guy opinion"

---

### ❤️ Couples Pack (`couples`)

**Rating:** teen

**Vibe:** Warmly chaotic. The humor comes from the intimacy and mutual dysfunction of being with someone, not from cruelty or contempt.

**Good topics:** IKEA trips, passive-aggressive text threads, who left the dishes, first apartment disasters, the way you argue about nothing important.

**Avoid:** breakup/divorce framing, anything that feels bitter. Couples should clearly like each other even when they're disasters.

**Example prompt labels:** "the thing we always argue about", "a noise you make when you're annoyed", "something that's definitely your fault", "an excuse that almost worked"

---

### 🦖 Kids Pack (`kids`)

**Rating:** kids

**Vibe:** Gleefully chaotic. Pure silliness. The humor is absurd but completely safe and parent-shareable.

**Good topics:** dinosaurs, snacks, weird animals, school, imaginative disasters, things that only make sense to a 7-year-old.

**Avoid:** anything a parent would cringe at, anything scary, any grown-up topics.

**Example prompt labels:** "your favorite dinosaur", "a really gross food", "a funny animal noise", "something you found on the playground"

---

## Coming-Soon Pack Strategy

Coming-soon packs appear in the browse UI to build anticipation. They are not playable.

Each coming-soon pack should have:

- Correct `status: "coming_soon"`
- Real metadata (title, description, emoji, tags, rating) — this is what players see
- At most 3 tags (keep them high-signal)
- `stories: []` — empty until the pack is ready to ship
- A tone that feels distinct and makes players want to come back

When a coming-soon pack is ready to publish, change `status` to `"available"` and populate `stories`.

### Pack statuses

- `"available"` — playable, appears normally in UI
- `"coming_soon"` — visible in browse, locked, shows teaser copy

Do not add intermediate states without updating the app UI to handle them.
