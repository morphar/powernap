package powernap

import "time"

type target struct {
	target  time.Time
	d       time.Duration
	elapsed time.Duration
}

func newTarget(d time.Duration) target {
	// If the duration is less than the cost of 2 time.Now() calls,
	// There's no reason to do any calculations.
	if d < (2 * timeNowCost) {
		return target{
			d: d,
		}
	}

	return target{
		target: time.Now().Add(d - timeNowCost),
		d:      d,
	}
}

func (t target) sleep() {
	// Return immediately if duration is less than it takes to get 2 timestamps
	if t.target.UnixNano() == 0 {
		return
	}

	if t.d-t.elapsed > safeDuration {
		time.Sleep(t.d - t.elapsed - safeDuration)
	}

	for time.Now().Before(t.target) {
	}
}

func (t target) sleepTight() {
	// Return immediately if duration is less than it takes to get 2 timestamps
	if t.target.UnixNano() == 0 {
		return
	}

	for time.Now().Before(t.target) {
	}
}
