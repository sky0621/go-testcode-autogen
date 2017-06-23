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

func (i *BadExprInspector) Inspect(node ast.Node, info *testinfo.TestInfo) error {
	be, ok := node.(*ast.BadExpr)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Printf("BadExprInspector: %#v\n", be)
	return nil
}
