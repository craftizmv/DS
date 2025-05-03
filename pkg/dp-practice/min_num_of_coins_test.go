package dp_practice

import "testing"

func Test_minNumberOfCoinsToPurchase(t *testing.T) {
	type args struct {
		coins  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic test",
			args: args{
				coins:  []int{1, 2, 5},
				target: 11,
			},
			want: 3,
		},
		{
			name: "basic test",
			args: args{
				coins:  []int{2},
				target: 3,
			},
			want: -1,
		},
		{
			name: "basic test",
			args: args{
				coins:  []int{1},
				target: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumberOfCoinsToPurchase(tt.args.coins, tt.args.target); got != tt.want {
				t.Errorf("minNumberOfCoinsToPurchase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minNumberOfCoinsToPurchaseWithDP(t *testing.T) {
	type args struct {
		coins  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic test",
			args: args{
				coins:  []int{1, 2, 5},
				target: 11,
			},
			want: 3,
		},
		{
			name: "basic test",
			args: args{
				coins:  []int{2},
				target: 3,
			},
			want: -1,
		},
		{
			name: "basic test",
			args: args{
				coins:  []int{1},
				target: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			targetDp := make(map[int]int, len(tt.args.coins)+1)
			if got := minNumberOfCoinsToPurchaseWithDp(tt.args.coins, tt.args.target, targetDp); got != tt.want {
				t.Errorf("minNumberOfCoinsToPurchase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minNumberOfCoinsToPurchaseWithTab(t *testing.T) {
	type args struct {
		coins  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic test",
			args: args{
				coins:  []int{1, 2, 5},
				target: 11,
			},
			want: 3,
		},
		{
			name: "basic test",
			args: args{
				coins:  []int{2},
				target: 3,
			},
			want: -1,
		},
		{
			name: "basic test",
			args: args{
				coins:  []int{1},
				target: 0,
			},
			want: 0,
		},
		{
			name: "leet test",
			args: args{
				coins:  []int{2, 5, 10, 1},
				target: 27,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumberOfCoinsToPurchaseWithTabulation(tt.args.coins, tt.args.target); got != tt.want {
				t.Errorf("minNumberOfCoinsToPurchase() = %v, want %v", got, tt.want)
			}
		})
	}
}
