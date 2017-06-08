package spec

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type ValueSpecInspector struct{}

func (i *ValueSpecInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.ValueSpec:
		return true
	}
	return false
}

func (i *ValueSpecInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	// FIXME
	fmt.Printf("ValueSpecInspector: %#v\n", node)
	return nil
}
