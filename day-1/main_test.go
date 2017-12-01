package main

import "testing"

func TestSumSequential(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"1111", 4},
		{"1122", 3},
		{"1234", 0},
		{"91212129", 9},
	}
	for _, test := range tests {
		if sum := sumSequential(test.input); sum != test.expected {
			t.Errorf("Expected sum to be %d, but it was %d instead.", test.expected, sum)
		}
	}
}

func TestSumMid(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	}
	for _, test := range tests {
		if sum := sumMid(test.input); sum != test.expected {
			t.Errorf("Expected sum to be %d, but it was %d instead.", test.expected, sum)
		}
	}
}
