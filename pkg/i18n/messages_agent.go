package i18n

func init() {
	register("en", map[string]string{
		"agent.migration_notice":         "USER.md found. User management has migrated to a new format (users.json).\nAsk in chat to migrate, or update manually.\n\nFor manual update, create ~/.clawdroid/data/users.json in this format:\n```json\n{\n  \"users\": [{\n    \"name\": \"Your Name\",\n    \"channels\": { \"websocket\": [\"default\"] },\n    \"memo\": [\"Preferred language: English\"]\n  }]\n}\n```",
		"agent.context_window_warning":   "⚠️ Context window exceeded. Compressing history and retrying...",
		"agent.memory_threshold_warning": "⚠️ Memory threshold reached. Optimizing conversation history...",
		"agent.rate_limited":             "Rate limited: %s. Please try again later.",
		"agent.rate_limited_tool":        "Rate limited: %s",
	})

	register("ru", map[string]string{
		"agent.migration_notice": "Найден USER.md. Управление пользователями переведено на новый формат users.json.\nПопросите выполнить перенос в чате или обновите данные вручную.\n\nДля ручного обновления создайте ~/.clawdroid/data/users.json в таком формате:\n```json\n{\n  \"users\": [{\n    \"name\": \"Ваше имя\",\n    \"channels\": { \"websocket\": [\"default\"] },\n    \"memo\": [\"Предпочитаемый язык: русский\"]\n  }]\n}\n```",
		"agent.context_window_warning": "⚠️ Превышен размер контекста. Сжимаю историю и повторяю запрос...",
		"agent.memory_threshold_warning": "⚠️ Достигнут предел памяти. Оптимизирую историю разговора...",
		"agent.rate_limited": "Превышен лимит запросов: %s. Попробуйте позже.",
		"agent.rate_limited_tool": "Превышен лимит запросов: %s",
	})
}
