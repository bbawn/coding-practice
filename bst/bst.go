// Binary search tree coding exercise. See Cormen et al Introduction to
// Algorithms, Ch 12

package bst

const (
	invalid = -1
)

// Node is a node in a binary search tree
type Node struct {
	parent      *Node
	left, right *Node
	val         int
}

// BST is a binary search tree
type BST struct {
	// root is a sentinel value with an invalid value. This simplifies
	// implementation because every actual node in the tree has a parent.
	root Node
}

// New creates and returns an empty tree
func New() *BST {
	return &BST{root: Node{val: invalid}}
}

// Walk traverses the tree in order, calling f for each node
func (t *BST) Walk(f func(val int)) {
	t.root.left.walk(f)
}

// Empty returns true for a tree with no elements, false otherwise
func (t *BST) Empty() bool {
	return t.root.left == nil
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
	if t.Empty() {
		return invalid
	}
	return t.root.left.min().val
}

func (n *Node) min() *Node {
	if n.left == nil {
		return n
	}
	return n.left.min()
}

// Max returns the maximum value in the tree or -1 for an empty tree
func (t *BST) Max() int {
	if t.Empty() {
		return invalid
	}
	return t.root.left.max().val
}

func (n *Node) max() *Node {
	if n.right == nil {
		return n
	}
	return n.right.max()
}

// Search returns a node in the tree with the given value or nil if not found
func (t *BST) Search(val int) *Node {
	return t.root.left.search(val)
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
	if n == nil {
		return nil
	}
	if n.right != nil {
		return n.right.min()
	}

	// Walk upward until path traverses a left edge or top reached
	var child *Node
	for child = n; child.parent.val != invalid && child != child.parent.left; child = child.parent {
	}

	if child.parent.val == invalid {
		return nil
	}

	return child.parent
}

func (n *Node) Predecessor() *Node {
	if n == nil {
		return nil
	}
	if n.left != nil {
		return n.left.max()
	}

	// Walk upward until path traverses a right edge or top reached
	var child *Node
	for child = n; child.parent.val != invalid && child != child.parent.right; child = child.parent {
	}

	if child.parent.val == invalid {
		return nil
	}

	return child.parent
}

// Insert adds a node with the given value to the tree if it's not already
// present. Returns the node with the value or nil for invalid value.
func (t *BST) Insert(val int) *Node {
	if !valid(val) {
		return nil
	}
	if t.Empty() {
		t.root.left = &Node{val: val, parent: &t.root}
		return t.root.left
	}

	n := t.root.left
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

func valid(val int) bool {
	return val >= 0
}

// Delete removes the node with the given value (if present) from the tree and
// returns it. Returns nil if tree has no such node.
func (t *BST) Delete(val int) *Node {
	n := t.Search(val)
	if n == nil {
		return nil
	}
	p := n.parent
	var child **Node
	if n == p.left {
		child = &p.left
	} else {
		child = &p.right
	}

	// Deleted node has no children
	if n.left == nil && n.right == nil {
		*child = nil
		return n
	}

	// Deleted node only has a left child
	if n.right == nil {
		n.left.parent = p
		*child = n.left
		return n
	}

	// Deleted node only has a right child
	if n.left == nil {
		n.right.parent = p
		*child = n.right
		return n
	}

	// Deleted node has two children
	n.left.parent = p
	*child = n.left

	// Add right to max leaf in left subtree
	max := n.left.max()
	max.right = n.right

	return n
}
