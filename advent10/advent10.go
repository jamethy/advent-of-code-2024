package advent10

import (
	"advent2024/util"
	"advent2024/util/set"
)

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)

	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = util.ParseIntList(line, "")
	}

	part1Score := 0
	for i, line := range grid {
		for j := range line {
			if grid[i][j] != 0 {
				continue
			}
			nines := set.NewSet[Point]()
			collectNines(grid, i, j, nines)
			part1Score += len(nines)
		}
	}

	return part1Score, 0
}

var directions = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

type Point struct {
	i, j int
}

func collectNines(grid [][]int, i, j int, res set.Set[Point]) {
	v := grid[i][j]

	if v == 9 {
		res.Add(Point{i: i, j: j})
		return
	}

	for _, d := range directions {
		ni, nj := i+d[0], j+d[1]
		if ni < 0 || nj < 0 || ni >= len(grid) || nj >= len(grid[0]) {
			continue
		}
		if grid[ni][nj] == v+1 {
			collectNines(grid, ni, nj, res)
		}
	}
}
