package tokens

import "fmt"

type tokenType int

type token struct {
	kind  tokenType
	value string
}

const (
	TOKEN_ERROR tokenType = iota
	TOKEN_EOF
)

const EOF rune = 0
const NEWLINE string = "\n"

func (t token) String() string {
	switch t.kind {
	case TOKEN_EOF:
		return "EOF"

	case TOKEN_ERROR:
		return t.value
	}

	if len(t.value) > 15 {
		return fmt.Sprintf("%.15q...", t.value)
	}

	return fmt.Sprintf("%q", t.value)
}
