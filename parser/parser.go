package parser

import "fmt"
import "os"

import "github.com/mfinelli/wmuc/lexer"
import "github.com/mfinelli/wmuc/tokens"

func Parse(input string) map[string]Project {
	l := lexer.Lex(input)
	results := make(map[string]Project)

	currentProject := findOrCreateProject("", results)
	var currentRepo Repo

	for {
		token := l.NextToken()

		if token.Kind == tokens.TOKEN_ERROR {
			fmt.Println(token.Value)
			os.Exit(1)
		}

		if token.Kind == tokens.TOKEN_PROJECT_VALUE {
			currentProject = findOrCreateProject(token.Value,
				results)
		}

		if token.Kind == tokens.TOKEN_REPO_VALUE {
			currentRepo = Repo{Url: token.Value}
			currentProject.Repos = append(currentProject.Repos,
				currentRepo)
			results[currentProject.Path] = currentProject
		}

		if token.Kind == tokens.TOKEN_EOF {
			break
		}
	}

	close(l.Items)
	return results
}

func findOrCreateProject(search string, m map[string]Project) Project {
	if p, ok := m[search]; ok {
		return p
	}

	// not found, so initialize a new project and add it to the map
	m[search] = Project{Path: search}
	return m[search]
}
