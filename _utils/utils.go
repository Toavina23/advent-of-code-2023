package utils

import (
	"os"
	"strings"
)

var NUMBERS = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func ReadInputFile(fileName string) []string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	strContent := string(content)
	// because it is fucking windows
	return strings.Split(strContent, "\r\n")
}

func IsInArray(target string, array []string) bool {
	for _, element := range array {
		if element == target {
			return true
		}
	}
	return false
}

func IsNumber(target string) bool {
	return IsInArray(target, NUMBERS)
}
