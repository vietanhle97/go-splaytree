package main

import (
	"fmt"
	. "go-splaytree/splay"
)

func preOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preOrder(root.Left, res)
	preOrder(root.Right, res)
}

func insertSplayTree(root *TreeNode, val []int) *TreeNode {
	for _,e := range val {
		if root == nil {
			root = &TreeNode{Val: e}
		} else {
			root = root.Insert(e)
		}
	}
	return root
}

func main() {
	val := []int{1,3,5,6,3,7,8,9,5,3,4,6,7,8,5,4,3,8,9,4}
	res := make([]int, 0)
	root := insertSplayTree(nil, val)
	preOrder(root, &res)
	fmt.Println(res)
	newRoot := root.Search(9)
	if newRoot != nil {
		fmt.Println(newRoot.Val)
	} else {
		fmt.Println(newRoot)
	}
	res = make([]int, 0)
	preOrder(newRoot, &res)
	fmt.Println(res)
}
