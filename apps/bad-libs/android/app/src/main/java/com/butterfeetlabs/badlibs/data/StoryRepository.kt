package com.butterfeetlabs.badlibs.data

import android.content.Context
import kotlinx.serialization.json.Json
import kotlinx.serialization.decodeFromString

class StoryRepository(
    private val context: Context,
    private val json: Json = Json { ignoreUnknownKeys = true }
) {
    fun loadStoryPacks(): Result<List<StoryPack>> {
        return runCatching {
            val packDir = "packs"
            val files = context.assets.list(packDir)?.filter { it.endsWith(".json") }.orEmpty()
            files.map { fileName ->
                val content = context.assets.open("$packDir/$fileName").bufferedReader().use { it.readText() }
                json.decodeFromString<StoryPack>(content)
            }
        }
    }
}
