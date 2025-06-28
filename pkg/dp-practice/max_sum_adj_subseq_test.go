// max sum adj sub-seq
package dp_practice

import "testing"

func Test_maxSumAdjSubSeq(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{

			name: "sample MAdjSubTest 1",
			args: args{
				input: []int{1, 2, 4},
			},
			want: 5,
		},
		{

			name: "sampe MAdjSubTest 1",
			args: args{
				input: []int{8, 8},
			},
			want: 8,
		},
		{

			name: "sampe MAdjSubTest 2",
			args: args{
				input: []int{2, 1, 4, 9},
			},
			want: 11,
		},
		{

			name: "sampe MAdjSubTest 2",
			args: args{
				input: []int{1, 2, 3, 1, 3, 5, 8, 1, 9},
			},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSumAdjSubSeq(tt.args.input); got != tt.want {
				t.Errorf("maxSumAdjSubSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}
