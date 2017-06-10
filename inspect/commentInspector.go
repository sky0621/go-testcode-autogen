package inspect

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type CommentInspector struct{}

func (i *CommentInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.Comment:
		return true
	}
	return false
}

func (i *CommentInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	cmt := node.(*ast.Comment)
	// FIXME
	fmt.Printf("CommentInspector: %#v\n", cmt)
	return nil
}
