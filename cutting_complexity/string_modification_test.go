package cuttingcomplexity

import "testing"

func TestStringModification(t *testing.T) {
	type args struct {
		S string
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"LMMMLM", args{"LMMMLM", 5}, 1},
		{"LMMMLM*3", args{"LMMMLMLMMMLMLMMMLM", 5}, 1},
		{"MLMMLLM", args{"MLMMLLM", 3}, 1},
		{"MLMMMLMMMM", args{"MLMMMLMMMM", 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringModification(tt.args.S, tt.args.K); got != tt.want {
				t.Errorf("StringModification() = %v, want %v", got, tt.want)
			}
		})
	}
}
