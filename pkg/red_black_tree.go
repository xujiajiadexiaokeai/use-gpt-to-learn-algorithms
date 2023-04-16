package main

type Color int

const (
	Red   Color = 0
	Black Color = 1
)

type Node struct {
	Key    int
	Val    string
	Left   *Node
	Right  *Node
	Color  Color
	Parent *Node
}

type RBTree struct {
	Root *Node
}

func (t *RBTree) Insert(key int, val string) {
	node := &Node{Key: key, Val: val}
	t.rootInsert(t.Root, node)
}

func (t *RBTree) rootInsert(root *Node, node *Node) {
	if t.Root == nil {
		t.Root = node
		t.Root.Color = Black
		return
	}
	if node.Key < root.Key {
		if root.Left == nil {
			root.Left = node
			node.Parent = root
			t.insertFixUp(node)
		} else {
			t.rootInsert(root.Left, node)
		}
	} else {
		if root.Right == nil {
			root.Right = node
			node.Parent = root
			t.insertFixUp(node)
		} else {
			t.rootInsert(root.Right, node)
		}
	}
}

// insertFixUp fix the red-black tree after insert
func (t *RBTree) insertFixUp(node *Node) {
	for node.Parent != nil && node.Parent.Color == Red {
		if node.Parent == node.Parent.Parent.Left {
			uncle := node.Parent.Parent.Right
			if uncle != nil && uncle.Color == Red {
				node.Parent.Color = Black
				uncle.Color = Black
				node.Parent.Parent.Color = Red
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Right {
					node = node.Parent
					t.rotateLeft(node)
				}
				node.Parent.Color = Black
				node.Parent.Parent.Color = Red
				t.rotateRight(node.Parent.Parent)
			}
		} else {
			uncle := node.Parent.Parent.Left
			if uncle != nil && uncle.Color == Red {
				node.Parent.Color = Black
				uncle.Color = Black
				node.Parent.Parent.Color = Red
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Left {
					node = node.Parent
					t.rotateRight(node)
				}
				node.Parent.Color = Black
				node.Parent.Parent.Color = Red
				t.rotateLeft(node.Parent.Parent)
			}
		}
	}
	t.Root.Color = Black
}

func (t *RBTree) Delete(key int) {
	node := t.search(key)
	if node == nil {
		return
	}
	t.rootDelete(t.Root, node)
}

func (t *RBTree) rootDelete(root *Node, node *Node) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		if node.Parent != nil {
			if node.Parent.Left == node {
				node.Parent.Left = nil
			} else {
				node.Parent.Right = nil
			}
		} else {
			t.Root = nil
		}
	} else if node.Left == nil {
		t.Transplant(node, node.Right)
	} else if node.Right == nil {
		t.Transplant(node, node.Left)
	} else {
		min := t.min(node.Right)
		if min.Parent != node {
			t.Transplant(min, min.Right)
			min.Right = node.Right
			min.Right.Parent = min
		}
		t.Transplant(node, min)
		min.Left = node.Left
		min.Left.Parent = min
	}
}

func (t *RBTree) Search(key int) *Node {
	return t.search(key)
}

func (t *RBTree) search(key int) *Node {
	return t.rootSearch(t.Root, key)
}

func (t *RBTree) rootSearch(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	if root.Key == key {
		return root
	}
	if root.Key > key {
		return t.rootSearch(root.Left, key)
	}
	return t.rootSearch(root.Right, key)
}

func (t *RBTree) Min() *Node {
	return t.min(t.Root)
}

func (t *RBTree) min(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root
	}
	return t.min(root.Left)
}

func (t *RBTree) Max() *Node {
	return t.max(t.Root)
}

func (t *RBTree) max(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Right == nil {
		return root
	}
	return t.max(root.Right)
}

// Transplant replace the subtree rooted at u with the subtree rooted at v
func (t *RBTree) Transplant(u *Node, v *Node) {
	if u.Parent != nil {
		if u.Parent.Left == u {
			u.Parent.Left = v
		} else {
			u.Parent.Right = v
		}
	} else {
		t.Root = v
	}
}

func (t *RBTree) rotateLeft(node *Node) *Node {
	right := node.Right
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Parent = node.Parent
	if node.Parent != nil {
		if node.Parent.Left == node {
			node.Parent.Left = right
		} else {
			node.Parent.Right = right
		}
	} else {
		t.Root = right
	}
	return right
}

func (t *RBTree) rotateRight(node *Node) *Node {
	left := node.Left
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Parent = node.Parent
	if node.Parent != nil {
		if node.Parent.Left == node {
			node.Parent.Left = left
		} else {
			node.Parent.Right = left
		}
	} else {
		t.Root = left
	}
	return left
}

func NewRBTree() *RBTree {
	return &RBTree{}
}
