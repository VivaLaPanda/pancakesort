package permutation

import (
	"reflect"
	"testing"
)

func TestMakeNode(t *testing.T) {
	testNode, err := MakeNode("1,4,8,9,11,6,2,7,5,3,10")

	if err != nil {
		t.Errorf("Fatal error while making new node: %s", err)
	}

	actual := testNode.contents
	expected := []int{1, 4, 8, 9, 11, 6, 2, 7, 5, 3, 10}

	if !(reflect.DeepEqual(actual, expected)) {
		t.Errorf("Error occured while testing MakeNode: '%v' != '%v'", expected, actual)
	}

	_, err_2 := MakeNode("a; 1, 12.. a")

	if err_2 == nil {
		t.Errorf("Error occured while testing MakeNode: Bad string is not rejected")
	}
}

func TestChildren(t *testing.T) {
	testNode, _ := MakeNode("1,4,8,9,11,6,2,7,5,3,10")

	actual := testNode.contents
	expected := []int{1, 4, 8, 9, 11, 6, 2, 7, 5, 3, 10}

	if !(reflect.DeepEqual(actual, expected)) {
		t.Errorf("Error occured while testing Childen, parent was modified: '%v' != '%v'", expected, actual)
	}

	actual_2 := testNode.Children()[0].(*Node).contents
	expected_2 := []int{3, 5, 7, 2, 6, 11, 9, 8, 4, 1, 10}

	if !(reflect.DeepEqual(actual_2, expected_2)) {
		t.Errorf("Error occured while testing Childen: '%v' != '%v'", expected_2, actual_2)
	}
}

func TestGerParent(t *testing.T) {
	testNode, _ := MakeNode("1,2,3,4,5")

	actual := testNode.Children()[0].GetParent().(*Node).contents
	expected := []int{1, 2, 3, 4, 5}

	if !(reflect.DeepEqual(actual, expected)) {
		t.Errorf("Error occured while testing GetParent: '%v' != '%v'", expected, actual)
	}
}

func TestSetParent(t *testing.T) {
	testNode, _ := MakeNode("1,2,3,4,5")
	testNode_2, _ := MakeNode("1,2,3,4,6")

	child := testNode.Children()[0].SetParent(testNode_2)
	actual := child.GetParent().(*Node).contents
	expected := []int{1, 2, 3, 4, 6}

	if !(reflect.DeepEqual(actual, expected)) {
		t.Errorf("Error occured while testing SetParent: '%v' != '%v'", expected, actual)
	}
}

func TestIsGoal(t *testing.T) {
	goalNode, _ := MakeNode("1,2,3,4,5")

	actual := goalNode.IsGoal()
	expected := true

	if actual != expected {
		t.Errorf("Error occured while testing IsGoal: e:'%v' != a:'%v'", expected, actual)
	}

	failNode, _ := MakeNode("1,2,3,5,4")

	actual_2 := failNode.IsGoal()
	expected_2 := false

	if actual_2 != expected_2 {
		t.Errorf("Error occured while testing IsGoal: e:'%v' != a:'%v'", expected_2, actual_2)
	}
}

func TestDump(t *testing.T) {
	goalNode, _ := MakeNode("1,2,3,4,5")

	actual := goalNode.Dump()
	expected := "1,2,3,4,5"

	if actual != expected {
		t.Errorf("Error occured while testing Dump: e:'%v' != a:'%v'", expected, actual)
	}
}
