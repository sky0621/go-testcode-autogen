package decl

import (
	"go/ast"

	"github.com/sky0621/go-testcode-autogen/inspect/result"

	"fmt"
)

type BadDeclInspector struct{}

func (i *BadDeclInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.BadDecl:
		return true
	}
	return false
}

func (i *BadDeclInspector) Inspect(node ast.Node, aggregater *result.Aggregater) error {
	bd, ok := node.(*ast.BadDecl)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Println("===== BadDeclInspector ===================================================================================")
	fmt.Printf("BadDecl: %#v\n", bd)
	return nil
}
