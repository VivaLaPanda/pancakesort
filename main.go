package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/vivalapanda/pancakesort/astar"
	"github.com/vivalapanda/pancakesort/heuristics"
	"github.com/vivalapanda/pancakesort/permutation"
)

func main() {
	fmt.Printf("Please enter the comma seperated string to sort.\n e.g. \"1,3,5,4,2\" without quotes: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	testNode, _ := permutation.MakeNode(scanner.Text())
	//testNode, _ := permutation.MakeNode("1,3,2,4")
	testGraph := astar.MakeGraph()

	tempNode, err := testGraph.GetGoal(testNode, heuristics.Breakpoints)
	if err != nil {
		fmt.Printf("Fatal Error: '%s'\n", err)
	} else if tempNode == nil {
		fmt.Printf("No solution found!\n")
		fmt.Printf("Nodes searched: %d\n", testGraph.GetNumExpanded())
	}
	goalNode := tempNode.(*permutation.Node)

	fmt.Printf("Goal found! Rotations to sort:\n")

	for tempNode := goalNode; tempNode != nil; tempNode = tempNode.GetParent().(*permutation.Node) {
		fmt.Printf("%q\n", tempNode.Dump())
	}

	fmt.Printf("\nIn total %f nodes were expanded in the search.", tempNode.GetDepth())
}
