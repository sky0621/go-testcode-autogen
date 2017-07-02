package inspect

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/inspect/result"
)

type CommentInspector struct{}

func (i *CommentInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.Comment:
		return true
	}
	return false
}

func (i *CommentInspector) Inspect(node ast.Node, aggregater *result.Aggregater) error {
	cmt, ok := node.(*ast.Comment)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// FIXME
	fmt.Println("===== CommentInspector ===================================================================================")
	fmt.Printf("Comment: %#v\n", cmt)
	return nil
}
