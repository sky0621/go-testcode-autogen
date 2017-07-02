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

func (i *CommentGroupInspector) Inspect(node ast.Node, info *testinfo.TestInfo) error {
	cg, ok := node.(*ast.CommentGroup)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Println("===== CommentGroupInspector ===================================================================================")
	fmt.Printf("CommentGroup: %#v\n", cg)
	return nil
}
