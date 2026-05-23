package com.butterfeetlabs.badlibs.data

import kotlinx.serialization.Serializable

@Serializable
data class StoryPack(
    val id: String,
    val title: String,
    val description: String,
    val stories: List<Story>
)

@Serializable
data class Story(
    val id: String,
    val title: String,
    val rating: String,
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
