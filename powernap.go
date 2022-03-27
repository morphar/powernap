package powernap

import (
	"time"
)

// Sleep is a drop-in replacement for time.Sleep.
// It tries to be more precise, but may fall short if safeDuration
// is set too low.
// For short sleep durations, where you want highest possible precision
// and are willing to spend CPU cycles, use SleepTight instead.
func Sleep(d time.Duration) {
	t := newTarget(d)
	t.sleep()
}

// SleepTight aims for the highest possible precision in
// exchange for CPU cycles.
func SleepTight(d time.Duration) {
	t := newTarget(d)
	t.sleepTight()
}
