package astar

import (
	ds "github.com/hishboy/gocommons/lang"
	"github.com/oleiade/lane"
	"github.com/vivalapanda/pancakesort/heuristics"
)

type GraphNode interface {
	Children() []*GraphNode
	IsGoal() bool
	GetParent() *GraphNode
	SetParent(*GraphNode) *GraphNode
}

type Graph struct {
	open        *lane.PQueue
	closed      *ds.Stack
	numExpanded int
}

func MakeGraph() *Graph {
	graph := &Graph{}

	graph.open = lane.NewPQueue(lane.MINPQ)
	graph.closed = ds.NewStack()

	return graph
}

func GetGoal(node *GraphNode, hfunc heuristics.Any) *GraphNode {

	return node
}
