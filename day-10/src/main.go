package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Point struct {
	i int
	j int
}

var possiblePipeConnections = map[rune][]Point{
	'|': {{i: -1, j: 0}, {i: 1, j: 0}},  // north, south
	'-': {{i: 0, j: -1}, {i: 0, j: 1}},  // west, east
	'L': {{i: -1, j: 0}, {i: 0, j: 1}},  // north, east
	'J': {{i: -1, j: 0}, {i: 0, j: -1}}, // north, west
	'7': {{i: 1, j: 0}, {i: 0, j: -1}},  // south, west
	'F': {{i: 1, j: 0}, {i: 0, j: 1}},   // south, east
}

var directionsToCheck = []Point{
	{i: 1, j: 0},
	{i: 0, j: 1},
	{i: -1, j: 0},
	{i: 0, j: -1},
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	twoDArr, sI, sJ := build2DArrayFromLines(string(dat))
	twoDArr = addPaddingTo2DArray(twoDArr, '.')

	i := sI + 1
	j := sJ + 1

	initial2Connections := make([]Point, 0)

	for _, directionToCheck := range directionsToCheck {
		checkingAt := Point{i: i + directionToCheck.i, j: j + directionToCheck.j}
		pipeToCheck := twoDArr[checkingAt.i][checkingAt.j]

		acceptsConnectionsFrom := possiblePipeConnections[pipeToCheck]
		if slices.ContainsFunc(acceptsConnectionsFrom, func(p Point) bool {
			return p.i == i-checkingAt.i && p.j == j-checkingAt.j
		}) {
			initial2Connections = append(initial2Connections, Point{i: checkingAt.i, j: checkingAt.j})
		}
	}

	for _, initialConnection := range initial2Connections {
		fmt.Println(string(twoDArr[initialConnection.i][initialConnection.j]))
	}

	steps := 0

	reachedEnd := false
	comingFrom := Point{i, j}
	goingTo := initial2Connections[0]

	for !reachedEnd {
		oldComingFrom := comingFrom
		comingFrom = goingTo

		possibleDirectionsToGo := possiblePipeConnections[twoDArr[comingFrom.i][comingFrom.j]]
		for _, directionToCheck := range possibleDirectionsToGo {
			checkingAt := Point{comingFrom.i + directionToCheck.i, comingFrom.j + directionToCheck.j}
			if (checkingAt.i == oldComingFrom.i) && (checkingAt.j == oldComingFrom.j) {
				continue
			}

			goingTo = checkingAt
			steps++
			fmt.Println("going to ", goingTo, " which is ", string(twoDArr[goingTo.i][goingTo.j]))
			if string(twoDArr[goingTo.i][goingTo.j]) == "S" {
				reachedEnd = true
			}
		}
	}

	fmt.Println(steps/2 + 1)
}

func build2DArrayFromLines(s string) ([][]rune, int, int) {
	sI := 0
	sJ := 0

	lines := strings.Split(s, "\n")
	m := len(lines)

	twoD := make([][]rune, m)
	for i, line := range lines {
		twoD[i] = []rune(line)

		for j, r := range twoD[i] {
			if r == 'S' {
				sI = i
				sJ = j
			}
		}
	}

	return twoD, sI, sJ
}

func addPaddingTo2DArray(twoDArr [][]rune, padding rune) [][]rune {
	for i := range twoDArr {
		twoDArr[i] = append(twoDArr[i], padding)
		twoDArr[i] = append([]rune{padding}, twoDArr[i]...)
	}

	paddingSlice := prefillSliceWithRunes([]rune{}, padding, len(twoDArr[0]))
	twoDArr = append([][]rune{paddingSlice}, twoDArr...)
	twoDArr = append(twoDArr, paddingSlice)

	return twoDArr
}

func prefillSliceWithRunes(s []rune, r rune, n int) []rune {
	for i := 0; i < n; i++ {
		s = append(s, r)
	}

	return s
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
