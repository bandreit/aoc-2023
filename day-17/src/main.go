package main

import (
	"container/heap"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.Index = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.Index = -1
	*pq = old[0 : n-1]
	return node
}

type Position struct {
	X, Y int
}

type Node struct {
	Pos       Position
	Direction Position
	Distance  int
	Straight  int
	Index     int
}

func Dijkstra(grid map[Position]int, start, target Position) []Position {
	dist := make(map[Position]int)
	prev := make(map[Position]Position)

	for pos := range grid {
		dist[pos] = int(^uint(0) >> 1) // Max int value
	}

	dist[start] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	dir1 := Position{start.X, start.Y + 1}
	dir2 := Position{start.X + 1, start.Y}

	prev[dir1] = start
	prev[dir2] = start

	heap.Push(&pq, &Node{Pos: start, Direction: dir1, Distance: 0, Straight: 0})
	heap.Push(&pq, &Node{Pos: start, Direction: dir2, Distance: 0, Straight: 0})

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)

		if node.Pos == target {
			break
		}

		for _, dir := range [3]Position{node.Direction, dirLeft(node.Direction), dirRight(node.Direction)} {
			// skip if dir is out of bounds
			nextPoint := Position{node.Pos.X + dir.X, node.Pos.Y + dir.Y}
			if _, ok := grid[nextPoint]; !ok {
				continue
			}

			alt := dist[node.Pos] + grid[nextPoint]

			newStraight := 1
			if node.Direction == dir {
				newStraight = node.Straight + 1
			}

			if ((node.Direction == dir) && node.Straight < 3) ||
				(dir != node.Direction) {
				if alt < dist[nextPoint] {
					dist[nextPoint] = alt
					prev[nextPoint] = node.Pos

					heap.Push(&pq, &Node{Pos: nextPoint, Direction: dir, Distance: alt, Straight: newStraight})
				}
			}
		}
	}

	path := make([]Position, 0)
	curr := target
	for curr != start {
		path = append(path, curr)
		curr = prev[curr]
	}
	path = append(path, start)
	reverse(path)

	return path
}

func getAvailableDirections(pos Position, prev map[Position]Position) []Position {
	neighbors := []Position{
		{pos.X - 1, pos.Y},
		{pos.X + 1, pos.Y},
		{pos.X, pos.Y - 1},
		{pos.X, pos.Y + 1},
	}

	prevNode, ok := prev[pos]
	if ok {
		neighbors = slices.DeleteFunc(neighbors, func(neighbor Position) bool {
			return neighbor == prevNode
		})
	}

	return neighbors
}

func reverse(path []Position) {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
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

	grid := map[Position]int{}
	start := Position{0, 0}
	target := Position{0, 0}

	for i, line := range lines {
		for j, char := range line {
			number, err := strconv.Atoi(string(char))
			check(err)

			grid[Position{i, j}] = number
			target = Position{i, j}
		}
	}

	path := Dijkstra(grid, start, target)

	fmt.Println("Shortest path:")
	for _, pos := range path {
		fmt.Println(pos)
	}

	for _, pos := range path {
		grid[pos] = 0
	}

	for i, line := range lines {
		for j := range line {
			fmt.Printf("%d", grid[Position{i, j}])
		}
		fmt.Println()
	}
}

func dirLeft(p Position) Position {
	return Position{p.Y, -p.X}
}

func dirRight(p Position) Position {
	return Position{-p.Y, p.X}
}
