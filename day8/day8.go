package day8

import (
	utils "advent_of_code/_utils"
	"fmt"
	"regexp"
	"strings"
)

type Node struct {
	Label      string
	LeftLabel  string
	RightLabel string
	Left       *Node
	Right      *Node
}

func traverseTree(currentNode *Node, target string, depth int, instructions string, instructionIndex int) int {
	fmt.Println(currentNode.Label)
	fmt.Println(depth)

	if currentNode.Label == target {
		return depth
	}
	if instructionIndex > len(instructions)-1 {
		instructionIndex = 0
	}
	currentDepth := 0
	if instructions[instructionIndex] == 'L' {
		currentDepth = traverseTree(currentNode.Left, target, depth+1, instructions, instructionIndex+1)
	} else if instructions[instructionIndex] == 'R' {
		currentDepth = traverseTree(currentNode.Right, target, depth+1, instructions, instructionIndex+1)
	}
	return currentDepth
}

func Solution() {
	inputs := utils.ReadInputFile("./day8/input.txt")
	var nodes []Node
	re := regexp.MustCompile(`\(([^)]*)\s*,\s*([^)]*)\)`)
	instructions := inputs[0]
	inputs = inputs[1:]
	for _, input := range inputs {
		if input == "" {
			continue
		}
		parts := strings.Split(input, "=")
		result := re.FindStringSubmatch(parts[1])
		if len(result) >= 3 {
			nodes = append(nodes, Node{Label: strings.Split(parts[0], " ")[0], LeftLabel: result[1], RightLabel: result[2]})
		} else {
			panic("bad input")
		}
	}
	for i := 0; i < len(nodes); i++ {
		for j := 0; j < len(nodes); j++ {
			if nodes[j].Label == nodes[i].LeftLabel {
				nodes[i].Left = &nodes[j]
			}
			if nodes[j].Label == nodes[i].RightLabel {
				nodes[i].Right = &nodes[j]
			}
		}
	}
	depth := traverseTree(&nodes[0], "ZZZ", 0, instructions, 0)
	fmt.Println(depth)
}
