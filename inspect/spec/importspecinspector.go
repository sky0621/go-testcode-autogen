package spec

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type ImportSpecInspector struct{}

func (i *ImportSpecInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.ImportSpec:
		return true
	}
	return false
}

func (i *ImportSpecInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	// FIXME
	fmt.Printf("ImportSpecInspector: %#v\n", node)
	return nil
}
