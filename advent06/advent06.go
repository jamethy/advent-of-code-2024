package advent06

import (
	"advent2024/util"
	"fmt"
)

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)
	grid := make([][]bool, len(lines))

	gi, gj := 0, 0
	gdi, gdj := 0, 0
	for i, line := range lines {
		grid[i] = make([]bool, len(line))
		for j, r := range line {
			switch r {
			case '#':
				grid[i][j] = true
			case '^':
				gi, gj = i, j
				gdi, gdj = -1, 0
			case '>':
				gi, gj = i, j
				gdi, gdj = 0, 1
			case '<':
				gi, gj = i, j
				gdi, gdj = 0, -1
			case 'v':
				gi, gj = i, j
				gdi, gdj = 1, 0
			}
		}
	}

	visited := make(map[int]map[int]struct{})
	for {
		i, j := gi+gdi, gj+gdj
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
			break
		}
		if !grid[i][j] {
			gi, gj = i, j
			if v, ok := visited[i]; !ok {
				visited[i] = map[int]struct{}{j: {}}
			} else {
				v[j] = struct{}{}
			}
		} else {
			gdi, gdj = gdj, -gdi
			fmt.Printf("turning right at %d, %d\n", gi, gj)
		}
	}

	visitCount := 0
	for _, v := range visited {
		visitCount += len(v)
	}

	return visitCount, 0
}
