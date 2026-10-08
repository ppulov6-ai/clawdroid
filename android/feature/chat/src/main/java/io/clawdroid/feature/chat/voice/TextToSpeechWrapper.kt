package io.clawdroid.feature.chat.voice

import android.content.Context
import android.speech.tts.TextToSpeech
import android.speech.tts.UtteranceProgressListener
import io.clawdroid.core.domain.model.TtsConfig
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.SupervisorJob
import kotlinx.coroutines.cancel
import kotlinx.coroutines.delay
import kotlinx.coroutines.withContext
import kotlinx.coroutines.withTimeoutOrNull
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.launch
import kotlinx.coroutines.suspendCancellableCoroutine
import kotlinx.coroutines.sync.Mutex
import kotlinx.coroutines.sync.withLock
import java.util.UUID
import kotlin.coroutines.resume

class TextToSpeechWrapper(
    private val context: Context,
    ttsConfigFlow: Flow<TtsConfig>
) {

    private val scope = CoroutineScope(SupervisorJob() + Dispatchers.Main)
    private val mutex = Mutex()

    private var tts: TextToSpeech? = null
    private var initialized = false
    private var engineGeneration = 0
    private var currentConfig = TtsConfig()
    private var currentEnginePackage: String? = null

    init {
        initTtsEngine(null)

        scope.launch {
            ttsConfigFlow.collect { config ->
                mutex.withLock {
                    currentConfig = config.normalized()
                    if (config.enginePackageName != currentEnginePackage) {
                        switchEngine(config.enginePackageName)
                    }
                }
            }
        }
    }

    private fun initTtsEngine(enginePackageName: String?) {
        tts?.stop()
        tts?.shutdown()
        initialized = false
        currentEnginePackage = enginePackageName
        val generation = ++engineGeneration

        val listener = TextToSpeech.OnInitListener { status ->
            if (generation == engineGeneration && status == TextToSpeech.SUCCESS) {
                initialized = true
            }
        }

        tts = if (enginePackageName != null) {
            TextToSpeech(context, listener, enginePackageName)
        } else {
            TextToSpeech(context, listener)
        }
    }

    private fun switchEngine(enginePackageName: String?) {
        initTtsEngine(enginePackageName)
    }

    private fun applyConfig(engine: TextToSpeech): Boolean {
        // Проверяем даже ранее сохранённый голос: только русский и без сети.
        val voices = engine.voices.orEmpty().filter {
            it.locale.language == "ru" && !it.isNetworkConnectionRequired
        }
        val voice = voices.firstOrNull { it.name == currentConfig.voiceName }
            ?: voices.sortedWith(
                compareByDescending<android.speech.tts.Voice> { it.locale.country == "RU" }
                    .thenByDescending { it.quality }
                    .thenBy { it.name }
            ).firstOrNull()
            ?: return false
        if (engine.setVoice(voice) != TextToSpeech.SUCCESS) return false
        val config = currentConfig.normalized()
        engine.setSpeechRate(config.speechRate)
        engine.setPitch(config.pitch)
        return true
    }

    suspend fun speak(text: String): Boolean = withContext(Dispatchers.Main) {
        // Привязка к движку асинхронная: первая команда ждёт её завершения.
        val ready = withTimeoutOrNull(5_000) {
            while (!initialized && tts != null) delay(50)
            initialized
        } ?: false
        if (!ready) return@withContext false
        withTimeoutOrNull(60_000) { speakReady(text) } ?: false
    }

    private suspend fun speakReady(text: String): Boolean = suspendCancellableCoroutine { cont ->
        val engine = tts
        if (engine == null || !initialized) {
            cont.resume(false)
            return@suspendCancellableCoroutine
        }

        if (!applyConfig(engine)) {
            cont.resume(false)
            return@suspendCancellableCoroutine
        }

        val utteranceId = UUID.randomUUID().toString()

        engine.setOnUtteranceProgressListener(object : UtteranceProgressListener() {
            override fun onStart(id: String?) {}

            override fun onDone(id: String?) {
                if (id == utteranceId && cont.isActive) {
                    cont.resume(true)
                }
            }

            override fun onStop(id: String?, interrupted: Boolean) {
                if (id == utteranceId && cont.isActive) cont.resume(false)
            }

            @Deprecated("Deprecated in Java")
            override fun onError(id: String?) {
                if (id == utteranceId && cont.isActive) {
                    cont.resume(false)
                }
            }

            override fun onError(id: String?, errorCode: Int) {
                if (id == utteranceId && cont.isActive) {
                    cont.resume(false)
                }
            }
        })

        cont.invokeOnCancellation {
            engine.stop()
        }

        if (engine.speak(text, TextToSpeech.QUEUE_FLUSH, null, utteranceId) == TextToSpeech.ERROR
            && cont.isActive
        ) {
            cont.resume(false)
        }
    }

    fun stop() {
        tts?.stop()
    }

    fun destroy() {
        ++engineGeneration
        initialized = false
        scope.cancel()
        tts?.stop()
        tts?.shutdown()
        tts = null
    }
}
