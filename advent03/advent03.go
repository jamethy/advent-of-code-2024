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

	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)

	doIndexes := doRegex.FindAllStringIndex(data, -1)
	dontIndexes := dontRegex.FindAllStringIndex(data, -1)
	matchIndexes := multiRegex.FindAllStringIndex(data, -1)

	enabled := true
	part2Total := 0
	for i := 0; i < len(data); i++ {
		if len(doIndexes) > 0 && doIndexes[0][0] == i {
			enabled = true
			i = doIndexes[0][1] - 1
			doIndexes = doIndexes[1:]
			continue
		}
		if len(dontIndexes) > 0 && dontIndexes[0][0] == i {
			enabled = false
			i = dontIndexes[0][1] - 1
			dontIndexes = dontIndexes[1:]
			continue
		}

		if len(matchIndexes) > 0 && matchIndexes[0][0] == i {
			if enabled {
				m := multiRegex.FindStringSubmatch(data[i:matchIndexes[0][1]])
				left, _ := strconv.Atoi(m[1])
				right, _ := strconv.Atoi(m[2])
				part2Total += left * right
			}
			i = matchIndexes[0][1] - 1
			matchIndexes = matchIndexes[1:]
		}
	}

	return part1Total, part2Total
}
