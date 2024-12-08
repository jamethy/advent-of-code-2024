package advent08

import (
	"advent2024/util"
	"advent2024/util/set"
)

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)
	maxY, maxX := len(lines)-1, len(lines[0])-1

	runePoints := make(map[rune][]Point)

	for y, line := range lines {
		for x, r := range line {
			if r == '.' {
				continue
			}
			runePoints[r] = append(runePoints[r], Point{x: x, y: y})
		}
	}

	uniquePoints := set.NewSet[Point]()
	for _, points := range runePoints {
		for _, a := range points {
			for _, b := range points {
				if a == b {
					continue
				}
				diff := a.minus(b)

				p := a.plus(diff)
				if p.within(maxX, maxY) {
					uniquePoints.Add(p)
				}
				p = b.minus(diff)
				if p.within(maxX, maxY) {
					uniquePoints.Add(p)
				}
			}
		}
	}

	return len(uniquePoints), 0
}

type Point struct {
	x, y int
}

func (p Point) minus(o Point) Point {
	return Point{
		x: p.x - o.x,
		y: p.y - o.y,
	}
}

func (p Point) plus(o Point) Point {
	return Point{
		x: p.x + o.x,
		y: p.y + o.y,
	}
}

func (p Point) within(maxX, maxY int) bool {
	return p.x >= 0 && p.y >= 0 && p.x <= maxX && p.y <= maxY
}
