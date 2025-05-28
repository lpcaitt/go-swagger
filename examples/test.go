
package main

import "fmt"

// TreeNode 定义二叉树的节点
type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

// BinaryTree 定义二叉树
type BinaryTree struct {
	root *TreeNode
}

// Insert 向二叉树中插入一个新节点
func (tree *BinaryTree) Insert(value int) {
	tree.root = insertRec(tree.root, value)
}

// insertRec 是递归插入节点的辅助函数
func insertRec(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return &TreeNode{value: value}
	}
	if value < node.value {
		node.left = insertRec(node.left, value)
	} else if value > node.value {
		node.right = insertRec(node.right, value)
	}
	return node
}

// PreOrder 前序遍历
func (tree *BinaryTree) PreOrder() {
	preOrderRec(tree.root)
	fmt.Println()
}

// preOrderRec 是前序遍历的递归辅助函数
func preOrderRec(node *TreeNode) {
	if node != nil {
		fmt.Print(node.value, " ")
		preOrderRec(node.left)
		preOrderRec(node.right)
	}
}

// InOrder 中序遍历
func (tree *BinaryTree) InOrder() {
	inOrderRec(tree.root)
	fmt.Println()
}

// inOrderRec 是中序遍历的递归辅助函数
func inOrderRec(node *TreeNode) {
	if node != nil {
		inOrderRec(node.left)
		fmt.Print(node.value, " ")
		inOrderRec(node.right)
	}
}

// PostOrder 后序遍历
func (tree *BinaryTree) PostOrder() {
	postOrderRec(tree.root)
	fmt.Println()
}

// postOrderRec 是后序遍历的递归辅助函数
func postOrderRec(node *TreeNode) {
	if node != nil {
		postOrderRec(node.left)
		postOrderRec(node.right)
		fmt.Print(node.value, " ")
	}
}

// Search 搜索值是否在二叉树中
func (tree *BinaryTree) Search(value int) bool {
	return searchRec(tree.root, value)
}

// searchRec 是递归搜索值的辅助函数
func searchRec(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}
	if node.value == value {
		return true
	}
	if value < node.value {
		return searchRec(node.left, value)
	}
	return searchRec(node.right, value)
}

// 主函数
func main() {
	tree := &BinaryTree{}

	// 插入节点
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(70)
	tree.Insert(60)
	tree.Insert(80)

	// 遍历
	fmt.Print("前序遍历: ")
	tree.PreOrder()

	fmt.Print("中序遍历: ")
	tree.InOrder()

	fmt.Print("后序遍历: ")
	tree.PostOrder()

	// 搜索
	fmt.Println("搜索 40:", tree.Search(40)) // 应该返回 true
	fmt.Println("搜索 90:", tree.Search(90)) // 应该返回 false
}
```

### 代码详解
- **TreeNode** 结构体：表示二叉树的节点，包含节点值、左子节点和右子节点。
- **BinaryTree** 结构体：表示整个二叉树，包含一个指向根节点的指针。
- **Insert** 方法：向树中插入新节点，使用递归实现。
- **PreOrder**, **InOrder**, **PostOrder** 方法：分别实现前序、中序和后序遍历，遍历的具体实现也采用递归辅助函数。
- **Search** 方法：在树中查找特定值，返回值是否存在。

### 如何运行
将代码保存到一个名为 `binary_tree.go` 的文件中，您可以使用以下命令运行：
```bash
go run binary_tree.go
