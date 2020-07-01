package splay

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rotateLeft(node *TreeNode) *TreeNode {
	newRoot := node.Right
	node.Right = newRoot.Left
	newRoot.Left = node
	return newRoot
}

func  rotateRight(node *TreeNode) *TreeNode {
	newRoot := node.Left
	node.Left = newRoot.Right
	newRoot.Right = node
	return newRoot
}

func splay(node *TreeNode, val int) *TreeNode {
	if node == nil || node.Val == val  {
		return node
	}
	if node.Val > val {
		if node.Left == nil {
			return node
		}
		if node.Left.Val > val {
			node.Left.Left = splay(node.Left.Left, val)
			node = rotateRight(node)
		} else if node.Left.Val < val {
			node.Left.Right = splay(node.Left.Right,  val)
			if node.Left.Right != nil {
				node.Left =rotateLeft( node.Left)
			}
		}
		if node.Left == nil {
			return node
		}
		return rotateRight(node)
	} else {
		if node.Right == nil {
			return node
		}
		if node.Right.Val < val {
			node.Right.Right = splay(node.Right.Right, val)
			node = rotateLeft(node)
		} else if node.Right.Val > val {
			node.Right.Left = splay(node.Right.Left,val)
			if node.Right.Left != nil {
				node.Right = rotateRight(node.Right)
			}
		}
		if node.Right == nil {
			return node
		}
		return rotateLeft(node)
	}
}

func (node *TreeNode) Insert(val int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: val}
	}
	root := splay(node, val)
	if root.Val == val {
		return root
	}
	newRoot := &TreeNode{Val: val}
	if root.Val < val {
		newRoot.Right = root.Right
		root.Right = nil
		newRoot.Left = root
	} else {
		newRoot.Left = root.Left
		root.Left = nil
		newRoot.Right = root
	}
	return newRoot
}

func (node *TreeNode) Search(val int) *TreeNode {
	return splay(node, val)
}