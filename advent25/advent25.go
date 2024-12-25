package advent25

import (
	"advent2024/util"
	"strings"
)

func Solution(inputFile string) (part1, part2 any) {
	keyLockStrings := util.ReadFileSplitBy(inputFile, "\n\n")
	keyLockHeight := len(strings.Split(keyLockStrings[0], "\n"))

	keys := make([][]int, 0, len(keyLockStrings))
	locks := make([][]int, 0, len(keyLockStrings))

	for _, str := range keyLockStrings {
		lines := strings.Split(str, "\n")
		heights := make([]int, len(lines[0]))
		for _, line := range lines {
			for i, c := range line {
				if c == '#' {
					heights[i] += 1
				}
			}
		}
		if lines[0] == "#####" {
			locks = append(locks, heights)
		} else {
			keys = append(keys, heights)
		}
	}

	part1Count := 0
	for _, key := range keys {
		for _, lock := range locks {
			fits := true
			for i, k := range key {
				if k+lock[i] > keyLockHeight {
					fits = false
					break
				}
			}
			if fits {
				part1Count++
			}
		}
	}

	return part1Count, 0
}
