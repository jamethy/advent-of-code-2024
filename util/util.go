package util

import (
	"io"
	"os"
	"strconv"
	"strings"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func StringsToInts(str []string) []int {
	ret := make([]int, 0, len(str))
	for _, s := range str {
		if i, err := strconv.Atoi(s); err == nil {
			ret = append(ret, i)
		}
	}
	return ret
}

func ParseIntList(str, sep string) []int {
	parts := strings.Split(str, sep)
	return StringsToInts(parts)
}

func IntsToStrings(ints []int) []string {
	str := make([]string, len(ints))
	for i, n := range ints {
		str[i] = strconv.Itoa(n)
	}
	return str
}

func IntsToString(ints []int, sep string) string {
	strs := IntsToStrings(ints)
	return strings.Join(strs, sep)
}

func ReadFile(name string) []string {
	lines := ReadFileSplitBy(name, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

func ReadFileAsString(name string) string {
	f, err := os.Open(name)
	if err != nil {
		panic("can't open file " + err.Error())
	}
	d, err := io.ReadAll(f)
	if err != nil {
		panic("can't read file " + err.Error())
	}

	return string(d)
}

func ReadFileSplitBy(name, delimiter string) []string {
	s := ReadFileAsString(name)
	return strings.Split(s, delimiter)
}

func FlipString(str string) string {
	str2 := ""
	for i := len(str) - 1; i >= 0; i-- {
		str2 += string(str[i])
	}
	return str2
}

func EqualIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
