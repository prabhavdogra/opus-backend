package parser

import (
	"opus-backend/ast"
	"opus-backend/lexer"
	"opus-backend/token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	// Used for parsing expression
	// These maps help check if the appropriate map has a parsing function associated with curToken.Type
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}
