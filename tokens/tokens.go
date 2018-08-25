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

// keywords
const KEYWORD_PROJECT = "PROJECT"
const KEYWORD_REPO = "REPO"
const KEYWORD_BRANCH = "BRANCH"
const KEYWORD_SINGLE_QUOTE = "'"
const KEYWORD_DOUBLE_QUOTE = "\""

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
