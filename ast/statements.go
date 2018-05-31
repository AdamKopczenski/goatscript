package ast

type Stmt interface {

}

type IfStmt struct {
	Condition  Expr
	Then       []Stmt
	Else       []Stmt
}

// stub
