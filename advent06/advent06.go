package advent06

import (
	"advent2024/util"
	"advent2024/util/set"
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

	visited, _ := calculate(grid, gi, gj, gdi, gdj)
	part1 = len(visited)
	delete(visited, Point{i: gi, j: gj})

	part2Count := 0
	for point := range visited {
		i, j := point.i, point.j
		grid[i][j] = true
		_, isLoop := calculate(grid, gi, gj, gdi, gdj)
		if isLoop {
			fmt.Printf("Loop if placed at %d, %d", i, j)
			part2Count++
		}
		grid[i][j] = false
	}

	return part1, part2Count
}

type Point struct {
	i, j int
}

func calculate(grid [][]bool, gi, gj, gdi, gdj int) (map[Point]set.Set[Point], bool) {
	visited := make(map[Point]set.Set[Point], len(grid)*len(grid[0]))
	for {
		i, j := gi+gdi, gj+gdj
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
			break
		}
		if !grid[i][j] {
			gi, gj = i, j

			p := Point{i: i, j: j}
			v := visited[p]
			if v == nil {
				v = set.NewSet[Point]()
			}

			d := Point{i: gdi, j: gdj}
			if v.Has(d) {
				return visited, true
			}
			v.Add(d)
			visited[p] = v

		} else {
			gdi, gdj = gdj, -gdi
		}
	}

	return visited, false
}
