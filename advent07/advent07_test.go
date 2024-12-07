package advent07

import (
	"fmt"
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
			wantPart1: uint64(3749),
			wantPart2: 0,
		},
		{
			name:      "input",
			wantPart1: uint64(12839601725877),
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

func TestTestValues(t *testing.T) {
	tests := []struct {
		testValue uint64
		values    []uint64
		want      bool
	}{
		{
			testValue: 190,
			values:    []uint64{10, 19},
			want:      true,
		},
		{
			testValue: 3267,
			values:    []uint64{81, 40, 27},
			want:      true,
		},
		{
			testValue: 83,
			values:    []uint64{17, 5},
			want:      false,
		},
		//156: 15 6
		//7290: 6 8 6 15
		{
			testValue: 161011,
			values:    []uint64{16, 10, 13},
			want:      false,
		},
		//192: 17 8 14
		{
			testValue: 21037,
			values:    []uint64{9, 7, 18, 13},
			want:      false,
		},
	}
	//292: 11 6 16 20
	for i, tt := range tests {
		t.Run(fmt.Sprintf("line %d", i), func(t *testing.T) {
			if got := TestValues(tt.testValue, tt.values); got != tt.want {
				t.Errorf("TestValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
