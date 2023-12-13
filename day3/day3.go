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
		} else if string(input[i]) == "*" {
			arr[i] = -2
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
			if arr[numberY-1][i] < 0 {
				isAdjacent = true
			} else if i > 0 && arr[numberY-1][i-1] < 0 {
				isAdjacent = true
			} else if i < len(arr[numberY-1])-1 && arr[numberY-1][i+1] < 0 {
				isAdjacent = true
			}
		}
		if numberY < len(arr)-1 {
			if arr[numberY+1][i] < 0 {
				isAdjacent = true
			} else if i > 0 && arr[numberY+1][i-1] < 0 {
				isAdjacent = true
			} else if i < len(arr[numberY+1])-1 && arr[numberY+1][i+1] < 0 {
				isAdjacent = true
			}
		}
		if i > 0 {
			if arr[numberY][i-1] < 0 {
				isAdjacent = true
			}
		}
		if i < len(arr[numberY])-1 {
			if arr[numberY][i+1] < 0 {
				isAdjacent = true
			}
		}
	}
	return isAdjacent
}

func IsGearRatio(arr [][]int, numberX int, numberY int) int {
	gearRatio := 1
	var gears []int
	var downDetected = false
	var upDetected = false
	if numberY > 0 {
		if arr[numberY-1][numberX] > 0 {
			gears = append(gears, arr[numberY-1][numberX])
			upDetected = true
		}
		if !upDetected {
			if numberX > 0 && arr[numberY-1][numberX-1] > 0 {
				gears = append(gears, arr[numberY-1][numberX-1])
			}
			if numberX < len(arr[numberY-1])-1 && arr[numberY-1][numberX+1] > 0 {
				gears = append(gears, arr[numberY-1][numberX+1])
			}
		}
	}
	if numberY < len(arr)-1 {
		if arr[numberY+1][numberX] > 0 {
			gears = append(gears, arr[numberY+1][numberX])
			downDetected = true
		}
		if !downDetected {
			if numberX > 0 && arr[numberY+1][numberX-1] > 0 {
				gears = append(gears, arr[numberY+1][numberX-1])
			}
			if numberX < len(arr[numberY+1])-1 && arr[numberY+1][numberX+1] > 0 {
				gears = append(gears, arr[numberY+1][numberX+1])
			}
		}
	}
	if numberX > 0 {
		if arr[numberY][numberX-1] > 0 {
			gears = append(gears, arr[numberY][numberX-1])
		}
	}
	if numberX < len(arr[numberY])-1 {
		if arr[numberY][numberX+1] > 0 {
			gears = append(gears, arr[numberY][numberX+1])
		}
	}
	if len(gears) > 1 {
		for _, v := range gears {
			gearRatio *= v
		}
		return gearRatio
	}
	return 0
}

func ComputePartSum(inputs []string) (int, int) {
	characterMap := make([][]int, len(inputs))
	partSum := 0
	gearRatioSum := 0
	for i, input := range inputs {
		arr := transformInput(input)
		characterMap[i] = arr
	}
	for y := 0; y < len(characterMap); y++ {
		for x := 0; x < len(characterMap[y]); x++ {
			if characterMap[y][x] <= 0 {
				if characterMap[y][x] == -2 {
					ratio := IsGearRatio(characterMap, x, y)
					gearRatioSum += ratio
				}
				continue
			} else if IsAdjacent(characterMap, x, y) {
				partSum += characterMap[y][x]
			}
			currentNumber := strconv.Itoa(characterMap[y][x])
			x += len(currentNumber) - 1
		}
	}
	return partSum, gearRatioSum
}

func Solution() {
	inputs := utils.ReadInputFile("./day3/input.txt")
	partSum, gearRatioSum := ComputePartSum(inputs)
	fmt.Println(partSum)
	fmt.Println(gearRatioSum)
}
