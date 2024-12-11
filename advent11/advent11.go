package advent11

import (
	"advent2024/util"
	"fmt"
	"strconv"
	"strings"
)

type Rule func(string string) (string, bool)

func Solution(inputFile string) (part1, part2 any) {
	numbers := util.ReadFileAsString(inputFile)

	rules := []Rule{
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

	for c := 0; c < 25; c++ {
		split := strings.Split(numbers, " ")
		for i, s := range split {
			for _, rule := range rules {
				if updated, applied := rule(s); applied {
					split[i] = updated
					break
				}
			}
		}
		numbers = strings.Join(split, " ")
		if c < 6 {
			fmt.Println(numbers)
		}
	}

	return len(strings.Split(numbers, " ")), 0
}
