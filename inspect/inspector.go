package inspect

import (
	"go/ast"

	"github.com/sky0621/go-testcode-autogen/inspect/result"
)

type Inspector interface {
	IsTarget(node ast.Node) bool
	Inspect(node ast.Node, aggregater *result.Aggregater) error
}
