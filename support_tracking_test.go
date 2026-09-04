package main

import "testing"

func TestFollowSupportXPureMove(t *testing.T) {
	area := ScreenRect{Left: 0, Top: 0, Right: 1920, Bottom: 1040}
	oldRect := ScreenRect{Left: 300, Top: 200, Right: 1100, Bottom: 800}
	newRect := ScreenRect{Left: 420, Top: 260, Right: 1220, Bottom: 860}
	x, target := followSupportX(520, 760, true, oldRect, newRect, area, 236)
	if x != 640 || target != 880 {
		t.Fatalf("pure move = (%d,%d), want (640,880)", x, target)
	}
}

func TestFollowSupportXResizeDoesNotJump(t *testing.T) {
	area := ScreenRect{Left: 0, Top: 0, Right: 1920, Bottom: 1040}
	oldRect := ScreenRect{Left: 300, Top: 200, Right: 1100, Bottom: 800}
	newRect := ScreenRect{Left: 420, Top: 200, Right: 1100, Bottom: 800}
	x, _ := followSupportX(650, 0, false, oldRect, newRect, area, 236)
	if x != 650 {
		t.Fatalf("left-edge resize moved pet to %d, want 650", x)
	}
}

func TestFollowSupportXResizeClampsToWindow(t *testing.T) {
	area := ScreenRect{Left: 0, Top: 0, Right: 1920, Bottom: 1040}
	oldRect := ScreenRect{Left: 300, Top: 200, Right: 1300, Bottom: 800}
	newRect := ScreenRect{Left: 300, Top: 200, Right: 700, Bottom: 800}
	x, _ := followSupportX(1000, 0, false, oldRect, newRect, area, 236)
	_, hi := supportXBounds(newRect, area, 236)
	if x != hi {
		t.Fatalf("narrow resize x=%d, want clamp %d", x, hi)
	}
}

func TestSupportXBoundsNegativeMonitor(t *testing.T) {
	area := ScreenRect{Left: -1920, Top: 0, Right: 0, Bottom: 1040}
	support := ScreenRect{Left: -1600, Top: 180, Right: -500, Bottom: 900}
	lo, hi := supportXBounds(support, area, 236)
	if lo >= 0 || hi >= 0 || hi <= lo {
		t.Fatalf("negative-monitor bounds = (%d,%d), want ordered negative coordinates", lo, hi)
	}
}
