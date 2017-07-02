package spec

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/inspect/result"
)

type TypeSpecInspector struct{}

func (i *TypeSpecInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.TypeSpec:
		return true
	}
	return false
}

func (i *TypeSpecInspector) Inspect(node ast.Node, aggregater *result.Aggregater) error {
	ts, ok := node.(*ast.TypeSpec)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Println("===== TypeSpecInspector ===================================================================================")
	fmt.Printf("TypeSpec: %#v\n", ts)
	return nil
}
