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
		"menu.greet":          "도롱아, 인사해",
		"menu.focus":          "25분 집중 시작",
		"menu.alarm":          "10분 알람 설정",
		"menu.sleep":          "잠깐 재우기",
		"menu.jump":           "점프시키기",
		"menu.wander":         "혼자 돌아다니기",
		"menu.topmost":        "항상 위에 표시",
		"menu.home":           "오른쪽 아래로 이동",
		"menu.lang_ko":        "언어: 한국어",
		"menu.lang_en":        "언어: English",
		"menu.exit":           "Dorong 종료",
		"greeting":            "안녕! 나는 도롱이야.",
		"pet_count":           "쓰담쓰담 %d번!",
		"happy":               "헤헤, 좋아!",
		"ticklish":            "간지러워!",
		"focus_started":       "25분 집중 시작!",
		"alarm_set":           "10분 뒤에 알려줄게.",
		"sleep":               "Zzz…",
		"jump":                "도롱!",
		"random.0":            "뭐 하고 있어?",
		"random.1":            "나 여기 있어.",
		"random.2":            "물 한 모금 어때?",
		"random.3":            "오늘도 같이 있자.",
		"random.4":            "하나씩 해보자!",
		"random.5":            "쓰담쓰담 대기 중!",
		"random.6":            "도롱도롱…",
		"sleep_random":        "조금만 잘게…",
		"hang":                "앗…!",
		"focus_done":          "집중 끝! 수고했어.",
		"focus_dialog":        "25분 집중이 끝났어! 잠깐 쉬자.",
		"focus_progress":      "집중 중 · %02d:%02d",
		"alarm_done":          "알람이야! 시간이 됐어.",
		"alarm_dialog":        "10분 알람 시간이 됐어!",
		"language_changed_ko": "한국어로 바꿨어!",
		"language_changed_en": "영어로 바꿨어!",
	},
	LangEN: {
		"menu.greet":          "Say hello",
		"menu.focus":          "Start 25-minute focus",
		"menu.alarm":          "Set 10-minute alarm",
		"menu.sleep":          "Take a short nap",
		"menu.jump":           "Jump",
		"menu.wander":         "Wander around",
		"menu.topmost":        "Always on top",
		"menu.home":           "Move to bottom-right",
		"menu.lang_ko":        "Language: 한국어",
		"menu.lang_en":        "Language: English",
		"menu.exit":           "Exit Dorong",
		"greeting":            "Hi! I'm Dorong.",
		"pet_count":           "Pets: %d!",
		"happy":               "Hehe, I like that!",
		"ticklish":            "That tickles!",
		"focus_started":       "25-minute focus started!",
		"alarm_set":           "I'll tell you in 10 minutes.",
		"sleep":               "Zzz…",
		"jump":                "Dorong!",
		"random.0":            "What are you up to?",
		"random.1":            "I'm right here.",
		"random.2":            "How about some water?",
		"random.3":            "Let's hang out today too.",
		"random.4":            "One thing at a time!",
		"random.5":            "Ready for head pats!",
		"random.6":            "Dorong dorong…",
		"sleep_random":        "Just a tiny nap…",
		"hang":                "Whoa…!",
		"focus_done":          "Focus done! Nice work.",
		"focus_dialog":        "Your 25-minute focus session is over. Take a short break!",
		"focus_progress":      "Focus · %02d:%02d",
		"alarm_done":          "Alarm! Time's up.",
		"alarm_dialog":        "Your 10-minute alarm is up!",
		"language_changed_ko": "Switched to Korean!",
		"language_changed_en": "Switched to English!",
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
