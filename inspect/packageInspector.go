package inspect

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type PackageInspector struct{}

func (i *PackageInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.Package:
		return true
	}
	return false
}

func (i *PackageInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	pkg, ok := node.(*ast.Package)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Printf("PackageInspector: %#v\n", pkg)
	return nil
}
