package advent14

import (
	"advent2024/util"
	"strconv"
	"strings"
)

type Robot struct {
	x, y   int
	vx, vy int
}

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)
	quadrantTotals := [][]int{
		{0, 0},
		{0, 0},
	}

	const durationInSeconds = 100
	roomHeight, roomWidth := 103, 101
	if strings.HasPrefix(inputFile, "sample") {
		roomHeight, roomWidth = 7, 11
	}

	for _, line := range lines {
		r := parseRobot(line)

		r.x = (r.x + r.vx*durationInSeconds) % roomWidth
		if r.x < 0 {
			r.x += roomWidth
		}
		r.y = (r.y + r.vy*durationInSeconds) % roomHeight
		if r.y < 0 {
			r.y += roomHeight
		}

		if r.x == roomWidth/2 || r.y == roomHeight/2 {
			continue
		}

		semi := (r.x + roomWidth/2) / roomWidth
		quad := (r.y + roomHeight/2) / roomHeight
		quadrantTotals[semi][quad] += 1
	}

	part1Total := 1
	for _, semi := range quadrantTotals {
		for _, q := range semi {
			part1Total *= q
		}
	}
	return part1Total, 0
}

func parseRobot(line string) Robot {
	r := Robot{}

	parts := strings.Split(line, " ")
	posParts := strings.Split(parts[0], ",")
	r.x, _ = strconv.Atoi(posParts[0][2:])
	r.y, _ = strconv.Atoi(posParts[1])

	velParts := strings.Split(parts[1], ",")
	r.vx, _ = strconv.Atoi(velParts[0][2:])
	r.vy, _ = strconv.Atoi(velParts[1])

	return r
}
