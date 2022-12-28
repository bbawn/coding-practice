package list

// linked-list coding exercise. See https://leetcode.com/problems/design-linked-list/

type Element struct {
	Val  int
	Next *Element
}

type LinkedList struct {
	Head *Element
}

func Constructor() LinkedList {
	return LinkedList{}
}

func (ll *LinkedList) Get(index int) int {
	e := ll.Head

	for i := 0; i < index && e != nil; i++ {
		e = e.Next
	}

	if e == nil {
		return -1
	}

	return e.Val
}

func (ll *LinkedList) AddAtHead(val int) {
	h := ll.Head
	ll.Head = &Element{val, h}
}

func (ll *LinkedList) AddAtTail(val int) {
	if ll.Head == nil {
		ll.Head = &Element{val, nil}
		return
	}

	var t *Element
	for t = ll.Head; t.Next != nil; t = t.Next {
	}
	t.Next = &Element{val, nil}
}

func (ll *LinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		return
	}

	if index == 0 {
		ll.AddAtHead(val)
		return
	}

	pred := ll.Head
	for i := 1; i < index; i++ {
		pred = pred.Next
		if pred == nil {
			return
		}
	}
	pred.Next = &Element{val, pred.Next}
}

func (ll *LinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}

	if index == 0 && ll.Head != nil {
		ll.Head = ll.Head.Next
		return
	}

	pred := ll.Head
	for i := 1; i < index; i++ {
		pred = pred.Next
		if pred.Next == nil {
			return
		}
	}
	pred.Next = pred.Next.Next
}
