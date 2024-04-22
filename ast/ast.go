package ast

import "interpreter-go/token"

type Node interface {
	TokenLiteral() string
}

// 표현식
type Statement interface {
	Node
	statementNode()
}

// Expression : 명령문
type Expression interface {
	Node
	expressionNode()
}

// Program : AST의 루트 노드 (=명령문 집합)
type Program struct {
	Statements []Statement
}

func (program *Program) TokenLiteral() string {
	if len(program.Statements) == 0 {
		return ""
	}

	return program.Statements[0].TokenLiteral()
}

// LetStatement : LET 표현식
type LetStatement struct {
	Token token.Token // token.LET 토큰
	Name  *Identifier // 변수명
	Value Expression  // 명령문
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier : 식별자(=변수명) 토큰
type Identifier struct {
	Token token.Token // token.IDENT 토큰
	Value string      // 변수명
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
// ReturnStatement : return 표현식
type ReturnStatement struct {
	Token       token.Token // token.RETURN 토큰
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
