package permutation

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/vivalapanda/pancakesort/helpers"
)

type Node struct {
	contents []int
	parent   *Node
	lock     *sync.Mutex
}

// Constructor for Node type
// Expects: a comma seperated string representing an arrangement of integers
// Returns: A reference to a new node representing the passed state
func MakeNode(arrange string) (*Node, error) {
	// Makes a new instance of the Node type
	node := &Node{}
	node.parent = nil
	node.lock = &sync.Mutex{}

	// Parses the passed comma seperated string into the contents of the node
	strSlice := strings.Split(arrange, ",") // Slice is composed of strings at this point
	intSlice := make([]int, len(strSlice))

	for i, element := range strSlice {
		intElement, err := strconv.Atoi(element)
		if err != nil {
			err = fmt.Errorf("Failure to parse given string (%s). Error: %v", element, err)
			return nil, err
		} else {
			intSlice[i] = intElement
		}
	}

	node.contents = intSlice

	// Returns a reference to the node
	return node, nil
}

// Gets the children of a node
// Expects: a valid rootNode
// Returns: a slice of Nodes
func (rootNode *Node) Children() []helpers.GraphNode {
	childrenSlice := []helpers.GraphNode{}

	// Iterate through all size 2 groupings of adjacent elements
	for subLen := len(rootNode.contents) - 1; subLen > 2; subLen-- {
		for i := 0; i+subLen < len(rootNode.contents); i++ {
			newNode := &Node{}
			newNode.parent = rootNode
			newNode.lock = &sync.Mutex{}

			// Copy the array associated with rootNode so as not to modify the original
			tempContents := make([]int, len(rootNode.contents), (cap(rootNode.contents)))
			copy(tempContents, rootNode.contents)

			// Do the rotation
			newContents := append(tempContents[:i], reverse(tempContents[i:subLen+i])...)
			newContents = append(newContents, tempContents[subLen+i:len(tempContents)]...)
			// Put the new node on the slice
			newNode.contents = newContents

			childrenSlice = append(childrenSlice, newNode)
		}
	}

	return childrenSlice
}

// Function which returns a pointer to the given node's Parent
// Expects: A valid node
// Returns: The parent of the node
func (node *Node) GetParent() helpers.GraphNode {
	return node.parent
}

// Function which modifies a node's parent, and then returns the original node
// Expects: A two valid nodes
// Returns: The node in the function reciever with the new parent
func (node *Node) SetParent(newParent helpers.GraphNode) helpers.GraphNode {
	node.lock.Lock()
	defer node.lock.Unlock()

	node.parent = newParent.(*Node)

	return node
}

// Helper function which reverses the slice passed
func reverse(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	return numbers
}

// Checks if the given node is a goal node
// Expects: A valid node
// Returns: A boolean representing whether the node was a goal
func (node *Node) IsGoal() bool {
	if node.CountBreakpoints() > 0 {
		return false
	}

	return true
}

// Returns the contents of the node so that it's state can be uniquely idenfitifed
// Expects: A valid node
// Returns: A slice of ints
func (node *Node) Key() interface{} {
	return fmt.Sprintf("%v", node.contents)
}

// Returns the number of breakpoints
// Expects: A valid node
// Returns: an int representing the number of breakpoints
func (node *Node) CountBreakpoints() int {
	numBreakpoints := 0

	// If we find any breakpoints return false
	// A breakpoint in a permutation X is a position j such that X(j) + 1 ≠ X(j+1)
	for i, element := range node.contents[:len(node.contents)-1] {
		if math.Abs(float64(element+1-node.contents[i+1])) > 1 {
			numBreakpoints++
		}
	}

	return numBreakpoints
}

func (node *Node) GetDepth() float64 {
	return float64(recDepth(node))
}

func recDepth(node helpers.GraphNode) int {
	if node.GetParent().(*Node) == nil {
		return 0
	} else {
		return recDepth(node.GetParent()) + 1
	}
}

// Gives the contents of a node in a comma seperated string
// Expects: a valid node
// Returns: A string of comma seperated integers
func (node *Node) Dump() string {
	strSlice := []string{}

	// Casting contents
	for _, elem := range node.contents {
		s := strconv.Itoa(elem)
		strSlice = append(strSlice, s)
	}

	dumpString := strings.Join(strSlice, ",")

	return dumpString
}
