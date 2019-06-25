package piscine

func BasicJoin(strs []string) string{
	result:=""
	for _,value:=range strs{
		result=concat(result,value)
	}
	return result	
}

func concat(str1,str2 string) string{
	result:=make([]byte,len(str1)+len(str2))
	j:=0
	for i,_:=range result{
		if(i<len(str1)){
			result[i]=str1[i]
		}else{
			result[i]=str2[j]
			j++
		}
		
	}
	return string(result)
}

