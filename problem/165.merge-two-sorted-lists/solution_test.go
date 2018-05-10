package lintcode

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"[Test Case 1]",
			args{
				&ListNode{
					1,
					&ListNode{
						3,
						&ListNode{
							5,
							&ListNode{
								8,
								&ListNode{
									11,
									nil,
								},
							},
						},
					},
				},
				&ListNode{
					2,
					nil,
				},
			},
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						&ListNode{
							5,
							&ListNode{
								8,
								&ListNode{
									11,
									nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
