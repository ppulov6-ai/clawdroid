package i18n

func init() {
	register("en", map[string]string{
		// Telegram
		"channel.thinking": "Thinking... 💭",

		// WebSocket
		"channel.config_required": "Configuration required",

		// Telegram commands (/help, /start, /show, /list)
		"cmd.help": `/start - Start the bot
/help - Show this help message
/show [model|channel] - Show current configuration
/list [models|channels] - List available options
`,
		"cmd.start":         "Hello! I am ClawDroid 🦞",
		"cmd.show.usage":    "Usage: /show [model|channel]",
		"cmd.show.model":    "Current Model: %s",
		"cmd.show.channel":  "Current Channel: telegram",
		"cmd.show.unknown":  "Unknown parameter: %s. Try 'model' or 'channel'.",
		"cmd.list.usage":    "Usage: /list [models|channels]",
		"cmd.list.models":   "Configured Model: %s\n\nTo change models, update config.json",
		"cmd.list.channels": "Enabled Channels:\n- %s",
		"cmd.list.unknown":  "Unknown parameter: %s. Try 'models' or 'channels'.",

		// Agent loop commands (/show, /list, /switch)
		// cmd.show.usage and cmd.list.usage are shared with Telegram commands
		"agent.cmd.show.model":        "Current model: %s",
		"agent.cmd.show.channel":      "Current channel: %s",
		"agent.cmd.show.unknown":      "Unknown show target: %s",
		"agent.cmd.list.models":       "Available models: glm-4.7, claude-3-5-sonnet, gpt-4o (configured in config.json/env)",
		"agent.cmd.list.no_channels":  "No channels enabled",
		"agent.cmd.list.channels":     "Enabled channels: %s",
		"agent.cmd.list.unknown":      "Unknown list target: %s",
		"agent.cmd.switch.usage":      "Usage: /switch [model|channel] to <name>",
		"agent.cmd.switch.model":      "Switched model from %s to %s",
		"agent.cmd.switch.channel":    "Switched target channel to %s (Note: this currently only validates existence)",
		"agent.cmd.switch.not_found":  "Channel '%s' not found or not enabled",
		"agent.cmd.switch.unknown":    "Unknown switch target: %s",
		"agent.cmd.channel_mgr_error": "Channel manager not initialized",
	})

	register("ru", map[string]string{
		// Telegram
		"channel.thinking": "Думаю... 💭",

		// WebSocket
		"channel.config_required": "Требуется настройка",

		// Telegram commands
		"cmd.help": "/start — Запустить бота\n/help — Показать справку\n/show [model|channel] — Показать текущие настройки\n/list [models|channels] — Показать доступные варианты\n",
		"cmd.start": "Здравствуйте! Я Джарвисджон.",
		"cmd.show.usage": "Использование: /show [model|channel]",
		"cmd.show.model": "Текущая модель: %s",
		"cmd.show.channel": "Текущий канал: telegram",
		"cmd.show.unknown": "Неизвестный параметр: %s. Укажите 'model' или 'channel'.",
		"cmd.list.usage": "Использование: /list [models|channels]",
		"cmd.list.models": "Настроенная модель: %s\n\nДля смены модели обновите config.json",
		"cmd.list.channels": "Включённые каналы:\n%s",
		"cmd.list.unknown": "Неизвестный параметр: %s. Укажите 'models' или 'channels'.",

		// Agent loop commands
		// cmd.show.usage and cmd.list.usage are shared with Telegram commands
		"agent.cmd.show.model": "Текущая модель: %s",
		"agent.cmd.show.channel": "Текущий канал: %s",
		"agent.cmd.show.unknown": "Неизвестный объект просмотра: %s",
		"agent.cmd.list.models": "Доступные модели: glm-4.7, claude-3-5-sonnet, gpt-4o (настройка в config.json или переменных окружения)",
		"agent.cmd.list.no_channels": "Нет включённых каналов",
		"agent.cmd.list.channels": "Включённые каналы: %s",
		"agent.cmd.list.unknown": "Неизвестный объект списка: %s",
		"agent.cmd.switch.usage": "Использование: /switch [model|channel] to <имя>",
		"agent.cmd.switch.model": "Модель изменена с %s на %s",
		"agent.cmd.switch.channel": "Выбран канал %s (пока проверяется только его наличие)",
		"agent.cmd.switch.not_found": "Канал '%s' не найден или отключён",
		"agent.cmd.switch.unknown": "Неизвестный объект переключения: %s",
		"agent.cmd.channel_mgr_error": "Управление каналами не инициализировано",
	})
}
