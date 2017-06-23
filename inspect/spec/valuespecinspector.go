package spec

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type ValueSpecInspector struct{}

func (i *ValueSpecInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.ValueSpec:
		return true
	}
	return false
}

func (i *ValueSpecInspector) Inspect(node ast.Node, info *testinfo.TestInfo) error {
	vs, ok := node.(*ast.ValueSpec)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Printf("ValueSpecInspector: %#v\n", vs)
	return nil
}
