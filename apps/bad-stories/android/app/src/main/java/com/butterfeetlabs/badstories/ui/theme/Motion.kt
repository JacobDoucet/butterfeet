package com.butterfeetlabs.badstories.ui.theme

import androidx.compose.animation.core.CubicBezierEasing
import androidx.compose.animation.core.Easing

object MotionTokens {
    const val DurationShortMs = 120
    const val DurationMediumMs = 220
    const val DurationExpressiveMs = 420

    val EaseOutStandard: Easing = CubicBezierEasing(0.2f, 0f, 0f, 1f)
    val EaseOutEmphasized: Easing = CubicBezierEasing(0.05f, 0.7f, 0.1f, 1f)
}
