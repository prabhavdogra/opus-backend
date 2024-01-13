package parser

import (
	"opus-backend/lexer"
	"opus-backend/token"
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token
}
