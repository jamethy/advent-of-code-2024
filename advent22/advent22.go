package advent22

import (
	"advent2024/util"
	"advent2024/util/bitutil"
	"strconv"
)

func calculateNextSecret(s uint) uint {
	a := 64 * s
	s = bitutil.XOR(a, s)
	s = s % 16777216

	// just this gives the same number 2k, without also gives the same number
	a = s / 32
	s = bitutil.XOR(a, s)
	s = s % 16777216

	a = 2048 * s
	s = bitutil.XOR(a, s)
	s = s % 16777216

	return s
}

func calculateNthSecret(s uint, n int) uint {
	for i := 0; i < n; i++ {
		s = calculateNextSecret(s)
	}
	return s
}

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)

	part1Sum := uint(0)
	for _, line := range lines {
		sInt, _ := strconv.Atoi(line)
		s := calculateNthSecret(uint(sInt), 2000)
		part1Sum += s
	}
	return part1Sum, 0
}
