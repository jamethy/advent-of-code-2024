package advent03

import (
	"advent2024/util"
	"regexp"
	"strconv"
	"strings"
)

func Solution(inputFile string) (part1, part2 any) {
	data := util.ReadFileAsString(inputFile)
	data = strings.ReplaceAll(data, "\n", "")
	data = strings.TrimSpace(data)

	multiRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := multiRegex.FindAllStringSubmatch(data, -1)

	part1Total := 0
	for _, m := range matches {
		left, _ := strconv.Atoi(m[1])
		right, _ := strconv.Atoi(m[2])
		part1Total += left * right
	}

	return part1Total, 0
}
