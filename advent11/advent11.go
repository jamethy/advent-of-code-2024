package advent11

import (
	"advent2024/util"
	"strconv"
	"strings"
)

type Rule func(string string) (string, bool)

var rules = []Rule{
	func(str string) (string, bool) {
		if str != "0" {
			return "", false
		}
		return "1", true
	},
	func(str string) (string, bool) {
		if len(str)%2 != 0 {
			return "", false
		}
		left := str[0 : len(str)/2]
		right := strings.TrimLeft(str[len(str)/2:], "0")
		if right == "" {
			right = "0"
		}
		return left + " " + right, true
	},
	func(str string) (string, bool) {
		value, _ := strconv.ParseUint(str, 10, 64)
		return strconv.FormatUint(value*2024, 10), true
	},
}

type Processor struct {
	cache map[string][]int
}

func applyRules(stoneStr string) string {
	for _, rule := range rules {
		if updated, applied := rule(stoneStr); applied {
			return updated
		}
	}
	panic("shouldn't get here")
}

func (p *Processor) process(stoneStr string, steps int) int {
	cached, ok := p.cache[stoneStr]
	if !ok {
		cached = make([]int, 75)
		p.cache[stoneStr] = cached
	}
	if cached[steps-1] != 0 {
		return cached[steps-1]
	}

	updated := applyRules(stoneStr)
	parts := strings.Split(updated, " ")

	if steps == 1 {
		cached[steps-1] = len(parts)
		return len(parts)
	}
	furtherSum := 0
	for _, part := range parts {
		furtherSum += p.process(part, steps-1)
	}
	cached[steps-1] = furtherSum
	return furtherSum
}

func Solution(inputFile string) (part1, part2 any) {
	numbers := util.ReadFileAsString(inputFile)

	p := Processor{cache: make(map[string][]int)}

	part1Total := 0
	for _, n := range strings.Split(numbers, " ") {
		part1Total += p.process(n, 25)
	}

	part2Total := 0
	for _, n := range strings.Split(numbers, " ") {
		part2Total += p.process(n, 75)
	}

	return part1Total, part2Total
}
