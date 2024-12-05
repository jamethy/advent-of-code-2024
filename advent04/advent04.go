package advent04

import (
	"advent2024/util"
	"regexp"
)

var xmasRgx = regexp.MustCompile("(XMAS|SAMX)")
var masRgx = regexp.MustCompile("(MAS|SAM)")

func Solution(inputFile string) (part1, part2 any) {
	grid := getGrid(inputFile)

	part1Total := 0
	part2Total := 0
	for _, line := range grid {
		part1Total += countInLine(line)
	}

	for i := range grid {
		for j := range grid[0] {
			if i < len(grid)-3 && j < len(grid[0])-3 {
				word := []byte{grid[i][j], grid[i+1][j+1], grid[i+2][j+2], grid[i+3][j+3]}
				if xmasRgx.Match(word) {
					part1Total++
				}
			}
			if i < len(grid)-3 && j >= 3 {
				word := []byte{grid[i][j], grid[i+1][j-1], grid[i+2][j-2], grid[i+3][j-3]}
				if xmasRgx.Match(word) {
					part1Total++
				}
			}

			if grid[i][j] == 'A' && i > 0 && j > 0 && i < len(grid)-1 && j < len(grid[0])-1 {
				downRight := masRgx.Match([]byte{grid[i-1][j-1], grid[i][j], grid[i+1][j+1]})
				upRight := masRgx.Match([]byte{grid[i+1][j-1], grid[i][j], grid[i-1][j+1]})
				if downRight && upRight {
					part2Total++
				}
			}
		}
	}

	grid = rotateGrid90(grid)
	for _, line := range grid {
		part1Total += countInLine(line)
	}

	return part1Total, part2Total
}

func getGrid(inputFile string) [][]byte {
	lines := util.ReadFile(inputFile)
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}
	return grid
}

func rotateGrid90(grid [][]byte) [][]byte {
	rotated := make([][]byte, len(grid[0]))
	for j := range grid[0] {
		rotated[j] = make([]byte, len(grid[0]))
		for i := range grid {
			rotated[j][i] = grid[i][j]
		}
	}
	return rotated
}
func countInLine(line []byte) int {

	total := 0
	for {
		res := xmasRgx.FindIndex(line)
		if len(res) == 0 {
			break
		}
		total++
		line = line[res[1]-1:]
	}
	return total
}
