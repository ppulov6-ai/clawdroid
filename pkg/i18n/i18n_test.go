package i18n

import (
	"regexp"
	"testing"
)

func TestNormalizeLocale(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", "en"},
		{"en", "en"},
		{"ru", "ru"},
		{"ru-RU", "ru"},
		{"en_US", "en"},
		{"EN", "en"},
		{"RU", "ru"},
		{"ru_RU", "ru"},
		{"  ru  ", "ru"},
		// Accept-Language header formats
		{"ru, en;q=0.9", "ru"},
		{"en-US,en;q=0.5", "en"},
		{"ru-RU, en-US;q=0.8, fr;q=0.5", "ru"},
		{"en;q=1.0", "en"},
	}
	for _, tt := range tests {
		got := NormalizeLocale(tt.input)
		if got != tt.want {
			t.Errorf("NormalizeLocale(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestT(t *testing.T) {
	// English
	got := T("en", "status.thinking")
	if got != "Thinking..." {
		t.Errorf("T(en, status.thinking) = %q", got)
	}

	// Russian
	got = T("ru", "status.thinking")
	if got != "Думаю..." {
		t.Errorf("T(ru, status.thinking) = %q", got)
	}

	// Fallback to en for unknown locale
	got = T("fr", "status.thinking")
	if got != "Thinking..." {
		t.Errorf("T(fr, status.thinking) = %q, want English fallback", got)
	}

	// Unknown key returns the key
	got = T("en", "nonexistent.key")
	if got != "nonexistent.key" {
		t.Errorf("T(en, nonexistent.key) = %q, want key itself", got)
	}
}

func TestTf(t *testing.T) {
	got := Tf("en", "status.searching_q", "golang")
	if got != "Searching... (golang)" {
		t.Errorf("Tf(en, status.searching_q, golang) = %q", got)
	}

	got = Tf("ru", "status.searching_q", "golang")
	if got != "Ищу... (golang)" {
		t.Errorf("Tf(ru, status.searching_q, golang) = %q", got)
	}
}

func TestConfigLabels(t *testing.T) {
	// Russian config label (namespaced with "config." prefix)
	got := T("ru", "config.Model")
	if got != "Модель" {
		t.Errorf("T(ru, config.Model) = %q, want Модель", got)
	}

	// English config label returns the struct tag value
	got = T("en", "config.Model")
	if got != "Model" {
		t.Errorf("T(en, config.Model) = %q, want Model", got)
	}
}

func TestAgentMessages(t *testing.T) {
	got := T("en", "agent.context_window_warning")
	if got == "agent.context_window_warning" {
		t.Error("expected English warning message, got key itself")
	}

	got = T("ru", "agent.context_window_warning")
	if got == "agent.context_window_warning" {
		t.Error("expected Russian warning message, got key itself")
	}
}

// TestFormatSpecifierConsistency verifies that en and ru translations have
// matching format specifiers (%s, %d, etc.) to prevent runtime panics in Tf().
func TestFormatSpecifierConsistency(t *testing.T) {
	re := regexp.MustCompile(`%[sdvfgqxobt]`)

	enMessages := messages["en"]
	ruMessages := messages["ru"]

	for key, enVal := range enMessages {
		ruVal, ok := ruMessages[key]
		if !ok {
			t.Errorf("missing Russian translation for %q", key)
			continue
		}

		enSpecs := re.FindAllString(enVal, -1)
		ruSpecs := re.FindAllString(ruVal, -1)

		if len(enSpecs) != len(ruSpecs) {
			t.Errorf("format specifier count mismatch for key %q: en has %d (%v), ru has %d (%v)",
				key, len(enSpecs), enSpecs, len(ruSpecs), ruSpecs)
		}
	}
}

func TestRussianCatalogCoverage(t *testing.T) {
	if len(messages["ru"]) != len(messages["en"]) {
		t.Errorf("catalog size mismatch: ru=%d, en=%d", len(messages["ru"]), len(messages["en"]))
	}
	for key := range messages["en"] {
		if value, ok := messages["ru"][key]; !ok || value == "" {
			t.Errorf("missing Russian translation for %q", key)
		}
	}
}
