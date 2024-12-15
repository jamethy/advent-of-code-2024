package advent15

import (
	"advent2024/util"
	"bytes"
	"strings"
)

func Solution(inputFile string) (part1, part2 any) {
	parts := util.ReadFileSplitBy(inputFile, "\n\n")

	grid := bytes.Split([]byte(parts[0]), []byte("\n"))
	directions := strings.ReplaceAll(parts[1], "\n", "")

	rI, rJ := 0, 0
robotSearch:
	for i, line := range grid {
		for j, c := range line {
			if c == '@' {
				rI, rJ = i, j
				break robotSearch
			}
		}
	}

	for _, dir := range directions {
		rI, rJ = move(grid, rI, rJ, dir)
	}

	part1Sum := 0
	for i, line := range grid {
		for j, c := range line {
			if c != 'O' {
				continue
			}
			part1Sum += 100*i + j
		}
	}

	//for _, line := range grid {
	//	fmt.Println(string(line))
	//}

	return part1Sum, 0
}

func move(grid [][]byte, i, j int, dir rune) (int, int) {

	ni, nj := 0, 0
	switch dir {
	case '<':
		ni, nj = i, j-1
	case '>':
		ni, nj = i, j+1
	case '^':
		ni, nj = i-1, j
	case 'v':
		ni, nj = i+1, j
	default:
		panic("unrecognized direction")
	}

	switch grid[ni][nj] {
	case '#':
		return i, j
	case '.':
		grid[ni][nj] = grid[i][j]
		grid[i][j] = '.'
		return ni, nj
	case 'O':
		move(grid, ni, nj, dir)
		if grid[ni][nj] == '.' {
			grid[ni][nj] = grid[i][j]
			grid[i][j] = '.'
			return ni, nj
		}
		return i, j
	default:
		panic("unrecognized character")
	}
}
