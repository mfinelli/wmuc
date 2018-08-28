package lexer_test

import "reflect"
import "testing"
import "github.com/mfinelli/wmuc/lexer"
import "github.com/mfinelli/wmuc/tokens"

func TestLex(t *testing.T) {
	wmuc_repo := "https://github.com/mfinelli/wmuc.git"

	tests := []struct {
		input string
		exp   []tokens.Token
	}{
		{"# only a comment", []tokens.Token{
			tokens.Token{tokens.TOKEN_EOF, ""}}},
		{`repo "https://github.com/mfinelli/wmuc.git"`,
			[]tokens.Token{
				tokens.Token{tokens.TOKEN_REPO, "repo"},
				tokens.Token{tokens.TOKEN_DOUBLE_QUOTE, "\""},
				tokens.Token{tokens.TOKEN_REPO_VALUE,
					wmuc_repo},
				tokens.Token{tokens.TOKEN_DOUBLE_QUOTE, "\""},
				tokens.Token{tokens.TOKEN_EOF, ""}}},
		{`repo "https://github.com/mfinelli/wmuc.git", branch: "dev"`,
			[]tokens.Token{
				tokens.Token{tokens.TOKEN_REPO, "repo"},
				tokens.Token{tokens.TOKEN_DOUBLE_QUOTE, "\""},
				tokens.Token{tokens.TOKEN_REPO_VALUE,
					wmuc_repo},
				tokens.Token{tokens.TOKEN_DOUBLE_QUOTE, "\""},
				tokens.Token{tokens.TOKEN_COMMA, ","},
				tokens.Token{tokens.TOKEN_BRANCH, "branch:"},
				tokens.Token{tokens.TOKEN_DOUBLE_QUOTE, "\""},
				tokens.Token{tokens.TOKEN_BRANCH_VALUE, "dev"},
				tokens.Token{tokens.TOKEN_DOUBLE_QUOTE, "\""},
				tokens.Token{tokens.TOKEN_EOF, ""}}},
		{`project "wmuc" do
			repo "https://github.com/mfinelli/wmuc.git"
		end`, []tokens.Token{
			tokens.Token{tokens.TOKEN_PROJECT, "project"},
			tokens.Token{tokens.TOKEN_DOUBLE_QUOTE, "\""},
			tokens.Token{tokens.TOKEN_PROJECT_VALUE, "wmuc"},
			tokens.Token{tokens.TOKEN_DOUBLE_QUOTE, "\""},
			tokens.Token{tokens.TOKEN_DO, "do"},
			tokens.Token{tokens.TOKEN_REPO, "repo"},
			tokens.Token{tokens.TOKEN_DOUBLE_QUOTE, "\""},
			tokens.Token{tokens.TOKEN_REPO_VALUE, wmuc_repo},
			tokens.Token{tokens.TOKEN_DOUBLE_QUOTE, "\""},
			tokens.Token{tokens.TOKEN_END, "end"},
			tokens.Token{tokens.TOKEN_EOF, ""}}},
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
