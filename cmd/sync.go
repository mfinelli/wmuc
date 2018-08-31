// wmuc: a git repository manager
// Copyright (C) 2018  Mario Finelli
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import "errors"
import "fmt"
import "io/ioutil"
import "os"
import "path/filepath"

import "github.com/spf13/cobra"
import "github.com/spf13/viper"
import "gopkg.in/src-d/go-git.v4"
import "gopkg.in/src-d/go-git.v4/config"
import "gopkg.in/src-d/go-git.v4/plumbing"

import "github.com/mfinelli/wmuc/parser"

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "ensure disk matches chuckfile",
	Long: `Ensure the repositories and configuration specified in the
chuckfile match what is on disk.`,
	Run: func(cmd *cobra.Command, args []string) {
		input, err := readChuckfile()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		results := parser.Parse(input)
		cloneRepos(results)
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
	syncCmd.Flags().BoolP("debug", "d", false, "enable lexer/parser debug")
	syncCmd.Flags().BoolP("verbose", "v", false, "enable verbose mode")
	viper.BindPFlag("debug", syncCmd.Flags().Lookup("debug"))
	viper.BindPFlag("verbose", syncCmd.Flags().Lookup("verbose"))
}

func readChuckfile() (string, error) {
	allowed := []string{"chuckfile", "Chuckfile", "CHUCKFILE"}

	for _, filename := range allowed {
		bytes, err := ioutil.ReadFile(filename)

		if err == nil {
			return string(bytes), nil
		}
	}

	return "", errors.New("unable to load a chuckfile in this directory")
}

func cloneRepos(results map[string]parser.Project) {
	cwd, _ := os.Getwd()

	for _, project := range results {
		createDirectory(project.Path)
		os.Chdir(filepath.Join(cwd, project.Path))

		for _, repo := range project.Repos {
			repodir := repo.DirFromPath()

			if _, err := os.Stat(repodir); os.IsNotExist(err) {
				r, err := git.PlainClone(repodir, false,
					&git.CloneOptions{
						URL:      repo.Url,
						Progress: os.Stdout,
					})

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				if repo.Branch != "" {
					w, _ := r.Worktree()

					err := r.Fetch(&git.FetchOptions{
						RefSpecs: []config.RefSpec{
							"refs/*:refs/*",
							"HEAD:refs/heads/HEAD"},
					})

					if err != nil {
						fmt.Println(err)
						// proceed anyway...
						continue
					}

					b := plumbing.ReferenceName(repo.Branch)
					err = w.Checkout(&git.CheckoutOptions{
						Branch: b,
					})

					if err != nil {
						fmt.Println(err)
					}
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
