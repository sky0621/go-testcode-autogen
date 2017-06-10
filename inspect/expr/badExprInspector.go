package expr

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type BadExprInspector struct{}

func (i *BadExprInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.BadExpr:
		return true
	}
	return false
}

func (i *BadExprInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	be := node.(*ast.BadExpr)
	// FIXME
	fmt.Printf("BadExprInspector: %#v\n", be)
	return nil
}
