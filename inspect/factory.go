package inspect

import (
	"go/ast"

	"github.com/sky0621/go-testcode-autogen/inspect/decl"
	"github.com/sky0621/go-testcode-autogen/inspect/expr"
	"github.com/sky0621/go-testcode-autogen/inspect/spec"
	"github.com/sky0621/go-testcode-autogen/inspect/stmt"
)

func GetInspector(node ast.Node) Inspector {
	if node == nil {
		return nil
	}
	inspectors := []Inspector{
		&decl.BadDeclInspector{},
		&decl.FuncDeclInspector{},
		&decl.GenDeclInspector{},

		&expr.FuncTypeInspector{},
		&expr.IdentInspector{},

		&spec.ImportSpecInspector{},
		&spec.TypeSpecInspector{},
		&spec.ValueSpecInspector{},

		&stmt.FuncDeclInspector{},

		// FIXME
	}

	for _, i := range inspectors {
		if i.IsTarget(node) {
			return i
		}
	}
	return nil
}
