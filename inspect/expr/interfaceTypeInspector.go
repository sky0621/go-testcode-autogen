package expr

import (
	"go/ast"

	"github.com/sky0621/go-testcode-autogen/inspect/result"

	"fmt"
)

type InterfaceTypeInspector struct{}

func (i *InterfaceTypeInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.InterfaceType:
		return true
	}
	return false
}

func (i *InterfaceTypeInspector) Inspect(node ast.Node, aggregater *result.Aggregater) error {
	it, ok := node.(*ast.InterfaceType)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Println("===== InterfaceTypeInspector ===================================================================================")
	fmt.Printf("InterfaceType: %#v\n", it)
	return nil
}
