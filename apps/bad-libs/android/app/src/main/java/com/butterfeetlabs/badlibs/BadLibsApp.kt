package com.butterfeetlabs.badlibs

import android.content.Intent
import android.content.Context
import androidx.compose.animation.AnimatedContent
import androidx.compose.animation.AnimatedVisibility
import androidx.compose.animation.animateContentSize
import androidx.compose.animation.animateColorAsState
import androidx.compose.animation.fadeIn
import androidx.compose.animation.fadeOut
import androidx.compose.animation.scaleIn
import androidx.compose.animation.core.LinearEasing
import androidx.compose.animation.core.RepeatMode
import androidx.compose.animation.core.animateFloatAsState
import androidx.compose.animation.core.animateFloat
import androidx.compose.animation.core.infiniteRepeatable
import androidx.compose.animation.core.rememberInfiniteTransition
import androidx.compose.animation.slideInHorizontally
import androidx.compose.animation.slideOutHorizontally
import androidx.compose.animation.core.tween
import androidx.compose.foundation.BorderStroke
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.BoxScope
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.imePadding
import androidx.compose.foundation.layout.navigationBarsPadding
import androidx.compose.foundation.layout.offset
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.WindowInsets
import androidx.compose.foundation.layout.ExperimentalLayoutApi
import androidx.compose.foundation.layout.isImeVisible
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.text.KeyboardActions
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Card
import androidx.compose.material3.CardDefaults
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.LinearProgressIndicator
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.OutlinedButton
import androidx.compose.material3.Scaffold
import androidx.compose.material3.SnackbarHost
import androidx.compose.material3.SnackbarHostState
import androidx.compose.material3.Text
import androidx.compose.material3.TextFieldDefaults
import androidx.compose.material3.TextButton
import androidx.compose.material3.TopAppBar
import androidx.compose.material3.TopAppBarDefaults
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateMapOf
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.foundation.interaction.MutableInteractionSource
import androidx.compose.foundation.interaction.collectIsPressedAsState
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.graphicsLayer
import androidx.compose.ui.hapticfeedback.HapticFeedbackType
import androidx.compose.ui.focus.FocusRequester
import androidx.compose.ui.focus.focusRequester
import androidx.compose.ui.platform.LocalFocusManager
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalHapticFeedback
import androidx.compose.ui.platform.LocalSoftwareKeyboardController
import androidx.compose.ui.text.input.ImeAction
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.sp
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.Dp
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavType
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import androidx.navigation.navArgument
import androidx.compose.foundation.gestures.awaitEachGesture
import androidx.compose.foundation.gestures.awaitFirstDown
import androidx.compose.ui.input.pointer.PointerEventPass
import androidx.compose.ui.input.pointer.changedToUp
import androidx.compose.ui.input.pointer.pointerInput
import androidx.compose.ui.input.pointer.positionChange
import androidx.compose.ui.platform.LocalDensity
import com.butterfeetlabs.badlibs.data.CompletedStory
import com.butterfeetlabs.badlibs.data.RenderedToken
import com.butterfeetlabs.badlibs.data.Story
import com.butterfeetlabs.badlibs.data.StoryLengthCategory
import com.butterfeetlabs.badlibs.data.StoryPack
import com.butterfeetlabs.badlibs.data.StoryRepository
import com.butterfeetlabs.badlibs.ui.theme.BadLibs
import com.butterfeetlabs.badlibs.ui.theme.BadLibsTheme
import com.butterfeetlabs.badlibs.ui.theme.ChaosOrange
import com.butterfeetlabs.badlibs.ui.theme.MotionTokens
import com.butterfeetlabs.badlibs.ui.theme.onColorFor
import com.butterfeetlabs.badlibs.ui.theme.paletteForPack
import com.butterfeetlabs.badlibs.ui.theme.softPackTint
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.text.style.TextDecoration
import androidx.compose.animation.core.FastOutSlowInEasing
import androidx.compose.animation.slideInVertically
import androidx.compose.foundation.layout.WindowInsetsSides
import androidx.compose.foundation.layout.statusBars
import androidx.compose.foundation.layout.statusBarsPadding
import androidx.compose.foundation.layout.windowInsetsPadding
import androidx.compose.foundation.layout.only
import java.time.LocalDate
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch

private sealed class Screen(val route: String) {
    data object Intro : Screen("intro")
    data object Home : Screen("home")
    data object Packs : Screen("packs")
    data object Stories : Screen("stories/{packId}")
    data object QuickChaos : Screen("quick-chaos")
    data object Prompts : Screen("prompts/{packId}/{storyId}")
    data object Result : Screen("result")
}

private enum class QuickChaosPhase {
    Spinning,
    Locked
}

private enum class RevealPhase {
    Anticipation,
    Reveal,
    Rest,
    Visible
}

private const val AppPrefs = "bad_libs_prefs"
private const val IntroSeenKey = "intro_seen"
private const val LastPackIdKey = "last_pack_id"
private const val LastStoryIdKey = "last_story_id"

private val PackAccents = listOf(
    Color(0xFFEA5B2A),
    Color(0xFF2E9E74),
    Color(0xFF7E53D7),
    Color(0xFFD5457A),
    Color(0xFF2E83C6)
)

private val ComingSoonAccent = com.butterfeetlabs.badlibs.ui.theme.ComingSoonPalette.primary

/**
 * Wraps a screen so an in-content swipe from left to right pops the back stack.
 * Gesture only fires when horizontal travel clearly dominates vertical motion,
 * so it leaves vertical scrolling untouched.
 */
@Composable
private fun SwipeBack(
    onBack: () -> Unit,
    content: @Composable () -> Unit
) {
    val density = LocalDensity.current
    val triggerPx = with(density) { 96.dp.toPx() }
    Box(
        modifier = Modifier
            .fillMaxSize()
            .pointerInput(onBack) {
                awaitEachGesture {
                    awaitFirstDown(requireUnconsumed = false, pass = PointerEventPass.Initial)
                    var totalX = 0f
                    var totalY = 0f
                    var triggered = false
                    while (true) {
                        val event = awaitPointerEvent(pass = PointerEventPass.Main)
                        val change = event.changes.firstOrNull() ?: break
                        if (!change.pressed) break
                        val delta = change.positionChange()
                        totalX += delta.x
                        totalY += delta.y
                        if (
                            !triggered &&
                            totalX > triggerPx &&
                            totalX > kotlin.math.abs(totalY) * 2f
                        ) {
                            triggered = true
                            change.consume()
                            onBack()
                            break
                        }
                        if (change.changedToUp()) break
                    }
                }
            }
    ) {
        content()
    }
}


private fun accentForSeed(seed: String): Color {
    if (seed.isEmpty()) return ChaosOrange
    val mapped = com.butterfeetlabs.badlibs.ui.theme.BadLibsPackPalettes[seed]
    if (mapped != null) return mapped.primary
    val index = seed.fold(0) { acc, c -> acc + c.code } % PackAccents.size
    return PackAccents[index]
}

private class MainViewModel(private val repository: StoryRepository) : ViewModel() {
    var isLoading by mutableStateOf(true)
        private set

    var packs by mutableStateOf<List<StoryPack>>(emptyList())
        private set

    var loadError by mutableStateOf<String?>(null)
        private set

    var completedStory by mutableStateOf<CompletedStory?>(null)
        private set

    init {
        loadContent()
    }

    fun loadContent() {
        isLoading = true
        loadError = null
        val result = repository.loadStoryPacks()
        result.onSuccess {
            packs = it
        }.onFailure {
            loadError = "Could not load local story packs."
        }
        isLoading = false
    }

    fun findPack(packId: String): StoryPack? = packs.firstOrNull { it.id == packId }

    fun findStory(packId: String, storyId: String): Story? {
        return findPack(packId)?.stories?.firstOrNull { it.id == storyId }
    }

    fun completeStory(packId: String, story: Story, values: Map<String, String>) {
        val rendered = values.entries.fold(story.template) { acc, entry ->
            acc.replace("{${entry.key}}", entry.value)
        }
        val labels = story.prompts.associate { it.key to it.label }
        val tokens = tokenizeTemplate(story.template, values, labels)
        completedStory = CompletedStory(
            packId = packId,
            storyId = story.id,
            storyTitle = story.title,
            values = values,
            renderedText = rendered,
            tokens = tokens
        )
    }

