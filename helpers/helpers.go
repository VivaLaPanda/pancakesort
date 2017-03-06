package helpers

type GraphNode interface {
	Children() []GraphNode
	IsGoal() bool
	GetParent() GraphNode
	SetParent(GraphNode) GraphNode
	GetDepth() float64
	Key() interface{}
}

type AnyHeuristic func(interface{}) float64
