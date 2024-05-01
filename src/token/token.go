package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // illegal token that don't know
	EOF     = "EOF"     // end of the file

	// Identifiers + literals
	IDENTIFIER = "IDENTIFIER" // add, foobar, x, y, z...
	INT        = "INT"        // 1,2,3,4,5...

	// Operators
	ASSIGNMENT = "="
	PLUS       = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LEFTPARENTHESIS  = "("
	RIGHTPARENTHESIS = ")"
	LEFTBRACKET      = "{"
	RIGHTBRACKET     = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
