package cuprum2018

import (
	"testing"
)

func TestBeautifulPassword(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		// TODO: Add test cases.
		{"Test 1", "6daa6ccaaa6d", 8},
		{"Test 2", "6daa6ccaaa6a6ccaaa6a6ccaaa6a6ccaaa6a6ccaaa6d", 40},
		{"Test 3", "6daa6ccaaa6a6ca6ccaaa6a6ccaaa6a6ccaaa6a6ccaaa6a6ccaaa6caaa6a6ccaaa6a6ccaaa6a6ccaaa6d", 80},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BeautifulPassword(tt.args); got != tt.want {
				t.Errorf("BeautifulPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isStringCharactersEven(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		// TODO: Add test cases.
		{"6daa6cca", "6daa6cca", false},
		{"6daa6ccd", "6daa6ccd", true},
		{"6daa6ccd * 5", "6daa6ccd6daa6ccd6daa6ccd6daa6ccd6daa6ccd", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStringCharactersEven(tt.args); got != tt.want {
				t.Errorf("isStringCharactersEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
