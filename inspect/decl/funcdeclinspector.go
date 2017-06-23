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

func (i *FuncDeclInspector) Inspect(node ast.Node, info *testinfo.TestInfo) error {
	fd, ok := node.(*ast.FuncDecl)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	if fd.Name == nil {
		return nil
	}
	info.Functions = append(info.Functions, &testinfo.Function{FunctionName: fd.Name.Name})
	// FIXME
	fmt.Printf("FuncDeclInspector: Name[%#v], Type[%#v], Body[%#v], Recv[%#v]\n", fd.Name, fd.Type, fd.Body, fd.Recv)
	return nil
}