    private fun tokenizeTemplate(
        template: String,
        values: Map<String, String>,
        labels: Map<String, String>
    ): List<com.butterfeetlabs.badlibs.data.RenderedToken> {
        val regex = Regex("\\{([^}]+)\\}")
        val out = mutableListOf<com.butterfeetlabs.badlibs.data.RenderedToken>()
        var cursor = 0
        regex.findAll(template).forEach { match ->
            if (match.range.first > cursor) {
                out += com.butterfeetlabs.badlibs.data.RenderedToken.Static(
                    template.substring(cursor, match.range.first)
                )
            }
            val key = match.groupValues[1]
            out += com.butterfeetlabs.badlibs.data.RenderedToken.Filled(
                value = values[key].orEmpty(),
                promptLabel = labels[key].orEmpty()
            )
            cursor = match.range.last + 1
        }
        if (cursor < template.length) {
            out += com.butterfeetlabs.badlibs.data.RenderedToken.Static(template.substring(cursor))
        }
        return out
    }
}

private class MainViewModelFactory(
    private val repository: StoryRepository
) : ViewModelProvider.Factory {
    override fun <T : ViewModel> create(modelClass: Class<T>): T {
        if (modelClass.isAssignableFrom(MainViewModel::class.java)) {
            @Suppress("UNCHECKED_CAST")
            return MainViewModel(repository) as T
        }
        throw IllegalArgumentException("Unknown ViewModel class")
    }
}

@Composable
fun BadLibsApp() {
    val context = LocalContext.current
    val prefs = remember { context.getSharedPreferences(AppPrefs, Context.MODE_PRIVATE) }
    val navController = rememberNavController()
    val snackbarHostState = remember { SnackbarHostState() }
    val appScope = rememberCoroutineScope()
    var hasSeenIntro by remember {
        mutableStateOf(
            prefs.getBoolean(IntroSeenKey, false)
        )
    }
    var lastPackId by remember { mutableStateOf(prefs.getString(LastPackIdKey, null)) }
    var lastStoryId by remember { mutableStateOf(prefs.getString(LastStoryIdKey, null)) }
    val viewModel: MainViewModel = viewModel(
        factory = MainViewModelFactory(StoryRepository(context))
    )

    fun rememberLastStory(packId: String, storyId: String) {
        prefs.edit()
            .putString(LastPackIdKey, packId)
            .putString(LastStoryIdKey, storyId)
            .apply()
        lastPackId = packId
        lastStoryId = storyId
    }

    fun startStory(packId: String, storyId: String) {
        rememberLastStory(packId, storyId)
        navController.navigate("prompts/$packId/$storyId")
    }

    val allStories = remember(viewModel.packs) {
        viewModel.packs
            .filter { it.status.equals("available", ignoreCase = true) }
            .flatMap { pack -> pack.stories.map { story -> pack.id to story.id } }
    }
    val dailyChallenge = remember(allStories) {
        if (allStories.isEmpty()) {
            null
        } else {
            val seed = LocalDate.now().toEpochDay().toInt()
            val idx = kotlin.math.abs(seed) % allStories.size
            allStories[idx]
        }
    }
    val lastStoryRef = if (lastPackId != null && lastStoryId != null) {
        val pack = viewModel.findPack(lastPackId!!)
        val story = viewModel.findStory(lastPackId!!, lastStoryId!!)
        if (pack != null && story != null) {
            Triple(pack.id, story.id, "${pack.title}: ${story.title}")
        } else {
            null
        }
    } else {
        null
    }

    BadLibsTheme {
        ChaosBackdrop {
            NavHost(
                navController = navController,
                startDestination = if (hasSeenIntro) Screen.Home.route else Screen.Intro.route,
                enterTransition = {
                    slideInHorizontally(
                        initialOffsetX = { width -> (width * 0.25f).toInt() },
                        animationSpec = tween(
                            durationMillis = MotionTokens.DurationMediumMs,
                            easing = MotionTokens.EaseOutStandard
                        )
                    ) + fadeIn(
                        animationSpec = tween(
                            durationMillis = MotionTokens.DurationMediumMs,
                            easing = MotionTokens.EaseOutEmphasized
                        )
                    )
                },
                exitTransition = {
                    slideOutHorizontally(
                        targetOffsetX = { width -> -(width * 0.15f).toInt() },
                        animationSpec = tween(
                            durationMillis = MotionTokens.DurationShortMs,
                            easing = MotionTokens.EaseOutStandard
                        )
                    ) + fadeOut(animationSpec = tween(MotionTokens.DurationShortMs))
                },
                popEnterTransition = {
                    slideInHorizontally(
                        initialOffsetX = { width -> -(width * 0.25f).toInt() },
                        animationSpec = tween(
                            durationMillis = MotionTokens.DurationMediumMs,
                            easing = MotionTokens.EaseOutStandard
                        )
                    ) + fadeIn(animationSpec = tween(MotionTokens.DurationMediumMs))
                },
                popExitTransition = {
                    slideOutHorizontally(
                        targetOffsetX = { width -> (width * 0.15f).toInt() },
                        animationSpec = tween(
                            durationMillis = MotionTokens.DurationShortMs,
                            easing = MotionTokens.EaseOutStandard
                        )
                    ) + fadeOut(animationSpec = tween(MotionTokens.DurationShortMs))
                }
            ) {
                composable(Screen.Intro.route) {
                    IntroScreen(
                        onStart = {
                            prefs.edit()
                                .putBoolean(IntroSeenKey, true)
                                .apply()
                            hasSeenIntro = true
                            navController.navigate(Screen.Home.route) {
                                popUpTo(Screen.Intro.route) { inclusive = true }
                            }
                        }
                    )
                }
                composable(Screen.Home.route) {
                    HomeHubScreen(
                        packs = viewModel.packs,
                        lastStoryLabel = lastStoryRef?.third,
                        onContinueLast = {
                            lastStoryRef?.let { (packId, storyId, _) -> startStory(packId, storyId) }
                        },
                        dailyChallengeLabel = dailyChallenge?.let { (packId, storyId) ->
                            val pack = viewModel.findPack(packId)
                            val story = viewModel.findStory(packId, storyId)
                            if (pack != null && story != null) "${pack.title}: ${story.title}" else null
                        },
                        onPlayDailyChallenge = {
                            dailyChallenge?.let { (packId, storyId) -> startStory(packId, storyId) }
                        },
                        onPlayPacks = { navController.navigate(Screen.Packs.route) },
                        quickPlayCandidates = allStories,
                        onQuickChaos = { navController.navigate(Screen.QuickChaos.route) },
                        onShowIntro = {
                            navController.navigate(Screen.Intro.route)
                        }
                    )
                }
                composable(Screen.Packs.route) {
                    SwipeBack(onBack = { navController.popBackStack() }) {
                        PackListScreen(
                            isLoading = viewModel.isLoading,
                            error = viewModel.loadError,
                            packs = viewModel.packs,
                            onBack = { navController.popBackStack() },
                            onRetry = { viewModel.loadContent() },
                            onOpenPack = { packId -> navController.navigate("stories/$packId") },
                            onPackLocked = {
                                appScope.launch {
                                    snackbarHostState.showSnackbar("This pack is still being cooked.")
                                }
                            }
                        )
                    }
                }
                composable(
                    route = Screen.Stories.route,
                    arguments = listOf(navArgument("packId") { type = NavType.StringType })
                ) { backStackEntry ->
                    val packId = backStackEntry.arguments?.getString("packId").orEmpty()
                    val pack = viewModel.findPack(packId)
                    SwipeBack(onBack = { navController.popBackStack() }) {
                        StoryListScreen(
                            pack = pack,
                            onBack = { navController.popBackStack() },
                            onOpenStory = { storyId -> navController.navigate("prompts/$packId/$storyId") }
                        )
                    }
                }
                composable(
                    route = Screen.QuickChaos.route
                ) { backStackEntry ->
                    fun quickPlayLabelFor(pair: Pair<String, String>): String {
                        val pack = viewModel.findPack(pair.first)
                        val story = viewModel.findStory(pair.first, pair.second)
                        return if (pack != null && story != null) {
                            "${pack.title}: ${story.title}"
                        } else {
                            "Random story"
                        }
                    }

                    QuickChaosScreen(
                        candidates = allStories,
                        labelFor = ::quickPlayLabelFor,
                        onContinue = { packId, storyId ->
                            rememberLastStory(packId, storyId)
                            navController.navigate("prompts/$packId/$storyId") {
                                popUpTo(Screen.Home.route) { inclusive = false }
                                launchSingleTop = true
                            }
                        }
                    )
                }
                composable(
                    route = Screen.Prompts.route,
                    arguments = listOf(
                        navArgument("packId") { type = NavType.StringType },
                        navArgument("storyId") { type = NavType.StringType }
                    )
                ) { backStackEntry ->
                    val packId = backStackEntry.arguments?.getString("packId").orEmpty()
                    val storyId = backStackEntry.arguments?.getString("storyId").orEmpty()
                    val story = viewModel.findStory(packId, storyId)
                    if (story != null) {
                        prefs.edit()
                            .putString(LastPackIdKey, packId)
                            .putString(LastStoryIdKey, storyId)
                            .apply()
                        lastPackId = packId
                        lastStoryId = storyId
                    }
                    PromptInputScreen(
                        story = story,
                        snackbarHostState = snackbarHostState,
                        onBack = { navController.popBackStack() },
                        onComplete = { values ->
                            if (story != null) {
                                viewModel.completeStory(packId, story, values)
                                navController.navigate(Screen.Result.route)
                            }
                        }
                    )
                }
                composable(Screen.Result.route) {
                    SwipeBack(onBack = { navController.popBackStack() }) {
                        ResultScreen(
                            completedStory = viewModel.completedStory,
                            onRemix = {
                                navController.popBackStack()
                            },
                            onPlayAgain = {
                                val moved = navController.popBackStack(Screen.Packs.route, inclusive = false)
                                if (!moved) {
                                    navController.navigate(Screen.Packs.route)
                                }
                            },
                            onBackHome = {
                                navController.popBackStack(Screen.Home.route, inclusive = false)
                            }
                        )
                    }
                }
            }

            SnackbarHost(
                hostState = snackbarHostState,
                modifier = Modifier.align(Alignment.BottomCenter)
            )
        }
    }
}

