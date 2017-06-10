package inspect

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type FieldListInspector struct{}

func (i *FieldListInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.FieldList:
		return true
	}
	return false
}

func (i *FieldListInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	// FIXME
	fmt.Printf("FieldListInspector: %#v\n", node)
	return nil
}
