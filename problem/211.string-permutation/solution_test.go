package lintcode

import "testing"

func TestPermutation(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				"ABCAE",
				"AACBD",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Permutation(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("Permutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
