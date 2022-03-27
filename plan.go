package powernap

import (
	"sort"
	"time"
)

type PlanResult struct {
	Start     time.Time
	End       time.Time
	ExecTimes []time.Time
}

type Plan struct {
	schedule map[time.Duration][]func()
}

func NewPlan() *Plan {
	return &Plan{
		schedule: map[time.Duration][]func(){},
	}
}

func (p *Plan) Schedule(d time.Duration, f func()) {
	if _, ok := p.schedule[d]; !ok {
		p.schedule[d] = []func(){}
	}
	p.schedule[d] = append(p.schedule[d], f)
}

func (p Plan) Start() chan PlanResult {
	ch := make(chan PlanResult)
	go func() {
		ch <- p.StartBlocking()
	}()
	return ch
}

func (p Plan) StartTight() chan PlanResult {
	ch := make(chan PlanResult)
	go func() {
		ch <- p.StartTightBlocking()
	}()
	return ch
}

func (p Plan) StartBlocking() PlanResult {
	if len(p.schedule) == 0 {
		return PlanResult{}
	}

	triggers := []time.Duration{}

	for d, _ := range p.schedule {
		triggers = append(triggers, d)
	}

	sort.Slice(triggers, func(i, j int) bool { return triggers[i] < triggers[j] })

	return p.iterateTriggers(triggers)
}

func (p Plan) StartTightBlocking() PlanResult {
	if len(p.schedule) == 0 {
		return PlanResult{}
	}

	triggers := []time.Duration{}

	for d, _ := range p.schedule {
		triggers = append(triggers, d)
	}

	sort.Slice(triggers, func(i, j int) bool { return triggers[i] < triggers[j] })

	return p.iterateTightTriggers(triggers)
}

func (p *Plan) iterateTriggers(triggers []time.Duration) PlanResult {
	res := PlanResult{
		Start:     time.Now(),
		ExecTimes: []time.Time{},
	}

	for _, d := range triggers {
		t := target{
			target:  res.Start.Add(d - timeNowCost),
			d:       d,
			elapsed: time.Now().Sub(res.Start),
		}

		t.sleep()

		for _, f := range p.schedule[d] {
			f()
			res.ExecTimes = append(res.ExecTimes, t.target)
		}
	}

	res.End = time.Now()

	return res
}

func (p *Plan) iterateTightTriggers(triggers []time.Duration) PlanResult {
	res := PlanResult{
		Start:     time.Now(),
		ExecTimes: []time.Time{},
	}

	for _, d := range triggers {
		t := target{
			target:  res.Start.Add(d - timeNowCost),
			d:       d,
			elapsed: time.Now().Sub(res.Start),
		}

		t.sleepTight()

		for _, f := range p.schedule[d] {
			f()
			res.ExecTimes = append(res.ExecTimes, t.target)
		}
	}

	res.End = time.Now()

	return res
}
