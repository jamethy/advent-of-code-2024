package advent16

import (
	"advent2024/util"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

var directions = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func Solution(inputFile string) (part1, part2 any) {
	grid := util.ReadFileAsByteGrid(inputFile)

	scoreGrid := make([][]int, len(grid))
	for i := range grid {
		scoreGrid[i] = make([]int, len(grid[i]))
	}

	i, j := findTile(grid, 'S')
	scoreGrid[i][j] = 1 // so I don't have to initialize entire grid
	stepCount := 0
	move(grid, scoreGrid, i, j, 0, stepCount)

	eI, eJ := findTile(grid, 'E')
	part1 = scoreGrid[eI][eJ] - 1

	//printScoreGrid(grid, scoreGrid)

	return part1, 0
}

func move(grid [][]byte, scoreGrid [][]int, i, j, dirIdx int, stepCount int) {
	for dirChange := 0; dirChange < 4; dirChange++ {
		nDirIdx := (dirIdx + dirChange) % 4
		dir := directions[nDirIdx]

		ni, nj := i+dir[0], j+dir[1]
		if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[0]) {
			continue
		}
		if grid[ni][nj] == '#' {
			continue
		}

		// todo figure out clever way
		var rotationCost int
		switch dirChange {
		case 0:
			rotationCost = 0
		case 1, 3:
			rotationCost = 1000
		case 2:
			rotationCost = 2000
		}

		potentialScore := scoreGrid[i][j] + rotationCost + 1
		if scoreGrid[ni][nj] == 0 || scoreGrid[ni][nj] >= potentialScore {
			scoreGrid[ni][nj] = potentialScore
			move(grid, scoreGrid, ni, nj, nDirIdx, stepCount+1)
		}
	}
}

func findTile(grid [][]byte, c byte) (int, int) {
	for i, line := range grid {
		for j, v := range line {
			if v == c {
				return i, j
			}
		}
	}
	return -1, -1
}

func printScoreGrid(grid [][]byte, scoreGrid [][]int) {
	sb := strings.Builder{}
	for i, line := range grid {
		for j, v := range line {
			if v == 'S' || v == 'E' {
				sb.Write(bytes.Repeat([]byte{v}, 4))
				sb.WriteRune(' ')
				continue
			}
			score := scoreGrid[i][j]
			s := strconv.Itoa(score)
			if score < 1000 {
				s = "0" + s
			}
			if score < 100 {
				s = "0" + s
			}
			if score < 10 {
				s = "0" + s
			}
			sb.WriteString(s + " ")
		}
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}
