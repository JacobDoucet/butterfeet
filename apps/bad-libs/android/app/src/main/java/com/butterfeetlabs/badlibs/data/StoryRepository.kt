package com.butterfeetlabs.badlibs.data

import android.content.Context
import kotlinx.serialization.json.Json
import kotlinx.serialization.decodeFromString

class StoryRepository(
    private val context: Context,
    private val json: Json = Json { ignoreUnknownKeys = true }
) {
    private val packDisplayOrder = mapOf(
        "party" to 0,
        "internet-brainrot" to 1,
        "couples" to 2,
        "kids" to 3,
        "office" to 4,
        "reality-tv" to 5,
        "science-fiction" to 6,
        "crime-documentary" to 7,
        "gen-alpha" to 8
    )

    fun loadStoryPacks(): Result<List<StoryPack>> {
        return runCatching {
            val packDir = "packs"
            val files = context.assets.list(packDir)?.filter { it.endsWith(".json") }.orEmpty()
            files.map { fileName ->
                val content = context.assets.open("$packDir/$fileName").bufferedReader().use { it.readText() }
                json.decodeFromString<StoryPack>(content)
            }.sortedWith(
                compareBy<StoryPack> { if (it.status.equals("available", ignoreCase = true)) 0 else 1 }
                    .thenBy { packDisplayOrder[it.id] ?: Int.MAX_VALUE }
                    .thenBy { it.title.lowercase() }
            )
        }
    }
}
