package lintcode

import (
	"reflect"
	"testing"
)

func Test_numbersByRecursion(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"[Test Case 1]",
			args{
				1,
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			"[Test Case 2]",
			args{
				0,
			},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numbersByRecursion(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numbersByRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}
