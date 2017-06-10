package stmt

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type GoStmtInspector struct{}

func (i *GoStmtInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.GoStmt:
		return true
	}
	return false
}

func (i *GoStmtInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	gs := node.(*ast.GoStmt)
	// FIXME
	fmt.Printf("GoStmtInspector: %#v\n", gs)
	return nil
}
