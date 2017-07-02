package expr

import (
	"go/ast"

	"github.com/sky0621/go-testcode-autogen/inspect/result"

	"fmt"
)

type StructTypeInspector struct{}

func (i *StructTypeInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.StructType:
		return true
	}
	return false
}

func (i *StructTypeInspector) Inspect(node ast.Node, aggregater *result.Aggregater) error {
	st, ok := node.(*ast.StructType)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Println("===== StructTypeInspector ===================================================================================")
	fmt.Printf("StructType: %#v\n", st)
	return nil
}
