package main

import "fmt"

import "github.com/mfinelli/wmuc/parser"

func main() {
	input := `# this is a comment
	repo "test one"
	repo 'test two'`
	results := parser.Parse(input)

	for _, token := range results {
		fmt.Print(token)
	}
}
