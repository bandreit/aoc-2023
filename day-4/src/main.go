package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strings"
)

var numberRegex = regexp.MustCompile("[0-9]")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func matchesForLine(line string) int {
	line = strings.Join(strings.Fields(line), " ")
	gameInfo := strings.Split(line, " | ")
	yourCards := strings.Split(gameInfo[1], " ")
	winningCards := strings.Split(strings.Split(gameInfo[0], ": ")[1], " ")
	matches := 0

	for _, card := range yourCards {
		if slices.Contains(winningCards, card) {
			matches++
		}
	}

	return matches
}

func calculateGamePoints(gameLines []string) int {
	s := 0
	for _, line := range gameLines {
		points := 0.0
		matches := matchesForLine(line)

		if matches == 0 {
			points = 0
		}

		points = math.Pow(2, float64(matches-1))
		s += int(points)
	}

	return s
}

type lineInfo struct {
	matches int
	copies  int
}

func createDefaultStatMap(gameLines []string) map[int]lineInfo {
	m := make(map[int]lineInfo)

	for i := 0; i < len(gameLines); i++ {
		matches := matchesForLine(gameLines[i])
		m[i] = lineInfo{matches: matches, copies: 1}
	}

	return m
}

func calculateTotalScratchCards(gameLines []string) int {
	mapOfRepeatingLines := createDefaultStatMap(gameLines)

	for i := 0; i < len(gameLines); i++ {
		matches := matchesForLine(gameLines[i])

		// for each copy
		copies := mapOfRepeatingLines[i].copies
		for k := 0; k < copies; k++ {
			// for the next {matches} lines
			for j := i + 1; j <= matches+i; j++ {
				// add a copy
				if val, ok := mapOfRepeatingLines[j]; ok {
					// create copes
					newCopies := val.copies + 1
					mapOfRepeatingLines[j] = lineInfo{matches: val.matches, copies: newCopies}
				}
			}
		}
	}

	s := 0
	for _, val := range mapOfRepeatingLines {
		s += val.copies
	}

	return s
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	gameLines := strings.Split(string(dat), "\n")

	sumOfPossibleGames := calculateGamePoints(gameLines)
	totalScratchCards := calculateTotalScratchCards(gameLines)

	fmt.Println("part 1: ", sumOfPossibleGames)
	fmt.Println("part 2: ", totalScratchCards)
}
