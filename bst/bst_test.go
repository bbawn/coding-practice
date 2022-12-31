package bst

import (
	"reflect"
	"testing"
)

func createSimple() *BST {
	n4 := &Node{val: 4}
	n7 := &Node{val: 7}
	n9 := &Node{val: 9}
	n11 := &Node{val: 11}
	n13 := &Node{val: 13}
	n15 := &Node{val: 15}
	n17 := &Node{val: 17}
	n22 := &Node{val: 22}
	n28 := &Node{val: 28}
	n30 := &Node{val: 30}
	n32 := &Node{val: 32}

	bst := New()
	bst.root.left = n22
	n22.parent = &bst.root
	n22.left = n9
	n9.parent = n22
	n22.right = n30
	n30.parent = n22
	n9.left = n4
	n4.parent = n9
	n9.right = n13
	n13.parent = n9
	n4.right = n7
	n7.parent = n4
	n13.left = n11
	n11.parent = n13
	n13.right = n17
	n17.parent = n13
	n17.left = n15
	n15.parent = n17
	n30.left = n28
	n28.parent = n30
	n30.right = n32
	n32.parent = n30

	return bst
}

func (n *Node) check(t *testing.T, exp *Node) {
	t.Helper()
	if exp == nil {
		if n != nil {
			t.Errorf("check() expected nil, got %v", n.val)
		}
		return
	}

	if n == nil {
		t.Errorf("check() expected %d, got nil", exp.val)
	}

	if n.val != exp.val {
		t.Errorf("check() expected %d, got %d", exp.val, n.val)
	}
}

func (bst *BST) check(t *testing.T, expected []int) {
	t.Helper()
	got := []int{}
	bst.Walk(func(val int) {
		got = append(got, val)
	})
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("check() expected %v, got %v", expected, got)
	}
}

func TestSimple(t *testing.T) {
	bst := createSimple()
	expected := []int{4, 7, 9, 11, 13, 15, 17, 22, 28, 30, 32}
	got := []int{}
	bst.Walk(func(val int) {
		got = append(got, val)
	})
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("after Walk() expected %v, got %v", expected, got)
	}

	min := bst.Min()
	if min != 4 {
		t.Errorf("Min() expected 4, got %v", min)
	}

	max := bst.Max()
	if max != 32 {
		t.Errorf("Min() expected 32, got %v", max)
	}

	node := bst.Search(11)
	node.check(t, &Node{val: 11})

	for _, v := range []int{3, 14, 55} {
		node = bst.Search(v)
		if node != nil {
			t.Errorf("Search(%d) expected nil", v)
		}
	}

	for _, v := range [][]int{{4, 7}, {9, 11}, {11, 13}} {
		node = bst.Search(v[0])
		node = node.Successor()
		node.check(t, &Node{val: v[1]})
	}

	node = bst.Search(32)
	node = node.Successor()
	if node != nil {
		t.Errorf("Successor of 32 expected nil got %d", node.val)
	}

	for _, v := range [][]int{{22, 17}, {11, 9}, {28, 22}} {
		node = bst.Search(v[0])
		node = node.Predecessor()
		node.check(t, &Node{val: v[1]})
	}

	node = bst.Search(4)
	node = node.Predecessor()
	if node != nil {
		t.Errorf("Predecessor of 4 expected nil got %d", node.val)
	}
}

func TestMutate(t *testing.T) {
	bst := New()

	table := []struct {
		new        int
		pred, succ *Node
	}{
		{7, nil, nil},
		{7, nil, nil},
		{27, &Node{val: 7}, nil},
		{17, &Node{val: 7}, &Node{val: 27}},
		{37, &Node{val: 27}, nil},
		{37, &Node{val: 27}, nil},
		{4, nil, &Node{val: 7}},
		{5, &Node{val: 4}, &Node{val: 7}},
		{12, &Node{val: 7}, &Node{val: 17}},
	}
	for _, item := range table {
		n := bst.Insert(item.new)
		if n == nil {
			t.Errorf("Insert(%d), expected val %d got nil", item.new, item.new)
		}
		s := n.Successor()
		if !match(item.succ, s) {
			t.Errorf("Insert(%d), Successor expected %v got %v", item.new, item.succ, n)
		}
		p := n.Predecessor()
		if !match(item.pred, p) {
			t.Errorf("Insert(%d), Predecessor expected %v got %v", item.new, item.pred, n)
		}
	}

	// Insert invalid value
	n := bst.Insert(-1)
	if n != nil {
		t.Errorf("Insert(-1), expected nil got val %d", n.val)
	}

	// Delete node with only left child
	n = bst.Delete(17)
	n.check(t, &Node{val: 17})
	bst.check(t, []int{4, 5, 7, 12, 27, 37})

	// Delete node with only right child
	n = bst.Delete(4)
	n.check(t, &Node{val: 4})
	bst.check(t, []int{5, 7, 12, 27, 37})

	// Delete node with no children
	n = bst.Delete(12)
	n.check(t, &Node{val: 12})
	bst.check(t, []int{5, 7, 27, 37})

	// Delete node with two children
	n = bst.Delete(27)
	n.check(t, &Node{val: 27})
	bst.check(t, []int{5, 7, 37})

	// Delete top node with two children
	n = bst.Delete(7)
	n.check(t, &Node{val: 7})
	bst.check(t, []int{5, 37})

	// Delete invalid value
	n = bst.Delete(-1)
	n.check(t, nil)
	bst.check(t, []int{5, 37})

	// Delete value not in tree
	n = bst.Delete(15)
	n.check(t, nil)
	bst.check(t, []int{5, 37})
}

func match(n1, n2 *Node) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if n1 != nil && n2 != nil && n1.val == n2.val {
		return true
	}

	return false
}
