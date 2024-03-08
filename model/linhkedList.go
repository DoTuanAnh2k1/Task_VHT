package model

type LinkedList struct {
	head *Item
	tail *Item
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		size: 0,
	}
}

func (ll *LinkedList) PopBack() *Item {
	if ll.size == 0 {
		return nil
	}

	if ll.size == 1 {
		result := ll.head
		ll.head.Next = nil
		ll.head.Prev = nil
		ll.head = nil
		ll.tail = nil
		ll.size--

		return result
	}
	result := ll.tail
	ll.tail = ll.tail.Prev

	result.Prev = nil
	result.Next = nil
	ll.size--

	return result
}

func (ll *LinkedList) PushBack(item *Item) {
	// newNode := &Node{Element: item, Next: nil, Prev: nil}

	if ll.size == 0 {
		ll.head = item
		ll.tail = item
		item.Prev = nil
		item.Next = nil
		ll.size++
		return
	}

	ll.tail.Next = item
	item.Prev = ll.tail
	item.Next = nil
	ll.tail = item
	ll.size++
}

func (ll *LinkedList) Len() int {
	return ll.size
}
