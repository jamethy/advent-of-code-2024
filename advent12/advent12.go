package advent12

import (
	"advent2024/util"
	"advent2024/util/set"
)

type Node struct {
	i, j   int
	letter rune
	id     int
	left   *Node
	right  *Node
	up     *Node
	down   *Node
}

var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

type Point struct {
	i, j int
}

func recursiveCalc(lines []string, i, j int, visited set.Set[Point]) (int, int) {
	visited.Add(Point{i: i, j: j})

	area := 1
	perim := 0

	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]
		if visited.Has(Point{i: ni, j: nj}) {
			continue
		}
		if ni < 0 || ni > len(lines)-1 || nj < 0 || nj > len(lines[0])-1 || lines[ni][nj] != lines[i][j] {
			perim++
			continue
		}
		a, p := recursiveCalc(lines, ni, nj, visited)
		area += a
		perim += p
	}
	return area, perim
}

func calcFor(lines []string, i, j int) (int, set.Set[Point]) {
	visited := set.NewSet[Point]()
	a, p := recursiveCalc(lines, i, j, visited)
	return a * p, visited
}

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)

	visited := set.NewSet[Point]()

	cost := 0
	for i, line := range lines {
		for j := range line {
			point := Point{i: i, j: j}
			if visited.Has(point) {
				continue
			}
			c, v := calcFor(lines, i, j)
			cost += c
			visited.AddAll(v)
		}
	}

	return cost, 0
}
