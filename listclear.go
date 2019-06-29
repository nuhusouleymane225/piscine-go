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

/*
func listPushBack(l *List, data interface{}) {
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
*/
func ListLast(l *List) *List {

	for l.Head != nil {
		if l.Head.Next == nil {
			return l
		}
		l.Head = l.Head.Next
	}
	return l
}

/*
//converts interface to int
func convertIntToInterface(t []int) []interface{} {
	RandLen := z01.RandIntBetween(0, len(t))
	s := make([]interface{}, RandLen)
	for j := 0; j < RandLen; j++ {
		for i := 0; i < z01.RandIntBetween(1, len(t)); i++ {
			s[j] = t[i]
		}
	}
	return s
}
//converts interface to string
func convertIntToStringface(t []string) []interface{} {
	RandLen := z01.RandIntBetween(0, len(t))
	s := make([]interface{}, RandLen)
	for j := 0; j < RandLen; j++ {
		for i := 0; i < z01.RandIntBetween(1, len(t)); i++ {
			s[j] = t[i]
		}
	}
	return s
}
func main() {
	link := &list{}
	type nodeTest struct {
		Data []interface{}
	}
	table := []nodeTest{}
	//in case the lists is empty
	table = append(table,
		nodeTest{
			data: []interface{}{},
		},
	)
	for i := 0; i < 15; i++ {
		val := nodeTest{
			data: convertIntToInterface(z01.MultRandInt()),
		}
		table = append(table, val)
	}
	for i := 0; i < 15; i++ {
		val := nodeTest{
			data: convertIntToStringface(z01.MultRandWords()),
		}
		table = append(table, val)
	}
	table = append(table,
		nodeTest{
			data: []interface{}{3, 2, 1},
		},
	)
	for _, arg := range table {
		for i := 0; i < len(arg.data); i++ {
			listPushBack(link, arg.data[i])
		}
		m := link.head
		for m != nil {
			fmt.Print(m.data, " -> ")
			m = m.next
		}
		fmt.Print(nil)
		fmt.Println()
		fmt.Println()
		fmt.Println(listLast(link).head)
	}
}
*/