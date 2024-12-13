package advent12

import (
	"advent2024/util"
	"advent2024/util/set"
)

var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

type Point struct {
	i, j int
}

func pointMatches(lines []string, i, j int, c uint8) bool {
	if i < 0 || i > len(lines)-1 || j < 0 || j > len(lines[0])-1 {
		return false
	}
	return lines[i][j] == c
}

func recursiveCalc(lines []string, i, j int, visited set.Set[Point]) (int, int, int) {
	if i < 0 || j < 0 || i > len(lines)-1 || j > len(lines[0])-1 {
		return 0, 0, 0
	}
	visited.Add(Point{i: i, j: j})

	area := 1
	perim := 0
	corners := 0

	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]
		if visited.Has(Point{i: ni, j: nj}) {
			continue
		}
		if !pointMatches(lines, ni, nj, lines[i][j]) {
			perim++
			ri, rj := i+dir[1], j-dir[0]
			if !pointMatches(lines, ri, rj, lines[i][j]) {
				corners++
			} else if pointMatches(lines, ri+dir[0], rj+dir[1], lines[i][j]) {
				corners++
			}
			continue
		}
		a, p, c := recursiveCalc(lines, ni, nj, visited)
		area += a
		perim += p
		corners += c
	}
	return area, perim, corners
}

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)

	visited := set.NewSet[Point]()

	part1Cost := 0
	part2Cost := 0
	for i, line := range lines {
		for j := range line {
			point := Point{i: i, j: j}
			if visited.Has(point) {
				continue
			}
			v := set.NewSet[Point]()
			a, p, c := recursiveCalc(lines, i, j, v)
			part1Cost += a * p
			part2Cost += a * c
			visited.AddAll(v)
		}
	}

	return part1Cost, part2Cost
}
