package lintcode

import "testing"

func Test_fibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				10,
			},
			34,
		},
		{
			"[Test Case 2]",
			args{
				47,
			},
			1836311903,
		},
		{
			"[Test Case 3]",
			args{
				1,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci(tt.args.n); got != tt.want {
				t.Errorf("fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
