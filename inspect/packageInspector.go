package inspect

import (
	"go/ast"

	"fmt"
)

type PackageInspector struct{}

func (i *PackageInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.Package:
		return true
	}
	return false
}

func (i *PackageInspector) Inspect(node ast.Node, info *ResultAggregater) error {
	pkg, ok := node.(*ast.Package)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Println("===== PackageInspector ===================================================================================")
	fmt.Printf("Package: %#v\n", pkg)
	return nil
}
