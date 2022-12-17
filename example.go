package main

import "fmt"

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
