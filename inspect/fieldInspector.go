package inspect

import (
	"go/ast"

	"github.com/sky0621/go-testcode-autogen/inspect/result"

	"fmt"
)

type FieldInspector struct{}

func (i *FieldInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.Field:
		return true
	}
	return false
}

func (i *FieldInspector) Inspect(node ast.Node, aggregater *result.Aggregater) error {
	fi, ok := node.(*ast.Field)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Println("===== FieldInspector ===================================================================================")
	fmt.Printf("Field: %#v\n", fi)
	return nil
}
