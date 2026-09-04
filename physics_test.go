package main

import "testing"

func TestNextFallStepStartsFallingAfterRelease(t *testing.T) {
	v, dy := nextFallStep(0)
	if v != 1.15 {
		t.Fatalf("velocity = %v, want 1.15", v)
	}
	if dy != 1 {
		t.Fatalf("deltaY = %d, want 1", dy)
	}
}

func TestNextFallStepAcceleratesDownward(t *testing.T) {
	v := 0.0
	previous := v
	for i := 0; i < 8; i++ {
		v, _ = nextFallStep(v)
		if v <= previous {
			t.Fatalf("step %d velocity did not increase: previous=%v current=%v", i, previous, v)
		}
		previous = v
	}
}

func TestNextFallStepCapsAtTerminalVelocity(t *testing.T) {
	v := 0.0
	for i := 0; i < 100; i++ {
		v, _ = nextFallStep(v)
	}
	if v != fallTerminalSpeed {
		t.Fatalf("velocity = %v, want terminal speed %v", v, fallTerminalSpeed)
	}
	_, dy := nextFallStep(v)
	if dy != 19 {
		t.Fatalf("terminal deltaY = %d, want 19", dy)
	}
}

func TestNegativeInitialVelocityAllowsJumpThenGravityWins(t *testing.T) {
	v, dy := nextFallStep(-6)
	if v >= 0 || dy >= 0 {
		t.Fatalf("first jump step should still move upward: velocity=%v deltaY=%d", v, dy)
	}
	for i := 0; i < 10 && v <= 0; i++ {
		v, _ = nextFallStep(v)
	}
	if v <= 0 {
		t.Fatalf("gravity never turned upward motion into downward motion: velocity=%v", v)
	}
}

func TestResolveFloorLandingSnapsExactlyToFloor(t *testing.T) {
	const petHeight int32 = 236
	const floorBottom int32 = 1040
	floorY := floorBottom - petHeight

	y, landed := resolveFloorLanding(floorY-1, petHeight, floorBottom)
	if landed || y != floorY-1 {
		t.Fatalf("premature landing: y=%d landed=%v", y, landed)
	}

	y, landed = resolveFloorLanding(floorY+12, petHeight, floorBottom)
	if !landed || y != floorY {
		t.Fatalf("landing = y:%d landed:%v, want y:%d landed:true", y, landed, floorY)
	}
}

func TestResolveWorkAreaLandingBottomTaskbar(t *testing.T) {
	area := ScreenRect{Left: 0, Top: 0, Right: 1920, Bottom: 1040} // 40 px bottom taskbar on 1080p
	x, y, landed := resolveWorkAreaLanding(1700, 900, 236, 236, area)
	if !landed {
		t.Fatal("expected landing on work-area floor")
	}
	if y != 804 { // 1040 - 236
		t.Fatalf("y = %d, want 804", y)
	}
	if x != 1684 { // 1920 - 236
		t.Fatalf("x = %d, want 1684", x)
	}
}

func TestResolveWorkAreaLandingLeftTaskbarClampsX(t *testing.T) {
	area := ScreenRect{Left: 64, Top: 0, Right: 1920, Bottom: 1080}
	x, y, landed := resolveWorkAreaLanding(0, 900, 236, 236, area)
	if !landed {
		t.Fatal("expected landing")
	}
	if x != 64 {
		t.Fatalf("x = %d, want 64 so Dorong does not overlap left taskbar", x)
	}
	if y != 844 {
		t.Fatalf("y = %d, want 844", y)
	}
}

func TestResolveWorkAreaLandingRightTaskbarClampsX(t *testing.T) {
	area := ScreenRect{Left: 0, Top: 0, Right: 1856, Bottom: 1080}
	x, _, _ := resolveWorkAreaLanding(1800, 900, 236, 236, area)
	if x != 1620 { // 1856 - 236
		t.Fatalf("x = %d, want 1620 so Dorong does not overlap right taskbar", x)
	}
}

func TestResolveWorkAreaLandingTopTaskbarStillUsesUsableFloor(t *testing.T) {
	area := ScreenRect{Left: 0, Top: 48, Right: 1920, Bottom: 1080}
	_, y, landed := resolveWorkAreaLanding(500, 1000, 236, 236, area)
	if !landed || y != 844 {
		t.Fatalf("landing = y:%d landed:%v, want y:844 landed:true", y, landed)
	}
}

func TestResolveWorkAreaLandingDoesNotLandEarly(t *testing.T) {
	area := ScreenRect{Left: 0, Top: 0, Right: 1920, Bottom: 1040}
	x, y, landed := resolveWorkAreaLanding(400, 700, 236, 236, area)
	if landed {
		t.Fatal("landed too early")
	}
	if x != 400 || y != 700 {
		t.Fatalf("position changed before landing: got %d,%d", x, y)
	}
}
