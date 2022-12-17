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

	// Print the tree in pre-order
	fmt.Println("PreOrder:")
	tree.PreOrder()

	// Print the tree in in-order
	fmt.Println("InOrder:")
	tree.InOrder()

	// Print the tree in post-order
	fmt.Println("PostOrder:")
	tree.PostOrder()

	// Print the tree in level-order
	fmt.Println("LevelOrder:")
	tree.LevelOrder()

	// Print the tree in reverse-level-order
	fmt.Println("ReverseLevelOrder:")
	tree.ReverseLevelOrder()

	// Print the tree in spiral-order
	fmt.Println("SpiralOrder:")
	tree.SpiralOrder()

	// Print the tree in reverse-spiral-order
	fmt.Println("ReverseSpiralOrder:")
	tree.ReverseSpiralOrder()

	// Print the tree in zigzag-order
	fmt.Println("ZigZagOrder:")
	tree.ZigZagOrder()

	// Print the tree in reverse-zigzag-order
	fmt.Println("ReverseZigZagOrder:")
	tree.ReverseZigZagOrder()

	// Print the tree in vertical-order
	fmt.Println("VerticalOrder:")
	tree.VerticalOrder()

	// Print the tree in reverse-vertical-order
	fmt.Println("ReverseVerticalOrder:")
	tree.ReverseVerticalOrder()

	// Print the tree in diagonal-order
	fmt.Println("DiagonalOrder:")
	tree.DiagonalOrder()

	// Print the tree in reverse-diagonal-order
	fmt.Println("ReverseDiagonalOrder:")
	tree.ReverseDiagonalOrder()

	// Print the tree in boundary-order
	fmt.Println("BoundaryOrder:")
	tree.BoundaryOrder()

	// Print the tree in reverse-boundary-order
	tree.ReverseBoundaryOrder()

	// Print the tree in left-boundary-order
	tree.LeftBoundaryOrder()

	// Print the tree in right-boundary-order
	tree.RightBoundaryOrder()

	// Print the tree in reverse-left-boundary-order
	tree.ReverseLeftBoundaryOrder()

	// Print the tree in reverse-right-boundary-order
	tree.ReverseRightBoundaryOrder()

	// Print the tree in left-view-order
	tree.LeftViewOrder()

	// Print the tree in right-view-order
	tree.RightViewOrder()

	// Print the tree in reverse-left-view-order
	tree.ReverseLeftViewOrder()

	// Print the tree in reverse-right-view-order
	tree.ReverseRightViewOrder()

	// Print the tree in top-view-order
	tree.TopViewOrder()

	// Print the tree in bottom-view-order
	tree.BottomViewOrder()

	// Print the tree in reverse-top-view-order
	tree.ReverseTopViewOrder()

	// Print the tree in reverse-bottom-view-order
	tree.ReverseBottomViewOrder()

	// Print the tree in left-cliff-order
	tree.LeftCliffOrder()

	// Print the tree in right-cliff-order
	tree.RightCliffOrder()

	// Print the tree in reverse-left-cliff-order
	tree.ReverseLeftCliffOrder()

	// Print the tree in reverse-right-cliff-order
	tree.ReverseRightCliffOrder()
}
