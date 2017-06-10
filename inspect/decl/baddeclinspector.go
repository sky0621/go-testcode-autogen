package decl

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type BadDeclInspector struct{}

func (i *BadDeclInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.BadDecl:
		return true
	}
	return false
}

func (i *BadDeclInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	bd := node.(*ast.BadDecl)
	// FIXME
	fmt.Printf("BadDeclInspector: %#v\n", bd)
	return nil
}
