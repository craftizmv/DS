package dp_practice

import "testing"

func TestPaintTheFence(t *testing.T) {
	tests := []struct {
		n, k     int
		expected int
	}{
		{0, 0, 0},
		{1, 0, 0},
		{0, 1, 0},
		{1, 2, 2},
		{2, 2, 4},
		{3, 2, 12},
		{4, 2, 48},
		{3, 3, 24},
		{4, 3, 120},
		{5, 2, 240},
	}

	for _, tt := range tests {
		got := paintTheFence(tt.n, tt.k)
		if got != tt.expected {
			t.Errorf("paintTheFence(%d, %d) = %d; want %d", tt.n, tt.k, got, tt.expected)
		}
	}
}
