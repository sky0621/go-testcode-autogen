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
	it := node.(*ast.InterfaceType)
	// FIXME
	fmt.Printf("InterfaceTypeInspector: %#v\n", it)
	return nil
}
