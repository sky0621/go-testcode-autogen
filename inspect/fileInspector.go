package inspect

import (
	"fmt"
	"go/ast"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type FileInspector struct{}

func (i *FileInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.File:
		return true
	}
	return false
}

func (i *FileInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	fl := node.(*ast.File)
	fmt.Println("[[[ FileInspector ]]]")
	fmt.Printf("Name: %#v\n", fl.Name)
	if fl.Scope != nil {
		fmt.Printf("Scope.Objects: %#v\n", fl.Scope.Objects)
		fmt.Printf("Scope.Outer: %#v\n", fl.Scope.Outer)
	}
	//for _, d := range fl.Decls {
	//
	//	fmt.Printf("Obj.Name: %#v\n", d.)
	//
	//}
	return nil
}
