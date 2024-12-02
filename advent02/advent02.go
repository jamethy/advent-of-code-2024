package advent02

import (
	"advent2024/util"
	"advent2024/util/mathutil"
	"strings"
)

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)
	var reports [][]int
	for _, line := range lines {
		report := util.StringsToInts(strings.Split(line, " "))
		reports = append(reports, report)
	}

	safeCount := 0
	for _, r := range reports {
		diffs := make([]int, len(r)-1)
		for i := range r {
			if i == 0 {
				continue
			}
			diffs[i-1] = r[i] - r[i-1]
		}

		isSafe := true
		s := mathutil.Sign(diffs[0])
		for _, d := range diffs {
			if s != mathutil.Sign(d) {
				isSafe = false
				break
			}
			delta := mathutil.AbsInt(d)
			if delta == 0 || delta > 3 {
				isSafe = false
				break
			}
		}
		if isSafe {
			safeCount++
		}
	}

	return safeCount, 0
}
