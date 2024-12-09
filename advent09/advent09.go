package advent09

import (
	"advent2024/util"
	"slices"
	"strconv"
)

const empty = -1

func Solution(inputFile string) (part1, part2 any) {
	str := util.ReadFileAsString(inputFile)

	disk1 := strToDisk(str)

	lastFree := 0
	for e := len(disk1) - 1; e > 0 && e > lastFree; e-- {
		if disk1[e] == empty {
			continue
		}
		for lastFree < e {
			lastFree++
			if disk1[lastFree] == empty {
				disk1[lastFree] = disk1[e]
				disk1[e] = empty
				break
			}
		}
	}

	disk2 := make([]Block, len(str))
	for i, v := range util.ParseIntList(str, "") {
		blockValue := empty
		if i%2 == 0 {
			blockValue = i / 2
		}
		disk2[i] = Block{
			Value: blockValue,
			Size:  v,
		}
	}

	for e := len(disk2) - 1; e > 0; e-- {
		block := disk2[e]
		if block.Value == empty {
			continue
		}
		emptySpace := 1
		for emptySpace < e {
			emptyBlock := disk2[emptySpace]
			if emptyBlock.Value != empty {
				emptySpace++
				continue
			}
			if emptyBlock.Size >= block.Size {
				newBlocks := []Block{
					block,
				}
				if emptyBlock.Size-block.Size != 0 {
					newBlocks = []Block{
						block,
						{
							Value: empty,
							Size:  emptyBlock.Size - block.Size,
						},
					}
				}

				disk2 = slices.Concat(
					disk2[0:emptySpace],
					newBlocks,
					disk2[emptySpace+1:e],
					[]Block{
						{
							Value: empty,
							Size:  block.Size,
						},
					},
					disk2[e+1:],
				)
				break
			}
			emptySpace++
		}
	}
	checksumableDisk2 := blocksToChecksumable(disk2)

	return checksum(disk1), checksum(checksumableDisk2)
}

type Block struct {
	Value int
	Size  int
}

func blocksToChecksumable(blocks []Block) []int {
	checksumableDisk := make([]int, 0, len(blocks)*10)
	for _, b := range blocks {
		for range b.Size {
			checksumableDisk = append(checksumableDisk, b.Value)
		}
	}
	return checksumableDisk
}

func strToDisk(str string) []int {
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
	return disk
}

func checksum(disk []int) int {
	sum := 0
	for i, v := range disk {
		if v != empty {
			sum += i * v
		}
	}
	return sum
}
