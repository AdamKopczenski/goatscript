package ast

type Token int

const (
	Unknown Token = iota
	LessThan
	EqualTo
	GreaterThan
	LessEqual
	GreaterEqual
	NotEqual
)

// stub
