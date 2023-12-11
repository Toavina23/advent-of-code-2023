package day1

import (
	utils "advent_of_code/_utils"
	"fmt"
	"strconv"
	"strings"
)

var NUMBERS = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var NUMBERS_TOSTR = map[string]string{
	"9": "nine",
	"8": "eight",
	"7": "seven",
	"6": "six",
	"5": "five",
	"4": "four",
	"3": "three",
	"2": "two",
	"1": "one",
}

func isInArray(target string, array []string) bool {
	for _, element := range array {
		if element == target {
			return true
		}
	}
	return false
}

func replaceNumberStrings(input string) []string {
	digits := []string{}
	for i := 0; i < len(input); i++ {
		if isInArray(string(input[i]), NUMBERS) {
			digits = append(digits, string(input[i]))
			continue
		}
		for key, value := range NUMBERS_TOSTR {
			if strings.HasPrefix(input[i:], value) {
				digits = append(digits, key)
				break
			}
		}
	}
	return digits
}
func computeCalibrationValue(input string) int {
	numbers := replaceNumberStrings(input)
	var calibrationValue int
	var err error
	if len(numbers) == 1 {
		calibrationValue, err = strconv.Atoi(numbers[0])
		calibrationValue *= 10
	} else {
		calibrationValue, err = strconv.Atoi(fmt.Sprintf("%s%s", numbers[0], numbers[len(numbers)-1]))
	}
	if err != nil {
		fmt.Println(err)
	}
	return calibrationValue
}

func Part1() {
	inputs := utils.ReadInputFile("./sample.txt")
	var sum int = 0
	for i := 0; i < len(inputs); i++ {
		var firstDigit = ""
		var lastDigit = ""
		for j := 0; j < len(inputs[i]); j++ {
			strlen := len(inputs[i]) - 1
			if firstDigit == "" && isInArray(string(inputs[i][j]), NUMBERS) {
				firstDigit = string(inputs[i][j])
			}
			if lastDigit == "" && isInArray(string(inputs[i][strlen-j]), NUMBERS) {
				lastDigit = string(inputs[i][strlen-j])
			}
		}
		var calibrationValue int
		var err error
		if firstDigit == "" {
			calibrationValue, err = strconv.Atoi(firstDigit)
		} else {
			calibrationValue, err = strconv.Atoi(fmt.Sprintf("%s%s", firstDigit, lastDigit))
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		sum += calibrationValue
	}
	fmt.Println(sum)
}

func Part2() {
	inputs := utils.ReadInputFile("./data")
	var sum int = 0
	for i := 0; i < len(inputs); i++ {
		sum += computeCalibrationValue(inputs[i])
	}
	fmt.Println(sum)
}
