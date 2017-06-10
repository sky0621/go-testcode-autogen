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

func (i *FieldInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	// FIXME
	fmt.Printf("FieldInspector: %#v\n", node)
	return nil
}
