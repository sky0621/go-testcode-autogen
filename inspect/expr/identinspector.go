package expr

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type IdentInspector struct{}

func (i *IdentInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.Ident:
		return true
	}
	return false
}

func (i *IdentInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	idt := node.(*ast.Ident)
	fmt.Println("[[[ IdentInspector ]]]")
	fmt.Printf("Name: %#v\n", idt.Name)
	if idt.Obj != nil {
		fmt.Printf("Obj.Name: %#v\n", idt.Obj.Name)
		fmt.Printf("Obj.Decl: %#v\n", idt.Obj.Decl)
		fmt.Printf("Obj.Type: %#v\n", idt.Obj.Type)
		fmt.Printf("Obj.Data: %#v\n", idt.Obj.Data)
		fmt.Printf("Obj.Kind: %#v\n", idt.Obj.Kind)

	}
	return nil
}
