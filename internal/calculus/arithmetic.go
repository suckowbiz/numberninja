package calculus

// Arithmetic represents a type of calculation.
type Arithmetic int

const (
	Multiplication Arithmetic = iota
	Division
	Addition
	Subtraction
)

// Rune returns the string representation of the given arithmetic type.
func (a Arithmetic) Rune() rune {
	return [...]rune{'*', '/', '+', '-'}[a]
}
