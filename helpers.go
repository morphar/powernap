package powernap

import (
	"time"
)

// timeNowCost is an estimate of the cost of calling time.Now.
// This is usually in ~100-200ns. But can of course vary.
var timeNowCost = 100 * time.Nanosecond

// safeDuration is used for deciding when to use native time.Sleep().
// It should be pretty high as overshooting can fluctuate wildly.
var safeDuration = 500 * time.Microsecond

// NOTE: It seems that actually getting the time from the OS should be pretty fast (ASM).
// So the cost of calling time.Now(), should be in the last part - after actually getting the time.

func init() {
	// Pre-calc timeNowCost
	runs := 10000
	start := time.Now()
	for i := 0; i < runs; i++ {
		tmp := time.Now()
		_ = tmp
	}
	timeNowCost = time.Now().Sub(start) / time.Duration(runs+1)

	// Sleep a little
	time.Sleep(100 * time.Millisecond)

	// Run pre-calc of timeNowCost again
	start = time.Now()
	for i := 0; i < runs; i++ {
		tmp := time.Now()
		_ = tmp
	}
	timeNowCost = (timeNowCost + (time.Now().Sub(start) / time.Duration(runs+1))) / 2

	// Pre-calc safeDuration
	worst := time.Nanosecond
	for i := 0; i < 1000; i++ {
		start := time.Now()
		time.Sleep(time.Nanosecond)
		duration := time.Since(start)
		if duration > worst {
			worst = duration
		}
	}

	// The iterations gives a rough idea of a worst case,
	// but we want to be far away from the danger zone.
	safeDuration = worst * 50
}
