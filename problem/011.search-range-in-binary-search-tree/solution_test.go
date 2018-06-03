package lintcode

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		root *TreeNode
		k1   int
		k2   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"[Test Case 1]",
			args{
				&TreeNode{
					20,
					&TreeNode{
						8,
						&TreeNode{
							4,
							nil,
							nil,
						},
						&TreeNode{
							12,
							nil,
							nil,
						},
					},
					&TreeNode{
						22,
						nil,
						nil,
					},
				},
				10,
				22,
			},
			[]int{12, 20, 22},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.root, tt.args.k1, tt.args.k2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
