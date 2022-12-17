/**
 *
 * @Name: Tree Traversals
 * @Author: Max Base
 * @Date: 2022/12/17
 * @Repository: https://github.com/BaseMax/TreeTraversalsGo
 * @License: GPL-3.0
 * @Description: Tree Traversals in Go
 **/

package main

import "fmt"

type Node struct {
	Value interface{}
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (tree *BinaryTree) Print() {
	tree.Root.Print()
}

func (node *Node) Print() {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
	node.Left.Print()
	node.Right.Print()
}

func (tree *BinaryTree) String() string {
	return tree.Root.String()
}

func (node *Node) String() string {
	if node == nil {
		return ""
	}
	return fmt.Sprintf("%v, ", node.Value) + node.Left.String() + node.Right.String()
}

func (tree *BinaryTree) PreOrder() {
	tree.Root.PreOrder()
}

func (node *Node) PreOrder() {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
	node.Left.PreOrder()
	node.Right.PreOrder()
}

func (tree *BinaryTree) InOrder() {
	tree.Root.InOrder()
}

func (node *Node) InOrder() {
	if node == nil {
		return
	}
	node.Left.InOrder()
	fmt.Println(node.Value)
	node.Right.InOrder()
}

func (tree *BinaryTree) PostOrder() {
	tree.Root.PostOrder()
}

func (node *Node) PostOrder() {
	if node == nil {
		return
	}
	node.Left.PostOrder()
	node.Right.PostOrder()
	fmt.Println(node.Value)
}

func (tree *BinaryTree) LevelOrder() {
	tree.Root.LevelOrder()
}

func (node *Node) LevelOrder() {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
	node.Left.LevelOrder()
	node.Right.LevelOrder()
}

func (tree *BinaryTree) ReverseLevelOrder() {
	tree.Root.ReverseLevelOrder()
}

func (node *Node) ReverseLevelOrder() {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
	node.Right.ReverseLevelOrder()
	node.Left.ReverseLevelOrder()
}

func (tree *BinaryTree) StringInOrder() string {
	return tree.Root.StringInOrder()
}

func (node *Node) StringInOrder() string {
	if node == nil {
		return ""
	}
	return node.Left.StringInOrder() + fmt.Sprintf("%v, ", node.Value) + node.Right.StringInOrder()
}

func (tree *BinaryTree) StringPreOrder() string {
	return tree.Root.StringPreOrder()
}

func (node *Node) StringPreOrder() string {
	if node == nil {
		return ""
	}
	return fmt.Sprintf("%v, ", node.Value) + node.Left.StringPreOrder() + node.Right.StringPreOrder()
}

func (tree *BinaryTree) StringPostOrder() string {
	return tree.Root.StringPostOrder()
}

func (node *Node) StringPostOrder() string {
	if node == nil {
		return ""
	}
	return node.Left.StringPostOrder() + node.Right.StringPostOrder() + fmt.Sprintf("%v, ", node.Value)
}

func (tree *BinaryTree) StringLevelOrder() string {
	return tree.Root.StringLevelOrder()
}

func (node *Node) StringLevelOrder() string {
	if node == nil {
		return ""
	}
	return fmt.Sprintf("%v, ", node.Value) + node.Left.StringLevelOrder() + node.Right.StringLevelOrder()
}

func (tree *BinaryTree) StringReverseLevelOrder() string {
	return tree.Root.StringReverseLevelOrder()
}

func (node *Node) StringReverseLevelOrder() string {
	if node == nil {
		return ""
	}
	return node.Right.StringReverseLevelOrder() + node.Left.StringReverseLevelOrder() + fmt.Sprintf("%v, ", node.Value)
}

func main() {
	// Create a large general-tree
	tree := BinaryTree{
		Root: &Node{
			Value: 1,
			Left: &Node{
				Value: 2,
				Left: &Node{
					Value: 4,
					Left: &Node{
						Value: 8,
					},
					Right: &Node{
						Value: 9,
					},
				},
				Right: &Node{
					Value: 5,
					Left: &Node{
						Value: 10,
					},
					Right: &Node{
						Value: 11,
					},
				},
			},
			Right: &Node{
				Value: 3,
				Right: &Node{
					Value: 7,
					Left: &Node{
						Value: 12,
					},
					Right: &Node{
						Value: 13,
					},
				},
			},
		},
	}

	// Print the tree
	fmt.Println("Tree:")
	// tree.Print()
	fmt.Println(tree.String())

	// Print the tree in pre-order
	fmt.Println("PreOrder:")
	// tree.PreOrder()
	fmt.Println(tree.StringPreOrder())

	// Print the tree in in-order
	fmt.Println("InOrder:")
	// tree.InOrder()
	fmt.Println(tree.StringInOrder())

	// Print the tree in post-order
	fmt.Println("PostOrder:")
	// tree.PostOrder()
	fmt.Println(tree.StringPostOrder())

	// Print the tree in level-order
	fmt.Println("LevelOrder:")
	// tree.LevelOrder()
	fmt.Println(tree.StringLevelOrder())

	// Print the tree in reverse-level-order
	fmt.Println("ReverseLevelOrder:")
	// tree.ReverseLevelOrder()
	fmt.Println(tree.StringReverseLevelOrder())

	// Print the tree in spiral-order
	// Print the tree in reverse-spiral-order
	// Print the tree in zigzag-order
	// Print the tree in reverse-zigzag-order
	// Print the tree in vertical-order
	// Print the tree in reverse-vertical-order
	// Print the tree in diagonal-order
	// Print the tree in reverse-diagonal-order
	// Print the tree in boundary-order
	// Print the tree in reverse-boundary-order
	// Print the tree in left-boundary-order
	// Print the tree in right-boundary-order
	// Print the tree in reverse-left-boundary-order
	// Print the tree in reverse-right-boundary-order
	// Print the tree in left-view-order
	// Print the tree in right-view-order
	// Print the tree in reverse-left-view-order
	// Print the tree in reverse-right-view-order
	// Print the tree in top-view-order
	// Print the tree in bottom-view-order
	// Print the tree in reverse-top-view-order
	// Print the tree in reverse-bottom-view-order
	// Print the tree in left-cliff-order
	// Print the tree in right-cliff-order
	// Print the tree in reverse-left-cliff-order
	// Print the tree in reverse-right-cliff-order
}
