package list

import "testing"

func TestEmpty(t *testing.T) {
	ll := Constructor()
	val := ll.Get(0)
	if val != -1 {
		t.Errorf("Get(0) = %d; want -1", val)
	}
}

func TestSimple(t *testing.T) {
	ll := Constructor()
	check(t, &ll)

	ll.AddAtHead(42)
	check(t, &ll, 42)

	ll = Constructor()
	check(t, &ll)

	ll.AddAtTail(42)
	check(t, &ll, 42)

	ll.AddAtHead(11)
	check(t, &ll, 11, 42)

	ll.AddAtTail(76)
	check(t, &ll, 11, 42, 76)

	ll.AddAtTail(234)
	check(t, &ll, 11, 42, 76, 234)

	ll.AddAtHead(13)
	check(t, &ll, 13, 11, 42, 76, 234)

	ll.AddAtTail(234)
	check(t, &ll, 13, 11, 42, 76, 234, 234)

	ll.AddAtIndex(1, 23)
	check(t, &ll, 13, 23, 11, 42, 76, 234, 234)

	ll.AddAtIndex(7, 5)
	check(t, &ll, 13, 23, 11, 42, 76, 234, 234, 5)

	ll.AddAtIndex(9, 99)
	check(t, &ll, 13, 23, 11, 42, 76, 234, 234, 5)

	ll.AddAtIndex(-1, 44)
	check(t, &ll, 13, 23, 11, 42, 76, 234, 234, 5)

	ll.DeleteAtIndex(-1)
	check(t, &ll, 13, 23, 11, 42, 76, 234, 234, 5)

	ll.DeleteAtIndex(9)
	check(t, &ll, 13, 23, 11, 42, 76, 234, 234, 5)

	ll.DeleteAtIndex(0)
	check(t, &ll, 23, 11, 42, 76, 234, 234, 5)

	ll.DeleteAtIndex(1)
	check(t, &ll, 23, 42, 76, 234, 234, 5)

	ll.DeleteAtIndex(5)
	check(t, &ll, 23, 42, 76, 234, 234)
}

func check(t *testing.T, ll *LinkedList, elts ...int) {
	t.Helper()
	for i := 0; i < len(elts); i++ {
		val := ll.Get(i)
		if val != elts[i] {
			t.Errorf("Get(%d) = %d; want %d", i, val, elts[i])
		}
	}
	val := ll.Get(len(elts))
	if val != -1 {
		t.Errorf("Get(%d) = %d; want -1", len(elts), val)
	}
}
