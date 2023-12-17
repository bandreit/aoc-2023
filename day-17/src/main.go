package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Point struct {
	X, Y int
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	distances := [][]int{}

	txt := strings.ReplaceAll(string(dat), "\n\n", "")
	input := strings.Split(txt, "\n")

	for _, line := range input {
		row := []int{}
		for _, char := range line {
			number, err := strconv.Atoi(string(char))
			check(err)

			row = append(row, number)
		}
		distances = append(distances, row)
	}

	rows := len(distances)
	cols := len(distances[0])

	start := Point{0, 0}
	goal := Point{rows - 1, cols - 1}

	shortestPath := AStar(start, goal, distances, ManhattanDistance)
	fmt.Println(shortestPath)

	for _, point := range shortestPath {
		distances[point.X][point.Y] = -1
	}

	for _, row := range distances {
		fmt.Println(row)
	}
}

func reconstructPath(cameFrom map[Point]Point, current Point) []Point {
	totalPath := []Point{current}
	for current != (Point{}) {
		current = cameFrom[current]
		totalPath = append([]Point{current}, totalPath...)
	}
	return totalPath
}

func AStar(start Point, goal Point, distances [][]int, h func(Point, Point) float64) []Point {
	rows := len(distances)
	cols := len(distances[0])

	openSet := map[Point]bool{start: true}
	cameFrom := make(map[Point]Point)
	gScore := make(map[Point]float64)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			gScore[Point{i, j}] = math.Inf(1)
		}
	}
	gScore[start] = 0
	fScore := make(map[Point]float64)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fScore[Point{i, j}] = math.Inf(1)
		}
	}
	fScore[start] = h(start, goal)

	for len(openSet) > 0 {
		var current Point
		minFScore := math.Inf(1)
		for node := range openSet {
			if fScore[node] < minFScore {
				minFScore = fScore[node]
				current = node
			}
		}

		if current == goal {
			return reconstructPath(cameFrom, current)
		}

		delete(openSet, current)

		neighbors := []Point{
			{current.X - 1, current.Y}, // left
			{current.X, current.Y - 1}, // up
			{current.X + 1, current.Y}, // right
			{current.X, current.Y + 1}, // down
		}

		for _, neighbor := range neighbors {
			// Skip if neighbor is out of bounds
			if neighbor.X < 0 || neighbor.X >= rows || neighbor.Y < 0 || neighbor.Y >= cols {
				continue
			}

			lastCameFroms := reconstructPath(cameFrom, current)
			last3CameFroms := make([]Point, 0)
			for i := len(lastCameFroms) - 1; i >= 0; i-- {
				last3CameFroms = append(last3CameFroms, lastCameFroms[i])
				if len(last3CameFroms) == 3 {
					break
				}
			}

			if len(last3CameFroms) == 3 {
				initialPosition := last3CameFroms[0]
				areOnTheSameAxis := areOnTheSameAxis(last3CameFroms, initialPosition, neighbor)

				if areOnTheSameAxis {
					if initialPosition.X == neighbor.X || initialPosition.Y == neighbor.Y {
						continue
					}
				}
			}

			tentativeGScore := gScore[current] + float64(distances[neighbor.X][neighbor.Y])
			if tentativeGScore < gScore[neighbor] {
				cameFrom[neighbor] = current
				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = tentativeGScore + h(neighbor, goal)
				openSet[neighbor] = true
			}
		}
	}

	return nil
}

func areOnTheSameAxis(last3CameFroms []Point, initialPosition Point, neighbor Point) bool {
	for i := 0; i < 2; i++ {
		if last3CameFroms[i].X != last3CameFroms[i+1].X && last3CameFroms[i].Y != last3CameFroms[i+1].Y {
			return false
		}
	}

	return true
}

func ManhattanDistance(p1 Point, p2 Point) float64 {
	return math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y))
}
