package inspect

import (
	"fmt"
	"go/ast"
)

type FileInspector struct{}

func (i *FileInspector) IsTarget(node ast.Node) bool {
	switch node.(type) {
	case *ast.File:
		return true
	}
	return false
}

func (i *FileInspector) Inspect(node ast.Node, info *ResultAggregater) error {
	fl, ok := node.(*ast.File)
	if !ok {
		return fmt.Errorf("Not target Node: %#v", node)
	}
	fmt.Println("===== FileInspector ===================================================================================")
	fmt.Printf("File: %#v\n", fl)
	// MEMO ast.Package で取得できないため、ここで取得
	if fl.Name != nil {
		info.PackageName = fl.Name.Name
	}
	return nil
}
