package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type appSettings struct {
	Language Language `json:"language"`
}

func settingsPath() string {
	dir, err := os.UserConfigDir()
	if err != nil || dir == "" {
		return "Dorong-settings.json"
	}
	return filepath.Join(dir, "Dorong", "settings.json")
}

func loadSettingsFrom(path string) appSettings {
	s := appSettings{Language: LangKO}
	data, err := os.ReadFile(path)
	if err != nil {
		return s
	}
	if json.Unmarshal(data, &s) != nil {
		return appSettings{Language: LangKO}
	}
	s.Language = normalizeLanguage(string(s.Language))
	return s
}

func saveSettingsTo(path string, s appSettings) error {
	s.Language = normalizeLanguage(string(s.Language))
	if dir := filepath.Dir(path); dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func loadSettings() {
	s := loadSettingsFrom(settingsPath())
	setLanguage(s.Language)
}

func saveSettings() error {
	return saveSettingsTo(settingsPath(), appSettings{Language: currentLanguage})
}
