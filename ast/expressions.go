package ast

import "strconv"

type Expr interface {

}

type BinaryExpr struct {
	Lhs Expr
	Op 	Token
	Rhs Expr
}

type Variable struct {
	Name string
}

type IntVariable struct {
	Variable
}

type Literal interface {
	text() string
}

type IntLiteral struct {
	Value int
}

func(i IntLiteral) text() string {
	return strconv.Itoa(i.Value)
}
// stub
