package model

type Node struct {
	Element *Item
	Next    *Node
	Prev    *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: &Node{
			Element: &Item{},
			Next:    nil,
			Prev:    nil,
		},
		tail: &Node{
			Element: &Item{},
			Next:    nil,
			Prev:    nil,
		},
		size: 0,
	}
}

func (ll *LinkedList) PopBack() *Item {
	if ll.size == 0 {
		return nil
	}

	lastNode := ll.tail.Prev
	ll.removeNode(lastNode)
	return lastNode.Element
}

func (ll *LinkedList) PushBack(item *Item) {
	newNode := &Node{Element: item, Next: nil, Prev: nil}

	if ll.size == 0 {
		ll.head.Next = newNode
		newNode.Prev = ll.head
	} else {
		lastNode := ll.tail.Prev
		lastNode.Next = newNode
		newNode.Prev = lastNode
	}

	ll.tail.Prev = newNode
	newNode.Next = ll.tail
	ll.size++
}

func (ll *LinkedList) Len() int {
	return ll.size
}

func (ll *LinkedList) removeNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	ll.size--
}
