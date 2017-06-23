package decl

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type GenDeclInspector struct{}

func (i *GenDeclInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.GenDecl:
		return true
	}
	return false
}

func (i *GenDeclInspector) Inspect(node ast.Node, info *testinfo.TestInfo) error {
	gd, ok := node.(*ast.GenDecl)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Printf("GenDeclInspector: %#v\n", gd)
	return nil
}
