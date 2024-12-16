package advent16

import (
	"advent2024/util"
	"advent2024/util/set"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

var directions = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

type Point struct {
	i, j int
}

type Path struct {
	Points set.Set[Point]
	Score  int
}

func (p Path) Clone() Path {
	return Path{
		Points: p.Points.Clone(),
		Score:  p.Score,
	}
}

var bestScore = -1

func Solution(inputFile string) (part1, part2 any) {
	bestScore = -1
	grid := util.ReadFileAsByteGrid(inputFile)

	scoreGrid := make([][]int, len(grid))
	for i := range grid {
		scoreGrid[i] = make([]int, len(grid[i]))
	}

	i, j := findTile(grid, 'S')
	paths := move(grid, scoreGrid, i, j, 0, Path{
		Points: set.NewSet(Point{i: i, j: j}),
		Score:  0,
	})

	//bestScore = paths[0].Score
	//for _, p := range paths {
	//	if p.Score < bestScore {
	//		bestScore = p.Score
	//	}
	//}
	//
	potentialSeats := set.NewSet[Point]()
	for _, path := range paths {
		if path.Score != bestScore {
			continue
		}
		potentialSeats.AddAll(path.Points)
	}

	//printScoreGrid(grid, scoreGrid)

	for seat := range potentialSeats {
		grid[seat.i][seat.j] = 'O'
	}
	for _, line := range grid {
		fmt.Println(string(line))
	}
	fmt.Println()

	return bestScore, len(potentialSeats)
}

func move(grid [][]byte, scoreGrid [][]int, i, j, dirIdx int, path Path) []Path {
	var paths []Path

	for dirChange := 0; dirChange < 4; dirChange++ {
		nDirIdx := (dirIdx + dirChange) % 4
		dir := directions[nDirIdx]

		ni, nj := i+dir[0], j+dir[1]
		if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[0]) {
			continue
		}
		if grid[ni][nj] == '#' {
			continue
		}

		// todo figure out clever way
		var rotationCost int
		switch dirChange {
		case 0:
			rotationCost = 0
		case 1, 3:
			rotationCost = 1000
		case 2:
			rotationCost = 2000
		}

		nextPoint := Point{i: ni, j: nj}
		nextScore := path.Score + rotationCost + 1
		if scoreGrid[ni][nj] == 0 || nextScore < scoreGrid[ni][nj] {
			scoreGrid[ni][nj] = nextScore
		} else if scoreGrid[ni][nj] != 0 && nextScore-2000 > scoreGrid[ni][nj] {
			continue
		}

		if !path.Points.Has(nextPoint) {
			newPath := path.Clone()
			newPath.Points.Add(nextPoint)
			newPath.Score = nextScore

			if bestScore > 0 && newPath.Score > bestScore {
				continue
			}

			if grid[ni][nj] == 'E' {
				bestScore = newPath.Score
				return []Path{newPath}
			}

			movePaths := move(grid, scoreGrid, ni, nj, nDirIdx, newPath)
			if len(movePaths) > 0 {
				paths = append(paths, movePaths...)
			}
		}
	}
	return paths
}

func findTile(grid [][]byte, c byte) (int, int) {
	for i, line := range grid {
		for j, v := range line {
			if v == c {
				return i, j
			}
		}
	}
	return -1, -1
}

func printScoreGrid(grid [][]byte, scoreGrid [][]int) {
	sb := strings.Builder{}
	for i, line := range grid {
		for j, v := range line {
			if v == 'S' || v == 'E' {
				sb.Write(bytes.Repeat([]byte{v}, 4))
				sb.WriteRune(' ')
				continue
			}
			score := scoreGrid[i][j]
			s := strconv.Itoa(score)
			if score < 1000 {
				s = "0" + s
			}
			if score < 100 {
				s = "0" + s
			}
			if score < 10 {
				s = "0" + s
			}
			sb.WriteString(s + " ")
		}
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}
