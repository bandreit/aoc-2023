package main

import (
	"fmt"
	"hash/fnv"
	"os"
	"slices"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, yl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	twodArr := make([][]string, 0)
	for _, line := range strings.Split(string(dat), "\n") {
		twodArr = append(twodArr, strings.Split(line, ""))
	}

	// for _, line := range twodArr {
	// 	fmt.Println(strings.Join(line, ""))
	// }
	// fmt.Println()

	hashes := make([]int, 0)
	loop := -1

	cyclesToDo := 1000000000
	for cycle := 0; cycle < cyclesToDo; cycle++ {
		fullString := ""
		for _, line := range twodArr {
			fullString += strings.Join(line, "")
		}
		hash := hash(fullString)

		if loop == -1 {
			if slices.Contains(hashes, int(hash)) {
				fmt.Println("Found a loop at cycle: ", cycle)
				loop = cycle

				toSkipTo := cyclesToDo - (cyclesToDo % cycle)
				skipping := toSkipTo - cycle
				fmt.Println("skipping to: ", toSkipTo)
				fmt.Println("Number of records skipped: ", skipping)

				cycle = toSkipTo
			}
			hashes = append(hashes, int(hash))
		}

		for k := 0; k < 4; k++ {
			// fmt.Println("Now at: ", k)

			if k == 0 || k == 2 {
				twodArr = transpose(twodArr)
			}

			if k == 2 || k == 3 {
				for _, line := range twodArr {
					for i := 0; i < len(line)/2; i++ {
						line[i], line[len(line)-i-1] = line[len(line)-i-1], line[i]
					}
				}
			}

			for i, line := range twodArr {
				row := strings.Join(line, "")
				rockIndexes := []int{0}

				for i, char := range row {
					if string(char) == "#" {
						rockIndexes = append(rockIndexes, i)
					}
				}

				if len(rockIndexes) == 1 {
					rockIndexes = append(rockIndexes, len(row))
				}

				rockIndexes = append(rockIndexes, len(row))

				for i := 0; i < len(rockIndexes)-1; i++ {
					substring := row[rockIndexes[i]:rockIndexes[i+1]]

					nrOfDots := 0
					for _, char := range substring {
						if string(char) == "." {
							nrOfDots++
						}
					}

					substring = strings.ReplaceAll(substring, ".", "")
					substring = substring + strings.Repeat(".", nrOfDots)

					row = strings.ReplaceAll(row, row[rockIndexes[i]:rockIndexes[i+1]], substring)
				}

				// for i, char := range row {
				// 	if string(char) == "O" {
				// 		s += len(row) - i
				// 	}
				// }

				twodArr[i] = strings.Split(row, "")
			}

			if k == 0 {
				twodArr = transpose(twodArr)
			}

			if k == 2 || k == 3 {
				for _, line := range twodArr {
					for i := 0; i < len(line)/2; i++ {
						line[i], line[len(line)-i-1] = line[len(line)-i-1], line[i]
					}
				}
			}

			if k == 2 {
				twodArr = transpose(twodArr)
			}
		}

		// fmt.Println("after cycle : ", cycle+1)
		// for _, line := range twodArr {
		// 	fmt.Println(strings.Join(line, ""))
		// }

		tempToCalc := transpose(twodArr)
		s := 0

		for _, line := range tempToCalc {
			for i, char := range line {
				if string(char) == "O" {
					s += len(line) - i
				}
			}
		}

		fmt.Println("s: ", s)
	}
}
