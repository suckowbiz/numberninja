package mistakes

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/suckowbiz/numberninja/internal/calculus"
)

func TestMistakes_Push(t *testing.T) {
	mistakes := Mistakes{}
	mistakes.Push(calculus.NewProblem(calculus.Division, 1, 1))
	mistakes.Push(calculus.NewProblem(calculus.Multiplication, 1, 1))
	assert.Len(t, mistakes.stack, 2)
	assert.Equal(t, mistakes.stack[0].Operator(), calculus.Division.Rune())
	assert.Equal(t, mistakes.stack[1].Operator(), calculus.Multiplication.Rune())
}

func TestMistakes_Pop(t *testing.T) {
	mistakes := Mistakes{[]calculus.Problemer{
		calculus.NewProblem(calculus.Division, 1, 1),
		calculus.NewProblem(calculus.Multiplication, 1, 1)}}

	got, err := mistakes.Pop()
	assert.NoError(t, err)
	assert.Equal(t, got.Operator(), calculus.Division.Rune())

	got, err = mistakes.Pop()
	assert.NoError(t, err)
	assert.Equal(t, got.Operator(), calculus.Multiplication.Rune())

	got, err = mistakes.Pop()
	assert.Error(t, err)
}
