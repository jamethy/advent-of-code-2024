package advent18

import (
	"advent2024/util"
	"fmt"
	"strings"
)

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)

	gridSize, part1Steps := 71, 1024
	if strings.HasPrefix(inputFile, "sample") {
		gridSize, part1Steps = 7, 12
	}

	grid := make([][]int, gridSize)
	for i := range grid {
		grid[i] = make([]int, gridSize)
	}

	for i := 0; i < part1Steps; i++ {
		corruptGrid(grid, lines[i])
	}

	fillGrid(grid, 0, 0)

	part1 = grid[gridSize-1][gridSize-1]

	for i := part1Steps; i < len(lines); i++ {
		corruptGrid(grid, lines[i])
		fillGrid(grid, 0, 0)
		if grid[gridSize-1][gridSize-1] == 0 {
			part2 = lines[i]
			break
		}
		resetFill(grid)
	}

	for _, line := range grid {
		fmt.Println(line)
	}

	return part1, part2
}

var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func corruptGrid(grid [][]int, line string) {
	coords := util.StringsToInts(strings.Split(line, ","))
	grid[coords[0]][coords[1]] = -1
}

func fillGrid(grid [][]int, i, j int) {
	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]
		if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[0]) {
			continue
		}
		if grid[ni][nj] == 0 || (grid[ni][nj] > 0 && grid[i][j]+1 < grid[ni][nj]) {
			grid[ni][nj] = grid[i][j] + 1
			fillGrid(grid, ni, nj)
		}
	}
}

func resetFill(grid [][]int) {
	for _, line := range grid {
		for j, v := range line {
			if v != -1 {
				line[j] = 0
			}
		}
	}
}
