package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadSettingsDefaultsForMissingFile(t *testing.T) {
	got := loadSettingsFrom(filepath.Join(t.TempDir(), "missing.json"))
	want := defaultAppSettings()
	if got != want {
		t.Fatalf("defaults mismatch: got %+v want %+v", got, want)
	}
}

func TestLoadSettingsBackfillsNewTimerFields(t *testing.T) {
	path := filepath.Join(t.TempDir(), "settings.json")
	if err := os.WriteFile(path, []byte(`{"language":"en"}`), 0o644); err != nil {
		t.Fatal(err)
	}
	got := loadSettingsFrom(path)
	if got.Language != LangEN || got.FocusMinutes != defaultFocusMinutes || got.AlarmMinutes != defaultAlarmMinutes {
		t.Fatalf("unexpected backfill result: %+v", got)
	}
}

func TestSettingsRoundTripAndClamp(t *testing.T) {
	path := filepath.Join(t.TempDir(), "settings.json")
	in := appSettings{
		Language:           LangEN,
		FocusMinutes:       999,
		AlarmMinutes:       9999,
		StartupWithWindows: true,
	}
	if err := saveSettingsTo(path, in); err != nil {
		t.Fatal(err)
	}
	got := loadSettingsFrom(path)
	if got.Language != LangEN || got.FocusMinutes != maxFocusMinutes || got.AlarmMinutes != maxAlarmMinutes || !got.StartupWithWindows {
		t.Fatalf("unexpected round trip: %+v", got)
	}
}
