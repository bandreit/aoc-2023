package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

var questionMarkRegexp = regexp.MustCompile(`(\?+)`)
var unknownMarkRegexp = regexp.MustCompile(`(\#+)`)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertToLenThBase(n int, arr []int, len int, L int) []rune {
	combo := make([]rune, 0)
	for i := 0; i < L; i++ {
		combo = append(combo, rune(arr[n%len]))
		n /= len
	}
	return combo
}

func getAllPermutations(arr []int, length int, L int) [][]rune {
	// There can be (length)^L permutations
	combos := make([][]rune, 0)
	for i := 0; i < int(math.Pow(float64(length), float64(L))); i++ {
		// Convert i to length th base
		combos = append(combos, convertToLenThBase(i, arr, length, L))
	}

	return combos
}

type memoized struct {
	f     func(head []string, tail ...[]string) [][]string
	cache map[string][][]string
}

func memoize(f func(head []string, tail ...[]string) [][]string) *memoized {
	return &memoized{f: f, cache: make(map[string][][]string)}
}

func (m *memoized) call(head []string, tail ...[]string) [][]string {
	key := fmt.Sprintf("%v:%v", head, tail)
	if v, ok := m.cache[key]; ok {
		return v
	}
	result := m.f(head, tail...)
	m.cache[key] = result
	return result
}

func cartesian(head []string, tail ...[]string) [][]string {
	if len(tail) == 0 {
		result := make([][]string, len(head))
		for i, h := range head {
			result[i] = []string{h}
		}
		return result
	}

	remainder := cartesian(tail[0], tail[1:]...)
	result := make([][]string, 0, len(head)*len(remainder))

	for _, h := range head {
		for _, r := range remainder {
			result = append(result, append([]string{h}, r...))
		}
	}

	return result
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	lines := strings.Split(string(dat), "\n")

	s := 0
	for i, line := range lines {
		if i%10 == 0 {
			fmt.Printf("Processing line %d\n", i)
		}
		s += part1(line)
	}

	fmt.Println(s)
}

func part1(line string) int {
	data := strings.Split(line, " ")
	// repeatingString := strings.Repeat(data[0]+"?", 5)
	// originalStr := repeatingString[:len(repeatingString)-1]
	// repeatingGroupSizes := strings.Repeat(data[1]+",", 5)
	// originalGroupSizes := repeatingGroupSizes[:len(repeatingGroupSizes)-1]

	originalStr := data[0]
	originalGroupSizes := data[1]

	possibleRunes := []int{'#', '.'}
	length := len(possibleRunes)

	originalIndices := questionMarkRegexp.FindAllIndex([]byte(originalStr), -1)

	mapOfPossibleStringsAtIndex := make(map[int][]string)

	for i, index := range originalIndices {
		L := index[1] - index[0]
		allPermutations := getAllPermutations(possibleRunes, length, L)

		strings := make([]string, 0)
		for _, permutation := range allPermutations {
			strings = append(strings, string(permutation))
		}

		mapOfPossibleStringsAtIndex[i] = strings
	}

	head := mapOfPossibleStringsAtIndex[0]
	tail := make([][]string, 0)
	for i := 1; i < len(mapOfPossibleStringsAtIndex); i++ {
		tail = append(tail, mapOfPossibleStringsAtIndex[i])
	}

	s := 0
	result := memoize(cartesian).call(head, tail...)
	for _, r := range result {
		for k, index := range originalIndices {
			// fmt.Println("121 currently at index", k, " from ", len(originalIndices))
			originalStr = originalStr[:index[0]] + r[k] + originalStr[index[1]:]
			fmt.Println(originalStr)
		}
		newIndices := unknownMarkRegexp.FindAllIndex([]byte(originalStr), -1)

		newGroupSizes := make([]string, 0)
		for _, index := range newIndices {
			newLength := index[1] - index[0]
			newGroupSizes = append(newGroupSizes, fmt.Sprintf("%d", newLength))
		}

		if strings.Join(newGroupSizes, ",") == originalGroupSizes {
			s++
		}
	}

	return s
}
