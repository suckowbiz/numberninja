package mistakes

import (
	"errors"

	"github.com/suckowbiz/mentalbreak/internal/calculus"
)

// Mistakes implements a LIFO stack.
type Mistakes struct {
	stack []calculus.Problemer
}

func (m *Mistakes) Len() int {
	return len(m.stack)
}

func (m *Mistakes) Push(problem calculus.Problemer) {
	m.stack = append(m.stack, problem)
}

func (m *Mistakes) Pop() (calculus.Problemer, error) {
	if m.Len() == 0 {
		return nil, errors.New("empty stack")
	}
	var res calculus.Problemer
	res, m.stack = m.stack[0], m.stack[1:]
	return res, nil
}
