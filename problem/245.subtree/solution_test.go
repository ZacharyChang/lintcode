package lintcode

import "testing"

func Test_isSubtree(t *testing.T) {
	type args struct {
		T1 *TreeNode
		T2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
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
					&TreeNode{
						3,
						&TreeNode{
							4,
							nil,
							nil,
						},
						nil,
					},
				},
				&TreeNode{
					3,
					&TreeNode{
						4,
						nil,
						nil,
					},
					nil,
				},
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				&TreeNode{
					1,
					nil,
					nil,
				},
				nil,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.T1, tt.args.T2); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
