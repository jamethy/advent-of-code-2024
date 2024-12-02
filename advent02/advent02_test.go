package advent02

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
			wantPart1: 2,
			wantPart2: 4,
		},
		{
			name:      "input",
			wantPart1: 218,
			wantPart2: 290,
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

func TestReportIsSafePart2(t *testing.T) {
	tests := []struct {
		values []int
		want   bool
	}{
		{
			values: []int{68, 73, 71, 74, 75},
			want:   true,
		},
		{
			values: []int{28, 31, 34, 35, 38, 39, 43},
			want:   true,
		},
		{
			values: []int{41, 43, 42, 44, 45, 45},
			want:   false,
		},
	}
	for _, tt := range tests {
		_, part2 := reportIsSafe(tt.values)
		if part2 != tt.want {
			t.Errorf("reportIsSafe part 2 = %v, want %v for %v", part2, tt.want, tt.values)
		}
	}
}
