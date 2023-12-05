package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/mpvl/unique"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// common

// part 1
func part1(conversionLines []string, seeds []int) {
	conversionGroups := conversionLines[1:]
	locations := make([]int, 0)

	// groups
	for _, seed := range seeds {
		x := seed

		for _, conversionGroup := range conversionGroups {
			conversionDetails := strings.Split(conversionGroup, ":\n")
			conversionMap := conversionDetails[1]

			mappings := strings.Split(conversionMap, "\n")
			twoD := make([][]int, 0)

			for _, conversion := range mappings {
				conversionDetails := strings.Split(conversion, " ")
				destinationRangeStart, err := strconv.Atoi(conversionDetails[0])
				check(err)
				sourceRangeStart, err := strconv.Atoi(conversionDetails[1])
				check(err)
				rangeLength, err := strconv.Atoi(conversionDetails[2])
				check(err)

				twoD = append(twoD, []int{destinationRangeStart, sourceRangeStart, rangeLength})
			}

			sort.SliceStable(twoD, func(i, j int) bool {
				return twoD[i][1] < twoD[j][1]
			})

			for _, conversionDetails := range twoD {
				destinationRangeStart := conversionDetails[0]
				sourceRangeStart := conversionDetails[1]
				rangeLength := conversionDetails[2]
				sourceRangeEnd := sourceRangeStart + rangeLength

				if x >= sourceRangeStart && x < sourceRangeEnd {
					x = x + (destinationRangeStart - sourceRangeStart)
					break
				}
			}
		}
		locations = append(locations, x)
	}

	fmt.Println(slices.Min(locations))
}

// part 2
func part2(conversionLines []string, seeds []int) {
	conversionGroups := conversionLines[1:]
	locations := make([]int, 0)
	allSeeds := make([]int, 0)

	for i := 0; i < len(seeds); i += 2 {
		println("i", i, "out of", len(seeds))
		for j := seeds[i]; j < seeds[i+1]+seeds[i]; j++ {
			allSeeds = append(allSeeds, j)
		}
		unique.Sort(unique.IntSlice{P: &allSeeds})
	}

	fmt.Println("allSeeds", len(allSeeds))
	twoD := make([][]int, 0)
	m := make(map[int][][]int)

	for i, conversionGroup := range conversionGroups {
		conversionDetails := strings.Split(conversionGroup, ":\n")
		conversionMap := conversionDetails[1]

		mappings := strings.Split(conversionMap, "\n")
		twoD = make([][]int, 0)

		for _, conversion := range mappings {
			conversionDetails := strings.Split(conversion, " ")
			destinationRangeStart, err := strconv.Atoi(conversionDetails[0])
			check(err)
			sourceRangeStart, err := strconv.Atoi(conversionDetails[1])
			check(err)
			rangeLength, err := strconv.Atoi(conversionDetails[2])
			check(err)

			twoD = append(twoD, []int{destinationRangeStart, sourceRangeStart, rangeLength})
		}

		sort.SliceStable(twoD, func(i, j int) bool {
			return twoD[i][1] < twoD[j][1]
		})

		m[i] = twoD
	}

	for k, seed := range allSeeds {
		x := seed

		if k%1000 == 0 {
			fmt.Println("k", k, "out of", len(allSeeds))
		}

		for i := 0; i < len(m); i++ {

			for _, conversionDetails := range m[i] {
				destinationRangeStart := conversionDetails[0]
				sourceRangeStart := conversionDetails[1]
				rangeLength := conversionDetails[2]
				sourceRangeEnd := sourceRangeStart + rangeLength

				if x >= sourceRangeStart && x < sourceRangeEnd {
					x = x + (destinationRangeStart - sourceRangeStart)
					break
				}
			}

		}
		locations = append(locations, x)
	}

	fmt.Println(slices.Min(locations))
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	conversionLines := strings.Split(string(dat), "\n\n")
	stringSeeds := strings.Split(strings.Split(conversionLines[0], ": ")[1], " ")
	seeds := make([]int, len(stringSeeds))

	for i, seed := range stringSeeds {
		seedInt, err := strconv.Atoi(seed)
		check(err)
		seeds[i] = seedInt
	}

	// part1(conversionLines, seeds)
	part2(conversionLines, seeds)
}

// destination range start  // source range start // range length
// water-to-light map:
// 88 18 7
// 18 25 70

// [18 - (18+7-1)] = x + (88-18)
// [25 - (25+70-1)] = x + (18-25)

// destination range start  // source range start // range length
// seed-to-soil map:
// 50 98 2
// 52 50 48

// 			  //97
// [50 - (50+48-1)] = x + (52-50)
// 				//99
// [98 - (98+2-1) = x + (50-98)
