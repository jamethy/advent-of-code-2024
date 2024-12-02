package advent02

import (
	"advent2024/util"
	"slices"
	"strings"
)

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)
	var reports [][]int
	for _, line := range lines {
		report := util.StringsToInts(strings.Split(line, " "))
		reports = append(reports, report)
	}

	part1SafeCount := 0
	part2SafeCount := 0
	for _, r := range reports {
		part1IsSafe, part2IsSafe := reportIsSafe(r)

		if part1IsSafe {
			part1SafeCount++
		}
		if part2IsSafe {
			part2SafeCount++
		}
		//if part2IsSafe != part2BruteForce(r) {
		//	fmt.Println("mismatch ", part2IsSafe, r)
		//}
	}

	return part1SafeCount, part2SafeCount
}

func part2BruteForce(r []int) bool {
	for i := range r {
		sub := slices.Concat(r[0:i], r[i+1:])
		safe, _ := reportIsSafe(sub)
		if safe {
			return true
		}
	}
	return false
}

func reportIsSafe(r []int) (bool, bool) {
	diffs := make([]int, len(r)-1)
	for i := range r {
		if i == 0 {
			continue
		}
		diffs[i-1] = r[i] - r[i-1]
	}
	diffs = positiveDiffs(diffs)

	part1UnsafeIdx := -1
	part1IsSafe := true
	part2IsSafe := true

	for i, d := range diffs {
		if d > 0 && d <= 3 {
			continue
		}
		if !part1IsSafe {
			if part1UnsafeIdx == i-1 {
				newDiff := d + diffs[i-1]
				if newDiff <= 0 || newDiff > 3 {
					part2IsSafe = false
					break
				}
			} else {
				part2IsSafe = false
				break
			}
		}
		part1IsSafe = false
		part1UnsafeIdx = i
		if i == 0 || i == len(diffs)-1 {
			continue
		}
		newDiff := d + diffs[i-1]
		if !(newDiff <= 0 || newDiff > 3) {
			continue
		}
		newDiff = d + diffs[i+1]
		if !(newDiff <= 0 || newDiff > 3) {
			continue
		}
		part2IsSafe = false
		break
	}

	return part1IsSafe, part2IsSafe
}

func positiveDiffs(diffs []int) []int {
	negativeCount := 0
	for _, d := range diffs {
		if d < 0 {
			negativeCount++
		}
	}
	if negativeCount > len(diffs)/2 {
		for i := range diffs {
			diffs[i] *= -1
		}
	}
	return diffs
}
