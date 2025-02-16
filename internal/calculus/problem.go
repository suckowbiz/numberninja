package calculus

// Problemer groups tasks that needs to be executed on a mathematical problem.
type Problemer interface {
	Solve() int
	Operator() rune
	LOperand() int
	ROperand() int
}

// Problem represents a mathematical calculation problem.
type Problem struct {
	operator    rune
	left, right int
}

// NewProblem creates a new problem of given input.
func NewProblem(operator Arithmetic, left, right int) Problemer {
	switch operator {
	case Multiplication:
		return Multiply{Problem{operator.Rune(), left, right}}
	case Division:
		return Divide{Problem{operator.Rune(), left, right}}
	case Addition:
		return Add{Problem{operator.Rune(), left, right}}
	case Subtraction:
		return Subtract{Problem{operator.Rune(), left, right}}
	default:
		panic("Unknown arithmetic operation")
	}
}

// LOperand returns the left side operand.
func (p Problem) LOperand() int {
	return p.left
}

// ROperand returns the right side operand.
func (p Problem) ROperand() int {
	return p.right
}

// Operator returns the calculation method operator.
func (p Problem) Operator() rune {
	return p.operator
}

// Multiply represents a multiplication problem.
type Multiply struct {
	Problem
}

// Solve solves the multiplication and returns the product.
func (m Multiply) Solve() int {
	return m.left * m.right
}

// Divide represents a division problem.
type Divide struct {
	Problem
}

// Solve solves the division and returns the quotient.
func (d Divide) Solve() int {
	return d.left / d.right
}

type Add struct {
	Problem
}

func (a Add) Solve() int {
	return a.left + a.right
}

type Subtract struct {
	Problem
}

func (s Subtract) Solve() int {
	return s.left - s.right
}
