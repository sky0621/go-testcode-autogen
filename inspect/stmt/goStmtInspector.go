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

func (i *GoStmtInspector) Inspect(node ast.Node, info *testinfo.TestInfo) error {
	gs, ok := node.(*ast.GoStmt)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Println("===== GoStmtInspector ===================================================================================")
	fmt.Printf("GoStmt: %#v\n", gs)
	return nil
}
