package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

func readCSV(r io.Reader) [][]string {
	rr := csv.NewReader(r)
	rows, err := rr.ReadAll()
	if err != nil {
		panic(err)
	}

	return rows
}

func selectColumns(row [][]string, columns []int) [][]string {
	newRows := make([][]string, 0, len(row))
	for _, row := range row {
		newRow := make([]string, 0)
		// if columns is empty, select all columns
		if len(columns) == 0 {
			newRow = append(newRow, row...)
		} else {
			for _, c := range columns {
				newRow = append(newRow, row[c])
			}
		}

		newRows = append(newRows, newRow)
	}

	return newRows
}

func main() {
	var (
		ranges = flag.String("s", "", "range of columns to convert")
	)

	flag.Parse()

	var f *os.File
	if flag.NArg() == 0 {
		f = os.Stdin
	} else {
		var err error
		f, err = os.Open(flag.Arg(0))
		if err != nil {
			panic(err)
		}
		defer f.Close()
	}

	rows := selectColumns(readCSV(f), A1ToIndeces(*ranges))
	for _, row := range rows {
		json_bytes, err := json.Marshal(row)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", json_bytes)
	}
}
