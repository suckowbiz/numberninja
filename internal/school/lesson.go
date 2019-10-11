package school

import (
	"github.com/suckowbiz/mentalbreak/internal/calculus"
	"github.com/suckowbiz/mentalbreak/internal/interaction"
)

// Lesson represents a school lesson.
type Lesson struct {
}

// NewLesson constructs a new Lesson.
func NewLesson() Lesson {
	return Lesson{}
}

// Run initiates and runs a lesson until duration elapsed.
func (l Lesson) Run(io interaction.Cmdliner, observer Observable) {
	var problem calculus.Problemer
	for {
		observer.RoundTick()
		if observer.RepeatRound() {
			problem, _ = observer.PopMissed()
		} else {
			problem = calculus.NewProblem([]calculus.Arithmetic{calculus.Multiplication, calculus.Division})
		}

		answer := io.AskProblem(observer.Rounds(), problem.Operator(), problem.LOperand(), problem.ROperand())
		if answer == problem.Solve() {
			if observer.Over() {
				io.PrintDone(observer.Rounds(), observer.Miscalculations())
				break
			}
		} else {
			io.PrintFailure(problem.Solve())
			observer.PushMissed(problem)
		}
	}
}
