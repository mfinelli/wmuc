package lexer

import "github.com/mfinelli/wmuc/tokens"

type lexer struct {
	name  string
	input string
	start int
	pos   int
	width int
	state stateFunc
	items chan tokens.Token
}

func Lex(name, input string) *lexer {
	l := &lexer{
		name:  name,
		input: input,
		// state: lexUnkown
		items: make(chan item, 2),
	}

	return l
}

func (l *lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}

	close(l.items)
}

func (l *lexer) emit(t tokens.TokenType) {
	l.items <- tokens.Token{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) nextToken() tokens.Token {
	for {
		select {
		case tok := <-l.items:
			return tok
		default:
			l.state = l.state(l)
		}
	}

	panic("nextToken reached an invalid state")
}

// return the next rune in the input
func (l *lexer) next() (rune int) {
	if l.pos >= len(l.input) {
		l.width = 0
		return tokens.EOF
	}

	rune, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return rune
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
func (l *lexer) peek() int {
	rune := l.next()
	l.backup()
	return rune
}
