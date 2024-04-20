package lexer

import (
	"github.com/ECecillo/redis-interpreter/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL" if we reach EOF or nothing to read
	} else {
		l.ch = l.input[l.readPosition] // Regarding input put next char in ch
	}
	l.position = l.readPosition // set input position to last read char
	l.readPosition += 1         // Go to next char
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()
	switch l.ch {
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '$':
		tok = newToken(token.DOLLAR, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '_':
		tok = newToken(token.UNDERSORE, l.ch)
	case '#':
		tok = newToken(token.HASH, l.ch)
	case ',':
		tok = newToken(token.COMA, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case '!':
		tok = newToken(token.BANG, l.ch)
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '%':
		tok = newToken(token.PRCT, l.ch)
	case '~':
		tok = newToken(token.TILDE, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '\r':
		if l.peekChar() == '\n' {
			l.readChar()
			tok = token.Token{Type: token.CRLF, Literal: "\r\n"}
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	case '\\':
		if l.peekChar() == 'r' {
			l.readChar()
			if l.peekChar() == '\\' {
				l.readChar()
				if l.peekChar() == 'n' {
					l.readChar()
					tok = token.Token{Type: token.CRLF, Literal: "\\r\\n"}
					return tok
				}
			}
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() token.Token {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	if l.ch == '.' {
		for isDigit(l.ch) {
			l.readChar()
		}
		return token.Token{Type: token.FLOAT, Literal: l.input[position:l.position]}
	}
	return token.Token{Type: token.INT, Literal: l.input[position:l.position]}
}

func (l *Lexer) readFloat() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	if l.ch == '.' {
		for isDigit(l.ch) {
			l.readChar()
		}
	}

	return l.input[position:l.position]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z')
}

func isDigit(ch byte) bool {
	return (ch >= '0' && ch <= '9')
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' {
		l.readChar()
	}
}
