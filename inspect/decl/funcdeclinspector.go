package decl

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
	"github.com/sky0621/go-testcode-autogen/util"
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
	if !util.IsFirstUpper(fd.Name.Name) {
		return nil
	}
	fmt.Println("===== FuncDeclInspector ===================================================================================")
	info.Functions = append(info.Functions, &testinfo.Function{
		FunctionName: fd.Name.Name},
	)
	// FIXME
	fmt.Printf("FuncDecl: \n>Name[%#v], \n>Type[%#v], \n>Body[%#v], \n>Recv[%#v]\n", fd.Name, fd.Type, fd.Body, fd.Recv)
	return nil
}
