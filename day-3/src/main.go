package main

import (
	"fmt"
	"os"
	"slices"
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

func buildCoordsToCheckForSpan(row int, startCol int, endCol int) [][]int {
	coordsToCheck := [][]int{
		{row - 1, startCol - 1},
		{row - 1, startCol},
		{row - 1, endCol + 1},
		{row - 1, endCol},
		{row, startCol - 1},
		{row, endCol + 1},
		{row + 1, startCol},
		{row + 1, startCol - 1},
		{row + 1, endCol + 1},
		{row + 1, endCol},
	}

	for j := startCol + 1; j < endCol; j++ {
		coordsToCheck = append(coordsToCheck, []int{row - 1, j})
		coordsToCheck = append(coordsToCheck, []int{row + 1, j})
	}

	return coordsToCheck
}

func buildCoordsToCheckForCell(row int, col int) [][]int {
	coordsToCheck := [][]int{
		{row - 1, col - 1},
		{row - 1, col},
		{row - 1, col + 1},
		{row, col - 1},
		{row, col + 1},
		{row + 1, col - 1},
		{row + 1, col},
		{row + 1, col + 1},
	}

	return coordsToCheck
}

func isAreaAdjacentToSymbol(twoDArr [][]rune, row int, startCol int, endCol int) bool {
	coordsToCheck := buildCoordsToCheckForSpan(row, startCol, endCol)

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

func findFullNumber(twoDArr [][]rune, row int, col int) []rune {
	number := []rune{}

	backWardCol := col

	for backWardCol >= 0 && unicode.IsDigit(twoDArr[row][backWardCol]) {
		number = append([]rune{twoDArr[row][backWardCol]}, number...)
		backWardCol--
	}

	forwardCol := col + 1

	for forwardCol < len(twoDArr[row]) && unicode.IsDigit(twoDArr[row][forwardCol]) {
		number = append(number, twoDArr[row][forwardCol])
		forwardCol++
	}

	return number
}

func getAdjacentNumbers(twoDArr [][]rune, row int, col int) []int {
	coordsToCheck := buildCoordsToCheckForCell(row, col)
	foundNumbers := []int{}

	for _, coords := range coordsToCheck {
		x := coords[0]
		y := coords[1]
		rune := twoDArr[x][y]
		if unicode.IsDigit(rune) {
			fullNumber := findFullNumber(twoDArr, x, y)
			subString := string(fullNumber)
			partNumber, err := strconv.Atoi(subString)
			check(err)

			if !slices.Contains(foundNumbers, partNumber) {
				foundNumbers = append(foundNumbers, partNumber)
			}
		}
	}

	return foundNumbers
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

func findGears(twoDArr [][]rune) []int {
	gears := []int{}

	m := len(twoDArr)
	n := len(twoDArr[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			char := twoDArr[i][j]
			isStar := char == '*'

			if isStar {
				adjacentNumbers := getAdjacentNumbers(twoDArr, i, j)
				if len(adjacentNumbers) == 2 {
					fmt.Println("found gear at", i, j)
					gears = append(gears, adjacentNumbers[0]*adjacentNumbers[1])
				}
			}
		}
	}

	return gears
}

func sumOfArray(partNumbers []int) int {
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
	sumOfPartNumbers := sumOfArray(partNumbers)
	println(sumOfPartNumbers)

	gears := findGears(engineSchema)
	sumOfGears := sumOfArray(gears)
	println(sumOfGears)
}
