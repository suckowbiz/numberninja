package calculus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArithmetic_Rune(t *testing.T) {
	tests := []struct {
		name string
		a    Arithmetic
		want rune
	}{
		{
			name: "Verify multiplication has operator '*'",
			a:    Multiplication,
			want: '*',
		},
		{
			name: "Verify division has operator '/'",
			a:    Division,
			want: '/',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Rune()
			assert.Equal(t, got, tt.want)
		})
	}
}
