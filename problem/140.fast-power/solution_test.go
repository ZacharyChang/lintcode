package lintcode

import "testing"

func Test_fastPower(t *testing.T) {
	type args struct {
		a int
		b int
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
				2,
				3,
				31,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fastPower(tt.args.a, tt.args.b, tt.args.n); got != tt.want {
				t.Errorf("fastPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
