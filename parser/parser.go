package parser

import "github.com/mfinelli/wmuc/lexer"
import "github.com/mfinelli/wmuc/tokens"

func Parse(input string) []string {
	l := lexer.Lex("test", input)
	var results []string

	for {
		token := l.NextToken()
		results = append(results, token.Value)

		if token.Kind == tokens.TOKEN_EOF {
			break
		}
	}

	return results
}
