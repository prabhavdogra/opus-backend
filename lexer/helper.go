package lexer

import (
	"opus/token"
	"opus/utils.go"
)

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Sets l.ch as the current character and increments
func (l *Lexer) readChar() {
	if l.readIndex >= len(l.input) {
		l.ch = 0 // ASCII for NULL character
	} else {
		l.ch = l.input[l.readIndex]
	}
	l.index = l.readIndex
	l.readIndex += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespaces()

	switch l.ch {
	// Handle case for = and ==
	case '=':
		if l.peekNextChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQUAL, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekNextChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQUAL, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LESSERTHAN, l.ch)
	case '>':
		tok = newToken(token.GREATERTHAN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if utils.IsLetter(l.ch) {
			tok = parseStringIdentifierToken(l)
			return tok
		} else if utils.IsDigit(l.ch) {
			tok = parseDigitIdentifierToken(l)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) peekNextChar() byte {
	if l.readIndex >= len(l.input) {
		return 0
	} else {
		return l.input[l.readIndex]
	}
}

func parseStringIdentifierToken(l *Lexer) token.Token {
	var tok token.Token
	tok.Literal = l.readIdentifier()
	tok.Type = token.LookupIdent(tok.Literal)
	return tok
}

func parseDigitIdentifierToken(l *Lexer) token.Token {
	var tok token.Token
	tok.Type = token.INT
	tok.Literal = l.readNumber()
	return tok
}

func (l *Lexer) skipWhitespaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	startIndex := l.index
	for utils.IsLetter(l.ch) {
		l.readChar()
	}
	return l.input[startIndex:l.index]
}

func (l *Lexer) readNumber() string {
	index := l.index
	for utils.IsDigit(l.ch) {
		l.readChar()
	}
	return l.input[index:l.index]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
