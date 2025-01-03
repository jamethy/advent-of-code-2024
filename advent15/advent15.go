package advent15

import (
	"advent2024/util"
	"bytes"
	"fmt"
	"strings"
)

func Solution(inputFile string) (part1, part2 any) {
	part1 = part1Solution(inputFile)
	part2 = part2Solution(inputFile)
	return part1, part2
}

func part1Solution(inputFile string) int {
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
		rI, rJ = part1Move(grid, rI, rJ, dir)
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

	return part1Sum
}

func part1Move(grid [][]byte, i, j int, dir rune) (int, int) {

	di, dj := dirFromRune(dir)
	ni, nj := i+di, j+dj

	switch grid[ni][nj] {
	case '#':
		return i, j
	case '.':
		grid[ni][nj] = grid[i][j]
		grid[i][j] = '.'
		return ni, nj
	case 'O':
		part1Move(grid, ni, nj, dir)
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

func part2Solution(inputFile string) int {
	parts := util.ReadFileSplitBy(inputFile, "\n\n")

	originalGrid := bytes.Split([]byte(parts[0]), []byte("\n"))
	directions := strings.ReplaceAll(parts[1], "\n", "")

	grid := make([][]byte, len(originalGrid))

	rI, rJ := 0, 0
	for i, line := range originalGrid {
		grid[i] = make([]byte, 2*len(originalGrid[0]))
		for j, c := range line {
			nj := 2 * j

			switch c {
			case '@':
				rI, rJ = i, nj
				grid[i][nj] = '@'
				grid[i][nj+1] = '.'
			case '.':
				grid[i][nj] = '.'
				grid[i][nj+1] = '.'
			case 'O':
				grid[i][nj] = '['
				grid[i][nj+1] = ']'
			case '#':
				grid[i][nj] = '#'
				grid[i][nj+1] = '#'
			}
		}
	}
	for _, line := range grid {
		fmt.Println(string(line))
	}

	for _, dir := range directions {
		//if idx == 1092 {
		//	for _, line := range grid {
		//		fmt.Println(string(line))
		//	}
		//}
		fn := makePart2MoveFunc(grid, rI, rJ, dir, false)
		if fn != nil {
			fn()
			di, dj := dirFromRune(dir)
			rI, rJ = rI+di, rJ+dj
			//validate(grid, idx)
		}

		//for _, line := range grid {
		//	fmt.Println(string(line))
		//}

		//fmt.Println("\nMove " + string(dir) + ":")
		//for _, line := range grid {
		//	fmt.Println(string(line))
		//}
		//time.Sleep(10 * time.Second)
	}

	//for _, line := range grid {
	//	fmt.Println(string(line))
	//}

	part2Sum := 0
	for i, line := range grid {
		for j, c := range line {
			if c != '[' {
				continue
			}
			part2Sum += 100*i + j
		}
	}

	return part2Sum
}

func validate(grid [][]byte, idx int) {
	for _, line := range grid {
		if bytes.Contains(line, []byte("[.")) {
			panic(fmt.Sprintf("panic %d", idx))
		}
		if bytes.Contains(line, []byte("[@")) {
			panic(fmt.Sprintf("panic %d", idx))
		}
	}
}

func dirFromRune(r rune) (int, int) {
	switch r {
	case '<':
		return 0, -1
	case '>':
		return 0, 1
	case '^':
		return -1, 0
	case 'v':
		return 1, 0
	default:
		panic("unrecognized direction")
	}
}

func makePart2MoveFunc(grid [][]byte, i, j int, dir rune, skipPushRight bool) func() {
	if grid[i][j] == '#' {
		return nil
	}
	if grid[i][j] == '.' {
		return func() {}
	}

	di, dj := dirFromRune(dir)
	ni, nj := i+di, j+dj

	if grid[i][j] == '@' {
		prev := makePart2MoveFunc(grid, ni, nj, dir, false)
		if prev == nil {
			return nil
		}
		return func() {
			prev()
			grid[ni][nj] = '@'
			grid[i][j] = '.'
		}
	}
	if grid[i][j] == ']' {
		j--
		nj--
	}

	switch dir {
	case '<':
		prev := makePart2MoveFunc(grid, ni, nj, dir, false)
		if prev == nil {
			return nil
		}
		return func() {
			prev()
			grid[i][nj] = '['
			grid[i][j] = ']'
			grid[i][j+1] = '.'
		}
	case '>':
		prev := makePart2MoveFunc(grid, ni, nj+1, dir, false)
		if prev == nil {
			return nil
		}
		return func() {
			prev()
			grid[i][j] = '.'
			grid[i][nj] = '['
			grid[i][nj+1] = ']'
		}
	case '^', 'v':

		rightPrev := makePart2MoveFunc(grid, ni, nj+1, dir, false)
		if rightPrev == nil {
			return nil
		}

		shouldSkipPushRight := grid[ni][nj] == ']' && grid[ni][nj+1] == '['
		leftPrev := makePart2MoveFunc(grid, ni, nj, dir, shouldSkipPushRight)
		if leftPrev == nil {
			return nil
		}

		return func() {
			if !skipPushRight {
				rightPrev()
			}
			leftPrev()
			grid[ni][nj] = '['
			grid[ni][nj+1] = ']'
			grid[i][j] = '.'
			grid[i][j+1] = '.'
		}
	default:
		panic("unexpected direction")
	}
}
