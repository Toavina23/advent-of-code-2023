package day2

import (
	utils "advent_of_code/_utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part1() {
	input := utils.ReadInputFile("./day2/input2.txt")
	var indexSum = 0
	var powerSum = 0
	for index, game := range input {
		var green = 0
		var red = 0
		var blue = 0
		gameData := (strings.Split(game, ":")[1])
		gameSets := strings.Split(gameData, ";")
		for _, gameSet := range gameSets {
			blocInfo := strings.Split(gameSet, ",")
			for _, bloc := range blocInfo {
				bloc = strings.Trim(bloc, " ")
				info := strings.Split(bloc, " ")
				secretNumber, err := strconv.Atoi(info[0])
				if err != nil {
					panic(err)
				}
				switch strings.Trim(info[1], " ") {
				case "blue":
					blue = int(math.Max(float64(blue), float64(secretNumber)))
				case "red":
					red = int(math.Max(float64(red), float64(secretNumber)))
				case "green":
					green = int(math.Max(float64(green), float64(secretNumber)))
				}
			}
		}
		powerSum += blue * red * green
		if blue <= 14 && red <= 12 && green <= 13 {
			indexSum += index + 1
		}
	}
	fmt.Println(indexSum)
	fmt.Println(powerSum)
}
