package io.clawdroid.core.domain.model

data class TtsConfig(
    val enginePackageName: String? = null,
    val voiceName: String? = null,
    val speechRate: Float = 1.0f,
    val pitch: Float = 1.0f
) {
    /** Ограничения применяются также к старым или повреждённым настройкам. */
    fun normalized(): TtsConfig = copy(
        speechRate = if (speechRate.isFinite()) speechRate.coerceIn(0.8f, 1.2f) else 1.0f,
        pitch = if (pitch.isFinite()) pitch.coerceIn(0.8f, 1.2f) else 1.0f
    )
}
