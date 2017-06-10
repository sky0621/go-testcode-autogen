package inspect

import (
	"go/ast"

	"fmt"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type FileInspector struct{}

func (i *FileInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.File:
		return true
	}
	return false
}

func (i *FileInspector) Inspect(node ast.Node, testinfo *testinfo.TestInfo) error {
	// FIXME
	fmt.Printf("FileInspector: %#v\n", node)
	return nil
}
