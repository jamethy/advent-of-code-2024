package advent15

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
			name:      "my-issue",
			wantPart1: 2708,
			wantPart2: 2818,
		},
		{
			name:      "small",
			wantPart1: 2028,
			wantPart2: 1751,
		},
		{
			name:      "small-2",
			wantPart1: 908,
			wantPart2: 618,
		},
		{
			name:      "sample",
			wantPart1: 10092,
			wantPart2: 9021,
		},
		{
			name:      "input",
			wantPart1: 1514353,
			wantPart2: 1533076,
		},
		{
			name:      "reddit-1",
			wantPart1: 1816,
			wantPart2: 1430,
		},
		{
			name:      "reddit-2",
			wantPart1: 403,
			wantPart2: 406,
		},
		{
			name:      "reddit-3",
			wantPart1: 504,
			wantPart2: 509,
		},
		{
			name:      "reddit-4",
			wantPart1: 1109,
			wantPart2: 822,
		},
		{
			name:      "reddit-5",
			wantPart1: 505,
			wantPart2: 511,
		},
		{
			name:      "reddit-6",
			wantPart1: 1207,
			wantPart2: 816,
		},
		{
			name:      "reddit-7",
			wantPart1: 2727,
			wantPart2: 3441,
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
			if tt.wantPart2 == nil {
				return
			}
			if !reflect.DeepEqual(gotPart2, tt.wantPart2) {
				t.Errorf("Solution() gotPart2 = %v, want %v", gotPart2, tt.wantPart2)
			}
		})
	}
}
