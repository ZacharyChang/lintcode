package lintcode

import "testing"

func Test_kthLargestElement(t *testing.T) {
	type args struct {
		n    int
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
				1,
				[]int{5, 4, 3, 2, 1},
			},
			5,
		},
		{
			"[Test Case 2]",
			args{
				2,
				[]int{6, 5, 2, 3, 4, 1},
			},
			5,
		},
		{
			"[Test Case 3]",
			args{
				1,
				[]int{3, 3, 3, 2, 2, 1, 6, 5, 8, 2, 2},
			},
			8,
		},
		{
			"[Test Case 4]",
			args{
				5,
				[]int{6, 5, 8, 6, 6},
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargestElement(tt.args.n, tt.args.nums); got != tt.want {
				t.Errorf("kthLargestElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
