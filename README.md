# Tree-Traversals Go

This is a simple Go program that traverses a tree in different ways. (Pre-order, In-order, Post-order, and Level-order)

## Methods

- Pre-order
- In-order
- Post-order
- Level-order
- Reverse level-order

## Output

**InOrder:**

```
8, 4, 9, 2, 10, 5, 11, 1, 3, 12, 7, 13
```

**PreOrder:**

```
1, 2, 4, 8, 9, 5, 10, 11, 3, 7, 12, 13
```

**PostOrder:**

```
8, 9, 4, 10, 11, 5, 2, 12, 13, 7, 3, 1
```

**LevelOrder:**

```
1, 2, 4, 8, 9, 5, 10, 11, 3, 7, 12, 13
```

**ReverseLevelOrder:**

```
13, 12, 7, 3, 11, 10, 5, 9, 8, 4, 2, 1
```

## Usage

```go
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
```

## License

This project is licensed under the GPL-3.0 License - see the [LICENSE.md](LICENSE.md) file for details.

Copyright (c) 2022, Max Base
