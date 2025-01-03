package advent13

import (
	"advent2024/util"
	"advent2024/util/mathutil"
	"math"
	"strconv"
	"strings"
)

const costOfPressA = 3
const costOfPressB = 1
const part2Offset = 10000000000000

type Machine struct {
	AX float64
	AY float64
	BX float64
	BY float64
	PX float64
	PY float64
}

func calcTokensUsed(m Machine) int {
	c1 := m.PY - m.BY/m.BX*m.PX
	d1 := m.AY/m.AX - m.BY/m.BX
	a1 := c1 / (d1 * m.AX)

	c2 := m.PY - m.AY/m.AX*m.PX
	d2 := -d1
	b1 := c2 / (d2 * m.BX)

	if !essentiallyInteger(a1) {
		return 0
	}

	if !essentiallyInteger(b1) {
		return 0
	}

	return costOfPressA*mathutil.Round(a1) + costOfPressB*mathutil.Round(b1)
}

func Solution(inputFile string) (part1, part2 any) {
	machineStrings := util.ReadFileSplitBy(inputFile, "\n\n")

	part1TokenUsed := 0
	part2TokenUsed := 0

	for _, machineString := range machineStrings {
		m := parseMachine(machineString)
		part1TokenUsed += calcTokensUsed(m)
		m.PX += part2Offset
		m.PY += part2Offset
		part2TokenUsed += calcTokensUsed(m)
	}

	return part1TokenUsed, part2TokenUsed
}

func essentiallyInteger(f float64) bool {
	return math.Abs(f-math.Round(f)) < 0.001
}

func parseMachine(machineString string) Machine {
	parts := strings.Split(machineString, "\n")
	m := Machine{}

	buttonParts := strings.Split(parts[0], ", Y+")
	m.AX, _ = strconv.ParseFloat(buttonParts[0][12:], 64)
	m.AY, _ = strconv.ParseFloat(buttonParts[1], 64)

	buttonParts = strings.Split(parts[1], ", Y+")
	m.BX, _ = strconv.ParseFloat(buttonParts[0][12:], 64)
	m.BY, _ = strconv.ParseFloat(buttonParts[1], 64)

	prizeParts := strings.Split(parts[2], ", Y=")
	m.PX, _ = strconv.ParseFloat(prizeParts[0][9:], 64)
	m.PY, _ = strconv.ParseFloat(prizeParts[1], 64)

	return m
}
