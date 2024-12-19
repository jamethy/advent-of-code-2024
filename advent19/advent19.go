package advent19

import (
	"advent2024/util"
	"advent2024/util/mathutil"
	"advent2024/util/set"
	"math"
	"strings"
)

func Solution(inputFile string) (part1, part2 any) {
	parts := util.ReadFileSplitBy(inputFile, "\n\n")

	patternSlice := strings.Split(parts[0], ", ")
	designer := NewDesigner(patternSlice)

	designs := strings.Split(parts[1], "\n")

	possibleDesigns := 0
	for _, design := range designs {
		if designer.isDesignPossible(design) {
			possibleDesigns++
		}
	}

	return possibleDesigns, 0
}

func (d Designer) isDesignPossible(design string) bool {
	if len(design) == 0 || d.patterns.Has(design) {
		return true
	}
	if len(design) < d.minLength {
		return false
	}
	maxSubDesignLength := mathutil.MinInt(d.maxLength, len(design))
	for l := d.minLength; l <= maxSubDesignLength; l++ {
		subDesign := design[0:l]
		if !d.patterns.Has(subDesign) {
			continue
		}
		remainingDesign := design[l:]
		if d.isDesignPossible(remainingDesign) {
			return true
		}
	}
	return false
}

type Designer struct {
	patterns             set.Set[string]
	minLength, maxLength int
}

func NewDesigner(patternSlice []string) Designer {
	patterns := set.NewSet[string]()
	minLength, maxLength := math.MaxInt, 0
	for _, pattern := range patternSlice {
		patterns.Add(pattern)
		minLength = mathutil.MinInt(minLength, len(pattern))
		maxLength = mathutil.MaxInt(maxLength, len(pattern))
	}
	return Designer{
		patterns:  patterns,
		minLength: minLength,
		maxLength: maxLength,
	}
}
