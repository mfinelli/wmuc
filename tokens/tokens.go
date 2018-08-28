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

package tokens

import "fmt"

type TokenType int

type Token struct {
	Kind  TokenType
	Value string
}

const (
	TOKEN_ERROR TokenType = iota
	TOKEN_EOF

	TOKEN_SINGLE_QUOTE
	TOKEN_DOUBLE_QUOTE
	TOKEN_COMMA

	TOKEN_PROJECT
	TOKEN_PROJECT_VALUE
	TOKEN_REPO
	TOKEN_REPO_VALUE
	TOKEN_BRANCH
	TOKEN_BRANCH_VALUE

	TOKEN_DO
	TOKEN_END
)

const EOF rune = 0
const NEWLINE = rune('\n')
const TAB = rune('\t')
const SPACE = rune(' ')
const COMMENT = rune('#')
const ESCAPE = rune('\\')
const SINGLE_QUOTE = rune('\'')
const DOUBLE_QUOTE = rune('"')
const COMMA = rune(',')

// keywords
const KEYWORD_PROJECT = "PROJECT"
const KEYWORD_REPO = "REPO"
const KEYWORD_BRANCH = "BRANCH:"
const KEYWORD_DO = "DO"
const KEYWORD_END = "END"
const KEYWORD_SINGLE_QUOTE = "'"
const KEYWORD_DOUBLE_QUOTE = "\""
const KEYWORD_COMMA = ","

func (t Token) String() string {
	switch t.Kind {
	case TOKEN_EOF:
		return "EOF"

	case TOKEN_ERROR:
		return t.Value
	}

	if len(t.Value) > 15 {
		return fmt.Sprintf("%.15q...", t.Value)
	}

	return fmt.Sprintf("%q", t.Value)
}
