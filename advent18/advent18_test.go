package advent18

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
			wantPart1: 22,
			wantPart2: "6,1",
		},
		{
			name:      "input",
			wantPart1: 356,
			wantPart2: "22,33",
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
