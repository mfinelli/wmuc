package cmd

import "fmt"
import "github.com/mfinelli/wmuc/parser"

func CloneRepos(results map[string]parser.Project) {
	for _, project := range results {
		fmt.Printf("For project: %q\n", project.Path)

		for _, repo := range project.Repos {
			fmt.Printf("\trepo: %s\n", repo.Url)
		}
	}
}
