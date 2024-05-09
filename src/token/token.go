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
	MINUS      = "-"
	BANG       = "!"
	ASTERISK   = "*"
	SLASH      = "/"

	LESS_THAN    = "<"
	GREATER_THAN = ">"

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

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// To check whether the given Identifier is keyword or not
func LookUpIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}

	return IDENTIFIER
}
