package inspect

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type FieldListInspector struct{}

func (i *FieldListInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.FieldList:
		return true
	}
	return false
}

func (i *FieldListInspector) Inspect(node ast.Node, info *testinfo.TestInfo) error {
	fl, ok := node.(*ast.FieldList)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Println("===== FieldListInspector ===================================================================================")
	fmt.Printf("FieldList: %#v\n", fl)
	return nil
}
