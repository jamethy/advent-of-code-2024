package advent20

import (
	"advent2024/util"
	"advent2024/util/mathutil"
	"fmt"
	"strings"
)

func Solution(inputFile string) (part1, part2 any) {
	racetrack := util.ReadFileAsByteGrid(inputFile)

	minimumSavedPart1, minimumSavedPart2 := 100, 100
	if inputFile == "sample.txt" {
		minimumSavedPart1, minimumSavedPart2 = 1, 50
	}

	paths := make([][]int, len(racetrack))
	for i := range paths {
		paths[i] = make([]int, len(racetrack[0]))
	}

	si, sj := 0, 0
sSearch:
	for i, line := range racetrack {
		for j, c := range line {
			switch c {
			case 'S':
				si, sj = i, j
				break sSearch
			}
		}
	}

	fillPaths(paths, racetrack, si, sj)
	paths[si][sj] = 0

	strGrid := util.IntGridToStringGrid(paths)
	l := len(strGrid[0][0])
	for i, line := range racetrack {
		for j, c := range line {
			if c != '.' {
				strGrid[i][j] = strings.Repeat(string(c), l)
			}
		}
	}

	for _, line := range strGrid {
		fmt.Println(line)
	}

	part1 = countCheats(racetrack, paths, 2, minimumSavedPart1)
	part2 = countCheats(racetrack, paths, 20, minimumSavedPart2)
	return part1, part2
}

func countCheats(racetrack [][]byte, paths [][]int, maxSkip int, minimumSaved int) int {
	cheatCount := 0
	for i, line := range racetrack {
		for j, c := range line {
			if c == '#' {
				continue
			}
			for di := -maxSkip; di <= maxSkip; di++ {
				for dj := -maxSkip; dj <= maxSkip; dj++ {
					dist := mathutil.AbsInt(di) + mathutil.AbsInt(dj)
					if dist > maxSkip {
						continue
					}
					ni, nj := i+di, j+dj
					if ni < 0 || nj < 0 || ni >= len(paths) || nj >= len(paths[0]) {
						continue
					}
					if racetrack[ni][nj] == '#' {
						continue
					}
					timeSaved := paths[ni][nj] - paths[i][j] - dist
					if timeSaved >= minimumSaved {
						cheatCount++
					}
				}
			}
		}
	}

	return cheatCount
}

var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func fillPaths(paths [][]int, racetrack [][]byte, i, j int) {
	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]
		if ni < 0 || ni >= len(paths) || nj < 0 || nj >= len(paths[0]) {
			continue
		}
		if racetrack[ni][nj] == '#' {
			continue
		}
		nextPostValue := paths[ni][nj]
		if nextPostValue == 0 || paths[i][j]+1 < nextPostValue {
			paths[ni][nj] = paths[i][j] + 1
			fillPaths(paths, racetrack, ni, nj)
		}
	}
}
