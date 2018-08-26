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

	inProject := false
	optionsAllowed := false
	allowedTokens := []tokens.TokenType{tokens.TOKEN_PROJECT,
		tokens.TOKEN_REPO, tokens.TOKEN_EOF}

	for {
		token := l.NextToken()

		if token.Kind == tokens.TOKEN_ERROR {
			fmt.Println(token.Value)
			os.Exit(1)
		}

		if !allowedToken(token.Kind, allowedTokens) {
			fmt.Printf("unexpected %q\n", token.Value)
			os.Exit(1)
		}

		if optionsAllowed && token.Kind == tokens.TOKEN_COMMA {
			optionsAllowed = false
			allowedTokens = []tokens.TokenType{tokens.TOKEN_BRANCH}
			continue
		}

		if !optionsAllowed && token.Kind == tokens.TOKEN_COMMA {
			fmt.Println("unexpected comma")
			os.Exit(1)
		}

		if token.Kind == tokens.TOKEN_PROJECT {
			inProject = true
			consumeQuote(l.NextToken())
			allowedTokens = []tokens.TokenType{
				tokens.TOKEN_PROJECT_VALUE}
		}

		if token.Kind == tokens.TOKEN_PROJECT_VALUE {
			currentProject = findOrCreateProject(token.Value,
				results)
			consumeQuote(l.NextToken())
			allowedTokens = []tokens.TokenType{tokens.TOKEN_DO}
		}

		if token.Kind == tokens.TOKEN_DO {
			allowedTokens = []tokens.TokenType{tokens.TOKEN_REPO,
				tokens.TOKEN_END}
		}

		if token.Kind == tokens.TOKEN_END {
			currentProject = findOrCreateProject("", results)
			inProject = false
			allowedTokens = []tokens.TokenType{tokens.TOKEN_PROJECT,
				tokens.TOKEN_REPO, tokens.TOKEN_EOF}
		}

		if token.Kind == tokens.TOKEN_REPO {
			consumeQuote(l.NextToken())
			allowedTokens = []tokens.TokenType{
				tokens.TOKEN_REPO_VALUE}
		}

		if token.Kind == tokens.TOKEN_REPO_VALUE {
			currentRepo = Repo{Url: token.Value}
			currentProject.Repos = append(currentProject.Repos,
				currentRepo)
			results[currentProject.Path] = currentProject
			consumeQuote(l.NextToken())

			optionsAllowed = true

			if inProject == true {
				allowedTokens = []tokens.TokenType{
					tokens.TOKEN_REPO, tokens.TOKEN_END,
					tokens.TOKEN_COMMA}
			} else {
				allowedTokens = []tokens.TokenType{
					tokens.TOKEN_REPO, tokens.TOKEN_PROJECT,
					tokens.TOKEN_COMMA, tokens.TOKEN_EOF}
			}
		}

		if token.Kind == tokens.TOKEN_BRANCH {
			consumeQuote(l.NextToken())
			allowedTokens = []tokens.TokenType{
				tokens.TOKEN_BRANCH_VALUE}
		}

		if token.Kind == tokens.TOKEN_BRANCH_VALUE {
			// remove the current repo, modify the branch and then
			// add it back... there's almost certainly a better way
			for i, r := range currentProject.Repos {
				if r.Url == currentRepo.Url {
					currentProject.Repos = append(
						currentProject.Repos[:i],
						currentProject.Repos[i+1:]...)
					break
				}
			}

			currentRepo.Branch = token.Value
			currentProject.Repos = append(currentProject.Repos,
				currentRepo)
			results[currentProject.Path] = currentProject
			consumeQuote(l.NextToken())

			fmt.Println(currentRepo)

			optionsAllowed = true

			if inProject == true {
				allowedTokens = []tokens.TokenType{
					tokens.TOKEN_REPO, tokens.TOKEN_END,
					tokens.TOKEN_COMMA}
			} else {
				allowedTokens = []tokens.TokenType{
					tokens.TOKEN_REPO, tokens.TOKEN_PROJECT,
					tokens.TOKEN_COMMA, tokens.TOKEN_EOF}
			}
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

func allowedToken(token tokens.TokenType, allowed []tokens.TokenType) bool {
	for _, a := range allowed {
		if token == a {
			return true
		}
	}
	return false
}

func consumeQuote(nextToken tokens.Token) {
	if nextToken.Kind != tokens.TOKEN_SINGLE_QUOTE &&
		nextToken.Kind != tokens.TOKEN_DOUBLE_QUOTE {
		fmt.Println("expecting quote")
		os.Exit(1)
	}
}
