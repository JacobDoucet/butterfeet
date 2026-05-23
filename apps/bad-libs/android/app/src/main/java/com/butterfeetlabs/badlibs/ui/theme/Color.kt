package com.butterfeetlabs.badlibs.ui.theme

import androidx.compose.ui.graphics.Color

// ---- Core identity palette ---------------------------------------------------
// Bad Libs ships on warm Paper with deep Ink and a single Highlight yellow.
val Ink = Color(0xFF0E0E10)
val Paper = Color(0xFFF6F1E8)
val Highlight = Color(0xFFFFE36E)
val ShadowSoft = Color(0x1F0E0E10)

// ---- Legacy aliases (kept so existing screens compile unchanged) -------------
val ChaosOrange = Color(0xFFFF8A3D)
val ChaosOrangeDark = Color(0xFFE16813)
val ChaosCream = Paper
val ChaosPaper = Paper
val ChaosInk = Ink
val ChaosBerry = Color(0xFFC53A73)
val ChaosMint = Color(0xFF2DAA7C)
val ChaosCard = Color(0xFFFFFAF4)

// ---- Per-pack identity -------------------------------------------------------
data class PackPalette(
    val primary: Color,
    val secondary: Color,
    val vibe: String
)

val ComingSoonPalette = PackPalette(
    primary = Color(0xFF8C9099),
    secondary = Color(0xFFB8BCC4),
    vibe = "Soon"
)

val BadLibsPackPalettes: Map<String, PackPalette> = mapOf(
    "party"              to PackPalette(Color(0xFFFF4D5A), Color(0xFF1E1A4A), "Loud"),
    "internet-brainrot"  to PackPalette(Color(0xFF7C5CFF), Color(0xFF00E0B8), "Glitched"),
    "couples"            to PackPalette(Color(0xFFE27D60), Color(0xFF3B5A57), "Tender chaos"),
    "kids"               to PackPalette(Color(0xFFFFB400), Color(0xFF3DA5D9), "Sticky"),
    "office"             to PackPalette(Color(0xFF5B6770), Color(0xFFC8A951), "Fluorescent"),
    "vacation-disaster"  to PackPalette(Color(0xFF00A6A6), Color(0xFFFF7B00), "Heatstroke"),
    "wedding"            to PackPalette(Color(0xFFD9C5B2), Color(0xFF7E1F2D), "Performative"),
    "holiday"            to PackPalette(Color(0xFFC0392B), Color(0xFF27632A), "Forced cheer"),
    "gen-alpha"          to PackPalette(Color(0xFFFF66C4), Color(0xFF39FF14), "Untranslatable")
)

fun paletteForPack(packId: String, available: Boolean = true): PackPalette {
    if (!available) return ComingSoonPalette
    return BadLibsPackPalettes[packId] ?: PackPalette(
        primary = ChaosOrange,
        secondary = ChaosBerry,
        vibe = "Chaos"
    )
}

/** Pick Ink or Paper as readable foreground for a given background. */
fun onColorFor(background: Color): Color {
    val r = background.red
    val g = background.green
    val b = background.blue
    val l = 0.2126f * r + 0.7152f * g + 0.0722f * b
    return if (l > 0.55f) Ink else Paper
}
