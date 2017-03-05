package heuristics

import (
	"testing"

	"github.com/vivalapanda/pancakesort/permutation"
)

func TestBuildEstimateDistance(t *testing.T) {
	testNode, _ := permutation.MakeNode("1,4,8,9,11,6,2,7,5,3,10")

	actual := Breakpoints(testNode)
	expected := float64(8)
	if actual != expected {
		t.Errorf("Error occured while testing BuildSudokuBoard: e:'%v' != a:'%v'", expected, actual)
	}
}
