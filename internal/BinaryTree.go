package BinaryTree

import "errors"

type Tree struct {
	Root *Node
}

type Node struct {
	Data  string
	Left  *Node
	Right *Node
}

func (t *Tree) Insert(data string) error {

	if t.Root == nil {
		t.Root = &Node{Data: data}
		return nil
	}
	return t.Root.Insert(data)
}

func (n *Node) Insert(data string) error {

	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}
	switch {

	case data == n.Data:
		return nil
	case data < n.Data:
		n.setValue(n.Left, data)
	case data > n.Data:
		n.setValue(n.Right, data)
	}
	return nil
}
func (n *Node) setValue(node *Node, data string) error {
	if node == nil {
		node = &Node{Data: data}
		return nil
	}
	return node.Insert(data)
}

func (t *Tree) Find(s string) (string, bool) {
	if t.Root == nil {
		return "", false
	}
	return t.Root.Find(s)
}

func (n *Node) Find(s string) (string, bool) {

	if n == nil {
		return "", false
	}

	switch {
	case s == n.Data:
		return n.Data, true
	case s < n.Data:
		return n.Left.Find(s)
	default:
		return n.Right.Find(s)
	}
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

func (n *Node) Delete(s string, parent *Node) error {
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
