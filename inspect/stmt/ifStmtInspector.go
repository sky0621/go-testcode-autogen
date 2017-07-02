package stmt

import (
	"go/ast"

	"github.com/sky0621/go-testcode-autogen/inspect/result"

	"fmt"
)

type IfStmtInspector struct{}

func (i *IfStmtInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.IfStmt:
		return true
	}
	return false
}

func (i *IfStmtInspector) Inspect(node ast.Node, aggregater *result.Aggregater) error {
	is, ok := node.(*ast.IfStmt)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Println("===== IfStmtInspector ===================================================================================")
	fmt.Printf("IfStmt: %#v\n", is)
	return nil
}
