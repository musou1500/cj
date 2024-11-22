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


## A1 like notation

A1 like notation is a way to specify columns in a CSV file.

unlike A1 notation, you can not specify a range of rows. (e.g. 'A1:B2' is not supported)

You can specify multiple ranges separated by commas.


## installation

download binary from [releases](https://github.com/musou1500/cj/releases) and place it in your PATH.
