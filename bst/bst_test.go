package bst

import (
	"reflect"
	"testing"
)

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
	if node == nil {
		t.Errorf("Search(11) expected non-nil ")
	} else if node.val != 11 {
		t.Errorf("Search(11) expected val 11, got %d", node.val)
	}

	for _, v := range []int{3, 14, 55} {
		node = bst.Search(v)
		if node != nil {
			t.Errorf("Search(%d) expected nil", v)
		}
	}

	for _, v := range [][]int{{4, 7}, {9, 11}, {11, 13}} {
		node = bst.Search(v[0])
		node = node.Successor()
		if node == nil {
			t.Errorf("Successor of %d expected %d got nil", v[0], v[1])
		} else if node.val != v[1] {
			t.Errorf("Successor of %d expected %d got %d", v[0], v[1], node.val)
		}
	}

	node = bst.Search(32)
	node = node.Successor()
	if node != nil {
		t.Errorf("Successor of 32 expected nil got %d", node.val)
	}

	for _, v := range [][]int{{22, 17}, {11, 9}, {28, 22}} {
		node = bst.Search(v[0])
		node = node.Predecessor()
		if node == nil {
			t.Errorf("Predecessor of %d expected %d got nil", v[0], v[1])
		} else if node.val != v[1] {
			t.Errorf("Predecessor of %d expected %d got %d", v[0], v[1], node.val)
		}
	}

	node = bst.Search(4)
	node = node.Predecessor()
	if node != nil {
		t.Errorf("Predecessor of 4 expected nil got %d", node.val)
	}
}

func createSimple() BST {
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

	bst := BST{root: n22}
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

func TestMutate(t *testing.T) {
	bst := BST{}

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
