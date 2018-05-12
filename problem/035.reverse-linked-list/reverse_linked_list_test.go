package lintcode

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		head *ListNode
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
						2,
						&ListNode{
							3,
							nil,
						},
					},
				},
			},
			&ListNode{
				3,
				&ListNode{
					2,
					&ListNode{
						1,
						nil,
					},
				},
			},
		},
		{
			"[Test Case 12]",
			args{
				&ListNode{
					1,
					&ListNode{
						2,
						nil,
					},
				},
			},
			&ListNode{
				2,
				&ListNode{
					1,
					nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
