package main

import "fmt"

import "github.com/mfinelli/wmuc/parser"

func main() {
	input := `repo "`
	results := parser.Parse(input)

	for _, token := range results {
		fmt.Print(token)
	}
}
