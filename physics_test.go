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
