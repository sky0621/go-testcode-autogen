package inspect

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type CommentGroupInspector struct{}

func (i *CommentGroupInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.CommentGroup:
		return true
	}
	return false
}

func (i *CommentGroupInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	cg := node.(*ast.CommentGroup)
	// FIXME
	fmt.Printf("CommentGroupInspector: %#v\n", cg)
	return nil
}
