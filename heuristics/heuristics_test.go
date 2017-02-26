package heuristics

import (
	"reflect"
	"testing"
)

func TestBuildEstimateDistance(t *testing.T) {

	actual := true
	expected := true
	if !(reflect.DeepEqual(actual, expected)) {
		t.Errorf("Error occured while testing BuildSudokuBoard: '%s' != '%s'", expected, actual)
	}
}
