package calculus

// Arithmetic represents a type of calculation.
type Arithmetic int

const (
	// Multiplication represents the arithmetic type of that name.
	Multiplication Arithmetic = iota
	// Division represents the arithmetic type of that name.
	Division
)

// Rune returns the string representation of the given arithmetic type.
func (a Arithmetic) Rune() rune {
	return [...]rune{'*', '/'}[a]
}
