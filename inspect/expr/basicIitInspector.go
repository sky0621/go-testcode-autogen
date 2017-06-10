package expr

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type BasicLitInspector struct{}

func (i *BasicLitInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.BasicLit:
		return true
	}
	return false
}

func (i *BasicLitInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	// FIXME
	fmt.Printf("BasicLitInspector: %#v\n", node)
	return nil
}
