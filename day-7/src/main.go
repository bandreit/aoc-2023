package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1(dat []byte) int {
	lines := strings.Split(string(dat), "\n")
	mapOfCardsByPower := make(map[int][]string)
	mapOfBidsByHand := make(map[string]int)

	for i := 0; i < len(lines); i++ {
		data := strings.Split(lines[i], " ")
		hand := data[0]
		bid, err := strconv.Atoi(data[1])
		check(err)

		cards := strings.Split(hand, "")

		mapOfOccurrences := make(map[string]int)
		occurrencesArray := make([]int, 0)

		for _, card := range cards {
			occurrences := strings.Count(hand, string(card))
			mapOfOccurrences[string(card)] = occurrences
		}

		for _, value := range mapOfOccurrences {
			occurrencesArray = append(occurrencesArray, value)
		}

		sort.Ints(occurrencesArray)
		combo := make([]string, 0)

		for _, value := range occurrencesArray {
			str := strconv.Itoa(value)
			combo = append(combo, str)
		}

		pokerHand := strings.Join(combo, "")
		mapOfBidsByHand[hand] = bid

		power := mapOfPowers[pokerHand]
		if mapOfCardsByPower[power] == nil {
			mapOfCardsByPower[power] = make([]string, 0)
		}

		mapOfCardsByPower[power] = append(mapOfCardsByPower[power], hand)
	}

	keys := make([]int, 0, len(mapOfCardsByPower))
	for k := range mapOfCardsByPower {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	totalWinningHands := 0
	rank := 0

	for _, v := range mapOfCardsByPower {
		rank += len(v)
	}

	fmt.Println("top rank", rank)

	for _, k := range keys {
		nrOfCombinations := len(mapOfCardsByPower[k])

		if nrOfCombinations > 1 {
			// sort the combinations by the highest card going through each card
			sort.Slice(mapOfCardsByPower[k], func(i, j int) bool {
				for card := 0; card < 5; card++ {
					if cardToValue(string(mapOfCardsByPower[k][i][card])) > cardToValue(string(mapOfCardsByPower[k][j][card])) {
						return true
					} else if cardToValue(string(mapOfCardsByPower[k][i][card])) < cardToValue(string(mapOfCardsByPower[k][j][card])) {
						return false
					}
				}
				return false
			})

			// print the combinations
			for _, combination := range mapOfCardsByPower[k] {
				totalWinningHands += mapOfBidsByHand[combination] * rank
				rank--
			}
		} else {
			totalWinningHands += mapOfBidsByHand[mapOfCardsByPower[k][0]] * rank
			rank--
		}
	}

	return totalWinningHands
}

var mapOfPowers = map[string]int{
	"5":     1,
	"14":    2,
	"23":    3,
	"113":   4,
	"122":   5,
	"1112":  6,
	"11111": 7,
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	fmt.Println(part1(dat))
}

func cardToValue(card string) int {
	switch card {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	case "T":
		return 10
	default:
		value, err := strconv.Atoi(card)
		check(err)
		return value
	}
}
