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

// ScreenRect is a platform-neutral rectangle used by the landing tests.
// On Windows it mirrors RECT values returned by SPI_GETWORKAREA.
type ScreenRect struct {
	Left, Top, Right, Bottom int32
}

// resolveWorkAreaLanding keeps Dorong inside the usable desktop work area.
// Because SPI_GETWORKAREA excludes a docked taskbar, this works for taskbars
// docked to the bottom, top, left, or right without special-casing their side.
func resolveWorkAreaLanding(nextX, nextY, petWidth, petHeight int32, area ScreenRect) (x, y int32, landed bool) {
	maxX := area.Right - petWidth
	if maxX < area.Left {
		maxX = area.Left
	}
	x = nextX
	if x < area.Left {
		x = area.Left
	}
	if x > maxX {
		x = maxX
	}

	y, landed = resolveFloorLanding(nextY, petHeight, area.Bottom)
	return x, y, landed
}
