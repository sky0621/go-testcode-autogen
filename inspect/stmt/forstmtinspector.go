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
	// FIXME
	fmt.Printf("ForStmtInspector: %#v\n", node)
	return nil
}
