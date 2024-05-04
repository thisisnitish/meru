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

// The basic structure of the NextToken() method, We look at the current character
// under examination (l.ch) and return a token depending on which character it is.
// Before returning the token we advance our pointers into the input so when we call NextToken()
// again the l.ch field is already updated. A small function called newToken helps us with initializing these tokens.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	// TODO: Add float token type tokenization

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
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
