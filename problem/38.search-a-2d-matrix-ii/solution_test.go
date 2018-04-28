package lintcode

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
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
				[][]int{
					{
						1, 3, 5, 7,
					},
					{
						10, 11, 16, 20,
					},
					{
						23, 30, 34, 50,
					},
				},
				3,
			},
			1,
		},
		{
			"[Test Case 2]",
			args{
				[][]int{},
				1,
			},
			0,
		},
		{
			"[Test Case 3]",
			args{
				[][]int{
					{1}, {3}, {5},
				},
				0,
			},
			0,
		},
		{
			"[Test Case 4]",
			args{
				[][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				5,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
