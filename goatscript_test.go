package goatscript

import (
	"fmt"
	// This path works for me but cannot possibly work for anyone else. Hmm.
	goatdata "github.com/AdamKopczenski/goatscript/ast"
)

// https://blog.golang.org/examples

func Example() {
	script := goatdata.IfStmt{
		Condition: goatdata.BinaryExpr{
				Lhs: goatdata.IntVariable{ "var1" },
				Op: goatdata.Token{ "<" },
				Rhs: goatdata.Int{ 2 },
			},
		Then: goatdata.IntIncrement{
				Variable: "var1",
				Amount: 1,
			},
		Else: goatdata.NoOp{},
	}

	templ, err := Precompile(script)
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
