package calculus

// Arithmetic represents a type of calculation.
type Arithmetic int

const (
	Multiplication Arithmetic = iota
	Division
	Addition
	Subtraction
)

var arithmeticSymbols = map[Arithmetic]rune{
	Multiplication: '*',
	Division:       '/',
	Addition:       '+',
	Subtraction:    '-',
}

func (a Arithmetic) Rune() rune {
	return arithmeticSymbols[a]
}
