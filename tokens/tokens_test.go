package tokens_test

import "fmt"
import "testing"
import "github.com/mfinelli/wmuc/tokens"

func TestString(t *testing.T) {
	eof := rune(0)
	errorText := "there was an error during lexing"
	shortStr := "project"
	shortStrExp := "\"project\""
	longStr := "https://github.com/mfinelli/wmuc"
	longStrExp := "\"https://github.\"..."

	tests := []struct {
		tok tokens.Token
		exp string
	}{
		{tokens.Token{tokens.TOKEN_EOF, fmt.Sprintf("%c", eof)}, "EOF"},
		{tokens.Token{tokens.TOKEN_ERROR, errorText}, errorText},
		{tokens.Token{tokens.TOKEN_PROJECT, shortStr}, shortStrExp},
		{tokens.Token{tokens.TOKEN_REPO_VALUE, longStr}, longStrExp},
	}

	for _, test := range tests {
		if test.tok.String() != test.exp {
			t.Errorf("String for %s didn't match: %s (got: %s)",
				test.tok.Value, test.exp, test.tok.String())
		}
	}
}
