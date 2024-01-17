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

type IntegerLiteral struct {
	Token token.Token // Token: token.INT
	Value int64
}

type PrefixExpression struct {
	Token      token.Token // Token: "!" / "-"
	Operator   string
	Expression Expression
}

type InfixExpression struct {
	Token           token.Token // Token: "*", "/", "+", "-" ..
	Operator        string
	RightExpression Expression
	LeftExpression  Expression
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

type BlockStatement struct {
	Token      token.Token // Token: token.LBRACE {
	Statements []Statement
}

type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

type StringLiteral struct {
	Token token.Token
	Value string
}

type ArrayLiteral struct {
	Token    token.Token // the '[' token
	Elements []Expression
}

type IndexExpression struct {
	Token token.Token // The [ token
	Left  Expression
	Index Expression
}

type HashLiteral struct {
	Token token.Token // the '{' token
	Pairs map[Expression]Expression
}

type Boolean struct {
	Token token.Token
	Value bool
}

type IfExpression struct {
	Token       token.Token // The 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

type FunctionLiteral struct {
	Token      token.Token // The 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}
