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
		// DECL
		&decl.BadDeclInspector{},
		&decl.FuncDeclInspector{},
		&decl.GenDeclInspector{},
		// EXPR
		&expr.ArrayTypeInspector{},
		&expr.BadExprInspector{},
		&expr.BasicLitInspector{},
		&expr.FuncTypeInspector{},
		&expr.IdentInspector{},
		&expr.InterfaceTypeInspector{},
		&expr.StructTypeInspector{},
		// SPEC
		&spec.ImportSpecInspector{},
		&spec.TypeSpecInspector{},
		&spec.ValueSpecInspector{},
		// STMT
		&stmt.ForStmtInspector{},
		&stmt.GoStmtInspector{},
		&stmt.IfStmtInspector{},
		&stmt.ReturnStmtInspector{},
		// NODE
		&CommentGroupInspector{},
		&CommentInspector{},
		&FieldInspector{},
		&FieldListInspector{},
		&FileInspector{},
		&PackageInspector{},
	}

	for _, i := range inspectors {
		if i.IsTarget(node) {
			return i
		}
	}
	return nil
}
