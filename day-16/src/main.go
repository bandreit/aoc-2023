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

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
	UpDown
	LeftRight
)

type Point struct {
	x         int
	y         int
	direction Direction
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	txt := strings.ReplaceAll(string(dat), "\n\n", "")
	input := strings.Split(txt, "\n")

	// part 1
	comingFrom := Point{0, 0, Right}
	fmt.Println(moveThroughMatrix(input, comingFrom))

	// part 2
	allEnergizedTileComboLengths := []int{}
	startingPoints := []Point{}

	maxY := len(input)
	maxX := len(input[0])
	for i := 0; i < maxX; i++ {
		startingPoints = append(startingPoints, Point{i, 0, Down})
		startingPoints = append(startingPoints, Point{i, maxY, Up})
	}

	for i := 1; i < maxY-1; i++ {
		startingPoints = append(startingPoints, Point{0, i, Right})
		startingPoints = append(startingPoints, Point{maxX, i, Left})
	}

	for i, startingPoint := range startingPoints {
		fmt.Printf("Starting point %d of %d\n", i+1, len(startingPoints))
		allEnergizedTileComboLengths = append(allEnergizedTileComboLengths, moveThroughMatrix(input, startingPoint))
	}

	fmt.Println(slices.Max(allEnergizedTileComboLengths))
}

func determineDirection(direction Direction, tile string) Direction {
	switch direction {
	case Up:
		switch tile {
		case "/":
			return Right
		case "\\":
			return Left
		case "|":
			return Up
		case "-":
			return LeftRight

		default:
			return Up
		}
	case Down:
		switch tile {
		case "/":
			return Left
		case "\\":
			return Right
		case "|":
			return Down
		case "-":
			return LeftRight

		default:
			return Down
		}
	case Left:
		switch tile {
		case "/":
			return Down
		case "\\":
			return Up
		case "|":
			return UpDown
		case "-":
			return Left

		default:
			return Left
		}
	case Right:
		switch tile {
		case "/":
			return Up
		case "\\":
			return Down
		case "|":
			return UpDown
		case "-":
			return Right

		default:
			return Right
		}
	}

	panic("Unknown direction")
}

func determineNewPointFromDirection(point Point) Point {
	switch point.direction {
	case Up:
		point.y--
	case Down:
		point.y++
	case Left:
		point.x--
	case Right:
		point.x++
	}
	return point
}

func moveThroughMatrix(matrix []string, startingPoint Point) int {
	energizedTiles := []Point{}

	maxY := len(matrix)
	maxX := len(matrix[0])

	var navigate func(point Point)
	navigate = func(goingTo Point) {
		if goingTo.x < 0 || goingTo.x >= maxX || goingTo.y < 0 || goingTo.y >= maxY {
			return
		}

		if slices.ContainsFunc(energizedTiles, func(energizedTile Point) bool {
			return energizedTile.x == goingTo.x && energizedTile.y == goingTo.y && energizedTile.direction == goingTo.direction
		}) {
			return
		}

		energizedTiles = append(energizedTiles, goingTo)
		tile := matrix[goingTo.y][goingTo.x]
		goingTo.direction = determineDirection(goingTo.direction, string(tile))

		if goingTo.direction == LeftRight {
			goingTo.direction = Left
			goingTo = determineNewPointFromDirection(goingTo)

			navigate(goingTo)

			goingTo.direction = Right
			goingTo = determineNewPointFromDirection(goingTo)

			navigate(goingTo)
		} else if goingTo.direction == UpDown {
			goingTo.direction = Up
			goingTo = determineNewPointFromDirection(goingTo)

			navigate(goingTo)

			goingTo.direction = Down
			goingTo = determineNewPointFromDirection(goingTo)

			navigate(goingTo)
		} else {
			goingTo = determineNewPointFromDirection(goingTo)
			navigate(goingTo)
		}
	}

	navigate(startingPoint)

	uniqueEnergizedTiles := make([]Point, 0)

	for _, energizedTile := range energizedTiles {
		if !slices.ContainsFunc(uniqueEnergizedTiles, func(uniqueEnergizedTile Point) bool {
			return uniqueEnergizedTile.x == energizedTile.x && uniqueEnergizedTile.y == energizedTile.y
		}) {
			uniqueEnergizedTiles = append(uniqueEnergizedTiles, energizedTile)
		}
	}

	return len(uniqueEnergizedTiles)
}
