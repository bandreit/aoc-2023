package main

import (
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func prefillSliceWithRunes(s []rune, r rune, n int) []rune {
	for i := 0; i < n; i++ {
		s = append(s, r)
	}

	return s
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

func build2DArrayFromLines(s string) [][]rune {
	lines := strings.Split(s, "\n")
	m := len(lines)

	twoD := make([][]rune, m)
	for i, line := range lines {
		twoD[i] = []rune(line)
	}

	return twoD
}

func isAreaAdjacentToSymbol(twoDArr [][]rune, row int, startCol int, endCol int) bool {
	coordsToCheck := [][]int{
		{row - 1, startCol - 1},
		{row - 1, startCol},
		{row, startCol - 1},
		{row + 1, startCol - 1},
		{row + 1, startCol},
		{row - 1, endCol + 1},
		{row - 1, endCol},
		{row, endCol + 1},
		{row + 1, endCol + 1},
		{row + 1, endCol},
	}

	for j := startCol + 1; j < endCol; j++ {
		coordsToCheck = append(coordsToCheck, []int{row - 1, j})
		coordsToCheck = append(coordsToCheck, []int{row + 1, j})
	}

	for _, coords := range coordsToCheck {
		x := coords[0]
		y := coords[1]
		rune := twoDArr[x][y]
		if rune != '.' {
			return true
		}
	}

	return false
}

func findPartNumbers(twoDArr [][]rune) []int {
	partNumbers := []int{}

	m := len(twoDArr)
	n := len(twoDArr[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			char := twoDArr[i][j]
			isDigit := unicode.IsDigit(char)

			number := []rune{}

			if isDigit {
				startJ := j
				endJ := j

				for j < n && unicode.IsDigit(twoDArr[i][j]) {
					number = append(number, twoDArr[i][j])
					endJ = j
					j++
				}

				if isAreaAdjacentToSymbol(twoDArr, i, startJ, endJ) {
					subString := string(number)
					partNumber, err := strconv.Atoi(subString)
					check(err)

					partNumbers = append(partNumbers, partNumber)
				}
			}

		}
	}

	return partNumbers
}

func sumOfPartNumbers(partNumbers []int) int {
	sum := 0
	for _, partNumber := range partNumbers {
		sum += partNumber
	}

	return sum
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	input := string(dat)
	engineSchema := build2DArrayFromLines(input)
	engineSchema = addPaddingTo2DArray(engineSchema, '.')

	partNumbers := findPartNumbers(engineSchema)
	sum := sumOfPartNumbers(partNumbers)

	println(sum)
}
