package token

import "strings"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"HELLO": HELLO,
	"PING":  PING,
	"PONG":  PONG,
	"SET":   SET,
	"GET":   GET,
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"   // 1343456
	FLOAT  = "FLOAT"
	STRING = "STRING"
	CRLF   = "CRLF"

	// Operators
	PLUS      = "+"
	MINUS     = "-"
	COLON     = ":"
	DOLLAR    = "$"
	ASTERISK  = "*"
	UNDERSORE = "_"
	HASH      = "#"
	COMA      = ","
	LPAREN    = "("
	BANG      = "!"
	ASSIGN    = "="
	PRCT      = "%"
	TILDE     = "~"
	GT        = ">"

	BACKSLASH = "\\"

	// Commands
	HELLO = "HELLO"
	PING  = "PING"
	PONG  = "PONG"
	SET   = "SET"
	GET   = "GET"
)

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[strings.ToUpper(ident)]; ok {
		return tok
	}
	return IDENT
}
