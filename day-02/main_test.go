package main

import "testing"

func TestCheckSum(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`5 1 9 5
		  7 5 3
          2 4 6 8`, 18},
	}

	for _, test := range tests {
		if sum := checksum(test.input); sum != test.expected {
			t.Errorf("Expected checksum to be %d, but it was %d instead.", test.expected, sum)
		}
	}
}

func TestCheckSumEven(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			`5 9 2 8
			 9 4 7 3
			 3 8 6 5`, 9,
		},
	}

	for _, test := range tests {
		if sum := checksumEven(test.input); sum != test.expected {
			t.Errorf("Expected checksum to be %d, but it ws %d instead.", test.expected, sum)
		}
	}
}
