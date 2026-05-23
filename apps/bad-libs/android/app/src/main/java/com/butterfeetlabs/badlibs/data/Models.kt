package com.butterfeetlabs.badlibs.data

import kotlinx.serialization.Serializable

@Serializable
data class StoryPack(
    val id: String,
    val title: String,
    val description: String,
    val rating: String,
    val status: String,
    val emoji: String,
    val tags: List<String> = emptyList(),
    val accentName: String? = null,
    val stories: List<Story>
)

@Serializable
data class Story(
    val id: String,
    val title: String,
    val tags: List<String> = emptyList(),
    val prompts: List<Prompt>,
    val template: String
)

@Serializable
data class Prompt(
    val key: String,
    val label: String
)

data class CompletedStory(
    val storyId: String,
    val storyTitle: String,
    val values: Map<String, String>,
    val renderedText: String
)
