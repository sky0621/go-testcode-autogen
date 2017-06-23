package stmt

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type IfStmtInspector struct{}

func (i *IfStmtInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.IfStmt:
		return true
	}
	return false
}

func (i *IfStmtInspector) Inspect(node ast.Node, info *testinfo.TestInfo) error {
	is, ok := node.(*ast.IfStmt)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Printf("IfStmtInspector: %#v\n", is)
	return nil
}
