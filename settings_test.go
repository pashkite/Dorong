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

func TestLoadSettingsBackfillsNewFields(t *testing.T) {
	path := filepath.Join(t.TempDir(), "settings.json")
	data := []byte("{\"language\":\"en\"}")
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatal(err)
	}
	got := loadSettingsFrom(path)
	if got.Language != LangEN || got.FocusMinutes != defaultFocusMinutes || got.AlarmMinutes != defaultAlarmMinutes || got.PetCount != 0 {
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
		PetCount:           87,
	}
	if err := saveSettingsTo(path, in); err != nil {
		t.Fatal(err)
	}
	got := loadSettingsFrom(path)
	if got.Language != LangEN || got.FocusMinutes != maxFocusMinutes || got.AlarmMinutes != maxAlarmMinutes || !got.StartupWithWindows || got.PetCount != 87 {
		t.Fatalf("unexpected round trip: %+v", got)
	}
}

func TestSettingsClampNegativePetCount(t *testing.T) {
	path := filepath.Join(t.TempDir(), "settings.json")
	if err := saveSettingsTo(path, appSettings{PetCount: -100}); err != nil {
		t.Fatal(err)
	}
	got := loadSettingsFrom(path)
	if got.PetCount != 0 {
		t.Fatalf("pet count = %d, want 0", got.PetCount)
	}
}
