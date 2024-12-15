package advent14

import (
	"advent2024/util"
	"bytes"
	"os"
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

	roomHeight, roomWidth := 103, 101
	if strings.HasPrefix(inputFile, "sample") {
		roomHeight, roomWidth = 7, 11
	}

	for _, line := range lines {
		r := parseRobot(line)
		r = move(r, 100, roomWidth, roomHeight)

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

	if strings.HasPrefix(inputFile, "sample") {
		return part1Total, 0
	}

	robots := make([]Robot, 0)
	for _, line := range lines {
		robots = append(robots, parseRobot(line))
	}

	for j, r := range robots {
		robots[j] = move(r, 8050, roomWidth, roomHeight)
	}

	//for i := 0; i < 200; i++ {
	//	for j, r := range robots {
	//		robots[j] = move(r, 101, roomWidth, roomHeight)
	//	}
	//
	//	num := fmt.Sprintf("%d", i+1)
	//	if i+1 < 100 {
	//		num = "0" + num
	//	}
	//	if i+1 < 10 {
	//		num = "0" + num
	//	}
	//
	//	f, _ := os.Create(fmt.Sprintf("output-%s.txt", num))
	//	_, _ = f.WriteString(displayString(robots, roomWidth, roomHeight))
	//	_ = f.Close()
	//}

	f, _ := os.Create("final.txt")
	_, _ = f.WriteString(displayString(robots, roomWidth, roomHeight))
	_ = f.Close()

	return part1Total, 0
}

func move(r Robot, seconds int, roomWidth, roomHeight int) Robot {
	r.x = (r.x + r.vx*seconds) % roomWidth
	for r.x < 0 {
		r.x += roomWidth
	}
	r.y = (r.y + r.vy*seconds) % roomHeight
	for r.y < 0 {
		r.y += roomHeight
	}
	return r
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

func displayString(robots []Robot, width, height int) string {
	lines := make([][]byte, height)
	for i := range lines {
		lines[i] = bytes.Repeat([]byte(" "), width)
	}

	for _, r := range robots {
		lines[r.y][r.x] = 'X'
	}

	return string(bytes.Join(lines, []byte("\n")))
}
