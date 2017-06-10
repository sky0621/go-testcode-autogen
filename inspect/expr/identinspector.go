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
	// FIXME
	fmt.Printf("IdentInspector: %#v\n", idt)
	return nil
}
