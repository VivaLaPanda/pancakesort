package astar

import (
	"fmt"
	"sync"

	heap "github.com/vivalapanda/go-datastructures/fibheap"
	"github.com/vivalapanda/pancakesort/helpers"
)

type Graph struct {
	open        heap.FloatingFibonacciHeap
	closed      map[interface{}]helpers.GraphNode
	numExpanded int
	lock        *sync.Mutex
}

func MakeGraph() *Graph {
	graph := &Graph{}

	graph.open = heap.NewFloatFibHeap()
	graph.closed = make(map[interface{}]helpers.GraphNode)
	graph.numExpanded = 0

	return graph
}

func (graph *Graph) GetGoal(rootNode helpers.GraphNode, hfunc helpers.AnyHeuristic) (helpers.GraphNode, error) {
	// Search until you run out of nodes
	for graph.open.Enqueue(rootNode, hfunc(rootNode)); !graph.open.IsEmpty(); {
		entry, err := graph.open.DequeueMin()
		if err != nil {
			return nil, fmt.Errorf("Fatal error while traversing tree!\n%v", err)
		}
		activeNode := entry.Value.(helpers.GraphNode)

		graph.closed[activeNode.Key()] = activeNode

		if activeNode.IsGoal() {
			return activeNode, nil
		} else {
			graph = graph.expand(activeNode, hfunc)
		}

		fmt.Printf("Set: %v\n", graph.open)
	}

	// We didn't find a goal node
	return nil, nil
}

// Method to expand thr first node on the open set
// Expects: A valid graph, the depth of nodeToExpand, and a heuristic with which to
// evaluate the nodes in the graph.
func (graph *Graph) expand(nodeToExpand helpers.GraphNode, hfunc helpers.AnyHeuristic) (newGraph *Graph) {
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

	return graph
}
