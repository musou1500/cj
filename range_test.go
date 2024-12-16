package main

import (
	"reflect"
	"testing"
)

func TestParseRange(t *testing.T) {
	tests := []struct {
		a1       string
		expected []int
	}{
		{"", []int{}},
		{"Z", []int{25}},
		{"z", []int{25}},
		{"3", []int{2}},
		{"AA", []int{26}},
		{"AB", []int{27}},
		{"AZ", []int{51}},
		{"AAA", []int{702}},
		{"AA", []int{26}},
		{"A:C", []int{0, 1, 2}},
		{"A:2", []int{0, 1}},
		{"C:A", []int{2, 1, 0}},
		{"A:C,E", []int{0, 1, 2, 4}},
		{"A:C,B:D", []int{0, 1, 2, 1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.a1, func(t *testing.T) {
			got, err := ParseRange(tt.a1)
			if err != nil {
				t.Errorf("ParseRange(%s) returned an error: %v", tt.a1, err)
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("A1ToIndeces(%s) == %v, want %v", tt.a1, got, tt.expected)
			}
		})
	}
}
