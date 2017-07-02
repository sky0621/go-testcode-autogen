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

func (i *ReturnStmtInspector) Inspect(node ast.Node, info *testinfo.TestInfo) error {
	rs, ok := node.(*ast.ReturnStmt)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Println("===== ReturnStmtInspector ===================================================================================")
	fmt.Printf("ReturnStmt: %#v\n", rs)
	return nil
}
