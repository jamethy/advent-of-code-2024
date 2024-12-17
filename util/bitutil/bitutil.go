package bitutil

import (
	"strconv"
	"strings"
)

func IsBitSet(i uint, pos uint) bool {
	return i&(1<<pos) > 0
}

func SetBit(i uint, pos uint, on bool) uint {
	if on {
		i |= 1 << pos
	} else {
		i &^= 1 << pos
	}
	return i
}

func IsBitSet64(i uint64, pos int) bool {
	return i&(1<<pos) > 0
}

func SetBit64(i uint64, pos int) uint64 {
	i |= 1 << pos
	return i
}

func FlipAllBits(i uint, bits uint) uint {
	maxVal := uint(1)<<bits - 1
	return i ^ maxVal
}

func ParseBinary(s string) (uint, error) {
	s = strings.TrimSpace(s)
	u, err := strconv.ParseUint(s, 2, len(s))
	return uint(u), err
}

func XOR(a uint, b uint) uint {
	var c uint
	var i uint
	for i = 0; i < 32; i++ {
		aSet := IsBitSet(a, i)
		bSet := IsBitSet(b, i)
		c = SetBit(c, i, (aSet || bSet) && !(aSet && bSet))
	}
	return c
}
