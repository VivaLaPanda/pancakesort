package astar

import (
	"reflect"
	"testing"

	"github.com/vivalapanda/pancakesort/heuristics"
	"github.com/vivalapanda/pancakesort/permutation"
)

func TestGetGoal(t *testing.T) {
	testNode, _ := permutation.MakeNode("1,4,8,9,11,6,2,7,5,3,10")
	testGraph := MakeGraph()

	tempNode, err := testGraph.GetGoal(testNode, heuristics.Breakpoints)
	if err != nil {
		t.Errorf("Error occured while testing GetGoal, method returned an error: '%s'", err)
	} else if tempNode == nil {
		t.Errorf("Error occured while testing GetGoal, solution not found")
	}
	goalNode := tempNode.(*permutation.Node)

	actual := goalNode.Key()
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	if !(reflect.DeepEqual(actual, expected)) {
		t.Errorf("Error occured while testing GetGoal: '%v' != '%v'", expected, actual)
	}
}
