package inspect

import (
	"fmt"
	"go/ast"

	"github.com/sky0621/go-testcode-autogen/inspect/result"
)

type FileInspector struct{}

func (i *FileInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.File:
		return true
	}
	return false
}

func (i *FileInspector) Inspect(node ast.Node, aggregater *result.Aggregater) error {
	fl, ok := node.(*ast.File)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	fmt.Println("===== FileInspector ===================================================================================")
	fmt.Printf("File: %#v\n", fl)
	// MEMO ast.Package で取得できないため、ここで取得
	if fl.Name != nil {
		aggregater.PackageName = fl.Name.Name
	}
	return nil
}
