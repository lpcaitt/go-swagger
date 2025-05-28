以下是用 Go 语言实现的一个简单的二叉树算法，包括插入、遍历（前序、中序、后序）和搜索操作：

```go
package main

import "fmt"

// 定义二叉树节点
type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

// 定义二叉树
type BinaryTree struct {
	root *TreeNode
}

// 插入节点
func (tree *BinaryTree) Insert(value int) {
	tree.root = insertRec(tree.root, value)
}

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

// 前序遍历
func (tree *BinaryTree) PreOrder() {
	preOrderRec(tree.root)
	fmt.Println()
}

func preOrderRec(node *TreeNode) {
	if node != nil {
		fmt.Print(node.value, " ")
		preOrderRec(node.left)
		preOrderRec(node.right)
	}
}

// 中序遍历
func (tree *BinaryTree) InOrder() {
	inOrderRec(tree.root)
	fmt.Println()
}

func inOrderRec(node *TreeNode) {
	if node != nil {
		inOrderRec(node.left)
		fmt.Print(node.value, " ")
		inOrderRec(node.right)
	}
}

// 后序遍历
func (tree *BinaryTree) PostOrder() {
	postOrderRec(tree.root)
	fmt.Println()
}

func postOrderRec(node *TreeNode) {
	if node != nil {
		postOrderRec(node.left)
		postOrderRec(node.right)
		fmt.Print(node.value, " ")
	}
}

// 搜索节点
func (tree *BinaryTree) Search(value int) bool {
	return searchRec(tree.root, value)
}

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

// 主函数，用于测试二叉树
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
	fmt.Println("搜索 40:", tree.Search(40))
	fmt.Println("搜索 90:", tree.Search(90))
}
```

### 功能说明
1. **插入节点**：通过递归方式将节点插入到适当的位置。
2. **遍历**：提供前序、后序和中序遍历的实现。
3. **搜索功能**：返回在树中是否找到指定的值。

### 如何运行
将上面的代码复制到一个 `.go` 文件中，比如 `binary_tree.go`，然后在终端中运行以下命令：
```bash
go run binary_tree.go
```

这将会执行这个程序并输出遍历结果，以及搜索指定节点的结果。如果你有更多问题，或者需要进一步的帮助，请告诉我！