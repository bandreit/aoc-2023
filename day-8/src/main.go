package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	left  string
	right string
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	lines := strings.Split(string(dat), "\n")
	instructions := lines[0]
	rawNodes := lines[2:]

	firstValue := "AAA"
	lastValue := "ZZZ"

	mapOfNodes := buildMap(rawNodes)
	// result := traverseMap(&mapOfNodes, firstValue, lastValue, strings.Split(instructions, ""), 0)

	s := 0
	i := 0
	nrOfInstructions := len(instructions)
	currentValue := firstValue
	for lastValue != currentValue {
		if i == nrOfInstructions {
			i = 0
		}
		instruction := string(instructions[i])
		i++
		s++

		if instruction == "L" {
			currentValue = mapOfNodes[currentValue].left
		} else if instruction == "R" {
			currentValue = mapOfNodes[currentValue].right
		}
	}

	fmt.Println(s)
}

// func traverseMap(mapOfNodes *map[string]Node, prevValue string, lastValue string, instructions []string, i int) int {
// 	count := 0

// 	// if prevValue == "PXP" || prevValue == "XBK" {
// 	// 	fmt.Printf("prevValue: %s, lastValue: %s, instructions: %s, i: %d\n", prevValue, lastValue, instructions, i)
// 	// }
// 	if prevValue == lastValue || prevValue == "" {
// 		return count
// 	}

// 	if i == len(instructions) {
// 		i = 0
// 	}

// 	instruction := string(instructions[i])
// 	var value string

// 	if instruction == "L" {
// 		value = mapOfNodes[prevValue].left
// 	} else if instruction == "R" {
// 		value = mapOfNodes[prevValue].right
// 	}

// 	count++
// 	count += traverseMap(mapOfNodes, value, lastValue, instructions, i+1)

// 	return count
// }

func buildMap(nodes []string) map[string]Node {
	mapOfNodes := make(map[string]Node)

	for _, node := range nodes {
		value, left, right := extractParentAndChild(node)

		if _, ok := mapOfNodes[value]; !ok {
			mapOfNodes[value] = Node{left, right}
		}
	}

	return mapOfNodes
}

func extractParentAndChild(node string) (string, string, string) {
	data := strings.Split(node, " = ")
	parent := data[0]
	childrenData := strings.Trim(data[1], "()")
	children := strings.Split(childrenData, ", ")

	return parent, children[0], children[1]
}

// func preOrder(node *Node) {
// 	if node == nil {
// 		return
// 	}
// 	fmt.Println(node.value)
// 	preOrder(node.left)
// 	preOrder(node.right)
// }

func check(e error) {
	if e != nil {
		panic(e)
	}
}
