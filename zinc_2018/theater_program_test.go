package zinc2018

import (
	"testing"
)

func TestTheaterTickets(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		// TODO: Add test cases.
		{"Test 1", []int{1, 2, 1, 1}, 3},
		{"Test 2", []int{1, 2, 3, 4}, 4},
		{"Test 3", []int{2, 2, 1, 2, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TheaterTickets(tt.args); got != tt.want {
				t.Errorf("TheaterTickets() = %v, want %v", got, tt.want)
			}
		})
	}
}
