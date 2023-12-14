package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)
	patterns := strings.Split(string(dat), "\n\n")

	s := 0
	for _, patternRaw := range patterns {
		pattern := strings.Split(patternRaw, "\n")

		// fmt.Println(i)

		mirrorRow := getReflectionRowIndex(pattern)
		if mirrorRow != -1 {
			s += (mirrorRow + 1) * 100
			fmt.Println("row", mirrorRow)
			continue
		}

		mirrorColumn := getReflectionColumnIndex(pattern)
		if mirrorColumn != -1 {
			s += mirrorColumn + 1
			fmt.Println("column", mirrorColumn)
		}

		fmt.Println()
	}

	fmt.Println(s)
}

func getAllColumns(pattern []string) []string {
	allColumnsRaw := make([][]string, 0)
	allColumns := make([]string, 0)

	for i := 0; i < len(pattern[0]); i++ {
		column := make([]string, 0)

		for _, row := range pattern {
			column = append(column, string(row[i]))
		}

		allColumnsRaw = append(allColumnsRaw, column)
	}

	for _, column := range allColumnsRaw {
		allColumns = append(allColumns, strings.Join(column, ""))
	}

	return allColumns
}

func getReflectionRowIndex(pattern []string) int {
	exclude := doWork1(pattern, "row")
	return doWork2(pattern, "row", exclude)

	// return doWork1(pattern, "row")
}

func getReflectionColumnIndex(pattern []string) int {
	allColumns := getAllColumns(pattern)
	exclude := doWork1(allColumns, "columns")
	return doWork2(allColumns, "columns", exclude)

	// return doWork1(allColumns, "columns")
}

func getDuplicatedSliceIndexes(allSlices []string) []int {
	duplicatedSliceIndexes := make([]int, 0)

	for i := 0; i < len(allSlices)-1; i++ {
		if allSlices[i] == allSlices[i+1] {
			if !slices.Contains(duplicatedSliceIndexes, i) {
				duplicatedSliceIndexes = append(duplicatedSliceIndexes, i)
			}
		}
	}

	return duplicatedSliceIndexes
}

func doWork1(pattern []string, typeOfArr string) int {
	indexes := getDuplicatedSliceIndexes(pattern)

	for _, row := range indexes {
		left := row - 1
		right := row + 2

		fullLength := len(pattern)

		if row == 0 || row == fullLength-2 {
			return row
		}

		reachedEnd := left < 0 || right >= len(pattern)
		if reachedEnd {
			continue
		}

		columnsAreEqual := pattern[left] == pattern[right]

		for !reachedEnd && columnsAreEqual {
			columnsAreEqual = pattern[left] == pattern[right]

			left--
			right++
			reachedEnd = left < 0 || right >= len(pattern)
		}

		if columnsAreEqual {
			// fmt.Println("found ", typeOfArr, row)
			return row
		}
	}

	return -1
}

func doWork2(pattern []string, typeOfArr string, exclude int) int {
	for row := range pattern {
		if row == exclude {
			// fmt.Println("Excluding ", row)
			continue
		}
		if row == len(pattern)-1 {
			continue
		}

		wasDiffUsed := false
		isDiff1 := getDiff(pattern[row], pattern[row+1]) == 1
		areTheyEqual := pattern[row] == pattern[row+1]

		if !(isDiff1 || areTheyEqual) {
			continue
		}

		if isDiff1 {
			wasDiffUsed = true
		}

		up := row - 1
		down := row + 2

		reachedEnd := up < 0 || down >= len(pattern)
		if reachedEnd {
			// fmt.Println("found row", row)
			return row
		}

		isDiff1 = getDiff(pattern[up], pattern[down]) == 1
		wasDiffUsed = false

		rowsAreEqual := pattern[up] == pattern[down] || isDiff1

		for !reachedEnd && rowsAreEqual && !wasDiffUsed {
			isDiff1 = getDiff(pattern[up], pattern[down]) == 1
			rowsAreEqual = pattern[up] == pattern[down] || isDiff1
			if isDiff1 {
				wasDiffUsed = true
			}

			up--
			down++
			reachedEnd = up < 0 || down >= len(pattern)
		}

		if rowsAreEqual {
			// fmt.Println("found ", typeOfArr, row)
			return row
		}
	}

	return -1
}

// func getReflectionColumnIndex(pattern []string) int {
// 	allColumns := getAllColumns(pattern)

// 	for column := range allColumns {
// 		if column == len(allColumns)-1 {
// 			continue
// 		}

// 		wasDiffUsed := false
// 		isDiff1 := getDiff(allColumns[column], allColumns[column+1]) == 1
// 		areTheyEqual := allColumns[column] == allColumns[column+1]

// 		if !(isDiff1 || areTheyEqual) {
// 			continue
// 		}

// 		if isDiff1 {
// 			wasDiffUsed = true
// 		}

// 		left := column - 1
// 		right := column + 2

// 		reachedEnd := left < 0 || right >= len(allColumns)
// 		if reachedEnd {
// 			fmt.Println("found column", column)
// 			return column
// 		}

// 		isDiff1 = getDiff(pattern[left], pattern[right]) == 1
// 		if isDiff1 {
// 			wasDiffUsed = true
// 		}

// 		rowsAreEqual := pattern[up] == pattern[down] || isDiff1

// 		for !reachedEnd && columnsAreEqual {
// 			columnsAreEqual = getDiff(allColumns[left], allColumns[right]) == 1

// 			left--
// 			right++
// 			reachedEnd = left < 0 || right >= len(allColumns)
// 		}

// 		if columnsAreEqual {
// 			fmt.Println("found column", column)
// 			return column
// 		}
// 	}

// 	return -1
// }

func getDiff(left string, right string) int {
	diffs := 0
	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			diffs++
		}
	}

	return diffs
}
