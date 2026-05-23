package com.butterfeetlabs.badstories.data

enum class StoryLengthCategory(
    val label: String,
    val emoji: String
) {
    Quick(label = "Quick", emoji = "⚡"),
    Medium(label = "Medium", emoji = "🍔"),
    BrainDamage(label = "Brain Damage", emoji = "🧠");

    companion object {
        fun fromPromptCount(count: Int): StoryLengthCategory {
            return when {
                count <= 5 -> Quick
                count <= 9 -> Medium
                else -> BrainDamage
            }
        }

        // TODO: Add estimated completion time buckets per category for richer UX hints.
        // TODO: Add multiplayer suitability tags based on category and prompt structure.
        // TODO: Add optional difficulty/chaos ratings as pack metadata evolves.
    }
}
