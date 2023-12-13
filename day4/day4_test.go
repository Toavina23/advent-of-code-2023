package day4

import (
	"strconv"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	testCases := []struct {
		input                  string
		expectedWinningNumbers string
		expectedElfNumbers     string
	}{
		{

			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			"41 48 83 86 17",
			"83 86 6 31 17 9 48 53",
		},
		{

			"Card 1: 41 48 83 86 17 41 | 83 86  6 31 17  9 48 53 9",
			"41 48 83 86 17",
			"83 86 6 31 17 9 48 53",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			winningNumbers, elfNumbers := parseInput(tc.input)
			elfNumbersStr := ""
			winningNumbersStr := ""
			for i := 0; i < len(elfNumbers); i++ {
				elfNumbersStr += " " + strconv.Itoa(elfNumbers[i])
			}
			for i := 0; i < len(winningNumbers); i++ {
				winningNumbersStr += " " + strconv.Itoa(winningNumbers[i])
			}
			elfNumbersStr = strings.Trim(elfNumbersStr, " ")
			winningNumbersStr = strings.Trim(winningNumbersStr, " ")
			if winningNumbersStr != tc.expectedWinningNumbers {
				t.Errorf("Expected winning numbers '%s', got: '%s'", tc.expectedWinningNumbers, winningNumbersStr)
			}
			if elfNumbersStr != tc.expectedElfNumbers {
				t.Errorf("Expected elf numbers '%s', got: '%s'", tc.expectedElfNumbers, elfNumbersStr)
			}
		})
	}
}
func TestCountMatch(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			4,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := countMatch(tc.input)
			if result != tc.expected {
				t.Errorf("Expected match number: %d, got: %d", tc.expected, result)
			}
		})
	}
}
