package dp_practice_test

import (
	"testing"

	dp_practice "github.com/craftizmv/DS/pkg/dp-practice"
)

func TestPerfectSquareSolve(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		target int
		want   int
	}{
		// {
		// 	"perfect square for 100",
		// 	100,
		// 	1,
		// },
		{
			"perfect square for 6",
			6,
			3, // possible combo - [ 1² + 1² + 2² ],
		},
		{
			"perfect square for 0",
			0,
			0,
		},
		{
			"perfect square for 4",
			4,
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dp_practice.PerfectSquareSolve(tt.target)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("PerfectSquareSolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
