package main

import (
	"strings"
)

func A1ToIndex(s string) int {
	if s == "" {
		panic("empty column name")
	}

	n := 0
	s = strings.ToUpper(s)
	for _, c := range s {
		if c < 'A' || c > 'Z' {
			panic("invalid column name:" + s)
		}

		n = n*26 + int(c-'A'+1)
	}
	return n - 1
}

func A1ToIndeces(s string) []int {
	if s == "" {
		return []int{}
	}

	ranges := strings.Split(s, ",")
	columns := make([]int, 0)
	for _, r := range ranges {
		start, end, _ := strings.Cut(r, ":")
		if end == "" {
			columns = append(columns, A1ToIndex(start))
			continue
		}

		startIndex := A1ToIndex(start)
		endIndex := A1ToIndex(end)
		step := 1
		if startIndex > endIndex {
			step = -1
		}

		for i := startIndex; i != endIndex+step; i += step {
			columns = append(columns, i)
		}
	}

	return columns
}
