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
import com.butterfeetlabs.badlibs.data.CompletedStory
import com.butterfeetlabs.badlibs.data.Story
import com.butterfeetlabs.badlibs.data.StoryLengthCategory
import com.butterfeetlabs.badlibs.data.StoryPack
import com.butterfeetlabs.badlibs.data.StoryRepository
import com.butterfeetlabs.badlibs.ui.theme.BadLibsTheme
import com.butterfeetlabs.badlibs.ui.theme.ChaosOrange
import com.butterfeetlabs.badlibs.ui.theme.MotionTokens
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

private fun accentForSeed(seed: String): Color {
    if (seed.isEmpty()) return ChaosOrange
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

    fun completeStory(story: Story, values: Map<String, String>) {
        val rendered = values.entries.fold(story.template) { acc, entry ->
            acc.replace("{${entry.key}}", entry.value)
        }
        completedStory = CompletedStory(
            storyId = story.id,
            storyTitle = story.title,
            values = values,
            renderedText = rendered
        )
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
        viewModel.packs.flatMap { pack -> pack.stories.map { story -> pack.id to story.id } }
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
                    PackListScreen(
                        isLoading = viewModel.isLoading,
                        error = viewModel.loadError,
                        packs = viewModel.packs,
                        onBack = { navController.popBackStack() },
                        onRetry = { viewModel.loadContent() },
                        onOpenPack = { packId -> navController.navigate("stories/$packId") }
                    )
                }
                composable(
                    route = Screen.Stories.route,
                    arguments = listOf(navArgument("packId") { type = NavType.StringType })
                ) { backStackEntry ->
                    val packId = backStackEntry.arguments?.getString("packId").orEmpty()
                    val pack = viewModel.findPack(packId)
                    StoryListScreen(
                        pack = pack,
                        onBack = { navController.popBackStack() },
                        onOpenStory = { storyId -> navController.navigate("prompts/$packId/$storyId") }
                    )
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
                                viewModel.completeStory(story, values)
                                navController.navigate(Screen.Result.route)
                            }
                        }
                    )
                }
                composable(Screen.Result.route) {
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

            SnackbarHost(
                hostState = snackbarHostState,
                modifier = Modifier.align(Alignment.BottomCenter)
            )
        }
    }
}

