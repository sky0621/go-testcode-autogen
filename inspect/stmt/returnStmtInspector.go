package stmt

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type ReturnStmtInspector struct{}

func (i *ReturnStmtInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.ReturnStmt:
		return true
	}
	return false
}

func (i *ReturnStmtInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	rs := node.(*ast.ReturnStmt)
	// FIXME
	fmt.Printf("ReturnStmtInspector: %#v\n", rs)
	return nil
}
