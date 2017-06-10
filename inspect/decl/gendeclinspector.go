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

func (i *GenDeclInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	gd := node.(*ast.GenDecl)
	// FIXME
	fmt.Printf("GenDeclInspector: %#v\n", gd)
	return nil
}
