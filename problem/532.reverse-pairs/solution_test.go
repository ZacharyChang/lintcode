package lintcode

import "testing"

func Test_reversePairs(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"[Test Case 1]",
			args{
				[]int{2, 3, 1, 55, 6, 4, 7, 3, 0},
			},
			18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePairs(tt.args.A); got != tt.want {
				t.Errorf("reversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
