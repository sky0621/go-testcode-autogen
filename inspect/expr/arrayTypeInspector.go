package expr

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type ArrayTypeInspector struct{}

func (i *ArrayTypeInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.ArrayType:
		return true
	}
	return false
}

func (i *ArrayTypeInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	at, ok := node.(*ast.ArrayType)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Printf("ArrayTypeInspector: %#v\n", at)
	return nil
}
