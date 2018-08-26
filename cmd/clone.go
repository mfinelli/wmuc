package cmd

import "fmt"
import "os"
import "path"
import "path/filepath"
import "strings"

import "gopkg.in/src-d/go-git.v4"
import "github.com/mfinelli/wmuc/parser"

func CloneRepos(results map[string]parser.Project) {
	cwd, _ := os.Getwd()

	for _, project := range results {
		createDirectory(project.Path)
		os.Chdir(filepath.Join(cwd, project.Path))

		for _, repo := range project.Repos {
			repopath := path.Base(repo.Url)
			repodir := strings.TrimSuffix(repopath,
				filepath.Ext(repopath))

			if _, err := os.Stat(repodir); os.IsNotExist(err) {
				_, err := git.PlainClone(repodir, false,
					&git.CloneOptions{
						URL:      repo.Url,
						Progress: os.Stdout,
					})

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		}

		if project.Path != "" {
			os.Chdir(cwd)
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
