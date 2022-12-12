package internal

import "fmt"

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

type BinaryTree struct {
	Root *BinaryNode
}

var size int64

func (t *BinaryTree) Insert(data int64) *BinaryTree {
	if t.Root == nil {
		t.Root = &BinaryNode{data: data, left: nil, right: nil}
		size++
	} else {
		t.Root.insert(data)
		size++
	}
	return t
}

func (n *BinaryNode) insert(data int64) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func (t *BinaryTree) Delete(data int64) *BinaryTree {
	if t.Root != nil {
		t.Root.delete(data)
		size--
	} else {
		t.Root = nil
		size = 0
	}
	return t
}

func (n *BinaryNode) delete(data int64) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			fmt.Println("No Data Found")
		} else if n.left.data == data {

		} else {
			n.left.delete(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func (n *BinaryTree) Size() int64 {
	return size
}
