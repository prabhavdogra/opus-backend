package ast

import "opus-backend/token"

// Common methods
type Node interface {
	TokenLiteral() string
	String() string
}

type Expression interface {
	Node
	expressionNode()
}

type Statement interface {
	Node
	statementNode()
}

type Program struct {
	Statements []Statement
}

type Identifier struct {
	Token token.Token // Token: token.IDENT
	Value string
}

type LetStatement struct {
	Token token.Token // Token: token.LET
	Name  *Identifier
	Value Expression
}

type ExpressionStatement struct {
	Token      token.Token // Token: First token in the expression
	Expression Expression
}

type ReturnStatement struct {
	Token       token.Token // Token: token.RETURN
	ReturnValue Expression
}
