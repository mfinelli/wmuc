package "lexer"

import "strings"

import "github.com/mfinelli/tokens"

// represent the state as a function that returns the next state
type stateFunc func(*lexer) stateFunc

func lexUnkown(l *lexer) stateFunc {
	for {
		if strings.HasPrefix(strings.ToUpper(l.input[l.pos:]),
			tokens.KEYWORD_REPO) {
			return lexRepo
		}

		// check and ignore whitespace: unicode.IsSpace

		if l.next() == tokens.EOF {
			break
		}
	}

	// we've reached the end of input, emit the last item and then EOF
	// if l.pos > l.start { l.emit(TODO) }
	l.emit(tokens.TOKEN_EOF)
	return nil
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
			return lexRepoValueSingleQuote
		}

		if strings.HasPrefix(l.input[l.pos:],
			tokens.KEYWORD_DOUBLE_QUOTE) {
			l.emit(tokens.TOKEN_DOUBLE_QUOTE)
			return lexRepoValueDoubleQuote
		}

		switch r := l.next(); {
		case r == tokens.EOF:
			return l.errorf("unexpected end of input")
		case r == "\t" || r == " " || r == "\n":
			l.ignore()
		}
	}
}

