package ast

type Expr interface {

}

type BinaryExpr struct {
	Lhs Expr
	Op 	Token
	Rhs Expr
}

// stub
