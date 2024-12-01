package bitutil

import (
	"strconv"
	"strings"
)

func IsBitSet(i uint, pos uint) bool {
	return i&(1<<pos) > 0
}

func SetBit(i uint, pos uint) uint {
	i |= 1 << pos
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
