package broadcast

import "testing"

func TestTargetPackage(t *testing.T) {
 t.Setenv("CLAWDROID_ANDROID_PACKAGE", "io.jarvisjon.android.debug")
 if got := targetPackage(); got != "io.jarvisjon.android.debug" { t.Fatalf("wrong target: %s", got) }
 t.Setenv("CLAWDROID_ANDROID_PACKAGE", "")
 if got := targetPackage(); got != Package { t.Fatalf("wrong fallback: %s", got) }
}
