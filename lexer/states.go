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
		case r == tokens.COMMENT:
			return lexCommentFromUnknown
		case r == tokens.EOF:
			l.emit(tokens.TOKEN_EOF)
			return nil
		}
	}
}

func lexCommentFromUnknown(l *lexer) stateFunc {
	for {
		switch r := l.next(); {
		case r == tokens.NEWLINE:
			return lexUnkown
		default:
			l.ignore()
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
			l.pos += len(tokens.KEYWORD_SINGLE_QUOTE)
			l.emit(tokens.TOKEN_SINGLE_QUOTE)
			return lexRepoValueSingleQuote
		}

		if strings.HasPrefix(l.input[l.pos:],
			tokens.KEYWORD_DOUBLE_QUOTE) {
			l.pos += len(tokens.KEYWORD_DOUBLE_QUOTE)
			l.emit(tokens.TOKEN_DOUBLE_QUOTE)
			return lexRepoValueDoubleQuote
		}

		switch r := l.next(); {
		case r == tokens.EOF:
			return l.errorf("unexpected end of input")
		case r == rune('\t') || r == rune(' ') || r == rune('\n'):
			l.ignore()
		}
	}
}

func lexRepoValueDoubleQuote(l *lexer) stateFunc {
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
		case r == tokens.DOUBLE_QUOTE:
			l.backup()
			l.emitWithEscapes(tokens.TOKEN_REPO_VALUE, "\"")
			return lexRepoEndDoubleQuote
		}
	}
}

func lexRepoValueSingleQuote(l *lexer) stateFunc {
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
		case r == tokens.SINGLE_QUOTE:
			l.backup()
			l.emitWithEscapes(tokens.TOKEN_REPO_VALUE, "'")
			return lexRepoEndSingleQuote
		}
	}
}

func lexRepoEndDoubleQuote(l *lexer) stateFunc {
	l.pos += len(tokens.KEYWORD_DOUBLE_QUOTE)
	l.emit(tokens.TOKEN_DOUBLE_QUOTE)
	return lexUnkown
}

func lexRepoEndSingleQuote(l *lexer) stateFunc {
	l.pos += len(tokens.KEYWORD_SINGLE_QUOTE)
	l.emit(tokens.TOKEN_SINGLE_QUOTE)
	return lexUnkown
}
