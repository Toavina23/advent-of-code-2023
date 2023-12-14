package day6

import (
	utils "advent_of_code/_utils"
	"fmt"
	"strconv"
	"strings"
)

func Part1() {
	inputs := utils.ReadInputFile("./day6/input.txt")
	times := utils.StrToNumberArray(strings.Trim(strings.Split(inputs[0], "Time:")[1], " "), " ")
	distances := utils.StrToNumberArray(strings.Trim(strings.Split(inputs[1], "Distance:")[1], " "), " ")
	result := 1
	for i := 0; i < len(times); i++ {
		solutions := 0
		for j := 0; j <= times[i]; j++ {
			d := j * (times[i] - j)
			if d > distances[i] {
				solutions++
			}
		}
		result *= solutions
	}
	fmt.Println(result)

}
func Part2() {
	inputs := utils.ReadInputFile("./day6/input.txt")
	time := strings.Join(strings.Split(strings.Trim(strings.Split(inputs[0], "Time:")[1], " "), " "), "")
	distance := strings.Join(strings.Split(strings.Trim(strings.Split(inputs[1], "Distance:")[1], " "), " "), "")
	timeNum, err := strconv.Atoi(time)
	if err != nil {
		panic(err)
	}
	distanceNum, err := strconv.Atoi(distance)
	if err != nil {
		panic(err)
	}
	solutions := 0
	for j := 0; j <= timeNum; j++ {
		d := j * (timeNum - j)
		if d > distanceNum {
			solutions++
		}
	}
	fmt.Println(solutions)
}
func Solution() {
	fmt.Println("Part 1")
	Part1()
	fmt.Println("Part 2")
	Part2()
}
