package advent22

import (
	"advent2024/util"
	"advent2024/util/bitutil"
	"advent2024/util/set"
	"fmt"
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
	fmt.Println("===================")
	fmt.Println(s)
	for i := 0; i < n; i++ {
		s = calculateNextSecret(s)
	}
	return s
}

func changeKey(a, b, c, d int32) int32 {
	// each is -9 to 9
	a += 9
	b += 9
	c += 9
	d += 9

	// each is 0 to 18 (5 bits)
	var k int32
	k |= a << 15
	k |= b << 10
	k |= c << 5
	k |= d
	return k
}

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)

	part1Sum := uint(0)
	prices := make([]int, 2001)
	changeMaps := make([]map[int32]byte, len(lines))
	allChangeKeys := set.NewSet[int32]()
	for lineNum, line := range lines {
		sInt, _ := strconv.Atoi(line)
		s := uint(sInt)

		changeMap := make(map[int32]byte, len(lines)*len(prices))
		changeMaps[lineNum] = changeMap

		prices[0] = int(s % 10)
		for i := 0; i < 2000; i++ {
			s = calculateNextSecret(s)
			prices[i+1] = int(s % 10)
		}

		part1Sum += s

		for i := 4; i < 2001; i++ {
			k := changeKey(
				int32(prices[i-3]-prices[i-4]),
				int32(prices[i-2]-prices[i-3]),
				int32(prices[i-1]-prices[i-2]),
				int32(prices[i]-prices[i-1]),
			)
			if _, ok := changeMap[k]; !ok {
				changeMap[k] = byte(prices[i])
				allChangeKeys.Add(k)
			}
		}
	}

	highestPrices := 0
	for k := range allChangeKeys {
		totalPrice := 0

		for _, changeMap := range changeMaps {
			totalPrice += int(changeMap[k])
		}

		if totalPrice > highestPrices {
			highestPrices = totalPrice
		}
	}

	return part1Sum, highestPrices
}
