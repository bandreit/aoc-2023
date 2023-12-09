package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	lines := make([][]int, 0)

	for _, stringLine := range strings.Split(string(dat), "\n") {
		if stringLine != "" {
			inputNumbers := make([]int, 0)
			input := strings.Split(stringLine, " ")
			for _, strings := range input {
				number, err := strconv.Atoi(strings)
				check(err)

				inputNumbers = append(inputNumbers, number)
			}
			lines = append(lines, inputNumbers)
		}
	}

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(lines [][]int) int {
	sumOfNextValues := 0
	for _, historyRecords := range lines {
		nextValue := calculateNext(historyRecords, 1)
		sumOfNextValues += nextValue
	}

	return sumOfNextValues
}

func part2(lines [][]int) int {
	sumOfNextValues := 0
	for _, historyRecords := range lines {
		nextValue := calculateNext(historyRecords, -1)
		sumOfNextValues += nextValue
	}

	return sumOfNextValues
}

func calculateNext(historyRecords []int, nTh int) int {
	lastValuesForSequences := getNthValuesForEachSequence(historyRecords, nTh)

	nextValue := 0
	for i := len(lastValuesForSequences) - 1; i >= 0; i-- {
		if nTh == -1 {
			nextValue = lastValuesForSequences[i] - nextValue
		} else {
			nextValue += lastValuesForSequences[i]
		}
	}

	return nextValue
}

func getNthValuesForEachSequence(historyRecords []int, nTh int) []int {
	lastValues := make([]int, 0)
	allRecordsAreZero := false

	for !allRecordsAreZero {
		newHistoryLength := len(historyRecords) - 1
		newHistoryRecords := make([]int, newHistoryLength)

		for i := 0; i < newHistoryLength; i++ {
			newHistoryRecords[i] = historyRecords[i+1] - historyRecords[i]
		}

		index := len(newHistoryRecords)
		if nTh == -1 {
			index = 0
		}
		lastValues = append(lastValues, historyRecords[index])
		allRecordsAreZero = checkAllRecordsAreZero(newHistoryRecords)
		historyRecords = newHistoryRecords
	}

	return lastValues
}

func checkAllRecordsAreZero(newHistoryRecords []int) bool {
	for _, record := range newHistoryRecords {
		if record != 0 {
			return false
		}
	}

	return true
}

// z6   z5   x4   x3  x2  x1  x
// 	 y5   y4   y3   y2   y1  y
// 		 z4   z3   z2   z1  z

// 0   3   6   9  12  15  x
//   3   3   3   3   3  y
//     0   0   0   0   z

// 10  13  16  21  30  45  a
//    3   3   5   9  15  b
//      0   2   4   6   c
//        2   2   2   d
//          0   0   e

// 10  13  16  21  30  45  a
//    									45-a

// 10  13  a4  a3   a2  a1  a
//    3   3   b3  b2  b1  b
//      0   2   c2  c1  c
//        2   2   d1   d
//          0   0   e

// a = a1 + b == a1 + b1 + c1 + d1 + e
// b = b1 + c == b1 + c1 + d1 + e
// c = c1 + d == c1 + d1 + e
// d = d1 + e

// x = x1 + y
// y = (x1-x2) + z
// z = y - y1

// x = 15 + y
// y = 3 + z
// z = 0

// 10  13  16  21  30  45  68
// 3   3   5   9   15  23
// 0   2   4   6   8
// 2   2   2   2
// 0   0   0

func check(e error) {
	if e != nil {
		panic(e)
	}
}
