package main

import (
	"errors"
	"strconv"
	"strings"
)

func parseA1(s string) (i int, err error) {
	n := 0
	s = strings.ToUpper(s)
	for _, c := range s {
		if c < 'A' || c > 'Z' {
			return 0, errors.New("invalid column:" + s)
		}

		n = n*26 + int(c-'A'+1)
	}
	return n - 1, nil
}

func parseColumn(s string) (i int, err error) {
	if s == "" {
		return 0, errors.New("invalid column:" + s)
	}

	fst := strings.ToUpper(s[:1])
	if fst >= "A" && fst <= "Z" {
		return parseA1(s)
	} else if fst >= "0" && fst <= "9" {
		i, err := strconv.Atoi(s)
		if i < 1 {
			return 0, errors.New("column must be greater than 0")
		} else if err != nil {
			return 0, err
		} else {
			return i - 1, nil
		}
	} else {
		return 0, errors.New("invalid column:" + s)
	}
}

func ParseRange(s string) ([]int, error) {
	if s == "" {
		return []int{}, nil
	}

	ranges := strings.Split(s, ",")
	columns := make([]int, 0)
	for _, r := range ranges {
		start, end, _ := strings.Cut(r, ":")
		if end == "" {
			if idx, err := parseColumn(start); err == nil {
				columns = append(columns, idx)
			} else {
				return nil, err
			}

			continue
		}

		startIndex, err := parseColumn(start)
		if err != nil {
			return nil, err
		}

		endIndex, err := parseColumn(end)
		if err != nil {
			return nil, err
		}

		step := 1
		if startIndex > endIndex {
			step = -1
		}

		for i := startIndex; i != endIndex+step; i += step {
			columns = append(columns, i)
		}
	}

	return columns, nil
}
