package advent05

import (
	"advent2024/util"
	"advent2024/util/set"
	"fmt"
	"sort"
	"strings"
)

func Solution(inputFile string) (part1, part2 any) {
	parts := util.ReadFileSplitBy(inputFile, "\n\n")
	ruleLines := strings.Split(parts[0], "\n")
	updateLines := strings.Split(parts[1], "\n")

	// rule: left should be before right
	// rules: for each right, all numbers that should be before
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

	part1MedianSum := 0
	part2MedianSum := 0
	for _, line := range updateLines {
		nums := util.StringsToInts(strings.Split(line, ","))
		if isRightOrder(nums, rules) {
			part1MedianSum += nums[len(nums)/2]
			continue
		}
		sort.Slice(nums, func(i, j int) bool {
			ni, nj := nums[i], nums[j]
			if s, ok := rules[ni]; ok && s.Has(nj) {
				return false
			}
			if s, ok := rules[nj]; ok && s.Has(ni) {
				return true
			}
			fmt.Println("should not see this. fingers crossed.")
			return false
		})
		part2MedianSum += nums[len(nums)/2]
	}

	return part1MedianSum, part2MedianSum
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