@Composable
private fun ChaosBackdrop(content: @Composable BoxScope.() -> Unit) {
    val transition = rememberInfiniteTransition(label = "chaos_backdrop")
    val orb1X = transition.animateFloat(
        initialValue = -120f,
        targetValue = 90f,
        animationSpec = infiniteRepeatable(
            animation = tween(8000, easing = LinearEasing),
            repeatMode = RepeatMode.Reverse
        ),
        label = "orb1x"
    )
    val orb1Y = transition.animateFloat(
        initialValue = 20f,
        targetValue = 220f,
        animationSpec = infiniteRepeatable(
            animation = tween(12000, easing = LinearEasing),
            repeatMode = RepeatMode.Reverse
        ),
        label = "orb1y"
    )
    val orb2X = transition.animateFloat(
        initialValue = 240f,
        targetValue = 20f,
        animationSpec = infiniteRepeatable(
            animation = tween(10000, easing = LinearEasing),
            repeatMode = RepeatMode.Reverse
        ),
        label = "orb2x"
    )
    val orb2Y = transition.animateFloat(
        initialValue = 420f,
        targetValue = 180f,
        animationSpec = infiniteRepeatable(
            animation = tween(14000, easing = LinearEasing),
            repeatMode = RepeatMode.Reverse
        ),
        label = "orb2y"
    )

    Box(
        modifier = Modifier
            .fillMaxSize()
            .background(
                Brush.verticalGradient(
                    colors = listOf(
                        Color(0xFFFFE5C7),
                        Color(0xFFFFC7A8),
                        Color(0xFFFFF2DE)
                    )
                )
            )
    ) {
        Box(
            modifier = Modifier
                .size(220.dp)
                .graphicsLayer {
                    translationX = orb1X.value
                    translationY = orb1Y.value
                    alpha = 0.22f
                }
                .background(Color(0x55FF9B6A), CircleShape)
        )
        Box(
            modifier = Modifier
                .size(180.dp)
                .graphicsLayer {
                    translationX = orb2X.value
                    translationY = orb2Y.value
                    alpha = 0.18f
                }
                .background(Color(0x55F06A99), CircleShape)
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
                    Text(text = title, style = MaterialTheme.typography.titleLarge)
                },
                navigationIcon = {
                    if (onBack != null) {
                        TextButton(onClick = onBack) { Text("Back") }
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
private fun ChaosPrimaryButton(
    label: String,
    onClick: () -> Unit,
    modifier: Modifier = Modifier
) {
    Button(
        onClick = onClick,
        modifier = modifier,
        colors = ButtonDefaults.buttonColors(containerColor = ChaosOrange),
        shape = RoundedCornerShape(16.dp)
    ) {
        Text(text = label, fontWeight = FontWeight.Bold)
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
    ChaosScaffold(title = "Bad Libs") { innerPadding ->
        LazyColumn(
            modifier = Modifier
                .padding(innerPadding)
                .fillMaxSize(),
            contentPadding = PaddingValues(horizontal = 18.dp, vertical = 10.dp),
            verticalArrangement = Arrangement.spacedBy(14.dp)
        ) {
            item {
                Column(
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(horizontal = 2.dp, vertical = 6.dp),
                    verticalArrangement = Arrangement.spacedBy(10.dp)
                ) {
                    Text(
                        text = "Make stupid stories with serious confidence.",
                        style = MaterialTheme.typography.headlineLarge,
                        lineHeight = 40.sp,
                        fontWeight = FontWeight.ExtraBold
                    )
                    Text(
                        text = "Pick a pack, add cursed words, and reveal your literary disaster.",
                        style = MaterialTheme.typography.bodyLarge,
                        color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.8f)
                    )
                }
            }
            item {
                androidx.compose.material3.Surface(
                    onClick = onStart,
                    modifier = Modifier.fillMaxWidth(),
                    shape = RoundedCornerShape(24.dp),
                    color = Color.Transparent,
                    border = BorderStroke(1.dp, Color(0xFFE5533D).copy(alpha = 0.6f))
                ) {
                    Row(
                        modifier = Modifier
                            .background(
                                Brush.linearGradient(
                                    colors = listOf(
                                        Color(0xFFFFB06F).copy(alpha = 0.4f),
                                        Color(0xFF8E2E63).copy(alpha = 0.28f)
                                    )
                                )
                            )
                            .padding(horizontal = 18.dp, vertical = 16.dp),
                        horizontalArrangement = Arrangement.SpaceBetween,
                        verticalAlignment = Alignment.CenterVertically
                    ) {
                        Row(
                            modifier = Modifier.weight(1f),
                            horizontalArrangement = Arrangement.spacedBy(12.dp),
                            verticalAlignment = Alignment.CenterVertically
                        ) {
                            Box(
                                modifier = Modifier
                                    .size(44.dp)
                                    .background(Color(0xFF8E2E63), CircleShape),
                                contentAlignment = Alignment.Center
                            ) {
                                Text("⚡", fontSize = 23.sp, color = Color.White)
                            }
                            Column {
                                Text(
                                    text = "Start The Chaos",
                                    style = MaterialTheme.typography.titleMedium,
                                    fontWeight = FontWeight.Bold
                                )
                                Spacer(modifier = Modifier.height(3.dp))
                                Text(
                                    text = "Jump in fast and build your cursed masterpiece.",
                                    style = MaterialTheme.typography.bodyMedium,
                                    color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.75f)
                                )
                            }
                        }
                        Text(
                            text = "→",
                            style = MaterialTheme.typography.titleLarge,
                            color = Color(0xFF8E2E63),
                            fontWeight = FontWeight.Bold
                        )
                    }
                }
            }
            item {
                Column(
                    modifier = Modifier.padding(horizontal = 4.dp),
                    verticalArrangement = Arrangement.spacedBy(2.dp)
                ) {
                    Text(
                        text = "How it works",
                        style = MaterialTheme.typography.titleMedium,
                        fontWeight = FontWeight.Bold
                    )
                    Text(
                        text = "Quick guide before you start",
                        style = MaterialTheme.typography.bodySmall,
                        color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.7f)
                    )
                }
            }
            item {
                IntroGuideTile(
                    emoji = "1",
                    title = "Pick a pack",
                    accent = Color(0xFF6C63FF),
                    copy = "Choose your flavor of nonsense and lock in the vibe."
                )
            }
            item {
                IntroGuideTile(
                    emoji = "2",
                    title = "Fill prompts",
                    accent = Color(0xFF8E2E63),
                    copy = "Use weird words. Regret nothing. Spell-check optional."
                )
            }
            item {
                IntroGuideTile(
                    emoji = "3",
                    title = "Reveal and share",
                    accent = Color(0xFFE5533D),
                    copy = "Drop your masterpiece in group chat and cause mild panic."
                )
            }
        }
    }
}

@Composable
private fun IntroGuideTile(
    emoji: String,
    title: String,
    copy: String,
    accent: Color
) {
    Box(
        modifier = Modifier
            .fillMaxWidth()
            .border(1.dp, Color.Black.copy(alpha = 0.08f), RoundedCornerShape(18.dp))
            .background(
                Color.White.copy(alpha = 0.58f),
                RoundedCornerShape(18.dp)
            )
            .padding(horizontal = 14.dp, vertical = 13.dp)
    ) {
        Row(
            horizontalArrangement = Arrangement.spacedBy(14.dp),
            verticalAlignment = Alignment.CenterVertically
        ) {
            Box(
                modifier = Modifier
                    .size(30.dp)
                    .border(1.dp, accent.copy(alpha = 0.45f), CircleShape)
                    .background(accent.copy(alpha = 0.08f), CircleShape),
                contentAlignment = Alignment.Center
            ) {
                Text(
                    text = emoji,
                    color = accent,
                    fontWeight = FontWeight.Bold,
                    style = MaterialTheme.typography.labelLarge
                )
            }
            Column {
                Text(text = title, style = MaterialTheme.typography.titleSmall, fontWeight = FontWeight.SemiBold)
                Spacer(modifier = Modifier.height(3.dp))
                Text(
                    text = copy,
                    style = MaterialTheme.typography.bodyMedium,
                    color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.74f)
                )
            }
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

    ChaosScaffold(title = "Bad Libs") { innerPadding ->
        Column(
            modifier = Modifier
                .padding(innerPadding)
                .fillMaxSize()
                .padding(horizontal = 18.dp, vertical = 10.dp),
            verticalArrangement = Arrangement.spacedBy(14.dp)
        ) {
            Box(
                modifier = Modifier
                    .fillMaxWidth()
                    .padding(horizontal = 2.dp, vertical = 6.dp)
            ) {
                Column(verticalArrangement = Arrangement.spacedBy(10.dp)) {
                    Text(
                        text = "Ready to wreck another perfectly good story?",
                        style = MaterialTheme.typography.headlineLarge,
                        fontWeight = FontWeight.ExtraBold,
                        lineHeight = 38.sp
                    )
                }
            }

            if (lastStoryLabel != null) {
                HomeActionTile(
                    title = "Continue Last Story",
                    subtitle = lastStoryLabel,
                    emoji = "↺",
                    accent = Color(0xFF8E2E63),
                    emojiSizeSp = 24,
                    emojiOffsetY = (-3).dp,
                    onClick = onContinueLast
                )
            }

            if (dailyChallengeLabel != null) {
                HomeActionTile(
                    title = "Daily Challenge",
                    subtitle = dailyChallengeLabel,
                    emoji = "☀",
                    accent = Color(0xFFFF7B54),
                    iconCircleColor = Color(0xFF5A2A14),
                    onClick = onPlayDailyChallenge
                )
            }

            HomeActionTile(
                title = "Browse Packs",
                subtitle = if (packs.isEmpty()) "No packs loaded yet" else "Pick a flavor and start causing damage.",
                emoji = "📚",
                accent = Color(0xFF6C63FF),
                onClick = onPlayPacks
            )

            HomeActionTile(
                title = if (canQuickPlay) "Quick Chaos" else "Quick Chaos",
                subtitle = if (canQuickPlay) "Random story, dramatic entrance." else "No stories loaded yet.",
                emoji = "🎲",
                accent = Color(0xFFFFB06F),
                enabled = canQuickPlay,
                onClick = onQuickChaos
            )

            if (!canQuickPlay) {
                Text(
                    text = "Add more story packs to unlock Quick Chaos.",
                    style = MaterialTheme.typography.bodyMedium,
                    color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.7f),
                    modifier = Modifier.padding(horizontal = 4.dp)
                )
            }

            Spacer(modifier = Modifier.weight(1f))

            Box(modifier = Modifier.fillMaxWidth(), contentAlignment = Alignment.Center) {
                TextButton(onClick = onShowIntro) {
                    Text("Replay Intro")
                }
            }
        }
    }
}

@Composable
private fun HomeActionTile(
    title: String,
    subtitle: String,
    emoji: String,
    accent: Color,
    onClick: () -> Unit,
    enabled: Boolean = true,
    iconCircleColor: Color = accent,
    emojiSizeSp: Int = 22,
    emojiOffsetY: Dp = 0.dp
) {
    val backgroundTint = if (enabled) accent.copy(alpha = 0.18f) else Color.Black.copy(alpha = 0.05f)
    val borderTint = if (enabled) accent.copy(alpha = 0.45f) else Color.Black.copy(alpha = 0.12f)
    val titleColor = if (enabled) MaterialTheme.colorScheme.onBackground else MaterialTheme.colorScheme.onBackground.copy(alpha = 0.45f)
    val subtitleColor = if (enabled) MaterialTheme.colorScheme.onBackground.copy(alpha = 0.78f) else MaterialTheme.colorScheme.onBackground.copy(alpha = 0.35f)

    androidx.compose.material3.Surface(
        onClick = onClick,
        enabled = enabled,
        modifier = Modifier.fillMaxWidth(),
        shape = RoundedCornerShape(24.dp),
        color = Color.Transparent,
        border = BorderStroke(1.dp, borderTint)
    ) {
        Box(
            modifier = Modifier
                .background(
                    Brush.linearGradient(
                        colors = listOf(
                            backgroundTint,
                            backgroundTint.copy(alpha = if (enabled) 0.26f else 0.10f)
                        )
                    )
                )
        ) {
            Row(
                modifier = Modifier.padding(horizontal = 16.dp, vertical = 16.dp),
                horizontalArrangement = Arrangement.spacedBy(14.dp),
                verticalAlignment = Alignment.CenterVertically
            ) {
                Box(
                    modifier = Modifier
                        .size(44.dp)
                        .background(iconCircleColor.copy(alpha = if (enabled) 0.95f else 0.35f), CircleShape),
                    contentAlignment = Alignment.Center
                ) {
                    Text(
                        text = emoji,
                        style = MaterialTheme.typography.titleMedium,
                        fontSize = emojiSizeSp.sp,
                        lineHeight = emojiSizeSp.sp,
                        color = Color.White,
                        fontWeight = FontWeight.Bold,
                        textAlign = TextAlign.Center,
                        modifier = Modifier
                            .offset(y = emojiOffsetY)
                            .fillMaxWidth()
                    )
                }

                Column(modifier = Modifier.weight(1f)) {
                    Text(
                        text = title,
                        style = MaterialTheme.typography.titleMedium,
                        fontWeight = FontWeight.Bold,
                        color = titleColor
                    )
                    Spacer(modifier = Modifier.height(4.dp))
                    Text(
                        text = subtitle,
                        style = MaterialTheme.typography.bodyMedium,
                        color = subtitleColor,
                        lineHeight = 20.sp
                    )
                }

                Text(
                    text = "→",
                    style = MaterialTheme.typography.titleLarge,
                    color = if (enabled) accent else subtitleColor,
                    fontWeight = FontWeight.Bold
                )
            }
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

@OptIn(ExperimentalMaterial3Api::class)
@Composable
private fun PackListScreen(
    isLoading: Boolean,
    error: String?,
    packs: List<StoryPack>,
    onBack: () -> Unit,
    onRetry: () -> Unit,
    onOpenPack: (String) -> Unit
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
                            val accent = accentForSeed(pack.id)
                            ChaosSurfaceCard(
                                modifier = Modifier.fillMaxWidth(),
                                accentColor = accent,
                                highlighted = roulettePackId == pack.id,
                                onClick = { onOpenPack(pack.id) }
                            ) {
                                Column(modifier = Modifier.padding(18.dp)) {
                                    Row(
                                        modifier = Modifier.fillMaxWidth(),
                                        horizontalArrangement = Arrangement.SpaceBetween,
                                        verticalAlignment = Alignment.CenterVertically
                                    ) {
                                        Text(
                                            text = pack.title,
                                            style = MaterialTheme.typography.titleLarge,
                                            modifier = Modifier.weight(1f)
                                        )
                                        PackCountPill(count = pack.stories.size, accent = accent)
                                    }
                                    Spacer(modifier = Modifier.height(8.dp))
                                    Text(pack.description, style = MaterialTheme.typography.bodyMedium)
                                    Spacer(modifier = Modifier.height(10.dp))
                                    StoryMetaChip(
                                        label = "Rated ${pack.rating.uppercase()}",
                                        accent = accent
                                    )
                                }
                            }
                        }
                    }

                    GameFooterBar(
                        primaryLabel = if (rouletteRunning) "Rolling..." else "Pick Random",
                        onPrimary = {
                            if (rouletteRunning || packs.isEmpty()) return@GameFooterBar
                            rouletteScope.launch {
                                rouletteRunning = true
                                repeat(8) { step ->
                                    roulettePackId = packs.random().id
                                    delay(55L + (step * 18L))
                                }
                                val winner = packs.random().id
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
private fun PackCountPill(count: Int, accent: Color = Color(0xFFFFB06F)) {
    Box(
        modifier = Modifier
            .background(accent.copy(alpha = 0.22f), RoundedCornerShape(50))
            .padding(horizontal = 12.dp, vertical = 6.dp)
    ) {
        Text(
            text = "$count stories",
            style = MaterialTheme.typography.labelLarge,
            color = MaterialTheme.colorScheme.onBackground
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
                            val lengthChipLabel = lengthChipLabelForPack(
                                category = lengthCategory,
                                packRating = pack.rating
                            )
                            ChaosSurfaceCard(
                                modifier = Modifier.fillMaxWidth(),
                                accentColor = accent,
                                highlighted = rouletteStoryId == story.id,
                                onClick = { onOpenStory(story.id) }
                            ) {
                                Column(modifier = Modifier.padding(18.dp)) {
                                    Text(story.title, style = MaterialTheme.typography.titleLarge)
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

private fun lengthChipLabelForPack(category: StoryLengthCategory, packRating: String): String {
    return if (packRating.equals("kids", ignoreCase = true) && category == StoryLengthCategory.BrainDamage) {
        "🐢 Long"
    } else {
        "${category.emoji} ${category.label}"
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
    val transition = rememberInfiniteTransition(label = "quick_chaos_screen")
    val driftX by transition.animateFloat(
        initialValue = -18f,
        targetValue = 18f,
        animationSpec = infiniteRepeatable(
            animation = tween(3400, easing = LinearEasing),
            repeatMode = RepeatMode.Reverse
        ),
        label = "quick_chaos_drift_x"
    )
    val driftY by transition.animateFloat(
        initialValue = 12f,
        targetValue = -12f,
        animationSpec = infiniteRepeatable(
            animation = tween(2800, easing = LinearEasing),
            repeatMode = RepeatMode.Reverse
        ),
        label = "quick_chaos_drift_y"
    )
    val spin by transition.animateFloat(
        initialValue = 0f,
        targetValue = 360f,
        animationSpec = infiniteRepeatable(
            animation = tween(4200, easing = LinearEasing)
        ),
        label = "quick_chaos_spin"
    )
    var phase by remember(candidates) { mutableStateOf(QuickChaosPhase.Spinning) }
    var selectedStory by remember(candidates) { mutableStateOf<Pair<String, String>?>(null) }
    var selectedLabel by remember(candidates) { mutableStateOf("Random story") }

    LaunchedEffect(candidates) {
        phase = QuickChaosPhase.Spinning
        selectedStory = null
        selectedLabel = "Random story"
        if (candidates.isEmpty()) {
            return@LaunchedEffect
        }
        delay(650L)
        selectedStory = candidates.random()
        selectedLabel = selectedStory?.let(labelFor) ?: "Random story"
        phase = QuickChaosPhase.Locked
        delay(1000L)
        selectedStory?.let { (packId, storyId) ->
            onContinue(packId, storyId)
        }
    }

    Box(
        modifier = Modifier
            .fillMaxSize()
            .background(
                Brush.verticalGradient(
                    colors = listOf(
                        Color(0xFFFFD7A8),
                        Color(0xFFFF9E5A),
                        Color(0xFF8E2E63)
                    )
                )
            )
    ) {
        Box(
            modifier = Modifier
                .size(240.dp)
                .graphicsLayer {
                    translationX = driftX * 1.2f
                    translationY = driftY
                    alpha = 0.18f
                }
                .background(Color.White.copy(alpha = 0.35f), CircleShape)
        )
        Box(
            modifier = Modifier
                .size(160.dp)
                .graphicsLayer {
                    translationX = -driftX
                    translationY = driftY * 0.8f
                    alpha = 0.22f
                }
                .background(Color(0x55F06A99), CircleShape)
        )

        Column(
            modifier = Modifier
                .fillMaxSize()
                .padding(horizontal = 24.dp, vertical = 28.dp),
            verticalArrangement = Arrangement.SpaceBetween,
            horizontalAlignment = Alignment.CenterHorizontally
        ) {
            Text(
                text = "Quick Chaos",
                style = MaterialTheme.typography.labelLarge,
                fontWeight = FontWeight.Bold,
                color = Color.White
            )

            Column(
                horizontalAlignment = Alignment.CenterHorizontally,
                verticalArrangement = Arrangement.spacedBy(16.dp)
            ) {
                Box(
                    modifier = Modifier
                        .size(184.dp)
                        .graphicsLayer {
                            rotationZ = spin
                        }
                        .background(Color.White.copy(alpha = 0.16f), CircleShape),
                    contentAlignment = Alignment.Center
                ) {
                    Box(
                        modifier = Modifier
                            .size(118.dp)
                            .background(Color.White.copy(alpha = 0.22f), CircleShape),
                        contentAlignment = Alignment.Center
                    ) {
                        Text(
                            text = "🎯",
                            style = MaterialTheme.typography.displayLarge
                        )
                    }
                }

                AnimatedContent(
                    targetState = phase,
                    label = "quick_chaos_phase"
                ) { currentPhase ->
                    Column(horizontalAlignment = Alignment.CenterHorizontally) {
                        Text(
                            text = when (currentPhase) {
                                QuickChaosPhase.Spinning -> "Landing the theme..."
                                QuickChaosPhase.Locked -> "Theme locked"
                            },
                            style = MaterialTheme.typography.headlineMedium,
                            color = Color.White,
                            fontWeight = FontWeight.Bold,
                            textAlign = TextAlign.Center
                        )
                        Spacer(modifier = Modifier.height(8.dp))
                        Text(
                            text = when (currentPhase) {
                                QuickChaosPhase.Spinning -> "Let the nonsense settle for a second."
                                QuickChaosPhase.Locked -> selectedLabel
                            },
                            style = MaterialTheme.typography.titleMedium,
                            color = Color.White.copy(alpha = 0.92f),
                            textAlign = TextAlign.Center
                        )
                    }
                }
            }

            Column(horizontalAlignment = Alignment.CenterHorizontally) {
                LinearProgressIndicator(
                    progress = {
                        if (phase == QuickChaosPhase.Spinning) 0.6f else 1f
                    },
                    modifier = Modifier.fillMaxWidth(),
                    color = Color.White,
                    trackColor = Color.White.copy(alpha = 0.22f)
                )
                Spacer(modifier = Modifier.height(10.dp))
                Text(
                    text = when (phase) {
                        QuickChaosPhase.Spinning -> "Holding the line..."
                        QuickChaosPhase.Locked -> "Launching in a beat"
                    },
                    style = MaterialTheme.typography.labelLarge,
                    color = Color.White.copy(alpha = 0.9f)
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
                submitStory()
            } else {
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
                    text = "One prompt at a time. Keep it weird.",
                    style = MaterialTheme.typography.titleSmall,
                    fontWeight = FontWeight.SemiBold,
                    color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.82f),
                    modifier = Modifier.padding(horizontal = 4.dp)
                )
            }

            LinearProgressIndicator(
                progress = { progress },
                modifier = Modifier.fillMaxWidth(),
                color = Color(0xFF8E2E63),
                trackColor = Color.White.copy(alpha = 0.45f)
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

                LaunchedEffect(displayedIndex) {
                    delay(120)
                    promptFocusRequester.requestFocus()
                    keyboardController?.show()
                }

                Box(
                    modifier = Modifier
                        .fillMaxSize()
                        .padding(top = 4.dp)
                ) {
                    Box(
                        modifier = Modifier
                            .fillMaxWidth()
                            .border(
                                width = if (promptHasError) 1.5.dp else 1.dp,
                                color = if (promptHasError) MaterialTheme.colorScheme.error else Color(0xFF8E2E63).copy(alpha = 0.28f),
                                shape = RoundedCornerShape(20.dp)
                            )
                            .background(
                                Color.White.copy(alpha = if (promptHasError) 0.82f else 0.72f),
                                RoundedCornerShape(20.dp)
                            )
                            .padding(16.dp)
                    ) {
                        Column {
                            Row(
                                verticalAlignment = Alignment.CenterVertically,
                                horizontalArrangement = Arrangement.spacedBy(8.dp)
                            ) {
                                Box(
                                    modifier = Modifier
                                        .size(28.dp)
                                        .background(Color(0xFF8E2E63), CircleShape),
                                    contentAlignment = Alignment.Center
                                ) {
                                    Text(
                                        text = "${displayedIndex + 1}",
                                        color = Color.White,
                                        style = MaterialTheme.typography.labelLarge,
                                        fontWeight = FontWeight.Bold
                                    )
                                }
                                Text(
                                    text = "Question ${displayedIndex + 1} of ${story.prompts.size}",
                                    style = MaterialTheme.typography.labelLarge,
                                    fontWeight = FontWeight.Bold
                                )
                            }
                            Spacer(modifier = Modifier.height(8.dp))
                            Text(
                                text = prompt.label,
                                style = MaterialTheme.typography.headlineSmall,
                                fontWeight = FontWeight.ExtraBold
                            )
                            Spacer(modifier = Modifier.height(8.dp))
                            OutlinedTextField(
                                value = inputMap[prompt.key].orEmpty(),
                                onValueChange = {
                                    inputMap[prompt.key] = it
                                    if (it.isNotBlank()) {
                                        errorMap[prompt.key] = false
                                    }
                                },
                                label = { Text(prompt.label) },
                                placeholder = { Text("Type something unhinged") },
                                modifier = Modifier
                                    .fillMaxWidth()
                                    .focusRequester(promptFocusRequester),
                                singleLine = true,
                                isError = promptHasError,
                                shape = RoundedCornerShape(14.dp),
                                keyboardOptions = KeyboardOptions(
                                    imeAction = if (displayedIsLast) ImeAction.Done else ImeAction.Next
                                ),
                                keyboardActions = KeyboardActions(
                                    onNext = { advancePrompt() },
                                    onDone = { advancePrompt() }
                                ),
                                colors = TextFieldDefaults.colors(
                                    focusedContainerColor = Color.White.copy(alpha = 0.9f),
                                    unfocusedContainerColor = Color.White.copy(alpha = 0.72f)
                                )
                            )
                            if (promptHasError) {
                                Spacer(modifier = Modifier.height(4.dp))
                                Text(
                                    text = "Need a value for this one.",
                                    color = MaterialTheme.colorScheme.error,
                                    style = MaterialTheme.typography.bodyMedium
                                )
                            }
                            Spacer(modifier = Modifier.height(8.dp))
                            Text(
                                text = if (displayedIsLast) "Final prompt. Lock it in." else "Answer this to unlock the next one.",
                                style = MaterialTheme.typography.bodyMedium,
                                color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.68f)
                            )
                        }
                    }
                }
            }

            if (!isKeyboardOpen) {
                Row(
                    modifier = Modifier
                        .fillMaxWidth()
                        .navigationBarsPadding(),
                    horizontalArrangement = Arrangement.spacedBy(12.dp)
                ) {
                    OutlinedButton(
                        onClick = {
                            focusManager.clearFocus()
                            if (safeIndex > 0) currentPromptIndex = safeIndex - 1
                        },
                        enabled = safeIndex > 0,
                        modifier = Modifier.weight(0.85f)
                    ) {
                        Text("Previous")
                    }

                    ChaosPrimaryButton(
                        label = if (isLastPrompt) "Reveal Story" else "Next Prompt",
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
    var revealPhase by remember(completedStory?.storyId) { mutableStateOf(RevealPhase.Anticipation) }
    val shareScale by animateFloatAsState(
        targetValue = if (revealPhase == RevealPhase.Visible) 1f else 0.96f,
        animationSpec = tween(
            durationMillis = MotionTokens.DurationMediumMs,
            easing = MotionTokens.EaseOutEmphasized
        ),
        label = "reveal_share_scale"
    )

    ChaosScaffold(title = "Your Masterpiece") { innerPadding ->
        if (completedStory == null) {
            Box(
                modifier = Modifier
                    .padding(innerPadding)
                    .fillMaxSize(),
                contentAlignment = Alignment.Center
            ) {
                Text("No completed story yet.")
            }
            return@ChaosScaffold
        }

        LaunchedEffect(completedStory.storyId) {
            revealPhase = RevealPhase.Anticipation
            haptics.performHapticFeedback(HapticFeedbackType.LongPress)
            delay(MotionTokens.DurationExpressiveMs.toLong())
            revealPhase = RevealPhase.Visible
            haptics.performHapticFeedback(HapticFeedbackType.LongPress)
        }

        Column(
            modifier = Modifier
                .padding(innerPadding)
                .padding(horizontal = 18.dp, vertical = 12.dp)
                .fillMaxSize()
                .verticalScroll(rememberScrollState())
                .navigationBarsPadding(),
            verticalArrangement = Arrangement.spacedBy(16.dp)
        ) {
            AnimatedContent(targetState = revealPhase, label = "reveal_phase_title") { phase ->
                Text(
                    text = if (phase == RevealPhase.Anticipation) "PREPARING THE STAGE" else "STORY REVEAL",
                    style = MaterialTheme.typography.labelLarge,
                    color = Color(0xFF8E2E63),
                    fontWeight = FontWeight.ExtraBold
                )
            }
            Text(
                text = completedStory.storyTitle,
                style = MaterialTheme.typography.headlineMedium,
                fontWeight = FontWeight.Bold
            )

            if (revealPhase == RevealPhase.Anticipation) {
                Box(
                    modifier = Modifier
                        .fillMaxWidth()
                        .border(1.dp, Color(0xFFFF7B54).copy(alpha = 0.45f), RoundedCornerShape(24.dp))
                        .background(
                            Brush.linearGradient(
                                colors = listOf(
                                    Color(0xFFFFD7A8).copy(alpha = 0.35f),
                                    Color(0xFFFF9E5A).copy(alpha = 0.26f)
                                )
                            ),
                            RoundedCornerShape(24.dp)
                        )
                ) {
                    Column(
                        modifier = Modifier.padding(18.dp),
                        horizontalAlignment = Alignment.CenterHorizontally
                    ) {
                        CircularProgressIndicator(color = Color(0xFF8E2E63))
                        Spacer(modifier = Modifier.height(12.dp))
                        Text(
                            text = "Matching your words with maximum chaos...",
                            style = MaterialTheme.typography.bodyMedium,
                            textAlign = TextAlign.Center
                        )
                    }
                }
            }

            AnimatedVisibility(
                visible = revealPhase == RevealPhase.Visible,
                enter = fadeIn(
                    animationSpec = tween(
                        durationMillis = MotionTokens.DurationMediumMs,
                        easing = MotionTokens.EaseOutStandard
                    )
                ) + scaleIn(
                    animationSpec = tween(
                        durationMillis = MotionTokens.DurationMediumMs,
                        easing = MotionTokens.EaseOutEmphasized
                    ),
                    initialScale = 0.96f
                )
            ) {
                Column(verticalArrangement = Arrangement.spacedBy(16.dp)) {
                    Box(
                        modifier = Modifier
                            .fillMaxWidth()
                            .border(1.dp, Color(0xFF6C63FF).copy(alpha = 0.45f), RoundedCornerShape(24.dp))
                            .background(
                                Brush.verticalGradient(
                                    colors = listOf(
                                        Color.White.copy(alpha = 0.92f),
                                        Color(0xFFFFF0E4).copy(alpha = 0.92f)
                                    )
                                ),
                                RoundedCornerShape(24.dp)
                            )
                            .padding(18.dp)
                    ) {
                        Column {
                            Text(
                                text = "Read this out loud with dramatic confidence:",
                                style = MaterialTheme.typography.bodyMedium,
                                color = MaterialTheme.colorScheme.onSurface.copy(alpha = 0.72f)
                            )
                            Spacer(modifier = Modifier.height(10.dp))
                            Text(
                                text = completedStory.renderedText,
                                style = MaterialTheme.typography.bodyLarge,
                                lineHeight = 29.sp
                            )
                        }
                    }

                    Spacer(modifier = Modifier.height(4.dp))
                    RevealPrimaryAction(
                        label = "Share This Disaster",
                        subtitle = "Drop this masterpiece in group chat",
                        accent = Color(0xFF8E2E63),
                        onClick = {
                            haptics.performHapticFeedback(HapticFeedbackType.LongPress)
                            val shareIntent = Intent(Intent.ACTION_SEND).apply {
                                type = "text/plain"
                                putExtra(Intent.EXTRA_TEXT, completedStory.renderedText)
                            }
                            context.startActivity(Intent.createChooser(shareIntent, "Share your Bad Libs story"))
                        },
                        modifier = Modifier
                            .fillMaxWidth()
                            .graphicsLayer {
                                scaleX = shareScale
                                scaleY = shareScale
                            }
                    )

                    Row(
                        modifier = Modifier.fillMaxWidth(),
                        horizontalArrangement = Arrangement.spacedBy(12.dp)
                    ) {
                        RevealSecondaryAction(
                            label = "Remix Words",
                            accent = Color(0xFF6C63FF),
                            onClick = {
                                haptics.performHapticFeedback(HapticFeedbackType.LongPress)
                                onRemix()
                            },
                            modifier = Modifier.weight(1f)
                        )
                        RevealSecondaryAction(
                            label = "Pick Another Story",
                            accent = Color(0xFFFF7B54),
                            onClick = {
                                haptics.performHapticFeedback(HapticFeedbackType.LongPress)
                                onPlayAgain()
                            },
                            modifier = Modifier.weight(1f)
                        )
                    }

                    Spacer(modifier = Modifier.height(4.dp))
                    Box(modifier = Modifier.fillMaxWidth(), contentAlignment = Alignment.Center) {
                        TextButton(onClick = onBackHome) {
                            Text("Back Home")
                        }
                    }
                }
            }
        }
    }
}

@Composable
private fun RevealPrimaryAction(
    label: String,
    subtitle: String,
    accent: Color,
    onClick: () -> Unit,
    modifier: Modifier = Modifier
) {
    androidx.compose.material3.Surface(
        onClick = onClick,
        modifier = modifier,
        shape = RoundedCornerShape(22.dp),
        color = Color.Transparent,
        border = BorderStroke(1.dp, accent.copy(alpha = 0.5f))
    ) {
        Row(
            modifier = Modifier
                .background(
                    Brush.linearGradient(
                        colors = listOf(
                            accent.copy(alpha = 0.22f),
                            accent.copy(alpha = 0.12f)
                        )
                    )
                )
                .padding(horizontal = 16.dp, vertical = 14.dp),
            horizontalArrangement = Arrangement.SpaceBetween,
            verticalAlignment = Alignment.CenterVertically
        ) {
            Column(modifier = Modifier.weight(1f)) {
                Text(text = label, style = MaterialTheme.typography.titleMedium, fontWeight = FontWeight.Bold)
                Spacer(modifier = Modifier.height(3.dp))
                Text(
                    text = subtitle,
                    style = MaterialTheme.typography.bodyMedium,
                    color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.75f)
                )
            }
            Text("↗", color = accent, style = MaterialTheme.typography.titleLarge, fontWeight = FontWeight.Bold)
        }
    }
}

@Composable
private fun RevealSecondaryAction(
    label: String,
    accent: Color,
    onClick: () -> Unit,
    modifier: Modifier = Modifier
) {
    androidx.compose.material3.Surface(
        onClick = onClick,
        modifier = modifier,
        shape = RoundedCornerShape(14.dp),
        color = accent.copy(alpha = 0.1f),
        border = BorderStroke(1.dp, accent.copy(alpha = 0.45f))
    ) {
        Box(
            modifier = Modifier.padding(vertical = 12.dp),
            contentAlignment = Alignment.Center
        ) {
            Text(
                text = label,
                style = MaterialTheme.typography.labelLarge,
                fontWeight = FontWeight.Bold,
                color = accent
            )
        }
    }
}
