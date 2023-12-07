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

func cardToValue(card string) int {
	switch card {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 1
	case "T":
		return 10
	default:
		value, err := strconv.Atoi(card)
		check(err)
		return value
	}
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

	fmt.Println("Part 2:", part1(dat))
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

		mapOfOccurrencesByCard := make(map[string]int)
		occurrencesArray := make([]int, 0)

		for _, card := range cards {
			occurrences := strings.Count(hand, string(card))
			mapOfOccurrencesByCard[string(card)] = occurrences
		}

		for _, value := range mapOfOccurrencesByCard {
			occurrencesArray = append(occurrencesArray, value)
		}

		sort.Ints(occurrencesArray)
		// ok till here

		if strings.Contains(hand, "J") {
			numberOfJacks := strings.Count(hand, "J")

			arrLen := len(occurrencesArray)
			if arrLen > 1 {
				if (occurrencesArray[arrLen-1]) == numberOfJacks {
					occurrencesArray[arrLen-1] = occurrencesArray[arrLen-1] + occurrencesArray[arrLen-2]
					occurrencesArray[arrLen-2] = 0
				} else if occurrencesArray[arrLen-1] == occurrencesArray[arrLen-2] {
					if numberOfJacks == 1 {
						occurrencesArray = []int{2, 3}
					}
				} else {
					occurrencesArray[arrLen-1] = occurrencesArray[arrLen-1] + numberOfJacks
					occurrencesArray[arrLen-2] = occurrencesArray[arrLen-2] - numberOfJacks
				}
			}

		}

		for i := 0; i < len(occurrencesArray); i++ {
			if occurrencesArray[i] == 0 {
				occurrencesArray = append(occurrencesArray[:i], occurrencesArray[i+1:]...)
			}
		}

		combo := make([]string, 0)

		for _, value := range occurrencesArray {
			str := strconv.Itoa(value)
			combo = append(combo, str)
		}
		mapOfBidsByHand[hand] = bid
		pokerHand := strings.Join(combo, "")
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

// 2345A 1
// J345A 2
// 2345J 3
// 32T3K 4 - one pair
// KK677 5 - two pairs
// T3Q33 6 - three of a kind
// Q2KJJ 7 - three of a kind

// T3T3J = T3T3T 10 - full house - should be here
// Q2Q2Q 8 - full house
// 2AAAA 9 - four of a kind
// T3T3J 10 - full house

// T55J5 11 - four of a kind
// QQQJA 12 - four of a kind
// KTJJT 13 - four of a kind
// JJJJJ 14
// JJJJ2 15
// JAAAA 16
// 2JJJJ 17
// AAAAJ 18
// AAAAA 19
