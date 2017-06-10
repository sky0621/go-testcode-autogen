package inspect

import (
	"go/ast"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type Inspector interface {
	IsTarget(node ast.Node) bool
	Inspect(node ast.Node, testinfo *testinfo.TestInfo) error
}

type DeclInspector interface {
	IsTargetDecl(decl ast.Decl) bool
	InspectDecl(decl ast.Decl, testinfo *testinfo.TestInfo) error
	Inspector
}
