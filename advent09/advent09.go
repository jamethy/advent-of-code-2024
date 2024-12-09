package advent09

import (
	"advent2024/util"
	"strconv"
)

const empty = -1

func Solution(inputFile string) (part1, part2 any) {
	str := util.ReadFileAsString(inputFile)

	disk := make([]int, 0, len(str)*10)
	for i, r := range str {
		v, _ := strconv.Atoi(string(r))
		blockValue := empty
		if i%2 == 0 {
			blockValue = i / 2
		}

		block := make([]int, v)
		for j := range block {
			block[j] = blockValue
		}

		disk = append(disk, block...)
	}

	lastFree := 0
	for e := len(disk) - 1; e > 0 && e > lastFree; e-- {
		if disk[e] == empty {
			continue
		}
		for lastFree < e {
			lastFree++
			if disk[lastFree] == empty {
				disk[lastFree] = disk[e]
				disk[e] = empty
				break
			}
		}
	}

	return checksum(disk), 0
}

func checksum(disk []int) int {
	sum := 0
	for i, v := range disk {
		if v == empty {
			break
		}
		sum += i * v
	}
	return sum
}
