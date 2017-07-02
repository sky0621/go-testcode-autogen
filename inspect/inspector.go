package inspect

import "go/ast"

type Inspector interface {
	IsTarget(node ast.Node) bool
	Inspect(node ast.Node, info *ResultAggregater) error
}
