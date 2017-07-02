package expr

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/inspect/result"
)

type FuncTypeInspector struct{}

func (i *FuncTypeInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.FuncType:
		return true
	}
	return false
}

func (i *FuncTypeInspector) Inspect(node ast.Node, aggregater *result.Aggregater) error {
	ft, ok := node.(*ast.FuncType)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Println("===== FuncTypeInspector ===================================================================================")
	fmt.Printf("FuncType: %#v\n", ft)
	return nil
}
