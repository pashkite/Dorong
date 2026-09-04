package main

import "math"

const (
	fallGravity       = 1.15
	fallTerminalSpeed = 19.0
)

// nextFallStep advances vertical velocity by one timer tick and returns the
// rounded pixel delta for that tick. Keeping this platform-neutral makes the
// desktop-pet gravity easy to test without a Windows GUI session.
func nextFallStep(v float64) (nextVelocity float64, deltaY int32) {
	nextVelocity = math.Min(v+fallGravity, fallTerminalSpeed)
	deltaY = int32(math.Round(nextVelocity))
	return nextVelocity, deltaY
}

// resolveFloorLanding snaps the pet to the work-area floor once its next
// position would cross it.
func resolveFloorLanding(nextY, petHeight, floorBottom int32) (resolvedY int32, landed bool) {
	floorY := floorBottom - petHeight
	if nextY >= floorY {
		return floorY, true
	}
	return nextY, false
}
