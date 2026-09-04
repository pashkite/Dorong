package main

import "testing"

func TestHangPoseLeftEdgeCentersOnWindowEdge(t *testing.T) {
	support := ScreenRect{Left: 400, Top: 300, Right: 1000, Bottom: 800}
	work := ScreenRect{Left: 0, Top: 0, Right: 1920, Bottom: 1040}
	x, y := hangPose(support, work, 236, 236, -1)
	if x != 282 { // 400 - 118
		t.Fatalf("x=%d, want 282", x)
	}
	if y != 108 { // 300 - 236 + 44
		t.Fatalf("y=%d, want 108", y)
	}
}

func TestHangPoseRightEdgeCentersOnWindowEdge(t *testing.T) {
	support := ScreenRect{Left: 400, Top: 300, Right: 1000, Bottom: 800}
	work := ScreenRect{Left: 0, Top: 0, Right: 1920, Bottom: 1040}
	x, _ := hangPose(support, work, 236, 236, 1)
	if x != 882 { // 1000 - 118
		t.Fatalf("x=%d, want 882", x)
	}
}

func TestHangPoseTracksMovedWindow(t *testing.T) {
	work := ScreenRect{Left: 0, Top: 0, Right: 1920, Bottom: 1040}
	before := ScreenRect{Left: 300, Top: 250, Right: 900, Bottom: 700}
	after := ScreenRect{Left: 430, Top: 330, Right: 1030, Bottom: 780}
	x1, y1 := hangPose(before, work, 236, 236, -1)
	x2, y2 := hangPose(after, work, 236, 236, -1)
	if x2-x1 != 130 || y2-y1 != 80 {
		t.Fatalf("hang pose did not follow moved support: before=%d,%d after=%d,%d", x1, y1, x2, y2)
	}
}

func TestHangPoseKeepsHalfPetReachableAtScreenEdge(t *testing.T) {
	work := ScreenRect{Left: 0, Top: 0, Right: 1920, Bottom: 1040}
	leftSupport := ScreenRect{Left: -500, Top: 200, Right: 400, Bottom: 700}
	x, _ := hangPose(leftSupport, work, 236, 236, -1)
	if x != -118 {
		t.Fatalf("left clamp x=%d, want -118", x)
	}

	rightSupport := ScreenRect{Left: 1500, Top: 200, Right: 2500, Bottom: 700}
	x, _ = hangPose(rightSupport, work, 236, 236, 1)
	if x != 1802 { // 1920 - 118
		t.Fatalf("right clamp x=%d, want 1802", x)
	}
}

func TestHangPoseAvoidsBecomingMostlyHiddenAboveTopEdge(t *testing.T) {
	work := ScreenRect{Left: 0, Top: 0, Right: 1920, Bottom: 1040}
	support := ScreenRect{Left: 300, Top: 0, Right: 900, Bottom: 700}
	_, y := hangPose(support, work, 236, 236, -1)
	if y != -118 {
		t.Fatalf("y=%d, want -118 so at least half remains reachable", y)
	}
}

func TestHangReleaseMovesOutwardAndDown(t *testing.T) {
	x, y := hangReleasePosition(500, 200, -1)
	if x != 490 || y != 207 {
		t.Fatalf("left release=%d,%d want 490,207", x, y)
	}
	x, y = hangReleasePosition(500, 200, 1)
	if x != 510 || y != 207 {
		t.Fatalf("right release=%d,%d want 510,207", x, y)
	}
}
