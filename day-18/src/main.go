package main

import (
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type point struct {
	x int
	y int
}

type instruction struct {
	dir    string
	meters int
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	txt := strings.ReplaceAll(string(dat), "\n\n", "")
	lines := strings.Split(txt, "\n")

	instructions := make([]instruction, 0)

	for _, line := range lines {
		data := strings.SplitAfter(line, " ")
		number, err := strconv.Atoi(strings.Trim(data[1], " "))
		check(err)

		instructions = append(instructions, instruction{
			dir:    strings.Trim(data[0], " "),
			meters: number,
		})
	}

	current := point{0, 0}

	path := make([]point, 0)
	for _, instruction := range instructions {
		if instruction.dir == "R" {
			for i := 0; i < instruction.meters; i++ {
				current = point{current.x, current.y + 1}
				path = append(path, current)
			}
		}

		if instruction.dir == "D" {
			for i := 0; i < instruction.meters; i++ {
				current = point{current.x + 1, current.y}
				path = append(path, current)
			}
		}

		if instruction.dir == "L" {
			for i := 0; i < instruction.meters; i++ {
				current = point{current.x, current.y - 1}
				path = append(path, current)
			}
		}

		if instruction.dir == "U" {
			for i := 0; i < instruction.meters; i++ {
				current = point{current.x - 1, current.y}
				path = append(path, current)
			}
		}
	}

	maxRow := 0
	minRow := 0

	maxCol := 0
	minCol := 0

	for _, p := range path {
		if p.x > maxRow {
			maxRow = p.x
		}

		if p.x < minRow {
			minRow = p.x
		}

		if p.y > maxCol {
			maxCol = p.y
		}

		if p.y < minCol {
			minCol = p.y
		}
	}

	for i, p := range path {
		p.x -= minRow
		p.y -= minCol
		path[i] = p
	}

	rows := maxRow - minRow + 1
	cols := maxCol - minCol + 1
	grid := make([][]rune, rows)
	for i := range grid {
		grid[i] = make([]rune, cols)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for _, p := range path {
		grid[p.x][p.y] = '#'
	}

	for _, line := range grid {
		println(string(line))
	}

	trappedDots := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// if the point is not a wall
			if grid[i][j] != '#' {
				isPointInPath := isPointInPath(i, j, path)

				if isPointInPath {
					trappedDots++
				}
			}
		}
	}

	println(trappedDots + len(path))
}

func isPointInPath(x, y int, mazePath []point) bool {
	num := len(mazePath)
	j := num - 1
	c := false
	for i := 0; i < num; i++ {
		if x == mazePath[i].x && y == mazePath[i].y {
			// point is a corner
			return true
		}
		if (mazePath[i].y > y) != (mazePath[j].y > y) {
			slope := (x-mazePath[i].x)*(mazePath[j].y-mazePath[i].y) - (mazePath[j].x-mazePath[i].x)*(y-mazePath[i].y)
			if slope == 0 {
				// point is on boundary
				return true
			}
			if (slope < 0) != (mazePath[j].y < mazePath[i].y) {
				c = !c
			}
		}
		j = i
	}
	return c
}