@Composable
private fun ChaosBackdrop(content: @Composable BoxScope.() -> Unit) {
    val tokens = BadLibs.tokens
    Box(
        modifier = Modifier
            .fillMaxSize()
            .background(tokens.paper)
    ) {
        Box(
            modifier = Modifier
                .size(360.dp)
                .offset(x = (-80).dp, y = (-120).dp)
                .background(tokens.highlight.copy(alpha = 0.55f), CircleShape)
        )
        Box(
            modifier = Modifier
                .size(220.dp)
                .offset(x = 260.dp, y = 520.dp)
                .background(tokens.ink.copy(alpha = 0.06f), CircleShape)
        )
        content()
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
private fun ChaosScaffold(
    title: String,
    onBack: (() -> Unit)? = null,
    content: @Composable (PaddingValues) -> Unit
) {
    Scaffold(
        containerColor = Color.Transparent,
        topBar = {
            TopAppBar(
                title = {
                    Text(
                        text = "// " + title.lowercase(),
                        style = MaterialTheme.typography.labelLarge,
                        color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.6f)
                    )
                },
                navigationIcon = {
                    if (onBack != null) {
                        TextButton(onClick = onBack) {
                            Text(
                                "← back",
                                style = MaterialTheme.typography.labelLarge,
                                color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.7f)
                            )
                        }
                    }
                },
                colors = TopAppBarDefaults.topAppBarColors(
                    containerColor = Color.Transparent,
                    titleContentColor = MaterialTheme.colorScheme.onBackground
                )
            )
        },
        content = content
    )
}

@Composable
private fun rememberPressBounce(): Pair<MutableInteractionSource, Float> {
    val source = remember { MutableInteractionSource() }
    val pressed by source.collectIsPressedAsState()
    val scale by animateFloatAsState(
        targetValue = if (pressed) 0.96f else 1f,
        animationSpec = tween(durationMillis = 140, easing = FastOutSlowInEasing),
        label = "press_bounce"
    )
    return source to scale
}

@Composable
private fun ChaosPrimaryButton(
    label: String,
    onClick: () -> Unit,
    modifier: Modifier = Modifier
) {
    val tokens = BadLibs.tokens
    val haptics = LocalHapticFeedback.current
    val (source, scale) = rememberPressBounce()
    Button(
        onClick = {
            haptics.performHapticFeedback(HapticFeedbackType.LongPress)
            onClick()
        },
        modifier = modifier.graphicsLayer { scaleX = scale; scaleY = scale },
        interactionSource = source,
        colors = ButtonDefaults.buttonColors(
            containerColor = tokens.ink,
            contentColor = tokens.paper
        ),
        shape = RoundedCornerShape(14.dp)
    ) {
        Text(
            text = label,
            style = MaterialTheme.typography.titleMedium,
            fontWeight = FontWeight.ExtraBold
        )
    }
}

@Composable
private fun GameFooterBar(
    primaryLabel: String,
    onPrimary: () -> Unit,
    onBack: () -> Unit,
    modifier: Modifier = Modifier,
    backLabel: String = "Back"
) {
    Row(
        modifier = modifier
            .fillMaxWidth()
            .animateContentSize(
                animationSpec = tween(
                    durationMillis = MotionTokens.DurationMediumMs,
                    easing = MotionTokens.EaseOutStandard
                )
            ),
        horizontalArrangement = Arrangement.spacedBy(12.dp)
    ) {
        OutlinedButton(onClick = onBack, modifier = Modifier.weight(0.8f)) {
            AnimatedContent(
                targetState = backLabel,
                label = "footer_back_label"
            ) { label ->
                Text(label)
            }
        }
        ChaosPrimaryButton(
            label = primaryLabel,
            onClick = onPrimary,
            modifier = Modifier.weight(1.2f)
        )
    }
}

