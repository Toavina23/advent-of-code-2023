package day4

import (
	utils "advent_of_code/_utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func parseInput(input string) ([]int, []int) {
	lists := strings.Split(input, "|")
	winningNumbers := strings.Split(strings.Trim(strings.Split(lists[0], ":")[1], " "), " ")
	elfNumbers := strings.Split(strings.Trim(lists[1], " "), " ")
	var winningNumbersSet []int
	var elfNumbersSet []int
	for _, v := range winningNumbers {
		if v != "" {
			val, err := strconv.Atoi(strings.Trim(v, " "))
			if err != nil {
				panic(err)
			}
			if !utils.IsInArray(val, winningNumbersSet) {
				winningNumbersSet = append(winningNumbersSet, val)
			}
		}
	}
	for _, v := range elfNumbers {
		if v != "" {
			val, err := strconv.Atoi(strings.Trim(v, " "))
			if err != nil {
				panic(err)
			}
			if !utils.IsInArray(val, elfNumbersSet) {
				elfNumbersSet = append(elfNumbersSet, val)
			}
		}
	}
	return winningNumbersSet, elfNumbersSet
}

func countMatch(input string) int {
	match := 0
	winningNumbers, elfNumber := parseInput(input)
	for i := 0; i < len(winningNumbers); i++ {
		for j := 0; j < len(elfNumber); j++ {
			if winningNumbers[i] == elfNumber[j] {
				match += 1
			}
		}
	}
	return match
}
func part2() {
	inputs := utils.ReadInputFile("./day4/input.txt")
	var queue []string
	var cardMappings = make(map[string]int)
	queue = append(queue, inputs...)
	for {
		if len(queue) == 0 {
			break
		}
		currentInput := queue[0]
		queue = queue[1:]
		mapKey := strings.Trim(strings.Split(strings.Split(currentInput, ":")[0], "Card")[1], " ")
		value, ok := cardMappings[mapKey]
		if ok {
			cardMappings[mapKey] = value + 1
		} else {
			cardMappings[mapKey] = 1
		}
		matches := countMatch(currentInput)
		gameNumber, err := strconv.Atoi(mapKey)
		if err != nil {
			panic(err)
		}
		for i := gameNumber; i < gameNumber+matches; i++ {
			queue = append(queue, inputs[i])
		}
	}
	cardNumbers := 0
	for _, value := range cardMappings {
		cardNumbers += value
	}
	fmt.Println(cardNumbers)
}

func part1() {
	inputs := utils.ReadInputFile("./day4/input.txt")
	pointsTotal := 0
	for _, input := range inputs {
		matches := countMatch(input)
		pointsTotal += int(math.Pow(2, float64(matches-1)))
	}
	fmt.Println(pointsTotal)
}
func Solution() {
	fmt.Println("Part 1")
	part1()
	fmt.Println("Part 2")
	part2()
}
