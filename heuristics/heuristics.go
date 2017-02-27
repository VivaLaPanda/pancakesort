package heuristics

import (
	"github.com/vivalapanda/pancakesort/permutation"
)

func Breakpoints(node *permutation.Node) int {
	return node.CountBreakpoints()
}
