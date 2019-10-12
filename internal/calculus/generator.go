package calculus

import (
	"math/rand"
	"time"
)

const (
	minFactor             = 2
	maxFactor             = 10
	maxDivisor            = 100
	maxDividend           = 10
	maxQuotient           = 10
	minDivisorAndDividend = 2
)

// GenerateProblem generates a new problem from the given set of arithmetic calculation methods.
func GenerateProblem(arithmetic []Arithmetic) Problemer {
	i := throwTheDice(0, len(arithmetic))
	a := arithmetic[i]

	var l, r int
	switch arithmetic[i] {
	case Multiplication:
		l, r = randFactors(minFactor, maxFactor)
	case Division:
		l, r = randDivisorAndDividend(minDivisorAndDividend, maxDivisor, maxDividend)
	default:
		panic("Not implemented, yet.")
	}

	result := NewProblem(a, l, r)
	return result
}

func randFactors(min, max int) (int, int) {
	return throwTheDice(min, max), throwTheDice(min, max)
}

func randDivisorAndDividend(min, maxDivisor, maxDividend int) (int, int) {
	var divisor, dividend int
	if min == 0 {
		panic("division by zero is illegal")
	}
	// meet the requirements of little multiplication table
	for divisor <= dividend || divisor%dividend != 0 || divisor/dividend > maxQuotient {
		divisor = throwTheDice(min, maxDivisor)
		dividend = throwTheDice(min, maxDividend)
	}
	return divisor, dividend
}

func throwTheDice(min, max int) int {
	if min == max {
		return min
	}

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
