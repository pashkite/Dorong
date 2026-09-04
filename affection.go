package main

type AffectionLevel int

const (
	AffectionShy AffectionLevel = iota
	AffectionFamiliar
	AffectionFriend
	AffectionClose
	AffectionBest
)

func normalizePetCount(v int) int {
	if v < 0 {
		return 0
	}
	return v
}

func affectionLevelFor(count int) AffectionLevel {
	count = normalizePetCount(count)
	switch {
	case count >= 100:
		return AffectionBest
	case count >= 60:
		return AffectionClose
	case count >= 30:
		return AffectionFriend
	case count >= 10:
		return AffectionFamiliar
	default:
		return AffectionShy
	}
}

func affectionNextThreshold(count int) int {
	switch affectionLevelFor(count) {
	case AffectionShy:
		return 10
	case AffectionFamiliar:
		return 30
	case AffectionFriend:
		return 60
	case AffectionClose:
		return 100
	default:
		return 0
	}
}

func affectionLevelKey(level AffectionLevel) string {
	switch level {
	case AffectionFamiliar:
		return "affection.level.1"
	case AffectionFriend:
		return "affection.level.2"
	case AffectionClose:
		return "affection.level.3"
	case AffectionBest:
		return "affection.level.4"
	default:
		return "affection.level.0"
	}
}

func affectionReactionKey(level AffectionLevel) string {
	switch level {
	case AffectionFamiliar:
		return "affection.react.1"
	case AffectionFriend:
		return "affection.react.2"
	case AffectionClose:
		return "affection.react.3"
	case AffectionBest:
		return "affection.react.4"
	default:
		return "affection.react.0"
	}
}
