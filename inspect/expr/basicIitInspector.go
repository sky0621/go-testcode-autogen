package expr

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type BasicLitInspector struct{}

func (i *BasicLitInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.BasicLit:
		return true
	}
	return false
}

func (i *BasicLitInspector) Inspect(node ast.Node, info *testinfo.TestInfo) error {
	bl, ok := node.(*ast.BasicLit)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Println("===== BasicLitInspector ===================================================================================")
	fmt.Printf("BasicLit: %#v\n", bl)
	return nil
}
