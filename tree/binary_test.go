package tree

import (
	"reflect"
	"testing"
)

func TestNewTree(t *testing.T) {
	tree := NewTree()

	got := tree.Left.Val
	want := "B"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}

func TestDepthFirstValues(t *testing.T) {
	tree := NewTree()

	got := DepthFirstValues(tree)
	want := []string{"A", "B", "D", "E", "C", "F", "G"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}

func TestMaxDepth(t *testing.T) {
	tree := NewTree()

	got := MaxDepth(tree)
	want := 4

	if want != got {
		t.Errorf("got %d want %d", got, want)
	}

}

func TestDepthFirstValuesRecursive(t *testing.T) {
	tree := NewTree()

	got := DepthFirstValuesRecursive(tree)
	want := []string{"A", "B", "D", "E", "C", "F", "G"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}

func TestBreadthFirstValues(t *testing.T) {
	tree := NewTree()

	got := BreadthFirstValues(tree)
	want := []string{"A", "B", "C", "D", "E", "F", "G"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}
