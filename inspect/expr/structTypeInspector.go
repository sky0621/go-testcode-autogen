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

func (i *StructTypeInspector) Inspect(node ast.Node, info *testinfo.TestInfo) error {
	st, ok := node.(*ast.StructType)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Println("===== StructTypeInspector ===================================================================================")
	fmt.Printf("StructType: %#v\n", st)
	return nil
}
