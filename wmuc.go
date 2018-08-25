package main

import "fmt"

import "github.com/mfinelli/wmuc/parser"

func main() {
	input := `# this is a comment
	repo "test one"
	repo 'test two'`
	results := parser.Parse(input)

	for _, project := range results {
		fmt.Printf("For project: %q\n", project.Path)

		for _, repo := range project.Repos {
			fmt.Printf("\trepo: %s\n", repo.Url)
		}
	}
}
