package piscine

func ListSize(l *List) int {
	if l.Head!=nil{
		count:=1
		cl:=l.Head
		for cl.Next!=nil{
			cl=cl.Next
			count++
		}
		return count
	}	
	return 0
}


