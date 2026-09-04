package main

// nextWalkX advances a horizontal walk toward target without overshooting.
// The target is first constrained to the current support range so a moved or
// resized application window cannot leave Dorong chasing an unreachable point.
func nextWalkX(current, target, speed, lo, hi int32) (next int32, arrived bool) {
	if hi < lo {
		hi = lo
	}
	if target < lo {
		target = lo
	}
	if target > hi {
		target = hi
	}
	if speed < 1 {
		speed = 1
	}
	if current < lo {
		current = lo
	}
	if current > hi {
		current = hi
	}
	if current == target {
		return target, true
	}
	if current < target {
		next = current + speed
		if next >= target {
			return target, true
		}
		return next, false
	}
	next = current - speed
	if next <= target {
		return target, true
	}
	return next, false
}
