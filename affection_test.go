package main

import "testing"

func TestAffectionLevelBoundaries(t *testing.T) {
	tests := []struct {
		count int
		want  AffectionLevel
	}{
		{-3, AffectionShy},
		{0, AffectionShy},
		{9, AffectionShy},
		{10, AffectionFamiliar},
		{29, AffectionFamiliar},
		{30, AffectionFriend},
		{59, AffectionFriend},
		{60, AffectionClose},
		{99, AffectionClose},
		{100, AffectionBest},
		{500, AffectionBest},
	}
	for _, tc := range tests {
		if got := affectionLevelFor(tc.count); got != tc.want {
			t.Fatalf("count %d: level=%d want=%d", tc.count, got, tc.want)
		}
	}
}

func TestAffectionNextThreshold(t *testing.T) {
	tests := []struct {
		count int
		want  int
	}{
		{0, 10},
		{10, 30},
		{30, 60},
		{60, 100},
		{100, 0},
	}
	for _, tc := range tests {
		if got := affectionNextThreshold(tc.count); got != tc.want {
			t.Fatalf("count %d: next=%d want=%d", tc.count, got, tc.want)
		}
	}
}

func TestAffectionKeysAreStable(t *testing.T) {
	levels := []AffectionLevel{AffectionShy, AffectionFamiliar, AffectionFriend, AffectionClose, AffectionBest}
	for i, level := range levels {
		if got := affectionLevelKey(level); got != "affection.level."+string(rune('0'+i)) {
			t.Fatalf("level %d key=%q", level, got)
		}
		if got := affectionReactionKey(level); got != "affection.react."+string(rune('0'+i)) {
			t.Fatalf("level %d reaction=%q", level, got)
		}
	}
}
