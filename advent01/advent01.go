package advent01

import (
	"advent2024/util"
	"advent2024/util/mathutil"
)

func Solution(inputFile string) (part1, part2 interface{}) {
	elves := parseElves(inputFile)

	top := []int{0, 0, 0} // descending order

	for _, elf := range elves {
		v := mathutil.SumInts(elf)
		for i, t := range top {
			if v > t {
				// shift remaining elements
				for j := len(top) - 1; j > i; j-- {
					top[j] = top[j-1]
				}
				// replace element
				top[i] = v
				break
			}
		}
	}
	return top[0], mathutil.SumInts(top)
}

func parseElves(inputFile string) [][]int {
	elfLines := util.ReadFileSplitBy(inputFile, "\n\n")
	res := make([][]int, len(elfLines))
	for i, line := range elfLines {
		res[i] = util.ParseIntList(line, "\n")
	}
	return res
}
