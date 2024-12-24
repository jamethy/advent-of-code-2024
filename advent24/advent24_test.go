package advent24

import (
	"advent2024/util/mathutil"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name      string
		wantPart1 any
		wantPart2 any
	}{
		{
			name:      "sample",
			wantPart1: uint(2024),
		},
		{
			name:      "input",
			wantPart1: uint(66055249060558),
			wantPart2: "fcd,fhp,hmk,rvf,tpc,z16,z20,z33",
		},
	}
	for _, tt := range tests {
		gotPart1, gotPart2 := Solution(tt.name + ".txt")
		t.Run(tt.name+"-part1", func(t *testing.T) {
			if tt.wantPart1 != nil && !reflect.DeepEqual(gotPart1, tt.wantPart1) {
				t.Errorf("Solution() gotPart1 = %v, want %v", gotPart1, tt.wantPart1)
			}
		})
		t.Run(tt.name+"-part2", func(t *testing.T) {
			if tt.wantPart2 != nil && !reflect.DeepEqual(gotPart2, tt.wantPart2) {
				t.Errorf("Solution() gotPart2 = %v, want %v", gotPart2, tt.wantPart2)
			}
		})
	}
}

func TestCalculation(t *testing.T) {
	nodes := parseNodes("input.txt")
	zKeys := findZKeys(nodes)

	swapNodes("z16", "hmk", nodes) // z16
	swapNodes("z20", "fhp", nodes) // z20
	swapNodes("tpc", "rvf", nodes) // z27
	swapNodes("z33", "fcd", nodes) // z33

	nodes["z15"].PrintInvolved()
	nodes["z16"].PrintInvolved()
	nodes["z45"].PrintInvolved()

	m := mathutil.IntPow(2, 33)

	x, y := uint(0), uint(m)
	setInputs(x, y, nodes)
	res := runCalculation(nodes, zKeys)
	if res != x+y {
		t.Errorf("Not equal: expected %v, got %v", x+y, res)
	}
}

func FuzzCalculation(f *testing.F) {
	nodes := parseNodes("input.txt")
	zKeys := findZKeys(nodes)

	f.Add(uint(0), uint(0))

	f.Fuzz(func(t *testing.T, x, y uint) {
		setInputs(x, y, nodes)
		res := runCalculation(nodes, zKeys)
		if res != x+y {
			t.Errorf("Not equal: expected %v, got %v", x+y, res)
		}
	})
}
