package powernap

import (
	"math"
	"sort"
	"testing"
	"time"
)

var planTests = []time.Duration{
	100 * time.Millisecond,
	100 * time.Millisecond,
	200 * time.Millisecond,
	300 * time.Millisecond,
	400 * time.Millisecond,
	500 * time.Millisecond,
	314 * time.Millisecond,
	201 * time.Millisecond,
}

func checkPlanResult(t *testing.T, plan *Plan, res PlanResult) {
	t.Helper()

	scheduleTotal := 0
	for _, s := range plan.schedule {
		scheduleTotal += len(s)
	}

	if scheduleTotal != len(planTests) {
		t.Fatalf("number of scheduled items: %d, doesn't match number of tests: %d", scheduleTotal, len(planTests))
	}

	if len(res.ExecTimes) != len(planTests) {
		t.Fatalf("number of recorded executions: %d, doesn't match number of tests: %d", len(res.ExecTimes), len(planTests))
	}

	sort.Slice(planTests, func(i, j int) bool { return planTests[i] < planTests[j] })

	for i, et := range res.ExecTimes {
		expected := res.Start.Add(planTests[i])
		diff := math.Abs(float64(et.Sub(expected)))
		diffRatio := diff / float64(planTests[i])

		if diffRatio > 0.01 {
			t.Fatalf("expected an error rate of less than 1%%, got: %0.2f", diffRatio*100)
		}
	}
}

func TestPlan(t *testing.T) {
	plan := NewPlan()
	for _, d := range planTests {
		plan.Schedule(d, func() { /* do nothing */ })
	}
	res := <-plan.Start()

	checkPlanResult(t, plan, res)
}

func TestPlanTight(t *testing.T) {
	plan := NewPlan()
	for _, d := range planTests {
		plan.Schedule(d, func() { /* do nothing */ })
	}
	res := <-plan.StartTight()

	checkPlanResult(t, plan, res)
}

func TestPlanBlocking(t *testing.T) {
	plan := NewPlan()
	for _, d := range planTests {
		plan.Schedule(d, func() { /* do nothing */ })
	}
	res := plan.StartBlocking()

	checkPlanResult(t, plan, res)
}

func TestPlanTightBlocking(t *testing.T) {
	plan := NewPlan()
	for _, d := range planTests {
		plan.Schedule(d, func() { /* do nothing */ })
	}
	res := plan.StartTightBlocking()

	checkPlanResult(t, plan, res)
}
