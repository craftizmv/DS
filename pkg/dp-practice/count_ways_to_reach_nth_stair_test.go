package dp_practice

import "testing"

func TestCountTheCostToReachNthStair(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test-CountTheCostToReachNthStair",
			args: args{
				cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountTheCostToReachNthStair(tt.args.cost); got != tt.want {
				t.Errorf("CountTheCostToReachNthStair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountTheCostToReachNthStairWithDp(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountTheCostToReachNthStairWithDp(tt.args.cost); got != tt.want {
				t.Errorf("CountTheCostToReachNthStairWithDp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountWaysToNthStair(t *testing.T) {
	type args struct {
		nStairs      int
		currentStair int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountWaysToNthStair(tt.args.nStairs, tt.args.currentStair); got != tt.want {
				t.Errorf("CountWaysToNthStair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve(t *testing.T) {
	type args struct {
		nStairs int
		cost    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.nStairs, tt.args.cost); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve2(t *testing.T) {
	type args struct {
		nStairs int
		cost    []int
		dp      []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve2(tt.args.nStairs, tt.args.cost, tt.args.dp); got != tt.want {
				t.Errorf("solve2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountTheCostToReachNthStairWithTabulation1D(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test-CountTheCostToReachNthStair-1",
			args: args{
				cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			},
			want: 6,
		},
		{
			name: "Test-CountTheCostToReachNthStair-2",
			args: args{
				cost: []int{10, 15, 20},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountTheCostToReachNthStairWithTabulation1D(tt.args.cost); got != tt.want {
				t.Errorf("CountTheCostToReachNthStairWithTabulation1D() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountTheCostToReachNthStairWithTabulationSpaceOptim(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test-CountTheCostToReachNthStair-1",
			args: args{
				cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			},
			want: 6,
		},
		{
			name: "Test-CountTheCostToReachNthStair-2",
			args: args{
				cost: []int{10, 15, 20},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountTheCostToReachNthStairWithTabulationSpaceOptim(tt.args.cost); got != tt.want {
				t.Errorf("CountTheCostToReachNthStairWithTabulationSpaceOptim() = %v, want %v", got, tt.want)
			}
		})
	}
}
