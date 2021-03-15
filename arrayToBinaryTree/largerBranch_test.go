package main

import (
	"testing"
)

func Test_largerBranch(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Base", args{[]int{3, 6, 2, 9, -1, 10}}, "Left"},
		{"Base", args{[]int{3, 6, 2, 9, -1, 10, 1, 4, 5, 6, 7, 2, 3, 4, 14}}, "Left"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largerBranch(tt.args.arr); got != tt.want {
				t.Errorf("largerBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}
