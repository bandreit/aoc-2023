package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var numberRegex = regexp.MustCompile("[0-9]")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isLinePossible(line string) int {
	gameReview := strings.Split(line, ": ")

	gameInfo := strings.Split(gameReview[0], " ")
	gameNumberString := gameInfo[1]

	cubeClusters := strings.Split(gameReview[1], "; ")

	for _, cubeCluster := range cubeClusters {
		var maxR = 12
		var maxG = 13
		var maxB = 14

		cubeInfos := strings.Split(cubeCluster, ", ")

		for _, cubeInfo := range cubeInfos {
			cubeData := strings.Split(cubeInfo, " ")
			color := cubeData[1]
			amount, err := strconv.Atoi(cubeData[0])
			check(err)

			if color == "red" {
				maxR -= amount
			} else if color == "green" {
				maxG -= amount
			} else if color == "blue" {
				maxB -= amount
			}
		}

		if maxR < 0 || maxG < 0 || maxB < 0 {
			return 0
		}
	}

	gameNumber, err := strconv.Atoi(gameNumberString)
	check(err)

	return gameNumber
}

func powerOfSet(line string) int {
	gameReview := strings.Split(line, ": ")
	cubeClusters := strings.Split(gameReview[1], "; ")

	var reds = make([]int, 0)
	var greens = make([]int, 0)
	var blues = make([]int, 0)

	for _, cubeCluster := range cubeClusters {
		cubeInfos := strings.Split(cubeCluster, ", ")

		for _, cubeInfo := range cubeInfos {
			cubeData := strings.Split(cubeInfo, " ")
			color := cubeData[1]
			amount, err := strconv.Atoi(cubeData[0])
			check(err)

			if color == "red" {
				reds = append(reds, amount)
			} else if color == "green" {
				greens = append(greens, amount)
			} else if color == "blue" {
				blues = append(blues, amount)
			}
		}
	}

	slices.Sort(reds[:])
	slices.Sort(greens[:])
	slices.Sort(blues[:])

	maxG := greens[len(greens)-1]
	maxB := blues[len(blues)-1]
	maxR := reds[len(reds)-1]

	return maxR * maxG * maxB
}

func calculatePossibleGames(gameLines []string) int {
	s := 0

	for _, line := range gameLines {
		s += isLinePossible(line)
	}

	return s
}

func calculatePowerOfGames(gameLines []string) int {
	s := 0

	for _, line := range gameLines {
		s += powerOfSet(line)
	}

	return s
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	gameLines := strings.Split(string(dat), "\n")

	sumOfPossibleGames := calculatePossibleGames(gameLines)
	sumOfPowerOfGames := calculatePowerOfGames(gameLines)

	fmt.Println(sumOfPossibleGames)
	fmt.Println(sumOfPowerOfGames)
}
