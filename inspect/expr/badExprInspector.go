package expr

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/inspect/result"
)

type BadExprInspector struct{}

func (i *BadExprInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.BadExpr:
		return true
	}
	return false
}

func (i *BadExprInspector) Inspect(node ast.Node, aggregater *result.Aggregater) error {
	be, ok := node.(*ast.BadExpr)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Println("===== BadExprInspector ===================================================================================")
	fmt.Printf("BadExpr: %#v\n", be)
	return nil
}
