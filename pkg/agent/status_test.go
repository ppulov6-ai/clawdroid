package agent

import (
	"testing"

	_ "github.com/KarakuriAgent/clawdroid/pkg/i18n"
)

// --- strArg ---

func TestStrArg(t *testing.T) {
	tests := []struct {
		name string
		args map[string]interface{}
		key  string
		want string
	}{
		{"existing key", map[string]interface{}{"k": "v"}, "k", "v"},
		{"missing key", map[string]interface{}{"k": "v"}, "other", ""},
		{"non-string value", map[string]interface{}{"k": 123}, "k", ""},
		{"nil map", nil, "k", ""},
		{"empty string value", map[string]interface{}{"k": ""}, "k", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strArg(tt.args, tt.key)
			if got != tt.want {
				t.Errorf("strArg(%v, %q) = %q, want %q", tt.args, tt.key, got, tt.want)
			}
		})
	}
}

// --- truncLabel ---

func TestTruncLabel(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		maxRunes int
		want     string
	}{
		{"short string", "hello", 10, "hello"},
		{"exact max", "hello", 5, "hello"},
		{"exceeds max", "hello world", 5, "hello..."},
		{"unicode within", "日本語テスト", 6, "日本語テスト"},
		{"unicode exceeds", "日本語テスト", 3, "日本語..."},
		{"empty", "", 5, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := truncLabel(tt.s, tt.maxRunes)
			if got != tt.want {
				t.Errorf("truncLabel(%q, %d) = %q, want %q", tt.s, tt.maxRunes, got, tt.want)
			}
		})
	}
}

// --- hostFromURL ---

func TestHostFromURL(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want string
	}{
		{"normal URL", "https://example.com/path", "example.com"},
		{"URL with port", "https://example.com:8080/path", "example.com:8080"},
		{"invalid URL", "://bad", "://bad"},
		{"empty string", "", ""},
		{"no scheme", "not-a-url", "not-a-url"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hostFromURL(tt.url)
			if got != tt.want {
				t.Errorf("hostFromURL(%q) = %q, want %q", tt.url, got, tt.want)
			}
		})
	}
}

// --- statusLabel (Russian locale, preserving original test expectations) ---

func TestStatusLabel(t *testing.T) {
	tests := []struct {
		name     string
		toolName string
		args     map[string]interface{}
		contains string
	}{
		{"web_search with query", "web_search", map[string]interface{}{"query": "golang"}, "golang"},
		{"web_search no query", "web_search", map[string]interface{}{}, "Ищу..."},
		{"web_fetch with url", "web_fetch", map[string]interface{}{"url": "https://example.com/page"}, "example.com"},
		{"web_fetch no url", "web_fetch", map[string]interface{}{}, "Загружаю страницу..."},
		{"read_file with path", "read_file", map[string]interface{}{"path": "/home/user/file.txt"}, "file.txt"},
		{"read_file no path", "read_file", map[string]interface{}{}, "Читаю файл..."},
		{"write_file", "write_file", map[string]interface{}{"path": "/tmp/out.txt"}, "out.txt"},
		{"edit_file", "edit_file", map[string]interface{}{}, "Редактирую файл..."},
		{"append_file", "append_file", map[string]interface{}{}, "Дополняю файл..."},
		{"list_dir with path", "list_dir", map[string]interface{}{"path": "/home/user/docs"}, "docs/"},
		{"list_dir no path", "list_dir", map[string]interface{}{}, "Проверяю папку..."},
		{"exec with command", "exec", map[string]interface{}{"command": "ls -la"}, "ls -la"},
		{"exec no command", "exec", map[string]interface{}{}, "Выполняю команду..."},
		{"memory", "memory", map[string]interface{}{"action": "read_long_term"}, "Загружаю память..."},
		{"skill", "skill", map[string]interface{}{"action": "skill_list"}, "Получаю список навыков..."},
		{"cron", "cron", map[string]interface{}{"action": "add"}, "Создаю напоминание..."},
		{"message", "message", map[string]interface{}{}, "Отправляю сообщение..."},
		{"spawn with label", "spawn", map[string]interface{}{"label": "task1"}, "task1"},
		{"spawn no label", "spawn", map[string]interface{}{}, "Запускаю подзадачу..."},
		{"subagent with label", "subagent", map[string]interface{}{"label": "sub1"}, "sub1"},
		{"subagent no label", "subagent", map[string]interface{}{}, "Выполняю подзадачу..."},
		{"android", "android", map[string]interface{}{"action": "screenshot"}, "Делаю снимок экрана..."},
		{"exit", "exit", map[string]interface{}{}, "Завершаю работу помощника..."},
		{"mcp", "mcp", map[string]interface{}{"action": "mcp_list"}, "Получаю список серверов MCP..."},
		{"unknown tool", "unknown_tool", map[string]interface{}{}, "Обрабатываю..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := statusLabel(tt.toolName, tt.args, "ru")
			if got == "" {
				t.Error("statusLabel returned empty string")
			}
			if !containsStr(got, tt.contains) {
				t.Errorf("statusLabel(%q, %v, ru) = %q, want to contain %q", tt.toolName, tt.args, got, tt.contains)
			}
		})
	}
}

