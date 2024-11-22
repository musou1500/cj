package main

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	tests := []struct {
		a1       string
		expected int
	}{
		{"A", 0},
		{"B", 1},
		{"Z", 25},
		{"AA", 26},
		{"AB", 27},
		{"AZ", 51},
	}
	for _, tt := range tests {
		t.Run(tt.a1, func(t *testing.T) {
			got := A1ToIndex(tt.a1)
			if got != tt.expected {
				t.Errorf("A1ToIndex(%s) == %d, want %d", tt.a1, got, tt.expected)
			}
		})
	}
}

func TestA1ToIndeces(t *testing.T) {
	tests := []struct {
		a1       string
		expected []int
	}{
		{"", []int{}},
		{"AA", []int{26}},
		{"A:C", []int{0, 1, 2}},
		{"A:C,E", []int{0, 1, 2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.a1, func(t *testing.T) {
			got := A1ToIndeces(tt.a1)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("A1ToIndeces(%s) == %v, want %v", tt.a1, got, tt.expected)
			}
		})
	}
}
