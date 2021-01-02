package tree

import "fmt"

//BinaryNode implementation
type BinaryNode struct {
	Data  int
	Left  *BinaryNode
	Right *BinaryNode
}

//BinaryTree implentation
type BinaryTree struct {
	root *BinaryNode
}

//Search element into the Binary tree
func (tree *BinaryTree) Search(data int) *BinaryNode {
	return Search(tree.root, data)
}

//PreOrderTraversel pre-order traversel of the Binary tree
func (tree *BinaryTree) PreOrderTraversel() []int {
	// data - left - right
	var traversel []int
	PreOrder(tree.root, &traversel)
	return traversel
}

//InOrderTraversel in-order traversel of the Binary tree
func (tree *BinaryTree) InOrderTraversel() []int {
	// data - left - right
	var traversel []int
	InOrder(tree.root, &traversel)
	return traversel
}

//PostOrderTraversel post-order traversel of the Binary tree
func (tree *BinaryTree) PostOrderTraversel() []int {
	// data - left - right
	var traversel []int
	PostOrder(tree.root, &traversel)
	return traversel
}

//LevelOrderTraversel level order traversel of the Binary tree
func (tree *BinaryTree) LevelOrderTraversel() [][]int {
	return LevelOrder(tree.root)
}

//Height returns the height of the tree
func (tree *BinaryTree) Height() int {
	return len(tree.LevelOrderTraversel())
}

//Size returns the size of the tree
func (tree *BinaryTree) Size() int {
	return Size(tree.root)
}

// Search BinaryNode
func Search(root *BinaryNode, data int) *BinaryNode {
	if root == nil {
		return root
	}
	if root.Data == data {
		return root
	}
	var node *BinaryNode
	if root.Left != nil {
		node = Search(root.Left, data)
	}
	if node == nil && root.Right != nil {
		node = Search(root.Right, data)
	}
	return node
}

//PreOrder pre-order traversel of the Binary tree
func PreOrder(root *BinaryNode, arr *[]int) {
	if root != nil {
		*arr = append(*arr, root.Data)
		PreOrder(root.Left, arr)
		PreOrder(root.Right, arr)
	}
}

//InOrder in-order traversel of the Binary tree
func InOrder(root *BinaryNode, arr *[]int) {
	if root != nil {
		InOrder(root.Left, arr)
		*arr = append(*arr, root.Data)
		InOrder(root.Right, arr)
	}
}

//PostOrder post-order traversel of the Binary tree
func PostOrder(root *BinaryNode, arr *[]int) {
	if root != nil {
		PostOrder(root.Left, arr)
		PostOrder(root.Right, arr)
		*arr = append(*arr, root.Data)
	}
}

//LevelOrder Traversel
func LevelOrder(root *BinaryNode) [][]int {
	levels := make([][]int, 0)
	currentLevel := make([]*BinaryNode, 0)
	currentLevel = append(currentLevel, root)
	for len(currentLevel) != 0 {
		n := len(currentLevel)
		temp := make([]int, n)
		for i, node := range currentLevel {
			temp[i] = node.Data
		}
		levels = append(levels, temp)
		nextLevel := make([]*BinaryNode, 0)
		for _, node := range currentLevel {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		currentLevel = nextLevel
	}
	return levels
}

//Size get size from the node
func Size(root *BinaryNode) int {
	if root != nil {
		return Size(root.Left) + Size(root.Right) + 1
	}
	return 0

}

func (root *BinaryNode) String() string {
	if root == nil {
		return fmt.Sprintf("BN{ <nil>}")
	}
	if root.Left == nil && root.Right != nil {
		return fmt.Sprintf("BN{ d: %d, l: <nil>, r: %d}", root.Data, root.Right.Data)
	}
	if root.Left != nil && root.Right == nil {
		return fmt.Sprintf("BN{ d: %d, l: %d, r: <nil>}", root.Data, root.Left.Data)
	}
	if root.Left == nil && root.Right == nil {
		return fmt.Sprintf("BN{ d: %d, l: <nil>, r: <nil>}", root.Data)
	}
	return fmt.Sprintf("BN{ d: %d, l: %d, r: %d}", root.Data, root.Left.Data, root.Right.Data)
}
