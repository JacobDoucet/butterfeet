package com.butterfeetlabs.badlibs

import android.content.Intent
import androidx.compose.foundation.BorderStroke
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.BoxScope
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.imePadding
import androidx.compose.foundation.layout.navigationBarsPadding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.WindowInsets
import androidx.compose.foundation.layout.isImeVisible
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.lazy.itemsIndexed
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.text.KeyboardActions
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Card
import androidx.compose.material3.CardDefaults
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.foundation.layout.ExperimentalLayoutApi
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
import androidx.compose.ui.platform.LocalFocusManager
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.text.input.ImeAction
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.focus.FocusDirection
import androidx.compose.ui.unit.sp
import androidx.compose.ui.unit.dp
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
import com.butterfeetlabs.badlibs.data.StoryPack
import com.butterfeetlabs.badlibs.data.StoryRepository
import com.butterfeetlabs.badlibs.ui.theme.BadLibsTheme
import com.butterfeetlabs.badlibs.ui.theme.ChaosOrange
import kotlinx.coroutines.launch

private sealed class Screen(val route: String) {
    data object Home : Screen("home")
    data object Packs : Screen("packs")
    data object Stories : Screen("stories/{packId}")
    data object Prompts : Screen("prompts/{packId}/{storyId}")
    data object Result : Screen("result")
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
    val navController = rememberNavController()
    val snackbarHostState = remember { SnackbarHostState() }
    val viewModel: MainViewModel = viewModel(
        factory = MainViewModelFactory(StoryRepository(context))
    )

