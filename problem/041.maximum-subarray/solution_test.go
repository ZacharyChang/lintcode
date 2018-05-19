package lintcode

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				[]int{
					-2, 2, -3, 4, -1, 2, 1, -5, 3,
				},
			},
			6,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					5,
				},
			},
			5,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					-2,
				},
			},
			-2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
