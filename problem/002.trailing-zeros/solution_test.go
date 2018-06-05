package lintcode

import "testing"

func Test_trailingZeros(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"[Test Case 1]",
			args{
				11,
			},
			2,
		},
		{
			"[Test Case 2]",
			args{
				1001171717,
			},
			250292920,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeros(tt.args.n); got != tt.want {
				t.Errorf("trailingZeros() = %v, want %v", got, tt.want)
			}
		})
	}
}
