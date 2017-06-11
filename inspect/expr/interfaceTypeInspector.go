package expr

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type InterfaceTypeInspector struct{}

func (i *InterfaceTypeInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.InterfaceType:
		return true
	}
	return false
}

func (i *InterfaceTypeInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	it, ok := node.(*ast.InterfaceType)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Printf("InterfaceTypeInspector: %#v\n", it)
	return nil
}
