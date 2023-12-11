package day3

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNumberParsing(t *testing.T) {
	testCases := []struct {
		input         string
		expected      int
		firstIndex    int
		expectedIndex int
	}{
		{"...324.-.&", 324, 3, 5},
		{"...324....", 324, 3, 5},
		{"34....", 34, 0, 1},
		{"...3.4", 3, 3, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result, nextIndex := parseNumber(tc.input, tc.firstIndex)
			if result != tc.expected {
				t.Errorf("Parsed number is not correct, expected %d got %d", tc.expected, result)
			}
			if nextIndex != tc.expectedIndex {
				t.Errorf("Next Index is not correct, expected %d got %d", tc.expected, result)
			}
		})
	}
}

func TestTransformInput(t *testing.T) {

	testCases := []struct {
		input    string
		expected []int
	}{
		{"...324.-.&", []int{0, 0, 0, 324, 0, -1, 0, -1}},
		{"34..*.", []int{34, 0, 0, -1, 0}},
		{"...3.4", []int{0, 0, 0, 3, 0, 4}},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := transformInput(tc.input)
			var resultStr = ""
			var expectedStr = ""
			for i := 0; i < len(tc.expected); i++ {
				fmt.Printf("normal: %d -- when converting: %s \n", result[i], strconv.Itoa(result[i]))
				resultStr += strconv.Itoa(result[i])
				expectedStr += strconv.Itoa(tc.expected[i])
			}
			if resultStr != expectedStr {
				t.Errorf("Expected %s got %s", expectedStr, resultStr)
			}
		})
	}
}
