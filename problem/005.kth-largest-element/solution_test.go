package lintcode

import "testing"

func Test_findKthMax(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				[]int{5, 4, 3, 2, 1},
				1,
			},
			5,
		},
		{
			"[Test Case 2]",
			args{
				[]int{6, 5, 2, 3, 4, 1},
				2,
			},
			5,
		},
		{
			"[Test Case 3]",
			args{
				[]int{3, 3, 3, 2, 2, 1, 6, 5, 8, 2, 2},
				1,
			},
			8,
		},
		{
			"[Test Case 4]",
			args{
				[]int{6, 5, 8, 6, 6},
				5,
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthMax(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
