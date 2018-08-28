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

import "strings"
import "unicode"

import "github.com/mfinelli/wmuc/tokens"

func lexComment(l *lexer, next stateFunc) stateFunc {
	for {
		switch r := l.next(); {
		case r == tokens.NEWLINE:
			return next
		case r == tokens.EOF:
			l.emit(tokens.TOKEN_EOF)
			return nil
		default:
			l.ignore()
		}
	}
}

func lexStr(l *lexer, q rune, tt tokens.TokenType, next stateFunc) stateFunc {
	skipNext := false

	for {
		// skip this character because we detected an escape sequence
		if skipNext {
			skipNext = false
			l.next()
			continue
		}

		switch r := l.next(); {
		case r == tokens.EOF:
			return l.errorf("unexpected end of input")
		case r == tokens.NEWLINE:
			return l.errorf("unexpected newline")
		case r == tokens.ESCAPE:
			skipNext = true
		case r == q:
			l.backup()
			l.emitWithEscapes(tt, string(q))
			return next
		}
	}
}

func lexQuote(l *lexer, nextS, nextD stateFunc) stateFunc {
	for {
		if strings.HasPrefix(l.input[l.pos:],
			tokens.KEYWORD_SINGLE_QUOTE) {
			l.pos += len(tokens.KEYWORD_SINGLE_QUOTE)
			l.emit(tokens.TOKEN_SINGLE_QUOTE)
			return nextS
		}

		if strings.HasPrefix(l.input[l.pos:],
			tokens.KEYWORD_DOUBLE_QUOTE) {
			l.pos += len(tokens.KEYWORD_DOUBLE_QUOTE)
			l.emit(tokens.TOKEN_DOUBLE_QUOTE)
			return nextD
		}

		switch r := l.next(); {
		case r == tokens.EOF:
			return l.errorf("unexpected end of input")
		case unicode.IsSpace(r):
			l.ignore()
		}
	}
}

func lexEndDoubleQuote(l *lexer, next stateFunc) stateFunc {
	l.pos += len(tokens.KEYWORD_DOUBLE_QUOTE)
	l.emit(tokens.TOKEN_DOUBLE_QUOTE)
	return next
}

func lexEndSingleQuote(l *lexer, next stateFunc) stateFunc {
	l.pos += len(tokens.KEYWORD_SINGLE_QUOTE)
	l.emit(tokens.TOKEN_SINGLE_QUOTE)
	return next
}
