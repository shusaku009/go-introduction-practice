package main

import "testing"

func TestIsEven(t *testing.T) {
	testcases := []struct {
		input int
		want bool
	}{
		{input: 1, want: false},
		{input: 2, want: true},
		{input: 3, want: false},
		{input: 4, want: true},
	}

	for _, tt := range testcases {
		got := isEven(tt.input)
		if got != tt.want {
			t.Errorf("isEven(%d) = %v, want %v", tt.input, got, tt.want)
		}
	}
}