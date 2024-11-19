package datastructures

type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

// AddToFront добавляет элемент в начало списка.
func (l *LinkedList) AddToFront(node *Node) {
	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}
	node.Next = l.Head
	l.Head.Prev = node
	l.Head = node
}

// RemoveNode удаляет конкретный узел.
func (l *LinkedList) RemoveNode(node *Node) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		l.Head = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		l.Tail = node.Prev
	}
}

// RemoveTail удаляет последний элемент.
func (l *LinkedList) RemoveTail() *Node {
	if l.Tail == nil {
		return nil
	}
	tail := l.Tail
	l.RemoveNode(tail)
	return tail
}
