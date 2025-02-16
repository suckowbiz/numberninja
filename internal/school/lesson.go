package school

import (
	"github.com/suckowbiz/1x1pass/internal/calculus"
	"github.com/suckowbiz/1x1pass/internal/interaction"
)

// Lesson represents a school lesson.
type Lesson struct {
}

// NewLesson constructs a new Lesson.
func NewLesson() Lesson {
	return Lesson{}
}

// Run initiates and runs a lesson until duration elapsed.
func (l Lesson) Run(io interaction.CmdLiner, observer Observable, mode interaction.Mode) {
	var problem calculus.Problemer
	io.PrintStart()
	for {
		observer.RoundTick()
		if observer.RepeatRound() {
			problem, _ = observer.PopMissed()
		} else {
			switch mode {
			case interaction.Multiplication:
				problem = calculus.GenerateProblem([]calculus.Arithmetic{calculus.Multiplication})
			case interaction.Addition:
				problem = calculus.GenerateProblem([]calculus.Arithmetic{calculus.Addition})
			case interaction.Subtraction:
				problem = calculus.GenerateProblem([]calculus.Arithmetic{calculus.Subtraction})
			case interaction.Division:
				problem = calculus.GenerateProblem([]calculus.Arithmetic{calculus.Division})
			case interaction.MultiplicationAndDivision:
				problem = calculus.GenerateProblem([]calculus.Arithmetic{calculus.Multiplication, calculus.Division})
			case interaction.AdditionAndSubtraction:
				problem = calculus.GenerateProblem([]calculus.Arithmetic{calculus.Addition, calculus.Subtraction})
			case interaction.All:
				problem = calculus.GenerateProblem([]calculus.Arithmetic{calculus.Multiplication, calculus.Division, calculus.Addition,
					calculus.Subtraction})
			}
		}

		answer := io.AskProblem(observer.Rounds(), problem.Operator(), problem.LOperand(), problem.ROperand())
		if answer == problem.Solve() {
			if observer.Over() {
				io.PrintDone(observer.Rounds(), observer.Miscalculations())
				break
			}
		} else {
			io.PrintFailure()
			observer.PushMissed(problem)
		}
	}
}
