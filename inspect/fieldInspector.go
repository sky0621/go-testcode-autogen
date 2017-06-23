package inspect

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type FieldInspector struct{}

func (i *FieldInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.Field:
		return true
	}
	return false
}

func (i *FieldInspector) Inspect(node ast.Node, info *testinfo.TestInfo) error {
	fi, ok := node.(*ast.Field)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Printf("FieldInspector: %#v\n", fi)
	return nil
}
