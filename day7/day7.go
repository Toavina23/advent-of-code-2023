package day7

import (
	utils "advent_of_code/_utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
	Label string
	Value int
}

type Hand struct {
	cards string
	bid   int
}

var cardsMap = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 1,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

func compareHand(hand1 string, hand2 string) int {
	hand1Occurences := evaluateHand(hand1)
	hand2Occurences := evaluateHand(hand2)
	for i, j := 0, 0; i < len(hand1Occurences) && j < len(hand2Occurences); i, j = i+1, j+1 {
		if hand1Occurences[i] > hand2Occurences[j] {
			return 0
		} else if hand1Occurences[i] < hand2Occurences[j] {
			return 1
		}
	}
	for i := 0; i < 5; i++ {
		v1 := cardsMap[string(hand1[i])]
		v2 := cardsMap[string(hand2[i])]
		if v1 > v2 {
			return 0
		} else if v1 < v2 {
			return 1
		}
	}
	return -1
}
func evaluateHand(hand string) []int {
	cardMap := make(map[string]int)
	var occurences []int
	jokerCount := 0
	for _, e := range hand {
		if string(e) == "J" {
			jokerCount++
			continue
		}
		val, ok := cardMap[string(e)]
		if ok {
			cardMap[string(e)] = val + 1
		} else {
			cardMap[string(e)] = 1
		}
	}
	for _, v := range cardMap {
		occurences = append(occurences, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(occurences)))
	if jokerCount > 0 {
		if len(occurences) != 0 {
			occurences[0] = occurences[0] + jokerCount
		} else {
			occurences = append(occurences, jokerCount)
		}
	}
	return occurences
}

func Solution() {
	inputs := utils.ReadInputFile("./day7/input.txt")
	var hands []Hand
	for _, input := range inputs {
		data := strings.Split(input, " ")
		hand := data[0]
		bid, err := strconv.Atoi(data[1])
		if err != nil {
			panic(err)
		}
		hands = append(hands, Hand{hand, bid})
	}
	sort.Slice(hands, func(i, j int) bool {
		return compareHand(hands[i].cards, hands[j].cards) == 1
	})
	sum := 0
	for i, hand := range hands {
		sum += hand.bid * (i + 1)
	}
	fmt.Println(sum)
}
