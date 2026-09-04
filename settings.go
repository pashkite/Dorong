package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	defaultFocusMinutes = 25
	defaultAlarmMinutes = 10
	maxFocusMinutes     = 180
	maxAlarmMinutes     = 1440
)

type appSettings struct {
	Language           Language `json:"language"`
	FocusMinutes       int      `json:"focus_minutes"`
	AlarmMinutes       int      `json:"alarm_minutes"`
	StartupWithWindows bool     `json:"startup_with_windows"`
	PetCount           int      `json:"pet_count"`
}

var (
	focusMinutes       = defaultFocusMinutes
	alarmMinutes       = defaultAlarmMinutes
	startupWithWindows bool
	persistentPetCount int
)

func defaultAppSettings() appSettings {
	return appSettings{
		Language:     LangKO,
		FocusMinutes: defaultFocusMinutes,
		AlarmMinutes: defaultAlarmMinutes,
		PetCount:     0,
	}
}

func normalizeMinutes(v, fallback, maxValue int) int {
	if v <= 0 {
		return fallback
	}
	if v > maxValue {
		return maxValue
	}
	return v
}

func normalizeAppSettings(s appSettings) appSettings {
	s.Language = normalizeLanguage(string(s.Language))
	s.FocusMinutes = normalizeMinutes(s.FocusMinutes, defaultFocusMinutes, maxFocusMinutes)
	s.AlarmMinutes = normalizeMinutes(s.AlarmMinutes, defaultAlarmMinutes, maxAlarmMinutes)
	s.PetCount = normalizePetCount(s.PetCount)
	return s
}

func settingsPath() string {
	dir, err := os.UserConfigDir()
	if err != nil || dir == "" {
		return "Dorong-settings.json"
	}
	return filepath.Join(dir, "Dorong", "settings.json")
}

func loadSettingsFrom(path string) appSettings {
	s := defaultAppSettings()
	data, err := os.ReadFile(path)
	if err != nil {
		return s
	}
	if json.Unmarshal(data, &s) != nil {
		return defaultAppSettings()
	}
	return normalizeAppSettings(s)
}

func saveSettingsTo(path string, s appSettings) error {
	s = normalizeAppSettings(s)
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
	focusMinutes = s.FocusMinutes
	alarmMinutes = s.AlarmMinutes
	startupWithWindows = s.StartupWithWindows
	persistentPetCount = s.PetCount
}

func saveSettings() error {
	return saveSettingsTo(settingsPath(), appSettings{
		Language:           currentLanguage,
		FocusMinutes:       focusMinutes,
		AlarmMinutes:       alarmMinutes,
		StartupWithWindows: startupWithWindows,
		PetCount:           persistentPetCount,
	})
}
