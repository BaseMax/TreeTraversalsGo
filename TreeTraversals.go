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
