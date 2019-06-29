package piscine

//node struct
type Node struct {
	Data interface{}
	Next *Node
}

type List struct {
	Head *Node
	Tail *Node
}

//main function
func ListPushBack(l *List, data interface{}) {

	n := &Node{}
	n.Data = data
	n.Next = nil

	if l.Head == nil {
		l.Head = n
		return
	}

	iterator := l.Head
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
}

//cleat List
func ListClear(l *List) {

	temp := l.Head
	Next := l.Head
	for temp != nil {
		Next = temp.Next
		temp = Next
	}
	l.Head = nil
}
