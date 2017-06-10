package expr

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type StructTypeInspector struct{}

func (i *StructTypeInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.StructType:
		return true
	}
	return false
}

func (i *StructTypeInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	st := node.(*ast.StructType)
	// FIXME
	fmt.Printf("StructTypeInspector: %#v\n", st)
	return nil
}
