package spec

import (
	"go/ast"

	"fmt"

	"strings"

	"github.com/sky0621/go-testcode-autogen/inspect/result"
)

type ImportSpecInspector struct{}

func (i *ImportSpecInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.ImportSpec:
		return true
	}
	return false
}

func (i *ImportSpecInspector) Inspect(node ast.Node, aggregater *result.Aggregater) error {
	is, ok := node.(*ast.ImportSpec)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	if is.Path == nil {
		return nil
	}
	fmt.Println("===== ImportSpecInspector ===================================================================================")
	fmt.Printf("ImportSpec: %#v\n", is)
	aggregater.ImportNames = append(aggregater.ImportNames, strings.Trim(is.Path.Value, `""`))
	return nil
}
