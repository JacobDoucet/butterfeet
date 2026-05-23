package com.butterfeetlabs.badlibs.ui.theme

import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Shapes
import androidx.compose.material3.darkColorScheme
import androidx.compose.material3.lightColorScheme
import androidx.compose.runtime.Composable
import androidx.compose.runtime.ReadOnlyComposable
import androidx.compose.runtime.compositionLocalOf
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.unit.dp

private val LightColors = lightColorScheme(
    primary = Ink,
    onPrimary = Paper,
    secondary = Highlight,
    onSecondary = Ink,
    tertiary = ChaosBerry,
    background = Paper,
    onBackground = Ink,
    surface = Paper,
    onSurface = Ink,
    surfaceVariant = Color(0xFFEDE6D6),
    onSurfaceVariant = Ink,
    outline = Color(0xFFCDC4B0),
    error = ChaosBerry
)

private val DarkColors = darkColorScheme(
    primary = Highlight,
    onPrimary = Ink,
    secondary = Highlight,
    tertiary = ChaosBerry,
    background = Ink,
    onBackground = Paper,
    surface = Color(0xFF1A1A1D),
    onSurface = Paper,
    error = ChaosBerry
)

private val BadLibsShapes = Shapes(
    extraSmall = androidx.compose.foundation.shape.RoundedCornerShape(8.dp),
    small = androidx.compose.foundation.shape.RoundedCornerShape(14.dp),
    medium = androidx.compose.foundation.shape.RoundedCornerShape(18.dp),
    large = androidx.compose.foundation.shape.RoundedCornerShape(24.dp)
)

/**
 * Bad Libs design tokens exposed via [LocalBadLibsTokens]. Holds families and
 * brand colors that exist outside of Material's color scheme so screens can
 * reach the full identity without going through MaterialTheme.
 */
data class BadLibsTokens(
    val displayFamily: FontFamily = BricolageFamily,
    val bodyFamily: FontFamily = InterFamily,
    val monoFamily: FontFamily = MonoFamily,
    val ink: Color = Ink,
    val paper: Color = Paper,
    val highlight: Color = Highlight,
    val shadowSoft: Color = ShadowSoft
) {
    fun palette(packId: String, available: Boolean = true) =
        paletteForPack(packId, available)
}

val LocalBadLibsTokens = compositionLocalOf { BadLibsTokens() }

object BadLibs {
    val tokens: BadLibsTokens
        @Composable
        @ReadOnlyComposable
        get() = LocalBadLibsTokens.current
}

@Composable
fun BadLibsTheme(
    darkTheme: Boolean = false,
    content: @Composable () -> Unit
) {
    val colors = if (darkTheme) DarkColors else LightColors

    androidx.compose.runtime.CompositionLocalProvider(
        LocalBadLibsTokens provides BadLibsTokens()
    ) {
        MaterialTheme(
            colorScheme = colors,
            typography = BadLibsTypography,
            shapes = BadLibsShapes,
            content = content
        )
    }
}
