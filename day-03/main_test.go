package main

import "testing"

func TestDistAccessPort(t *testing.T) {
	tests := []struct {
		input    float64
		expected int
	}{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31},
	}

	for _, test := range tests {
		if expected := distAccessPort(test.input); expected != test.expected {
			t.Errorf("Expected distance to be %d, but it was %d instead. \n", test.expected, expected)
		}
	}
}
