package set

import (
	"reflect"
	"testing"
)

func TestStringSet_Retain(t *testing.T) {
	type args struct {
		str []string
	}
	tests := []struct {
		name string
		s    Set[string]
		args args
		want Set[string]
	}{
		{
			name: "test1",
			s:    NewSet("one", "two"),
			args: args{
				str: []string{"two", "three"},
			},
			want: NewSet("two"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Retain(tt.args.str...)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("Retain() gotPart1 = %v, want %v", tt.s, tt.want)
			}
		})
	}
}
