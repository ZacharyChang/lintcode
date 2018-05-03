package lintcode

import (
	"reflect"
	"testing"
)

func Test_partitionArray(t *testing.T) {
	type args struct {
		nums []int
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
			},
			[]int{1, 3, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