    BadLibsTheme {
        ChaosBackdrop {
            NavHost(navController = navController, startDestination = Screen.Home.route) {
                composable(Screen.Home.route) {
                    HomeScreen(onStart = { navController.navigate("packs") })
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
                    route = Screen.Prompts.route,
                    arguments = listOf(
                        navArgument("packId") { type = NavType.StringType },
                        navArgument("storyId") { type = NavType.StringType }
                    )
                ) { backStackEntry ->
                    val packId = backStackEntry.arguments?.getString("packId").orEmpty()
                    val storyId = backStackEntry.arguments?.getString("storyId").orEmpty()
                    val story = viewModel.findStory(packId, storyId)
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
            ),
        content = content
    )
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
private fun ChaosSurfaceCard(
    modifier: Modifier = Modifier,
    onClick: (() -> Unit)? = null,
    content: @Composable () -> Unit
) {
    if (onClick != null) {
        Card(
            modifier = modifier,
            onClick = onClick,
            shape = RoundedCornerShape(22.dp),
            border = BorderStroke(1.dp, Color.White.copy(alpha = 0.7f)),
            elevation = CardDefaults.cardElevation(defaultElevation = 3.dp),
            colors = CardDefaults.cardColors(containerColor = Color.White.copy(alpha = 0.9f))
        ) {
            content()
        }
    } else {
        Card(
            modifier = modifier,
            shape = RoundedCornerShape(22.dp),
            border = BorderStroke(1.dp, Color.White.copy(alpha = 0.7f)),
            elevation = CardDefaults.cardElevation(defaultElevation = 3.dp),
            colors = CardDefaults.cardColors(containerColor = Color.White.copy(alpha = 0.9f))
        ) {
            content()
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
private fun HomeScreen(onStart: () -> Unit) {
    ChaosScaffold(title = "Bad Libs") { innerPadding ->
        LazyColumn(
            modifier = Modifier
                .padding(innerPadding)
                .fillMaxSize(),
            contentPadding = PaddingValues(horizontal = 20.dp, vertical = 8.dp),
            verticalArrangement = Arrangement.spacedBy(14.dp)
        ) {
            item {
                ChaosSurfaceCard(modifier = Modifier.fillMaxWidth()) {
                    Column(modifier = Modifier.padding(20.dp)) {
                        Text(
                            text = "Make stupid stories with serious confidence.",
                            style = MaterialTheme.typography.headlineLarge,
                            lineHeight = 38.sp,
                            color = MaterialTheme.colorScheme.onSurface
                        )
                        Spacer(modifier = Modifier.height(10.dp))
                        Text(
                            text = "Pick a pack, add cursed words, and reveal your literary disaster.",
                            style = MaterialTheme.typography.bodyLarge
                        )
                        Spacer(modifier = Modifier.height(18.dp))
                        ChaosPrimaryButton(
                            label = "Start The Chaos",
                            onClick = onStart,
                            modifier = Modifier.fillMaxWidth()
                        )
                    }
                }
            }
            item {
                Text(
                    text = "How it works",
                    style = MaterialTheme.typography.titleMedium,
                    fontWeight = FontWeight.Bold,
                    modifier = Modifier.padding(horizontal = 4.dp)
                )
            }
            item {
                ChaosStepCard(
                    emoji = "1",
                    title = "Pick a pack",
                    copy = "Choose your flavor of nonsense and lock in the vibe."
                )
            }
            item {
                ChaosStepCard(
                    emoji = "2",
                    title = "Fill prompts",
                    copy = "Use weird words. Regret nothing. Spell-check optional."
                )
            }
            item {
                ChaosStepCard(
                    emoji = "3",
                    title = "Reveal and share",
                    copy = "Drop your masterpiece in group chat and cause mild panic."
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
                        item {
                            Text(
                                text = "Pick your disaster theme",
                                style = MaterialTheme.typography.titleMedium,
                                modifier = Modifier.padding(horizontal = 4.dp, vertical = 6.dp)
                            )
                        }
                        items(packs, key = { it.id }) { pack ->
                            ChaosSurfaceCard(
                                modifier = Modifier.fillMaxWidth(),
                                onClick = { onOpenPack(pack.id) }
                            ) {
                                Column(modifier = Modifier.padding(18.dp)) {
                                    Row(
                                        modifier = Modifier.fillMaxWidth(),
                                        horizontalArrangement = Arrangement.SpaceBetween,
                                        verticalAlignment = Alignment.CenterVertically
                                    ) {
                                        Text(pack.title, style = MaterialTheme.typography.titleLarge)
                                        PackCountPill(count = pack.stories.size)
                                    }
                                    Spacer(modifier = Modifier.height(8.dp))
                                    Text(pack.description, style = MaterialTheme.typography.bodyMedium)
                                }
                            }
                        }
                    }

                    Row(
                        modifier = Modifier
                            .fillMaxWidth()
                            .padding(horizontal = 16.dp, vertical = 12.dp)
                            .navigationBarsPadding(),
                        horizontalArrangement = Arrangement.spacedBy(12.dp)
                    ) {
                        OutlinedButton(onClick = onBack, modifier = Modifier.weight(0.8f)) {
                            Text("Back")
                        }
                        ChaosPrimaryButton(
                            label = "Pick Random",
                            onClick = {
                                packs.randomOrNull()?.let { onOpenPack(it.id) }
                            },
                            modifier = Modifier.weight(1.2f)
                        )
                    }
                }
            }
        }
    }
}

@Composable
private fun PackCountPill(count: Int) {
    Box(
        modifier = Modifier
            .background(Color(0xFFFFE7D2), RoundedCornerShape(50))
            .padding(horizontal = 12.dp, vertical = 6.dp)
    ) {
        Text(
            text = "$count stories",
            style = MaterialTheme.typography.labelLarge,
            color = MaterialTheme.colorScheme.onBackground
        )
    }
}

@OptIn(ExperimentalMaterial3Api::class)
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
                        item {
                            Text(
                                text = "Pick a story and make it worse.",
                                style = MaterialTheme.typography.titleMedium,
                                modifier = Modifier.padding(horizontal = 4.dp, vertical = 4.dp)
                            )
                        }
                        items(pack.stories, key = { it.id }) { story ->
                            ChaosSurfaceCard(
                                modifier = Modifier.fillMaxWidth(),
                                onClick = { onOpenStory(story.id) }
                            ) {
                                Column(modifier = Modifier.padding(18.dp)) {
                                    Text(story.title, style = MaterialTheme.typography.titleLarge)
                                    Spacer(modifier = Modifier.height(8.dp))
                                    Text(
                                        text = "${story.prompts.size} prompts • ~${(story.prompts.size * 10).coerceAtMost(90)} sec",
                                        style = MaterialTheme.typography.bodyMedium
                                    )
                                    Spacer(modifier = Modifier.height(10.dp))
                                    Row(horizontalArrangement = Arrangement.spacedBy(8.dp)) {
                                        StoryMetaChip(label = "Rated ${story.rating}")
                                        StoryMetaChip(label = "Chaos-ready")
                                    }
                                }
                            }
                        }
                    }

                    Row(
                        modifier = Modifier
                            .fillMaxWidth()
                            .padding(horizontal = 16.dp, vertical = 12.dp)
                            .navigationBarsPadding(),
                        horizontalArrangement = Arrangement.spacedBy(12.dp)
                    ) {
                        OutlinedButton(onClick = onBack, modifier = Modifier.weight(0.8f)) {
                            Text("Back")
                        }
                        ChaosPrimaryButton(
                            label = "Pick Random",
                            onClick = {
                                pack.stories.randomOrNull()?.let { onOpenStory(it.id) }
                            },
                            modifier = Modifier.weight(1.2f)
                        )
                    }
                }
            }
        }
    }
}

