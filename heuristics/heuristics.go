package heuristics

import (
	"github.com/vivalapanda/pancakesort/permutation"
)

type Any func(interface{}) float64

func Breakpoints(node *permutation.Node) float64 {
	return float64(node.CountBreakpoints())
}
