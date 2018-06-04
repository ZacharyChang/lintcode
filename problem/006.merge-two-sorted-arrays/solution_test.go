package lintcode

import (
	"reflect"
	"testing"
)

func Test_mergeSortedArray(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"[Test Case 1]",
			args{
				[]int{1, 2, 3, 4},
				[]int{2, 4, 5, 6},
			},
			[]int{1, 2, 2, 3, 4, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSortedArray_2(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
