package lexer

import "meru/src/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

// The purpose of this function is to give the next character and advance our position in the input string
// l.position always points to the position where we last read
// l.readPosition always points to the next position where we're going to read
func (l *Lexer) readChar() {

	// If we reach the end of the file or haven't read anything
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition] // read through the next character
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGNMENT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LEFTPARENTHESIS, l.ch)
	case ')':
		tok = newToken(token.RIGHTPARENTHESIS, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LEFTBRACKET, l.ch)
	case '}':
		tok = newToken(token.RIGHTBRACKET, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
