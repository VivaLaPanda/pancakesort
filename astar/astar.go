package astar

import (
	"fmt"
	"sync"

	heap "github.com/vivalapanda/go-datastructures/fibheap"
	"github.com/vivalapanda/pancakesort/heuristics"
)

type GraphNode interface {
	Children() []GraphNode
	IsGoal() bool
	GetParent() GraphNode
	SetParent(GraphNode) GraphNode
	GetDepth() float64
	Key() interface{}
}

type Graph struct {
	open        heap.FloatingFibonacciHeap
	closed      map[interface{}]GraphNode
	numExpanded int
	lock        *sync.Mutex
}

func MakeGraph() *Graph {
	graph := &Graph{}

	graph.open = heap.NewFloatFibHeap()
	graph.closed = make(map[interface{}]GraphNode)
	graph.numExpanded = 0

	return graph
}

func (graph *Graph) GetGoal(rootNode GraphNode, hfunc heuristics.Any) (GraphNode, error) {
	graph.lock.Lock()
	graph.open.Enqueue(rootNode, hfunc(rootNode))
	graph.lock.Unlock()

	// Search until you run out of nodes
	for graph.open.IsEmpty() {
		entry, err := graph.open.DequeueMin()
		if err != nil {
			return nil, fmt.Errorf("Fatal error while traversing tree!\n%v", err)
		}
		activeNode := entry.Value.(GraphNode)

		graph.closed[activeNode.Key()] = activeNode

		if activeNode.IsGoal() {
			return activeNode, nil
		} else {
			graph = graph.expand(activeNode, hfunc)
		}

	}

	// We didn't find a goal node
	return nil, nil
}

// Method to expand thr first node on the open set
// Expects: A valid graph, the depth of nodeToExpand, and a heuristic with which to
// evaluate the nodes in the graph.
func (graph Graph) expand(nodeToExpand GraphNode, hfunc heuristics.Any) (newGraph *Graph) {
	// Making sure the graph object we have is completely dereferenced from the
	// one passed in

	// Expand all of the nodes in the
	newNodes := nodeToExpand.Children()

	for _, node := range newNodes {
		// Calculate w(n) + f(n)
		estimatedCost := hfunc(node) + node.GetDepth()

		oldNode, exists := graph.closed[node.Key()]

		if exists {
			oldCost := oldNode.GetDepth() + hfunc(oldNode)
			if estimatedCost < oldCost {
				delete(graph.closed, oldNode)
			}
		}

		graph.open.Enqueue(node, estimatedCost)
	}

	return &graph
}
