package piscine

func MakeRange(min, max int) []int{
	if max<=min{
		return nil
	}
	result:=make([]int,max-min)
	j:=0
	for i:=min;i< max;i++{
		result[j]=i
		j++
	}
	return result
}
