package tokens

import "fmt"

type tokenType int

type Token struct {
	Kind  tokenType
	Value string
}

const (
	TOKEN_ERROR tokenType = iota
	TOKEN_EOF

	TOKEN_VALUE
)

const EOF rune = 0
const NEWLINE string = "\n"

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
