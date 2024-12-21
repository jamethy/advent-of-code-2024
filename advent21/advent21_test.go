package advent21

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
			wantPart1: 126384,
		},
		{
			name:      "input",
			wantPart1: 163086,
			wantPart2: 198466286401228,
		},
		{
			name:      "reddit-1",
			wantPart1: 151826,
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

func TestSequenceLength(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "159A",
			want:  82,
		},
		{
			input: "375A",
			want:  70,
		},
		{
			input: "613A",
			want:  62,
		},
		{
			input: "894A",
			want:  78,
		},
		{
			input: "080A",
			want:  60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			res := sequenceLength(tt.input, 2)
			if res != tt.want {
				t.Errorf("sequenceLength got = %v, want %v", res, tt.want)
			}
		})
	}
}
