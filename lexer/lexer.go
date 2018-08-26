package lexer

import "fmt"
import "strings"
import "unicode/utf8"

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
		state: lexUnkown,
		Items: make(chan tokens.Token, 2),
	}

	return l
}

func (l *lexer) run() {
	for state := lexUnkown; state != nil; {
		state = state(l)
	}

	close(l.Items)
}

func (l *lexer) emit(t tokens.TokenType) {
	l.Items <- tokens.Token{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) emitWithEscapes(t tokens.TokenType, escape string) {
	escaped := strings.Replace(l.input[l.start:l.pos], fmt.Sprintf("\\%s", escape), escape, -1)
	l.Items <- tokens.Token{t, escaped}
	l.start = l.pos
}

func (l *lexer) errorf(f string, args ...interface{}) stateFunc {
	l.Items <- tokens.Token{tokens.TOKEN_ERROR, fmt.Sprintf(f, args...)}
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

	panic("nextToken reached an invalid state")
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
