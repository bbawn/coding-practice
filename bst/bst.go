// Binary search tree coding exercise. See Cormen et al Introduction to
// Algorithms, Ch 12

package bst

// Node is a node in a binary search tree
type Node struct {
	parent      *Node
	left, right *Node
	val         int
}

// BST is a binary search tree
type BST struct {
	root *Node
}

// Walk traverses the tree in order, calling f for each node
func (t *BST) Walk(f func(val int)) {
	t.root.walk(f)
}

func (n *Node) walk(f func(val int)) {
	if n == nil {
		return
	}
	n.left.walk(f)
	f(n.val)
	n.right.walk(f)
}

// Min returns the minimum value in the tree or -1 for an empty tree
func (t *BST) Min() int {
	if t.root == nil {
		return -1
	}
	return t.root.min().val
}

func (n *Node) min() *Node {
	if n.left == nil {
		return n
	}
	return n.left.min()
}

// Max returns the maximum value in the tree or -1 for an empty tree
func (t *BST) Max() int {
	if t.root == nil {
		return -1
	}
	return t.root.max().val
}

func (n *Node) max() *Node {
	if n.right == nil {
		return n
	}
	return n.right.max()
}

// Search returns a node in the tree with the given value or nil if not found
func (t *BST) Search(val int) *Node {
	return t.root.search(val)
}

func (n *Node) search(val int) *Node {
	if n == nil {
		return nil
	}

	if n.val == val {
		return n
	}

	match := n.left.search(val)
	if match != nil {
		return match
	}

	return n.right.search(val)
}

func (n *Node) Successor() *Node {
	if n.right != nil {
		return n.right.min()
	}

	// Walk upward until path traverses a left edge or top reached
	var child *Node
	for child = n; child.parent != nil && child != child.parent.left; child = child.parent {
	}

	return child.parent
}

func (n *Node) Predecessor() *Node {
	if n.left != nil {
		return n.left.max()
	}

	// Walk upward until path traverses a right edge or top reached
	var child *Node
	for child = n; child.parent != nil && child != child.parent.right; child = child.parent {
	}

	return child.parent
}

// Insert adds a node with the given value to the tree if it's not already
// present. Returns the node with the value.
func (t *BST) Insert(val int) *Node {
	if t.root == nil {
		t.root = &Node{val: val}
		return t.root
	}

	n := t.root
	for {
		if val == n.val {
			return n
		} else if val < n.val {
			if n.left == nil {
				n.left = &Node{val: val, parent: n}
				return n.left
			}
			n = n.left
		} else {
			if n.right == nil {
				n.right = &Node{val: val, parent: n}
				return n.right
			}
			n = n.right
		}
	}
}

// Delete removes the node with the given value (if present) from the tree and
// returns it. Returns nil if tree has no such node.
func (t *BST) Delete(val int) *Node {
	n := t.root.search(val)
	if n == nil {
		return nil
	}
	return n
}
