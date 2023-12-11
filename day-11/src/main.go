package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Point struct {
	i int
	j int
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	twoDArr, coords := build2DArrayFromLines(string(dat))

	for _, line := range twoDArr {
		fmt.Println(string(line))
	}

	combos := make([][]Point, 0)
	comb([]Point{}, coords, 2, &combos)

	s := 0

	for _, combo := range combos {
		distance := getManhattanDistance(combo[0], combo[1])
		s += distance
	}

	fmt.Println(s)
}

func getManhattanDistance(p1 Point, p2 Point) int {
	return int(math.Abs(float64(p1.i-p2.i)) + math.Abs(float64(p1.j-p2.j)))
}

func comb(sofar []Point, rest []Point, n int, combos *[][]Point) {
	if n == 0 {
		*combos = append(*combos, sofar)
	} else {
		for i := 0; i < len(rest); i++ {
			comb(append(sofar, rest[i]), rest[i+1:], n-1, combos)
		}
	}
}

func build2DArrayFromLines(s string) ([][]rune, []Point) {
	lines := strings.Split(s, "\n")
	twoD := make([][]rune, 0)
	coords := make([]Point, 0)

	for i := 0; i < len(lines); i++ {
		row := []rune(lines[i])
		twoD = append(twoD, row)

		if !slices.Contains(row, '#') {
			twoD = append(twoD, row)
		}
	}

	colIndexesToRepeat := make([]int, 0)
	for j := 0; j < len(twoD[0]); j++ {
		col := make([]rune, 0)
		for i := 0; i < len(twoD); i++ {
			col = append(col, twoD[i][j])
		}

		if !slices.Contains(col, '#') {
			colIndexesToRepeat = append(colIndexesToRepeat, j)
		}
	}

	fmt.Println(colIndexesToRepeat)
	for k, colIndex := range colIndexesToRepeat {
		for i := 0; i < len(twoD); i++ {
			twoD[i] = insert(twoD[i], colIndex+1+k, '.')
		}
	}

	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[i]); j++ {
			if twoD[i][j] == '#' {
				coords = append(coords, Point{i: i, j: j})
			}
		}
	}

	return twoD, coords
}

func insert(a []rune, index int, value rune) []rune {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
