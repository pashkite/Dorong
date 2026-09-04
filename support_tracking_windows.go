//go:build windows

package main

var (
	trackedSupportHwnd uintptr
	trackedSupportRect RECT
	trackedSupportOK   bool
)

func rememberSupportWindow(hwnd uintptr, r RECT) {
	trackedSupportHwnd = hwnd
	trackedSupportRect = r
	trackedSupportOK = hwnd != 0
}

func clearSupportWindowTracking() {
	trackedSupportHwnd = 0
	trackedSupportRect = RECT{}
	trackedSupportOK = false
}

// updateSupportMotion runs on a dedicated ~60 Hz timer. The normal 40 ms pet
// tick keeps gameplay/animation timing unchanged, while this lightweight path
// only follows the window Dorong is standing on or hanging from.
func updateSupportMotion() {
	if pet.supportHwnd == 0 {
		clearSupportWindowTracking()
		return
	}
	if pet.dragging || pet.falling {
		return
	}

	r, ok := validSupportWindow(pet.supportHwnd)
	if !ok {
		startFall(0)
		clearSupportWindowTracking()
		return
	}

	if pet.hanging {
		if !syncHangPose() {
			startFall(0)
			clearSupportWindowTracking()
			return
		}
		rememberSupportWindow(pet.supportHwnd, r)
		return
	}

	x, y := currentPos()
	wa := workAreaForWindow(pet.supportHwnd)
	area := ScreenRect{Left: wa.Left, Top: wa.Top, Right: wa.Right, Bottom: wa.Bottom}
	current := ScreenRect{Left: r.Left, Top: r.Top, Right: r.Right, Bottom: r.Bottom}

	if !trackedSupportOK || trackedSupportHwnd != pet.supportHwnd {
		rememberSupportWindow(pet.supportHwnd, r)
		lo, hi := supportXBounds(current, area, PET_W)
		nx := clampSupportX(x, lo, hi)
		ny := r.Top - PET_H
		if nx != x || ny != y {
			setPos(nx, ny)
		}
		return
	}

	previous := ScreenRect{Left: trackedSupportRect.Left, Top: trackedSupportRect.Top, Right: trackedSupportRect.Right, Bottom: trackedSupportRect.Bottom}
	nx, nextTarget := followSupportX(x, pet.targetX, pet.walking, previous, current, area, PET_W)
	if pet.walking {
		pet.targetX = nextTarget
	}
	ny := r.Top - PET_H
	if nx != x || ny != y {
		setPos(nx, ny)
	}
	rememberSupportWindow(pet.supportHwnd, r)
}
