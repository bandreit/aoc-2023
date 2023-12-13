package main

import (
	"fmt"
	"os"
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
	for i, patternRaw := range patterns {
		pattern := strings.Split(patternRaw, "\n")

		fmt.Println(i)
		mirrorRow := getReflectionRowIndex(pattern)
		mirrorColumn := getReflectionColumnIndex(pattern)

		if mirrorRow != -1 {
			s += (mirrorRow + 1) * 100
			continue
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

func getReflectionRowIndex(pattern []string) int {
	return doWork(pattern, "row")
}

func getReflectionColumnIndex(pattern []string) int {
	allColumns := getAllColumns(pattern)
	return doWork(allColumns, "columns")
}

func doWork(pattern []string, typeOfArr string) int {
	for row := range pattern {
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
			fmt.Println("found row", row)
			return row
		}

		isDiff1 = getDiff(pattern[up], pattern[down]) == 1
		if isDiff1 {
			wasDiffUsed = true
		}

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
			fmt.Println("found ", typeOfArr, row)
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
