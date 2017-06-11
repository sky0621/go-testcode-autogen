package spec

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type ImportSpecInspector struct{}

func (i *ImportSpecInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.ImportSpec:
		return true
	}
	return false
}

func (i *ImportSpecInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	is, ok := node.(*ast.ImportSpec)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Printf("ImportSpecInspector: %#v\n", is)
	return nil
}
