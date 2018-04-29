package lintcode

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"[Test Case 1]",
			args{
				[]int{1, 2, 3},
				[]int{1, 3, 2},
			},
			&TreeNode{
				2,
				&TreeNode{
					1,
					nil,
					nil,
				},
				&TreeNode{
					3,
					nil,
					nil,
				},
			},
		},
		{
			"[Test Case 2]",
			args{
				[]int{},
				[]int{},
			},
			nil,
		},
		{
			"[Test Case 3]",
			args{
				[]int{2, 1},
				[]int{2, 1},
			},
			&TreeNode{
				1,
				&TreeNode{
					2,
					nil,
					nil,
				},
				nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.inorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
