package lexer

import "strings"
import "unicode"

import "github.com/mfinelli/wmuc/tokens"

// represent the state as a function that returns the next state
type stateFunc func(*lexer) stateFunc

func lexUnkown(l *lexer) stateFunc {
	for {
		if strings.HasPrefix(strings.ToUpper(l.input[l.pos:]),
			tokens.KEYWORD_REPO) {
			return lexRepo
		}

		switch r := l.next(); {
		case unicode.IsSpace(r):
			l.ignore()
		case r == tokens.EOF:
			l.emit(tokens.TOKEN_EOF)
			return nil
		}
	}
}

func lexRepo(l *lexer) stateFunc {
	l.pos += len(tokens.KEYWORD_REPO)
	l.emit(tokens.TOKEN_REPO)
	return lexRepoBeginQuote
}

func lexRepoBeginQuote(l *lexer) stateFunc {
	for {
		if strings.HasPrefix(l.input[l.pos:],
			tokens.KEYWORD_SINGLE_QUOTE) {
			l.emit(tokens.TOKEN_SINGLE_QUOTE)
			// return lexRepoValueSingleQuote
			return lexUnkown
		}

		if strings.HasPrefix(l.input[l.pos:],
			tokens.KEYWORD_DOUBLE_QUOTE) {
			l.emit(tokens.TOKEN_DOUBLE_QUOTE)
			// return lexRepoValueDoubleQuote
			return lexUnkown
		}

		switch r := l.next(); {
		case r == tokens.EOF:
			return l.errorf("unexpected end of input")
		case r == rune('\t') || r == rune(' ') || r == rune('\n'):
			l.ignore()
		}
	}
}

