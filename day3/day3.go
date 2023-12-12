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
			arr[i] = 0
		} else if utils.IsNumber(string(input[i])) {
			number, newIndex := parseNumber(input, i)
			arr[i] = number
			for j := i + 1; j <= newIndex; j++ {
				arr[j] = number
			}
			i = newIndex
		} else {
			arr[i] = -1
		}

	}
	return arr
}
func IsAdjacent(arr [][]int, numberX int, numberY int) bool {
	isAdjacent := false
	numberStr := strconv.Itoa(arr[numberY][numberX])
	for i := numberX; i < numberX+(len(numberStr)); i++ {
		if numberY > 0 {
			if arr[numberY-1][i] == -1 {
				isAdjacent = true
			} else if i > 0 && arr[numberY-1][i-1] == -1 {
				isAdjacent = true
			} else if i < len(arr[numberY-1])-1 && arr[numberY-1][i+1] == -1 {
				isAdjacent = true
			}
		}
		if numberY < len(arr)-1 {
			if arr[numberY+1][i] == -1 {
				isAdjacent = true
			} else if i > 0 && arr[numberY+1][i-1] == -1 {
				isAdjacent = true
			} else if i < len(arr[numberY+1])-1 && arr[numberY+1][i+1] == -1 {
				isAdjacent = true
			}
		}
		if i > 0 {
			if arr[numberY][i-1] == -1 {
				isAdjacent = true
			}
		}
		if i < len(arr[numberY])-1 {
			if arr[numberY][i+1] == -1 {
				isAdjacent = true
			}
		}
	}
	return isAdjacent
}

func IsGearRatio(arr [][]int, numberX int, numberY int) int {
	gearRatio := 1
	if numberY > 0 {
		if arr[numberY-1][numberX] > 0 {
			gearRatio *= arr[numberY-1][numberX]
		} else if numberX > 0 && arr[numberY-1][numberX-1] > 0 {
			gearRatio *= arr[numberY-1][numberX-1]
		} else if numberX < len(arr[numberY-1])-1 && arr[numberY-1][numberX+1] > 0 {
			gearRatio *= arr[numberY-1][numberX+1]
		}
	}
	if numberY < len(arr)-1 {
		if arr[numberY+1][numberX] > 0 {
			gearRatio *= arr[numberY+1][numberX]
		} else if numberX > 0 && arr[numberY+1][numberX-1] > 0 {
			gearRatio *= arr[numberY+1][numberX-1]
		} else if numberX < len(arr[numberY+1])-1 && arr[numberY+1][numberX+1] > 0 {
			gearRatio *= arr[numberY+1][numberX+1]
		}
	}
	if numberX > 0 {
		if arr[numberY][numberX-1] > 0 {
			gearRatio *= arr[numberY][numberX-1]
		}
	}
	if numberX < len(arr[numberY])-1 {
		if arr[numberY][numberX+1] > 0 {
			gearRatio *= arr[numberY][numberX+1]
		}
	}
	return gearRatio
}

func ComputePartSum(inputs []string) int {
	characterMap := make([][]int, len(inputs))
	partSum := 0
	for i, input := range inputs {
		arr := transformInput(input)
		characterMap[i] = arr
	}
	for y := 0; y < len(characterMap); y++ {
		for x := 0; x < len(characterMap[y]); x++ {
			if characterMap[y][x] == 0 || characterMap[y][x] == -1 || characterMap[y][x] == -2 {
				continue
			} else if IsAdjacent(characterMap, x, y) {
				partSum += characterMap[y][x]
			}
			currentNumber := strconv.Itoa(characterMap[y][x])
			x += len(currentNumber) - 1
		}
	}
	return partSum
}

func Part1() {
	inputs := utils.ReadInputFile("./day3/input.txt")
	result := ComputePartSum(inputs)
	fmt.Println(result)
}
