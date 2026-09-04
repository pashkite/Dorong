package main

import (
	"path/filepath"
	"testing"
)

func TestNormalizeLanguage(t *testing.T) {
	if normalizeLanguage("en") != LangEN {
		t.Fatal("en should select English")
	}
	if normalizeLanguage("ko") != LangKO || normalizeLanguage("unknown") != LangKO {
		t.Fatal("Korean should be the safe default")
	}
}

func TestTranslationChangesWithLanguage(t *testing.T) {
	old := currentLanguage
	defer setLanguage(old)
	setLanguage(LangKO)
	ko := tr("greeting")
	setLanguage(LangEN)
	en := tr("greeting")
	if ko == en || en != "Hi! I'm Dorong." {
		t.Fatalf("unexpected translations: ko=%q en=%q", ko, en)
	}
}

func TestLanguageSettingPersists(t *testing.T) {
	path := filepath.Join(t.TempDir(), "settings.json")
	if err := saveSettingsTo(path, appSettings{Language: LangEN}); err != nil {
		t.Fatal(err)
	}
	got := loadSettingsFrom(path)
	if got.Language != LangEN {
		t.Fatalf("language = %q, want en", got.Language)
	}
}

func TestInvalidStoredLanguageFallsBackToKorean(t *testing.T) {
	path := filepath.Join(t.TempDir(), "settings.json")
	if err := saveSettingsTo(path, appSettings{Language: Language("xx")}); err != nil {
		t.Fatal(err)
	}
	if got := loadSettingsFrom(path); got.Language != LangKO {
		t.Fatalf("language = %q, want ko", got.Language)
	}
}

func TestTranslationKeyParity(t *testing.T) {
	for key := range translations[LangKO] {
		if _, ok := translations[LangEN][key]; !ok {
			t.Fatalf("English translation missing key %q", key)
		}
	}
	for key := range translations[LangEN] {
		if _, ok := translations[LangKO][key]; !ok {
			t.Fatalf("Korean translation missing key %q", key)
		}
	}
}
