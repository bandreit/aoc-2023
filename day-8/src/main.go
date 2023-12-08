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

	result1 := part1(instructions, mapOfNodes, firstValue, lastValue)
	result2 := part2(instructions, mapOfNodes)
	fmt.Println(result1, result2)
}

func part1(instructions string, mapOfNodes map[string]Node, firstValue, lastValue string) int {
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

	return s
}

func part2(instructions string, mapOfNodes map[string]Node) []int {
	nrOfInstructions := len(instructions)
	currentValues := []string{}
	results := []int{}

	for k := range mapOfNodes {
		if k[len(k)-1:] == "A" {
			currentValues = append(currentValues, k)
		}
	}

	endsInZ := false

	for _, currentValue := range currentValues {
		s := 0
		i := 0
		for !endsInZ {
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

			if currentValue[len(currentValue)-1:] == "Z" {
				endsInZ = true
				results = append(results, s)
			}
		}
		endsInZ = false
	}

	return results
}

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
