package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// BinarySearchTree 二叉树
type BinarySearchTree struct {
	//根节点
	Root *BinarySearchTreeNode
}

// BinarySearchTreeNode 节点
type BinarySearchTreeNode struct {
	Value int64                 //具体值
	Times int64                 //重复次数
	Left  *BinarySearchTreeNode //左节点
	Right *BinarySearchTreeNode //右节点
}

// NewBinarySearchTree 初始化
func NewBinarySearchTree() *BinarySearchTree {
	return new(BinarySearchTree) //
}

// Add is a method of the BinarySearchTree struct that adds a new node with the given value to the tree.
func (tree *BinarySearchTree) Add(value int64) {
	// Check if the tree is empty
	if tree.Root == nil {
		// If empty, create a new node with the given value and set it as the root
		tree.Root = &BinarySearchTreeNode{Value: value}
		return
	}

	// If the tree is not empty, call the Add method of the root node to add the new value
	tree.Root.Add(value)
}

// Add 子节点
func (node *BinarySearchTreeNode) Add(value int64) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &BinarySearchTreeNode{
				Value: value,
			}
		} else {
			node.Left.Add(value) //递归处理
		}
	} else if value > node.Value {
		if node.Right == nil {
			node.Right = &BinarySearchTreeNode{Value: value}
		} else {
			node.Right.Add(value)
		}
	} else {
		//重复的值
		node.Times = node.Times + 1
	}
}

func (tree *BinarySearchTree) Search(value int64) *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.Search(value)
}
func (node *BinarySearchTreeNode) Search(value int64) *BinarySearchTreeNode {
	//找到了
	if value == node.Value {
		return node
	}
	if value < node.Value {
		if node.Left == nil {
			return nil
		}
		return node.Left.Search(value)
	} else {
		if node.Right == nil {
			return nil
		}
		return node.Right.Search(value)
	}
}

func (tree *BinarySearchTree) FindMax() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMax()
}
func (node *BinarySearchTreeNode) FindMax() *BinarySearchTreeNode {
	if node.Right == nil {
		return node
	}
	return node.Right.FindMax()
}
func (tree *BinarySearchTree) FindMin() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMax()
}
func (node *BinarySearchTreeNode) Findmin() *BinarySearchTreeNode {
	if node.Left == nil {
		return node
	}
	return node.Left.Findmin()
}

//遍历树
//先序遍历
//中序遍历
//后序遍历

func (tree *BinarySearchTree) PreOrder() []int64 {
	var res []int64
	if tree.Root == nil {
		return res
	}
	tree.Root.preOrder(&res)
	return res
}

func (node *BinarySearchTreeNode) preOrder(res *[]int64) {
	if node == nil {
		return
	}
	*res = append(*res, node.Value)
	node.Left.preOrder(res)
	node.Right.preOrder(res)
}

func GetTree() *BinarySearchTree {
	b := NewBinarySearchTree()
	b.Add(5)
	b.Add(3)
	b.Add(7)
	b.Add(2)
	b.Add(4)
	return b
}
func TestTree_Add_Search(t *testing.T) {
	b := GetTree()
	if b.Search(5) == nil {
		t.Errorf("BinarySearchTree.Add(5) faild,expected:%v :go %v",
			&BinarySearchTreeNode{Value: 5}, nil)
	}
	if b.Search(3) == nil {
		t.Errorf("BinarySearchTree.Add(3) faild,expected:%v :go %v",
			&BinarySearchTreeNode{Value: 3}, nil)
	}
	if b.Search(7) == nil {
		t.Errorf("BinarySearchTree.Add(7) faild,expected:%v :go %v",
			&BinarySearchTreeNode{Value: 7}, nil)
	}
	if b.Search(2) == nil {
		t.Errorf("BinarySearchTree.Add(2) faild,expected:%v :go %v",
			&BinarySearchTreeNode{Value: 2}, nil)
	}
	if b.Search(4) == nil {
		t.Errorf("BinarySearchTree.Add(4) faild,expected:%v :go %v",
			&BinarySearchTreeNode{Value: 4}, nil)
	}
}
func TestTree_Max_Min(t *testing.T) {
	b := NewBinarySearchTree()
	b.Add(5)
	b.Add(3)
	b.Add(7)
	b.Add(2)
	b.Add(4)
	if b.FindMax() == nil {
		t.Errorf("BinarySearchTree.Add(7) faild,expected:%v :go %v",
			&BinarySearchTreeNode{Value: 7}, nil)
	}
	if b.FindMin() == nil {
		t.Errorf("BinarySearchTree.Add(2) faild,expected:%v :go %v",
			&BinarySearchTreeNode{Value: 2}, nil)
	}

	bnil := NewBinarySearchTree()
	assert.Equal(t, bnil.FindMax(), bnil.FindMin())
	n := bnil.FindMax()
	assert.Nil(t, n) //对比nil
}

func TestTree_PreOder(t *testing.T) {
	b := GetTree()
	res := b.PreOrder()
	fmt.Println(res)
}

func (tree *BinarySearchTree) GetDepth() int {
	if tree.Root == nil {
		return 0
	}
	return tree.Root.GetDepth()
}
func (node *BinarySearchTreeNode) GetDepth() int {
	if node == nil {
		return 0
	}
	leftDepth := node.Left.GetDepth()
	rightDepth := node.Right.GetDepth()
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

func TestTree_GetDepth(t *testing.T) {
	b := GetTree()
	d := b.GetDepth()
	fmt.Println(d)
}
