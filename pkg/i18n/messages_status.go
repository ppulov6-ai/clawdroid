package i18n

func init() {
	register("en", map[string]string{
		// status labels
		"status.thinking":    "Thinking...",
		"status.processing":  "Processing...",
		"status.interrupted": "[Response was interrupted]",

		// web
		"status.searching":     "Searching...",
		"status.searching_q":   "Searching... (%s)",
		"status.fetching_page": "Fetching page...",
		"status.fetching_q":    "Fetching page... (%s)",

		// file operations
		"status.reading_file":     "Reading file...",
		"status.reading_file_q":   "Reading file... (%s)",
		"status.writing_file":     "Writing file...",
		"status.writing_file_q":   "Writing file... (%s)",
		"status.editing_file":     "Editing file...",
		"status.editing_file_q":   "Editing file... (%s)",
		"status.appending_file":   "Appending to file...",
		"status.appending_file_q": "Appending to file... (%s)",

		// directory
		"status.listing_dir":   "Checking folder...",
		"status.listing_dir_q": "Checking folder... (%s)",

		// exec
		"status.running_command":   "Running command...",
		"status.running_command_q": "Running command... (%s)",

		// memory
		"status.memory_read":         "Loading memory...",
		"status.memory_read_daily":   "Loading today's memo...",
		"status.memory_write":        "Writing memory...",
		"status.memory_append_daily": "Appending to today's memo...",
		"status.memory_default":      "Memory operation...",

		// skill
		"status.skill_list":    "Getting skill list...",
		"status.skill_read":    "Loading skill...",
		"status.skill_read_q":  "Loading skill... (%s)",
		"status.skill_default": "Skill operation...",

		// cron
		"status.cron_add":     "Setting reminder...",
		"status.cron_list":    "Getting schedule...",
		"status.cron_remove":  "Removing schedule...",
		"status.cron_default": "Updating schedule...",

		// message
		"status.sending_message": "Sending message...",

		// spawn/subagent
		"status.spawn":      "Starting subtask...",
		"status.spawn_q":    "Starting subtask... (%s)",
		"status.subagent":   "Running subtask...",
		"status.subagent_q": "Running subtask... (%s)",

		// android
		"status.android_search_apps":  "Searching apps...",
		"status.android_app_info":     "Getting app info...",
		"status.android_app_info_q":   "Getting app info... (%s)",
		"status.android_launch_app":   "Launching app...",
		"status.android_launch_app_q": "Launching app... (%s)",
		"status.android_screenshot":   "Taking screenshot...",
		"status.android_get_ui_tree":  "Getting UI elements...",
		"status.android_tap":          "Tapping...",
		"status.android_swipe":        "Swiping...",
		"status.android_text":         "Entering text...",
		"status.android_keyevent":     "Key operation...",
		"status.android_keyevent_q":   "Key operation... (%s)",
		"status.android_broadcast":    "Sending broadcast...",
		"status.android_intent":       "Sending intent...",
		"status.android_default":      "Device operation...",

		// exit
		"status.exit": "Shutting down assistant...",

		// mcp
		"status.mcp_list":    "Getting MCP server list...",
		"status.mcp_tools":   "Getting MCP tools...",
		"status.mcp_tools_q": "Getting MCP tools... (%s)",
		"status.mcp_call":    "Running MCP tool...",
		"status.mcp_call_q":  "Running MCP tool... (%s)",
		"status.mcp_call_sq": "Running MCP tool... (%s/%s)",
		"status.mcp_default": "MCP operation...",
	})

	register("ru", map[string]string{
		// status labels
		"status.thinking": "Думаю...",
		"status.processing": "Обрабатываю...",
		"status.interrupted": "[Ответ прерван]",

		// web
		"status.searching": "Ищу...",
		"status.searching_q": "Ищу... (%s)",
		"status.fetching_page": "Загружаю страницу...",
		"status.fetching_q": "Загружаю страницу... (%s)",

		// file operations
		"status.reading_file": "Читаю файл...",
		"status.reading_file_q": "Читаю файл... (%s)",
		"status.writing_file": "Записываю файл...",
		"status.writing_file_q": "Записываю файл... (%s)",
		"status.editing_file": "Редактирую файл...",
		"status.editing_file_q": "Редактирую файл... (%s)",
		"status.appending_file": "Дополняю файл...",
		"status.appending_file_q": "Дополняю файл... (%s)",

		// directory
		"status.listing_dir": "Проверяю папку...",
		"status.listing_dir_q": "Проверяю папку... (%s)",

		// exec
		"status.running_command": "Выполняю команду...",
		"status.running_command_q": "Выполняю команду... (%s)",

		// memory
		"status.memory_read": "Загружаю память...",
		"status.memory_read_daily": "Загружаю сегодняшние заметки...",
		"status.memory_write": "Сохраняю в память...",
		"status.memory_append_daily": "Дополняю сегодняшние заметки...",
		"status.memory_default": "Работаю с памятью...",

		// skill
		"status.skill_list": "Получаю список навыков...",
		"status.skill_read": "Загружаю навык...",
		"status.skill_read_q": "Загружаю навык... (%s)",
		"status.skill_default": "Работаю с навыком...",

		// cron
		"status.cron_add": "Создаю напоминание...",
		"status.cron_list": "Получаю расписание...",
		"status.cron_remove": "Удаляю задачу из расписания...",
		"status.cron_default": "Изменяю расписание...",

		// message
		"status.sending_message": "Отправляю сообщение...",

		// spawn/subagent
		"status.spawn": "Запускаю подзадачу...",
		"status.spawn_q": "Запускаю подзадачу... (%s)",
		"status.subagent": "Выполняю подзадачу...",
		"status.subagent_q": "Выполняю подзадачу... (%s)",

		// android
		"status.android_search_apps": "Ищу приложения...",
		"status.android_app_info": "Получаю сведения о приложении...",
		"status.android_app_info_q": "Получаю сведения о приложении... (%s)",
		"status.android_launch_app": "Открываю приложение...",
		"status.android_launch_app_q": "Открываю приложение... (%s)",
		"status.android_screenshot": "Делаю снимок экрана...",
		"status.android_get_ui_tree": "Получаю элементы экрана...",
		"status.android_tap": "Нажимаю...",
		"status.android_swipe": "Прокручиваю...",
		"status.android_text": "Ввожу текст...",
		"status.android_keyevent": "Нажимаю клавишу...",
		"status.android_keyevent_q": "Нажимаю клавишу... (%s)",
		"status.android_broadcast": "Отправляю системное событие...",
		"status.android_intent": "Отправляю системную команду...",
		"status.android_default": "Управляю устройством...",

		// exit
		"status.exit": "Завершаю работу помощника...",

		// mcp
		"status.mcp_list": "Получаю список серверов MCP...",
		"status.mcp_tools": "Получаю инструменты MCP...",
		"status.mcp_tools_q": "Получаю инструменты MCP... (%s)",
		"status.mcp_call": "Выполняю инструмент MCP...",
		"status.mcp_call_q": "Выполняю инструмент MCP... (%s)",
		"status.mcp_call_sq": "Выполняю инструмент MCP... (%s/%s)",
		"status.mcp_default": "Работаю с MCP...",
	})
}
