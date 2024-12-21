package advent21

import (
	"advent2024/util"
	"advent2024/util/mathutil"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// +---+---+---+
// | 7 | 8 | 9 |
// +---+---+---+
// | 4 | 5 | 6 |
// +---+---+---+
// | 1 | 2 | 3 |
// +---+---+---+
//     | 0 | A |
//     +---+---+

func numpadMoves(a, b rune, transformer func(string) string) string {
	aN, _ := strconv.Atoi(string(a))
	bN, _ := strconv.Atoi(string(b))

	moveString := string(a) + " to " + string(b)
	_ = moveString

	aColumn := (aN + 2) % 3
	if a == '0' {
		aColumn = 1
	}
	bColumn := (bN + 2) % 3
	if b == '0' {
		bColumn = 1
	}

	var leftRights string
	if aColumn < bColumn {
		leftRights = strings.Repeat(">", mathutil.AbsInt(bColumn-aColumn))
	} else if aColumn > bColumn {
		leftRights = strings.Repeat("<", mathutil.AbsInt(aColumn-bColumn))
	}

	aLevel := (aN + 2) / 3
	bLevel := (bN + 2) / 3

	var upDowns string
	if aLevel < bLevel {
		upDowns = strings.Repeat("^", mathutil.AbsInt(bLevel-aLevel))
	} else if aLevel > bLevel {
		upDowns = strings.Repeat("v", mathutil.AbsInt(aLevel-bLevel))
	}

	if len(upDowns) == 0 || len(leftRights) == 0 {
		// doesn't matter which is better, they are the same
		t := transformer(upDowns + leftRights)
		return t
	}

	if aLevel == 0 && bColumn == 0 {
		// avoid going left into gap
		t := transformer(upDowns + leftRights)
		return t
	}
	if aColumn == 0 && bLevel == 0 {
		// avoid going down into gap
		t := transformer(leftRights + upDowns)
		return t
	}

	upDownsFirst := transformer(upDowns + leftRights)
	leftRightsFirst := transformer(leftRights + upDowns)
	if len(leftRightsFirst) <= len(upDownsFirst) {
		return leftRightsFirst
	}
	return upDownsFirst
}

//     +---+---+
//     | ^ | A |
// +---+---+---+
// | < | v | > |
// +---+---+---+

func getArrowKeysLevelAndColumn(a rune) (int, int) {
	level, column := 0, 0

	switch a {
	case '^', 'A':
		level = 1
	default:
		level = 0
	}

	switch a {
	case 'A', '>':
		column = 2
	case '^', 'v':
		column = 1
	default:
		column = 0
	}

	return level, column
}

func transformTwoArrowKeys(a, b rune, transformer func(string) string) string {
	aLevel, aColumn := getArrowKeysLevelAndColumn(a)
	bLevel, bColumn := getArrowKeysLevelAndColumn(b)

	moveString := string(a) + " to " + string(b)
	_ = moveString

	var leftRights string
	if aColumn < bColumn {
		leftRights = strings.Repeat(">", mathutil.AbsInt(bColumn-aColumn))
	} else if aColumn > bColumn {
		leftRights = strings.Repeat("<", mathutil.AbsInt(aColumn-bColumn))
	}

	var upDowns string
	if aLevel < bLevel {
		upDowns = strings.Repeat("^", mathutil.AbsInt(bLevel-aLevel))
	} else if aLevel > bLevel {
		upDowns = strings.Repeat("v", mathutil.AbsInt(aLevel-bLevel))
	}

	if len(upDowns) == 0 || len(leftRights) == 0 {
		// doesn't matter, they are the same
		return transformer(upDowns + leftRights)
	}

	if aLevel == 1 && bColumn == 0 {
		// avoid going left into gap
		return transformer(upDowns + leftRights)
	}
	if aColumn == 0 {
		// avoid going up into gap
		return transformer(leftRights + upDowns)
	}

	upDownsFirst := transformer(upDowns + leftRights)
	leftRightsFirst := transformer(leftRights + upDowns)
	if len(leftRightsFirst) <= len(upDownsFirst) {
		return leftRightsFirst
	}
	return upDownsFirst
}

// <A^A>^^AvvvA
// A< -> v<<A

// v<

var cache = make(map[string]string)

func transformOnce(input string, depth int) string {
	input += "A"
	if depth == 0 {
		return input
	}

	pos := 'A'
	str := ""
	for _, newPos := range input {
		c := transformTwoArrowKeys(pos, newPos, func(s string) string {
			return transformOnce(s, depth-1)
		})
		str += c
		pos = newPos
	}

	return str
}

func sequenceLength(line string, depth int) int {
	seq := ""
	numPadPos := 'A'
	for _, newPos := range line {

		moveString := string(numPadPos) + " to " + string(newPos)
		_ = moveString

		mv := numpadMoves(numPadPos, newPos, func(s string) string {
			//return transformArrowKeys(s+"A", 'A', depth)
			return transformOnce(s, depth)
		})
		numPadPos = newPos
		seq += mv
	}
	return len(seq)
}

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)

	part1Complexity, part2Complexity := 0, 0
	for _, line := range lines {
		depthOf2 := sequenceLength(line, 2)
		//depthOf25 := sequenceLength(line, 25)

		num, _ := strconv.Atoi(line[:len(line)-1])
		fmt.Printf("%s: %d * %d\n", line, depthOf2, num)
		part1Complexity += depthOf2 * num
		//part2Complexity += depthOf25 * num
	}

	return part1Complexity, part2Complexity
}

func reverseString(str string) string {
	b := []byte(str)
	slices.Reverse(b)
	return string(b)
}

func stringIdentity(str string) string {
	return str
}
