# cj

convert CSV to JSON Lines

## usage

```sh
# convert CSV file to JSON Lines
$ cj input.csv > output.jsonl

# read from stdin if no file is specified
$ cj < input.csv > output.jsonl

# specify columns to convert in A1 like notation
$ cj -s AA:AC,BC:CD input.csv > output.jsonl
```


## installation

download binary from [releases](https://github.com/musou1500/cj/releases) and place it in your PATH.


## A1 like notation

`-s` option is used to specify columns to convert in A1 like notation.
but unlike A1 notation, you can not specify a range of rows. (e.g. 'A1:B2' is not supported)

by using A1 like notation you can

* select single column
  * e.g. `A`
* select range of columns
  * e.g. `A:C` 
* select multiple ranges of columns
  * e.g. `A:C,E:G`

in the following examples, notation is written in upper case. but it is case-insensitive.

suppose you have a CSV file like this.

```csv
A,B,C,...,Z,AA,AB,AC
0,1,2,...,25,26,27,28
```

### select single column

You can specify single column by sequence of alphabets.

```bash
$ cj -s AA input.csv # or `cj -s aa input.csv`
["AA"]
["26"]
```

### select range of columns

You can select range of columns by following notation.

```bash
$ cj -s AA:AC input.csv
["AA","AB","AC"]
["26","27","28"]
```

If left side of colon is larger than right side, select columns in reverse order.

```bash
$ cj -s AC:AA input.csv
["AC","AB","AA"]
["28","27","26"]
```

### select multiple ranges of columns

You can specify multiple ranges by separating them with comma.

```bash
$ cj -s A:B,D:E input.csv
["A","B","D","E"]
["0","1","3","4"]
```

overlapping ranges are supported. if then, concatenated columns are selected.

```bash
$ cj -s A:C,B:D ./input.csv
["A","B","C","B","C","D"]
["0","1","2","1","2","3"]
```
