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

package lexer

import "fmt"
import "strings"
import "unicode/utf8"

import "github.com/spf13/viper"

import "github.com/mfinelli/wmuc/tokens"

type lexer struct {
	input string
	start int
	pos   int
	width int
	state stateFunc
	Items chan tokens.Token
}

func Lex(input string) *lexer {
	l := &lexer{
		input: input,
		state: lexGeneric,
		Items: make(chan tokens.Token, 2),
	}

	return l
}

func (l *lexer) run() {
	for state := lexGeneric; state != nil; {
		state = state(l)
	}

	close(l.Items)
}

func (l *lexer) emit(t tokens.TokenType) {
	if viper.GetBool("debug") {
		fmt.Printf("emitted %q\n", l.input[l.start:l.pos])
	}

	l.Items <- tokens.Token{Kind: t, Value: l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) emitWithEscapes(t tokens.TokenType, escape string) {
	escaped := strings.Replace(l.input[l.start:l.pos],
		fmt.Sprintf("\\%s", escape), escape, -1)

	if viper.GetBool("debug") {
		fmt.Printf("emitted %q\n", escaped)
	}

	l.Items <- tokens.Token{Kind: t, Value: escaped}
	l.start = l.pos
}

func (l *lexer) errorf(f string, args ...interface{}) stateFunc {
	l.Items <- tokens.Token{Kind: tokens.TOKEN_ERROR,
		Value: fmt.Sprintf(f, args...)}
	return nil
}

func (l *lexer) NextToken() tokens.Token {
	for {
		select {
		case tok := <-l.Items:
			return tok
		default:
			l.state = l.state(l)
		}
	}
}

// return the next rune in the input
func (l *lexer) next() rune {
	if l.pos >= len(l.input) {
		l.width = 0
		return tokens.EOF
	}

	result, width := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = width
	l.pos += l.width
	return result
}

// skips any consumed input
func (l *lexer) ignore() {
	l.start = l.pos
}

// step back one rune, can only be called once per call of next()
func (l *lexer) backup() {
	l.pos -= l.width
}

// return the next rune without consuming it
func (l *lexer) peek() rune {
	rune := l.next()
	l.backup()
	return rune
}
