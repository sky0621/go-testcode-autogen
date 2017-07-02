package spec

import (
	"go/ast"

	"github.com/sky0621/go-testcode-autogen/inspect/result"

	"fmt"
)

type ValueSpecInspector struct{}

func (i *ValueSpecInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.ValueSpec:
		return true
	}
	return false
}

func (i *ValueSpecInspector) Inspect(node ast.Node, aggregater *result.Aggregater) error {
	vs, ok := node.(*ast.ValueSpec)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Println("===== ValueSpecInspector ===================================================================================")
	fmt.Printf("ValueSpec: %#v\n", vs)
	return nil
}
