package lintcode

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"[Test Case 1]",
			args{
				[][]int{
					{
						1, 2, 3,
					},
					{
						4, 5, 6,
					},
					{
						7, 8, 9,
					},
				},
			},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			"[Test Case 2]",
			args{
				[][]int{},
			},
			[]int{},
		},
		{
			"[Test Case 3]",
			args{
				[][]int{
					{1, 11},
					{2, 12},
					{3, 13},
					{4, 14},
					{5, 15},
					{6, 16},
					{7, 17},
					{8, 18},
					{9, 19},
					{10, 20},
				},
			},
			[]int{
				1, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 10, 9, 8, 7, 6, 5, 4, 3, 2,
			},
		},
		{
			"[Test Case 4]",
			args{
				[][]int{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
				},
			},
			[]int{
				1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
