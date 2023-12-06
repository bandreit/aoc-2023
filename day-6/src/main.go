package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var numberRegex = regexp.MustCompile(`\d+`)

func getNumberOfWaysToWin(times []string, distances []string) []int {
	numberOfWaysToWin := make([]int, len(times))

	for i := 0; i < len(times); i++ {
		time, err := strconv.Atoi(times[i])
		check(err)
		distance, err := strconv.Atoi(distances[i])
		check(err)

		minSpeed := math.Ceil(float64(distance) / float64(time))

		for hold := int(minSpeed); hold < time; hold++ {
			if hold*(time-hold) > distance {
				numberOfWaysToWin[i]++
			}
		}
	}

	return numberOfWaysToWin
}

func getNumberOfWaysToWin2(times []string, distances []string) int {
	numberOfWaysToWin := 0

	time, err := strconv.Atoi(strings.Join(times, ""))
	check(err)
	distance, err := strconv.Atoi(strings.Join(distances, ""))
	check(err)

	minSpeed := math.Ceil(float64(distance) / float64(time))

	for hold := int(minSpeed); hold < time; hold++ {
		if hold*(time-hold) > distance {
			numberOfWaysToWin++
		}
	}

	return numberOfWaysToWin
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	lines := strings.Split(string(dat), "\n")
	times := numberRegex.FindAllString(lines[0], -1)
	distances := numberRegex.FindAllString(lines[1], -1)

	productOfWaysToWin := 1

	numberOfWaysToWin := getNumberOfWaysToWin(times, distances)

	for _, waysToWin := range numberOfWaysToWin {
		productOfWaysToWin *= waysToWin
	}

	fmt.Println(productOfWaysToWin)
	fmt.Println(getNumberOfWaysToWin2(times, distances))
}
