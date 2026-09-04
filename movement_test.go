package main

import "testing"

func TestNextWalkXMovesRightWithoutOvershoot(t *testing.T) {
	next, arrived := nextWalkX(100, 105, 2, 0, 500)
	if next != 102 || arrived {
		t.Fatalf("next=%d arrived=%v", next, arrived)
	}
	next, arrived = nextWalkX(104, 105, 2, 0, 500)
	if next != 105 || !arrived {
		t.Fatalf("final next=%d arrived=%v", next, arrived)
	}
}

func TestNextWalkXMovesLeftWithoutOvershoot(t *testing.T) {
	next, arrived := nextWalkX(105, 100, 2, 0, 500)
	if next != 103 || arrived {
		t.Fatalf("next=%d arrived=%v", next, arrived)
	}
	next, arrived = nextWalkX(101, 100, 2, 0, 500)
	if next != 100 || !arrived {
		t.Fatalf("final next=%d arrived=%v", next, arrived)
	}
}

func TestNextWalkXSupportsTargetAtZero(t *testing.T) {
	next, arrived := nextWalkX(2, 0, 2, 0, 500)
	if next != 0 || !arrived {
		t.Fatalf("next=%d arrived=%v; x=0 must be a valid destination", next, arrived)
	}
}

func TestNextWalkXClampsTargetAfterWindowResize(t *testing.T) {
	next, arrived := nextWalkX(118, 300, 2, 40, 120)
	if next != 120 || !arrived {
		t.Fatalf("next=%d arrived=%v, want 120,true", next, arrived)
	}
}