@Composable
private fun ChaosSurfaceCard(
    modifier: Modifier = Modifier,
    onClick: (() -> Unit)? = null,
    accentColor: Color? = null,
    highlighted: Boolean = false,
    content: @Composable () -> Unit
) {
    val targetBorder = when {
        highlighted -> (accentColor ?: ChaosOrange).copy(alpha = 0.95f)
        else -> Color.White
    }
    val animatedBorder by animateColorAsState(
        targetValue = targetBorder,
        animationSpec = tween(
            durationMillis = MotionTokens.DurationMediumMs,
            easing = MotionTokens.EaseOutStandard
        ),
        label = "card_border"
    )

    if (onClick != null) {
        Card(
            modifier = modifier,
            onClick = onClick,
            shape = RoundedCornerShape(22.dp),
            border = BorderStroke(if (highlighted) 2.dp else 1.dp, animatedBorder),
            elevation = CardDefaults.cardElevation(defaultElevation = 0.dp),
            colors = CardDefaults.cardColors(containerColor = Color.White)
        ) {
            content()
        }
    } else {
        Card(
            modifier = modifier,
            shape = RoundedCornerShape(22.dp),
            border = BorderStroke(if (highlighted) 2.dp else 1.dp, animatedBorder),
            elevation = CardDefaults.cardElevation(defaultElevation = 0.dp),
            colors = CardDefaults.cardColors(containerColor = Color.White)
        ) {
            content()
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
private fun IntroScreen(onStart: () -> Unit) {
    val tokens = BadLibs.tokens
    ChaosScaffold(title = "Bad Libs") { innerPadding ->
        Column(
            modifier = Modifier
                .padding(innerPadding)
                .fillMaxSize()
                .verticalScroll(rememberScrollState())
                .padding(horizontal = 24.dp)
                .padding(bottom = 24.dp)
                .navigationBarsPadding(),
            verticalArrangement = Arrangement.spacedBy(20.dp)
        ) {
            Text(
                text = "// terrible stories, on demand",
                style = MaterialTheme.typography.labelMedium,
                color = tokens.ink.copy(alpha = 0.6f)
            )
            Text(
                text = "BAD\nLIBS.",
                style = MaterialTheme.typography.displayLarge,
                color = tokens.ink,
                lineHeight = 64.sp
            )
            Text(
                text = "A fill-in-the-blank disaster generator for road trips, dinner parties, and the group chat that never sleeps.",
                style = MaterialTheme.typography.bodyLarge,
                color = tokens.ink.copy(alpha = 0.78f)
            )

            ChaosPrimaryButton(
                label = "start the chaos →",
                onClick = onStart,
                modifier = Modifier
                    .fillMaxWidth()
                    .padding(top = 6.dp)
            )

            Spacer(modifier = Modifier.height(8.dp))
            Text(
                text = "// how it works",
                style = MaterialTheme.typography.labelMedium,
                color = tokens.ink.copy(alpha = 0.55f)
            )
            IntroStep(number = "01", title = "Pick a pack.", body = "Choose your flavor of nonsense.")
            IntroStep(number = "02", title = "Fill the prompts.", body = "Weird words win. Spell-check optional.")
            IntroStep(number = "03", title = "Reveal & share.", body = "Drop the masterpiece in the group chat.")
        }
    }
}

@Composable
private fun IntroStep(number: String, title: String, body: String) {
    val tokens = BadLibs.tokens
    Row(
        modifier = Modifier.fillMaxWidth(),
        horizontalArrangement = Arrangement.spacedBy(16.dp)
    ) {
        Text(
            text = number,
            style = MaterialTheme.typography.labelLarge,
            color = tokens.ink.copy(alpha = 0.5f),
            modifier = Modifier.padding(top = 4.dp)
        )
        Column(modifier = Modifier.weight(1f)) {
            Text(
                text = title,
                style = MaterialTheme.typography.titleLarge,
                color = tokens.ink
            )
            Spacer(modifier = Modifier.height(2.dp))
            Text(
                text = body,
                style = MaterialTheme.typography.bodyMedium,
                color = tokens.ink.copy(alpha = 0.72f)
            )
        }
    }
}

@Composable
private fun HomeHubScreen(
    packs: List<StoryPack>,
    lastStoryLabel: String?,
    onContinueLast: () -> Unit,
    dailyChallengeLabel: String?,
    onPlayDailyChallenge: () -> Unit,
    onPlayPacks: () -> Unit,
    quickPlayCandidates: List<Pair<String, String>>,
    onQuickChaos: () -> Unit,
    onShowIntro: () -> Unit
) {
    val canQuickPlay = quickPlayCandidates.isNotEmpty()
    val tokens = BadLibs.tokens

    ChaosScaffold(title = "Bad Libs") { innerPadding ->
        Column(
            modifier = Modifier
                .padding(innerPadding)
                .fillMaxSize()
                .verticalScroll(rememberScrollState())
                .padding(horizontal = 24.dp)
                .padding(bottom = 16.dp)
                .navigationBarsPadding(),
            verticalArrangement = Arrangement.spacedBy(14.dp)
        ) {
            Text(
                text = "// home",
                style = MaterialTheme.typography.labelMedium,
                color = tokens.ink.copy(alpha = 0.55f)
            )
            Text(
                text = "Make stories\nworse, together.",
                style = MaterialTheme.typography.displayMedium,
                color = tokens.ink,
                lineHeight = 50.sp
            )
            Spacer(modifier = Modifier.height(4.dp))

            if (lastStoryLabel != null) {
                HomeActionTile(
                    eyebrow = "// continue",
                    title = "Pick up where you left off",
                    subtitle = lastStoryLabel,
                    primary = true,
                    onClick = onContinueLast
                )
            }

            if (dailyChallengeLabel != null) {
                HomeActionTile(
                    eyebrow = "// today",
                    title = "Daily Challenge",
                    subtitle = dailyChallengeLabel,
                    primary = lastStoryLabel == null,
                    onClick = onPlayDailyChallenge
                )
            }

            HomeActionTile(
                eyebrow = "// library",
                title = "Browse Packs",
                subtitle = if (packs.isEmpty()) "No packs loaded yet" else "Pick a flavor and start causing damage.",
                primary = false,
                onClick = onPlayPacks
            )

            HomeActionTile(
                eyebrow = "// random",
                title = "Quick Chaos",
                subtitle = if (canQuickPlay) "Random story, dramatic entrance." else "No stories loaded yet.",
                primary = false,
                enabled = canQuickPlay,
                onClick = onQuickChaos
            )

            Spacer(modifier = Modifier.height(8.dp))

            Box(modifier = Modifier.fillMaxWidth(), contentAlignment = Alignment.Center) {
                TextButton(onClick = onShowIntro) {
                    Text(
                        "// replay intro",
                        style = MaterialTheme.typography.labelMedium,
                        color = tokens.ink.copy(alpha = 0.55f)
                    )
                }
            }
        }
    }
}

@Composable
private fun HomeActionTile(
    eyebrow: String,
    title: String,
    subtitle: String,
    primary: Boolean,
    onClick: () -> Unit,
    enabled: Boolean = true
) {
    val tokens = BadLibs.tokens
    val haptics = LocalHapticFeedback.current
    val (source, scale) = rememberPressBounce()
    val container = when {
        !enabled -> tokens.ink.copy(alpha = 0.06f)
        primary -> tokens.ink
        else -> Color.Transparent
    }
    val onContainer = when {
        !enabled -> tokens.ink.copy(alpha = 0.4f)
        primary -> tokens.paper
        else -> tokens.ink
    }
    val border = when {
        !enabled -> tokens.ink.copy(alpha = 0.12f)
        primary -> Color.Transparent
        else -> tokens.ink.copy(alpha = 0.45f)
    }

    androidx.compose.material3.Surface(
        onClick = {
            haptics.performHapticFeedback(HapticFeedbackType.LongPress)
            onClick()
        },
        enabled = enabled,
        interactionSource = source,
        modifier = Modifier
            .fillMaxWidth()
            .graphicsLayer { scaleX = scale; scaleY = scale },
        shape = RoundedCornerShape(22.dp),
        color = container,
        border = BorderStroke(if (primary) 0.dp else 1.5.dp, border)
    ) {
        Row(
            modifier = Modifier.padding(horizontal = 20.dp, vertical = 18.dp),
            horizontalArrangement = Arrangement.SpaceBetween,
            verticalAlignment = Alignment.CenterVertically
        ) {
            Column(modifier = Modifier.weight(1f)) {
                Text(
                    text = eyebrow,
                    style = MaterialTheme.typography.labelMedium,
                    color = onContainer.copy(alpha = 0.62f)
                )
                Spacer(modifier = Modifier.height(4.dp))
                Text(
                    text = title,
                    style = MaterialTheme.typography.titleLarge,
                    color = onContainer
                )
                Spacer(modifier = Modifier.height(4.dp))
                Text(
                    text = subtitle,
                    style = MaterialTheme.typography.bodyMedium,
                    color = onContainer.copy(alpha = 0.72f),
                    lineHeight = 19.sp
                )
            }
            Text(
                text = "→",
                style = MaterialTheme.typography.displaySmall,
                color = onContainer
            )
        }
    }
}

@Composable
private fun ChaosStepCard(
    emoji: String,
    title: String,
    copy: String
) {
    ChaosSurfaceCard(modifier = Modifier.fillMaxWidth()) {
        Row(
            modifier = Modifier
                .fillMaxWidth()
                .padding(16.dp),
            horizontalArrangement = Arrangement.spacedBy(14.dp)
        ) {
            Box(
                modifier = Modifier
                    .size(34.dp)
                    .background(ChaosOrange, RoundedCornerShape(10.dp)),
                contentAlignment = Alignment.Center
            ) {
                Text(text = emoji, color = Color.White, fontWeight = FontWeight.Bold)
            }
            Column {
                Text(text = title, style = MaterialTheme.typography.titleMedium)
                Spacer(modifier = Modifier.height(4.dp))
                Text(text = copy, style = MaterialTheme.typography.bodyMedium)
            }
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class, ExperimentalLayoutApi::class)
@Composable
private fun PackListScreen(
    isLoading: Boolean,
    error: String?,
    packs: List<StoryPack>,
    onBack: () -> Unit,
    onRetry: () -> Unit,
    onOpenPack: (String) -> Unit,
    onPackLocked: () -> Unit
) {
    ChaosScaffold(title = "Choose A Pack") { innerPadding ->
        when {
            isLoading -> {
                Box(
                    modifier = Modifier
                        .padding(innerPadding)
                        .fillMaxSize(),
                    contentAlignment = Alignment.Center
                ) {
                    Column(horizontalAlignment = Alignment.CenterHorizontally) {
                        CircularProgressIndicator(color = ChaosOrange)
                        Spacer(modifier = Modifier.height(12.dp))
                        Text("Warming up nonsense generator...")
                    }
                }
            }

            error != null -> {
                Column(
                    modifier = Modifier
                        .padding(innerPadding)
                        .padding(24.dp)
                        .fillMaxSize(),
                    verticalArrangement = Arrangement.Center,
                    horizontalAlignment = Alignment.CenterHorizontally
                ) {
                    Text(
                        text = "Oops. The story goblins dropped the packs.",
                        style = MaterialTheme.typography.titleMedium,
                        textAlign = TextAlign.Center
                    )
                    Spacer(modifier = Modifier.height(8.dp))
                    Text(error, textAlign = TextAlign.Center)
                    Spacer(modifier = Modifier.height(12.dp))
                    ChaosPrimaryButton(label = "Reload Packs", onClick = onRetry)
                }
            }

            packs.isEmpty() -> {
                Box(
                    modifier = Modifier
                        .padding(innerPadding)
                        .fillMaxSize(),
                    contentAlignment = Alignment.Center
                ) {
                    ChaosSurfaceCard(modifier = Modifier.padding(20.dp)) {
                        Column(
                            modifier = Modifier.padding(18.dp),
                            horizontalAlignment = Alignment.CenterHorizontally
                        ) {
                            Text("No packs yet", style = MaterialTheme.typography.titleMedium)
                            Spacer(modifier = Modifier.height(8.dp))
                            Text(
                                "Drop a pack JSON into assets/packs and come back.",
                                textAlign = TextAlign.Center
                            )
                        }
                    }
                }
            }

            else -> {
                val rouletteScope = rememberCoroutineScope()
                val availablePacks = remember(packs) {
                    packs.filter { it.status.equals("available", ignoreCase = true) }
                }
                var roulettePackId by remember(packs) { mutableStateOf<String?>(null) }
                var rouletteRunning by remember { mutableStateOf(false) }

                Column(
                    modifier = Modifier
                        .padding(innerPadding)
                        .fillMaxSize()
                ) {
                    LazyColumn(
                        modifier = Modifier.weight(1f),
                        contentPadding = PaddingValues(horizontal = 16.dp, vertical = 8.dp),
                        verticalArrangement = Arrangement.spacedBy(12.dp)
                    ) {
                        items(packs, key = { it.id }) { pack ->
                            val isAvailable = pack.status.equals("available", ignoreCase = true)
                            val palette = paletteForPack(pack.id, available = isAvailable)
                            val cardColor = softPackTint(palette, strength = if (isAvailable) 0.18f else 0.10f)
                            val onPaletteColor = onColorFor(cardColor)
                            val accentColor = palette.primary
                            val packIcon = if (isAvailable) pack.emoji else "🔒"
                            val highlighted = isAvailable && roulettePackId == pack.id
                            val haptics = LocalHapticFeedback.current
                            val (source, scale) = rememberPressBounce()

                            androidx.compose.material3.Surface(
                                onClick = {
                                    haptics.performHapticFeedback(HapticFeedbackType.LongPress)
                                    if (isAvailable) onOpenPack(pack.id) else onPackLocked()
                                },
                                interactionSource = source,
                                modifier = Modifier
                                    .fillMaxWidth()
                                    .graphicsLayer { scaleX = scale; scaleY = scale },
                                shape = RoundedCornerShape(22.dp),
                                color = cardColor,
                                border = if (highlighted) BorderStroke(2.dp, BadLibs.tokens.ink) else null
                            ) {
                                Column(
                                    modifier = Modifier
                                        .padding(20.dp)
                                        .graphicsLayer { alpha = if (isAvailable) 1f else 0.85f }
                                ) {
                                    Text(
                                        text = "// ${palette.vibe.lowercase()}",
                                        style = MaterialTheme.typography.labelMedium,
                                        color = accentColor
                                    )
                                    Spacer(modifier = Modifier.height(6.dp))
                                    Row(
                                        modifier = Modifier.fillMaxWidth(),
                                        horizontalArrangement = Arrangement.SpaceBetween,
                                        verticalAlignment = Alignment.Top
                                    ) {
                                        Text(
                                            text = "$packIcon ${pack.title}",
                                            style = MaterialTheme.typography.headlineMedium,
                                            color = onPaletteColor,
                                            modifier = Modifier.weight(1f),
                                            lineHeight = 32.sp
                                        )
                                        Text(
                                            text = if (isAvailable) "${pack.stories.size}" else "—",
                                            style = MaterialTheme.typography.displaySmall,
                                            color = accentColor.copy(alpha = 0.75f)
                                        )
                                    }
                                    Spacer(modifier = Modifier.height(10.dp))
                                    Text(
                                        text = pack.description,
                                        style = MaterialTheme.typography.bodyMedium,
                                        color = onPaletteColor.copy(alpha = 0.85f),
                                        lineHeight = 19.sp
                                    )
                                    Spacer(modifier = Modifier.height(14.dp))
                                    FlowRow(
                                        horizontalArrangement = Arrangement.spacedBy(8.dp),
                                        verticalArrangement = Arrangement.spacedBy(8.dp)
                                    ) {
                                        val visiblePackTags = pack.tags.take(3)
                                        val hiddenPackTagCount = (pack.tags.size - visiblePackTags.size).coerceAtLeast(0)

                                        visiblePackTags.forEach { tag ->
                                            PackTagChip(label = formatStoryTag(tag), onColor = onPaletteColor)
                                        }
                                        if (hiddenPackTagCount > 0) {
                                            PackTagChip(label = "+$hiddenPackTagCount", onColor = onPaletteColor)
                                        }
                                        if (!isAvailable) {
                                            PackTagChip(label = "coming soon", onColor = onPaletteColor)
                                        }
                                    }
                                }
                            }
                        }
                    }

                    GameFooterBar(
                        primaryLabel = if (rouletteRunning) "Rolling..." else "Pick Random",
                        onPrimary = {
                            if (rouletteRunning || availablePacks.isEmpty()) return@GameFooterBar
                            rouletteScope.launch {
                                rouletteRunning = true
                                repeat(8) { step ->
                                    roulettePackId = availablePacks.random().id
                                    delay(55L + (step * 18L))
                                }
                                val winner = availablePacks.random().id
                                roulettePackId = winner
                                delay(140L)
                                rouletteRunning = false
                                onOpenPack(winner)
                            }
                        },
                        onBack = onBack,
                        modifier = Modifier
                            .padding(horizontal = 16.dp, vertical = 12.dp)
                            .navigationBarsPadding()
                    )
                }
            }
        }
    }
}

@Composable
private fun PackCountPill(label: String, accent: Color = Color(0xFFFFB06F)) {
    Box(
        modifier = Modifier
            .background(accent.copy(alpha = 0.22f), RoundedCornerShape(50))
            .padding(horizontal = 12.dp, vertical = 6.dp)
    ) {
        Text(
            text = label,
            style = MaterialTheme.typography.labelLarge,
            color = MaterialTheme.colorScheme.onBackground
        )
    }
}

@Composable
private fun PackTagChip(label: String, onColor: Color) {
    Box(
        modifier = Modifier
            .background(onColor.copy(alpha = 0.14f), RoundedCornerShape(50))
            .padding(horizontal = 10.dp, vertical = 5.dp)
    ) {
        Text(
            text = label,
            style = MaterialTheme.typography.labelMedium,
            color = onColor.copy(alpha = 0.95f)
        )
    }
}

@OptIn(ExperimentalMaterial3Api::class, ExperimentalLayoutApi::class)
@Composable
private fun StoryListScreen(
    pack: StoryPack?,
    onBack: () -> Unit,
    onOpenStory: (String) -> Unit
) {
    ChaosScaffold(title = pack?.title ?: "Stories") { innerPadding ->
        when {
            pack == null -> {
                Box(
                    modifier = Modifier
                        .padding(innerPadding)
                        .fillMaxSize(),
                    contentAlignment = Alignment.Center
                ) {
                    Text("Pack not found.")
                }
            }

            pack.stories.isEmpty() -> {
                Column(
                    modifier = Modifier
                        .padding(innerPadding)
                        .padding(24.dp)
                        .fillMaxSize(),
                    verticalArrangement = Arrangement.Center,
                    horizontalAlignment = Alignment.CenterHorizontally
                ) {
                    Text("This pack is empty. Add more stories and try again.")
                    Spacer(modifier = Modifier.height(12.dp))
                    TextButton(onClick = onBack) {
                        Text("Back")
                    }
                }
            }

            else -> {
                val rouletteScope = rememberCoroutineScope()
                var rouletteStoryId by remember(pack.id) { mutableStateOf<String?>(null) }
                var rouletteRunning by remember { mutableStateOf(false) }

                Column(
                    modifier = Modifier
                        .padding(innerPadding)
                        .fillMaxSize()
                ) {
                    LazyColumn(
                        modifier = Modifier.weight(1f),
                        contentPadding = PaddingValues(horizontal = 16.dp, vertical = 8.dp),
                        verticalArrangement = Arrangement.spacedBy(12.dp)
                    ) {
                        items(pack.stories, key = { it.id }) { story ->
                            val accent = accentForSeed(story.id)
                            val lengthCategory = StoryLengthCategory.fromPromptCount(story.prompts.size)
                            val lengthChipLabel = "${lengthCategory.emoji} ${lengthCategory.label}"
                            ChaosSurfaceCard(
                                modifier = Modifier.fillMaxWidth(),
                                accentColor = accent,
                                highlighted = rouletteStoryId == story.id,
                                onClick = { onOpenStory(story.id) }
                            ) {
                                Column(modifier = Modifier.padding(18.dp)) {
                                    Text(story.title, style = MaterialTheme.typography.titleLarge)
                                    Spacer(modifier = Modifier.height(6.dp))
                                    Text(
                                        text = "${story.prompts.size} prompts",
                                        style = MaterialTheme.typography.bodyMedium,
                                        color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.72f)
                                    )
                                    Spacer(modifier = Modifier.height(8.dp))
                                    FlowRow(
                                        horizontalArrangement = Arrangement.spacedBy(8.dp),
                                        verticalArrangement = Arrangement.spacedBy(8.dp)
                                    ) {
                                        StoryMetaChip(
                                            label = lengthChipLabel,
                                            accent = accent
                                        )
                                        story.tags.forEach { tag ->
                                            StoryMetaChip(
                                                label = formatStoryTag(tag),
                                                accent = accent
                                            )
                                        }
                                    }
                                }
                            }
                        }
                    }

                    GameFooterBar(
                        primaryLabel = if (rouletteRunning) "Rolling..." else "Pick Random",
                        onPrimary = {
                            if (rouletteRunning || pack.stories.isEmpty()) return@GameFooterBar
                            rouletteScope.launch {
                                rouletteRunning = true
                                repeat(8) { step ->
                                    rouletteStoryId = pack.stories.random().id
                                    delay(55L + (step * 18L))
                                }
                                val winner = pack.stories.random().id
                                rouletteStoryId = winner
                                delay(140L)
                                rouletteRunning = false
                                onOpenStory(winner)
                            }
                        },
                        onBack = onBack,
                        modifier = Modifier
                            .padding(horizontal = 16.dp, vertical = 12.dp)
                            .navigationBarsPadding()
                    )
                }
            }
        }
    }
}

private fun formatStoryTag(tag: String): String {
    return tag
        .trim()
        .split('-', '_', ' ')
        .filter { it.isNotBlank() }
        .joinToString(" ") { part ->
            part.lowercase().replaceFirstChar { char ->
                if (char.isLowerCase()) char.titlecase() else char.toString()
            }
        }
}

@Composable
private fun QuickChaosScreen(
    candidates: List<Pair<String, String>>,
    labelFor: (Pair<String, String>) -> String,
    onContinue: (String, String) -> Unit
) {
    val tokens = BadLibs.tokens
    val ink = tokens.ink
    val paper = tokens.paper
    val highlight = tokens.highlight

    var phase by remember(candidates) { mutableStateOf(QuickChaosPhase.Spinning) }
    var shuffleLabel by remember(candidates) {
        mutableStateOf(candidates.firstOrNull()?.let(labelFor).orEmpty())
    }

    LaunchedEffect(candidates) {
        phase = QuickChaosPhase.Spinning
        if (candidates.isEmpty()) return@LaunchedEffect
        val started = System.currentTimeMillis()
        val spinDurationMs = 1500L
        while (System.currentTimeMillis() - started < spinDurationMs) {
            shuffleLabel = labelFor(candidates.random())
            delay(85L)
        }
        val pick = candidates.random()
        shuffleLabel = labelFor(pick)
        phase = QuickChaosPhase.Locked
        delay(950L)
        onContinue(pick.first, pick.second)
    }

    val (packTitle, storyTitle) = remember(shuffleLabel) {
        val parts = shuffleLabel.split(": ", limit = 2)
        if (parts.size == 2) parts[0] to parts[1] else "" to shuffleLabel
    }

    val cardBg by animateColorAsState(
        targetValue = if (phase == QuickChaosPhase.Locked) highlight else paper,
        animationSpec = tween(durationMillis = 220, easing = MotionTokens.EaseOutEmphasized),
        label = "quick_chaos_card_bg"
    )
    val progressTarget = if (phase == QuickChaosPhase.Spinning) 0.55f else 1f
    val progress by animateFloatAsState(
        targetValue = progressTarget,
        animationSpec = tween(durationMillis = 320, easing = MotionTokens.EaseOutEmphasized),
        label = "quick_chaos_progress"
    )

    Box(
        modifier = Modifier
            .fillMaxSize()
            .background(paper)
            .padding(horizontal = 28.dp, vertical = 36.dp),
        contentAlignment = Alignment.Center
    ) {
        Column(
            modifier = Modifier.fillMaxSize(),
            verticalArrangement = Arrangement.SpaceBetween,
            horizontalAlignment = Alignment.CenterHorizontally
        ) {
            Column(horizontalAlignment = Alignment.CenterHorizontally) {
                Text(
                    text = "// quick chaos",
                    style = MaterialTheme.typography.labelLarge.copy(fontFamily = tokens.monoFamily),
                    color = ink
                )
                Spacer(modifier = Modifier.height(6.dp))
                Text(
                    text = if (phase == QuickChaosPhase.Spinning) "rolling a story…" else "locked in.",
                    style = MaterialTheme.typography.titleMedium,
                    color = ink.copy(alpha = 0.65f)
                )
            }

            Box(contentAlignment = Alignment.Center) {
                Box(
                    modifier = Modifier
                        .offset(x = 10.dp, y = 12.dp)
                        .size(width = 280.dp, height = 200.dp)
                        .background(highlight, RoundedCornerShape(28.dp))
                )
                Column(
                    modifier = Modifier
                        .size(width = 280.dp, height = 200.dp)
                        .background(cardBg, RoundedCornerShape(28.dp))
                        .border(2.dp, ink, RoundedCornerShape(28.dp))
                        .padding(horizontal = 22.dp, vertical = 24.dp),
                    verticalArrangement = Arrangement.Center,
                    horizontalAlignment = Alignment.CenterHorizontally
                ) {
                    Text(
                        text = if (packTitle.isNotEmpty()) packTitle.uppercase() else "RANDOM",
                        style = MaterialTheme.typography.labelMedium.copy(fontFamily = tokens.monoFamily),
                        color = ink.copy(alpha = 0.7f),
                        textAlign = TextAlign.Center
                    )
                    Spacer(modifier = Modifier.height(10.dp))
                    Text(
                        text = if (storyTitle.isNotEmpty()) storyTitle else "—",
                        style = MaterialTheme.typography.headlineSmall,
                        fontWeight = FontWeight.ExtraBold,
                        color = ink,
                        textAlign = TextAlign.Center,
                        maxLines = 3
                    )
                }
            }

            Column(
                modifier = Modifier.fillMaxWidth(),
                horizontalAlignment = Alignment.CenterHorizontally,
                verticalArrangement = Arrangement.spacedBy(10.dp)
            ) {
                LinearProgressIndicator(
                    progress = { progress },
                    modifier = Modifier.fillMaxWidth(),
                    color = ink,
                    trackColor = ink.copy(alpha = 0.12f)
                )
                Text(
                    text = if (phase == QuickChaosPhase.Spinning) "// holding the line" else "// launching",
                    style = MaterialTheme.typography.labelMedium.copy(fontFamily = tokens.monoFamily),
                    color = ink.copy(alpha = 0.7f)
                )
            }
        }
    }
}

@Composable
private fun StoryMetaChip(label: String, accent: Color = Color(0xFFA64D79)) {
    Box(
        modifier = Modifier
            .border(1.dp, accent.copy(alpha = 0.35f), RoundedCornerShape(50))
            .background(accent.copy(alpha = 0.12f), RoundedCornerShape(50))
            .padding(horizontal = 10.dp, vertical = 5.dp)
    ) {
        Text(text = label, style = MaterialTheme.typography.labelLarge)
    }
}

@OptIn(ExperimentalMaterial3Api::class, ExperimentalLayoutApi::class)
@Composable
private fun PromptInputScreen(
    story: Story?,
    snackbarHostState: SnackbarHostState,
    onBack: () -> Unit,
    onComplete: (Map<String, String>) -> Unit
) {
    val scope = rememberCoroutineScope()
    val focusManager = LocalFocusManager.current
    val keyboardController = LocalSoftwareKeyboardController.current
    val haptics = LocalHapticFeedback.current
    val isKeyboardOpen = WindowInsets.isImeVisible
    val inputMap = remember(story?.id) { mutableStateMapOf<String, String>() }
    val errorMap = remember(story?.id) { mutableStateMapOf<String, Boolean>() }
    var introVisible by remember(story?.id) { mutableStateOf(false) }
    var currentPromptIndex by remember(story?.id) { mutableStateOf(0) }

    LaunchedEffect(story?.id) {
        introVisible = false
        currentPromptIndex = 0
        if (story != null) {
            delay(90)
            introVisible = true
        }
    }

    ChaosScaffold(title = story?.title ?: "Prompts", onBack = onBack) { innerPadding ->
        if (story == null) {
            Box(
                modifier = Modifier
                    .padding(innerPadding)
                    .fillMaxSize(),
                contentAlignment = Alignment.Center
            ) {
                Text("Story not found.")
            }
            return@ChaosScaffold
        }

        val totalPrompts = story.prompts.size.coerceAtLeast(1)
        val filledPrompts = story.prompts.count { !inputMap[it.key].orEmpty().isBlank() }
        val progress = filledPrompts.toFloat() / totalPrompts.toFloat()
        val safeIndex = currentPromptIndex.coerceIn(0, story.prompts.lastIndex)
        val activePrompt = story.prompts[safeIndex]
        val isLastPrompt = safeIndex == story.prompts.lastIndex

        fun submitStory() {
            val missing = story.prompts.filter { inputMap[it.key].orEmpty().isBlank() }
            errorMap.clear()
            missing.forEach { errorMap[it.key] = true }

            if (missing.isNotEmpty()) {
                scope.launch {
                    snackbarHostState.showSnackbar("Missing value for: ${missing.first().label}")
                }
            } else {
                focusManager.clearFocus()
                onComplete(story.prompts.associate { it.key to inputMap[it.key].orEmpty().trim() })
            }
        }

        fun advancePrompt() {
            val value = inputMap[activePrompt.key].orEmpty().trim()
            if (value.isBlank()) {
                errorMap[activePrompt.key] = true
                scope.launch {
                    snackbarHostState.showSnackbar("Missing value for: ${activePrompt.label}")
                }
                return
            }

            errorMap[activePrompt.key] = false
            if (isLastPrompt) {
                haptics.performHapticFeedback(HapticFeedbackType.LongPress)
                submitStory()
            } else {
                haptics.performHapticFeedback(HapticFeedbackType.TextHandleMove)
                currentPromptIndex = (safeIndex + 1).coerceAtMost(story.prompts.lastIndex)
            }
        }

        Column(
            modifier = Modifier
                .padding(innerPadding)
                .imePadding()
                .fillMaxSize()
                .padding(horizontal = 16.dp, vertical = 8.dp),
            verticalArrangement = Arrangement.spacedBy(12.dp)
        ) {
            AnimatedVisibility(
                visible = introVisible,
                enter = fadeIn(
                    animationSpec = tween(
                        durationMillis = MotionTokens.DurationMediumMs,
                        easing = MotionTokens.EaseOutEmphasized
                    )
                ) + scaleIn(
                    initialScale = 0.985f,
                    animationSpec = tween(
                        durationMillis = MotionTokens.DurationMediumMs,
                        easing = MotionTokens.EaseOutEmphasized
                    )
                )
            ) {
                Text(
                    text = "// one at a time. keep it weird.",
                    style = MaterialTheme.typography.labelMedium,
                    color = BadLibs.tokens.ink.copy(alpha = 0.55f),
                    modifier = Modifier.padding(horizontal = 4.dp)
                )
            }

            LinearProgressIndicator(
                progress = { progress },
                modifier = Modifier
                    .fillMaxWidth()
                    .height(3.dp),
                color = BadLibs.tokens.ink,
                trackColor = BadLibs.tokens.ink.copy(alpha = 0.12f)
            )

            AnimatedContent(
                targetState = safeIndex,
                modifier = Modifier
                    .weight(1f)
                    .fillMaxWidth(),
                label = "prompt_single_step"
            ) { displayedIndex ->
                val prompt = story.prompts[displayedIndex]
                val promptHasError = errorMap[prompt.key] == true
                val displayedIsLast = displayedIndex == story.prompts.lastIndex
                val promptFocusRequester = remember(displayedIndex) { FocusRequester() }
                val tokens = BadLibs.tokens

                LaunchedEffect(displayedIndex) {
                    delay(120)
                    promptFocusRequester.requestFocus()
                    keyboardController?.show()
                }

                Column(
                    modifier = Modifier
                        .fillMaxSize()
                        .padding(top = 12.dp),
                    verticalArrangement = Arrangement.spacedBy(18.dp)
                ) {
                    Text(
                        text = "// prompt ${"%02d".format(displayedIndex + 1)} of ${"%02d".format(story.prompts.size)}",
                        style = MaterialTheme.typography.labelMedium,
                        color = tokens.ink.copy(alpha = 0.55f)
                    )
                    Text(
                        text = prompt.label,
                        style = MaterialTheme.typography.displaySmall,
                        color = tokens.ink,
                        lineHeight = 44.sp
                    )
                    OutlinedTextField(
                        value = inputMap[prompt.key].orEmpty(),
                        onValueChange = {
                            inputMap[prompt.key] = it
                            if (it.isNotBlank()) {
                                errorMap[prompt.key] = false
                            }
                        },
                        placeholder = {
                            Text(
                                "type something unhinged",
                                style = MaterialTheme.typography.titleMedium,
                                color = tokens.ink.copy(alpha = 0.35f)
                            )
                        },
                        modifier = Modifier
                            .fillMaxWidth()
                            .focusRequester(promptFocusRequester),
                        singleLine = true,
                        isError = promptHasError,
                        textStyle = MaterialTheme.typography.titleLarge.copy(
                            color = tokens.ink,
                            fontWeight = FontWeight.ExtraBold
                        ),
                        shape = RoundedCornerShape(14.dp),
                        keyboardOptions = KeyboardOptions(
                            imeAction = if (displayedIsLast) ImeAction.Done else ImeAction.Next
                        ),
                        keyboardActions = KeyboardActions(
                            onNext = { advancePrompt() },
                            onDone = { advancePrompt() }
                        ),
                        colors = TextFieldDefaults.colors(
                            focusedContainerColor = tokens.paper,
                            unfocusedContainerColor = tokens.paper,
                            focusedIndicatorColor = tokens.ink,
                            unfocusedIndicatorColor = tokens.ink.copy(alpha = 0.4f),
                            cursorColor = tokens.ink,
                            errorContainerColor = tokens.paper
                        )
                    )
                    if (promptHasError) {
                        Text(
                            text = "// needs a value",
                            style = MaterialTheme.typography.labelMedium,
                            color = MaterialTheme.colorScheme.error
                        )
                    }
                    Text(
                        text = if (displayedIsLast) "// final one. lock it in." else "// next: ${story.prompts.getOrNull(displayedIndex + 1)?.label?.lowercase() ?: ""}",
                        style = MaterialTheme.typography.labelMedium,
                        color = tokens.ink.copy(alpha = 0.55f)
                    )
                }
            }

            if (!isKeyboardOpen) {
                Row(
                    modifier = Modifier
                        .fillMaxWidth()
                        .navigationBarsPadding(),
                    horizontalArrangement = Arrangement.spacedBy(12.dp),
                    verticalAlignment = Alignment.CenterVertically
                ) {
                    OutlinedButton(
                        onClick = {
                            focusManager.clearFocus()
                            if (safeIndex > 0) currentPromptIndex = safeIndex - 1
                        },
                        enabled = safeIndex > 0,
                        shape = RoundedCornerShape(14.dp),
                        border = BorderStroke(1.dp, BadLibs.tokens.ink.copy(alpha = 0.45f)),
                        colors = ButtonDefaults.outlinedButtonColors(contentColor = BadLibs.tokens.ink),
                        modifier = Modifier.weight(0.85f)
                    ) {
                        Text(
                            "← back",
                            style = MaterialTheme.typography.labelLarge
                        )
                    }

                    ChaosPrimaryButton(
                        label = if (isLastPrompt) "reveal story →" else "next →",
                        onClick = { advancePrompt() },
                        modifier = Modifier.weight(1.15f)
                    )
                }
            }
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
private fun ResultScreen(
    completedStory: CompletedStory?,
    onRemix: () -> Unit,
    onPlayAgain: () -> Unit,
    onBackHome: () -> Unit
) {
    val context = LocalContext.current
    val haptics = LocalHapticFeedback.current
    val tokens = BadLibs.tokens
    val palette = paletteForPack(completedStory?.packId.orEmpty(), available = true)
    val backgroundColor = softPackTint(palette, strength = 0.16f)
    val onBackgroundColor = onColorFor(backgroundColor)
    // Action surface contrasts with background: dark bg gets Highlight, light bg gets Ink.
    val isDarkBackground = onBackgroundColor == tokens.paper
    val actionBackground = if (isDarkBackground) tokens.highlight else tokens.ink
    val onActionBackground = onColorFor(actionBackground)
    // Revealed words wear the pack's saturated color as a typographic accent.
    val wordAccent = palette.primary

    var revealPhase by remember(completedStory?.storyId) { mutableStateOf(RevealPhase.Anticipation) }
    var revealedFilled by remember(completedStory?.storyId) { mutableStateOf(0) }
    var anticipationIndex by remember(completedStory?.storyId) { mutableStateOf(0) }

    val filledTotal = completedStory?.tokens?.count { it is RenderedToken.Filled } ?: 0
    val anticipationPhrases = remember {
        listOf(
            "loading the bit",
            "consulting the group chat",
            "asking your ex",
            "rendering nonsense",
            "running with scissors"
        )
    }

    Box(
        modifier = Modifier
            .fillMaxSize()
            .background(backgroundColor)
            .windowInsetsPadding(WindowInsets.statusBars.only(WindowInsetsSides.Top))
    ) {
        if (completedStory == null) {
            Box(modifier = Modifier.fillMaxSize(), contentAlignment = Alignment.Center) {
                Text("No completed story yet.", color = onBackgroundColor)
            }
            return@Box
        }

        LaunchedEffect(completedStory.storyId) {
            revealPhase = RevealPhase.Anticipation
            revealedFilled = 0
            anticipationIndex = 0
            haptics.performHapticFeedback(HapticFeedbackType.LongPress)
            repeat(4) {
                delay(180)
                anticipationIndex = (anticipationIndex + 1) % anticipationPhrases.size
                haptics.performHapticFeedback(HapticFeedbackType.TextHandleMove)
            }
            delay(120)
            revealPhase = RevealPhase.Reveal
            haptics.performHapticFeedback(HapticFeedbackType.LongPress)
            delay(420)
            while (revealedFilled < filledTotal) {
                delay(220)
                revealedFilled += 1
                haptics.performHapticFeedback(HapticFeedbackType.TextHandleMove)
            }
            delay(380)
            revealPhase = RevealPhase.Rest
            haptics.performHapticFeedback(HapticFeedbackType.LongPress)
        }

        Column(
            modifier = Modifier
                .fillMaxSize()
                .verticalScroll(rememberScrollState())
                .padding(horizontal = 24.dp)
                .padding(top = 24.dp, bottom = 16.dp)
                .navigationBarsPadding(),
            verticalArrangement = Arrangement.spacedBy(20.dp)
        ) {
            // Top mono label — rotates during anticipation, sticks during reveal
            AnimatedContent(
                targetState = when (revealPhase) {
                    RevealPhase.Anticipation -> "// " + anticipationPhrases[anticipationIndex]
                    RevealPhase.Reveal -> "// reveal in progress"
                    else -> "// the masterpiece"
                },
                label = "reveal_mono_label"
            ) { label ->
                Text(
                    text = label.uppercase(),
                    style = MaterialTheme.typography.labelMedium,
                    color = onBackgroundColor.copy(alpha = 0.62f)
                )
            }

            // Title — drops in on Reveal
            AnimatedVisibility(
                visible = revealPhase != RevealPhase.Anticipation,
                enter = slideInVertically(
                    initialOffsetY = { -it / 2 },
                    animationSpec = tween(420, easing = FastOutSlowInEasing)
                ) + fadeIn(animationSpec = tween(300))
            ) {
                Text(
                    text = completedStory.storyTitle,
                    style = MaterialTheme.typography.displayMedium,
                    color = onBackgroundColor,
                    fontFamily = tokens.displayFamily
                )
            }

            // Story body — shows blanks then progressively fills
            if (revealPhase != RevealPhase.Anticipation) {
                val annotated = remember(completedStory.storyId, revealedFilled) {
                    buildRevealedStory(
                        storyTokens = completedStory.tokens,
                        revealedCount = revealedFilled,
                        onBackground = onBackgroundColor,
                        accent = wordAccent,
                        monoFamily = tokens.monoFamily,
                        displayFamily = tokens.displayFamily
                    )
                }
                Text(
                    text = annotated,
                    style = MaterialTheme.typography.bodyLarge.copy(
                        color = onBackgroundColor,
                        fontSize = 19.sp,
                        lineHeight = 30.sp
                    ),
                    modifier = Modifier.fillMaxWidth()
                )
            } else {
                // Anticipation hero — big mono dots so the screen isn't empty
                Spacer(modifier = Modifier.height(12.dp))
                Text(
                    text = "• • •",
                    style = MaterialTheme.typography.displayLarge,
                    color = onBackgroundColor.copy(alpha = 0.55f)
                )
            }

            Spacer(modifier = Modifier.height(8.dp))

            // Actions — fade in at Rest
            AnimatedVisibility(
                visible = revealPhase == RevealPhase.Rest,
                enter = fadeIn(tween(360)) + slideInVertically(
                    initialOffsetY = { it / 4 },
                    animationSpec = tween(360, easing = FastOutSlowInEasing)
                )
            ) {
                Column(verticalArrangement = Arrangement.spacedBy(12.dp)) {
                    RevealPrimaryAction(
                        label = "Send to the group chat",
                        sublabel = "↗ share this disaster",
                        background = actionBackground,
                        onLabel = onActionBackground,
                        onClick = {
                            haptics.performHapticFeedback(HapticFeedbackType.LongPress)
                            val shareIntent = Intent(Intent.ACTION_SEND).apply {
                                type = "text/plain"
                                putExtra(Intent.EXTRA_TEXT, completedStory.renderedText)
                            }
                            context.startActivity(
                                Intent.createChooser(shareIntent, "Share your Bad Libs story")
                            )
                        },
                        modifier = Modifier.fillMaxWidth()
                    )

                    Row(
                        modifier = Modifier.fillMaxWidth(),
                        horizontalArrangement = Arrangement.spacedBy(12.dp)
                    ) {
                        RevealMonoPill(
                            label = "remix words",
                            onLabel = onBackgroundColor,
                            onClick = {
                                haptics.performHapticFeedback(HapticFeedbackType.LongPress)
                                onRemix()
                            },
                            modifier = Modifier.weight(1f)
                        )
                        RevealMonoPill(
                            label = "next story",
                            onLabel = onBackgroundColor,
                            onClick = {
                                haptics.performHapticFeedback(HapticFeedbackType.LongPress)
                                onPlayAgain()
                            },
                            modifier = Modifier.weight(1f)
                        )
                    }

                    Box(modifier = Modifier.fillMaxWidth(), contentAlignment = Alignment.Center) {
                        TextButton(onClick = onBackHome) {
                            Text(
                                "← home",
                                style = MaterialTheme.typography.labelMedium,
                                color = onBackgroundColor.copy(alpha = 0.7f)
                            )
                        }
                    }
                }
            }
        }
    }
}

private fun buildRevealedStory(
    storyTokens: List<RenderedToken>,
    revealedCount: Int,
    onBackground: Color,
    accent: Color,
    monoFamily: androidx.compose.ui.text.font.FontFamily,
    displayFamily: androidx.compose.ui.text.font.FontFamily
) = buildAnnotatedString {
    var filledSeen = 0
    storyTokens.forEach { token ->
        when (token) {
            is RenderedToken.Static -> {
                withStyle(SpanStyle(color = onBackground)) {
                    append(token.text)
                }
            }
            is RenderedToken.Filled -> {
                val isRevealed = filledSeen < revealedCount
                filledSeen += 1
                if (isRevealed) {
                    withStyle(
                        SpanStyle(
                            color = accent,
                            fontWeight = FontWeight.ExtraBold,
                            fontFamily = displayFamily,
                            textDecoration = TextDecoration.Underline
                        )
                    ) {
                        append(token.value)
                    }
                } else {
                    withStyle(
                        SpanStyle(
                            color = onBackground.copy(alpha = 0.32f),
                            fontFamily = monoFamily
                        )
                    ) {
                        append("_".repeat((token.value.length).coerceAtLeast(4)))
                    }
                }
            }
        }
    }
}

// Legacy ChaosScaffold no longer used by ResultScreen, kept for unrelated screens.
@Composable
private fun RevealPrimaryAction(
    label: String,
    sublabel: String,
    background: Color,
    onLabel: Color,
    onClick: () -> Unit,
    modifier: Modifier = Modifier
) {
    androidx.compose.material3.Surface(
        onClick = onClick,
        modifier = modifier,
        shape = RoundedCornerShape(20.dp),
        color = background
    ) {
        Column(
            modifier = Modifier.padding(horizontal = 18.dp, vertical = 16.dp),
        ) {
            Text(
                text = label,
                style = MaterialTheme.typography.titleLarge,
                color = onLabel,
                fontWeight = FontWeight.ExtraBold
            )
            Spacer(modifier = Modifier.height(2.dp))
            Text(
                text = sublabel,
                style = MaterialTheme.typography.labelMedium,
                color = onLabel.copy(alpha = 0.7f)
            )
        }
    }
}

@Composable
private fun RevealMonoPill(
    label: String,
    onLabel: Color,
    onClick: () -> Unit,
    modifier: Modifier = Modifier
) {
    androidx.compose.material3.Surface(
        onClick = onClick,
        modifier = modifier,
        shape = RoundedCornerShape(14.dp),
        color = Color.Transparent,
        border = BorderStroke(1.dp, onLabel.copy(alpha = 0.45f))
    ) {
        Box(
            modifier = Modifier.padding(vertical = 14.dp),
            contentAlignment = Alignment.Center
        ) {
            Text(
                text = label,
                style = MaterialTheme.typography.labelLarge,
                color = onLabel
            )
        }
    }
}

