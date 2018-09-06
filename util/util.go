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

package util

import "os"
import "path"
import "path/filepath"

import "gopkg.in/src-d/go-git.v4"
import "gopkg.in/src-d/go-git.v4/plumbing"

import "github.com/mfinelli/wmuc/parser"

func InSync(expected map[string]parser.Project) map[string][]string {
	cwd, _ := os.Getwd()

	sync := make(map[string][]string)

	for _, project := range expected {
		if !directoryExists(project.Path) {
			sync["missingProject"] = append(sync["missingProject"],
				project.Path)
			continue
		}

		os.Chdir(filepath.Join(cwd, project.Path))

		for _, repo := range project.Repos {
			repodir := repo.DirFromPath()

			if !directoryExists(repodir) {
				sync["missingRepo"] = append(
					sync["missingRepo"], path.Join(
						project.Path, repodir))
				continue
			}

			r, err := git.PlainOpen(repodir)

			if err != nil {
				sync["unableToOpen"] = append(
					sync["unableToOpen"], path.Join(
						project.Path, repodir))
				continue
			}

			if GetOriginRemote(r) != repo.Url {
				sync["wrongRepo"] = append(sync["wrongRepo"],
					path.Join(project.Path, repodir))
				continue
			}

			if !correctBranch(r, repo.Branch) {
				sync["wrongBranch"] = append(
					sync["wrongBranch"], path.Join(
						project.Path, repodir))
			}
		}

		os.Chdir(cwd)
	}

	return sync
}

func directoryExists(path string) bool {
	if path == "" {
		return true
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func GetOriginRemote(repo *git.Repository) string {
	remotes, _ := repo.Remotes()

	for _, remote := range remotes {
		if remote.Config().Name == "origin" {
			return remote.Config().URLs[0]
		}
	}

	return ""
}

func correctBranch(repo *git.Repository, branch string) bool {
	if branch == "" {
		return true
	}

	b, _ := repo.Head()
	// cast HEAD to a referenceName so we can call short
	ref := plumbing.ReferenceName(b.Name())

	if ref.Short() == branch {
		return true
	}

	return false
}
