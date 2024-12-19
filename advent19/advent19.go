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
	possibleDesignOptions := 0
	for _, design := range designs {
		c := designer.possibleDesignCount(design)
		if c != 0 {
			possibleDesigns++
			possibleDesignOptions += c
		}
	}

	return possibleDesigns, possibleDesignOptions
}

func (d Designer) possibleDesignCount(design string) int {
	if len(design) == 0 {
		return 1
	}
	if len(design) < d.minLength {
		return 0
	}
	if cached, ok := d.countCache[design]; ok {
		return cached
	}

	possibilities := 0
	maxSubDesignLength := mathutil.MinInt(d.maxLength, len(design))
	for l := d.minLength; l <= maxSubDesignLength; l++ {
		subDesign := design[0:l]
		if !d.patterns.Has(subDesign) {
			continue
		}
		remainingDesign := design[l:]
		possibilities += d.possibleDesignCount(remainingDesign)
	}
	d.countCache[design] = possibilities
	return possibilities
}

type Designer struct {
	patterns             set.Set[string]
	countCache           map[string]int
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
		patterns:   patterns,
		countCache: make(map[string]int, len(patterns)*1000),
		minLength:  minLength,
		maxLength:  maxLength,
	}
}
