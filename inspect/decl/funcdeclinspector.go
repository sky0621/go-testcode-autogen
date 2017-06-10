package decl

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type FuncDeclInspector struct{}

func (i *FuncDeclInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.FuncDecl:
		return true
	}
	return false
}

func (i *FuncDeclInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	fd := node.(*ast.FuncDecl)
	// FIXME
	fmt.Printf("FuncDeclInspector: %#v\n", fd)
	return nil
}
