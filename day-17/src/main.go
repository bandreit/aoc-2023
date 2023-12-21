package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HeatState struct {
	point, dir Point
	streak     int
}

type Point struct {
	x int
	y int
}

func Dijkstra(grid map[Point]int, start, target Point) int {
	dir1 := Point{start.x, start.y + 1}
	dir2 := Point{start.x + 1, start.y}
	pointsToCheck := []HeatState{{start, dir1, 0}, {start, dir2, 0}}
	dist := map[HeatState]int{{start, Point{0, 0}, 0}: 0}
	minHeatLoss := int(^uint(0) >> 1)

	for len(pointsToCheck) > 0 {
		node := pointsToCheck[0]
		pointsToCheck = pointsToCheck[1:]

		if node.point == target && node.streak >= 4 {
			minHeatLoss = min(minHeatLoss, dist[node])
		}

		for _, dir := range [3]Point{node.dir, dirLeft(node.dir), dirRight(node.dir)} {
			nextPoint := Point{node.point.x + dir.x, node.point.y + dir.y}
			if _, ok := grid[nextPoint]; !ok {
				continue
			}

			alt := dist[node] + grid[nextPoint]
			newStraight := 1
			if node.dir == dir {
				newStraight = node.streak + 1
			}

			if ((node.dir == dir) && node.streak < 10) ||
				(dir != node.dir && node.streak >= 4) {
				nextState := HeatState{point: nextPoint, dir: dir, streak: newStraight}
				if val, found := dist[nextState]; !found || val > alt {
					dist[nextState] = alt
					pointsToCheck = append(pointsToCheck, nextState)
				}
			}
		}
	}

	return minHeatLoss
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	txt := strings.ReplaceAll(string(dat), "\n\n", "")
	lines := strings.Split(txt, "\n")

	grid := map[Point]int{}
	start := Point{0, 0}
	target := Point{0, 0}

	for i, line := range lines {
		for j, char := range line {
			number, err := strconv.Atoi(string(char))
			check(err)

			grid[Point{i, j}] = number
			target = Point{i, j}
		}
	}

	minPath := Dijkstra(grid, start, target)
	fmt.Println(minPath)
}

func dirLeft(p Point) Point {
	return Point{p.y, -p.x}
}

func dirRight(p Point) Point {
	return Point{-p.y, p.x}
}
