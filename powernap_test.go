package powernap

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	// Run some test sleeps at different durations
	d := time.Nanosecond
	for i := 6; i < 10; i++ {
		d = time.Duration(math.Pow10(i))

		reps := int(math.Pow10(9 - i))

		// Iterate a bunch of times, to rule out most of the time keeping
		start := time.Now()
		for n := 0; n < reps; n++ {
			Sleep(d)
		}
		total := time.Since(start)
		avg := total / time.Duration(reps)
		if math.Abs(float64(avg-d)) > float64(d)/100 {
			t.Fatalf("expected: %v, got: %v, diff: %v\n", d, avg, math.Abs(float64(avg-d)))
		}
	}
}

func TestSleepTight(t *testing.T) {
	// Run some test sleeps at different durations
	d := time.Nanosecond
	for i := 6; i < 10; i++ {
		d = time.Duration(math.Pow10(i))

		reps := int(math.Pow10(9 - i))

		// Iterate a bunch of times, to rule out most of the time keeping
		start := time.Now()
		for n := 0; n < reps; n++ {
			SleepTight(d)
		}
		total := time.Since(start)
		avg := total / time.Duration(reps)
		if math.Abs(float64(avg-d)) > float64(d)/100 {
			t.Fatalf("expected: %v, got: %v, diff: %v\n", d, avg, math.Abs(float64(avg-d)))
		}
	}
}

func BenchmarkNativeSleep(b *testing.B) {
	d := time.Nanosecond
	for i := 0; i < 10; i++ {
		d = time.Duration(math.Pow10(i))

		b.Run(fmt.Sprintf("%v", d), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				time.Sleep(d)
			}
		})
	}
}

func BenchmarkSleep(b *testing.B) {
	d := time.Nanosecond
	for i := 0; i < 10; i++ {
		d = time.Duration(math.Pow10(i))

		b.Run(fmt.Sprintf("%v", d), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Sleep(d)
			}
		})
	}
}

func BenchmarkSleepTight(b *testing.B) {
	d := time.Nanosecond
	for i := 0; i < 10; i++ {
		d = time.Duration(math.Pow10(i))

		b.Run(fmt.Sprintf("%v", d), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SleepTight(d)
			}
		})
	}
}
