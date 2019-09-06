package expressionevaluator

import "testing"

func Test_eval(t *testing.T) {
	tests := []struct {
		args string
		want int32
	}{
		{"1+1", 2},
		{"1", 1},
		{"(1+5)", 6},
		{"(1+((2+3)*(4*5)))", 101},
		{"(1+((2+3)*(4*5)))-200", -99},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if got := eval(tt.args); got != tt.want {
				t.Errorf("eval() = %v, want %v", got, tt.want)
			}
		})
	}
}
