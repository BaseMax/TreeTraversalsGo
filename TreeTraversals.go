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
	tree.Print()
	fmt.Println(tree.String())

	// Print the tree in pre-order
	fmt.Println("PreOrder:")
	tree.PreOrder()
	fmt.Println(tree.StringPreOrder())

	// Print the tree in in-order
	fmt.Println("InOrder:")
	tree.InOrder()
	fmt.Println(tree.StringInOrder())

	// Print the tree in post-order
	fmt.Println("PostOrder:")
	tree.PostOrder()
	fmt.Println(tree.StringPostOrder())

	// Print the tree in level-order
	fmt.Println("LevelOrder:")
	tree.LevelOrder()
	fmt.Println(tree.StringLevelOrder())

	// Print the tree in reverse-level-order
	fmt.Println("ReverseLevelOrder:")
	tree.ReverseLevelOrder()
	fmt.Println(tree.StringReverseLevelOrder())

	// // Print the tree
	// fmt.Println("Tree:")
	// tree.Print()

	// // Print the tree in pre-order
	// fmt.Println("PreOrder:")
	// tree.PreOrder()

	// // Print the tree in in-order
	// fmt.Println("InOrder:")
	// tree.InOrder()

	// // Print the tree in post-order
	// fmt.Println("PostOrder:")
	// tree.PostOrder()

	// // Print the tree in level-order
	// fmt.Println("LevelOrder:")
	// tree.LevelOrder()

	// // Print the tree in reverse-level-order
	// fmt.Println("ReverseLevelOrder:")
	// tree.ReverseLevelOrder()

	// // Print the tree in spiral-order
	// fmt.Println("SpiralOrder:")
	// tree.SpiralOrder()

	// // Print the tree in reverse-spiral-order
	// fmt.Println("ReverseSpiralOrder:")
	// tree.ReverseSpiralOrder()

	// // Print the tree in zigzag-order
	// fmt.Println("ZigZagOrder:")
	// tree.ZigZagOrder()

	// // Print the tree in reverse-zigzag-order
	// fmt.Println("ReverseZigZagOrder:")
	// tree.ReverseZigZagOrder()

	// // Print the tree in vertical-order
	// fmt.Println("VerticalOrder:")
	// tree.VerticalOrder()

	// // Print the tree in reverse-vertical-order
	// fmt.Println("ReverseVerticalOrder:")
	// tree.ReverseVerticalOrder()

	// // Print the tree in diagonal-order
	// fmt.Println("DiagonalOrder:")
	// tree.DiagonalOrder()

	// // Print the tree in reverse-diagonal-order
	// fmt.Println("ReverseDiagonalOrder:")
	// tree.ReverseDiagonalOrder()

	// // Print the tree in boundary-order
	// fmt.Println("BoundaryOrder:")
	// tree.BoundaryOrder()

	// // Print the tree in reverse-boundary-order
	// fmt.Println("ReverseBoundaryOrder:")
	// tree.ReverseBoundaryOrder()

	// // Print the tree in left-boundary-order
	// fmt.Println("LeftBoundaryOrder:")
	// tree.LeftBoundaryOrder()

	// // Print the tree in right-boundary-order
	// fmt.Println("RightBoundaryOrder:")
	// tree.RightBoundaryOrder()

	// // Print the tree in reverse-left-boundary-order
	// fmt.Println("ReverseLeftBoundaryOrder:")
	// tree.ReverseLeftBoundaryOrder()

	// // Print the tree in reverse-right-boundary-order
	// fmt.Println("ReverseRightBoundaryOrder:")
	// tree.ReverseRightBoundaryOrder()

	// // Print the tree in left-view-order
	// fmt.Println("LeftViewOrder:")
	// tree.LeftViewOrder()

	// // Print the tree in right-view-order
	// fmt.Println("RightViewOrder:")
	// tree.RightViewOrder()

	// // Print the tree in reverse-left-view-order
	// fmt.Println("ReverseLeftViewOrder:")
	// tree.ReverseLeftViewOrder()

	// // Print the tree in reverse-right-view-order
	// fmt.Println("ReverseRightViewOrder:")
	// tree.ReverseRightViewOrder()

	// // Print the tree in top-view-order
	// fmt.Println("TopViewOrder:")
	// tree.TopViewOrder()

	// // Print the tree in bottom-view-order
	// fmt.Println("BottomViewOrder:")
	// tree.BottomViewOrder()

	// // Print the tree in reverse-top-view-order
	// fmt.Println("ReverseTopViewOrder:")
	// tree.ReverseTopViewOrder()

	// // Print the tree in reverse-bottom-view-order
	// fmt.Println("ReverseBottomViewOrder:")
	// tree.ReverseBottomViewOrder()

	// // Print the tree in left-cliff-order
	// fmt.Println("LeftCliffOrder:")
	// tree.LeftCliffOrder()

	// // Print the tree in right-cliff-order
	// fmt.Println("RightCliffOrder:")
	// tree.RightCliffOrder()

	// // Print the tree in reverse-left-cliff-order
	// fmt.Println("ReverseLeftCliffOrder:")
	// tree.ReverseLeftCliffOrder()

	// // Print the tree in reverse-right-cliff-order
	// fmt.Println("ReverseRightCliffOrder:")
	// tree.ReverseRightCliffOrder()
}
