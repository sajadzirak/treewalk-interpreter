package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers (variable names, ...)
	IDENT = "IDENT"

	// Literals
	INT = "INT"

	// Operators
	PLUS   = "+"
	ASSIGN = "="

	// Delimiters
	SEMICOLON = ";"
	COMMA     = ","

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"
)

type Token struct {
	Type    TokenType
	Literal string
}
