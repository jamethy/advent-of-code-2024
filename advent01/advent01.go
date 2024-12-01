package advent01

import (
	"advent2024/util"
	"advent2024/util/mathutil"
	"slices"
	"strconv"
	"strings"
)

func Solution(inputFile string) (part1, part2 interface{}) {
	lines := util.ReadFile(inputFile)

	var lefts []int
	var rights []int

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, "   ")
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])
		lefts = append(lefts, left)
		rights = append(rights, right)
	}

	slices.Sort(lefts)
	slices.Sort(rights)

	totalDiff := 0
	for i := range lefts {
		totalDiff += mathutil.AbsInt(lefts[i] - rights[i])
	}
	return totalDiff, nil
}
