package advent17

import (
	"advent2024/util"
	"advent2024/util/bitutil"
	"advent2024/util/mathutil"
	"reflect"
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
	for c.InstructionPointer < uint(len(c.Program)) {
		c.Step()
	}
}

func (c *Computer) Step() {

	opcode := c.Program[c.InstructionPointer]
	operand := c.Program[c.InstructionPointer+1]
	c.InstructionPointer += 2

	switch opcode {
	case 0: // adv
		c.A = c.A / uint(mathutil.IntPow(2, int(c.ComboValue(operand))))
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
		c.B = c.A / uint(mathutil.IntPow(2, int(c.ComboValue(operand))))
	case 7: // cdv
		c.C = c.A / uint(mathutil.IntPow(2, int(c.ComboValue(operand))))
	}
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

func (c *Computer) OutputString() string {
	sb := strings.Builder{}
	for i, n := range c.Output {
		if i != 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(strconv.FormatUint(uint64(n), 10))
	}
	return sb.String()
}

func findPart2A(program []uint, want []uint, prev uint) (uint, bool) {
	for a := uint(0); a < 8; a++ {
		newA := 8*prev + a

		sub := Computer{A: newA, Program: program}
		sub.Run()

		if !reflect.DeepEqual(sub.Output, want) {
			continue
		}
		if len(program) == len(want) {
			return newA, true
		}

		subA, ok := findPart2A(program, program[len(program)-len(want)-1:], newA)
		if ok {
			return subA, true
		}
	}
	return 0, false
}

func Solution(inputFile string) (part1, part2 any) {

	c := parseComputer(inputFile)
	c.Run()
	part1 = c.OutputString()
	part2, _ = findPart2A(c.Program, []uint{0}, 0)

	return part1, part2
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
