package day3

import (
	utils "advent_of_code/_utils"
	"fmt"
	"strconv"
)

func parseNumber(source string, first int) (int, int) {
	var numStr = ""
	i := first
	for {
		if i >= len(source) || !utils.IsNumber(string(source[i])) {
			break
		}
		numStr += string(source[i])
		i++
	}
	number, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, i
	}
	return number, i - 1
}

func transformInput(input string) []int {
	var arr = make([]int, len(input))
	for i := 0; i < len(input); i++ {
		if string(input[i]) == "." {
			fmt.Printf("the character: %s\n", string(input[i]))
			arr[i] = 0
		} else if utils.IsNumber(string(input[i])) {
			number, newIndex := parseNumber(input, i)
			arr[i] = number
			fmt.Printf("the character: %d\n", number)
			for j := i + 1; j < newIndex; j++ {
				arr[j] = 0
			}
			i = newIndex
		} else {
			fmt.Printf("the character: %s\n", string(input[i]))
			arr[i] = -1
		}
	}
	for _, i := range arr {
		fmt.Printf(" %d", i)
		fmt.Printf("\n")
	}
	fmt.Printf(" ----------------\n")
	return arr
}

/*func Part1() []int {
	inputs := utils.ReadInputFile("./day3/sample.txt")
	for _, input := range inputs {
		arr := transformInput(input)
	}
}*/
