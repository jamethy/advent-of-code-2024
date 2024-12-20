package advent20

import (
	"advent2024/util"
	"advent2024/util/mathutil"
	"fmt"
	"strings"
)

func Solution(inputFile string) (part1, part2 any) {
	racetrack := util.ReadFileAsByteGrid(inputFile)

	minimumSaved := 100
	if inputFile == "sample.txt" {
		minimumSaved = 0
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

	strGrid := util.IntGridToStringGrid(paths)
	l := len(strGrid[0][0])
	for i, line := range racetrack {
		for j, c := range line {
			if c != '.' {
				strGrid[i][j] = strings.Repeat(string(c), l)
			}
		}
	}

	shortcutCount := 0
	for i, line := range racetrack {
		for j, c := range line {
			if c != '#' {
				continue
			}
			if isSkippableLeftRight(i, j, racetrack) {
				left, right := paths[i][j-1], paths[i][j+1]
				timeSaved := mathutil.AbsInt(left-right) - 2
				if timeSaved >= minimumSaved {
					shortcutCount++
				}
			}
			if isSkippableUpDown(i, j, racetrack) {
				up, down := paths[i-1][j], paths[i+1][j]
				timeSaved := mathutil.AbsInt(up-down) - 2
				if timeSaved >= minimumSaved {
					shortcutCount++
				}
			}
		}
	}

	for _, line := range strGrid {
		fmt.Println(line)
	}

	return shortcutCount, 0
}

func isSkippableLeftRight(i, j int, racetrack [][]byte) bool {
	if j <= 1 || j >= len(racetrack[0])-2 {
		return false
	}
	return racetrack[i][j-1] != '#' && racetrack[i][j+1] != '#'
}

func isSkippableUpDown(i, j int, racetrack [][]byte) bool {
	if i <= 1 || i >= len(racetrack)-2 {
		return false
	}
	return racetrack[i-1][j] != '#' && racetrack[i+1][j] != '#'
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
