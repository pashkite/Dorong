package main

import "fmt"

type Language string

const (
	LangKO Language = "ko"
	LangEN Language = "en"
)

var currentLanguage = LangKO

var translations = map[Language]map[string]string{
	LangKO: {
		"menu.greet":             "도롱아, 인사해",
		"menu.focus":             "%d분 집중 시작",
		"menu.alarm":             "%d분 알람 설정",
		"menu.sleep":             "잠깐 재우기",
		"menu.jump":              "점프시키기",
		"menu.wander":            "혼자 돌아다니기",
		"menu.topmost":           "항상 위에 표시",
		"menu.home":              "오른쪽 아래로 이동",
		"menu.affection":         "호감도 보기",
		"menu.lang_ko":           "언어: 한국어",
		"menu.lang_en":           "언어: English",
		"menu.settings":          "설정...",
		"menu.exit":              "Dorong 종료",
		"greeting":               "안녕! 나는 도롱이야.",
		"pet_count":              "쓰담쓰담 %d번!",
		"affection.level.0":      "낯가림",
		"affection.level.1":      "익숙함",
		"affection.level.2":      "친구",
		"affection.level.3":      "단짝",
		"affection.level.4":      "최고의 친구",
		"affection.level_up":     "호감도 상승! %s",
		"affection.status":       "호감도 %s · 쓰담 %d회 · 다음 %d회",
		"affection.status_max":   "호감도 %s · 쓰담 %d회!",
		"affection.react.0":      "헤헤, 좋아!",
		"affection.react.1":      "응, 이제 좀 익숙해!",
		"affection.react.2":      "또 쓰다듬어 줘!",
		"affection.react.3":      "너랑 있으면 좋아!",
		"affection.react.4":      "역시 네가 제일 좋아!",
		"happy":                  "헤헤, 좋아!",
		"ticklish":               "간지러워!",
		"focus_started":          "%d분 집중 시작!",
		"alarm_set":              "%d분 뒤에 알려줄게.",
		"sleep":                  "Zzz…",
		"jump":                   "도롱!",
		"random.0":               "뭐 하고 있어?",
		"random.1":               "나 여기 있어.",
		"random.2":               "물 한 모금 어때?",
		"random.3":               "오늘도 같이 있자.",
		"random.4":               "하나씩 해보자!",
		"random.5":               "쓰담쓰담 대기 중!",
		"random.6":               "도롱도롱…",
		"sleep_random":           "조금만 잘게…",
		"hang":                   "앗…!",
		"focus_done":             "집중 끝! 수고했어.",
		"focus_dialog":           "집중 시간이 끝났어! 잠깐 쉬자.",
		"focus_progress":         "집중 중 · %02d:%02d",
		"alarm_done":             "알람이야! 시간이 됐어.",
		"alarm_dialog":           "알람 시간이 됐어!",
		"language_changed_ko":    "한국어로 바꿨어!",
		"language_changed_en":    "영어로 바꿨어!",
		"settings.title":         "Dorong 설정",
		"settings.focus_label":   "집중 시간 (분, 1~180)",
		"settings.alarm_label":   "알람 시간 (분, 1~1440)",
		"settings.startup":       "Windows 시작 시 Dorong 자동 실행",
		"settings.save":          "저장",
		"settings.cancel":        "취소",
		"settings.saved":         "설정을 저장했어!",
		"settings.invalid":       "시간은 표시된 범위 안의 숫자로 입력해줘.",
		"settings.startup_error": "Windows 자동 실행 설정을 변경하지 못했어.",
		"settings.save_error":    "설정 파일을 저장하지 못했어.",
		"settings.open_error":    "설정창을 열지 못했어.",
	},
	LangEN: {
		"menu.greet":             "Say hello",
		"menu.focus":             "Start %d-minute focus",
		"menu.alarm":             "Set %d-minute alarm",
		"menu.sleep":             "Take a short nap",
		"menu.jump":              "Jump",
		"menu.wander":            "Wander around",
		"menu.topmost":           "Always on top",
		"menu.home":              "Move to bottom-right",
		"menu.affection":         "Show affection",
		"menu.lang_ko":           "Language: 한국어",
		"menu.lang_en":           "Language: English",
		"menu.settings":          "Settings...",
		"menu.exit":              "Exit Dorong",
		"greeting":               "Hi! I'm Dorong.",
		"pet_count":              "Pets: %d!",
		"affection.level.0":      "Shy",
		"affection.level.1":      "Familiar",
		"affection.level.2":      "Friend",
		"affection.level.3":      "Bestie",
		"affection.level.4":      "Best friend",
		"affection.level_up":     "Affection up! %s",
		"affection.status":       "Bond %s · %d pets · next %d",
		"affection.status_max":   "Bond %s · %d pets!",
		"affection.react.0":      "Hehe, I like that!",
		"affection.react.1":      "I'm getting used to you!",
		"affection.react.2":      "Pet me again!",
		"affection.react.3":      "I like being with you!",
		"affection.react.4":      "You're my favorite!",
		"happy":                  "Hehe, I like that!",
		"ticklish":               "That tickles!",
		"focus_started":          "%d-minute focus started!",
		"alarm_set":              "I'll tell you in %d minutes.",
		"sleep":                  "Zzz…",
		"jump":                   "Dorong!",
		"random.0":               "What are you up to?",
		"random.1":               "I'm right here.",
		"random.2":               "How about some water?",
		"random.3":               "Let's hang out today too.",
		"random.4":               "One thing at a time!",
		"random.5":               "Ready for head pats!",
		"random.6":               "Dorong dorong…",
		"sleep_random":           "Just a tiny nap…",
		"hang":                   "Whoa…!",
		"focus_done":             "Focus done! Nice work.",
		"focus_dialog":           "Your focus session is over. Take a short break!",
		"focus_progress":         "Focus · %02d:%02d",
		"alarm_done":             "Alarm! Time's up.",
		"alarm_dialog":           "Your alarm is up!",
		"language_changed_ko":    "Switched to Korean!",
		"language_changed_en":    "Switched to English!",
		"settings.title":         "Dorong Settings",
		"settings.focus_label":   "Focus minutes (1-180)",
		"settings.alarm_label":   "Alarm minutes (1-1440)",
		"settings.startup":       "Start Dorong with Windows",
		"settings.save":          "Save",
		"settings.cancel":        "Cancel",
		"settings.saved":         "Settings saved!",
		"settings.invalid":       "Enter a number within the shown range.",
		"settings.startup_error": "Could not change the Windows startup setting.",
		"settings.save_error":    "Could not save the settings file.",
		"settings.open_error":    "Could not open the settings window.",
	},
}

func normalizeLanguage(v string) Language {
	switch Language(v) {
	case LangEN:
		return LangEN
	default:
		return LangKO
	}
}

func setLanguage(lang Language) {
	currentLanguage = normalizeLanguage(string(lang))
}

func tr(key string, args ...any) string {
	lang := normalizeLanguage(string(currentLanguage))
	text, ok := translations[lang][key]
	if !ok {
		text = translations[LangKO][key]
	}
	if len(args) > 0 {
		return fmt.Sprintf(text, args...)
	}
	return text
}

func randomPhrases() []string {
	return []string{tr("random.0"), tr("random.1"), tr("random.2"), tr("random.3"), tr("random.4"), tr("random.5"), tr("random.6")}
}
