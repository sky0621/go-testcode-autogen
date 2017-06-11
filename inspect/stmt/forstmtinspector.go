package stmt

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type ForStmtInspector struct{}

func (i *ForStmtInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.ForStmt:
		return true
	}
	return false
}

func (i *ForStmtInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	fs, ok := node.(*ast.ForStmt)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Printf("ForStmtInspector: %#v\n", fs)
	return nil
}
