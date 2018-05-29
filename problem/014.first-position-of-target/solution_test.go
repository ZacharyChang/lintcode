package lintcode

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
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
					1, 4, 4, 5, 7, 7, 8, 9, 9, 10,
				},
				1,
			},
			0,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					2, 6, 8, 13, 15, 17, 17, 18, 19, 20,
				},
				15,
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
