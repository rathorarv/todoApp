package test

import (
	"testing"
)

func Test_getQueries(t *testing.T) {
	type args struct {
		s string
	}
	type Input struct {
		name string
		args args
		want string
	}
	tests := []Input{
		{"testing of getQueries", args{"fetch"}, "select * from todo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getQueries(tt.args.s); got != tt.want {
				t.Errorf("getQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
