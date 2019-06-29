package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n:=&NodeL{Data:data}
	if l.Head==nil{
		l.Head=n
	}else{
		elemt:=l.Head
		for elemt.Next!=nil{
			elemt=elemt.Next
		}
		elemt.Next=n
	}
}


