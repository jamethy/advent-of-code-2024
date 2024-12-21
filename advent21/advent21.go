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

	if aLevel == 0 {
		t := transformer(upDowns + leftRights)
		return t
	}
	if aColumn == 0 {
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

type ArrowKeys struct {
	cached map[string]string
}

//     +---+---+
//     | ^ | A |
// +---+---+---+
// | < | v | > |
// +---+---+---+

func (ak ArrowKeys) getLevelAndColumn(a rune) (int, int) {
	level, column := 0, 0
	if strings.ContainsRune("^A", a) {
		level = 1
	}
	if strings.ContainsRune("^v", a) {
		column = 1
	} else if strings.ContainsRune("A>", a) {
		column = 2
	}

	return level, column
}

func (ak ArrowKeys) transformTwo(a, b rune, transformer func(string) string) string {
	aLevel, aColumn := ak.getLevelAndColumn(a)
	bLevel, bColumn := ak.getLevelAndColumn(b)

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
		return transformer(upDowns + leftRights)
	}

	if aLevel == 1 {
		return transformer(upDowns + leftRights)
	}
	if aColumn == 0 {
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

func (ak ArrowKeys) transform(input string, pos rune, transformer func(string) string) string {
	if input == "" {
		return ""
	}

	//if c, ok := a.cached[input]; ok {
	//	return c
	//}
	//if c, ok := a.cached[reverseString(input)]; ok {
	//	return reverseString(c)
	//}

	c := ak.transformTwo(pos, rune(input[0]), transformer)
	c += "A"

	rest := ak.transform(input[1:], rune(input[0]), transformer)
	c += rest
	//ak.cached[input] = c

	//c := ""
	//
	//for i, k := range input {
	//	if i == 0 {
	//		continue
	//	}
	//
	//}
	//
	//a.cached[input] = c
	return c
}

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)

	ak := ArrowKeys{cached: make(map[string]string)}

	part1Complexity := 0
	for _, line := range lines {
		seq := ""
		numPadPos := 'A'
		transformer1Pos := 'A'
		//transformer2Pos := 'A'
		//line = strings.Join(strings.Split(line, ""), "A")
		for _, newPos := range line {

			moveString := string(numPadPos) + " to " + string(newPos)
			_ = moveString

			mv := numpadMoves(numPadPos, newPos, func(s string) string {
				return s
			})
			mv += "A"
			mv = ak.transform(mv, transformer1Pos, func(s string) string {
				return s
			})
			mv = ak.transform(mv, transformer1Pos, func(s string) string {
				return s
			})
			numPadPos = newPos
			seq += mv
		}

		//seq = ak.transform(seq, 'A')
		//seq = ak.transform(seq, 'A')

		num, _ := strconv.Atoi(line[:len(line)-1])
		fmt.Printf("%s: %d * %d\n", line, len(seq), num)
		part1Complexity += len(seq) * num
	}

	return part1Complexity, 0
}

func reverseString(str string) string {
	b := []byte(str)
	slices.Reverse(b)
	return string(b)
}

func stringIdentity(str string) string {
	return str
}
