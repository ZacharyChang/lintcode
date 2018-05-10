package lintcode

import (
	"reflect"
	"testing"
)

func Test_cloneTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"[Test Case 1]",
			args{
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
		{
			"[Test Case 2]",
			args{
				&TreeNode{
					1,
					&TreeNode{
						2,
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
			&TreeNode{
				1,
				&TreeNode{
					2,
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cloneTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cloneTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
