package tree

import (
	"testing"
)

func TestBinaryTreeSearch(t *testing.T) {
	tree := getSampleTree()

	node := tree.Search(2)
	nodeAssert(*tree.root.Left, *node, t)

	node = tree.Search(1)
	nodeAssert(*tree.root, *node, t)

	node = tree.Search(4)
	nodeAssert(*tree.root.Left.Left, *node, t)

}

func TestBinaryTreePreOrderTraversel(t *testing.T) {
	root := getSampleTree()
	traverel := root.PreOrderTraversel()
	expected := []int{1, 2, 4, 5, 8, 3, 6, 7}
	if !isEqual(expected, traverel) {
		t.Errorf("traversel did not matched. expected: %v, got %v", expected, traverel)
	}
}

func TestBinaryTreePostOrderTraversel(t *testing.T) {
	root := getSampleTree()
	traverel := root.PostOrderTraversel()
	expected := []int{4, 8, 5, 2, 6, 7, 3, 1}
	if !isEqual(expected, traverel) {
		t.Errorf("traversel did not matched. expected: %v, got %v", expected, traverel)
	}
}

func TestBinaryTreeInOrderTraversel(t *testing.T) {
	root := getSampleTree()
	traverel := root.InOrderTraversel()
	expected := []int{4, 2, 8, 5, 1, 6, 3, 7}
	if !isEqual(expected, traverel) {
		t.Errorf("traversel did not matched. expected: %v, got %v", expected, traverel)
	}
}

func TestBinaryTreeLevelOrderTraversel(t *testing.T) {
	tree := getSampleTree()
	traversel := tree.LevelOrderTraversel()

	if len(traversel) != 4 {
		t.Fatalf("level size did not matched. expected: 4, got: %d", len(traversel))
	}

	if !isEqual([]int{1}, traversel[0]) {
		t.Errorf("traversel did not matched. expected: [1], got %v", traversel[0])
	}

	if !isEqual([]int{2, 3}, traversel[1]) {
		t.Errorf("traversel did not matched. expected: [2 3], got %v", traversel[1])
	}

	if !isEqual([]int{4, 5, 6, 7}, traversel[2]) {
		t.Errorf("traversel did not matched. expected: [4 5 6 7], got %v", traversel[2])
	}

	if !isEqual([]int{8}, traversel[3]) {
		t.Errorf("traversel did not matched. expected: [8], got %v", traversel[3])
	}

}

func TestBinaryTreeSize(t *testing.T) {
	var tree BinaryTree
	size := tree.Size()
	if size != 0 {
		t.Errorf("Size did not matched. expected: 0, got: %d", size)
	}
	tree.root = &BinaryNode{Data: 1}
	size = tree.Size()
	if size != 1 {
		t.Errorf("Size did not matched. expected: 1, got: %d", size)
	}

	tree.root.Left = &BinaryNode{Data: 2}
	size = tree.Size()
	if size != 2 {
		t.Errorf("Size did not matched. expected: 2, got: %d", size)
	}

}

func nodeAssert(expected BinaryNode, got BinaryNode, t *testing.T) {
	if expected.Data != got.Data || expected.Left != got.Left || expected.Right != got.Right {
		t.Errorf("Search data did not matched. expected: %s, got %s", &expected, &got)
	}
}

func isEqual(expected []int, got []int) bool {

	if (expected == nil) != (got == nil) {
		return false
	}

	if len(expected) != len(got) {
		return false
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != got[i] {
			return false
		}
	}

	return true
}

func getSampleTree() BinaryTree {
	root := BinaryNode{Data: 1}
	root.Left = &BinaryNode{Data: 2}
	root.Right = &BinaryNode{Data: 3}
	root.Left.Left = &BinaryNode{Data: 4}
	root.Left.Right = &BinaryNode{Data: 5}
	root.Right.Left = &BinaryNode{Data: 6}
	root.Right.Right = &BinaryNode{Data: 7}
	root.Left.Right.Left = &BinaryNode{Data: 8}
	// fmt.Printf("\t\t%d\t\t\n", root.Data)
	// fmt.Printf("\t%d\t\t%d\t\n", root.Left.Data, root.Right.Data)
	// fmt.Printf("    %d\t    %d\t    %d\t\t%d\n", root.Left.Left.Data, root.Left.Right.Data,
	// 	root.Right.Left.Data, root.Right.Right.Data)
	// fmt.Printf("\t   %d\n", root.Left.Right.Left.Data)
	return BinaryTree{root: &root}
}
