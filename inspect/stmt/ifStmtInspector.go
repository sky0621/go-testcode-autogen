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

func (i *IfStmtInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	is := node.(*ast.IfStmt)
	// FIXME
	fmt.Printf("IfStmtInspector: %#v\n", is)
	return nil
}
