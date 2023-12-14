package day5

import (
	utils "advent_of_code/_utils"
	"fmt"
	"sort"
	"strings"
)

type ElementData struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}
type SeedToLocation struct {
	seed     int
	location int
}

type SeedPair struct {
	seed      int
	seedRange int
}

func computeSeedPairs(seedData []int) []SeedPair {
	var seedPairs []SeedPair
	for i := 0; i < len(seedData)-1; i += 2 {
		seedPairs = append(seedPairs, SeedPair{seed: seedData[i], seedRange: seedData[i+1]})
	}
	return seedPairs
}

func organiseInput(path string) ([]int, map[int][]ElementData) {
	inputs := utils.ReadInputFile(path)
	seedInput := inputs[0]
	inputs = inputs[1:]
	seeds := utils.StrToNumberArray(strings.Split(seedInput, ":")[1], " ")

	for i := 0; i < len(seeds)-1; i++ {

	}
	parameterMap := make(map[int][]ElementData)
	parameterIndex := -1
	var elementDataTempList []ElementData
	for _, input := range inputs {
		if input == "" {
			if parameterIndex >= 0 {
				parameterMap[parameterIndex] = elementDataTempList
				elementDataTempList = make([]ElementData, 0)
			}
			continue
		}
		if strings.HasSuffix(input, ":") {
			parameterIndex++
			continue
		}
		params := utils.StrToNumberArray(input, " ")
		newElementData := ElementData{
			destinationRangeStart: params[0],
			sourceRangeStart:      params[1],
			rangeLength:           params[2],
		}
		elementDataTempList = append(elementDataTempList, newElementData)
	}
	parameterMap[parameterIndex] = elementDataTempList
	return seeds, parameterMap
}

func getMapValue(number int, elementDataList []ElementData) int {
	for _, elementData := range elementDataList {
		sourceEnd := elementData.sourceRangeStart + elementData.rangeLength
		if number <= sourceEnd && number >= elementData.sourceRangeStart {
			return elementData.destinationRangeStart + (number - elementData.sourceRangeStart)
		}
	}
	return number
}

func getMapping(elementData []ElementData, key int, mapList *[]int) {
	val := getMapValue(key, elementData)
	*mapList = append(*mapList, val)
}
func Part1() {
	seeds, parameterMap := organiseInput("./day5/input.txt")
	var seedsToLocations []SeedToLocation
	for _, seed := range seeds {
		var route []int
		route = append(route, seed)
		for i := 0; i < len(parameterMap); i++ {
			elementData := parameterMap[i]
			getMapping(elementData, route[len(route)-1], &route)
		}
		seedsToLocations = append(seedsToLocations, SeedToLocation{
			seed:     seed,
			location: route[len(route)-1],
		})
	}
	sort.Slice(seedsToLocations, func(i, j int) bool {
		return seedsToLocations[i].location < seedsToLocations[j].location
	})
	fmt.Println("Part 1")
	fmt.Printf("The nearest is %+v\n", seedsToLocations[0])
}

func Part2() {
}

func Solution() {
	Part1()
	Part2()
}
