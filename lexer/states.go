package lexer

import "strings"
import "unicode"

import "github.com/mfinelli/wmuc/tokens"

// represent the state as a function that returns the next state
type stateFunc func(*lexer) stateFunc

func lexGeneric(l *lexer) stateFunc {
	for {
		if strings.HasPrefix(strings.ToUpper(l.input[l.pos:]),
			tokens.KEYWORD_REPO) {
			return lexRepo
		}

		switch r := l.next(); {
		case unicode.IsSpace(r):
			l.ignore()
		case r == tokens.COMMENT:
			return lexComment(l, lexGeneric)
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
	return lexQuote(l, lexRepoValueSingleQuote, lexRepoValueDoubleQuote)
}

func lexRepoValueDoubleQuote(l *lexer) stateFunc {
	return lexStr(l, tokens.DOUBLE_QUOTE, tokens.TOKEN_REPO_VALUE,
		lexRepoEndDoubleQuote)
}

func lexRepoValueSingleQuote(l *lexer) stateFunc {
	return lexStr(l, tokens.DOUBLE_QUOTE, tokens.TOKEN_REPO_VALUE,
		lexRepoEndSingleQuote)
}

func lexRepoEndDoubleQuote(l *lexer) stateFunc {
	return lexEndDoubleQuote(l, lexGeneric)
}

func lexRepoEndSingleQuote(l *lexer) stateFunc {
	return lexEndSingleQuote(l, lexGeneric)
}
