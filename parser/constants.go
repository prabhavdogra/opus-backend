package parser

import "opus-backend/token"

// Like an enum, interrelated constants
// LOWEST = 1, EQUALS = 2, ...., CALL = 7
// This helps operators with higher precedence to be deeper in the AST than expressions with lower precedence operators
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
	INDEX       // arr[index]
)

var precedences = map[token.TokenType]int{
	token.EQUAL:       EQUALS,
	token.NOT_EQUAL:   EQUALS,
	token.LESSERTHAN:  LESSGREATER,
	token.GREATERTHAN: LESSGREATER,
	token.PLUS:        SUM,
	token.MINUS:       SUM,
	token.SLASH:       PRODUCT,
	token.ASTERISK:    PRODUCT,
	token.LPAREN:      CALL,
	token.LBRACKET:    INDEX,
}
