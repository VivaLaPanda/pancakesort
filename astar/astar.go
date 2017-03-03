package astar

import (
	ds "github.com/hishboy/gocommons/lang"
	"github.com/oleiade/lane"
	"github.com/vivalapanda/pancakesort/heuristics"
	"sync"
)

type GraphNode interface {
	Children() []*GraphNode
	IsGoal() bool
	GetParent() *GraphNode
	SetParent(*GraphNode) *GraphNode
	GetDepth() int
}

type Graph struct {
	open        *lane.PQueue
	closed      *ds.HashSet
	numExpanded int
	lock        *sync.Mutex
}

func MakeGraph() *Graph {
	graph := &Graph{}

	graph.open = lane.NewPQueue(lane.MINPQ)
	graph.closed = ds.NewHashSet()
	graph.numExpanded = 0

	return graph
}

func (graph *Graph) GetGoal(rootNode *GraphNode, hfunc heuristics.Any) *GraphNode {

	return node
}

// Method to expand thr first node on the open set
// Expects: A valid graph, the depth of nodeToExpand, and a heuristic with which to
// evaluate the nodes in the graph.
func (graph Graph) expand(hfunc heuristics.Any) (newGraph Graph, nodeExpanded *GraphNode) {
	// Making sure the graph object we have is completely dereferenced from the
	// one passed in
	graph.open = &graph.open
	graph.closed = &graph.closed

	// Expand all of the nodes in the
	nodeToExpand := graph.open.Pop().(*GraphNode)
	newNodes := nodeToExpand.Children()
	graph.closed.Push(nodeToExpand)

	for _, node := range newNodes {
		// Calculate w(n) + f(n)
		distance := hfunc(node)
		cost := node.GetDepth()

		graph.open.Push(node, distance+cost)
	}

	return graph, nodeToExpand
}
