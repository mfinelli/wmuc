package parser

import "github.com/mfinelli/wmuc/lexer"
import "github.com/mfinelli/wmuc/tokens"

type Project struct {
	Path  string
	Repos []Repo
}

type Repo struct {
	Url    string
	Branch string
}

func Parse(input string) map[string]Project {
	l := lexer.Lex("test", input)
	results := make(map[string]Project)

	currentProject := findOrCreateProject("", results)
	var currentRepo Repo

	for {
		token := l.NextToken()

		if token.Kind == tokens.TOKEN_PROJECT_VALUE {
			currentProject = findOrCreateProject(token.Value,
				results)
		}

		if token.Kind == tokens.TOKEN_REPO_VALUE {
			currentRepo = Repo{Url: token.Value}
			currentProject.Repos = append(currentProject.Repos,
				currentRepo)
		}

		if token.Kind == tokens.TOKEN_EOF {
			break
		}
	}

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
