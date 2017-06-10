package inspect

import (
	"fmt"
	"go/ast"

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
	fl := node.(*ast.File)
	// FIXME
	fmt.Printf("FileInspector: %#v\n", fl)
	return nil
}
