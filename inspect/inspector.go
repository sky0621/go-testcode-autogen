package inspect

import (
	"go/ast"

	"github.com/sky0621/go-testcode-autogen/testinfo"
)

type Inspector interface {
	IsTarget(node ast.Node) bool
	Inspect(node ast.Node, info *testinfo.TestInfo) error
}
