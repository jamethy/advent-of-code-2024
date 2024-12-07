package advent07

import (
	"advent2024/util"
	"strconv"
	"strings"
)

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)

	var part1Sum uint64
	var part2Sum uint64
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		testValue, _ := strconv.ParseUint(parts[0], 10, 64)
		values := make([]uint64, 0, 10)
		for _, s := range strings.Split(parts[1], " ") {
			value, err := strconv.ParseUint(s, 10, 64)
			if err == nil {
				values = append(values, value)
			}
		}

		if TestValues(testValue, values, false) {
			part1Sum += testValue
			part2Sum += testValue
		} else if TestValues(testValue, values, true) {
			part2Sum += testValue
		}

	}
	return part1Sum, part2Sum
}

func TestValues(testValue uint64, values []uint64, allowConcat bool) bool {
	if len(values) == 1 {
		return values[0] == testValue
	}
	//if values[0] >= testValue {
	//	return false
	//}

	if TestValues(testValue, append([]uint64{values[0] * values[1]}, values[2:]...), allowConcat) {
		return true
	}
	if TestValues(testValue, append([]uint64{values[0] + values[1]}, values[2:]...), allowConcat) {
		return true
	}
	if allowConcat {
		newNum, _ := strconv.ParseUint(strconv.FormatUint(values[0], 10)+strconv.FormatUint(values[1], 10), 10, 64)
		if TestValues(testValue, append([]uint64{newNum}, values[2:]...), allowConcat) {
			return true
		}
	}
	return false
}
