package school

import (
	"time"

	"github.com/suckowbiz/numberninja/internal/calculus"

	"github.com/suckowbiz/numberninja/internal/mistakes"
)

// Observable is a thing that remembers statistics and facts about the current lesson.
type Observable interface {
	Over() bool
	RepeatRound() bool
	RoundTick()
	PushMissed(problem calculus.Problemer)
	PopMissed() (calculus.Problemer, error)
	Duration() time.Duration
	Rounds() int
	Miscalculations() int
}

type observer struct {
	round               int
	started             time.Time
	mistakes            mistakes.Mistakes
	duration            time.Duration
	repeatModulus       int
	miscalculationCount int
}

// NewObserver returns a new Observable
func NewObserver(duration time.Duration, repeatModulus int) Observable {
	return &observer{
		round:         0,
		started:       time.Now(),
		mistakes:      mistakes.Mistakes{},
		duration:      duration,
		repeatModulus: repeatModulus,
	}
}

func (o *observer) Miscalculations() int {
	return o.miscalculationCount
}

// PushMissed adds a mistakes problem that needs to be repeated.
func (o *observer) PushMissed(problem calculus.Problemer) {
	o.mistakes.Push(problem)
	o.miscalculationCount++
}

// PopMissed returns a mistakes problem and removes it from the remembered miscalculationCount.
func (o *observer) PopMissed() (calculus.Problemer, error) {
	return o.mistakes.Pop()
}

// Over indicates the lesson is over.
func (o *observer) Over() bool {
	elapsed := time.Since(o.started)
	return elapsed.Minutes() >= o.duration.Minutes()
}

// Duration returns the duration of the lesson.
func (o *observer) Duration() time.Duration {
	return o.duration
}

// RepeatRound indicates it is time to repeat a mistake.
func (o *observer) RepeatRound() bool {
	return o.round%o.repeatModulus == 0 && o.mistakes.Len() > 0
}

// RoundTick remembers another round is done.
func (o *observer) RoundTick() {
	o.round++
}

// Rounds returns the round count.
func (o *observer) Rounds() int {
	return o.round
}
