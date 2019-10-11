package calculus

// Problem represents a mathematical calculation problem.
type Problem struct {
	operator    rune
	left, right int
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
