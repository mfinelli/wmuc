package lexer

import "strings"
import "unicode"

import "github.com/mfinelli/wmuc/tokens"

// represent the state as a function that returns the next state
type stateFunc func(*lexer) stateFunc

func lexGeneric(l *lexer) stateFunc {
	for {
		if strings.HasPrefix(strings.ToUpper(l.input[l.pos:]),
			tokens.KEYWORD_PROJECT) {
			return lexProject
		}

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

func lexProject(l *lexer) stateFunc {
	l.pos += len(tokens.KEYWORD_PROJECT)
	l.emit(tokens.TOKEN_PROJECT)
	return lexProjectBeginQuote
}

func lexProjectBeginQuote(l *lexer) stateFunc {
	return lexQuote(l, lexProjectValueSingleQuote,
		lexProjectValueDoubleQuote)
}

func lexProjectValueDoubleQuote(l *lexer) stateFunc {
	return lexStr(l, tokens.DOUBLE_QUOTE, tokens.TOKEN_PROJECT_VALUE,
		lexProjectEndDoubleQuote)
}

func lexProjectValueSingleQuote(l *lexer) stateFunc {
	return lexStr(l, tokens.DOUBLE_QUOTE, tokens.TOKEN_PROJECT_VALUE,
		lexProjectEndSingleQuote)
}

func lexProjectEndDoubleQuote(l *lexer) stateFunc {
	return lexEndDoubleQuote(l, lexProjectDo)
}

func lexProjectEndSingleQuote(l *lexer) stateFunc {
	return lexEndSingleQuote(l, lexProjectDo)
}

func lexProjectDo(l *lexer) stateFunc {
	for {
		if strings.HasPrefix(strings.ToUpper(l.input[l.pos:]),
			tokens.KEYWORD_DO) {
			l.pos += len(tokens.KEYWORD_DO)
			l.emit(tokens.TOKEN_DO)
			return lexProjectContext
		}

		switch r := l.next(); {
		case unicode.IsSpace(r):
			l.ignore()
		case r == tokens.EOF:
			return l.errorf("unexpected end of input")
		}
	}
}

func lexProjectContext(l *lexer) stateFunc {
	for {
		if strings.HasPrefix(strings.ToUpper(l.input[l.pos:]),
			tokens.KEYWORD_REPO) {
			return lexProjectRepo
		}

		if strings.HasPrefix(strings.ToUpper(l.input[l.pos:]),
			tokens.KEYWORD_END) {
			return lexProjectEnd
		}

		switch r := l.next(); {
		case unicode.IsSpace(r):
			l.ignore()
		case r == tokens.COMMENT:
			return lexComment(l, lexProjectContext)
		case r == tokens.EOF:
			return l.errorf("unexpected end of input")
		}
	}
}

func lexProjectEnd(l *lexer) stateFunc {
	l.pos += len(tokens.KEYWORD_END)
	l.emit(tokens.TOKEN_END)
	return lexGeneric
}

func lexProjectRepo(l *lexer) stateFunc {
	l.pos += len(tokens.KEYWORD_REPO)
	l.emit(tokens.TOKEN_REPO)
	return lexProjectRepoBeginQuote
}

func lexProjectRepoBeginQuote(l *lexer) stateFunc {
	return lexQuote(l, lexProjectRepoValueSingleQuote,
		lexProjectRepoValueDoubleQuote)
}

func lexProjectRepoValueDoubleQuote(l *lexer) stateFunc {
	return lexStr(l, tokens.DOUBLE_QUOTE, tokens.TOKEN_REPO_VALUE,
		lexProjectRepoEndDoubleQuote)
}

func lexProjectRepoValueSingleQuote(l *lexer) stateFunc {
	return lexStr(l, tokens.DOUBLE_QUOTE, tokens.TOKEN_REPO_VALUE,
		lexProjectRepoEndSingleQuote)
}

func lexProjectRepoEndDoubleQuote(l *lexer) stateFunc {
	return lexEndDoubleQuote(l, lexProjectContext)
}

func lexProjectRepoEndSingleQuote(l *lexer) stateFunc {
	return lexEndSingleQuote(l, lexProjectContext)
}
