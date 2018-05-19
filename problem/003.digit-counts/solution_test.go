package lintcode

import "testing"

func Test_digitCounts(t *testing.T) {
	type args struct {
		k int
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
				1,
				12,
			},
			5,
		},
		{
			"[Test Case 2]",
			args{
				0,
				19,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digitCounts(tt.args.k, tt.args.n); got != tt.want {
				t.Errorf("digitCounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
