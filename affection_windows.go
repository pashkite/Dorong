//go:build windows

package main

import "time"

const ID_AFFECTION = 1013

func syncAffectionFromSettings() {
	pet.petCount = normalizePetCount(persistentPetCount)
}

func handleHeadPat() {
	oldLevel := affectionLevelFor(pet.petCount)
	pet.petCount++
	persistentPetCount = pet.petCount
	_ = saveSettings()

	newLevel := affectionLevelFor(pet.petCount)
	pet.happyUntil = time.Now().Add(2300 * time.Millisecond)

	if newLevel > oldLevel {
		showBubble(tr("affection.level_up", tr(affectionLevelKey(newLevel))), 3*time.Second)
		return
	}
	if pet.petCount%10 == 0 {
		showBubble(tr("pet_count", pet.petCount), 2200*time.Millisecond)
		return
	}
	showBubble(tr(affectionReactionKey(newLevel)), 2*time.Second)
}

func showAffectionStatus() {
	level := affectionLevelFor(pet.petCount)
	levelName := tr(affectionLevelKey(level))
	next := affectionNextThreshold(pet.petCount)
	if next == 0 {
		showBubble(tr("affection.status_max", levelName, pet.petCount), 4*time.Second)
		return
	}
	showBubble(tr("affection.status", levelName, pet.petCount, next), 4*time.Second)
}
