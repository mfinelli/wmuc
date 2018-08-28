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
		{tokens.Token{Kind: tokens.TOKEN_EOF,
			Value: fmt.Sprintf("%c", eof)}, "EOF"},
		{tokens.Token{Kind: tokens.TOKEN_ERROR, Value: errorText},
			errorText},
		{tokens.Token{Kind: tokens.TOKEN_PROJECT, Value: shortStr},
			shortStrExp},
		{tokens.Token{Kind: tokens.TOKEN_REPO_VALUE, Value: longStr},
			longStrExp},
	}

	for _, test := range tests {
		if test.tok.String() != test.exp {
			t.Errorf("String for %s didn't match: %s (got: %s)",
				test.tok.Value, test.exp, test.tok.String())
		}
	}
}
