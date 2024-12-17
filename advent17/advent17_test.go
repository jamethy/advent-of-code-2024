package advent17

import (
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
			wantPart1: "4,6,3,5,6,3,5,2,1,0",
			wantPart2: 0,
		},
		{
			name:      "input",
			wantPart1: "2,1,3,0,5,2,3,7,1",
			wantPart2: 0,
		},
	}
	for _, tt := range tests {
		gotPart1, gotPart2 := Solution(tt.name + ".txt")
		t.Run(tt.name+"-part1", func(t *testing.T) {
			if !reflect.DeepEqual(gotPart1, tt.wantPart1) {
				t.Errorf("Solution() gotPart1 = %v, want %v", gotPart1, tt.wantPart1)
			}
		})
		t.Run(tt.name+"-part2", func(t *testing.T) {
			if !reflect.DeepEqual(gotPart2, tt.wantPart2) {
				t.Errorf("Solution() gotPart2 = %v, want %v", gotPart2, tt.wantPart2)
			}
		})
	}
}

func TestComputer(t *testing.T) {
	//If register C contains 9, the program 2,6 would set register B to 1.
	t.Run("2,6", func(t *testing.T) {
		c := Computer{
			C:       9,
			Program: []uint{2, 6},
		}
		c.Run()
		if c.B != 1 {
			t.Errorf("TestComputer() got = %v", c.B)
		}
	})
	//If register A contains 10, the program 5,0,5,1,5,4 would output 0,1,2.
	t.Run("5,0,5,1,5,4", func(t *testing.T) {
		c := Computer{
			A:       10,
			Program: []uint{5, 0, 5, 1, 5, 4},
		}
		c.Run()
		if !reflect.DeepEqual(c.Output, []uint{0, 1, 2}) {
			t.Errorf("TestComputer() got = %v", c.Output)
		}
	})
	//If register A contains 2024, the program 0,1,5,4,3,0 would output 4,2,5,6,7,7,7,7,3,1,0 and leave 0 in register A.
	t.Run("0,1,5,4,3,0", func(t *testing.T) {
		c := Computer{
			A:       2024,
			Program: []uint{0, 1, 5, 4, 3, 0},
		}
		c.Run()
		if !reflect.DeepEqual(c.Output, []uint{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0}) {
			t.Errorf("TestComputer() got = %v", c.Output)
		}
		if c.A != 0 {
			t.Errorf("TestComputer() got = %v", c.A)
		}
	})
	//If register B contains 29, the program 1,7 would set register B to 26.
	t.Run("1,7", func(t *testing.T) {
		c := Computer{
			B:       29,
			Program: []uint{1, 7},
		}
		c.Run()
		if c.B != 26 {
			t.Errorf("TestComputer() got = %v", c.B)
		}
	})
	//If register B contains 2024 and register C contains 43690, the program 4,0 would set register B to 44354.
	t.Run("4,0", func(t *testing.T) {
		c := Computer{
			B:       2024,
			C:       43690,
			Program: []uint{4, 0},
		}
		c.Run()
		if c.B != 44354 {
			t.Errorf("TestComputer() got = %v", c.B)
		}
	})
}
