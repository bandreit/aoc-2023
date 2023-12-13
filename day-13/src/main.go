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
	// LevenshteinDistance("acc", "abc")

	dat, err := os.ReadFile("../input.txt")
	check(err)
	patterns := strings.Split(string(dat), "\n\n")

	s := 0
	for _, patternRaw := range patterns {
		pattern := strings.Split(patternRaw, "\n")

		mirrorRow := getReflectionRowIndex(pattern)
		mirrorColumn := getReflectionColumnIndex(pattern)

		if mirrorRow != -1 {
			s += (mirrorRow + 1) * 100
		}

		if mirrorColumn != -1 {
			s += mirrorColumn + 1
		}
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

func getReflectionRowIndex(pattern []string) int {
	duplicatedRowsIndexes := getDuplicatedSliceIndexes(pattern)

	for _, row := range duplicatedRowsIndexes {
		up := row - 1
		down := row + 2

		if row == 0 || row == len(pattern)-2 {
			return row
		}

		reachedEnd := up < 0 || down >= len(pattern)
		if reachedEnd {
			continue
		}

		rowsAreEqual := pattern[up] == pattern[down]

		for !reachedEnd && rowsAreEqual {
			rowsAreEqual = pattern[up] == pattern[down]

			up--
			down++
			reachedEnd = up < 0 || down >= len(pattern)
		}

		if rowsAreEqual {
			return row
		}
	}

	return -1
}

func getReflectionColumnIndex(pattern []string) int {
	allColumns := getAllColumns(pattern)
	duplicatedColumnIndexes := getDuplicatedSliceIndexes(allColumns)

	for _, column := range duplicatedColumnIndexes {
		left := column - 1
		right := column + 2

		if column == 0 || column == len(pattern[0])-2 {
			return column
		}

		reachedEnd := left < 0 || right >= len(allColumns)
		if reachedEnd {
			continue
		}

		columnsAreEqual := allColumns[left] == allColumns[right]

		for !reachedEnd && columnsAreEqual {
			columnsAreEqual = allColumns[left] == allColumns[right]

			left--
			right++
			reachedEnd = left < 0 || right >= len(allColumns)
		}

		if columnsAreEqual {
			return column
		}
	}

	return -1
}

func LevenshteinDistance(s string, t string) int {
	// create two work vectors of integer distances
	v0 := make([]int, len(t)+1)
	v1 := make([]int, len(t)+1)

	// initialize v0 (the previous row of distances)
	// this row is A[0][i]: edit distance from an empty s to t;
	// that distance is the number of characters to append to s to make t.
	for i := 0; i <= len(t); i++ {
		v0[i] = i
	}

	for i := 0; i < len(s); i++ {
		// calculate v1 (current row distances) from the previous row v0

		// first element of v1 is A[i + 1][0]
		//   edit distance is delete (i + 1) chars from s to match empty t
		v1[0] = i + 1

		// use formula to fill in the rest of the row
		for j := 0; j < len(t); j++ {
			// calculating costs for A[i + 1][j + 1]
			deletionCost := v0[j+1] + 1
			insertionCost := v1[j] + 1
			substitutionCost := v0[j]
			if s[i] != t[j] {
				substitutionCost++
			}

			v1[j+1] = min(deletionCost, insertionCost, substitutionCost)
		}

		// copy v1 (current row) to v0 (previous row) for next iteration
		// since data in v1 is always invalidated, a swap without copy could be more efficient
		v0, v1 = v1, v0
	}

	fmt.Println(v0)
	// after the last swap, the results of v1 are now in v0
	return v0[len(t)]
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	} else {
		return c
	}
}
