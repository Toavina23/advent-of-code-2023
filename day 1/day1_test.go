package day1

import "testing"

func TestReplaceStringNumbers(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"fivezg8jmf6hrxnhgxxttwoneg", 51},
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
		{"eighthree", 83},
		{"sevenine", 79},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := computeCalibrationValue(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %d got %d", tc.expected, result)
			}
		})
	}
}
