package com.butterfeetlabs.badlibs.ui.theme

import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Shapes
import androidx.compose.material3.darkColorScheme
import androidx.compose.material3.lightColorScheme
import androidx.compose.runtime.Composable
import androidx.compose.ui.unit.dp

private val LightColors = lightColorScheme(
    primary = ChaosOrange,
    onPrimary = ChaosPaper,
    secondary = ChaosBerry,
    onSecondary = ChaosPaper,
    tertiary = ChaosMint,
    background = ChaosCream,
    onBackground = ChaosInk,
    surface = ChaosCard,
    onSurface = ChaosInk,
    error = ChaosBerry
)

private val DarkColors = darkColorScheme(
    primary = ChaosOrange,
    onPrimary = ChaosInk,
    secondary = ChaosBerry,
    tertiary = ChaosMint,
    background = ChaosInk,
    onBackground = ChaosPaper,
    surface = ChaosOrangeDark,
    onSurface = ChaosPaper,
    error = ChaosBerry
)

private val BadLibsShapes = Shapes(
    extraSmall = androidx.compose.foundation.shape.RoundedCornerShape(8.dp),
    small = androidx.compose.foundation.shape.RoundedCornerShape(14.dp),
    medium = androidx.compose.foundation.shape.RoundedCornerShape(18.dp),
    large = androidx.compose.foundation.shape.RoundedCornerShape(24.dp)
)

@Composable
fun BadLibsTheme(
    darkTheme: Boolean = false,
    content: @Composable () -> Unit
) {
    val colors = if (darkTheme) DarkColors else LightColors

    MaterialTheme(
        colorScheme = colors,
        typography = BadLibsTypography,
        shapes = BadLibsShapes,
        content = content
    )
}
