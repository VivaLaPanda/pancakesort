package heuristics

import (
	"github.com/vivalapanda/pancakesort/permutation"
)

func Breakpoints(node interface{}) float64 {
	return float64(node.(*permutation.Node).CountBreakpoints())
}
