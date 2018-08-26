package cmd

import "fmt"
import "os"
import "path/filepath"

import "github.com/mfinelli/wmuc/parser"

func CloneRepos(results map[string]parser.Project) {
	for _, project := range results {
		createDirectory(project.Path)

		for _, repo := range project.Repos {
			fmt.Printf("\trepo: %s\n", repo.Url)
		}
	}
}

func createDirectory(path string) {
	if path == "" {
		return
	}

	cwd, _ := os.Getwd()
	err := os.MkdirAll(filepath.Join(cwd, path), 0755)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
