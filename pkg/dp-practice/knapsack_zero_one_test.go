package dp_practice

import (
	"testing"
)

func TestKnapsack(t *testing.T) {
	tests := []struct {
		name     string
		weights  []int
		values   []int
		capacity int
		want     int
	}{
		{
			name:     "basic example",
			weights:  []int{1, 2, 3},
			values:   []int{6, 10, 12},
			capacity: 5,
			want:     22, // items 2 and 3
		},
		{
			name:     "zero capacity",
			weights:  []int{1, 2, 3},
			values:   []int{6, 10, 12},
			capacity: 0,
			want:     0,
		},
		{
			name:     "empty weights",
			weights:  []int{},
			values:   []int{1, 2, 3},
			capacity: 5,
			want:     0,
		},
		{
			name:     "empty values",
			weights:  []int{1, 2, 3},
			values:   []int{},
			capacity: 5,
			want:     0,
		},
		{
			name:     "single item fits",
			weights:  []int{2},
			values:   []int{8},
			capacity: 2,
			want:     8,
		},
		{
			name:     "single item does not fit",
			weights:  []int{5},
			values:   []int{10},
			capacity: 3,
			want:     0,
		},
		{
			name:     "all items fit",
			weights:  []int{1, 2, 3},
			values:   []int{4, 5, 6},
			capacity: 6,
			want:     15,
		},
		{
			name:     "choose best combination",
			weights:  []int{4, 2, 3},
			values:   []int{10, 4, 7},
			capacity: 5,
			want:     11, // items 2 and 3
		},
	}

	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		// got := Knapsack(tt.weights, tt.values, tt.capacity)
	// 		got := Knapsack(tt.weights, tt.values, tt.capacity)
	// 		if got != tt.want {
	// 			t.Errorf("Knapsack() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dp := make([][]int, len(tt.weights))
			for i := range dp {
				dp[i] = make([]int, tt.capacity+1)
			}

			for i := 0; i < len(tt.weights); i++ {
				for j := 0; j < tt.capacity+1; j++ {
					dp[i][j] = -1
				}
			}

			got := KnapsackMem(tt.weights, tt.values, tt.capacity, dp)
			if got != tt.want {
				t.Errorf("Knapsack() = %v, want %v", got, tt.want)
			}
		})
	}
}

// testing tabulation impl of knapsack algo
func TestKnapsackTabulation(t *testing.T) {
	tests := []struct {
		name     string
		weights  []int
		values   []int
		capacity int
		want     int
	}{
		{
			name:     "basic example",
			weights:  []int{1, 2, 3},
			values:   []int{6, 10, 12},
			capacity: 5,
			want:     22, // items 2 and 3
		},
		{
			name:     "zero capacity",
			weights:  []int{1, 2, 3},
			values:   []int{6, 10, 12},
			capacity: 0,
			want:     0,
		},
		{
			name:     "empty weights",
			weights:  []int{},
			values:   []int{1, 2, 3},
			capacity: 5,
			want:     0,
		},
		{
			name:     "empty values",
			weights:  []int{1, 2, 3},
			values:   []int{},
			capacity: 5,
			want:     0,
		},
		{
			name:     "single item fits",
			weights:  []int{2},
			values:   []int{8},
			capacity: 2,
			want:     8,
		},
		{
			name:     "single item does not fit",
			weights:  []int{5},
			values:   []int{10},
			capacity: 3,
			want:     0,
		},
		{
			name:     "all items fit",
			weights:  []int{1, 2, 3},
			values:   []int{4, 5, 6},
			capacity: 6,
			want:     15,
		},
		{
			name:     "choose best combination",
			weights:  []int{4, 2, 3},
			values:   []int{10, 4, 7},
			capacity: 5,
			want:     11, // items 2 and 3
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// got := Knapsack(tt.weights, tt.values, tt.capacity)
			got := KnapsackTab(tt.weights, tt.values, tt.capacity)
			if got != tt.want {
				t.Errorf("Knapsack() = %v, want %v", got, tt.want)
			}
		})
	}
}

// testing tabulation impl of knapsack algo
// with slight modification in code to build the dp table.
func TestKnapsackTabulation2(t *testing.T) {
	tests := []struct {
		name     string
		weights  []int
		values   []int
		capacity int
		want     int
	}{
		{
			name:     "basic example",
			weights:  []int{1, 2, 3},
			values:   []int{6, 10, 12},
			capacity: 5,
			want:     22, // items 2 and 3
		},
		{
			name:     "zero capacity",
			weights:  []int{1, 2, 3},
			values:   []int{6, 10, 12},
			capacity: 0,
			want:     0,
		},
		{
			name:     "empty weights",
			weights:  []int{},
			values:   []int{1, 2, 3},
			capacity: 5,
			want:     0,
		},
		{
			name:     "empty values",
			weights:  []int{1, 2, 3},
			values:   []int{},
			capacity: 5,
			want:     0,
		},
		{
			name:     "single item fits",
			weights:  []int{2},
			values:   []int{8},
			capacity: 2,
			want:     8,
		},
		{
			name:     "single item does not fit",
			weights:  []int{5},
			values:   []int{10},
			capacity: 3,
			want:     0,
		},
		{
			name:     "all items fit",
			weights:  []int{1, 2, 3},
			values:   []int{4, 5, 6},
			capacity: 6,
			want:     15,
		},
		{
			name:     "choose best combination",
			weights:  []int{4, 2, 3},
			values:   []int{10, 4, 7},
			capacity: 5,
			want:     11, // items 2 and 3
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// got := Knapsack(tt.weights, tt.values, tt.capacity)
			got := KnapsackTab2(tt.weights, tt.values, tt.capacity)
			if got != tt.want {
				t.Errorf("Knapsack() = %v, want %v", got, tt.want)
			}
		})
	}
}
