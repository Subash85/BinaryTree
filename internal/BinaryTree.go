package BinaryTree

import "errors"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func New(data int) *Node {
	return &Node{
		Left:  nil,
		Right: nil,
		Data:  data,
	}
}

func (n *Node) Insert(node *Node, data int) (*Node, error) {
	var err error = nil
	if node == nil {
		return New(data), nil
	}
	switch {
	case data == node.Data:
		return nil, errors.New("Duplicate error")
	case data < node.Data:
		node.Left, err = node.setValue(node.Left, data)
	case data > n.Data:
		node.Right, err = node.setValue(node.Right, data)
	}
	if err != nil {
		return nil, err
	}

	return node, nil
}
func (n *Node) setValue(node *Node, data int) (*Node, error) {
	if node == nil {
		node = &Node{Data: data}
		return node, nil
	}
	return node.Insert(node, data)
}

func (n *Node) Find(node *Node, val int) (*Node, bool) {
	if node == nil {
		return nil, false
	}

	if node.Data == val {
		return node, true
	}
	if val < node.Data {
		return node.Left.Find(node.Left, val)
	} else if val > node.Data {
		return node.Right.Find(node.Right, val)
	}
	return nil, false
}

func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n == nil {
		return nil, parent
	}
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(n)
}

func (n *Node) Delete(s int, parent *Node) error {
	if n == nil {
		return errors.New("Value to be deleted does not exist in the tree")
	}

	switch {
	case s < n.Data:
		return n.Left.Delete(s, n)
	case s > n.Data:
		return n.Right.Delete(s, n)
	default:

		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
			return nil
		}

		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}
		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return nil
		}

		replacement, replParent := n.Left.findMax(n)

		n.Data = replacement.Data

		return replacement.Delete(replacement.Data, replParent)
	}
}

func (n *Node) replaceNode(parent, replacement *Node) error {
	if n == nil {
		return errors.New("replaceNode() not allowed on a nil node")
	}

	if n == parent.Left {
		parent.Left = replacement
		return nil
	}
	parent.Right = replacement
	return nil
}
