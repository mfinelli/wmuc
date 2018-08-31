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

			if getOriginRemote(r) != repo.Url {
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

func getOriginRemote(repo *git.Repository) string {
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
