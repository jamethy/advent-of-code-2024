package advent22

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
			wantPart1: uint(37327623),
		},
		{
			name:      "sample-2",
			wantPart2: 23,
		},
		{
			name:      "input",
			wantPart1: uint(14691757043),
			wantPart2: 1831,
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

func TestPart1Numbers(t *testing.T) {
	tests := []struct {
		input uint
		want  uint
	}{
		{
			input: 1,
			want:  8685429,
		},
		{
			input: 10,
			want:  4700978,
		},
		{
			input: 100,
			want:  15273692,
		},
		{
			input: 2024,
			want:  8667524,
		},
	}
	for _, tt := range tests {
		res := calculateNthSecret(tt.input, 2000)
		if res != tt.want {
			t.Errorf("calculateNthSecret got = %v, want = %v", res, tt.want)
		}
	}
}
