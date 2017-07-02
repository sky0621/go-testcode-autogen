package decl

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/inspect/result"
)

type GenDeclInspector struct{}

func (i *GenDeclInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.GenDecl:
		return true
	}
	return false
}

func (i *GenDeclInspector) Inspect(node ast.Node, aggregater *result.Aggregater) error {
	gd, ok := node.(*ast.GenDecl)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Println("===== GenDeclInspector ===================================================================================")
	fmt.Printf("GenDecl: %#v\n", gd)
	return nil
}
