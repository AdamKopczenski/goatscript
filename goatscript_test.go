package goatscript

import (
	"fmt"
	// This path works for me but cannot possibly work for anyone else. Hmm.
	ast "github.com/AdamKopczenski/goatscript/ast"
	template "github.com/AdamKopczenski/goatscript/template"
	// "Imported and not used" feels like a wrong description of this. Hmm.
	//instance "github.com/AdamKopczenski/goatscript/instance"
)

// https://blog.golang.org/examples

func Example() {
	// if var1 < 2 {
	// 	++var1;
	// } else {
	// }
	script := ast.IfStmt{
		Condition: ast.BinaryExpr{
				Lhs: ast.IntVariable{ ast.Variable{"var1"} },
				Op: ast.Token(ast.LessThan),
				Rhs: ast.IntLiteral{2},
			},
		Then: []ast.Stmt{ ast.IncrementStmt{
				Operand: ast.Variable{"var1"} ,
			} },
		Else: []ast.Stmt{ ast.NoOpStmt{} },
	}

	templ, err := template.Precompile(script)
	if err != nil {
		panic(err)
	}

	inst1 := templ.CreateInst()
	inst1.SetIntVariable("var1", 1)
	inst1.Run()
	var1, _ := inst1.GetIntVariable("var1")
	fmt.Println(var1)

	inst2 := templ.CreateInst()
	inst2.SetIntVariable("var1", 6)
	inst2.Run()
	var1again, _ := inst2.GetIntVariable("var1")
	fmt.Println(var1again)

	// Output:
	// 2
	// 6
}
