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

package chuckfile

import "fmt"
import "io/ioutil"
import "os"
import "path"

import "gopkg.in/src-d/go-git.v4"
import "gopkg.in/src-d/go-git.v4/plumbing"

import "github.com/mfinelli/wmuc/parser"
import "github.com/mfinelli/wmuc/util"

var allowed = [...]string{"chuckfile", "Chuckfile", "CHUCKFILE"}

func ChuckfileExists() bool {
	for _, filename := range allowed {
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			continue
		} else {
			return true
		}
	}

	return false
}

func currentChuckfile() string {
	for _, filename := range allowed {
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			continue
		} else {
			return filename
		}
	}

	return ""
}

func RemoveExistingChuckfile() {
	cf := currentChuckfile()

	if cf == "" {
		return
	}

	err := os.Remove(cf)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func WalkDirectory() []parser.Project {
	items, err := ioutil.ReadDir(".")

	projects := make([]parser.Project, 0)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	project := parser.Project{Path: ""}

	for _, item := range items {
		p := path.Join(".", item.Name())

		if !isDir(p) {
			continue
		}

		r, err := git.PlainOpen(p)

		if err == nil {
			origin := util.GetOriginRemote(r)
			branch := getCurrentBranch(r)

			// don't specify master if we're creating new
			if branch == "master" {
				branch = ""
			}

			repo := parser.Repo{Url: origin, Branch: branch}
			project.Repos = append(project.Repos, repo)
		} else {
			projects = append(projects, getReposInDir(p))
		}
	}

	return append(projects, project)
}

func getReposInDir(dir string) parser.Project {
	items, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	project := parser.Project{Path: dir}

	for _, item := range items {
		p := path.Join(".", dir, item.Name())

		if !isDir(p) {
			continue
		}

		r, err := git.PlainOpen(p)

		if err == nil {
			origin := util.GetOriginRemote(r)
			branch := getCurrentBranch(r)

			// don't specify master if we're creating new
			if branch == "master" {
				branch = ""
			}

			repo := parser.Repo{Url: origin, Branch: branch}
			project.Repos = append(project.Repos, repo)
		}
	}

	return project
}

func getCurrentBranch(repo *git.Repository) string {
	b, _ := repo.Head()
	// cast HEAD to a referenceName so we can call short
	ref := plumbing.ReferenceName(b.Name())
	return ref.Short()
}

func isDir(p string) bool {
	info, err := os.Stat(p)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return info.IsDir()
}
