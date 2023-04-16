package binaryTree

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func New() BinaryTree {
	return BinaryTree{}
}

func (bt *BinaryTree) Add(data int) {
	inputNode := &Node{
		data: data,
	}

	if bt.root == nil {
		bt.root = inputNode
		return
	}

	currentNode := bt.root
	for {
		if data >= currentNode.data {
			if currentNode.right == nil {
				currentNode.right = inputNode
				return
			}
			currentNode = currentNode.right
		} else {
			if currentNode.left == nil {
				currentNode.left = inputNode
				return
			}

			currentNode = currentNode.left
		}
	}
}

func (bt *BinaryTree) List() {
	fmt.Println("*** Pre Order ***")
	ListPreOrder(bt.root)
	fmt.Println()
	fmt.Println("*** Post Order ***")
	ListPostOrder(bt.root)
	fmt.Println()
	fmt.Println("*** In Order ***")
	ListInOrder(bt.root)
	fmt.Println()
	fmt.Println("*** Level Order ***")
	ListLevelOrder(bt.root)
	fmt.Println()
}
func ListPreOrder(focusNode *Node) {
	if focusNode == nil {
		return
	}

	fmt.Printf("%d ", focusNode.data)
	ListPreOrder(focusNode.left)
	ListPreOrder(focusNode.right)

}

func ListPostOrder(focusNode *Node) {
	if focusNode == nil {
		return
	}

	ListPreOrder(focusNode.left)
	ListPreOrder(focusNode.right)
	fmt.Printf("%d ", focusNode.data)
}

func ListInOrder(focusNode *Node) {
	if focusNode == nil {
		return
	}

	ListPreOrder(focusNode.left)
	fmt.Printf("%d ", focusNode.data)
	ListPreOrder(focusNode.right)
}

func ListLevelOrder(focusNode *Node) {
	h := height(focusNode)

	for i := 1; i <= h; i++ {
		printCurrentLevel(focusNode, i)
		fmt.Println()
	}
}

func height(node *Node) int {
	if node == nil {
		return 0
	}

	lheight := height(node.left)
	rheight := height(node.right)
	if lheight > rheight {
		return lheight + 1
	} else {
		return rheight + 1
	}
}
func printCurrentLevel(root *Node, level int) {
	if root == nil {
		return
	}

	if level == 1 {
		fmt.Printf("%d ", root.data)
		return
	}

	if level > 1 {
		printCurrentLevel(root.left, level-1)
		printCurrentLevel(root.right, level-1)
	}
}

func (bt *BinaryTree) Search() {

}

func BFS(focusNode *Node) {
	fmt.Print(focusNode.data)
	ListPreOrder(focusNode.left)
	ListPreOrder(focusNode.right)
}

func DFS(focusNode *Node) {
	fmt.Print(focusNode.data)
	ListPreOrder(focusNode.left)
	ListPreOrder(focusNode.right)
}

func (bt *BinaryTree) Remove() {

}

/*
Pre-order Traversal of the above tree: 1-2-4-5-3-6-7
In-order Traversal of the above tree: 4-2-5-1-6-3-7
Post-order Traversal of the above tree: 4-5-2-6-7-3-1
Level-order Traversal of the above tree: 1-2-3-4-5-6-7
*/
