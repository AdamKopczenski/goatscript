package ast

type Stmt interface {

}

type IfStmt struct {
	Condition  Expr
	Then       []Stmt
	Else       []Stmt
}

type IncrementStmt struct {
	Operand Variable
}

type NoOpStmt struct {
}
