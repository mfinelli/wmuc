package lexer_test

import "reflect"
import "testing"
import "github.com/mfinelli/wmuc/lexer"
import "github.com/mfinelli/wmuc/tokens"

func TestLex(t *testing.T) {
	tests := []struct {
		input string
		exp   []tokens.Token
	}{
		{"# only a comment", []tokens.Token{tokens.Token{tokens.TOKEN_EOF, ""}}},
	}

	for _, test := range tests {
		l := lexer.Lex(test.input)
		results := make([]tokens.Token, 0)

		for {
			tok := l.NextToken()
			results = append(results, tok)

			if tok.Kind == tokens.TOKEN_EOF {
				break
			}
		}

		if !reflect.DeepEqual(test.exp, results) {
			t.Errorf("%s didn't return the expected tokens",
				test.input)
		}
	}
}
