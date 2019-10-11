package calculus

import (
	"math/rand"
	"time"
)

// Problemer groups tasks that needs to be executed on a mathematical problem.
type Problemer interface {
	Solve() int
	Operator() rune
	LOperand() int
	ROperand() int
}

// ProblemGenerator generates mathematical problems.
type ProblemGenerator struct {
}

// NewProblem generates a new problem from the given set of arithmetic calculation methods.
func NewProblem(calculationTypes []Arithmetic) Problemer {
	i := throwTheDice(0, len(calculationTypes))
	t := calculationTypes[i]
	var result Problemer
	switch t {
	case Division:
		left, right := findDivisors(2, 100, 10)
		result = Divide{Problem{
			operator: '/',
			left:     left,
			right:    right,
		}}
	case Multiplication:
		left := throwTheDice(2, 10)
		right := throwTheDice(2, 10)
		result = Multiply{Problem{
			operator: '*',
			left:     left,
			right:    right,
		}}
	default:
		panic("Not implemented, yet.")
	}
	return result
}

func findDivisors(min, lmax, rmax int) (int, int) {
	var left, right int
	if min == 0 {
		// avoid division by zero
		min = 1
	}
	for {
		left = throwTheDice(min, lmax)
		right = throwTheDice(min, rmax)
		// meet the requirements of little multiplication table
		if left > right && left%right == 0 && left/right <= 10 {
			break
		}
	}
	return left, right
}

func throwTheDice(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var n int
	for {
		n = r.Intn(max)
		if n >= min {
			break
		}
	}
	return n
}
