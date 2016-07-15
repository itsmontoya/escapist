# Escapist [![GoDoc](https://godoc.org/github.com/itsmontoya/escapist?status.svg)](https://godoc.org/github.com/itsmontoya/escapist) ![Status](https://img.shields.io/badge/status-alpha-red.svg)
Escapist is a byteslice-focused HTML escaping library

## Benchmarks
```bash
# Basic test, very simple replacement
BenchmarkBasic-4        10000000               142 ns/op              48 B/op          1 allocs/op
BenchmarkHTMLBasic-4     5000000               271 ns/op              96 B/op          2 allocs/op
BenchmarkAdvBasic-4     10000000               173 ns/op              48 B/op          1 allocs/op

# No replacement
BenchmarkNoRep-4        30000000                58.3 ns/op             0 B/op          0 allocs/op
BenchmarkHTMLNoRep-4    20000000                64.1 ns/op             0 B/op          0 allocs/op
BenchmarkAdvNoRep-4     20000000               100 ns/op               0 B/op          0 allocs/op

# Note: escapist.Escape and html.EscapeString look at: ', ", <, >, and &
# This is considered to be a very basic escape.

# In contrast, escapist.EscapeAdv looks at:
# &, <, >, ", ', `, !, @, $, %, (, ), =, +, \{, \}, \[, and \]*
```

## Usage
```go
package main

import "fmt"

func main() {
	b := Escape([]byte("Hello there Mr. Joe <inject stuff>"))
	fmt.Println(string(b))
}
```