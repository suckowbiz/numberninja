package school

import (
	"github.com/suckowbiz/numberninja/internal/calculus"
	"github.com/suckowbiz/numberninja/internal/interaction"
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

	var arithmeticOps = map[interaction.Mode][]calculus.Arithmetic{
		interaction.Multiplication:            {calculus.Multiplication},
		interaction.Division:                  {calculus.Division},
		interaction.Addition:                  {calculus.Addition},
		interaction.Subtraction:               {calculus.Subtraction},
		interaction.MultiplicationAndDivision: {calculus.Multiplication, calculus.Division},
		interaction.AdditionAndSubtraction:    {calculus.Addition, calculus.Subtraction},
		interaction.All:                       {calculus.Multiplication, calculus.Division, calculus.Addition, calculus.Subtraction},
	}

	for {
		observer.RoundTick()
		if observer.RepeatRound() {
			problem, _ = observer.PopMissed()
		} else {
			problem = calculus.GenerateProblem(arithmeticOps[mode])
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
