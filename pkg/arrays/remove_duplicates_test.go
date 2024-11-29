package arrays

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "When contain duplicate, then should return just 1 result",
			args: args{nums: []int{1, 1, 2}},
			want: 2,
		},
		{
			name: "When contain duplicate, then should return just 1 result",
			args: args{nums: []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