// --- statusLabel (English locale) ---

func TestStatusLabelEnglish(t *testing.T) {
	tests := []struct {
		name     string
		toolName string
		args     map[string]interface{}
		contains string
	}{
		{"web_search", "web_search", map[string]interface{}{}, "Searching..."},
		{"read_file", "read_file", map[string]interface{}{}, "Reading file..."},
		{"exec", "exec", map[string]interface{}{}, "Running command..."},
		{"memory", "memory", map[string]interface{}{"action": "read_long_term"}, "Loading memory..."},
		{"exit", "exit", map[string]interface{}{}, "Shutting down assistant..."},
		{"unknown", "unknown_tool", map[string]interface{}{}, "Processing..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := statusLabel(tt.toolName, tt.args, "en")
			if !containsStr(got, tt.contains) {
				t.Errorf("statusLabel(%q, %v, en) = %q, want to contain %q", tt.toolName, tt.args, got, tt.contains)
			}
		})
	}
}

// --- fileStatusLabel ---

func TestFileStatusLabel(t *testing.T) {
	got := fileStatusLabel("ru", "status.reading_file", "status.reading_file_q", map[string]interface{}{"path": "/home/user/test.go"})
	if got != "Читаю файл... (test.go)" {
		t.Errorf("got %q", got)
	}

	got = fileStatusLabel("ru", "status.reading_file", "status.reading_file_q", map[string]interface{}{})
	if got != "Читаю файл..." {
		t.Errorf("got %q", got)
	}
}

// --- memoryStatusLabel ---

func TestMemoryStatusLabel(t *testing.T) {
	tests := []struct {
		action string
		want   string
	}{
		{"read_long_term", "Загружаю память..."},
		{"read_daily", "Загружаю сегодняшние заметки..."},
		{"write_long_term", "Сохраняю в память..."},
		{"append_daily", "Дополняю сегодняшние заметки..."},
		{"unknown", "Работаю с памятью..."},
	}
	for _, tt := range tests {
		t.Run(tt.action, func(t *testing.T) {
			got := memoryStatusLabel(map[string]interface{}{"action": tt.action}, "ru")
			if got != tt.want {
				t.Errorf("memoryStatusLabel(%q) = %q, want %q", tt.action, got, tt.want)
			}
		})
	}
}

// --- skillStatusLabel ---

