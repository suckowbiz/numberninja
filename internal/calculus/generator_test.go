package calculus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateProblem(t *testing.T) {
	type args struct {
		arithmetic []Arithmetic
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			name: "Verify returns only possible arithmetic",
			args: args{[]Arithmetic{Multiplication}},
			want: '*',
		},
		{
			name: "Verify returns only possible arithmetic",
			args: args{[]Arithmetic{Division}},
			want: '/',
		},
		{
			name: "Verify returns only possible arithmetic",
			args: args{[]Arithmetic{Addition}},
			want: '+',
		},
		{
			name: "Verify returns only possible arithmetic",
			args: args{[]Arithmetic{Subtraction}},
			want: '-',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateProblem(tt.args.arithmetic)
			assert.Equal(t, tt.want, got.Operator())
		})
	}
}

func Test_findFactors(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Verify min and max values are in range.",
			args: args{
				min: 2,
				max: 2,
			},
			want: []int{2},
		},
		{
			name: "Verify min and max values are in range.",
			args: args{
				min: 10,
				max: 10,
			},
			want: []int{10},
		},
		{
			name: "Verify min and max values are in range.",
			args: args{
				min: 1,
				max: 10,
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := randFactors(tt.args.min, tt.args.max)
			assert.Contains(t, tt.want, got)
			assert.Contains(t, tt.want, got1)
		})
	}
}

func Test_findDivisorAndDividend(t *testing.T) {
	type args struct {
		min         int
		maxDivisor  int
		maxDividend int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		{
			name: "Verify dividend and divisor are in expected range",
			args: args{
				min:         2,
				maxDivisor:  10,
				maxDividend: 10,
			},
			want:  []int{2, 3, 4, 5, 6, 7, 8, 9, 10},
			want1: []int{2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := randDivisorAndDividend(tt.args.min, tt.args.maxDivisor, tt.args.maxDividend)
			assert.GreaterOrEqual(t, got, got1)
			assert.Contains(t, tt.want, got)
			assert.Contains(t, tt.want1, got1)
		})
	}
}

func Test_findAddends(t *testing.T) {
	type args struct {
		maxSum    int
		minAddend int
		maxAddend int
	}
	tests := []struct {
		name            string
		args            args
		acceptedAddends []int
	}{
		{
			name: "Verify addends are in expected range",
			args: args{
				maxSum:    10,
				minAddend: 1,
				maxAddend: 9,
			},
			acceptedAddends: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			laddend, raddend := randAddends(tt.args.minAddend, tt.args.maxAddend, tt.args.maxAddend)
			assert.GreaterOrEqual(t, laddend, tt.args.minAddend)
			assert.LessOrEqual(t, laddend, tt.args.maxAddend)

			assert.GreaterOrEqual(t, raddend, tt.args.minAddend)
			assert.LessOrEqual(t, raddend, tt.args.maxAddend)

			assert.NotZero(t, laddend)
			assert.NotZero(t, raddend)

			assert.Contains(t, tt.acceptedAddends, laddend)
			assert.Contains(t, tt.acceptedAddends, raddend)
		})
	}
}

func Test_findMinuendAndSubtrahend(t *testing.T) {
	type args struct {
		minMinuend int
		maxMinuend int
	}
	tests := []struct {
		name                string
		args                args
		acceptedMinuends    []int
		acceptedSubtrahends []int
	}{
		{
			name: "Verify minuend and subtrahend are in expected range",
			args: args{
				minMinuend: 2,
				maxMinuend: 9,
			},
			acceptedMinuends:    []int{2, 3, 4, 5, 6, 7, 8, 9, 10},
			acceptedSubtrahends: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minuend, subtrahend := randMinuendAndSubtrahend(tt.args.minMinuend, tt.args.maxMinuend)
			assert.GreaterOrEqual(t, minuend, tt.args.minMinuend)
			assert.LessOrEqual(t, minuend, tt.args.maxMinuend)

			assert.Greater(t, minuend, subtrahend)

			assert.NotZero(t, minuend)
			assert.NotZero(t, subtrahend)
		})
	}
}
