package main

const supportFootInset int32 = 42

// supportXBounds returns the horizontal range in which Dorong's visible feet
// remain over a supporting application window while also staying on the
// monitor's usable work area.
func supportXBounds(support, area ScreenRect, petWidth int32) (lo, hi int32) {
	lo = support.Left - petWidth/2 + supportFootInset
	hi = support.Right - petWidth/2 - supportFootInset
	if lo < area.Left {
		lo = area.Left
	}
	maxX := area.Right - petWidth
	if hi > maxX {
		hi = maxX
	}
	if hi < lo {
		hi = lo
	}
	return lo, hi
}

func clampSupportX(v, lo, hi int32) int32 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

// followSupportX keeps Dorong visually attached to a moving support window.
// Pure translations shift Dorong and an active walking target by the same
// amount. During resize, Dorong keeps its screen X where possible and is only
// clamped back onto the resized window, which avoids sideways jumps when a
// single edge is dragged.
func followSupportX(petX, targetX int32, walking bool, previous, current, area ScreenRect, petWidth int32) (x, nextTarget int32) {
	x = petX
	nextTarget = targetX
	oldWidth := previous.Right - previous.Left
	newWidth := current.Right - current.Left
	if oldWidth == newWidth {
		dx := current.Left - previous.Left
		x += dx
		if walking {
			nextTarget += dx
		}
	}

	lo, hi := supportXBounds(current, area, petWidth)
	x = clampSupportX(x, lo, hi)
	if walking {
		nextTarget = clampSupportX(nextTarget, lo, hi)
	}
	return x, nextTarget
}
