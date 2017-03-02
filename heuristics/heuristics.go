package heuristics

import (
	"github.com/vivalapanda/pancakesort/permutation"
)

type Any func(*permutation.Node) int

func Breakpoints(node *permutation.Node) int {
	return node.CountBreakpoints()
}
