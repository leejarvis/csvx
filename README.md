# csvx

csvx is a small wrapper around the `encoding/csv` package which aims to
abstract and build on the existing functionality.

## Install

    go get github.com/leejarvis/csvx

## Usage

```go
csv, err := csvx.Parse("somefile.csv")

// csv.Rows() returns a slice of csvx.Row types
// csv.Headers() returns a slice of strings
```
