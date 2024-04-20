package lexer

import (
	"testing"

	"github.com/ECecillo/redis-interpreter/token"
)

type TokenOutputExpected struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func checkToken(input string, tests []TokenOutputExpected, t *testing.T) {
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		// fmt.Println(tok)
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
	}
}

func TestOperators(t *testing.T) {
	input := "+-:$*_#,(!=%~>"

	tests := []TokenOutputExpected{
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.COLON, ":"},
		{token.DOLLAR, "$"},
		{token.ASTERISK, "*"},
		{token.UNDERSORE, "_"},
		{token.HASH, "#"},
		{token.COMA, ","},
		{token.LPAREN, "("},
		{token.BANG, "!"},
		{token.ASSIGN, "="},
		{token.PRCT, "%"},
		{token.TILDE, "~"},
		{token.GT, ">"},
	}

	checkToken(input, tests, t)
}

func TestSkipCRLF(t *testing.T) {
	input := "\r\n+\r\n"

	tests := []TokenOutputExpected{
		{token.CRLF, "\r\n"},
		{token.PLUS, "+"},
		{token.CRLF, "\r\n"},
	}
	checkToken(input, tests, t)
}

func TestPing(t *testing.T) {
	input := "*1\r\n$4\r\nping\r\n"

	tests := []TokenOutputExpected{
		{token.ASTERISK, "*"},
		{token.INT, "1"},
		{token.CRLF, "\r\n"},
		{token.DOLLAR, "$"},
		{token.INT, "4"},
		{token.CRLF, "\r\n"},
		{token.PING, "ping"},
		{token.CRLF, "\r\n"},
	}
	checkToken(input, tests, t)
}