@Composable
private fun StoryMetaChip(label: String) {
    Box(
        modifier = Modifier
            .border(1.dp, Color(0x26A64D79), RoundedCornerShape(50))
            .background(Color(0x14A64D79), RoundedCornerShape(50))
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
    val inputMap = remember(story?.id) { mutableStateMapOf<String, String>() }
    val errorMap = remember(story?.id) { mutableStateMapOf<String, Boolean>() }
    val isKeyboardOpen = WindowInsets.isImeVisible

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

        Column(
            modifier = Modifier
                .padding(innerPadding)
                .imePadding()
                .fillMaxSize()
        ) {
            LazyColumn(
                modifier = Modifier
                    .weight(1f)
                    .fillMaxWidth(),
                contentPadding = PaddingValues(horizontal = 16.dp, vertical = 8.dp),
                verticalArrangement = Arrangement.spacedBy(12.dp)
            ) {
                item {
                    ChaosSurfaceCard(modifier = Modifier.fillMaxWidth()) {
                        Column(modifier = Modifier.padding(16.dp)) {
                            Text(
                                text = "Fill these in. Keep it weird.",
                                style = MaterialTheme.typography.titleMedium
                            )
                            Spacer(modifier = Modifier.height(10.dp))
                            LinearProgressIndicator(
                                progress = { progress },
                                modifier = Modifier.fillMaxWidth(),
                                color = ChaosOrange,
                                trackColor = Color.White.copy(alpha = 0.5f)
                            )
                            Spacer(modifier = Modifier.height(6.dp))
                            Text(
                                text = "$filledPrompts / ${story.prompts.size} prompts filled",
                                style = MaterialTheme.typography.labelLarge
                            )
                        }
                    }
                }

                itemsIndexed(story.prompts, key = { _, prompt -> prompt.key }) { index, prompt ->
                    val isLast = index == story.prompts.lastIndex
                    val hasError = errorMap[prompt.key] == true

                    ChaosSurfaceCard(modifier = Modifier.fillMaxWidth()) {
                        Column(modifier = Modifier.padding(14.dp)) {
                            Text(
                                text = "${index + 1}. ${prompt.label}",
                                style = MaterialTheme.typography.labelLarge
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
                                modifier = Modifier.fillMaxWidth(),
                                singleLine = true,
                                isError = hasError,
                                keyboardOptions = KeyboardOptions(
                                    imeAction = if (isLast) ImeAction.Done else ImeAction.Next
                                ),
                                keyboardActions = KeyboardActions(
                                    onNext = { focusManager.moveFocus(FocusDirection.Down) },
                                    onDone = { submitStory() }
                                ),
                                colors = TextFieldDefaults.colors(
                                    focusedContainerColor = Color.White.copy(alpha = 0.7f),
                                    unfocusedContainerColor = Color.White.copy(alpha = 0.55f)
                                )
                            )
                            if (hasError) {
                                Spacer(modifier = Modifier.height(4.dp))
                                Text(
                                    text = "Need a value for this one.",
                                    color = MaterialTheme.colorScheme.error,
                                    style = MaterialTheme.typography.bodyMedium
                                )
                            }
                        }
                    }
                }

                item {
                    // Extra tail space keeps the final field visible above the sticky action bar + keyboard.
                    Spacer(modifier = Modifier.height(120.dp))
                }
            }

            if (!isKeyboardOpen) {
                Row(
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(horizontal = 16.dp, vertical = 12.dp)
                        .navigationBarsPadding(),
                    horizontalArrangement = Arrangement.spacedBy(12.dp)
                ) {
                    OutlinedButton(onClick = onBack, modifier = Modifier.weight(0.8f)) {
                        Text("Back")
                    }
                    ChaosPrimaryButton(
                        label = "Reveal Story",
                        onClick = { submitStory() },
                        modifier = Modifier.weight(1.2f)
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

        Column(
            modifier = Modifier
                .padding(innerPadding)
                .padding(20.dp)
                .fillMaxSize(),
            verticalArrangement = Arrangement.spacedBy(16.dp)
        ) {
            Text(
                text = "Story Reveal",
                style = MaterialTheme.typography.labelLarge,
                color = MaterialTheme.colorScheme.secondary
            )
            Text(
                text = completedStory.storyTitle,
                style = MaterialTheme.typography.headlineMedium,
                fontWeight = FontWeight.Bold
            )
            ChaosSurfaceCard(modifier = Modifier.fillMaxWidth()) {
                Column(modifier = Modifier.padding(18.dp)) {
                    Text(
                        text = "Read this out loud with dramatic confidence:",
                        style = MaterialTheme.typography.bodyMedium,
                        color = MaterialTheme.colorScheme.onSurface.copy(alpha = 0.7f)
                    )
                    Spacer(modifier = Modifier.height(10.dp))
                    Text(
                        text = completedStory.renderedText,
                        style = MaterialTheme.typography.bodyLarge,
                        lineHeight = 28.sp
                    )
                }
            }

            Spacer(modifier = Modifier.height(4.dp))
            ChaosPrimaryButton(
                label = "Share This Disaster",
                onClick = {
                    val shareIntent = Intent(Intent.ACTION_SEND).apply {
                        type = "text/plain"
                        putExtra(Intent.EXTRA_TEXT, completedStory.renderedText)
                    }
                    context.startActivity(Intent.createChooser(shareIntent, "Share your Bad Libs story"))
                },
                modifier = Modifier.fillMaxWidth()
            )

            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.spacedBy(12.dp)
            ) {
                OutlinedButton(
                    onClick = onRemix,
                    modifier = Modifier.weight(1f),
                    shape = RoundedCornerShape(12.dp)
                ) {
                    Text("Remix Words")
                }
                OutlinedButton(
                    onClick = onPlayAgain,
                    modifier = Modifier.weight(1f),
                    shape = RoundedCornerShape(12.dp)
                ) {
                    Text("Pick Another Story")
                }
            }

            Spacer(modifier = Modifier.height(4.dp))
            Box(modifier = Modifier.fillMaxWidth(), contentAlignment = Alignment.Center) {
                TextButton(onClick = onBackHome) {
                    Text("Back Home")
                }
            }
            Text(
                text = "Pro tip: bad grammar makes it better.",
                style = MaterialTheme.typography.bodyMedium,
                color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.65f)
            )
        }
    }
}
