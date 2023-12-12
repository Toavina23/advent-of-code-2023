package day3

import (
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
		{"...324.-.&", []int{0, 0, 0, 324, 324, 324, 0, -1, 0, -1}},
		{"34..*.", []int{34, 34, 0, 0, -1, 0}},
		{"...3.4", []int{0, 0, 0, 3, 0, 4}},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := transformInput(tc.input)
			var resultStr = ""
			var expectedStr = ""
			for i := 0; i < len(tc.expected); i++ {
				resultStr += strconv.Itoa(result[i])
				expectedStr += strconv.Itoa(tc.expected[i])
			}
			if resultStr != expectedStr {
				t.Errorf("Expected %s got %s", expectedStr, resultStr)
			}
		})
	}
}

func TestComputePartSum(t *testing.T) {
	testCases := []struct {
		title    string
		input    []string
		expected int
	}{
		{"down check", []string{
			"...*......",
			"......633.",
			"......#...",
		}, 633},
		{"up check", []string{
			"...*..#...",
			"......633.",
			"..........",
		}, 633},
		{"left check", []string{
			"...*......",
			".....#633.",
			"..........",
		}, 633},
		{"upper left check", []string{
			"...*.#....",
			"......633.",
			"..........",
		}, 633},
		{"lower left check", []string{
			"...*......",
			"......633.",
			".....#....",
		}, 633},
		{"right check", []string{
			"...*......",
			"......633#",
			"..........",
		}, 633},
		{"upper right check", []string{
			"...*.....#",
			"......633.",
			"..........",
		}, 633},
		{"lower right check", []string{
			"...*......",
			"......633.",
			".........#",
		}, 633},
		{"Real test", []string{
			"....411...............838......721.....44..............................................607..................................................",
			"...&......519..................*..........#.97.........994..............404..............*...&43........440...882.......673.505.............",
			".....*......*...892.........971...%....131....*..........*.......515...$.......157.....412.............-.....*.............*............594.",
			"..856.495....13...-...............602..........36...$.985....341*.........88.....*.921....................122..................806..508.....",
		}, 10081},
		{"test left", []string{
			"...&......519..................*..........#.97.........994..............404..............*...&43........440...882.......673.505.............",
		}, 43},
	}
	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			result, _ := ComputePartSum(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %d got %d", tc.expected, result)
			}
		})
	}
}
