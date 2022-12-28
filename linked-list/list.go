package list

// linked-list coding exercise. See https://leetcode.com/problems/design-linked-list/

type element struct {
	val  int
	next *element
}

type LinkedList struct {
	// root is a sentinel element - it's next element is the start of the list
	root element
}

func Constructor() LinkedList {
	return LinkedList{}
}

func (ll *LinkedList) Get(index int) int {
	e := ll.root.next

	for i := 0; i < index && e != nil; i++ {
		e = e.next
	}

	if e == nil {
		return -1
	}

	return e.val
}

func (ll *LinkedList) AddAtHead(val int) {
	h := ll.root.next
	ll.root.next = &element{val, h}
}

func (ll *LinkedList) AddAtTail(val int) {
	var t *element
	for t = &ll.root; t.next != nil; t = t.next {
	}
	t.next = &element{val, nil}
}

func (ll *LinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		return
	}

	pred := &ll.root
	for i := 0; i < index; i++ {
		pred = pred.next
		if pred == nil {
			return
		}
	}
	pred.next = &element{val, pred.next}
}

func (ll *LinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}

	pred := &ll.root
	for i := 0; i < index; i++ {
		pred = pred.next
		if pred.next == nil {
			return
		}
	}
	pred.next = pred.next.next
}
