package calculus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProblem(t *testing.T) {
	type args struct {
		operator Arithmetic
		left     int
		right    int
	}
	tests := []struct {
		name string
		args args
		want Problem
	}{
		{
			name: "Verify left and right operands are not mixed",
			args: args{
				left:  1,
				right: 2,
			},
			want: Problem{
				left:  1,
				right: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewProblem(tt.args.operator, tt.args.left, tt.args.right)
			assert.NotNil(t, got)
			assert.Equal(t, tt.args.left, got.LOperand())
			assert.Equal(t, tt.args.right, got.ROperand())
		})
	}
}

func TestMultiply_Solve(t *testing.T) {
	type fields struct {
		Problem Multiply
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Verify multiplication is calculated correctly 1",
			fields: fields{Multiply{Problem{
				operator: '*',
				left:     1,
				right:    2,
			}}},
			want: 2,
		},
		{
			name: "Verify multiplication is calculated correctly 2",
			fields: fields{Multiply{Problem{
				operator: '*',
				left:     5,
				right:    5,
			}}},
			want: 25,
		},
		{
			name: "Verify multiplication is calculated correctly 2",
			fields: fields{Multiply{Problem{
				operator: '*',
				left:     10,
				right:    10,
			}}},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.Problem.Solve()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDivide_Solve(t *testing.T) {
	type fields struct {
		Problem Divide
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Verify division works as expected 1",
			fields: fields{Divide{Problem{
				left:  4,
				right: 2,
			}}},
			want: 2,
		},
		{
			name: "Verify division works as expected 2",
			fields: fields{Divide{Problem{
				left:  100,
				right: 10,
			}}},
			want: 10,
		},
		{
			name: "Verify division works as expected 3",
			fields: fields{Divide{Problem{
				left:  25,
				right: 5,
			}}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.Problem.Solve()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAdd_Solve(t *testing.T) {
	type fields struct {
		Problem Add
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Verify addition works as expected",
			fields: fields{Add{Problem{
				left:  4,
				right: 2,
			}}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.Problem.Solve()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSubtract_Solve(t *testing.T) {
	type fields struct {
		Problem Subtract
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Verify subtraction works as expected 1",
			fields: fields{Subtract{Problem{
				left:  4,
				right: 2,
			}}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.Problem.Solve()
			assert.Equal(t, tt.want, got)
		})
	}
}
