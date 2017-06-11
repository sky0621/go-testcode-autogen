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
	fl, ok := node.(*ast.File)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	// MEMO ast.Package で取得できないため、ここで取得
	if fl.Name != nil {
		testinfo.PackageName = fl.Name.Name
	}
	return nil
}
