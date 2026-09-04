package main

const hangDropPixels int32 = 44

// hangPose returns the desktop-pet window position while Dorong is hanging
// from the left or right edge of a supporting application window.
// side < 0 means left edge; side >= 0 means right edge.
func hangPose(support, workArea ScreenRect, petWidth, petHeight int32, side int) (x, y int32) {
	if petWidth < 1 {
		petWidth = 1
	}
	if petHeight < 1 {
		petHeight = 1
	}

	edgeX := support.Right
	if side < 0 {
		edgeX = support.Left
	}

	// Center Dorong on the window edge. The normal standing range keeps the
	// feet about 42px inside the support; hanging shifts that contact point to
	// the actual edge so the body visibly moves outward.
	x = edgeX - petWidth/2
	y = support.Top - petHeight + hangDropPixels

	// Keep at least half of Dorong reachable on-screen while still allowing a
	// convincing partial overhang when the supporting window touches a screen edge.
	minX := workArea.Left - petWidth/2
	maxX := workArea.Right - petWidth/2
	if maxX < minX {
		maxX = minX
	}
	if x < minX {
		x = minX
	}
	if x > maxX {
		x = maxX
	}

	minY := workArea.Top - petHeight/2
	maxY := workArea.Bottom - petHeight
	if maxY < minY {
		maxY = minY
	}
	if y < minY {
		y = minY
	}
	if y > maxY {
		y = maxY
	}
	return x, y
}

// hangReleasePosition gives Dorong a small outward/downward nudge before
// gravity takes over, making the release from an edge read as an actual drop.
func hangReleasePosition(x, y int32, side int) (int32, int32) {
	if side < 0 {
		x -= 10
	} else {
		x += 10
	}
	return x, y + 7
}
