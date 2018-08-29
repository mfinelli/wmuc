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
		{`project "wmuc" do
			repo "https://github.com/mfinelli/wmuc.git",
				branch: "dev"
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
			tokens.Token{tokens.TOKEN_COMMA, ","},
			tokens.Token{tokens.TOKEN_BRANCH, "branch:"},
			tokens.Token{tokens.TOKEN_DOUBLE_QUOTE, "\""},
			tokens.Token{tokens.TOKEN_BRANCH_VALUE, "dev"},
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
