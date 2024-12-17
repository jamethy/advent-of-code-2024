package advent17

import (
	"advent2024/util"
	"advent2024/util/bitutil"
	"advent2024/util/mathutil"
	"strconv"
	"strings"
)

type Computer struct {
	A, B, C            uint
	InstructionPointer uint
	Program            []uint
	Output             []uint
}

func (c *Computer) Run() {
	for c.Next() {
		c.Step()
	}
}

func (c *Computer) Next() bool {
	return c.InstructionPointer < uint(len(c.Program))
}

func (c *Computer) Step() {

	opcode := c.Program[c.InstructionPointer]
	operand := c.Program[c.InstructionPointer+1]
	c.InstructionPointer += 2

	switch opcode {
	case 0: // adv
		c.A = c.advValue(operand)
	case 1: // bxl
		c.B = bitutil.XOR(c.B, operand)
	case 2: // bst
		c.B = c.ComboValue(operand) % 8
	case 3: // jnz
		if c.A != 0 {
			c.InstructionPointer = operand
		}
	case 4: // bxc
		c.B = bitutil.XOR(c.B, c.C)
	case 5: // out
		c.Output = append(c.Output, c.ComboValue(operand)%8)
	case 6: // bdv
		c.B = c.advValue(operand)
	case 7: // cdv
		c.C = c.advValue(operand)
	}
}

func (c *Computer) advValue(operand uint) uint {
	numerator := c.A
	denominator := uint(mathutil.IntPow(2, int(c.ComboValue(operand))))
	return numerator / denominator
}

func (c *Computer) ComboValue(v uint) uint {
	switch v {
	case 0, 1, 2, 3:
		return v
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	default:
		panic("invalid combo value")
	}
}

func Solution(inputFile string) (part1, part2 any) {

	c := parseComputer(inputFile)

	for c.Next() {
		c.Step()
	}

	sb := strings.Builder{}
	for i, n := range c.Output {
		if i != 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(strconv.FormatUint(uint64(n), 10))
	}

	return sb.String(), 0
}

func parseComputer(inputFile string) Computer {
	parts := util.ReadFileSplitBy(inputFile, "\n\n")

	registers := strings.Split(parts[0], "\n")
	a, _ := strconv.Atoi(registers[0][12:])
	b, _ := strconv.Atoi(registers[1][12:])
	c, _ := strconv.Atoi(registers[2][12:])

	program := util.StringsToInts(strings.Split(parts[1][9:], ","))
	programUints := make([]uint, len(program))
	for i, p := range program {
		programUints[i] = uint(p)
	}

	return Computer{
		A:                  uint(a),
		B:                  uint(b),
		C:                  uint(c),
		InstructionPointer: 0,
		Program:            programUints,
	}
}
