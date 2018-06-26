package template

import (
	ast "github.com/AdamKopczenski/goatscript/ast"
	instance "github.com/AdamKopczenski/goatscript/instance"
)

type Template struct {
}

func Precompile(scr ast.IfStmt) (Template, error) {
  return *new(Template), nil
}

func (templ Template) CreateInst() instance.Instance {
  return *new(instance.Instance)
}
