package advent05

import (
	"advent2024/util"
	"advent2024/util/set"
	"strings"
)

func Solution(inputFile string) (part1, part2 any) {
	parts := util.ReadFileSplitBy(inputFile, "\n\n")
	ruleLines := strings.Split(parts[0], "\n")
	updateLines := strings.Split(parts[1], "\n")

	rules := make(map[int]set.Set[int])
	for _, line := range ruleLines {
		nums := util.StringsToInts(strings.Split(line, "|"))
		s, ok := rules[nums[1]]
		if !ok {
			s = set.NewSet[int]()
		}
		s.Add(nums[0])
		rules[nums[1]] = s
	}

	medianSum := 0
	for _, line := range updateLines {
		nums := util.StringsToInts(strings.Split(line, ","))
		if !isRightOrder(nums, rules) {
			continue
		}
		medianSum += nums[len(nums)/2]
	}

	return medianSum, 0
}

func isRightOrder(nums []int, rules map[int]set.Set[int]) bool {
	for i := 0; i < len(nums)-1; i++ {
		if s, ok := rules[nums[i]]; ok {
			for j := i + 1; j < len(nums); j++ {
				if s.Has(nums[j]) {
					return false
				}
			}
		}
	}
	return true
}