func TestSkillStatusLabel(t *testing.T) {
	tests := []struct {
		name string
		args map[string]interface{}
		want string
	}{
		{"skill_list", map[string]interface{}{"action": "skill_list"}, "Получаю список навыков..."},
		{"skill_read with name", map[string]interface{}{"action": "skill_read", "name": "github"}, "Загружаю навык... (github)"},
		{"skill_read no name", map[string]interface{}{"action": "skill_read"}, "Загружаю навык..."},
		{"unknown", map[string]interface{}{"action": "other"}, "Работаю с навыком..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := skillStatusLabel(tt.args, "ru")
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

// --- cronStatusLabel ---

func TestCronStatusLabel(t *testing.T) {
	tests := []struct {
		action string
		want   string
	}{
		{"add", "Создаю напоминание..."},
		{"list", "Получаю расписание..."},
		{"remove", "Удаляю задачу из расписания..."},
		{"unknown", "Изменяю расписание..."},
	}
	for _, tt := range tests {
		t.Run(tt.action, func(t *testing.T) {
			got := cronStatusLabel(map[string]interface{}{"action": tt.action}, "ru")
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

// --- androidStatusLabel ---

func TestAndroidStatusLabel(t *testing.T) {
	tests := []struct {
		name string
		args map[string]interface{}
		want string
	}{
		{"search_apps", map[string]interface{}{"action": "search_apps"}, "Ищу приложения..."},
		{"app_info with pkg", map[string]interface{}{"action": "app_info", "package_name": "com.example"}, "Получаю сведения о приложении... (com.example)"},
		{"app_info no pkg", map[string]interface{}{"action": "app_info"}, "Получаю сведения о приложении..."},
		{"launch_app with pkg", map[string]interface{}{"action": "launch_app", "package_name": "com.test"}, "Открываю приложение... (com.test)"},
		{"launch_app no pkg", map[string]interface{}{"action": "launch_app"}, "Открываю приложение..."},
		{"screenshot", map[string]interface{}{"action": "screenshot"}, "Делаю снимок экрана..."},
		{"get_ui_tree", map[string]interface{}{"action": "get_ui_tree"}, "Получаю элементы экрана..."},
		{"tap", map[string]interface{}{"action": "tap"}, "Нажимаю..."},
		{"swipe", map[string]interface{}{"action": "swipe"}, "Прокручиваю..."},
		{"text", map[string]interface{}{"action": "text"}, "Ввожу текст..."},
		{"keyevent with key", map[string]interface{}{"action": "keyevent", "key": "BACK"}, "Нажимаю клавишу... (BACK)"},
		{"keyevent no key", map[string]interface{}{"action": "keyevent"}, "Нажимаю клавишу..."},
		{"broadcast", map[string]interface{}{"action": "broadcast"}, "Отправляю системное событие..."},
		{"intent", map[string]interface{}{"action": "intent"}, "Отправляю системную команду..."},
		{"unknown", map[string]interface{}{"action": "other"}, "Управляю устройством..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := androidStatusLabel(tt.args, "ru")
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

// --- mcpStatusLabel ---

func TestMcpStatusLabel(t *testing.T) {
	tests := []struct {
		name string
		args map[string]interface{}
		want string
	}{
		{"mcp_list", map[string]interface{}{"action": "mcp_list"}, "Получаю список серверов MCP..."},
		{"mcp_tools with server", map[string]interface{}{"action": "mcp_tools", "server": "myserver"}, "Получаю инструменты MCP... (myserver)"},
		{"mcp_tools no server", map[string]interface{}{"action": "mcp_tools"}, "Получаю инструменты MCP..."},
		{"mcp_call tool+server", map[string]interface{}{"action": "mcp_call", "tool": "mytool", "server": "srv"}, "Выполняю инструмент MCP... (srv/mytool)"},
		{"mcp_call tool only", map[string]interface{}{"action": "mcp_call", "tool": "mytool"}, "Выполняю инструмент MCP... (mytool)"},
		{"mcp_call no args", map[string]interface{}{"action": "mcp_call"}, "Выполняю инструмент MCP..."},
		{"unknown", map[string]interface{}{"action": "other"}, "Работаю с MCP..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mcpStatusLabel(tt.args, "ru")
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func containsStr(s, substr string) bool {
	return len(substr) == 0 || len(s) >= len(substr) && (s == substr || len(s) > 0 && containsSubstring(s, substr))
}

func containsSubstring(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
